// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionHandlerGetArgs : Pulumi.ResourceArgs
    {
        [Input("authFailAction")]
        public Input<string>? AuthFailAction { get; set; }

        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("redirectHttpResponseCode")]
        public Input<string>? RedirectHttpResponseCode { get; set; }

        [Input("script")]
        public Input<Inputs.StandardAppVersionHandlerScriptGetArgs>? Script { get; set; }

        [Input("securityLevel")]
        public Input<string>? SecurityLevel { get; set; }

        [Input("staticFiles")]
        public Input<Inputs.StandardAppVersionHandlerStaticFilesGetArgs>? StaticFiles { get; set; }

        [Input("urlRegex")]
        public Input<string>? UrlRegex { get; set; }

        public StandardAppVersionHandlerGetArgs()
        {
        }
    }
}
