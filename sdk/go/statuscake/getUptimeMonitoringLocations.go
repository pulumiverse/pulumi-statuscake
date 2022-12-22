// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetUptimeMonitoringLocations(ctx *pulumi.Context, args *GetUptimeMonitoringLocationsArgs, opts ...pulumi.InvokeOption) (*GetUptimeMonitoringLocationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetUptimeMonitoringLocationsResult
	err := ctx.Invoke("statuscake:index/getUptimeMonitoringLocations:getUptimeMonitoringLocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUptimeMonitoringLocations.
type GetUptimeMonitoringLocationsArgs struct {
	RegionCode *string `pulumi:"regionCode"`
}

// A collection of values returned by getUptimeMonitoringLocations.
type GetUptimeMonitoringLocationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                 `pulumi:"id"`
	Locations  []GetUptimeMonitoringLocationsLocation `pulumi:"locations"`
	RegionCode *string                                `pulumi:"regionCode"`
}

func GetUptimeMonitoringLocationsOutput(ctx *pulumi.Context, args GetUptimeMonitoringLocationsOutputArgs, opts ...pulumi.InvokeOption) GetUptimeMonitoringLocationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUptimeMonitoringLocationsResult, error) {
			args := v.(GetUptimeMonitoringLocationsArgs)
			r, err := GetUptimeMonitoringLocations(ctx, &args, opts...)
			var s GetUptimeMonitoringLocationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUptimeMonitoringLocationsResultOutput)
}

// A collection of arguments for invoking getUptimeMonitoringLocations.
type GetUptimeMonitoringLocationsOutputArgs struct {
	RegionCode pulumi.StringPtrInput `pulumi:"regionCode"`
}

func (GetUptimeMonitoringLocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUptimeMonitoringLocationsArgs)(nil)).Elem()
}

// A collection of values returned by getUptimeMonitoringLocations.
type GetUptimeMonitoringLocationsResultOutput struct{ *pulumi.OutputState }

func (GetUptimeMonitoringLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUptimeMonitoringLocationsResult)(nil)).Elem()
}

func (o GetUptimeMonitoringLocationsResultOutput) ToGetUptimeMonitoringLocationsResultOutput() GetUptimeMonitoringLocationsResultOutput {
	return o
}

func (o GetUptimeMonitoringLocationsResultOutput) ToGetUptimeMonitoringLocationsResultOutputWithContext(ctx context.Context) GetUptimeMonitoringLocationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetUptimeMonitoringLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUptimeMonitoringLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetUptimeMonitoringLocationsResultOutput) Locations() GetUptimeMonitoringLocationsLocationArrayOutput {
	return o.ApplyT(func(v GetUptimeMonitoringLocationsResult) []GetUptimeMonitoringLocationsLocation { return v.Locations }).(GetUptimeMonitoringLocationsLocationArrayOutput)
}

func (o GetUptimeMonitoringLocationsResultOutput) RegionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUptimeMonitoringLocationsResult) *string { return v.RegionCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUptimeMonitoringLocationsResultOutput{})
}
