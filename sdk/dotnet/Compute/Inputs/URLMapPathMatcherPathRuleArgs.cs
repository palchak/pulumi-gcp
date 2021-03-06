// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class URLMapPathMatcherPathRuleArgs : Pulumi.ResourceArgs
    {
        [Input("paths", required: true)]
        private InputList<string>? _paths;
        public InputList<string> Paths
        {
            get => _paths ?? (_paths = new InputList<string>());
            set => _paths = value;
        }

        [Input("routeAction")]
        public Input<Inputs.URLMapPathMatcherPathRuleRouteActionArgs>? RouteAction { get; set; }

        [Input("service")]
        public Input<string>? Service { get; set; }

        [Input("urlRedirect")]
        public Input<Inputs.URLMapPathMatcherPathRuleUrlRedirectArgs>? UrlRedirect { get; set; }

        public URLMapPathMatcherPathRuleArgs()
        {
        }
    }
}
