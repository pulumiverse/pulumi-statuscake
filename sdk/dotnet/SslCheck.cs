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
    [StatuscakeResourceType("statuscake:index/sslCheck:SslCheck")]
    public partial class SslCheck : Pulumi.CustomResource
    {
        /// <summary>
        /// Alert configuration block
        /// </summary>
        [Output("alertConfig")]
        public Output<Outputs.SslCheckAlertConfig> AlertConfig { get; private set; } = null!;

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
        /// Whether to follow redirects when testing. Disabled by default
        /// </summary>
        [Output("followRedirects")]
        public Output<bool?> FollowRedirects { get; private set; } = null!;

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Output("monitoredResource")]
        public Output<Outputs.SslCheckMonitoredResource> MonitoredResource { get; private set; } = null!;

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Output("paused")]
        public Output<bool?> Paused { get; private set; } = null!;

        /// <summary>
        /// Custom user agent string set when testing
        /// </summary>
        [Output("userAgent")]
        public Output<string?> UserAgent { get; private set; } = null!;


        /// <summary>
        /// Create a SslCheck resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SslCheck(string name, SslCheckArgs args, CustomResourceOptions? options = null)
            : base("statuscake:index/sslCheck:SslCheck", name, args ?? new SslCheckArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SslCheck(string name, Input<string> id, SslCheckState? state = null, CustomResourceOptions? options = null)
            : base("statuscake:index/sslCheck:SslCheck", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SslCheck resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SslCheck Get(string name, Input<string> id, SslCheckState? state = null, CustomResourceOptions? options = null)
        {
            return new SslCheck(name, id, state, options);
        }
    }

    public sealed class SslCheckArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert configuration block
        /// </summary>
        [Input("alertConfig", required: true)]
        public Input<Inputs.SslCheckAlertConfigArgs> AlertConfig { get; set; } = null!;

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
        /// Whether to follow redirects when testing. Disabled by default
        /// </summary>
        [Input("followRedirects")]
        public Input<bool>? FollowRedirects { get; set; }

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Input("monitoredResource", required: true)]
        public Input<Inputs.SslCheckMonitoredResourceArgs> MonitoredResource { get; set; } = null!;

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// Custom user agent string set when testing
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public SslCheckArgs()
        {
        }
    }

    public sealed class SslCheckState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Alert configuration block
        /// </summary>
        [Input("alertConfig")]
        public Input<Inputs.SslCheckAlertConfigGetArgs>? AlertConfig { get; set; }

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
        /// Whether to follow redirects when testing. Disabled by default
        /// </summary>
        [Input("followRedirects")]
        public Input<bool>? FollowRedirects { get; set; }

        /// <summary>
        /// Monitored resource configuration block. The describes server under test
        /// </summary>
        [Input("monitoredResource")]
        public Input<Inputs.SslCheckMonitoredResourceGetArgs>? MonitoredResource { get; set; }

        /// <summary>
        /// Whether the check should be run
        /// </summary>
        [Input("paused")]
        public Input<bool>? Paused { get; set; }

        /// <summary>
        /// Custom user agent string set when testing
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public SslCheckState()
        {
        }
    }
}