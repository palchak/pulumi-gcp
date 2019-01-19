// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * An attestor that attests to container image artifacts.
 * 
 * > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
 * See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
 * 
 * To get more information about Attestor, see:
 * 
 * * [API documentation](https://cloud.google.com/binary-authorization/docs/reference/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/binary-authorization/)
 * 
 * ## Example Usage - Binary Authorization Attestor Basic
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * 
 * const google_container_analysis_note_note = new gcp.containeranalysis.Note("note", {
 *     attestationAuthority: {
 *         hint: {
 *             humanReadableName: "Attestor Note",
 *         },
 *     },
 *     name: "test-attestor-note",
 * });
 * const google_binary_authorization_attestor_attestor = new gcp.binaryauthorization.Attestor("attestor", {
 *     attestationAuthorityNote: {
 *         noteReference: google_container_analysis_note_note.name,
 *         publicKeys: [{
 *             asciiArmoredPgpPublicKey: "mQENBFtP0doBCADF+joTiXWKVuP8kJt3fgpBSjT9h8ezMfKA4aXZctYLx5wslWQl\nbB7Iu2ezkECNzoEeU7WxUe8a61pMCh9cisS9H5mB2K2uM4Jnf8tgFeXn3akJDVo0\noR1IC+Dp9mXbRSK3MAvKkOwWlG99sx3uEdvmeBRHBOO+grchLx24EThXFOyP9Fk6\nV39j6xMjw4aggLD15B4V0v9JqBDdJiIYFzszZDL6pJwZrzcP0z8JO4rTZd+f64bD\nMpj52j/pQfA8lZHOaAgb1OrthLdMrBAjoDjArV4Ek7vSbrcgYWcI6BhsQrFoxKdX\n83TZKai55ZCfCLIskwUIzA1NLVwyzCS+fSN/ABEBAAG0KCJUZXN0IEF0dGVzdG9y\nIiA8ZGFuYWhvZmZtYW5AZ29vZ2xlLmNvbT6JAU4EEwEIADgWIQRfWkqHt6hpTA1L\nuY060eeM4dc66AUCW0/R2gIbLwULCQgHAgYVCgkICwIEFgIDAQIeAQIXgAAKCRA6\n0eeM4dc66HdpCAC4ot3b0OyxPb0Ip+WT2U0PbpTBPJklesuwpIrM4Lh0N+1nVRLC\n51WSmVbM8BiAFhLbN9LpdHhds1kUrHF7+wWAjdR8sqAj9otc6HGRM/3qfa2qgh+U\nWTEk/3us/rYSi7T7TkMuutRMIa1IkR13uKiW56csEMnbOQpn9rDqwIr5R8nlZP5h\nMAU9vdm1DIv567meMqTaVZgR3w7bck2P49AO8lO5ERFpVkErtu/98y+rUy9d789l\n+OPuS1NGnxI1YKsNaWJF4uJVuvQuZ1twrhCbGNtVorO2U12+cEq+YtUxj7kmdOC1\nqoIRW6y0+UlAc+MbqfL0ziHDOAmcqz1GnROg\n=6Bvm\n",
 *         }],
 *     },
 *     name: "test-attestor",
 * });
 * ```
 */
export class Attestor extends pulumi.CustomResource {
    /**
     * Get an existing Attestor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttestorState, opts?: pulumi.CustomResourceOptions): Attestor {
        return new Attestor(name, <any>state, { ...opts, id: id });
    }

    public readonly attestationAuthorityNote: pulumi.Output<{ delegationServiceAccountEmail: string, noteReference: string, publicKeys?: { asciiArmoredPgpPublicKey: string, comment?: string, id: string }[] }>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly project: pulumi.Output<string>;

    /**
     * Create a Attestor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttestorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttestorArgs | AttestorState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: AttestorState = argsOrState as AttestorState | undefined;
            inputs["attestationAuthorityNote"] = state ? state.attestationAuthorityNote : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as AttestorArgs | undefined;
            if (!args || args.attestationAuthorityNote === undefined) {
                throw new Error("Missing required property 'attestationAuthorityNote'");
            }
            inputs["attestationAuthorityNote"] = args ? args.attestationAuthorityNote : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        }
        super("gcp:binaryauthorization/attestor:Attestor", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Attestor resources.
 */
export interface AttestorState {
    readonly attestationAuthorityNote?: pulumi.Input<{ delegationServiceAccountEmail?: pulumi.Input<string>, noteReference: pulumi.Input<string>, publicKeys?: pulumi.Input<pulumi.Input<{ asciiArmoredPgpPublicKey: pulumi.Input<string>, comment?: pulumi.Input<string>, id?: pulumi.Input<string> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Attestor resource.
 */
export interface AttestorArgs {
    readonly attestationAuthorityNote: pulumi.Input<{ delegationServiceAccountEmail?: pulumi.Input<string>, noteReference: pulumi.Input<string>, publicKeys?: pulumi.Input<pulumi.Input<{ asciiArmoredPgpPublicKey: pulumi.Input<string>, comment?: pulumi.Input<string>, id?: pulumi.Input<string> }>[]> }>;
    readonly description?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly project?: pulumi.Input<string>;
}