// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ContainerAnalysis.Outputs
{

    [OutputType]
    public sealed class NoteAttestationAuthority
    {
        public readonly Outputs.NoteAttestationAuthorityHint Hint;

        [OutputConstructor]
        private NoteAttestationAuthority(Outputs.NoteAttestationAuthorityHint hint)
        {
            Hint = hint;
        }
    }
}
