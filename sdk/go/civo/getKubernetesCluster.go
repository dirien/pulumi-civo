// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Civo Kubernetes cluster data source.
//
// **Note:** This data source returns a single kubernetes cluster. When specifying a `name`, an
// error is triggered if more than one kubernetes Cluster is found.
//
// ## Example Usage
//
// Get the Kubernetes Cluster by name:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-civo/sdk/go/civo"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "my-super-cluster"
// 		my_cluster, err := civo.LookupKubernetesCluster(ctx, &civo.LookupKubernetesClusterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("kubernetesClusterOutput", my_cluster.MasterIp)
// 		return nil
// 	})
// }
// ```
//
// Get the Kubernetes Cluster by id:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-civo/sdk/go/civo"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "40ac97ee-b82b-4231-9b60-079c7e2e5d79"
// 		_, err := civo.LookupKubernetesCluster(ctx, &civo.LookupKubernetesClusterArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("civo:index/getKubernetesCluster:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterArgs struct {
	// The ID of the kubernetes Cluster
	Id *string `pulumi:"id"`
	// The name of the kubernetes Cluster.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResult struct {
	// The base URL of the API server on the Kubernetes master node.
	ApiEndpoint string `pulumi:"apiEndpoint"`
	// A list of application installed.
	Applications string `pulumi:"applications"`
	// The date where the Kubernetes cluster was build.
	BuiltAt string `pulumi:"builtAt"`
	// The date where the Kubernetes cluster was create.
	CreatedAt string `pulumi:"createdAt"`
	// The unique dns entry for the cluster in this case point to the master.
	DnsEntry string `pulumi:"dnsEntry"`
	// The ID of the pool
	Id *string `pulumi:"id"`
	// A unique ID that can be used to identify and reference a Kubernetes cluster.
	InstalledApplications []GetKubernetesClusterInstalledApplication `pulumi:"installedApplications"`
	// A list of instance inside the pool
	Instances []GetKubernetesClusterInstance `pulumi:"instances"`
	// A representation of the Kubernetes cluster's kubeconfig in yaml format.
	Kubeconfig string `pulumi:"kubeconfig"`
	// The version of Kubernetes.
	KubernetesVersion string `pulumi:"kubernetesVersion"`
	// The Ip of the Kubernetes master node.
	MasterIp string `pulumi:"masterIp"`
	// The name of your cluster,.
	Name *string `pulumi:"name"`
	// The size of the Kubernetes cluster.
	NumTargetNodes int `pulumi:"numTargetNodes"`
	// A list of node pools associated with the cluster. Each node pool exports the following attributes:
	Pools  []GetKubernetesClusterPool `pulumi:"pools"`
	Ready  bool                       `pulumi:"ready"`
	Region *string                    `pulumi:"region"`
	// The status of Kubernetes cluster.
	// * `ready` -If the Kubernetes cluster is ready.
	Status string `pulumi:"status"`
	// The tag of the instances
	Tags string `pulumi:"tags"`
	// The size of each node.
	TargetNodesSize string `pulumi:"targetNodesSize"`
}
