// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GameServices.Inputs
{

    public sealed class GameServerDeploymentRolloutGameServerConfigOverrideGetArgs : Pulumi.ResourceArgs
    {
        [Input("configVersion")]
        public Input<string>? ConfigVersion { get; set; }

        [Input("realmsSelector")]
        public Input<Inputs.GameServerDeploymentRolloutGameServerConfigOverrideRealmsSelectorGetArgs>? RealmsSelector { get; set; }

        public GameServerDeploymentRolloutGameServerConfigOverrideGetArgs()
        {
        }
    }
}
