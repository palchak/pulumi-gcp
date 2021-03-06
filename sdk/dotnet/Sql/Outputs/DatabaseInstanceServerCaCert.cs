// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Outputs
{

    [OutputType]
    public sealed class DatabaseInstanceServerCaCert
    {
        public readonly string? Cert;
        public readonly string? CommonName;
        public readonly string? CreateTime;
        /// <summary>
        /// The [RFC 3339](https://tools.ietf.org/html/rfc3339)
        /// formatted date time string indicating when this whitelist expires.
        /// </summary>
        public readonly string? ExpirationTime;
        public readonly string? Sha1Fingerprint;

        [OutputConstructor]
        private DatabaseInstanceServerCaCert(
            string? cert,

            string? commonName,

            string? createTime,

            string? expirationTime,

            string? sha1Fingerprint)
        {
            Cert = cert;
            CommonName = commonName;
            CreateTime = createTime;
            ExpirationTime = expirationTime;
            Sha1Fingerprint = sha1Fingerprint;
        }
    }
}
