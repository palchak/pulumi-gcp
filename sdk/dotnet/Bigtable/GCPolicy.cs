// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Bigtable
{
    /// <summary>
    /// Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
    /// [the official documentation](https://cloud.google.com/bigtable/) and
    /// [API](https://cloud.google.com/bigtable/docs/go/reference).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigtable_gc_policy.html.markdown.
    /// </summary>
    public partial class GCPolicy : Pulumi.CustomResource
    {
        [Output("columnFamily")]
        public Output<string> ColumnFamily { get; private set; } = null!;

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// GC policy that applies to all cells older than the given age.
        /// </summary>
        [Output("maxAges")]
        public Output<ImmutableArray<Outputs.GCPolicyMaxAges>> MaxAges { get; private set; } = null!;

        /// <summary>
        /// GC policy that applies to all versions of a cell except for the most recent.
        /// </summary>
        [Output("maxVersions")]
        public Output<ImmutableArray<Outputs.GCPolicyMaxVersions>> MaxVersions { get; private set; } = null!;

        /// <summary>
        /// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("table")]
        public Output<string> Table { get; private set; } = null!;


        /// <summary>
        /// Create a GCPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GCPolicy(string name, GCPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigtable/gCPolicy:GCPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GCPolicy(string name, Input<string> id, GCPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/gCPolicy:GCPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GCPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GCPolicy Get(string name, Input<string> id, GCPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new GCPolicy(name, id, state, options);
        }
    }

    public sealed class GCPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("columnFamily", required: true)]
        public Input<string> ColumnFamily { get; set; } = null!;

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        [Input("maxAges")]
        private InputList<Inputs.GCPolicyMaxAgesArgs>? _maxAges;

        /// <summary>
        /// GC policy that applies to all cells older than the given age.
        /// </summary>
        public InputList<Inputs.GCPolicyMaxAgesArgs> MaxAges
        {
            get => _maxAges ?? (_maxAges = new InputList<Inputs.GCPolicyMaxAgesArgs>());
            set => _maxAges = value;
        }

        [Input("maxVersions")]
        private InputList<Inputs.GCPolicyMaxVersionsArgs>? _maxVersions;

        /// <summary>
        /// GC policy that applies to all versions of a cell except for the most recent.
        /// </summary>
        public InputList<Inputs.GCPolicyMaxVersionsArgs> MaxVersions
        {
            get => _maxVersions ?? (_maxVersions = new InputList<Inputs.GCPolicyMaxVersionsArgs>());
            set => _maxVersions = value;
        }

        /// <summary>
        /// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public GCPolicyArgs()
        {
        }
    }

    public sealed class GCPolicyState : Pulumi.ResourceArgs
    {
        [Input("columnFamily")]
        public Input<string>? ColumnFamily { get; set; }

        /// <summary>
        /// The name of the Bigtable instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        [Input("maxAges")]
        private InputList<Inputs.GCPolicyMaxAgesGetArgs>? _maxAges;

        /// <summary>
        /// GC policy that applies to all cells older than the given age.
        /// </summary>
        public InputList<Inputs.GCPolicyMaxAgesGetArgs> MaxAges
        {
            get => _maxAges ?? (_maxAges = new InputList<Inputs.GCPolicyMaxAgesGetArgs>());
            set => _maxAges = value;
        }

        [Input("maxVersions")]
        private InputList<Inputs.GCPolicyMaxVersionsGetArgs>? _maxVersions;

        /// <summary>
        /// GC policy that applies to all versions of a cell except for the most recent.
        /// </summary>
        public InputList<Inputs.GCPolicyMaxVersionsGetArgs> MaxVersions
        {
            get => _maxVersions ?? (_maxVersions = new InputList<Inputs.GCPolicyMaxVersionsGetArgs>());
            set => _maxVersions = value;
        }

        /// <summary>
        /// If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("table")]
        public Input<string>? Table { get; set; }

        public GCPolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class GCPolicyMaxAgesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days before applying GC policy.
        /// </summary>
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        public GCPolicyMaxAgesArgs()
        {
        }
    }

    public sealed class GCPolicyMaxAgesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of days before applying GC policy.
        /// </summary>
        [Input("days", required: true)]
        public Input<int> Days { get; set; } = null!;

        public GCPolicyMaxAgesGetArgs()
        {
        }
    }

    public sealed class GCPolicyMaxVersionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of version before applying the GC policy.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        public GCPolicyMaxVersionsArgs()
        {
        }
    }

    public sealed class GCPolicyMaxVersionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of version before applying the GC policy.
        /// </summary>
        [Input("number", required: true)]
        public Input<int> Number { get; set; } = null!;

        public GCPolicyMaxVersionsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GCPolicyMaxAges
    {
        /// <summary>
        /// Number of days before applying GC policy.
        /// </summary>
        public readonly int Days;

        [OutputConstructor]
        private GCPolicyMaxAges(int days)
        {
            Days = days;
        }
    }

    [OutputType]
    public sealed class GCPolicyMaxVersions
    {
        /// <summary>
        /// Number of version before applying the GC policy.
        /// </summary>
        public readonly int Number;

        [OutputConstructor]
        private GCPolicyMaxVersions(int number)
        {
            Number = number;
        }
    }
    }
}
