// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source allows you to use data encrypted with Google Cloud KMS
 * within your resource definitions.
 * 
 * For more information see
 * [the official documentation](https://cloud.google.com/kms/docs/encrypt-decrypt).
 * 
 * > **NOTE**: Using this data provider will allow you to conceal secret data within your
 * resource definitions, but it does not take care of protecting that data in the
 * logging output, plan output, or state output.  Please take care to secure your secret
 * data outside of resource definitions.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/kms_secret.html.markdown.
 */
export function getKMSSecret(args: GetKMSSecretArgs, opts?: pulumi.InvokeOptions): Promise<GetKMSSecretResult> & GetKMSSecretResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetKMSSecretResult> = pulumi.runtime.invoke("gcp:kms/getKMSSecret:getKMSSecret", {
        "ciphertext": args.ciphertext,
        "cryptoKey": args.cryptoKey,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getKMSSecret.
 */
export interface GetKMSSecretArgs {
    /**
     * The ciphertext to be decrypted, encoded in base64
     */
    readonly ciphertext: string;
    /**
     * The id of the CryptoKey that will be used to
     * decrypt the provided ciphertext. This is represented by the format
     * `{projectId}/{location}/{keyRingName}/{cryptoKeyName}`.
     */
    readonly cryptoKey: string;
}

/**
 * A collection of values returned by getKMSSecret.
 */
export interface GetKMSSecretResult {
    readonly ciphertext: string;
    readonly cryptoKey: string;
    /**
     * Contains the result of decrypting the provided ciphertext.
     */
    readonly plaintext: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
