// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AppEngine.Inputs
{

    public sealed class StandardAppVersionHandlerScriptGetArgs : Pulumi.ResourceArgs
    {
        [Input("scriptPath", required: true)]
        public Input<string> ScriptPath { get; set; } = null!;

        public StandardAppVersionHandlerScriptGetArgs()
        {
        }
    }
}
