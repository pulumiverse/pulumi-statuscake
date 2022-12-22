// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Statuscake.Outputs
{

    [OutputType]
    public sealed class GetUptimeMonitoringLocationsLocationResult
    {
        public readonly string Description;
        public readonly string Ipv4;
        public readonly string Ipv6;
        public readonly string Region;
        public readonly string RegionCode;
        public readonly string Status;

        [OutputConstructor]
        private GetUptimeMonitoringLocationsLocationResult(
            string description,

            string ipv4,

            string ipv6,

            string region,

            string regionCode,

            string status)
        {
            Description = description;
            Ipv4 = ipv4;
            Ipv6 = ipv6;
            Region = region;
            RegionCode = regionCode;
            Status = status;
        }
    }
}
