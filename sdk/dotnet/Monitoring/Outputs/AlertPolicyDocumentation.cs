// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Outputs
{

    [OutputType]
    public sealed class AlertPolicyDocumentation
    {
        public readonly string? Content;
        public readonly string? MimeType;

        [OutputConstructor]
        private AlertPolicyDocumentation(
            string? content,

            string? mimeType)
        {
            Content = content;
            MimeType = mimeType;
        }
    }
}