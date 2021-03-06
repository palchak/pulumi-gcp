// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Inputs
{

    public sealed class DomainMappingStatusArgs : Pulumi.ResourceArgs
    {
        [Input("conditions")]
        private InputList<Inputs.DomainMappingStatusConditionArgs>? _conditions;
        public InputList<Inputs.DomainMappingStatusConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.DomainMappingStatusConditionArgs>());
            set => _conditions = value;
        }

        [Input("mappedRouteName")]
        public Input<string>? MappedRouteName { get; set; }

        [Input("observedGeneration")]
        public Input<int>? ObservedGeneration { get; set; }

        [Input("resourceRecords")]
        private InputList<Inputs.DomainMappingStatusResourceRecordArgs>? _resourceRecords;
        public InputList<Inputs.DomainMappingStatusResourceRecordArgs> ResourceRecords
        {
            get => _resourceRecords ?? (_resourceRecords = new InputList<Inputs.DomainMappingStatusResourceRecordArgs>());
            set => _resourceRecords = value;
        }

        public DomainMappingStatusArgs()
        {
        }
    }
}
