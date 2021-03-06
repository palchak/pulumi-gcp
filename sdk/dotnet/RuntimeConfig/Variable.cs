// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.RuntimeConfig
{
    /// <summary>
    /// Manages a RuntimeConfig variable in Google Cloud. For more information, see the
    /// [official documentation](https://cloud.google.com/deployment-manager/runtime-configurator/),
    /// or the
    /// [JSON API](https://cloud.google.com/deployment-manager/runtime-configurator/reference/rest/).
    /// </summary>
    public partial class Variable : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the variable to manage. Note that variable
        /// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the RuntimeConfig resource containing this
        /// variable.
        /// </summary>
        [Output("parent")]
        public Output<string> Parent { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("text")]
        public Output<string?> Text { get; private set; } = null!;

        /// <summary>
        /// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds, representing when the variable was last updated.
        /// Example: "2016-10-09T12:33:37.578138407Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a Variable resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Variable(string name, VariableArgs args, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/variable:Variable", name, args ?? new VariableArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Variable(string name, Input<string> id, VariableState? state = null, CustomResourceOptions? options = null)
            : base("gcp:runtimeconfig/variable:Variable", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Variable resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Variable Get(string name, Input<string> id, VariableState? state = null, CustomResourceOptions? options = null)
        {
            return new Variable(name, id, state, options);
        }
    }

    public sealed class VariableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the variable to manage. Note that variable
        /// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the RuntimeConfig resource containing this
        /// variable.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("text")]
        public Input<string>? Text { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public VariableArgs()
        {
        }
    }

    public sealed class VariableState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the variable to manage. Note that variable
        /// names can be hierarchical using slashes (e.g. "prod-variables/hostname").
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the RuntimeConfig resource containing this
        /// variable.
        /// </summary>
        [Input("parent")]
        public Input<string>? Parent { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("text")]
        public Input<string>? Text { get; set; }

        /// <summary>
        /// (Computed) The timestamp in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds, representing when the variable was last updated.
        /// Example: "2016-10-09T12:33:37.578138407Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public VariableState()
        {
        }
    }
}
