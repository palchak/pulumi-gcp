// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get OpenID userinfo about the credentials used with the Google provider,
 * specifically the email.
 * 
 * When the `https://www.googleapis.com/auth/userinfo.email` scope is enabled in
 * your provider block, this datasource enables you to export the email of the
 * account you've authenticated the provider with; this can be used alongside
 * `data.google_client_config`'s `accessToken` to perform OpenID Connect
 * authentication with GKE and configure an RBAC role for the email used.
 * 
 * > This resource will only work as expected if the provider is configured to
 * use the `https://www.googleapis.com/auth/userinfo.email` scope! You will
 * receive an error otherwise.
 * 
 * ## Example Usage - exporting an email
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const me = gcp.organizations.getClientOpenIdUserInfo({});
 * 
 * export const myEmail = me.email;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/client_openid_userinfo.html.markdown.
 */
export function getClientOpenIdUserInfo(opts?: pulumi.InvokeOptions): Promise<GetClientOpenIdUserInfoResult> & GetClientOpenIdUserInfoResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetClientOpenIdUserInfoResult> = pulumi.runtime.invoke("gcp:organizations/getClientOpenIdUserInfo:getClientOpenIdUserInfo", {
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of values returned by getClientOpenIdUserInfo.
 */
export interface GetClientOpenIdUserInfoResult {
    /**
     * The email of the account used by the provider to authenticate with GCP.
     */
    readonly email: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
