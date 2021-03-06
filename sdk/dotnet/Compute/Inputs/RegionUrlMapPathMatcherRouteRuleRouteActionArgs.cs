// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionUrlMapPathMatcherRouteRuleRouteActionArgs : Pulumi.ResourceArgs
    {
        [Input("corsPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionCorsPolicyArgs>? CorsPolicy { get; set; }

        [Input("faultInjectionPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionFaultInjectionPolicyArgs>? FaultInjectionPolicy { get; set; }

        [Input("requestMirrorPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionRequestMirrorPolicyArgs>? RequestMirrorPolicy { get; set; }

        [Input("retryPolicy")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionRetryPolicyArgs>? RetryPolicy { get; set; }

        [Input("timeout")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionTimeoutArgs>? Timeout { get; set; }

        [Input("urlRewrite")]
        public Input<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionUrlRewriteArgs>? UrlRewrite { get; set; }

        [Input("weightedBackendServices")]
        private InputList<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArgs>? _weightedBackendServices;
        public InputList<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArgs> WeightedBackendServices
        {
            get => _weightedBackendServices ?? (_weightedBackendServices = new InputList<Inputs.RegionUrlMapPathMatcherRouteRuleRouteActionWeightedBackendServiceArgs>());
            set => _weightedBackendServices = value;
        }

        public RegionUrlMapPathMatcherRouteRuleRouteActionArgs()
        {
        }
    }
}
