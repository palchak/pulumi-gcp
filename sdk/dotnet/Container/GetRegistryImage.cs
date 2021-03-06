// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container
{
    public static class GetRegistryImage
    {
        /// <summary>
        /// This data source fetches the project name, and provides the appropriate URLs to use for container registry for this project.
        /// 
        /// The URLs are computed entirely offline - as long as the project exists, they will be valid, but this data source does not contact Google Container Registry (GCR) at any point.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryImageResult> InvokeAsync(GetRegistryImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryImageResult>("gcp:container/getRegistryImage:getRegistryImage", args ?? new GetRegistryImageArgs(), options.WithVersion());
    }


    public sealed class GetRegistryImageArgs : Pulumi.InvokeArgs
    {
        [Input("digest")]
        public string? Digest { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("region")]
        public string? Region { get; set; }

        [Input("tag")]
        public string? Tag { get; set; }

        public GetRegistryImageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryImageResult
    {
        public readonly string? Digest;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ImageUrl;
        public readonly string Name;
        public readonly string Project;
        public readonly string? Region;
        public readonly string? Tag;

        [OutputConstructor]
        private GetRegistryImageResult(
            string? digest,

            string id,

            string imageUrl,

            string name,

            string project,

            string? region,

            string? tag)
        {
            Digest = digest;
            Id = id;
            ImageUrl = imageUrl;
            Name = name;
            Project = project;
            Region = region;
            Tag = tag;
        }
    }
}
