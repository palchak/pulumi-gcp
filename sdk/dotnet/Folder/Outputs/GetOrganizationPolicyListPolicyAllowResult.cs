// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Folder.Outputs
{

    [OutputType]
    public sealed class GetOrganizationPolicyListPolicyAllowResult
    {
        public readonly bool All;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetOrganizationPolicyListPolicyAllowResult(
            bool all,

            ImmutableArray<string> values)
        {
            All = all;
            Values = values;
        }
    }
}
