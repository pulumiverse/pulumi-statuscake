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

    public sealed class UptimeCheckHttpCheckContentMatchersGetArgs : Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        [Input("includeHeaders")]
        public Input<bool>? IncludeHeaders { get; set; }

        [Input("matcher")]
        public Input<string>? Matcher { get; set; }

        public UptimeCheckHttpCheckContentMatchersGetArgs()
        {
        }
    }
}
