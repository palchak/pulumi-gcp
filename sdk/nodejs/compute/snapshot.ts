// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a new snapshot of a disk within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/disks/create-snapshots)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/snapshots).
 */
export class Snapshot extends pulumi.CustomResource {
    /**
     * Get an existing Snapshot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SnapshotState): Snapshot {
        return new Snapshot(name, <any>state, { id });
    }

    /**
     * The unique fingerprint of the labels.
     */
    public /*out*/ readonly labelFingerprint: pulumi.Output<string>;
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    public readonly labels: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to encrypt this snapshot.
     */
    public readonly snapshotEncryptionKeyRaw: pulumi.Output<string | undefined>;
    /**
     * The [RFC 4648 base64]
     * (https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the
     * [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects this resource.
     */
    public /*out*/ readonly snapshotEncryptionKeySha256: pulumi.Output<string>;
    /**
     * The disk which will be used as the source of the snapshot.
     */
    public readonly sourceDisk: pulumi.Output<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to decrypt the source disk.
     */
    public readonly sourceDiskEncryptionKeyRaw: pulumi.Output<string | undefined>;
    /**
     * The [RFC 4648 base64]
     * (https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the
     * [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects the source disk.
     */
    public /*out*/ readonly sourceDiskEncryptionKeySha256: pulumi.Output<string>;
    /**
     * The URI of the source disk.
     */
    public /*out*/ readonly sourceDiskLink: pulumi.Output<string>;
    /**
     * The zone where the source disk is located.
     */
    public readonly zone: pulumi.Output<string>;

    /**
     * Create a Snapshot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SnapshotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SnapshotArgs | SnapshotState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SnapshotState = argsOrState as SnapshotState | undefined;
            inputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["snapshotEncryptionKeyRaw"] = state ? state.snapshotEncryptionKeyRaw : undefined;
            inputs["snapshotEncryptionKeySha256"] = state ? state.snapshotEncryptionKeySha256 : undefined;
            inputs["sourceDisk"] = state ? state.sourceDisk : undefined;
            inputs["sourceDiskEncryptionKeyRaw"] = state ? state.sourceDiskEncryptionKeyRaw : undefined;
            inputs["sourceDiskEncryptionKeySha256"] = state ? state.sourceDiskEncryptionKeySha256 : undefined;
            inputs["sourceDiskLink"] = state ? state.sourceDiskLink : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as SnapshotArgs | undefined;
            if (!args || args.sourceDisk === undefined) {
                throw new Error("Missing required property 'sourceDisk'");
            }
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["snapshotEncryptionKeyRaw"] = args ? args.snapshotEncryptionKeyRaw : undefined;
            inputs["sourceDisk"] = args ? args.sourceDisk : undefined;
            inputs["sourceDiskEncryptionKeyRaw"] = args ? args.sourceDiskEncryptionKeyRaw : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["snapshotEncryptionKeySha256"] = undefined /*out*/;
            inputs["sourceDiskEncryptionKeySha256"] = undefined /*out*/;
            inputs["sourceDiskLink"] = undefined /*out*/;
        }
        super("gcp:compute/snapshot:Snapshot", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Snapshot resources.
 */
export interface SnapshotState {
    /**
     * The unique fingerprint of the labels.
     */
    readonly labelFingerprint?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to encrypt this snapshot.
     */
    readonly snapshotEncryptionKeyRaw?: pulumi.Input<string>;
    /**
     * The [RFC 4648 base64]
     * (https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the
     * [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects this resource.
     */
    readonly snapshotEncryptionKeySha256?: pulumi.Input<string>;
    /**
     * The disk which will be used as the source of the snapshot.
     */
    readonly sourceDisk?: pulumi.Input<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to decrypt the source disk.
     */
    readonly sourceDiskEncryptionKeyRaw?: pulumi.Input<string>;
    /**
     * The [RFC 4648 base64]
     * (https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the
     * [customer-supplied encryption key](https://cloud.google.com/compute/docs/disks/customer-supplied-encryption)
     * that protects the source disk.
     */
    readonly sourceDiskEncryptionKeySha256?: pulumi.Input<string>;
    /**
     * The URI of the source disk.
     */
    readonly sourceDiskLink?: pulumi.Input<string>;
    /**
     * The zone where the source disk is located.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Snapshot resource.
 */
export interface SnapshotArgs {
    /**
     * A set of key/value label pairs to assign to the snapshot.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to encrypt this snapshot.
     */
    readonly snapshotEncryptionKeyRaw?: pulumi.Input<string>;
    /**
     * The disk which will be used as the source of the snapshot.
     */
    readonly sourceDisk: pulumi.Input<string>;
    /**
     * A 256-bit [customer-supplied encryption key]
     * (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption),
     * encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4)
     * to decrypt the source disk.
     */
    readonly sourceDiskEncryptionKeyRaw?: pulumi.Input<string>;
    /**
     * The zone where the source disk is located.
     */
    readonly zone?: pulumi.Input<string>;
}
