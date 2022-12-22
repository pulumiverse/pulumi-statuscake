// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Statuscake.Inputs
{

    public sealed class UptimeCheckTcpCheckArgs : Pulumi.ResourceArgs
    {
        [Input("authentication")]
        public Input<Inputs.UptimeCheckTcpCheckAuthenticationArgs>? Authentication { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public UptimeCheckTcpCheckArgs()
        {
        }
    }
}
