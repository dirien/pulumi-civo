// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo
{
    /// <summary>
    /// Provides a Civo Kubernetes node pool resource. While the default node pool must be defined in the civo.KubernetesCluster resource, this resource can be used to add additional ones to a cluster.
    /// 
    /// ## Import
    /// 
    /// Then the Kubernetes cluster node pool can be imported using the cluster's and pool id `cluster_id:node_pool_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import civo:index/kubernetesNodePool:KubernetesNodePool my-pool 1b8b2100-0e9f-4e8f-ad78-9eb578c2a0af:502c1130-cb9b-4a88-b6d2-307bd96d946a
    /// ```
    /// </summary>
    [CivoResourceType("civo:index/kubernetesNodePool:KubernetesNodePool")]
    public partial class KubernetesNodePool : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Kubernetes cluster to which the node pool is associated.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The number of instances to create (The default at the time of writing is 3).
        /// </summary>
        [Output("numTargetNodes")]
        public Output<int> NumTargetNodes { get; private set; } = null!;

        /// <summary>
        /// The size of each node.
        /// </summary>
        [Output("targetNodesSize")]
        public Output<string> TargetNodesSize { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesNodePool(string name, KubernetesNodePoolArgs args, CustomResourceOptions? options = null)
            : base("civo:index/kubernetesNodePool:KubernetesNodePool", name, args ?? new KubernetesNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesNodePool(string name, Input<string> id, KubernetesNodePoolState? state = null, CustomResourceOptions? options = null)
            : base("civo:index/kubernetesNodePool:KubernetesNodePool", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubernetesNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesNodePool Get(string name, Input<string> id, KubernetesNodePoolState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesNodePool(name, id, state, options);
        }
    }

    public sealed class KubernetesNodePoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Kubernetes cluster to which the node pool is associated.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The number of instances to create (The default at the time of writing is 3).
        /// </summary>
        [Input("numTargetNodes")]
        public Input<int>? NumTargetNodes { get; set; }

        /// <summary>
        /// The size of each node.
        /// </summary>
        [Input("targetNodesSize")]
        public Input<string>? TargetNodesSize { get; set; }

        public KubernetesNodePoolArgs()
        {
        }
    }

    public sealed class KubernetesNodePoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Kubernetes cluster to which the node pool is associated.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The number of instances to create (The default at the time of writing is 3).
        /// </summary>
        [Input("numTargetNodes")]
        public Input<int>? NumTargetNodes { get; set; }

        /// <summary>
        /// The size of each node.
        /// </summary>
        [Input("targetNodesSize")]
        public Input<string>? TargetNodesSize { get; set; }

        public KubernetesNodePoolState()
        {
        }
    }
}