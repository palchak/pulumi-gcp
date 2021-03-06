// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class BackendServiceCircuitBreakersArgs : Pulumi.ResourceArgs
    {
        [Input("connectTimeout")]
        public Input<Inputs.BackendServiceCircuitBreakersConnectTimeoutArgs>? ConnectTimeout { get; set; }

        [Input("maxConnections")]
        public Input<int>? MaxConnections { get; set; }

        [Input("maxPendingRequests")]
        public Input<int>? MaxPendingRequests { get; set; }

        [Input("maxRequests")]
        public Input<int>? MaxRequests { get; set; }

        [Input("maxRequestsPerConnection")]
        public Input<int>? MaxRequestsPerConnection { get; set; }

        [Input("maxRetries")]
        public Input<int>? MaxRetries { get; set; }

        public BackendServiceCircuitBreakersArgs()
        {
        }
    }
}
