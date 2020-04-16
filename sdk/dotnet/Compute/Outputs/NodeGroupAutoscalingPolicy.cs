// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class NodeGroupAutoscalingPolicy
    {
        public readonly int? MaxNodes;
        public readonly int? MinNodes;
        public readonly string? Mode;

        [OutputConstructor]
        private NodeGroupAutoscalingPolicy(
            int? maxNodes,

            int? minNodes,

            string? mode)
        {
            MaxNodes = maxNodes;
            MinNodes = minNodes;
            Mode = mode;
        }
    }
}