// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Endpoints.Outputs
{

    [OutputType]
    public sealed class ServiceApi
    {
        public readonly ImmutableArray<Outputs.ServiceApiMethod> Methods;
        public readonly string? Name;
        public readonly string? Syntax;
        public readonly string? Version;

        [OutputConstructor]
        private ServiceApi(
            ImmutableArray<Outputs.ServiceApiMethod> methods,

            string? name,

            string? syntax,

            string? version)
        {
            Methods = methods;
            Name = name;
            Syntax = syntax;
            Version = version;
        }
    }
}