// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class URLMapPathMatcherRouteRuleMatchRuleMetadataFilter
    {
        public readonly ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel> FilterLabels;
        public readonly string FilterMatchCriteria;

        [OutputConstructor]
        private URLMapPathMatcherRouteRuleMatchRuleMetadataFilter(
            ImmutableArray<Outputs.URLMapPathMatcherRouteRuleMatchRuleMetadataFilterFilterLabel> filterLabels,

            string filterMatchCriteria)
        {
            FilterLabels = filterLabels;
            FilterMatchCriteria = filterMatchCriteria;
        }
    }
}
