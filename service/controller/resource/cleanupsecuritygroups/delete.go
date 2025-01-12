package cleanupsecuritygroups

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/operatorkit/v5/pkg/controller/context/finalizerskeptcontext"

	"github.com/giantswarm/aws-operator/service/controller/controllercontext"
	"github.com/giantswarm/aws-operator/service/controller/key"
)

func (r *Resource) EnsureDeleted(ctx context.Context, obj interface{}) error {
	cr, err := key.ToCluster(ctx, obj)
	if err != nil {
		return microerror.Mask(err)
	}
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	var groups []*ec2.SecurityGroup
	{
		r.logger.Debugf(ctx, "finding security groups for tenant cluster %#q", key.ClusterID(&cr))

		i := &ec2.DescribeSecurityGroupsInput{
			Filters: []*ec2.Filter{
				// The filter here matches all security groups managed by ingress
				// controllers running within the tenant cluster. The cluster deletion
				// process might not be perfect and the ingress controllers ran by
				// customers might not cleanup properly when the whole tenant cluster is
				// torn down.
				{
					Name: aws.String(fmt.Sprintf("tag:kubernetes.io/cluster/%s", key.ClusterID(&cr))),
					Values: []*string{
						aws.String("owned"),
					},
				},
			},
		}

		o, err := cc.Client.TenantCluster.AWS.EC2.DescribeSecurityGroups(i)
		if err != nil {
			return microerror.Mask(err)
		}

		groups = o.SecurityGroups

		r.logger.Debugf(ctx, "found security groups for tenant cluster %#q", key.ClusterID(&cr))
	}

	if len(groups) > 0 {
		r.logger.Debugf(ctx, "deleting %d security groups for tenant cluster %#q", len(groups), key.ClusterID(&cr))

		var deleted int
		for _, g := range groups {
			r.logger.Debugf(ctx, "deleting security group %#q for tenant cluster %#q", *g.GroupId, key.ClusterID(&cr))

			i := &ec2.DeleteSecurityGroupInput{
				GroupId: g.GroupId,
			}

			_, err := cc.Client.TenantCluster.AWS.EC2.DeleteSecurityGroup(i)
			if IsDependencyViolation(err) {
				r.logger.Debugf(ctx, "security group %#q for tenant cluster %#q still has dependency", *g.GroupId, key.ClusterID(&cr))
				r.logger.Debugf(ctx, "skipping security group %#q for tenant cluster %#q", *g.GroupId, key.ClusterID(&cr))

				r.logger.Debugf(ctx, "keeping finalizers")
				finalizerskeptcontext.SetKept(ctx)

				continue

			} else if err != nil {
				return microerror.Mask(err)
			}

			deleted++

			r.logger.Debugf(ctx, "deleted security group %#q for tenant cluster %#q", *g.GroupId, key.ClusterID(&cr))
		}

		r.logger.Debugf(ctx, "deleted %d security groups for tenant cluster %#q", deleted, key.ClusterID(&cr))
	} else {
		r.logger.Debugf(ctx, "no security groups to be deleted for tenant cluster %#q", key.ClusterID(&cr))
	}

	return nil
}
