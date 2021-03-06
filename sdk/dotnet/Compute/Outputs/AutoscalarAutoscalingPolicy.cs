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
    public sealed class AutoscalarAutoscalingPolicy
    {
        public readonly int? CooldownPeriod;
        public readonly Outputs.AutoscalarAutoscalingPolicyCpuUtilization? CpuUtilization;
        public readonly Outputs.AutoscalarAutoscalingPolicyLoadBalancingUtilization? LoadBalancingUtilization;
        public readonly int MaxReplicas;
        public readonly ImmutableArray<Outputs.AutoscalarAutoscalingPolicyMetric> Metrics;
        public readonly int MinReplicas;

        [OutputConstructor]
        private AutoscalarAutoscalingPolicy(
            int? cooldownPeriod,

            Outputs.AutoscalarAutoscalingPolicyCpuUtilization? cpuUtilization,

            Outputs.AutoscalarAutoscalingPolicyLoadBalancingUtilization? loadBalancingUtilization,

            int maxReplicas,

            ImmutableArray<Outputs.AutoscalarAutoscalingPolicyMetric> metrics,

            int minReplicas)
        {
            CooldownPeriod = cooldownPeriod;
            CpuUtilization = cpuUtilization;
            LoadBalancingUtilization = loadBalancingUtilization;
            MaxReplicas = maxReplicas;
            Metrics = metrics;
            MinReplicas = minReplicas;
        }
    }
}
