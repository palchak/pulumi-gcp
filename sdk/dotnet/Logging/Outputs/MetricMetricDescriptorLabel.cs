// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging.Outputs
{

    [OutputType]
    public sealed class MetricMetricDescriptorLabel
    {
        public readonly string? Description;
        public readonly string Key;
        public readonly string? ValueType;

        [OutputConstructor]
        private MetricMetricDescriptorLabel(
            string? description,

            string key,

            string? valueType)
        {
            Description = description;
            Key = key;
            ValueType = valueType;
        }
    }
}
