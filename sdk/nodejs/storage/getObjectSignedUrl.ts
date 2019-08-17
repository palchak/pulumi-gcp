// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The Google Cloud storage signed URL data source generates a signed URL for a given storage object. Signed URLs provide a way to give time-limited read or write access to anyone in possession of the URL, regardless of whether they have a Google account.
 * 
 * For more info about signed URL's is available [here](https://cloud.google.com/storage/docs/access-control/signed-urls).
 * 
 * ## Full Example
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const getUrl = gcp.storage.getObjectSignedUrl({
 *     bucket: "friedChicken",
 *     contentMd5: "pRviqwS4c4OTJRTe03FD1w==",
 *     contentType: "text/plain",
 *     credentials: fs.readFileSync("path/to/credentials.json", "utf-8"),
 *     duration: "2d",
 *     extensionHeaders: {
 *         "x-goog-if-generation-match": 1,
 *     },
 *     path: "path/to/file",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_object_signed_url.html.markdown.
 */
export function getObjectSignedUrl(args: GetObjectSignedUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectSignedUrlResult> & GetObjectSignedUrlResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetObjectSignedUrlResult> = pulumi.runtime.invoke("gcp:storage/getObjectSignedUrl:getObjectSignedUrl", {
        "bucket": args.bucket,
        "contentMd5": args.contentMd5,
        "contentType": args.contentType,
        "credentials": args.credentials,
        "duration": args.duration,
        "extensionHeaders": args.extensionHeaders,
        "httpMethod": args.httpMethod,
        "path": args.path,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getObjectSignedUrl.
 */
export interface GetObjectSignedUrlArgs {
    /**
     * The name of the bucket to read the object from
     */
    readonly bucket: string;
    /**
     * The [MD5 digest](https://cloud.google.com/storage/docs/hashes-etags#_MD5) value in Base64.
     * Typically retrieved from `google_storage_bucket_object.object.md5hash` attribute.
     * If you provide this in the datasource, the client (e.g. browser, curl) must provide the `Content-MD5` HTTP header with this same value in its request.
     */
    readonly contentMd5?: string;
    /**
     * If you specify this in the datasource, the client must provide the `Content-Type` HTTP header with the same value in its request.
     */
    readonly contentType?: string;
    /**
     * What Google service account credentials json should be used to sign the URL. 
     * This data source checks the following locations for credentials, in order of preference: data source `credentials` attribute, provider `credentials` attribute and finally the GOOGLE_APPLICATION_CREDENTIALS environment variable.
     */
    readonly credentials?: string;
    /**
     * For how long shall the signed URL be valid (defaults to 1 hour - i.e. `1h`). 
     * See [here](https://golang.org/pkg/time/#ParseDuration) for info on valid duration formats.
     */
    readonly duration?: string;
    /**
     * As needed. The server checks to make sure that the client provides matching values in requests using the signed URL. 
     * Any header starting with `x-goog-` is accepted but see the [Google Docs](https://cloud.google.com/storage/docs/xml-api/reference-headers) for list of headers that are supported by Google.
     */
    readonly extensionHeaders?: {[key: string]: string};
    /**
     * What HTTP Method will the signed URL allow (defaults to `GET`)
     */
    readonly httpMethod?: string;
    /**
     * The full path to the object inside the bucket
     */
    readonly path: string;
}

/**
 * A collection of values returned by getObjectSignedUrl.
 */
export interface GetObjectSignedUrlResult {
    readonly bucket: string;
    readonly contentMd5?: string;
    readonly contentType?: string;
    readonly credentials?: string;
    readonly duration?: string;
    readonly extensionHeaders?: {[key: string]: string};
    readonly httpMethod?: string;
    readonly path: string;
    /**
     * The signed URL that can be used to access the storage object without authentication.
     */
    readonly signedUrl: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
