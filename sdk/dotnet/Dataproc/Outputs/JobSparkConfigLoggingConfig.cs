// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class JobSparkConfigLoggingConfig
    {
        public readonly ImmutableDictionary<string, string> DriverLogLevels;

        [OutputConstructor]
        private JobSparkConfigLoggingConfig(ImmutableDictionary<string, string> driverLogLevels)
        {
            DriverLogLevels = driverLogLevels;
        }
    }
}
