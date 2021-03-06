// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class SubscriptionPushConfigArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("oidcToken")]
        public Input<Inputs.SubscriptionPushConfigOidcTokenArgs>? OidcToken { get; set; }

        [Input("pushEndpoint", required: true)]
        public Input<string> PushEndpoint { get; set; } = null!;

        public SubscriptionPushConfigArgs()
        {
        }
    }
}
