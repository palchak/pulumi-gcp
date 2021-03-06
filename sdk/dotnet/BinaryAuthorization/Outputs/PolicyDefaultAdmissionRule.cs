// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BinaryAuthorization.Outputs
{

    [OutputType]
    public sealed class PolicyDefaultAdmissionRule
    {
        public readonly string EnforcementMode;
        public readonly string EvaluationMode;
        public readonly ImmutableArray<string> RequireAttestationsBies;

        [OutputConstructor]
        private PolicyDefaultAdmissionRule(
            string enforcementMode,

            string evaluationMode,

            ImmutableArray<string> requireAttestationsBies)
        {
            EnforcementMode = enforcementMode;
            EvaluationMode = evaluationMode;
            RequireAttestationsBies = requireAttestationsBies;
        }
    }
}
