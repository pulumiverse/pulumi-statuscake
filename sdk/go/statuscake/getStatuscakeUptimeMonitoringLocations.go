// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStatuscakeUptimeMonitoringLocations(ctx *pulumi.Context, args *GetStatuscakeUptimeMonitoringLocationsArgs, opts ...pulumi.InvokeOption) (*GetStatuscakeUptimeMonitoringLocationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetStatuscakeUptimeMonitoringLocationsResult
	err := ctx.Invoke("statuscake:index/getStatuscakeUptimeMonitoringLocations:getStatuscakeUptimeMonitoringLocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStatuscakeUptimeMonitoringLocations.
type GetStatuscakeUptimeMonitoringLocationsArgs struct {
	RegionCode *string `pulumi:"regionCode"`
}

// A collection of values returned by getStatuscakeUptimeMonitoringLocations.
type GetStatuscakeUptimeMonitoringLocationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                           `pulumi:"id"`
	Locations  []GetStatuscakeUptimeMonitoringLocationsLocation `pulumi:"locations"`
	RegionCode *string                                          `pulumi:"regionCode"`
}

func GetStatuscakeUptimeMonitoringLocationsOutput(ctx *pulumi.Context, args GetStatuscakeUptimeMonitoringLocationsOutputArgs, opts ...pulumi.InvokeOption) GetStatuscakeUptimeMonitoringLocationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStatuscakeUptimeMonitoringLocationsResult, error) {
			args := v.(GetStatuscakeUptimeMonitoringLocationsArgs)
			r, err := GetStatuscakeUptimeMonitoringLocations(ctx, &args, opts...)
			var s GetStatuscakeUptimeMonitoringLocationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStatuscakeUptimeMonitoringLocationsResultOutput)
}

// A collection of arguments for invoking getStatuscakeUptimeMonitoringLocations.
type GetStatuscakeUptimeMonitoringLocationsOutputArgs struct {
	RegionCode pulumi.StringPtrInput `pulumi:"regionCode"`
}

func (GetStatuscakeUptimeMonitoringLocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatuscakeUptimeMonitoringLocationsArgs)(nil)).Elem()
}

// A collection of values returned by getStatuscakeUptimeMonitoringLocations.
type GetStatuscakeUptimeMonitoringLocationsResultOutput struct{ *pulumi.OutputState }

func (GetStatuscakeUptimeMonitoringLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatuscakeUptimeMonitoringLocationsResult)(nil)).Elem()
}

func (o GetStatuscakeUptimeMonitoringLocationsResultOutput) ToGetStatuscakeUptimeMonitoringLocationsResultOutput() GetStatuscakeUptimeMonitoringLocationsResultOutput {
	return o
}

func (o GetStatuscakeUptimeMonitoringLocationsResultOutput) ToGetStatuscakeUptimeMonitoringLocationsResultOutputWithContext(ctx context.Context) GetStatuscakeUptimeMonitoringLocationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetStatuscakeUptimeMonitoringLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStatuscakeUptimeMonitoringLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStatuscakeUptimeMonitoringLocationsResultOutput) Locations() GetStatuscakeUptimeMonitoringLocationsLocationArrayOutput {
	return o.ApplyT(func(v GetStatuscakeUptimeMonitoringLocationsResult) []GetStatuscakeUptimeMonitoringLocationsLocation {
		return v.Locations
	}).(GetStatuscakeUptimeMonitoringLocationsLocationArrayOutput)
}

func (o GetStatuscakeUptimeMonitoringLocationsResultOutput) RegionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStatuscakeUptimeMonitoringLocationsResult) *string { return v.RegionCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStatuscakeUptimeMonitoringLocationsResultOutput{})
}