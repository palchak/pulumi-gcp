// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherPathRuleRouteActionArgs : Pulumi.ResourceArgs
    {
        [Input("corsPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionCorsPolicyArgs>? CorsPolicy { get; set; }

        [Input("faultInjectionPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyArgs>? FaultInjectionPolicy { get; set; }

        [Input("requestMirrorPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicyArgs>? RequestMirrorPolicy { get; set; }

        [Input("retryPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionRetryPolicyArgs>? RetryPolicy { get; set; }

        [Input("timeout")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionTimeoutArgs>? Timeout { get; set; }

        [Input("urlRewrite")]
        public Input<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionUrlRewriteArgs>? UrlRewrite { get; set; }

        [Input("weightedBackendServices")]
        private InputList<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs>? _weightedBackendServices;
        public InputList<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs> WeightedBackendServices
        {
            get => _weightedBackendServices ?? (_weightedBackendServices = new InputList<Inputs.RegionUrlMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs>());
            set => _weightedBackendServices = value;
        }

        public RegionUrlMapPathMatcherPathRuleRouteActionArgs()
        {
        }
    }
}
