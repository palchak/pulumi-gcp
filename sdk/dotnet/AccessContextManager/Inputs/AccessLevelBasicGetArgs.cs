// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.AccessContextManager.Inputs
{

    public sealed class AccessLevelBasicGetArgs : Pulumi.ResourceArgs
    {
        [Input("combiningFunction")]
        public Input<string>? CombiningFunction { get; set; }

        [Input("conditions", required: true)]
        private InputList<Inputs.AccessLevelBasicConditionGetArgs>? _conditions;
        public InputList<Inputs.AccessLevelBasicConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.AccessLevelBasicConditionGetArgs>());
            set => _conditions = value;
        }

        public AccessLevelBasicGetArgs()
        {
        }
    }
}
