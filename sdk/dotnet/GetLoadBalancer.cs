// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo
{
    public static class GetLoadBalancer
    {
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("civo:index/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithVersion());
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hostname of the Load Balancer.
        /// </summary>
        [Input("hostname")]
        public string? Hostname { get; set; }

        /// <summary>
        /// The ID of the Load Balancer.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        public GetLoadBalancerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        /// <summary>
        /// A list of backend instances
        /// </summary>
        public readonly ImmutableArray<Outputs.GetLoadBalancerBackendResult> Backends;
        /// <summary>
        /// The wait time until the backend is marked as a failure
        /// </summary>
        public readonly int FailTimeout;
        /// <summary>
        /// The path to check the health of the backend
        /// </summary>
        public readonly string HealthCheckPath;
        /// <summary>
        /// The hostname of the Load Balancer
        /// </summary>
        public readonly string? Hostname;
        /// <summary>
        /// The ID of the Load Balancer
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Should self-signed/invalid certificates be ignored from the backend servers
        /// </summary>
        public readonly bool IgnoreInvalidBackendTls;
        /// <summary>
        /// How many concurrent connections can each backend handle
        /// </summary>
        public readonly int MaxConns;
        /// <summary>
        /// The max request size set in the configuration
        /// </summary>
        public readonly int MaxRequestSize;
        /// <summary>
        /// The policy set in the Load Balancer
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// The port set in the configuration.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The protocol used in the configuration.
        /// </summary>
        public readonly string Protocol;
        public readonly string? Region;
        /// <summary>
        /// If is set will be returned
        /// </summary>
        public readonly string TlsCertificate;
        /// <summary>
        /// If is set will be returned
        /// </summary>
        public readonly string TlsKey;

        [OutputConstructor]
        private GetLoadBalancerResult(
            ImmutableArray<Outputs.GetLoadBalancerBackendResult> backends,

            int failTimeout,

            string healthCheckPath,

            string? hostname,

            string? id,

            bool ignoreInvalidBackendTls,

            int maxConns,

            int maxRequestSize,

            string policy,

            int port,

            string protocol,

            string? region,

            string tlsCertificate,

            string tlsKey)
        {
            Backends = backends;
            FailTimeout = failTimeout;
            HealthCheckPath = healthCheckPath;
            Hostname = hostname;
            Id = id;
            IgnoreInvalidBackendTls = ignoreInvalidBackendTls;
            MaxConns = maxConns;
            MaxRequestSize = maxRequestSize;
            Policy = policy;
            Port = port;
            Protocol = protocol;
            Region = region;
            TlsCertificate = tlsCertificate;
            TlsKey = tlsKey;
        }
    }
}
