// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RouterNatLogConfigArgs : Pulumi.ResourceArgs
    {
        [Input("enable", required: true)]
        public Input<bool> Enable { get; set; } = null!;

        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        public RouterNatLogConfigArgs()
        {
        }
    }
}
