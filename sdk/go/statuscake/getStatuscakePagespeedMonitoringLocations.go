// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStatuscakePagespeedMonitoringLocations(ctx *pulumi.Context, args *GetStatuscakePagespeedMonitoringLocationsArgs, opts ...pulumi.InvokeOption) (*GetStatuscakePagespeedMonitoringLocationsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetStatuscakePagespeedMonitoringLocationsResult
	err := ctx.Invoke("statuscake:index/getStatuscakePagespeedMonitoringLocations:getStatuscakePagespeedMonitoringLocations", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStatuscakePagespeedMonitoringLocations.
type GetStatuscakePagespeedMonitoringLocationsArgs struct {
	RegionCode *string `pulumi:"regionCode"`
}

// A collection of values returned by getStatuscakePagespeedMonitoringLocations.
type GetStatuscakePagespeedMonitoringLocationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                                              `pulumi:"id"`
	Locations  []GetStatuscakePagespeedMonitoringLocationsLocation `pulumi:"locations"`
	RegionCode *string                                             `pulumi:"regionCode"`
}

func GetStatuscakePagespeedMonitoringLocationsOutput(ctx *pulumi.Context, args GetStatuscakePagespeedMonitoringLocationsOutputArgs, opts ...pulumi.InvokeOption) GetStatuscakePagespeedMonitoringLocationsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetStatuscakePagespeedMonitoringLocationsResult, error) {
			args := v.(GetStatuscakePagespeedMonitoringLocationsArgs)
			r, err := GetStatuscakePagespeedMonitoringLocations(ctx, &args, opts...)
			var s GetStatuscakePagespeedMonitoringLocationsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetStatuscakePagespeedMonitoringLocationsResultOutput)
}

// A collection of arguments for invoking getStatuscakePagespeedMonitoringLocations.
type GetStatuscakePagespeedMonitoringLocationsOutputArgs struct {
	RegionCode pulumi.StringPtrInput `pulumi:"regionCode"`
}

func (GetStatuscakePagespeedMonitoringLocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatuscakePagespeedMonitoringLocationsArgs)(nil)).Elem()
}

// A collection of values returned by getStatuscakePagespeedMonitoringLocations.
type GetStatuscakePagespeedMonitoringLocationsResultOutput struct{ *pulumi.OutputState }

func (GetStatuscakePagespeedMonitoringLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStatuscakePagespeedMonitoringLocationsResult)(nil)).Elem()
}

func (o GetStatuscakePagespeedMonitoringLocationsResultOutput) ToGetStatuscakePagespeedMonitoringLocationsResultOutput() GetStatuscakePagespeedMonitoringLocationsResultOutput {
	return o
}

func (o GetStatuscakePagespeedMonitoringLocationsResultOutput) ToGetStatuscakePagespeedMonitoringLocationsResultOutputWithContext(ctx context.Context) GetStatuscakePagespeedMonitoringLocationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetStatuscakePagespeedMonitoringLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStatuscakePagespeedMonitoringLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStatuscakePagespeedMonitoringLocationsResultOutput) Locations() GetStatuscakePagespeedMonitoringLocationsLocationArrayOutput {
	return o.ApplyT(func(v GetStatuscakePagespeedMonitoringLocationsResult) []GetStatuscakePagespeedMonitoringLocationsLocation {
		return v.Locations
	}).(GetStatuscakePagespeedMonitoringLocationsLocationArrayOutput)
}

func (o GetStatuscakePagespeedMonitoringLocationsResultOutput) RegionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStatuscakePagespeedMonitoringLocationsResult) *string { return v.RegionCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStatuscakePagespeedMonitoringLocationsResultOutput{})
}
