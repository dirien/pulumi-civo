// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo.Outputs
{

    [OutputType]
    public sealed class GetRegionFilterResult
    {
        public readonly bool? All;
        /// <summary>
        /// Filter the sizes by this key. This may be one of `code`, `name`, `country`, `default`.
        /// </summary>
        public readonly string Key;
        public readonly string? MatchBy;
        /// <summary>
        /// Only retrieves region which keys has value that matches one of the values provided here.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetRegionFilterResult(
            bool? all,

            string key,

            string? matchBy,

            ImmutableArray<string> values)
        {
            All = all;
            Key = key;
            MatchBy = matchBy;
            Values = values;
        }
    }
}