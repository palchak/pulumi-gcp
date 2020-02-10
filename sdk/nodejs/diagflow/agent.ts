// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dialogflow_agent.html.markdown.
 */
export class Agent extends pulumi.CustomResource {
    /**
     * Get an existing Agent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AgentState, opts?: pulumi.CustomResourceOptions): Agent {
        return new Agent(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:diagflow/agent:Agent';

    /**
     * Returns true if the given object is an instance of Agent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Agent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Agent.__pulumiType;
    }

    /**
     * API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
     * different service endpoints for different API versions. However, bots connectors and webhook calls will follow the
     * specified API version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1
     * API.
     */
    public readonly apiVersion!: pulumi.Output<string>;
    /**
     * The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into
     * this field, the Dialogflow will save the image in the backend. The address of the backend image returned from the
     * API will be shown in the [avatarUriBackend] field.
     */
    public readonly avatarUri!: pulumi.Output<string | undefined>;
    /**
     * The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar,
     * the [avatarUri] field can be used.
     */
    public /*out*/ readonly avatarUriBackend!: pulumi.Output<string>;
    /**
     * To filter out false positive results and still get variety in matched natural language inputs for your agent, you
     * can tune the machine learning classification threshold. If the returned score value is less than the threshold
     * value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
     * triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
     * default of 0.3 is used.
     */
    public readonly classificationThreshold!: pulumi.Output<number | undefined>;
    /**
     * The default language of the agent as a language tag. [See Language
     * Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
     * codes. This field cannot be updated after creation.
     */
    public readonly defaultLanguageCode!: pulumi.Output<string>;
    /**
     * The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of this agent.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Determines whether this agent should log conversation queries.
     */
    public readonly enableLogging!: pulumi.Output<boolean | undefined>;
    /**
     * Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
     * examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used
     * for agents with a large number of examples in intents, especially the ones using @sys.any or very large developer
     * entities.
     */
    public readonly matchMode!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The list of all languages supported by this agent (except for the defaultLanguageCode).
     */
    public readonly supportedLanguageCodes!: pulumi.Output<string[] | undefined>;
    /**
     * The agent tier. If not specified, TIER_STANDARD is assumed. * TIER_STANDARD: Standard tier. * TIER_ENTERPRISE:
     * Enterprise tier (Essentials). * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus). NOTE: This field seems to have
     * eventual consistency in the API. Updating this field to a new value, or even creating a new agent with a tier that
     * is different from a previous agent in the same project will take some time to propagate. The provider will wait for
     * the API to show consistency, which can lead to longer apply times.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
     * Europe/Paris.
     */
    public readonly timeZone!: pulumi.Output<string>;

    /**
     * Create a Agent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AgentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AgentArgs | AgentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AgentState | undefined;
            inputs["apiVersion"] = state ? state.apiVersion : undefined;
            inputs["avatarUri"] = state ? state.avatarUri : undefined;
            inputs["avatarUriBackend"] = state ? state.avatarUriBackend : undefined;
            inputs["classificationThreshold"] = state ? state.classificationThreshold : undefined;
            inputs["defaultLanguageCode"] = state ? state.defaultLanguageCode : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["enableLogging"] = state ? state.enableLogging : undefined;
            inputs["matchMode"] = state ? state.matchMode : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["supportedLanguageCodes"] = state ? state.supportedLanguageCodes : undefined;
            inputs["tier"] = state ? state.tier : undefined;
            inputs["timeZone"] = state ? state.timeZone : undefined;
        } else {
            const args = argsOrState as AgentArgs | undefined;
            if (!args || args.defaultLanguageCode === undefined) {
                throw new Error("Missing required property 'defaultLanguageCode'");
            }
            if (!args || args.displayName === undefined) {
                throw new Error("Missing required property 'displayName'");
            }
            if (!args || args.timeZone === undefined) {
                throw new Error("Missing required property 'timeZone'");
            }
            inputs["apiVersion"] = args ? args.apiVersion : undefined;
            inputs["avatarUri"] = args ? args.avatarUri : undefined;
            inputs["classificationThreshold"] = args ? args.classificationThreshold : undefined;
            inputs["defaultLanguageCode"] = args ? args.defaultLanguageCode : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["enableLogging"] = args ? args.enableLogging : undefined;
            inputs["matchMode"] = args ? args.matchMode : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["supportedLanguageCodes"] = args ? args.supportedLanguageCodes : undefined;
            inputs["tier"] = args ? args.tier : undefined;
            inputs["timeZone"] = args ? args.timeZone : undefined;
            inputs["avatarUriBackend"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Agent.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Agent resources.
 */
export interface AgentState {
    /**
     * API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
     * different service endpoints for different API versions. However, bots connectors and webhook calls will follow the
     * specified API version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1
     * API.
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into
     * this field, the Dialogflow will save the image in the backend. The address of the backend image returned from the
     * API will be shown in the [avatarUriBackend] field.
     */
    readonly avatarUri?: pulumi.Input<string>;
    /**
     * The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar,
     * the [avatarUri] field can be used.
     */
    readonly avatarUriBackend?: pulumi.Input<string>;
    /**
     * To filter out false positive results and still get variety in matched natural language inputs for your agent, you
     * can tune the machine learning classification threshold. If the returned score value is less than the threshold
     * value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
     * triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
     * default of 0.3 is used.
     */
    readonly classificationThreshold?: pulumi.Input<number>;
    /**
     * The default language of the agent as a language tag. [See Language
     * Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
     * codes. This field cannot be updated after creation.
     */
    readonly defaultLanguageCode?: pulumi.Input<string>;
    /**
     * The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of this agent.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Determines whether this agent should log conversation queries.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
     * examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used
     * for agents with a large number of examples in intents, especially the ones using @sys.any or very large developer
     * entities.
     */
    readonly matchMode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The list of all languages supported by this agent (except for the defaultLanguageCode).
     */
    readonly supportedLanguageCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The agent tier. If not specified, TIER_STANDARD is assumed. * TIER_STANDARD: Standard tier. * TIER_ENTERPRISE:
     * Enterprise tier (Essentials). * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus). NOTE: This field seems to have
     * eventual consistency in the API. Updating this field to a new value, or even creating a new agent with a tier that
     * is different from a previous agent in the same project will take some time to propagate. The provider will wait for
     * the API to show consistency, which can lead to longer apply times.
     */
    readonly tier?: pulumi.Input<string>;
    /**
     * The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
     * Europe/Paris.
     */
    readonly timeZone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Agent resource.
 */
export interface AgentArgs {
    /**
     * API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
     * different service endpoints for different API versions. However, bots connectors and webhook calls will follow the
     * specified API version. * API_VERSION_V1: Legacy V1 API. * API_VERSION_V2: V2 API. * API_VERSION_V2_BETA_1: V2beta1
     * API.
     */
    readonly apiVersion?: pulumi.Input<string>;
    /**
     * The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered into
     * this field, the Dialogflow will save the image in the backend. The address of the backend image returned from the
     * API will be shown in the [avatarUriBackend] field.
     */
    readonly avatarUri?: pulumi.Input<string>;
    /**
     * To filter out false positive results and still get variety in matched natural language inputs for your agent, you
     * can tune the machine learning classification threshold. If the returned score value is less than the threshold
     * value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
     * triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
     * default of 0.3 is used.
     */
    readonly classificationThreshold?: pulumi.Input<number>;
    /**
     * The default language of the agent as a language tag. [See Language
     * Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language
     * codes. This field cannot be updated after creation.
     */
    readonly defaultLanguageCode: pulumi.Input<string>;
    /**
     * The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of this agent.
     */
    readonly displayName: pulumi.Input<string>;
    /**
     * Determines whether this agent should log conversation queries.
     */
    readonly enableLogging?: pulumi.Input<boolean>;
    /**
     * Determines how intents are detected from user queries. * MATCH_MODE_HYBRID: Best for agents with a small number of
     * examples in intents and/or wide use of templates syntax and composite entities. * MATCH_MODE_ML_ONLY: Can be used
     * for agents with a large number of examples in intents, especially the ones using @sys.any or very large developer
     * entities.
     */
    readonly matchMode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The list of all languages supported by this agent (except for the defaultLanguageCode).
     */
    readonly supportedLanguageCodes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The agent tier. If not specified, TIER_STANDARD is assumed. * TIER_STANDARD: Standard tier. * TIER_ENTERPRISE:
     * Enterprise tier (Essentials). * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus). NOTE: This field seems to have
     * eventual consistency in the API. Updating this field to a new value, or even creating a new agent with a tier that
     * is different from a previous agent in the same project will take some time to propagate. The provider will wait for
     * the API to show consistency, which can lead to longer apply times.
     */
    readonly tier?: pulumi.Input<string>;
    /**
     * The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
     * Europe/Paris.
     */
    readonly timeZone: pulumi.Input<string>;
}
