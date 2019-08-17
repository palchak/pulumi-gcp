// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents an Image resource.
 * 
 * Google Compute Engine uses operating system images to create the root
 * persistent disks for your instances. You specify an image when you create
 * an instance. Images contain a boot loader, an operating system, and a
 * root file system. Linux operating system images are also capable of
 * running containers on Compute Engine.
 * 
 * Images can be either public or custom.
 * 
 * Public images are provided and maintained by Google, open-source
 * communities, and third-party vendors. By default, all projects have
 * access to these images and can use them to create instances.  Custom
 * images are available only to your project. You can create a custom image
 * from root persistent disks and other images. Then, use the custom image
 * to create an instance.
 * 
 * 
 * To get more information about Image, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/images)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/images)
 * 
 * ## Example Usage - Image Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const example = new gcp.compute.Image("example", {
 *     rawDisk: {
 *         source: "https://storage.googleapis.com/bosh-cpi-artifacts/bosh-stemcell-3262.4-google-kvm-ubuntu-trusty-go_agent-raw.tar.gz",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_image.html.markdown.
 */
export class Image extends pulumi.CustomResource {
    /**
     * Get an existing Image resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ImageState, opts?: pulumi.CustomResourceOptions): Image {
        return new Image(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/image:Image';

    /**
     * Returns true if the given object is an instance of Image.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Image {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Image.__pulumiType;
    }

    public /*out*/ readonly archiveSizeBytes!: pulumi.Output<number>;
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly diskSizeGb!: pulumi.Output<number>;
    public readonly family!: pulumi.Output<string | undefined>;
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly licenses!: pulumi.Output<string[]>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    public readonly rawDisk!: pulumi.Output<outputs.compute.ImageRawDisk | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    public readonly sourceDisk!: pulumi.Output<string | undefined>;

    /**
     * Create a Image resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ImageArgs | ImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ImageState | undefined;
            inputs["archiveSizeBytes"] = state ? state.archiveSizeBytes : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["diskSizeGb"] = state ? state.diskSizeGb : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["licenses"] = state ? state.licenses : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["rawDisk"] = state ? state.rawDisk : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sourceDisk"] = state ? state.sourceDisk : undefined;
        } else {
            const args = argsOrState as ImageArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["diskSizeGb"] = args ? args.diskSizeGb : undefined;
            inputs["family"] = args ? args.family : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["licenses"] = args ? args.licenses : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["rawDisk"] = args ? args.rawDisk : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["archiveSizeBytes"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Image.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Image resources.
 */
export interface ImageState {
    readonly archiveSizeBytes?: pulumi.Input<number>;
    readonly creationTimestamp?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly diskSizeGb?: pulumi.Input<number>;
    readonly family?: pulumi.Input<string>;
    readonly labelFingerprint?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly licenses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly rawDisk?: pulumi.Input<inputs.compute.ImageRawDisk>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly sourceDisk?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Image resource.
 */
export interface ImageArgs {
    readonly description?: pulumi.Input<string>;
    readonly diskSizeGb?: pulumi.Input<number>;
    readonly family?: pulumi.Input<string>;
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly licenses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    readonly rawDisk?: pulumi.Input<inputs.compute.ImageRawDisk>;
    readonly sourceDisk?: pulumi.Input<string>;
}
