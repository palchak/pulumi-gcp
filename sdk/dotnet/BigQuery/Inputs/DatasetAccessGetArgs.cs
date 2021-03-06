// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery.Inputs
{

    public sealed class DatasetAccessGetArgs : Pulumi.ResourceArgs
    {
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("groupByEmail")]
        public Input<string>? GroupByEmail { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("specialGroup")]
        public Input<string>? SpecialGroup { get; set; }

        [Input("userByEmail")]
        public Input<string>? UserByEmail { get; set; }

        [Input("view")]
        public Input<Inputs.DatasetAccessViewGetArgs>? View { get; set; }

        public DatasetAccessGetArgs()
        {
        }
    }
}
