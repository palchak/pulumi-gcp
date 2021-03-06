// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherPathRuleRouteActionGetArgs : Pulumi.ResourceArgs
    {
        [Input("corsPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionCorsPolicyGetArgs>? CorsPolicy { get; set; }

        [Input("faultInjectionPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionFaultInjectionPolicyGetArgs>? FaultInjectionPolicy { get; set; }

        [Input("requestMirrorPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionRequestMirrorPolicyGetArgs>? RequestMirrorPolicy { get; set; }

        [Input("retryPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionRetryPolicyGetArgs>? RetryPolicy { get; set; }

        [Input("timeout")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionTimeoutGetArgs>? Timeout { get; set; }

        [Input("urlRewrite")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionUrlRewriteGetArgs>? UrlRewrite { get; set; }

        [Input("weightedBackendServices")]
        private InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceGetArgs>? _weightedBackendServices;
        public InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceGetArgs> WeightedBackendServices
        {
            get => _weightedBackendServices ?? (_weightedBackendServices = new InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceGetArgs>());
            set => _weightedBackendServices = value;
        }

        public URLMapPathMatcherPathRuleRouteActionGetArgs()
        {
        }
    }
}
