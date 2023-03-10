// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class SslCheck extends pulumi.CustomResource {
    /**
     * Get an existing SslCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SslCheckState, opts?: pulumi.CustomResourceOptions): SslCheck {
        return new SslCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'statuscake:index/sslCheck:SslCheck';

    /**
     * Returns true if the given object is an instance of SslCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SslCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SslCheck.__pulumiType;
    }

    /**
     * Alert configuration block
     */
    public readonly alertConfig!: pulumi.Output<outputs.SslCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    public readonly checkInterval!: pulumi.Output<number>;
    /**
     * List of contact group IDs
     */
    public readonly contactGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to follow redirects when testing. Disabled by default
     */
    public readonly followRedirects!: pulumi.Output<boolean | undefined>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    public readonly monitoredResource!: pulumi.Output<outputs.SslCheckMonitoredResource>;
    /**
     * Whether the check should be run
     */
    public readonly paused!: pulumi.Output<boolean | undefined>;
    /**
     * Custom user agent string set when testing
     */
    public readonly userAgent!: pulumi.Output<string | undefined>;

    /**
     * Create a SslCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SslCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SslCheckArgs | SslCheckState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SslCheckState | undefined;
            resourceInputs["alertConfig"] = state ? state.alertConfig : undefined;
            resourceInputs["checkInterval"] = state ? state.checkInterval : undefined;
            resourceInputs["contactGroups"] = state ? state.contactGroups : undefined;
            resourceInputs["followRedirects"] = state ? state.followRedirects : undefined;
            resourceInputs["monitoredResource"] = state ? state.monitoredResource : undefined;
            resourceInputs["paused"] = state ? state.paused : undefined;
            resourceInputs["userAgent"] = state ? state.userAgent : undefined;
        } else {
            const args = argsOrState as SslCheckArgs | undefined;
            if ((!args || args.alertConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alertConfig'");
            }
            if ((!args || args.checkInterval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'checkInterval'");
            }
            if ((!args || args.monitoredResource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitoredResource'");
            }
            resourceInputs["alertConfig"] = args ? args.alertConfig : undefined;
            resourceInputs["checkInterval"] = args ? args.checkInterval : undefined;
            resourceInputs["contactGroups"] = args ? args.contactGroups : undefined;
            resourceInputs["followRedirects"] = args ? args.followRedirects : undefined;
            resourceInputs["monitoredResource"] = args ? args.monitoredResource : undefined;
            resourceInputs["paused"] = args ? args.paused : undefined;
            resourceInputs["userAgent"] = args ? args.userAgent : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SslCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SslCheck resources.
 */
export interface SslCheckState {
    /**
     * Alert configuration block
     */
    alertConfig?: pulumi.Input<inputs.SslCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    checkInterval?: pulumi.Input<number>;
    /**
     * List of contact group IDs
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to follow redirects when testing. Disabled by default
     */
    followRedirects?: pulumi.Input<boolean>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    monitoredResource?: pulumi.Input<inputs.SslCheckMonitoredResource>;
    /**
     * Whether the check should be run
     */
    paused?: pulumi.Input<boolean>;
    /**
     * Custom user agent string set when testing
     */
    userAgent?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SslCheck resource.
 */
export interface SslCheckArgs {
    /**
     * Alert configuration block
     */
    alertConfig: pulumi.Input<inputs.SslCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    checkInterval: pulumi.Input<number>;
    /**
     * List of contact group IDs
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to follow redirects when testing. Disabled by default
     */
    followRedirects?: pulumi.Input<boolean>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    monitoredResource: pulumi.Input<inputs.SslCheckMonitoredResource>;
    /**
     * Whether the check should be run
     */
    paused?: pulumi.Input<boolean>;
    /**
     * Custom user agent string set when testing
     */
    userAgent?: pulumi.Input<string>;
}
