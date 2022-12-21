// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class StatuscakePagespeedCheck extends pulumi.CustomResource {
    /**
     * Get an existing StatuscakePagespeedCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StatuscakePagespeedCheckState, opts?: pulumi.CustomResourceOptions): StatuscakePagespeedCheck {
        return new StatuscakePagespeedCheck(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'statuscake:index/statuscakePagespeedCheck:StatuscakePagespeedCheck';

    /**
     * Returns true if the given object is an instance of StatuscakePagespeedCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StatuscakePagespeedCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StatuscakePagespeedCheck.__pulumiType;
    }

    /**
     * Alert configuration block. Omitting this block disabled all alerts
     */
    public readonly alertConfig!: pulumi.Output<outputs.StatuscakePagespeedCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    public readonly checkInterval!: pulumi.Output<number>;
    /**
     * List of contact group IDs
     */
    public readonly contactGroups!: pulumi.Output<string[] | undefined>;
    /**
     * Assigned monitoring location on which checks will be run
     */
    public /*out*/ readonly location!: pulumi.Output<string>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    public readonly monitoredResource!: pulumi.Output<outputs.StatuscakePagespeedCheckMonitoredResource>;
    /**
     * Name of the check
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether the check should be run
     */
    public readonly paused!: pulumi.Output<boolean | undefined>;
    /**
     * Region on which to run checks
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a StatuscakePagespeedCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StatuscakePagespeedCheckArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StatuscakePagespeedCheckArgs | StatuscakePagespeedCheckState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StatuscakePagespeedCheckState | undefined;
            resourceInputs["alertConfig"] = state ? state.alertConfig : undefined;
            resourceInputs["checkInterval"] = state ? state.checkInterval : undefined;
            resourceInputs["contactGroups"] = state ? state.contactGroups : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["monitoredResource"] = state ? state.monitoredResource : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paused"] = state ? state.paused : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as StatuscakePagespeedCheckArgs | undefined;
            if ((!args || args.alertConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alertConfig'");
            }
            if ((!args || args.checkInterval === undefined) && !opts.urn) {
                throw new Error("Missing required property 'checkInterval'");
            }
            if ((!args || args.monitoredResource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'monitoredResource'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["alertConfig"] = args ? args.alertConfig : undefined;
            resourceInputs["checkInterval"] = args ? args.checkInterval : undefined;
            resourceInputs["contactGroups"] = args ? args.contactGroups : undefined;
            resourceInputs["monitoredResource"] = args ? args.monitoredResource : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paused"] = args ? args.paused : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["location"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StatuscakePagespeedCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StatuscakePagespeedCheck resources.
 */
export interface StatuscakePagespeedCheckState {
    /**
     * Alert configuration block. Omitting this block disabled all alerts
     */
    alertConfig?: pulumi.Input<inputs.StatuscakePagespeedCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    checkInterval?: pulumi.Input<number>;
    /**
     * List of contact group IDs
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Assigned monitoring location on which checks will be run
     */
    location?: pulumi.Input<string>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    monitoredResource?: pulumi.Input<inputs.StatuscakePagespeedCheckMonitoredResource>;
    /**
     * Name of the check
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the check should be run
     */
    paused?: pulumi.Input<boolean>;
    /**
     * Region on which to run checks
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StatuscakePagespeedCheck resource.
 */
export interface StatuscakePagespeedCheckArgs {
    /**
     * Alert configuration block. Omitting this block disabled all alerts
     */
    alertConfig: pulumi.Input<inputs.StatuscakePagespeedCheckAlertConfig>;
    /**
     * Number of seconds between checks
     */
    checkInterval: pulumi.Input<number>;
    /**
     * List of contact group IDs
     */
    contactGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Monitored resource configuration block. The describes server under test
     */
    monitoredResource: pulumi.Input<inputs.StatuscakePagespeedCheckMonitoredResource>;
    /**
     * Name of the check
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the check should be run
     */
    paused?: pulumi.Input<boolean>;
    /**
     * Region on which to run checks
     */
    region: pulumi.Input<string>;
}
