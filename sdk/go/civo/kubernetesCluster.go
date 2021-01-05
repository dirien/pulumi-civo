// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Then the Kubernetes cluster can be imported using the cluster's `id`, e.g.
//
// ```sh
//  $ pulumi import civo:index/kubernetesCluster:KubernetesCluster my-cluster 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af
// ```
type KubernetesCluster struct {
	pulumi.CustomResourceState

	// The base URL of the API server on the Kubernetes master node.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik
	Applications pulumi.StringPtrOutput `pulumi:"applications"`
	// The date where the Kubernetes cluster was build.
	BuiltAt pulumi.StringOutput `pulumi:"builtAt"`
	// The date where the Kubernetes cluster was create.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The unique dns entry for the cluster in this case point to the master.
	DnsEntry pulumi.StringOutput `pulumi:"dnsEntry"`
	// A unique ID that can be used to identify and reference a Kubernetes cluster.
	InstalledApplications KubernetesClusterInstalledApplicationArrayOutput `pulumi:"installedApplications"`
	// In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported.
	Instances KubernetesClusterInstanceArrayOutput `pulumi:"instances"`
	// A representation of the Kubernetes cluster's kubeconfig in yaml format.
	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
	// The version of k3s to install (The default is currently the latest available).
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The Ip of the Kubernetes master node.
	MasterIp pulumi.StringOutput `pulumi:"masterIp"`
	// A name for the Kubernetes cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of instances to create (The default at the time of writing is 3).
	NumTargetNodes pulumi.IntPtrOutput `pulumi:"numTargetNodes"`
	Ready          pulumi.BoolOutput   `pulumi:"ready"`
	// The status of Kubernetes cluster.
	// * `ready` -If the Kubernetes cluster is ready.
	Status pulumi.StringOutput `pulumi:"status"`
	// A space separated list of tags, to be used freely as required.
	Tags pulumi.StringPtrOutput `pulumi:"tags"`
	// The size of each node (The default is currently g2.small)
	TargetNodesSize pulumi.StringPtrOutput `pulumi:"targetNodesSize"`
}

// NewKubernetesCluster registers a new resource with the given unique name, arguments, and options.
func NewKubernetesCluster(ctx *pulumi.Context,
	name string, args *KubernetesClusterArgs, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	if args == nil {
		args = &KubernetesClusterArgs{}
	}
	var resource KubernetesCluster
	err := ctx.RegisterResource("civo:index/kubernetesCluster:KubernetesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesCluster gets an existing KubernetesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesClusterState, opts ...pulumi.ResourceOption) (*KubernetesCluster, error) {
	var resource KubernetesCluster
	err := ctx.ReadResource("civo:index/kubernetesCluster:KubernetesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesCluster resources.
type kubernetesClusterState struct {
	// The base URL of the API server on the Kubernetes master node.
	ApiEndpoint *string `pulumi:"apiEndpoint"`
	// A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik
	Applications *string `pulumi:"applications"`
	// The date where the Kubernetes cluster was build.
	BuiltAt *string `pulumi:"builtAt"`
	// The date where the Kubernetes cluster was create.
	CreatedAt *string `pulumi:"createdAt"`
	// The unique dns entry for the cluster in this case point to the master.
	DnsEntry *string `pulumi:"dnsEntry"`
	// A unique ID that can be used to identify and reference a Kubernetes cluster.
	InstalledApplications []KubernetesClusterInstalledApplication `pulumi:"installedApplications"`
	// In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported.
	Instances []KubernetesClusterInstance `pulumi:"instances"`
	// A representation of the Kubernetes cluster's kubeconfig in yaml format.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// The version of k3s to install (The default is currently the latest available).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The Ip of the Kubernetes master node.
	MasterIp *string `pulumi:"masterIp"`
	// A name for the Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The number of instances to create (The default at the time of writing is 3).
	NumTargetNodes *int  `pulumi:"numTargetNodes"`
	Ready          *bool `pulumi:"ready"`
	// The status of Kubernetes cluster.
	// * `ready` -If the Kubernetes cluster is ready.
	Status *string `pulumi:"status"`
	// A space separated list of tags, to be used freely as required.
	Tags *string `pulumi:"tags"`
	// The size of each node (The default is currently g2.small)
	TargetNodesSize *string `pulumi:"targetNodesSize"`
}

type KubernetesClusterState struct {
	// The base URL of the API server on the Kubernetes master node.
	ApiEndpoint pulumi.StringPtrInput
	// A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik
	Applications pulumi.StringPtrInput
	// The date where the Kubernetes cluster was build.
	BuiltAt pulumi.StringPtrInput
	// The date where the Kubernetes cluster was create.
	CreatedAt pulumi.StringPtrInput
	// The unique dns entry for the cluster in this case point to the master.
	DnsEntry pulumi.StringPtrInput
	// A unique ID that can be used to identify and reference a Kubernetes cluster.
	InstalledApplications KubernetesClusterInstalledApplicationArrayInput
	// In addition to the arguments provided, these additional attributes about the cluster's default node instance are exported.
	Instances KubernetesClusterInstanceArrayInput
	// A representation of the Kubernetes cluster's kubeconfig in yaml format.
	Kubeconfig pulumi.StringPtrInput
	// The version of k3s to install (The default is currently the latest available).
	KubernetesVersion pulumi.StringPtrInput
	// The Ip of the Kubernetes master node.
	MasterIp pulumi.StringPtrInput
	// A name for the Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The number of instances to create (The default at the time of writing is 3).
	NumTargetNodes pulumi.IntPtrInput
	Ready          pulumi.BoolPtrInput
	// The status of Kubernetes cluster.
	// * `ready` -If the Kubernetes cluster is ready.
	Status pulumi.StringPtrInput
	// A space separated list of tags, to be used freely as required.
	Tags pulumi.StringPtrInput
	// The size of each node (The default is currently g2.small)
	TargetNodesSize pulumi.StringPtrInput
}

func (KubernetesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterState)(nil)).Elem()
}

type kubernetesClusterArgs struct {
	// A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik
	Applications *string `pulumi:"applications"`
	// The version of k3s to install (The default is currently the latest available).
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// A name for the Kubernetes cluster.
	Name *string `pulumi:"name"`
	// The number of instances to create (The default at the time of writing is 3).
	NumTargetNodes *int `pulumi:"numTargetNodes"`
	// A space separated list of tags, to be used freely as required.
	Tags *string `pulumi:"tags"`
	// The size of each node (The default is currently g2.small)
	TargetNodesSize *string `pulumi:"targetNodesSize"`
}

// The set of arguments for constructing a KubernetesCluster resource.
type KubernetesClusterArgs struct {
	// A comma separated list of applications to install. Spaces within application names are fine, but shouldn't be either side of the comma. If you want to remove a default installed application, prefix it with a '-', e.g. -traefik
	Applications pulumi.StringPtrInput
	// The version of k3s to install (The default is currently the latest available).
	KubernetesVersion pulumi.StringPtrInput
	// A name for the Kubernetes cluster.
	Name pulumi.StringPtrInput
	// The number of instances to create (The default at the time of writing is 3).
	NumTargetNodes pulumi.IntPtrInput
	// A space separated list of tags, to be used freely as required.
	Tags pulumi.StringPtrInput
	// The size of each node (The default is currently g2.small)
	TargetNodesSize pulumi.StringPtrInput
}

func (KubernetesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesClusterArgs)(nil)).Elem()
}

type KubernetesClusterInput interface {
	pulumi.Input

	ToKubernetesClusterOutput() KubernetesClusterOutput
	ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput
}

func (KubernetesCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesCluster)(nil)).Elem()
}

func (i KubernetesCluster) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return i.ToKubernetesClusterOutputWithContext(context.Background())
}

func (i KubernetesCluster) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesClusterOutput)
}

type KubernetesClusterOutput struct {
	*pulumi.OutputState
}

func (KubernetesClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubernetesClusterOutput)(nil)).Elem()
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutput() KubernetesClusterOutput {
	return o
}

func (o KubernetesClusterOutput) ToKubernetesClusterOutputWithContext(ctx context.Context) KubernetesClusterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(KubernetesClusterOutput{})
}
