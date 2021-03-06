// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudBuild.Inputs
{

    public sealed class TriggerBuildArgs : Pulumi.ResourceArgs
    {
        [Input("images")]
        private InputList<string>? _images;
        public InputList<string> Images
        {
            get => _images ?? (_images = new InputList<string>());
            set => _images = value;
        }

        [Input("steps", required: true)]
        private InputList<Inputs.TriggerBuildStepArgs>? _steps;
        public InputList<Inputs.TriggerBuildStepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.TriggerBuildStepArgs>());
            set => _steps = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        public TriggerBuildArgs()
        {
        }
    }
}
