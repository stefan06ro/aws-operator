package tcnp

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
	infrastructurev1alpha3 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/microerror"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/aws-operator/pkg/awstags"
	"github.com/giantswarm/aws-operator/pkg/label"
	"github.com/giantswarm/aws-operator/service/controller/controllercontext"
	"github.com/giantswarm/aws-operator/service/controller/key"
	"github.com/giantswarm/aws-operator/service/controller/resource/tccpnoutputs"
	"github.com/giantswarm/aws-operator/service/controller/resource/tcnp/template"
	cloudformationutils "github.com/giantswarm/aws-operator/service/internal/cloudformation"
	"github.com/giantswarm/aws-operator/service/internal/encrypter/kms"
)

const (
	capabilityNamesIAM = "CAPABILITY_NAMED_IAM"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cr, err := key.ToMachineDeployment(obj)
	if err != nil {
		return microerror.Mask(err)
	}
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	// Ensure some preconditions are met so we have all necessary information
	_, err = r.encrypter.EncryptionKey(ctx, key.ClusterID(&cr))
	if kms.IsKeyNotFound(err) {
		r.logger.Debugf(ctx, "canceling resource", "reason", "encryption key not available yet")
		return nil

	} else if err != nil {
		return microerror.Mask(err)
	}

	// Ensure some preconditions are met so we have all neccessary information
	// available to manage the TCNP CF stack.
	{
		if !cc.Status.TenantCluster.S3Object.Uploaded {
			r.logger.Debugf(ctx, "s3 object not available yet")
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		}

		if len(cc.Spec.TenantCluster.TCNP.AvailabilityZones) == 0 {
			r.logger.Debugf(ctx, "availability zone information not available yet")
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		}

		if len(cc.Status.TenantCluster.TCCP.AvailabilityZones) == 0 {
			r.logger.Debugf(ctx, "availability zone information not available yet")
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		}

		if len(cc.Status.TenantCluster.TCCP.SecurityGroups) == 0 {
			r.logger.Debugf(ctx, "security group information not available yet")
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		}

		if cc.Status.TenantCluster.TCCP.VPC.PeeringConnectionID == "" {
			r.logger.Debugf(ctx, "vpc peering connection id not available yet")
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		}
	}

	{
		r.logger.Debugf(ctx, "finding the tenant cluster's node pool cloud formation stack")

		i := &cloudformation.DescribeStacksInput{
			StackName: aws.String(key.StackNameTCNP(&cr)),
		}

		o, err := cc.Client.TenantCluster.AWS.CloudFormation.DescribeStacks(i)
		if IsNotExists(err) {
			r.logger.Debugf(ctx, "did not find the tenant cluster's node pool cloud formation stack")

			err = r.createStack(ctx, cr)
			if err != nil {
				return microerror.Mask(err)
			}

			return nil

		} else if err != nil {
			return microerror.Mask(err)

		} else if len(o.Stacks) != 1 {
			return microerror.Maskf(executionFailedError, "expected one stack, got %d", len(o.Stacks))

		} else if *o.Stacks[0].StackStatus == cloudformation.StackStatusCreateFailed {
			return microerror.Maskf(eventCFCreateError, "expected successful status, got %#q", *o.Stacks[0].StackStatus)
		} else if *o.Stacks[0].StackStatus == cloudformation.StackStatusRollbackFailed {
			return microerror.Maskf(eventCFRollbackError, "expected successful status, got %#q", *o.Stacks[0].StackStatus)
		} else if *o.Stacks[0].StackStatus == cloudformation.StackStatusUpdateRollbackFailed {
			return microerror.Maskf(eventCFUpdateRollbackError, "expected successful status, got %#q", *o.Stacks[0].StackStatus)

		} else if key.StackInProgress(*o.Stacks[0].StackStatus) {
			r.logger.Debugf(ctx, "the tenant cluster's node pool cloud formation stack has stack status %#q", *o.Stacks[0].StackStatus)
			r.event.Emit(ctx, &cr, "CFInProgress", fmt.Sprintf("the tenant cluster's node pool cloud formation stack has stack status %#q", *o.Stacks[0].StackStatus))
			r.logger.Debugf(ctx, "canceling resource")
			return nil
		} else if key.StackComplete(*o.Stacks[0].StackStatus) {
			r.event.Emit(ctx, &cr, "CFCompleted", fmt.Sprintf("the tenant cluster's control plane cloud formation stack has stack status %#q", *o.Stacks[0].StackStatus))
		}

		r.logger.Debugf(ctx, "found the tenant cluster's node pool cloud formation stack already exists")
	}

	{
		scale, err := r.detection.ShouldScale(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}
		update, err := r.detection.ShouldUpdate(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}

		if scale {
			err = r.updateStack(ctx, cr)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		if update {
			// only allow tcnp CF stack update when a tccpn CF stack has finished updating
			tccpnUpdated, err := isTCCPNUpdated(ctx, cr)
			if IsTccpnNotUpdated(err) {
				r.logger.Debugf(ctx, "waiting for tccpn stack to finish an update before executing an tcnp update")

				return nil
			} else if err != nil {
				return microerror.Mask(err)
			}

			if tccpnUpdated {
				err = r.updateStack(ctx, cr)
				if err != nil {
					return microerror.Mask(err)
				}
			}
		}
	}

	return nil
}

func (r *Resource) createStack(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	var templateBody string
	{
		r.logger.Debugf(ctx, "computing the template of the tenant cluster's node pool cloud formation stack")

		params, err := r.newTemplateParams(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}

		templateBody, err = template.Render(params)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "computed the template of the tenant cluster's node pool cloud formation stack")
	}

	{
		r.logger.Debugf(ctx, "requesting the creation of the tenant cluster's node pool cloud formation stack")

		tags, err := r.getCloudFormationTags(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}

		i := &cloudformation.CreateStackInput{
			Capabilities: []*string{
				aws.String(capabilityNamesIAM),
			},
			EnableTerminationProtection: aws.Bool(true),
			StackName:                   aws.String(key.StackNameTCNP(&cr)),
			Tags:                        tags,
			TemplateBody:                aws.String(templateBody),
		}

		_, err = cc.Client.TenantCluster.AWS.CloudFormation.CreateStack(i)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "requested the creation of the tenant cluster's node pool cloud formation stack")
		r.event.Emit(ctx, &cr, "CFCreateRequested", "requested the creation of the tenant cluster's node pool cloud formation stack")
	}

	return nil
}

func (r *Resource) getCloudFormationTags(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) ([]*cloudformation.Tag, error) {
	tags := key.AWSTags(&cr, r.installationName)
	tags[key.TagStack] = key.StackTCNP
	tags[key.TagMachineDeployment] = key.MachineDeploymentID(&cr)

	cloudtags, err := r.cloudtags.GetTagsByCluster(ctx, key.ClusterID(&cr))
	if err != nil {
		return nil, microerror.Mask(err)
	}
	for k, v := range cloudtags {
		tags[k] = v
	}

	return awstags.NewCloudFormation(tags), nil
}

func (r *Resource) updateStack(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	var templateBody string
	{
		r.logger.Debugf(ctx, "computing the template of the tenant cluster's node pool cloud formation stack")

		params, err := r.newTemplateParams(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}

		templateBody, err = template.Render(params)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "computed the template of the tenant cluster's node pool cloud formation stack")
	}

	{
		r.logger.Debugf(ctx, "requesting the update of the tenant cluster's node pool cloud formation stack")

		tags, err := r.getCloudFormationTags(ctx, cr)
		if err != nil {
			return microerror.Mask(err)
		}

		i := &cloudformation.UpdateStackInput{
			Capabilities: []*string{
				aws.String(capabilityNamesIAM),
			},
			StackName:    aws.String(key.StackNameTCNP(&cr)),
			Tags:         tags,
			TemplateBody: aws.String(templateBody),
		}
		_, err = cc.Client.TenantCluster.AWS.CloudFormation.UpdateStack(i)
		if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "requested the update of the tenant cluster's node pool cloud formation stack")
		r.event.Emit(ctx, &cr, "CFUpdateRequested", "requested the update of the tenant cluster's node pool cloud formation stack")
	}

	return nil
}

// minDesiredWorkers calculates appropriate minimum value to be set for ASG
// Desired value and to be used for computation of workerCountRatio.
//
// When cluster-autoscaler has scaled cluster and ASG's Desired value is higher
// than minimum number of instances allowed for that ASG, then it makes sense to
// consider Desired value as minimum number of running instances for further
// operational computations.
//
// Example:
// Initially ASG has minimum of 3 workers and maximum of 10. Due to amount of
// workload deployed on workers, cluster-autoscaler has scaled current Desired
// number of instances to 5. Therefore it makes sense to consider 5 as minimum
// number of nodes also when working on batch updates on ASG instances.
//
// Example 2:
// When end user is scaling cluster and adding restrictions to its size, it
// might be that initial ASG configuration is following:
// 		- Min: 3
//		- Max: 10
// 		- Desired: 10
//
// Now end user decides that it must be scaled down so maximum size is decreased
// to 7. When desired number of instances is temporarily bigger than maximum
// number of instances, it must be fixed to be maximum number of instances.
//
func minDesiredWorkers(minWorkers, maxWorkers, statusDesiredCapacity int) int {
	if statusDesiredCapacity > maxWorkers {
		return maxWorkers
	}

	if statusDesiredCapacity > minWorkers {
		return statusDesiredCapacity
	}

	return minWorkers
}

func (r *Resource) newAutoScalingGroup(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainAutoScalingGroup, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	var cl infrastructurev1alpha3.AWSCluster
	{
		var list infrastructurev1alpha3.AWSClusterList
		err := r.k8sClient.CtrlClient().List(
			ctx,
			&list,
			client.InNamespace(cr.Namespace),
			client.MatchingLabels{label.Cluster: key.ClusterID(&cr)},
		)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if len(list.Items) != 1 {
			return nil, microerror.Maskf(executionFailedError, "expected 1 CR got %d", len(list.Items))
		}

		cl = list.Items[0]
	}

	var subnets []string
	for _, az := range cc.Spec.TenantCluster.TCNP.AvailabilityZones {
		subnets = append(subnets, key.SanitizeCFResourceName(key.PrivateSubnetName(az.Name)))
	}

	minDesiredNodes := minDesiredWorkers(key.MachineDeploymentScalingMin(cr), key.MachineDeploymentScalingMax(cr), cc.Status.TenantCluster.ASG.DesiredCapacity)

	var launchTemplateOverride []template.LaunchTemplateOverride
	{
		val, ok := r.alikeInstances[key.MachineDeploymentInstanceType(cr)]

		if cr.Spec.Provider.Worker.UseAlikeInstanceTypes && ok {
			launchTemplateOverride = val
		}
	}

	var maxBatchSize string
	var minInstancesInService string
	{
		// try read the value from cluster CR
		if val, ok := cl.Annotations[annotation.AWSUpdateMaxBatchSize]; ok {
			maxBatchSize = key.MachineDeploymentParseMaxBatchSize(val, minDesiredNodes)

			r.logger.Debugf(ctx, "value of MaxBatchSize for ASG updates set by annotation from AWSCluster CR")
		}
		// override the value with machine deployment value if its set
		if val, ok := cr.Annotations[annotation.AWSUpdateMaxBatchSize]; ok {
			maxBatchSize = key.MachineDeploymentParseMaxBatchSize(val, minDesiredNodes)

			r.logger.Debugf(ctx, "value of MaxBatchSize for ASG updates overridden by annotation from AWSMachineDeployment CR")
		}
		// if nothing is set use the default
		if maxBatchSize == "" {
			maxBatchSize = key.MachineDeploymentWorkerCountRatio(minDesiredNodes, 0.3)
		}
		// set minInstancesInService based on the maxBatchSize value
		minInstancesInService, err = key.MachineDeploymentMinInstanceInServiceFromMaxBatchSize(maxBatchSize, minDesiredNodes)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var pauseTime string
	{
		// try read the value from cluster CR
		if val, ok := cl.Annotations[annotation.AWSUpdatePauseTime]; ok {
			if key.MachineDeploymentPauseTimeIsValid(val) {
				pauseTime = val

				r.logger.Debugf(ctx, "value of PauseTime for ASG updates set by annotation from AWSCLuster CR")
			}
		}
		// override the value with machine deployment value if its set
		if val, ok := cr.Annotations[annotation.AWSUpdatePauseTime]; ok {
			if key.MachineDeploymentPauseTimeIsValid(val) {
				pauseTime = val

				r.logger.Debugf(ctx, "value of PauseTime for ASG updates overridden by annotation from AWSMachineDeployment CR")
			}
		}
		// if nothing is set use the default
		if pauseTime == "" {
			pauseTime = key.DefaultPauseTimeBetweenUpdates
		}
	}

	autoScalingGroup := &template.ParamsMainAutoScalingGroup{
		AvailabilityZones: key.MachineDeploymentAvailabilityZones(cr),
		Cluster: template.ParamsMainAutoScalingGroupCluster{
			ID: key.ClusterID(&cr),
		},
		DesiredCapacity:                     minDesiredNodes,
		MaxBatchSize:                        maxBatchSize,
		MaxSize:                             key.MachineDeploymentScalingMax(cr),
		MinInstancesInService:               minInstancesInService,
		MinSize:                             key.MachineDeploymentScalingMin(cr),
		Subnets:                             subnets,
		PauseTime:                           pauseTime,
		OnDemandPercentageAboveBaseCapacity: key.MachineDeploymentOnDemandPercentageAboveBaseCapacity(cr),
		OnDemandBaseCapacity:                key.MachineDeploymentOnDemandBaseCapacity(cr),
		SpotInstancePools:                   key.MachineDeploymentSpotInstancePools(launchTemplateOverride),
		SpotAllocationStrategy:              "lowest-price",
		LaunchTemplateOverrides:             launchTemplateOverride,
		LifeCycleHookName:                   key.LifeCycleHookNodePool,
	}

	return autoScalingGroup, nil
}

func (r *Resource) newIAMPolicies(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainIAMPolicies, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	ek, err := r.encrypter.EncryptionKey(ctx, key.ClusterID(&cr))
	if kms.IsKeyNotFound(err) {
		r.logger.Debugf(ctx, "canceling resource", "reason", "encryption key not available yet")
		return nil, nil

	} else if err != nil {
		return nil, microerror.Mask(err)
	}

	var iamPolicies *template.ParamsMainIAMPolicies
	{
		iamPolicies = &template.ParamsMainIAMPolicies{
			Cluster: template.ParamsMainIAMPoliciesCluster{
				ID: key.ClusterID(&cr),
			},
			EC2ServiceDomain: key.EC2ServiceDomain(cc.Status.TenantCluster.AWS.Region),
			KMSKeyARN:        ek,
			NodePool: template.ParamsMainIAMPoliciesNodePool{
				ID: key.MachineDeploymentID(&cr),
			},
			RegionARN: key.RegionARN(cc.Status.TenantCluster.AWS.Region),
			S3Bucket:  key.BucketName(&cr, cc.Status.TenantCluster.AWS.AccountID),
		}
	}

	return iamPolicies, nil
}

func (r *Resource) newLaunchTemplate(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainLaunchTemplate, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var ami string
	{
		ami, err = r.images.AMI(ctx, &cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	launchTemplate := &template.ParamsMainLaunchTemplate{
		BlockDeviceMapping: template.ParamsMainLaunchTemplateBlockDeviceMapping{
			Docker: template.ParamsMainLaunchTemplateBlockDeviceMappingDocker{
				Volume: template.ParamsMainLaunchTemplateBlockDeviceMappingDockerVolume{
					Size: key.MachineDeploymentDockerVolumeSizeGB(cr),
				},
			},
			Kubelet: template.ParamsMainLaunchTemplateBlockDeviceMappingKubelet{
				Volume: template.ParamsMainLaunchTemplateBlockDeviceMappingKubeletVolume{
					Size: key.MachineDeploymentKubeletVolumeSizeGB(cr),
				},
			},
			Logging: template.ParamsMainLaunchTemplateBlockDeviceMappingLogging{
				Volume: template.ParamsMainLaunchTemplateBlockDeviceMappingLoggingVolume{
					Size: 100,
				},
			},
		},
		Instance: template.ParamsMainLaunchTemplateInstance{
			Image:      ami,
			Monitoring: true,
			Type:       key.MachineDeploymentInstanceType(cr),
		},
		Metadata: template.ParamsMainLaunchTemplateMetadata{
			HttpTokens: key.MachineDeploymentMetadataV2(cr),
		},
		Name:           key.MachineDeploymentLaunchTemplateName(cr),
		ReleaseVersion: key.ReleaseVersion(&cr),
		SmallCloudConfig: template.ParamsMainLaunchTemplateSmallCloudConfig{
			S3URL: fmt.Sprintf("s3://%s/%s", key.BucketName(&cr, cc.Status.TenantCluster.AWS.AccountID), key.S3ObjectPathTCNP(&cr)),
		},
	}

	return launchTemplate, nil
}

func (r *Resource) newOutputs(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainOutputs, error) {
	var err error

	var ami string
	{
		ami, err = r.images.AMI(ctx, &cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	outputs := &template.ParamsMainOutputs{
		DockerVolumeSizeGB: key.MachineDeploymentDockerVolumeSizeGB(cr),
		Instance: template.ParamsMainOutputsInstance{
			Image: ami,
			Type:  key.MachineDeploymentInstanceType(cr),
		},
		OperatorVersion: key.OperatorVersion(&cr),
		ReleaseVersion:  key.ReleaseVersion(&cr),
	}

	return outputs, nil
}

func (r *Resource) newRouteTables(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainRouteTables, error) {
	var routeTables template.ParamsMainRouteTables

	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	for _, a := range cc.Spec.TenantCluster.TCNP.AvailabilityZones {
		r := template.ParamsMainRouteTablesListItem{
			AvailabilityZone: a.Name,
			ClusterID:        key.ClusterID(&cr),
			NodePoolID:       cr.GetName(),
			Name:             key.SanitizeCFResourceName(key.PrivateRouteTableName(a.Name)),
			Route: template.ParamsMainRouteTablesListItemRoute{
				Name: key.SanitizeCFResourceName(key.NATRouteName(a.Name)),
			},
			TCCP: template.ParamsMainRouteTablesListItemTCCP{
				NATGateway: template.ParamsMainRouteTablesListItemTCCPNATGateway{
					ID: a.NATGateway.ID,
				},
				VPC: template.ParamsMainRouteTablesListItemTCCPVPC{
					ID: cc.Status.TenantCluster.TCCP.VPC.ID,
				},
			},
		}
		routeTables.List = append(routeTables.List, r)
	}

	return &routeTables, nil
}

func (r *Resource) newSecurityGroups(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainSecurityGroups, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var cl infrastructurev1alpha3.AWSCluster
	{
		var list infrastructurev1alpha3.AWSClusterList
		err := r.k8sClient.CtrlClient().List(
			ctx,
			&list,
			client.InNamespace(cr.Namespace),
			client.MatchingLabels{label.Cluster: key.ClusterID(&cr)},
		)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		if len(list.Items) != 1 {
			return nil, microerror.Maskf(executionFailedError, "expected 1 CR got %d", len(list.Items))
		}

		cl = list.Items[0]
	}

	var networkCIDR string
	{
		if cl.Status.Provider.Network.CIDR == "" {
			r.logger.Debugf(ctx, "canceling resource", "reason", "tenant cluster network cidr not set")
			return nil, nil
		}
		networkCIDR = cl.Status.Provider.Network.CIDR
	}

	var nodePools []template.ParamsMainSecurityGroupsTenantClusterNodePool
	for _, ID := range cc.Spec.TenantCluster.TCNP.SecurityGroupIDs {
		np := template.ParamsMainSecurityGroupsTenantClusterNodePool{
			ID:           ID,
			ResourceName: key.SanitizeCFResourceName(ID),
		}

		nodePools = append(nodePools, np)
	}

	securityGroups := &template.ParamsMainSecurityGroups{
		ClusterID: key.ClusterID(&cr),
		ControlPlane: template.ParamsMainSecurityGroupsControlPlane{
			VPC: template.ParamsMainSecurityGroupsControlPlaneVPC{
				CIDR: cc.Status.ControlPlane.VPC.CIDR,
			},
		},
		TenantCluster: template.ParamsMainSecurityGroupsTenantCluster{
			InternalAPI: template.ParamsMainSecurityGroupsTenantClusterInternalAPI{
				ID: idFromGroups(cc.Status.TenantCluster.TCCP.SecurityGroups, key.SecurityGroupName(&cr, "internal-api")),
			},
			Master: template.ParamsMainSecurityGroupsTenantClusterMaster{
				ID: idFromGroups(cc.Status.TenantCluster.TCCP.SecurityGroups, key.SecurityGroupName(&cr, "master")),
			},
			AWSCNI: template.ParamsMainSecurityGroupsTenantClusterAWSCNI{
				ID: idFromGroups(cc.Status.TenantCluster.TCCP.SecurityGroups, key.SecurityGroupName(&cr, "aws-cni")),
			},
			NodePools: nodePools,
			VPC: template.ParamsMainSecurityGroupsTenantClusterVPC{
				ID:   cc.Status.TenantCluster.TCCP.VPC.ID,
				CIDR: networkCIDR,
			},
		},
	}

	return securityGroups, nil
}

func (r *Resource) newSubnets(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainSubnets, error) {
	var subnets template.ParamsMainSubnets

	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	for _, a := range cc.Spec.TenantCluster.TCNP.AvailabilityZones {
		s := template.ParamsMainSubnetsListItem{
			AvailabilityZone: a.Name,
			CIDR:             a.Subnet.Private.CIDR.String(),
			Name:             key.SanitizeCFResourceName(key.PrivateSubnetName(a.Name)),
			RouteTable: template.ParamsMainSubnetsListItemRouteTable{
				Name: key.SanitizeCFResourceName(key.PrivateRouteTableName(a.Name)),
			},
			RouteTableAssociation: template.ParamsMainSubnetsListItemRouteTableAssociation{
				Name: key.SanitizeCFResourceName(key.PrivateSubnetRouteTableAssociationName(a.Name)),
			},
			TCCP: template.ParamsMainSubnetsListItemTCCP{
				VPC: template.ParamsMainSubnetsListItemTCCPVPC{
					ID: cc.Status.TenantCluster.TCCP.VPC.ID,
				},
			},
		}

		subnets.List = append(subnets.List, s)
	}

	return &subnets, nil
}

func (r *Resource) newTemplateParams(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMain, error) {
	var params *template.ParamsMain
	{
		autoScalingGroup, err := r.newAutoScalingGroup(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		iamPolicies, err := r.newIAMPolicies(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		launchTemplate, err := r.newLaunchTemplate(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		outputs, err := r.newOutputs(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		routeTables, err := r.newRouteTables(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		securityGroups, err := r.newSecurityGroups(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		subnets, err := r.newSubnets(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		vpc, err := r.newVPC(ctx, cr)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		params = &template.ParamsMain{
			AutoScalingGroup: autoScalingGroup,
			IAMPolicies:      iamPolicies,
			LaunchTemplate:   launchTemplate,
			Outputs:          outputs,
			RouteTables:      routeTables,
			SecurityGroups:   securityGroups,
			Subnets:          subnets,
			VPC:              vpc,
		}
	}

	return params, nil
}

func (r *Resource) newVPC(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (*template.ParamsMainVPC, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var routeTables []template.ParamsMainVPCRouteTable
	for _, a := range cc.Spec.TenantCluster.TCNP.AvailabilityZones {
		r := template.ParamsMainVPCRouteTable{
			ControlPlane: template.ParamsMainVPCRouteTableControlPlane{
				VPC: template.ParamsMainVPCRouteTableControlPlaneVPC{
					CIDR: cc.Status.ControlPlane.VPC.CIDR,
				},
			},
			Route: template.ParamsMainVPCRouteTableRoute{
				Name: key.SanitizeCFResourceName(key.VPCPeeringRouteName(a.Name)),
			},
			RouteTable: template.ParamsMainVPCRouteTableRouteTable{
				Name: key.SanitizeCFResourceName(key.PrivateRouteTableName(a.Name)),
			},
			TenantCluster: template.ParamsMainVPCRouteTableTenantCluster{
				PeeringConnectionID: cc.Status.TenantCluster.TCCP.VPC.PeeringConnectionID,
			},
		}
		routeTables = append(routeTables, r)
	}

	vpc := &template.ParamsMainVPC{
		Cluster: template.ParamsMainVPCCluster{
			ID: key.ClusterID(&cr),
		},
		Region: template.ParamsMainVPCRegion{
			ARN:  key.RegionARN(cc.Status.TenantCluster.AWS.Region),
			Name: cc.Status.TenantCluster.AWS.Region,
		},
		RouteTables: routeTables,
		TCCP: template.ParamsMainVPCTCCP{
			VPC: template.ParamsMainVPCTCCPVPC{
				ID: cc.Status.TenantCluster.TCCP.VPC.ID,
			},
		},
		TCNP: template.ParamsMainVPCTCNP{
			CIDR: key.MachineDeploymentSubnet(cr),
		},
	}

	return vpc, nil
}

func idFromGroups(groups []*ec2.SecurityGroup, name string) string {
	for _, g := range groups {
		if awstags.ValueForKey(g.Tags, "Name") == name {
			return *g.GroupId
		}
	}

	return ""
}

func isTCCPNUpdated(ctx context.Context, cr infrastructurev1alpha3.AWSMachineDeployment) (bool, error) {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return false, microerror.Mask(err)
	}

	// check if TCCPN CF stack is updated
	{
		var cloudFormation *cloudformationutils.CloudFormation
		{
			c := cloudformationutils.Config{
				Client: cc.Client.TenantCluster.AWS.CloudFormation,
			}

			cloudFormation, err = cloudformationutils.New(c)
			if err != nil {
				return false, microerror.Mask(err)
			}
		}

		o, s, err := cloudFormation.DescribeOutputsAndStatus(key.StackNameTCCPN(&cr))
		if cloudformationutils.IsOutputsNotAccessible(err) {
			// outputsNotAccessible can occur when is CF in updating status
			return false, microerror.Mask(tccpnNotUpdatedError)
		} else if err != nil {
			return false, microerror.Mask(err)
		}

		if s != cloudformation.StackStatusUpdateComplete {
			// when TCCPN stack is updated, only  good status that we want to see is `StackStatusUpdateComplete`
			// anything else indicate either CF stack not updated, update in progress or an error
			return false, microerror.Mask(tccpnNotUpdatedError)
		}

		v, err := cloudFormation.GetOutputValue(o, tccpnoutputs.OperatorVersionKey)
		if err != nil {
			return false, microerror.Mask(err)
		}

		if v != key.OperatorVersion(&cr) {
			// CF output value for Operator version do not match with operator version on CR
			// CF is not yet updated
			return false, microerror.Mask(tccpnNotUpdatedError)
		}
	}

	// check if master nodes in tc k8s api are on the latest operator version
	// the value is stored in node labels
	{
		nodes := &corev1.NodeList{}

		err = cc.Client.TenantCluster.K8s.CtrlClient().List(
			ctx,
			nodes,
			client.MatchingLabels{key.NodeRoleLabel: key.MasterNodeRoleLabel},
		)
		if err != nil {
			return false, microerror.Mask(err)
		}

		// check if all master nodes have proper operator version
		for _, n := range nodes.Items {
			if v, ok := n.GetLabels()[label.OperatorVersion]; ok {
				if v != key.OperatorVersion(&cr) {
					// operator version mismatch, CP is not updated yet
					return false, microerror.Mask(tccpnNotUpdatedError)
				}
			} else {
				// node dont have proper node label, its probably in a update phase
				return false, microerror.Mask(tccpnNotUpdatedError)
			}
		}
	}
	// tccpn CF stack is updated and all master nodes have new operator version
	return true, nil
}
