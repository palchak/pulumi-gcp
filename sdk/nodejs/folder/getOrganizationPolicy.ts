// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows management of Organization policies for a Google Folder. For more information see
 * [the official
 * documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const policy = gcp.folder.getOrganizationPolicy({
 *     constraint: "constraints/compute.trustedImageProjects",
 *     folder: "folders/folderid",
 * });
 * 
 * export const version = policy.version;
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/folder_organization_policy.html.markdown.
 */
export function getOrganizationPolicy(args: GetOrganizationPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationPolicyResult> & GetOrganizationPolicyResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetOrganizationPolicyResult> = pulumi.runtime.invoke("gcp:folder/getOrganizationPolicy:getOrganizationPolicy", {
        "constraint": args.constraint,
        "folder": args.folder,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getOrganizationPolicy.
 */
export interface GetOrganizationPolicyArgs {
    /**
     * (Required) The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
     */
    readonly constraint: string;
    /**
     * The resource name of the folder to set the policy for. Its format is folders/{folder_id}.
     */
    readonly folder: string;
}

/**
 * A collection of values returned by getOrganizationPolicy.
 */
export interface GetOrganizationPolicyResult {
    readonly booleanPolicies: outputs.folder.GetOrganizationPolicyBooleanPolicy[];
    readonly constraint: string;
    readonly etag: string;
    readonly folder: string;
    readonly listPolicies: outputs.folder.GetOrganizationPolicyListPolicy[];
    readonly restorePolicies: outputs.folder.GetOrganizationPolicyRestorePolicy[];
    readonly updateTime: string;
    readonly version: number;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
