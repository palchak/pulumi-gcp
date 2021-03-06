// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a TargetHttpsProxy resource, which is used by one or more
 * global forwarding rule to route incoming HTTPS requests to a URL map.
 * 
 * 
 * To get more information about TargetHttpsProxy, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetHttpsProxies)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/compute/docs/load-balancing/http/target-proxies)
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_target_https_proxy.html.markdown.
 */
export class TargetHttpsProxy extends pulumi.CustomResource {
    /**
     * Get an existing TargetHttpsProxy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetHttpsProxyState, opts?: pulumi.CustomResourceOptions): TargetHttpsProxy {
        return new TargetHttpsProxy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/targetHttpsProxy:TargetHttpsProxy';

    /**
     * Returns true if the given object is an instance of TargetHttpsProxy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TargetHttpsProxy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TargetHttpsProxy.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The unique identifier for the resource.
     */
    public /*out*/ readonly proxyId!: pulumi.Output<number>;
    /**
     * Specifies the QUIC override policy for this resource. This determines whether the load balancer will attempt to
     * negotiate QUIC with clients or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is specified, uses the QUIC
     * policy with no user overrides, which is equivalent to DISABLE. Not specifying this field is equivalent to specifying
     * NONE.
     */
    public readonly quicOverride!: pulumi.Output<string | undefined>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * A list of SslCertificate resources that are used to authenticate connections between users and the load balancer. At
     * least one SSL certificate must be specified.
     */
    public readonly sslCertificates!: pulumi.Output<string[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the
     * TargetHttpsProxy resource will not have any SSL policy configured.
     */
    public readonly sslPolicy!: pulumi.Output<string | undefined>;
    /**
     * A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
     */
    public readonly urlMap!: pulumi.Output<string>;

    /**
     * Create a TargetHttpsProxy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetHttpsProxyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetHttpsProxyArgs | TargetHttpsProxyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetHttpsProxyState | undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["proxyId"] = state ? state.proxyId : undefined;
            inputs["quicOverride"] = state ? state.quicOverride : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["sslCertificates"] = state ? state.sslCertificates : undefined;
            inputs["sslPolicy"] = state ? state.sslPolicy : undefined;
            inputs["urlMap"] = state ? state.urlMap : undefined;
        } else {
            const args = argsOrState as TargetHttpsProxyArgs | undefined;
            if (!args || args.sslCertificates === undefined) {
                throw new Error("Missing required property 'sslCertificates'");
            }
            if (!args || args.urlMap === undefined) {
                throw new Error("Missing required property 'urlMap'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["quicOverride"] = args ? args.quicOverride : undefined;
            inputs["sslCertificates"] = args ? args.sslCertificates : undefined;
            inputs["sslPolicy"] = args ? args.sslPolicy : undefined;
            inputs["urlMap"] = args ? args.urlMap : undefined;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["proxyId"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TargetHttpsProxy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TargetHttpsProxy resources.
 */
export interface TargetHttpsProxyState {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The unique identifier for the resource.
     */
    readonly proxyId?: pulumi.Input<number>;
    /**
     * Specifies the QUIC override policy for this resource. This determines whether the load balancer will attempt to
     * negotiate QUIC with clients or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is specified, uses the QUIC
     * policy with no user overrides, which is equivalent to DISABLE. Not specifying this field is equivalent to specifying
     * NONE.
     */
    readonly quicOverride?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    /**
     * A list of SslCertificate resources that are used to authenticate connections between users and the load balancer. At
     * least one SSL certificate must be specified.
     */
    readonly sslCertificates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the
     * TargetHttpsProxy resource will not have any SSL policy configured.
     */
    readonly sslPolicy?: pulumi.Input<string>;
    /**
     * A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
     */
    readonly urlMap?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TargetHttpsProxy resource.
 */
export interface TargetHttpsProxyArgs {
    /**
     * An optional description of this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long,
     * and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
     * '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * Specifies the QUIC override policy for this resource. This determines whether the load balancer will attempt to
     * negotiate QUIC with clients or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is specified, uses the QUIC
     * policy with no user overrides, which is equivalent to DISABLE. Not specifying this field is equivalent to specifying
     * NONE.
     */
    readonly quicOverride?: pulumi.Input<string>;
    /**
     * A list of SslCertificate resources that are used to authenticate connections between users and the load balancer. At
     * least one SSL certificate must be specified.
     */
    readonly sslCertificates: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A reference to the SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the
     * TargetHttpsProxy resource will not have any SSL policy configured.
     */
    readonly sslPolicy?: pulumi.Input<string>;
    /**
     * A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
     */
    readonly urlMap: pulumi.Input<string>;
}
