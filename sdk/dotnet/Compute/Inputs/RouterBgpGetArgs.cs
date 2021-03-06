// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RouterBgpGetArgs : Pulumi.ResourceArgs
    {
        [Input("advertiseMode")]
        public Input<string>? AdvertiseMode { get; set; }

        [Input("advertisedGroups")]
        private InputList<string>? _advertisedGroups;
        public InputList<string> AdvertisedGroups
        {
            get => _advertisedGroups ?? (_advertisedGroups = new InputList<string>());
            set => _advertisedGroups = value;
        }

        [Input("advertisedIpRanges")]
        private InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs>? _advertisedIpRanges;
        public InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs> AdvertisedIpRanges
        {
            get => _advertisedIpRanges ?? (_advertisedIpRanges = new InputList<Inputs.RouterBgpAdvertisedIpRangeGetArgs>());
            set => _advertisedIpRanges = value;
        }

        [Input("asn", required: true)]
        public Input<int> Asn { get; set; } = null!;

        public RouterBgpGetArgs()
        {
        }
    }
}
