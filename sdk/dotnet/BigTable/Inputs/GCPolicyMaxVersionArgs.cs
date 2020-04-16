// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable.Inputs
{

    public sealed class GCPolicyMaxVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of version before applying the GC policy.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        public GCPolicyMaxVersionArgs()
        {
        }
    }
}