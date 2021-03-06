// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager.Inputs
{

    public sealed class AccessLevelBasicArgs : Pulumi.ResourceArgs
    {
        [Input("combiningFunction")]
        public Input<string>? CombiningFunction { get; set; }

        [Input("conditions", required: true)]
        private InputList<Inputs.AccessLevelBasicConditionArgs>? _conditions;
        public InputList<Inputs.AccessLevelBasicConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AccessLevelBasicConditionArgs>());
            set => _conditions = value;
        }

        public AccessLevelBasicArgs()
        {
        }
    }
}
