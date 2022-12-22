// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContactGroup(ctx *pulumi.Context, args *LookupContactGroupArgs, opts ...pulumi.InvokeOption) (*LookupContactGroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupContactGroupResult
	err := ctx.Invoke("statuscake:index/getContactGroup:getContactGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContactGroup.
type LookupContactGroupArgs struct {
	Id string `pulumi:"id"`
}

// A collection of values returned by getContactGroup.
type LookupContactGroupResult struct {
	EmailAddresses []string `pulumi:"emailAddresses"`
	Id             string   `pulumi:"id"`
	Integrations   []string `pulumi:"integrations"`
	MobileNumbers  []string `pulumi:"mobileNumbers"`
	Name           string   `pulumi:"name"`
	PingUrl        string   `pulumi:"pingUrl"`
}

func LookupContactGroupOutput(ctx *pulumi.Context, args LookupContactGroupOutputArgs, opts ...pulumi.InvokeOption) LookupContactGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactGroupResult, error) {
			args := v.(LookupContactGroupArgs)
			r, err := LookupContactGroup(ctx, &args, opts...)
			var s LookupContactGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactGroupResultOutput)
}

// A collection of arguments for invoking getContactGroup.
type LookupContactGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupContactGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactGroupArgs)(nil)).Elem()
}

// A collection of values returned by getContactGroup.
type LookupContactGroupResultOutput struct{ *pulumi.OutputState }

func (LookupContactGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactGroupResult)(nil)).Elem()
}

func (o LookupContactGroupResultOutput) ToLookupContactGroupResultOutput() LookupContactGroupResultOutput {
	return o
}

func (o LookupContactGroupResultOutput) ToLookupContactGroupResultOutputWithContext(ctx context.Context) LookupContactGroupResultOutput {
	return o
}

func (o LookupContactGroupResultOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContactGroupResult) []string { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupContactGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContactGroupResultOutput) Integrations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContactGroupResult) []string { return v.Integrations }).(pulumi.StringArrayOutput)
}

func (o LookupContactGroupResultOutput) MobileNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupContactGroupResult) []string { return v.MobileNumbers }).(pulumi.StringArrayOutput)
}

func (o LookupContactGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContactGroupResultOutput) PingUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactGroupResult) string { return v.PingUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactGroupResultOutput{})
}