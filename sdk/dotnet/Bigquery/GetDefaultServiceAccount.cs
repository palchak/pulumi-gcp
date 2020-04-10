// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    public static partial class Invokes
    {
        /// <summary>
        /// Get the email address of a project's unique BigQuery service account.
        /// 
        /// Each Google Cloud project has a unique service account used by BigQuery. When using
        /// BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
        /// this account needs to be granted the
        /// `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
        /// 
        /// For more information see
        /// [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_bigquery_default_service_account.html.markdown.
        /// </summary>
        [Obsolete("Use GetDefaultServiceAccount.InvokeAsync() instead")]
        public static Task<GetDefaultServiceAccountResult> GetDefaultServiceAccount(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetDefaultServiceAccount
    {
        /// <summary>
        /// Get the email address of a project's unique BigQuery service account.
        /// 
        /// Each Google Cloud project has a unique service account used by BigQuery. When using
        /// BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
        /// this account needs to be granted the
        /// `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
        /// 
        /// For more information see
        /// [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_bigquery_default_service_account.html.markdown.
        /// </summary>
        public static Task<GetDefaultServiceAccountResult> InvokeAsync(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDefaultServiceAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project the unique service account was created for. If it is not provided, the provider project is used.
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
        /// The email address of the service account. This value is often used to refer to the service account
        /// in order to grant IAM permissions.
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
