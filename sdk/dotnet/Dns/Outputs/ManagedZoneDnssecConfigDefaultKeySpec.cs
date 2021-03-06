// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Outputs
{

    [OutputType]
    public sealed class ManagedZoneDnssecConfigDefaultKeySpec
    {
        public readonly string? Algorithm;
        public readonly int? KeyLength;
        public readonly string? KeyType;
        public readonly string? Kind;

        [OutputConstructor]
        private ManagedZoneDnssecConfigDefaultKeySpec(
            string? algorithm,

            int? keyLength,

            string? keyType,

            string? kind)
        {
            Algorithm = algorithm;
            KeyLength = keyLength;
            KeyType = keyType;
            Kind = kind;
        }
    }
}
