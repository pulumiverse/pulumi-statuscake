// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Statuscake
{
    [StatuscakeResourceType("statuscake:index/statuscakePagespeedCheck:StatuscakePagespeedCheck")]
    public partial class StatuscakePagespeedCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// Alert configuration block. Omitting this block disabled all alerts
        /// </summary>
        [Output("alertConfig")]
        public Output<Outputs.StatuscakePagespeedCheckAlertConfig> AlertConfig { get; private set; } = null!;

        /// <summary>
        /// Number of seconds between checks
        /// </summary>
        [Output("checkInterval")]
        public Output<int> CheckInterval { get; private set; } = null!;

        /// <summary>
        /// List of contact group IDs
        /// </summary>
        [Output("contactGroups")]
        public Output<ImmutableArray<string>> ContactGroups { get; private set; } = null!;

        /// <summary>
        /// Assigned monitoring location on which checks will be run
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Output("monitoredResource")]
        public Output<Outputs.StatuscakePagespeedCheckMonitoredResource> MonitoredResource { get; private set; } = null!;

        /// <summary>
        /// Name of the check
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Output("paused")]
        public Output<bool?> Paused { get; private set; } = null!;

        /// <summary>
        /// Region on which to run checks
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a StatuscakePagespeedCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StatuscakePagespeedCheck(string name, StatuscakePagespeedCheckArgs args, CustomResourceOptions? options = null)
            : base("statuscake:index/statuscakePagespeedCheck:StatuscakePagespeedCheck", name, args ?? new StatuscakePagespeedCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StatuscakePagespeedCheck(string name, Input<string> id, StatuscakePagespeedCheckState? state = null, CustomResourceOptions? options = null)
            : base("statuscake:index/statuscakePagespeedCheck:StatuscakePagespeedCheck", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StatuscakePagespeedCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StatuscakePagespeedCheck Get(string name, Input<string> id, StatuscakePagespeedCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new StatuscakePagespeedCheck(name, id, state, options);
        }
    }

    public sealed class StatuscakePagespeedCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert configuration block. Omitting this block disabled all alerts
        /// </summary>
        [Input("alertConfig", required: true)]
        public Input<Inputs.StatuscakePagespeedCheckAlertConfigArgs> AlertConfig { get; set; } = null!;

        /// <summary>
        /// Number of seconds between checks
        /// </summary>
        [Input("checkInterval", required: true)]
        public Input<int> CheckInterval { get; set; } = null!;

        [Input("contactGroups")]
        private InputList<string>? _contactGroups;

        /// <summary>
        /// List of contact group IDs
        /// </summary>
        public InputList<string> ContactGroups
        {
            get => _contactGroups ?? (_contactGroups = new InputList<string>());
            set => _contactGroups = value;
        }

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Input("monitoredResource", required: true)]
        public Input<Inputs.StatuscakePagespeedCheckMonitoredResourceArgs> MonitoredResource { get; set; } = null!;

        /// <summary>
        /// Name of the check
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// Region on which to run checks
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public StatuscakePagespeedCheckArgs()
        {
        }
    }

    public sealed class StatuscakePagespeedCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert configuration block. Omitting this block disabled all alerts
        /// </summary>
        [Input("alertConfig")]
        public Input<Inputs.StatuscakePagespeedCheckAlertConfigGetArgs>? AlertConfig { get; set; }

        /// <summary>
        /// Number of seconds between checks
        /// </summary>
        [Input("checkInterval")]
        public Input<int>? CheckInterval { get; set; }

        [Input("contactGroups")]
        private InputList<string>? _contactGroups;

        /// <summary>
        /// List of contact group IDs
        /// </summary>
        public InputList<string> ContactGroups
        {
            get => _contactGroups ?? (_contactGroups = new InputList<string>());
            set => _contactGroups = value;
        }

        /// <summary>
        /// Assigned monitoring location on which checks will be run
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.StatuscakePagespeedCheckMonitoredResourceGetArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// Name of the check
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// Region on which to run checks
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public StatuscakePagespeedCheckState()
        {
        }
    }
}
