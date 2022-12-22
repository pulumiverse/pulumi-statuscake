// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the statuscake package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'statuscake';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * The API token for operations. This can also be provided as an environment variable `STATUSCAKE_API_TOKEN`
     */
    public readonly apiToken!: pulumi.Output<string | undefined>;
    /**
     * Custom endpoint to which request will be made. This can also be provided as an environment variable
     * `STATUCAKE_CUSTOM_ENDPOINT`
     */
    public readonly statuscakeCustomEndpoint!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["apiToken"] = (args?.apiToken ? pulumi.secret(args.apiToken) : undefined) ?? utilities.getEnv("STATUSCAKE_API_TOKEN");
            resourceInputs["maxBackoff"] = pulumi.output(args ? args.maxBackoff : undefined).apply(JSON.stringify);
            resourceInputs["minBackoff"] = pulumi.output(args ? args.minBackoff : undefined).apply(JSON.stringify);
            resourceInputs["retries"] = pulumi.output(args ? args.retries : undefined).apply(JSON.stringify);
            resourceInputs["rps"] = pulumi.output(args ? args.rps : undefined).apply(JSON.stringify);
            resourceInputs["statuscakeCustomEndpoint"] = args ? args.statuscakeCustomEndpoint : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["apiToken"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The API token for operations. This can also be provided as an environment variable `STATUSCAKE_API_TOKEN`
     */
    apiToken?: pulumi.Input<string>;
    /**
     * Maximum backoff period in seconds after failed API calls. This can also be provided as an environment variable
     * `STATUSCAKE_MAX_BACKOFF`
     */
    maxBackoff?: pulumi.Input<number>;
    /**
     * Minimum backoff period in seconds after failed API calls. This can also be provided as an environment variable
     * `STATUSCAKE_MIN_BACKOFF`
     */
    minBackoff?: pulumi.Input<number>;
    /**
     * Maximum number of retries to perform when an API request fails. This can also be provided as an environment variable
     * `STATUSCAKE_RETRIES`
     */
    retries?: pulumi.Input<number>;
    /**
     * RPS limit to apply when making calls to the API. This can also be provided as an environment variable `STATUSCAKE_RPS`
     */
    rps?: pulumi.Input<number>;
    /**
     * Custom endpoint to which request will be made. This can also be provided as an environment variable
     * `STATUCAKE_CUSTOM_ENDPOINT`
     */
    statuscakeCustomEndpoint?: pulumi.Input<string>;
}