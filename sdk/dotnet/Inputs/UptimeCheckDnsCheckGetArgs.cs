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

    public sealed class UptimeCheckDnsCheckGetArgs : Pulumi.ResourceArgs
    {
        [Input("dnsIps", required: true)]
        private InputList<string>? _dnsIps;
        public InputList<string> DnsIps
        {
            get => _dnsIps ?? (_dnsIps = new InputList<string>());
            set => _dnsIps = value;
        }

        [Input("dnsServer")]
        public Input<string>? DnsServer { get; set; }

        public UptimeCheckDnsCheckGetArgs()
        {
        }
    }
}
