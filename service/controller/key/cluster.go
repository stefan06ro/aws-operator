package key

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strconv"
	"strings"
	"time"

	infrastructurev1alpha3 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
	"github.com/giantswarm/microerror"

	"github.com/giantswarm/aws-operator/pkg/project"
)

const (
	// CloudConfigVersion defines the version of k8scloudconfig in use. It is used
	// in the main stack output and S3 object paths.
	CloudConfigVersion = "v_6_1_0"
	CloudProvider      = "aws"
)

const (
	EC2RoleK8s   = "EC2-K8S-Role"
	EC2PolicyK8s = "EC2-K8S-Policy"
)

const (
	EtcdPort                     = 2379
	EtcdPrefix                   = "giantswarm.io"
	KubernetesSecurePort         = 443
	KubernetesApiHealthCheckPort = 8089
)

const (
	HAMasterSnapshotIDValue = "ha-master-migration"
)

const (
	// KubernetesAPIHealthzVersion is a tag representing the version of
	// https://github.com/giantswarm/k8s-api-healthz/ used.
	KubernetesAPIHealthzVersion = "0.1.1"
	// K8sSetupNetworkEnvironment is a tag representing the version of
	// https://github.com/giantswarm/k8s-setup-network-environment used.
	K8sSetupNetworkEnvironment = "0.2.0"
)

// AWS Tags used for cost analysis and general resource tagging.
const (
	TagAvailabilityZone  = "giantswarm.io/availability-zone"
	TagCluster           = "giantswarm.io/cluster"
	TagClusterType       = "giantswarm.io/cluster-type"
	TagControlPlane      = "giantswarm.io/control-plane"
	TagInstallation      = "giantswarm.io/installation"
	TagMachineDeployment = "giantswarm.io/machine-deployment"
	TagName              = "Name"
	TagOrganization      = "giantswarm.io/organization"
	TagRouteTableType    = "giantswarm.io/route-table-type"
	TagStack             = "giantswarm.io/stack"
	TagSnapshot          = "giantswarm.io/snapshot"
	TagSubnetType        = "giantswarm.io/subnet-type"
)

const (
	StackTCCP  = "tccp"
	StackTCCPF = "tccpf"
	StackTCCPI = "tccpi"
	StackTCCPN = "tccpn"
	StackTCNP  = "tcnp"
	StackTCNPF = "tcnpf"
)

const (
	LifeCycleHookControlPlane = "ControlPlane"
	LifeCycleHookNodePool     = "NodePool"
)

const (
	RefWorkerASG = "workerAutoScalingGroup"
)

const (
	// ComponentOS is the name of the component specified in a Release CR which
	// determines the version of the OS to be used for tenant cluster nodes and
	// is ultimately transformed into an AMI based on TC region.
	ComponentOS = "containerlinux"
)

func ClusterAPIEndpoint(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("api.%s", TenantClusterBaseDomain(cluster))
}

func ClusterBaseDomain(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Cluster.DNS.Domain
}

func ClusterEtcdEndpoint(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("etcd.%s", TenantClusterBaseDomain(cluster))
}

func ClusterEtcdEndpointWithPort(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s:2379", ClusterEtcdEndpoint(cluster))
}

func ClusterKubeletEndpoint(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("worker.%s", TenantClusterBaseDomain(cluster))
}

func ClusterNamespace(cluster infrastructurev1alpha3.AWSCluster) string {
	return ClusterID(&cluster)
}

func CredentialName(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.CredentialSecret.Name
}

func CredentialNamespace(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.CredentialSecret.Namespace
}

func ExternalSNAT(cluster infrastructurev1alpha3.AWSCluster) *bool {
	return cluster.Spec.Provider.Pods.ExternalSNAT
}

func PodsCIDRBlock(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.Pods.CIDRBlock
}

func IsChinaRegion(awsRegion string) bool {
	return strings.HasPrefix(awsRegion, "cn-")
}

func IsNewCluster(cluster infrastructurev1alpha3.AWSCluster) bool {
	// If  condition list is empty then this is a new cluster.
	if len(cluster.Status.Cluster.Conditions) == 0 {
		return true
	}
	// If versions list is empty then this is a new cluster.
	if len(cluster.Status.Cluster.Versions) == 0 {
		return true
	}
	// Check if there is transition state Updated or Updating,
	// this indicates cluster is not new but updated from other version.
	if cluster.Status.Cluster.HasUpdatedCondition() || cluster.Status.Cluster.HasUpdatingCondition() {
		return false
	}

	// Check if there is transition state from other version,
	// this indicates cluster is not new but updated from other version.
	for _, version := range cluster.Status.Cluster.Versions {
		if version.Version != project.Version() {
			return false
		}
	}

	// No update event is registered in the versions or conditions lists in cr status so this is new cluster.
	return true
}

func IsAlreadyCreatedCluster(cluster infrastructurev1alpha3.AWSCluster) bool {
	// if cluster has Created status it has been already provisioned
	return cluster.Status.Cluster.HasCreatedCondition()
}

func MasterAvailabilityZone(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.Master.AvailabilityZone
}

func MasterCount(cluster infrastructurev1alpha3.AWSCluster) int {
	return 1
}

func MasterInstanceResourceName(cr infrastructurev1alpha3.AWSCluster, t time.Time) string {
	return getResourcenameWithTimeHash("MasterInstance", cr, t)
}

func MasterInstanceName(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-master", ClusterID(&cluster))
}

func MasterInstanceType(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.Master.InstanceType
}

func ManagedRecordSets(cluster infrastructurev1alpha3.AWSCluster) []string {
	tcBaseDomain := TenantClusterBaseDomain(cluster)
	return []string{
		fmt.Sprintf("%s.", tcBaseDomain),
		fmt.Sprintf("\\052.%s.", tcBaseDomain), // \\052 - `*` wildcard record
		fmt.Sprintf("api.%s.", tcBaseDomain),
		fmt.Sprintf("etcd.%s.", tcBaseDomain),
		fmt.Sprintf("internal-api.%s.", tcBaseDomain),
	}
}

func OIDCClientID(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Cluster.OIDC.ClientID
}
func OIDCIssuerURL(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Cluster.OIDC.IssuerURL
}
func OIDCUsernameClaim(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Cluster.OIDC.Claims.Username
}
func OIDCGroupsClaim(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Cluster.OIDC.Claims.Groups
}

func PolicyNameMaster(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-master-%s", ClusterID(&cluster), EC2PolicyK8s)
}

func ProfileNameMaster(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-master-%s", ClusterID(&cluster), EC2RoleK8s)
}

func Region(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Spec.Provider.Region
}

func RoleNameMaster(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-master-%s", ClusterID(&cluster), EC2RoleK8s)
}

func RolePeerAccess(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-vpc-peer-access", ClusterID(&cluster))
}

func RouteTableName(cluster infrastructurev1alpha3.AWSCluster, suffix, az string) string {
	return fmt.Sprintf("%s-%s-%s", ClusterID(&cluster), suffix, az)
}

func StatusClusterNetworkCIDR(cluster infrastructurev1alpha3.AWSCluster) string {
	return cluster.Status.Provider.Network.CIDR
}

func TenantClusterBaseDomain(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s.k8s.%s", ClusterID(&cluster), ClusterBaseDomain(cluster))
}

func ToCluster(ctx context.Context, v interface{}) (infrastructurev1alpha3.AWSCluster, error) {
	if v == nil {
		return infrastructurev1alpha3.AWSCluster{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &infrastructurev1alpha3.AWSCluster{}, v)
	}

	p, ok := v.(*infrastructurev1alpha3.AWSCluster)
	if !ok {
		return infrastructurev1alpha3.AWSCluster{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &infrastructurev1alpha3.AWSCluster{}, v)
	}

	c := p.DeepCopy()

	return *c, nil
}

func VolumeNameDocker(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-docker", ClusterID(&cluster))
}

func VolumeNameEtcd(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-etcd", ClusterID(&cluster))
}

func VolumeNameLog(cluster infrastructurev1alpha3.AWSCluster) string {
	return fmt.Sprintf("%s-log", ClusterID(&cluster))
}

func ensureLabel(labels string, key string, value string) string {
	if key == "" {
		return labels
	}
	if value == "" {
		return labels
	}

	var split []string
	if labels != "" {
		split = strings.Split(labels, ",")
	}

	var found bool
	for i, l := range split {
		if !strings.HasPrefix(l, key+"=") {
			continue
		}

		found = true
		split[i] = key + "=" + value
	}

	if !found {
		split = append(split, key+"="+value)
	}

	joined := strings.Join(split, ",")

	return joined
}

// getResourcenameWithTimeHash returns a string cromprised of some prefix, a
// time hash and a cluster ID.
func getResourcenameWithTimeHash(prefix string, cluster infrastructurev1alpha3.AWSCluster, t time.Time) string {
	id := strings.Replace(ClusterID(&cluster), "-", "", -1)

	h := sha512.New()
	_, err := h.Write([]byte(strconv.FormatInt(t.UnixNano(), 10)))
	if err != nil {
		panic(microerror.JSON(err))
	}
	timeHash := fmt.Sprintf("%x", h.Sum(nil))[0:5]

	upperTimeHash := strings.ToUpper(timeHash)
	upperClusterID := strings.ToUpper(id)

	return fmt.Sprintf("%s%s%s", prefix, upperClusterID, upperTimeHash)
}
