// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPagespeedMonitoringLocations(ctx *pulumi.Context, args *GetPagespeedMonitoringLocationsArgs, opts ...pulumi.InvokeOption) (*GetPagespeedMonitoringLocationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPagespeedMonitoringLocationsResult
	err := ctx.Invoke("statuscake:index/getPagespeedMonitoringLocations:getPagespeedMonitoringLocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPagespeedMonitoringLocations.
type GetPagespeedMonitoringLocationsArgs struct {
	RegionCode *string `pulumi:"regionCode"`
}

// A collection of values returned by getPagespeedMonitoringLocations.
type GetPagespeedMonitoringLocationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                    `pulumi:"id"`
	Locations  []GetPagespeedMonitoringLocationsLocation `pulumi:"locations"`
	RegionCode *string                                   `pulumi:"regionCode"`
}

func GetPagespeedMonitoringLocationsOutput(ctx *pulumi.Context, args GetPagespeedMonitoringLocationsOutputArgs, opts ...pulumi.InvokeOption) GetPagespeedMonitoringLocationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPagespeedMonitoringLocationsResult, error) {
			args := v.(GetPagespeedMonitoringLocationsArgs)
			r, err := GetPagespeedMonitoringLocations(ctx, &args, opts...)
			var s GetPagespeedMonitoringLocationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPagespeedMonitoringLocationsResultOutput)
}

// A collection of arguments for invoking getPagespeedMonitoringLocations.
type GetPagespeedMonitoringLocationsOutputArgs struct {
	RegionCode pulumi.StringPtrInput `pulumi:"regionCode"`
}

func (GetPagespeedMonitoringLocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPagespeedMonitoringLocationsArgs)(nil)).Elem()
}

// A collection of values returned by getPagespeedMonitoringLocations.
type GetPagespeedMonitoringLocationsResultOutput struct{ *pulumi.OutputState }

func (GetPagespeedMonitoringLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPagespeedMonitoringLocationsResult)(nil)).Elem()
}

func (o GetPagespeedMonitoringLocationsResultOutput) ToGetPagespeedMonitoringLocationsResultOutput() GetPagespeedMonitoringLocationsResultOutput {
	return o
}

func (o GetPagespeedMonitoringLocationsResultOutput) ToGetPagespeedMonitoringLocationsResultOutputWithContext(ctx context.Context) GetPagespeedMonitoringLocationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetPagespeedMonitoringLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPagespeedMonitoringLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetPagespeedMonitoringLocationsResultOutput) Locations() GetPagespeedMonitoringLocationsLocationArrayOutput {
	return o.ApplyT(func(v GetPagespeedMonitoringLocationsResult) []GetPagespeedMonitoringLocationsLocation {
		return v.Locations
	}).(GetPagespeedMonitoringLocationsLocationArrayOutput)
}

func (o GetPagespeedMonitoringLocationsResultOutput) RegionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPagespeedMonitoringLocationsResult) *string { return v.RegionCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPagespeedMonitoringLocationsResultOutput{})
}
