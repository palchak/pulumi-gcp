// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class PacketMirroringFilterGetArgs : Pulumi.ResourceArgs
    {
        [Input("cidrRanges")]
        private InputList<string>? _cidrRanges;
        public InputList<string> CidrRanges
        {
            get => _cidrRanges ?? (_cidrRanges = new InputList<string>());
            set => _cidrRanges = value;
        }

        [Input("ipProtocols")]
        private InputList<string>? _ipProtocols;
        public InputList<string> IpProtocols
        {
            get => _ipProtocols ?? (_ipProtocols = new InputList<string>());
            set => _ipProtocols = value;
        }

        public PacketMirroringFilterGetArgs()
        {
        }
    }
}
