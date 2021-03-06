// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherPathRuleRouteActionArgs : Pulumi.ResourceArgs
    {
        [Input("corsPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionCorsPolicyArgs>? CorsPolicy { get; set; }

        [Input("faultInjectionPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionFaultInjectionPolicyArgs>? FaultInjectionPolicy { get; set; }

        [Input("requestMirrorPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionRequestMirrorPolicyArgs>? RequestMirrorPolicy { get; set; }

        [Input("retryPolicy")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionRetryPolicyArgs>? RetryPolicy { get; set; }

        [Input("timeout")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionTimeoutArgs>? Timeout { get; set; }

        [Input("urlRewrite")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionUrlRewriteArgs>? UrlRewrite { get; set; }

        [Input("weightedBackendServices")]
        private InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs>? _weightedBackendServices;
        public InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs> WeightedBackendServices
        {
            get => _weightedBackendServices ?? (_weightedBackendServices = new InputList<Inputs.URLMapPathMatcherPathRuleRouteActionWeightedBackendServiceArgs>());
            set => _weightedBackendServices = value;
        }

        public URLMapPathMatcherPathRuleRouteActionArgs()
        {
        }
    }
}
