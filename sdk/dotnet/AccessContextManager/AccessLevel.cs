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
    /// An AccessLevel is a label that can be applied to requests to GCP services,
    /// along with a list of requirements necessary for the label to be applied.
    /// 
    /// 
    /// To get more information about AccessLevel, see:
    /// 
    /// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
    /// * How-to Guides
    ///     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
    /// </summary>
    public partial class AccessLevel : Pulumi.CustomResource
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Output("basic")]
        public Output<Outputs.AccessLevelBasic?> Basic { get; private set; } = null!;

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
        /// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a AccessLevel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessLevel(string name, AccessLevelArgs args, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevel:AccessLevel", name, args ?? new AccessLevelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessLevel(string name, Input<string> id, AccessLevelState? state = null, CustomResourceOptions? options = null)
            : base("gcp:accesscontextmanager/accessLevel:AccessLevel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccessLevel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessLevel Get(string name, Input<string> id, AccessLevelState? state = null, CustomResourceOptions? options = null)
        {
            return new AccessLevel(name, id, state, options);
        }
    }

    public sealed class AccessLevelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Input("basic")]
        public Input<Inputs.AccessLevelBasicArgs>? Basic { get; set; }

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
        /// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public AccessLevelArgs()
        {
        }
    }

    public sealed class AccessLevelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A set of predefined conditions for the access level and a combining function.
        /// </summary>
        [Input("basic")]
        public Input<Inputs.AccessLevelBasicGetArgs>? Basic { get; set; }

        /// <summary>
        /// Description of the AccessLevel and its use. Does not affect behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
        /// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// Human readable title. Must be unique within the Policy.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public AccessLevelState()
        {
        }
    }
}
