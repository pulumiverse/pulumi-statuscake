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
    public static class GetStatuscakeContactGroup
    {
        public static Task<GetStatuscakeContactGroupResult> InvokeAsync(GetStatuscakeContactGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStatuscakeContactGroupResult>("statuscake:index/getStatuscakeContactGroup:getStatuscakeContactGroup", args ?? new GetStatuscakeContactGroupArgs(), options.WithDefaults());

        public static Output<GetStatuscakeContactGroupResult> Invoke(GetStatuscakeContactGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetStatuscakeContactGroupResult>("statuscake:index/getStatuscakeContactGroup:getStatuscakeContactGroup", args ?? new GetStatuscakeContactGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStatuscakeContactGroupArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetStatuscakeContactGroupArgs()
        {
        }
    }

    public sealed class GetStatuscakeContactGroupInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetStatuscakeContactGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStatuscakeContactGroupResult
    {
        public readonly ImmutableArray<string> EmailAddresses;
        public readonly string Id;
        public readonly ImmutableArray<string> Integrations;
        public readonly ImmutableArray<string> MobileNumbers;
        public readonly string Name;
        public readonly string PingUrl;

        [OutputConstructor]
        private GetStatuscakeContactGroupResult(
            ImmutableArray<string> emailAddresses,

            string id,

            ImmutableArray<string> integrations,

            ImmutableArray<string> mobileNumbers,

            string name,

            string pingUrl)
        {
            EmailAddresses = emailAddresses;
            Id = id;
            Integrations = integrations;
            MobileNumbers = mobileNumbers;
            Name = name;
            PingUrl = pingUrl;
        }
    }
}
