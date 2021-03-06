// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices
{
    /// <summary>
    /// A game server config resource. Configs are global and immutable.
    /// 
    /// To get more information about GameServerConfig, see:
    /// 
    /// * [API documentation](https://cloud.google.com/game-servers/docs/reference/rest/v1beta/projects.locations.gameServerDeployments.configs)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/game-servers/docs)
    /// </summary>
    public partial class GameServerConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique id for the deployment config.
        /// </summary>
        [Output("configId")]
        public Output<string> ConfigId { get; private set; } = null!;

        /// <summary>
        /// A unique id for the deployment.
        /// </summary>
        [Output("deploymentId")]
        public Output<string> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// The description of the game server config.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The fleet config contains list of fleet specs. In the Single Cloud, there will be only one.
        /// </summary>
        [Output("fleetConfigs")]
        public Output<ImmutableArray<Outputs.GameServerConfigFleetConfig>> FleetConfigs { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this game server config. Each label is a key-value pair.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Location of the Deployment.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the game server config, in the form:
        /// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}/configs/{config_id}'.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. This contains the autoscaling settings.
        /// </summary>
        [Output("scalingConfigs")]
        public Output<ImmutableArray<Outputs.GameServerConfigScalingConfig>> ScalingConfigs { get; private set; } = null!;


        /// <summary>
        /// Create a GameServerConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GameServerConfig(string name, GameServerConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:gameservices/gameServerConfig:GameServerConfig", name, args ?? new GameServerConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GameServerConfig(string name, Input<string> id, GameServerConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gameservices/gameServerConfig:GameServerConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GameServerConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GameServerConfig Get(string name, Input<string> id, GameServerConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new GameServerConfig(name, id, state, options);
        }
    }

    public sealed class GameServerConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique id for the deployment config.
        /// </summary>
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        /// <summary>
        /// A unique id for the deployment.
        /// </summary>
        [Input("deploymentId", required: true)]
        public Input<string> DeploymentId { get; set; } = null!;

        /// <summary>
        /// The description of the game server config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fleetConfigs", required: true)]
        private InputList<Inputs.GameServerConfigFleetConfigArgs>? _fleetConfigs;

        /// <summary>
        /// The fleet config contains list of fleet specs. In the Single Cloud, there will be only one.
        /// </summary>
        public InputList<Inputs.GameServerConfigFleetConfigArgs> FleetConfigs
        {
            get => _fleetConfigs ?? (_fleetConfigs = new InputList<Inputs.GameServerConfigFleetConfigArgs>());
            set => _fleetConfigs = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server config. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Deployment.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("scalingConfigs")]
        private InputList<Inputs.GameServerConfigScalingConfigArgs>? _scalingConfigs;

        /// <summary>
        /// Optional. This contains the autoscaling settings.
        /// </summary>
        public InputList<Inputs.GameServerConfigScalingConfigArgs> ScalingConfigs
        {
            get => _scalingConfigs ?? (_scalingConfigs = new InputList<Inputs.GameServerConfigScalingConfigArgs>());
            set => _scalingConfigs = value;
        }

        public GameServerConfigArgs()
        {
        }
    }

    public sealed class GameServerConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique id for the deployment config.
        /// </summary>
        [Input("configId")]
        public Input<string>? ConfigId { get; set; }

        /// <summary>
        /// A unique id for the deployment.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The description of the game server config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fleetConfigs")]
        private InputList<Inputs.GameServerConfigFleetConfigGetArgs>? _fleetConfigs;

        /// <summary>
        /// The fleet config contains list of fleet specs. In the Single Cloud, there will be only one.
        /// </summary>
        public InputList<Inputs.GameServerConfigFleetConfigGetArgs> FleetConfigs
        {
            get => _fleetConfigs ?? (_fleetConfigs = new InputList<Inputs.GameServerConfigFleetConfigGetArgs>());
            set => _fleetConfigs = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this game server config. Each label is a key-value pair.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Location of the Deployment.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the game server config, in the form:
        /// 'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}/configs/{config_id}'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("scalingConfigs")]
        private InputList<Inputs.GameServerConfigScalingConfigGetArgs>? _scalingConfigs;

        /// <summary>
        /// Optional. This contains the autoscaling settings.
        /// </summary>
        public InputList<Inputs.GameServerConfigScalingConfigGetArgs> ScalingConfigs
        {
            get => _scalingConfigs ?? (_scalingConfigs = new InputList<Inputs.GameServerConfigScalingConfigGetArgs>());
            set => _scalingConfigs = value;
        }

        public GameServerConfigState()
        {
        }
    }
}
