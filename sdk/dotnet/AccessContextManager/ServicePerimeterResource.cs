// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager
{
    /// <summary>
    /// Allows configuring a single GCP resource that should be inside of a service perimeter.
    /// This resource is intended to be used in cases where it is not possible to compile a full list
    /// of projects to include in a `gcp.accesscontextmanager.ServicePerimeter` resource,
    /// to enable them to be added separately.
    /// 
    /// &gt; **Note:** If this resource is used alongside a `gcp.accesscontextmanager.ServicePerimeter` resource,
    /// the service perimeter resource must have a `lifecycle` block with `ignore_changes = [status[0].resources]` so
    /// they don't fight over which resources should be in the policy.
    /// 
    /// 
    /// To get more information about ServicePerimeterResource, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
    /// * How-to Guides
    ///     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
    /// </summary>
    public partial class ServicePerimeterResource : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the Service Perimeter to add this resource to.
        /// </summary>
        [Output("perimeterName")]
        public Output<string> PerimeterName { get; private set; } = null!;

        /// <summary>
        /// A GCP resource that is inside of the service perimeter. Currently only projects are allowed. Format:
        /// projects/{project_number}
        /// </summary>
        [Output("resource")]
        public Output<string> Resource { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePerimeterResource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePerimeterResource(string name, ServicePerimeterResourceArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, args ?? new ServicePerimeterResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePerimeterResource(string name, Input<string> id, ServicePerimeterResourceState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePerimeterResource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePerimeterResource Get(string name, Input<string> id, ServicePerimeterResourceState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePerimeterResource(name, id, state, options);
        }
    }

    public sealed class ServicePerimeterResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Service Perimeter to add this resource to.
        /// </summary>
        [Input("perimeterName", required: true)]
        public Input<string> PerimeterName { get; set; } = null!;

        /// <summary>
        /// A GCP resource that is inside of the service perimeter. Currently only projects are allowed. Format:
        /// projects/{project_number}
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        public ServicePerimeterResourceArgs()
        {
        }
    }

    public sealed class ServicePerimeterResourceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Service Perimeter to add this resource to.
        /// </summary>
        [Input("perimeterName")]
        public Input<string>? PerimeterName { get; set; }

        /// <summary>
        /// A GCP resource that is inside of the service perimeter. Currently only projects are allowed. Format:
        /// projects/{project_number}
        /// </summary>
        [Input("resource")]
        public Input<string>? Resource { get; set; }

        public ServicePerimeterResourceState()
        {
        }
    }
}
