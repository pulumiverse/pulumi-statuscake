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

    public sealed class PagespeedCheckAlertConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("alertBigger")]
        public Input<int>? AlertBigger { get; set; }

        [Input("alertSlower")]
        public Input<int>? AlertSlower { get; set; }

        [Input("alertSmaller")]
        public Input<int>? AlertSmaller { get; set; }

        public PagespeedCheckAlertConfigGetArgs()
        {
        }
    }
}