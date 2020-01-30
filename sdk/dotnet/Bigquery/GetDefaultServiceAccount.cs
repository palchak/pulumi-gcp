// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    public static partial class Invokes
    {
        /// <summary>
        /// Use this data source to retrieve default service account for this project
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/bigquery_default_service_account.html.markdown.
        /// </summary>
        public static Task<GetDefaultServiceAccountResult> GetDefaultServiceAccount(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDefaultServiceAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project ID. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDefaultServiceAccountArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDefaultServiceAccountResult
    {
        /// <summary>
        /// Email address of the default service account used by bigquery encryption in this project
        /// </summary>
        public readonly string Email;
        public readonly string Project;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDefaultServiceAccountResult(
            string email,
            string project,
            string id)
        {
            Email = email;
            Project = project;
            Id = id;
        }
    }
}