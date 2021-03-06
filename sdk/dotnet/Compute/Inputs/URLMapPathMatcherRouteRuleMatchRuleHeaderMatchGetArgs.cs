// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherRouteRuleMatchRuleHeaderMatchGetArgs : Pulumi.ResourceArgs
    {
        [Input("exactMatch")]
        public Input<string>? ExactMatch { get; set; }

        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        [Input("invertMatch")]
        public Input<bool>? InvertMatch { get; set; }

        [Input("prefixMatch")]
        public Input<string>? PrefixMatch { get; set; }

        [Input("presentMatch")]
        public Input<bool>? PresentMatch { get; set; }

        [Input("rangeMatch")]
        public Input<Inputs.URLMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatchGetArgs>? RangeMatch { get; set; }

        [Input("regexMatch")]
        public Input<string>? RegexMatch { get; set; }

        [Input("suffixMatch")]
        public Input<string>? SuffixMatch { get; set; }

        public URLMapPathMatcherRouteRuleMatchRuleHeaderMatchGetArgs()
        {
        }
    }
}
