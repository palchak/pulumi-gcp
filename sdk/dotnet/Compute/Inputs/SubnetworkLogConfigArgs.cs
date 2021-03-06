// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class SubnetworkLogConfigArgs : Pulumi.ResourceArgs
    {
        [Input("aggregationInterval")]
        public Input<string>? AggregationInterval { get; set; }

        [Input("flowSampling")]
        public Input<double>? FlowSampling { get; set; }

        [Input("metadata")]
        public Input<string>? Metadata { get; set; }

        public SubnetworkLogConfigArgs()
        {
        }
    }
}
