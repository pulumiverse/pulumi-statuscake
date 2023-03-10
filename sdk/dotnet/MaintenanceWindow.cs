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
    [StatuscakeResourceType("statuscake:index/maintenanceWindow:MaintenanceWindow")]
    public partial class MaintenanceWindow : Pulumi.CustomResource
    {
        /// <summary>
        /// End of the maintenance window (RFC3339 format)
        /// </summary>
        [Output("end")]
        public Output<string> End { get; private set; } = null!;

        /// <summary>
        /// Name of the maintenance window
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// How often the maintenance window should occur
        /// </summary>
        [Output("repeatInterval")]
        public Output<string?> RepeatInterval { get; private set; } = null!;

        /// <summary>
        /// Start of the maintenance window (RFC3339 format)
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;

        /// <summary>
        /// List of tags used to include matching uptime checks in this maintenance window
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// List of uptime check IDs explicitly included in this maintenance window
        /// </summary>
        [Output("tests")]
        public Output<ImmutableArray<string>> Tests { get; private set; } = null!;

        /// <summary>
        /// Standard timezone associated with this maintenance window
        /// </summary>
        [Output("timezone")]
        public Output<string> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a MaintenanceWindow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MaintenanceWindow(string name, MaintenanceWindowArgs args, CustomResourceOptions? options = null)
            : base("statuscake:index/maintenanceWindow:MaintenanceWindow", name, args ?? new MaintenanceWindowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MaintenanceWindow(string name, Input<string> id, MaintenanceWindowState? state = null, CustomResourceOptions? options = null)
            : base("statuscake:index/maintenanceWindow:MaintenanceWindow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MaintenanceWindow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MaintenanceWindow Get(string name, Input<string> id, MaintenanceWindowState? state = null, CustomResourceOptions? options = null)
        {
            return new MaintenanceWindow(name, id, state, options);
        }
    }

    public sealed class MaintenanceWindowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of the maintenance window (RFC3339 format)
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// Name of the maintenance window
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often the maintenance window should occur
        /// </summary>
        [Input("repeatInterval")]
        public Input<string>? RepeatInterval { get; set; }

        /// <summary>
        /// Start of the maintenance window (RFC3339 format)
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags used to include matching uptime checks in this maintenance window
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tests")]
        private InputList<string>? _tests;

        /// <summary>
        /// List of uptime check IDs explicitly included in this maintenance window
        /// </summary>
        public InputList<string> Tests
        {
            get => _tests ?? (_tests = new InputList<string>());
            set => _tests = value;
        }

        /// <summary>
        /// Standard timezone associated with this maintenance window
        /// </summary>
        [Input("timezone", required: true)]
        public Input<string> Timezone { get; set; } = null!;

        public MaintenanceWindowArgs()
        {
        }
    }

    public sealed class MaintenanceWindowState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of the maintenance window (RFC3339 format)
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// Name of the maintenance window
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// How often the maintenance window should occur
        /// </summary>
        [Input("repeatInterval")]
        public Input<string>? RepeatInterval { get; set; }

        /// <summary>
        /// Start of the maintenance window (RFC3339 format)
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags used to include matching uptime checks in this maintenance window
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("tests")]
        private InputList<string>? _tests;

        /// <summary>
        /// List of uptime check IDs explicitly included in this maintenance window
        /// </summary>
        public InputList<string> Tests
        {
            get => _tests ?? (_tests = new InputList<string>());
            set => _tests = value;
        }

        /// <summary>
        /// Standard timezone associated with this maintenance window
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public MaintenanceWindowState()
        {
        }
    }
}
