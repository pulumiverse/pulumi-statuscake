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
    public sealed class PagespeedCheckAlertConfig
    {
        public readonly int? AlertBigger;
        public readonly int? AlertSlower;
        public readonly int? AlertSmaller;

        [OutputConstructor]
        private PagespeedCheckAlertConfig(
            int? alertBigger,

            int? alertSlower,

            int? alertSmaller)
        {
            AlertBigger = alertBigger;
            AlertSlower = alertSlower;
            AlertSmaller = alertSmaller;
        }
    }
}