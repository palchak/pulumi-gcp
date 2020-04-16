// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class BackendServiceCdnPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("cacheKeyPolicy")]
        public Input<Inputs.BackendServiceCdnPolicyCacheKeyPolicyGetArgs>? CacheKeyPolicy { get; set; }

        [Input("signedUrlCacheMaxAgeSec")]
        public Input<int>? SignedUrlCacheMaxAgeSec { get; set; }

        public BackendServiceCdnPolicyGetArgs()
        {
        }
    }
}