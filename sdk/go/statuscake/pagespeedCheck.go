// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PagespeedCheck struct {
	pulumi.CustomResourceState

	// Alert configuration block. Omitting this block disabled all alerts
	AlertConfig PagespeedCheckAlertConfigOutput `pulumi:"alertConfig"`
	// Number of seconds between checks
	CheckInterval pulumi.IntOutput `pulumi:"checkInterval"`
	// List of contact group IDs
	ContactGroups pulumi.StringArrayOutput `pulumi:"contactGroups"`
	// Assigned monitoring location on which checks will be run
	Location pulumi.StringOutput `pulumi:"location"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource PagespeedCheckMonitoredResourceOutput `pulumi:"monitoredResource"`
	// Name of the check
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the check should be run
	Paused pulumi.BoolPtrOutput `pulumi:"paused"`
	// Region on which to run checks
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPagespeedCheck registers a new resource with the given unique name, arguments, and options.
func NewPagespeedCheck(ctx *pulumi.Context,
	name string, args *PagespeedCheckArgs, opts ...pulumi.ResourceOption) (*PagespeedCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertConfig == nil {
		return nil, errors.New("invalid value for required argument 'AlertConfig'")
	}
	if args.CheckInterval == nil {
		return nil, errors.New("invalid value for required argument 'CheckInterval'")
	}
	if args.MonitoredResource == nil {
		return nil, errors.New("invalid value for required argument 'MonitoredResource'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource PagespeedCheck
	err := ctx.RegisterResource("statuscake:index/pagespeedCheck:PagespeedCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPagespeedCheck gets an existing PagespeedCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPagespeedCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PagespeedCheckState, opts ...pulumi.ResourceOption) (*PagespeedCheck, error) {
	var resource PagespeedCheck
	err := ctx.ReadResource("statuscake:index/pagespeedCheck:PagespeedCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PagespeedCheck resources.
type pagespeedCheckState struct {
	// Alert configuration block. Omitting this block disabled all alerts
	AlertConfig *PagespeedCheckAlertConfig `pulumi:"alertConfig"`
	// Number of seconds between checks
	CheckInterval *int `pulumi:"checkInterval"`
	// List of contact group IDs
	ContactGroups []string `pulumi:"contactGroups"`
	// Assigned monitoring location on which checks will be run
	Location *string `pulumi:"location"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource *PagespeedCheckMonitoredResource `pulumi:"monitoredResource"`
	// Name of the check
	Name *string `pulumi:"name"`
	// Whether the check should be run
	Paused *bool `pulumi:"paused"`
	// Region on which to run checks
	Region *string `pulumi:"region"`
}

type PagespeedCheckState struct {
	// Alert configuration block. Omitting this block disabled all alerts
	AlertConfig PagespeedCheckAlertConfigPtrInput
	// Number of seconds between checks
	CheckInterval pulumi.IntPtrInput
	// List of contact group IDs
	ContactGroups pulumi.StringArrayInput
	// Assigned monitoring location on which checks will be run
	Location pulumi.StringPtrInput
	// Monitored resource configuration block. The describes server under test
	MonitoredResource PagespeedCheckMonitoredResourcePtrInput
	// Name of the check
	Name pulumi.StringPtrInput
	// Whether the check should be run
	Paused pulumi.BoolPtrInput
	// Region on which to run checks
	Region pulumi.StringPtrInput
}

func (PagespeedCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*pagespeedCheckState)(nil)).Elem()
}

type pagespeedCheckArgs struct {
	// Alert configuration block. Omitting this block disabled all alerts
	AlertConfig PagespeedCheckAlertConfig `pulumi:"alertConfig"`
	// Number of seconds between checks
	CheckInterval int `pulumi:"checkInterval"`
	// List of contact group IDs
	ContactGroups []string `pulumi:"contactGroups"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource PagespeedCheckMonitoredResource `pulumi:"monitoredResource"`
	// Name of the check
	Name *string `pulumi:"name"`
	// Whether the check should be run
	Paused *bool `pulumi:"paused"`
	// Region on which to run checks
	Region string `pulumi:"region"`
}

// The set of arguments for constructing a PagespeedCheck resource.
type PagespeedCheckArgs struct {
	// Alert configuration block. Omitting this block disabled all alerts
	AlertConfig PagespeedCheckAlertConfigInput
	// Number of seconds between checks
	CheckInterval pulumi.IntInput
	// List of contact group IDs
	ContactGroups pulumi.StringArrayInput
	// Monitored resource configuration block. The describes server under test
	MonitoredResource PagespeedCheckMonitoredResourceInput
	// Name of the check
	Name pulumi.StringPtrInput
	// Whether the check should be run
	Paused pulumi.BoolPtrInput
	// Region on which to run checks
	Region pulumi.StringInput
}

func (PagespeedCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pagespeedCheckArgs)(nil)).Elem()
}

type PagespeedCheckInput interface {
	pulumi.Input

	ToPagespeedCheckOutput() PagespeedCheckOutput
	ToPagespeedCheckOutputWithContext(ctx context.Context) PagespeedCheckOutput
}

func (*PagespeedCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**PagespeedCheck)(nil)).Elem()
}

func (i *PagespeedCheck) ToPagespeedCheckOutput() PagespeedCheckOutput {
	return i.ToPagespeedCheckOutputWithContext(context.Background())
}

func (i *PagespeedCheck) ToPagespeedCheckOutputWithContext(ctx context.Context) PagespeedCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagespeedCheckOutput)
}

// PagespeedCheckArrayInput is an input type that accepts PagespeedCheckArray and PagespeedCheckArrayOutput values.
// You can construct a concrete instance of `PagespeedCheckArrayInput` via:
//
//	PagespeedCheckArray{ PagespeedCheckArgs{...} }
type PagespeedCheckArrayInput interface {
	pulumi.Input

	ToPagespeedCheckArrayOutput() PagespeedCheckArrayOutput
	ToPagespeedCheckArrayOutputWithContext(context.Context) PagespeedCheckArrayOutput
}

type PagespeedCheckArray []PagespeedCheckInput

func (PagespeedCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PagespeedCheck)(nil)).Elem()
}

func (i PagespeedCheckArray) ToPagespeedCheckArrayOutput() PagespeedCheckArrayOutput {
	return i.ToPagespeedCheckArrayOutputWithContext(context.Background())
}

func (i PagespeedCheckArray) ToPagespeedCheckArrayOutputWithContext(ctx context.Context) PagespeedCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagespeedCheckArrayOutput)
}

// PagespeedCheckMapInput is an input type that accepts PagespeedCheckMap and PagespeedCheckMapOutput values.
// You can construct a concrete instance of `PagespeedCheckMapInput` via:
//
//	PagespeedCheckMap{ "key": PagespeedCheckArgs{...} }
type PagespeedCheckMapInput interface {
	pulumi.Input

	ToPagespeedCheckMapOutput() PagespeedCheckMapOutput
	ToPagespeedCheckMapOutputWithContext(context.Context) PagespeedCheckMapOutput
}

type PagespeedCheckMap map[string]PagespeedCheckInput

func (PagespeedCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PagespeedCheck)(nil)).Elem()
}

func (i PagespeedCheckMap) ToPagespeedCheckMapOutput() PagespeedCheckMapOutput {
	return i.ToPagespeedCheckMapOutputWithContext(context.Background())
}

func (i PagespeedCheckMap) ToPagespeedCheckMapOutputWithContext(ctx context.Context) PagespeedCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PagespeedCheckMapOutput)
}

type PagespeedCheckOutput struct{ *pulumi.OutputState }

func (PagespeedCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PagespeedCheck)(nil)).Elem()
}

func (o PagespeedCheckOutput) ToPagespeedCheckOutput() PagespeedCheckOutput {
	return o
}

func (o PagespeedCheckOutput) ToPagespeedCheckOutputWithContext(ctx context.Context) PagespeedCheckOutput {
	return o
}

// Alert configuration block. Omitting this block disabled all alerts
func (o PagespeedCheckOutput) AlertConfig() PagespeedCheckAlertConfigOutput {
	return o.ApplyT(func(v *PagespeedCheck) PagespeedCheckAlertConfigOutput { return v.AlertConfig }).(PagespeedCheckAlertConfigOutput)
}

// Number of seconds between checks
func (o PagespeedCheckOutput) CheckInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.IntOutput { return v.CheckInterval }).(pulumi.IntOutput)
}

// List of contact group IDs
func (o PagespeedCheckOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.StringArrayOutput { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

// Assigned monitoring location on which checks will be run
func (o PagespeedCheckOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Monitored resource configuration block. The describes server under test
func (o PagespeedCheckOutput) MonitoredResource() PagespeedCheckMonitoredResourceOutput {
	return o.ApplyT(func(v *PagespeedCheck) PagespeedCheckMonitoredResourceOutput { return v.MonitoredResource }).(PagespeedCheckMonitoredResourceOutput)
}

// Name of the check
func (o PagespeedCheckOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the check should be run
func (o PagespeedCheckOutput) Paused() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.BoolPtrOutput { return v.Paused }).(pulumi.BoolPtrOutput)
}

// Region on which to run checks
func (o PagespeedCheckOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PagespeedCheck) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type PagespeedCheckArrayOutput struct{ *pulumi.OutputState }

func (PagespeedCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PagespeedCheck)(nil)).Elem()
}

func (o PagespeedCheckArrayOutput) ToPagespeedCheckArrayOutput() PagespeedCheckArrayOutput {
	return o
}

func (o PagespeedCheckArrayOutput) ToPagespeedCheckArrayOutputWithContext(ctx context.Context) PagespeedCheckArrayOutput {
	return o
}

func (o PagespeedCheckArrayOutput) Index(i pulumi.IntInput) PagespeedCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PagespeedCheck {
		return vs[0].([]*PagespeedCheck)[vs[1].(int)]
	}).(PagespeedCheckOutput)
}

type PagespeedCheckMapOutput struct{ *pulumi.OutputState }

func (PagespeedCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PagespeedCheck)(nil)).Elem()
}

func (o PagespeedCheckMapOutput) ToPagespeedCheckMapOutput() PagespeedCheckMapOutput {
	return o
}

func (o PagespeedCheckMapOutput) ToPagespeedCheckMapOutputWithContext(ctx context.Context) PagespeedCheckMapOutput {
	return o
}

func (o PagespeedCheckMapOutput) MapIndex(k pulumi.StringInput) PagespeedCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PagespeedCheck {
		return vs[0].(map[string]*PagespeedCheck)[vs[1].(string)]
	}).(PagespeedCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PagespeedCheckInput)(nil)).Elem(), &PagespeedCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*PagespeedCheckArrayInput)(nil)).Elem(), PagespeedCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PagespeedCheckMapInput)(nil)).Elem(), PagespeedCheckMap{})
	pulumi.RegisterOutputType(PagespeedCheckOutput{})
	pulumi.RegisterOutputType(PagespeedCheckArrayOutput{})
	pulumi.RegisterOutputType(PagespeedCheckMapOutput{})
}
