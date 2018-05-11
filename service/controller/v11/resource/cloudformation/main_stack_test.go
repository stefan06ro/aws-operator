package cloudformation

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	"github.com/giantswarm/micrologger/microloggertest"

	"github.com/giantswarm/aws-operator/service/controller/v11/adapter"
	"github.com/giantswarm/aws-operator/service/controller/v11/cloudconfig"
	cloudformationservice "github.com/giantswarm/aws-operator/service/controller/v11/cloudformation"
	"github.com/giantswarm/aws-operator/service/controller/v11/key"
)

func testConfig() Config {
	c := Config{}

	c.Clients = &adapter.Clients{}
	c.EBS = &EBSServiceMock{}
	c.HostClients = &adapter.Clients{}
	c.Logger = microloggertest.New()
	c.Service = &cloudformationservice.CloudFormation{}

	c.InstallationName = "myinstallation"

	return c
}

func TestMainGuestTemplateGetEmptyBody(t *testing.T) {
	t.Parallel()
	customObject := v1alpha1.AWSConfig{}

	c := testConfig()
	c.Clients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
		KMS: &adapter.KMSClientMock{},
	}
	c.HostClients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
	}

	newResource, err := New(c)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	_, err = newResource.getMainGuestTemplateBody(customObject, StackState{})
	if err == nil {
		t.Fatal("expected", nil, "got", err)
	}
}

func TestMainGuestTemplateExistingFields(t *testing.T) {
	t.Parallel()
	// customObject with example fields for both asg and launch config
	customObject := v1alpha1.AWSConfig{
		Spec: v1alpha1.AWSConfigSpec{
			Cluster: v1alpha1.Cluster{
				ID:      "test-cluster",
				Version: "myversion",
				Kubernetes: v1alpha1.ClusterKubernetes{
					API: v1alpha1.ClusterKubernetesAPI{
						Domain:     "api.domain",
						SecurePort: 443,
					},
					IngressController: v1alpha1.ClusterKubernetesIngressController{
						Domain:       "ingress.domain",
						InsecurePort: 30010,
						SecurePort:   30011,
					},
				},
				Etcd: v1alpha1.ClusterEtcd{
					Domain: "etcd.domain",
				},
			},
			AWS: v1alpha1.AWSConfigSpecAWS{
				API: v1alpha1.AWSConfigSpecAWSAPI{
					ELB: v1alpha1.AWSConfigSpecAWSAPIELB{
						IdleTimeoutSeconds: 3600,
					},
				},
				Region: "eu-central-1",
				AZ:     "eu-central-1a",
				Masters: []v1alpha1.AWSConfigSpecAWSNode{
					{
						ImageID:      "ami-1234-master",
						InstanceType: "m3.large",
					},
				},
				Workers: []v1alpha1.AWSConfigSpecAWSNode{
					{
						ImageID:      "ami-1234-worker",
						InstanceType: "m3.large",
					},
				},
				Ingress: v1alpha1.AWSConfigSpecAWSIngress{
					ELB: v1alpha1.AWSConfigSpecAWSIngressELB{
						IdleTimeoutSeconds: 60,
					},
				},
				VPC: v1alpha1.AWSConfigSpecAWSVPC{
					CIDR:              "10.1.1.0/24",
					PublicSubnetCIDR:  "10.1.1.0/25",
					PrivateSubnetCIDR: "10.1.2.0/25",
				},
			},
		},
	}

	imageID, err := key.ImageID(customObject)
	if err != nil {
		t.Fatalf("expected %#v got %#v", nil, err)
	}

	stackState := StackState{
		Name: key.MainGuestStackName(customObject),

		MasterImageID:              imageID,
		MasterInstanceResourceName: key.MasterInstanceResourceName(customObject),
		MasterInstanceType:         key.MasterInstanceType(customObject),
		MasterCloudConfigVersion:   cloudconfig.MasterCloudConfigVersion,
		MasterInstanceMonitoring:   false,

		WorkerCount:              strconv.Itoa(key.WorkerCount(customObject)),
		WorkerImageID:            imageID,
		WorkerInstanceMonitoring: true,
		WorkerInstanceType:       key.WorkerInstanceType(customObject),
		WorkerCloudConfigVersion: cloudconfig.WorkerCloudConfigVersion,

		VersionBundleVersion: key.VersionBundleVersion(customObject),
	}

	cfg := testConfig()
	cfg.Clients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
		KMS: &adapter.KMSClientMock{},
		ELB: &adapter.ELBClientMock{},
	}
	cfg.HostClients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
	}
	cfg.AdvancedMonitoringEC2 = true
	newResource, err := New(cfg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	body, err := newResource.getMainGuestTemplateBody(customObject, stackState)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !strings.Contains(body, "Description: Main Guest CloudFormation stack.") {
		t.Fatal("stack header not found")
	}

	if !strings.Contains(body, "  workerLaunchConfiguration:") {
		t.Fatal("launch configuration header not found")
	}

	if !strings.Contains(body, "  workerAutoScalingGroup:") {
		t.Fatal("asg header not found")
	}

	if !strings.Contains(body, "InstanceType: m3.large") {
		t.Fatal("launch configuration element not found")
	}

	if !strings.Contains(body, "AvailabilityZones: [eu-central-1a]") {
		fmt.Println(body)
		t.Fatal("asg element not found")
	}

	if !strings.Contains(body, "Outputs:") {
		fmt.Println(body)
		t.Fatal("outputs header not found")
	}

	if !strings.Contains(body, key.MasterImageIDKey+":") {
		fmt.Println(body)
		t.Fatal("MasterImageID output element not found")
	}
	if !strings.Contains(body, key.MasterInstanceMonitoring+": true") {
		fmt.Println(body)
		t.Fatal("MasterInstanceMonitoring output element not found")
	}
	if !strings.Contains(body, key.MasterInstanceTypeKey+":") {
		fmt.Println(body)
		t.Fatal("MasterInstanceType output element not found")
	}
	if !strings.Contains(body, key.MasterCloudConfigVersionKey+":") {
		fmt.Println(body)
		t.Fatal("master CloudConfig version output element not found")
	}
	if !strings.Contains(body, key.WorkerCountKey+":") {
		fmt.Println(body)
		t.Fatal("workers output element not found")
	}
	if !strings.Contains(body, key.WorkerImageIDKey+":") {
		fmt.Println(body)
		t.Fatal("WorkerImageID output element not found")
	}
	if !strings.Contains(body, key.WorkerInstanceTypeKey+":") {
		fmt.Println(body)
		t.Fatal("WorkerInstanceType output element not found")
	}
	if !strings.Contains(body, key.WorkerCloudConfigVersionKey+":") {
		fmt.Println(body)
		t.Fatal("worker CloudConfig version output element not found")
	}
	if !strings.Contains(body, key.WorkerInstanceMonitoring+": true") {
		fmt.Println(body)
		t.Fatal("WorkerInstanceMonitoring output element not found")
	}

	if !strings.Contains(body, "Value: "+cloudconfig.MasterCloudConfigVersion) {
		fmt.Println(body)
		t.Fatal("output element not found")
	}

	if !strings.Contains(body, workerRoleKey+":") {
		fmt.Println(body)
		t.Fatal("workerRole output element not found")
	}
	if !strings.Contains(body, "PolicyName: test-cluster-master") {
		fmt.Println(body)
		t.Fatal("PolicyName output element not found")
	}
	if !strings.Contains(body, "PolicyName: test-cluster-worker") {
		fmt.Println(body)
		t.Fatal("PolicyName output element not found")
	}
	if !strings.Contains(body, "ApiRecordSet:") {
		fmt.Println(body)
		t.Fatal("ApiRecordSet element not found")
	}
	if !strings.Contains(body, "EtcdRecordSet:") {
		fmt.Println(body)
		t.Fatal("EtcdRecordSet element not found")
	}
	if !strings.Contains(body, "IngressRecordSet:") {
		fmt.Println(body)
		t.Fatal("IngressRecordSet element not found")
	}
	if !strings.Contains(body, "IngressWildcardRecordSet:") {
		fmt.Println(body)
		t.Fatal("ingressWildcardRecordSet element not found")
	}
	if !strings.Contains(body, "MasterInstance") {
		fmt.Println(body)
		t.Fatal("MasterInstance element not found")
	}
	if !strings.Contains(body, "ApiLoadBalancer:") {
		fmt.Println(body)
		t.Fatal("ApiLoadBalancer element not found")
	}
	if !strings.Contains(body, "IngressLoadBalancer:") {
		fmt.Println(body)
		t.Fatal("IngressLoadBalancer element not found")
	}
	if !strings.Contains(body, "InternetGateway:") {
		fmt.Println(body)
		t.Fatal("InternetGateway element not found")
	}
	if !strings.Contains(body, "NATGateway:") {
		fmt.Println(body)
		t.Fatal("NATGateway element not found")
	}
	if !strings.Contains(body, "PublicRouteTable:") {
		fmt.Println(body)
		t.Fatal("PublicRouteTable element not found")
	}
	if !strings.Contains(body, "PublicSubnet:") {
		fmt.Println(body)
		t.Fatal("PublicSubnet element not found")
	}
	if !strings.Contains(body, "PrivateRouteTable:") {
		fmt.Println(body)
		t.Fatal("PrivateRouteTable element not found")
	}
	if !strings.Contains(body, "PrivateSubnet:") {
		fmt.Println(body)
		t.Fatal("PrivateSubnet element not found")
	}
	if !strings.Contains(body, "MasterSecurityGroup:") {
		fmt.Println(body)
		t.Fatal("MasterSecurityGroup element not found")
	}
	if !strings.Contains(body, "WorkerSecurityGroup:") {
		fmt.Println(body)
		t.Fatal("WorkerSecurityGroup element not found")
	}
	if !strings.Contains(body, "IngressSecurityGroup:") {
		fmt.Println(body)
		t.Fatal("IngressSecurityGroup element not found")
	}
	if !strings.Contains(body, " VPC:") {
		fmt.Println(body)
		t.Fatal("VPC element not found")
	}
	if !strings.Contains(body, "CidrBlock: 10.1.1.0/24") {
		fmt.Println(body)
		t.Fatal("CidrBlock element not found")
	}

	// image ids should be fixed despite the values in the custom object
	if !strings.Contains(body, "ImageId: ami-604e118b") {
		fmt.Println(body)
		t.Fatal("Fixed image ID not found")
	}
}

func TestMainHostPreTemplateExistingFields(t *testing.T) {
	t.Parallel()
	// customObject with example fields for both asg and launch config
	customObject := v1alpha1.AWSConfig{
		Spec: v1alpha1.AWSConfigSpec{
			Cluster: v1alpha1.Cluster{
				ID: "test-cluster",
			},
		},
	}

	cfg := testConfig()
	cfg.Clients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
	}
	cfg.HostClients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
		IAM: &adapter.IAMClientMock{},
	}
	newResource, err := New(cfg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	body, err := newResource.getMainHostPreTemplateBody(customObject)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !strings.Contains(body, "Description: Main Host Pre-Guest CloudFormation stack.") {
		fmt.Println(body)
		t.Fatal("stack header not found")
	}

	if !strings.Contains(body, "  PeerRole:") {
		fmt.Println(body)
		t.Fatal("peer role header not found")
	}

	if !strings.Contains(body, "  RoleName: test-cluster-vpc-peer-access") {
		fmt.Println(body)
		t.Fatal("role name not found")
	}
}

func TestMainHostPostTemplateExistingFields(t *testing.T) {
	t.Parallel()
	// customObject with example fields for both asg and launch config
	customObject := v1alpha1.AWSConfig{
		Spec: v1alpha1.AWSConfigSpec{
			Cluster: v1alpha1.Cluster{
				ID: "test-cluster",
			},
			AWS: v1alpha1.AWSConfigSpecAWS{
				VPC: v1alpha1.AWSConfigSpecAWSVPC{
					RouteTableNames: []string{
						"route_table_1",
						"route_table_2",
					},
					PeerID: "mypeerid",
				},
			},
		},
	}

	cfg := testConfig()
	cfg.Clients = &adapter.Clients{
		EC2: &adapter.EC2ClientMock{},
	}
	ec2Mock := &adapter.EC2ClientMock{}
	ec2Mock.SetMatchingRouteTables(1)
	cfg.HostClients = &adapter.Clients{
		EC2: ec2Mock,
		IAM: &adapter.IAMClientMock{},
	}
	newResource, err := New(cfg)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	body, err := newResource.getMainHostPostTemplateBody(customObject)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !strings.Contains(body, "Description: Main Host Post-Guest CloudFormation stack.") {
		fmt.Println(body)
		t.Fatal("stack header not found")
	}

	if !strings.Contains(body, "  PrivateRoute1:") {
		fmt.Println(body)
		t.Fatal("route header not found")
	}
}
