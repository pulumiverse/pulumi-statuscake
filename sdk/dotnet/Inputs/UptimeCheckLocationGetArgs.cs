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

    public sealed class UptimeCheckLocationGetArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ipv4")]
        public Input<string>? Ipv4 { get; set; }

        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("regionCode")]
        public Input<string>? RegionCode { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public UptimeCheckLocationGetArgs()
        {
        }
    }
}
