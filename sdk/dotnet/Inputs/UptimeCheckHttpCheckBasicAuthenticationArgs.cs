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

    public sealed class UptimeCheckHttpCheckBasicAuthenticationArgs : Pulumi.ResourceArgs
    {
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UptimeCheckHttpCheckBasicAuthenticationArgs()
        {
        }
    }
}
