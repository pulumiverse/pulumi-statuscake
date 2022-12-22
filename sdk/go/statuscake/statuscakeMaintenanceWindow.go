// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StatuscakeMaintenanceWindow struct {
	pulumi.CustomResourceState

	// End of the maintenance window (RFC3339 format)
	End pulumi.StringOutput `pulumi:"end"`
	// Name of the maintenance window
	Name pulumi.StringOutput `pulumi:"name"`
	// How often the maintenance window should occur
	RepeatInterval pulumi.StringPtrOutput `pulumi:"repeatInterval"`
	// Start of the maintenance window (RFC3339 format)
	Start pulumi.StringOutput `pulumi:"start"`
	// List of tags used to include matching uptime checks in this maintenance window
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// List of uptime check IDs explicitly included in this maintenance window
	Tests pulumi.StringArrayOutput `pulumi:"tests"`
	// Standard timezone associated with this maintenance window
	Timezone pulumi.StringOutput `pulumi:"timezone"`
}

// NewStatuscakeMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewStatuscakeMaintenanceWindow(ctx *pulumi.Context,
	name string, args *StatuscakeMaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*StatuscakeMaintenanceWindow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	if args.Timezone == nil {
		return nil, errors.New("invalid value for required argument 'Timezone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StatuscakeMaintenanceWindow
	err := ctx.RegisterResource("statuscake:index/statuscakeMaintenanceWindow:StatuscakeMaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatuscakeMaintenanceWindow gets an existing StatuscakeMaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatuscakeMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatuscakeMaintenanceWindowState, opts ...pulumi.ResourceOption) (*StatuscakeMaintenanceWindow, error) {
	var resource StatuscakeMaintenanceWindow
	err := ctx.ReadResource("statuscake:index/statuscakeMaintenanceWindow:StatuscakeMaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatuscakeMaintenanceWindow resources.
type statuscakeMaintenanceWindowState struct {
	// End of the maintenance window (RFC3339 format)
	End *string `pulumi:"end"`
	// Name of the maintenance window
	Name *string `pulumi:"name"`
	// How often the maintenance window should occur
	RepeatInterval *string `pulumi:"repeatInterval"`
	// Start of the maintenance window (RFC3339 format)
	Start *string `pulumi:"start"`
	// List of tags used to include matching uptime checks in this maintenance window
	Tags []string `pulumi:"tags"`
	// List of uptime check IDs explicitly included in this maintenance window
	Tests []string `pulumi:"tests"`
	// Standard timezone associated with this maintenance window
	Timezone *string `pulumi:"timezone"`
}

type StatuscakeMaintenanceWindowState struct {
	// End of the maintenance window (RFC3339 format)
	End pulumi.StringPtrInput
	// Name of the maintenance window
	Name pulumi.StringPtrInput
	// How often the maintenance window should occur
	RepeatInterval pulumi.StringPtrInput
	// Start of the maintenance window (RFC3339 format)
	Start pulumi.StringPtrInput
	// List of tags used to include matching uptime checks in this maintenance window
	Tags pulumi.StringArrayInput
	// List of uptime check IDs explicitly included in this maintenance window
	Tests pulumi.StringArrayInput
	// Standard timezone associated with this maintenance window
	Timezone pulumi.StringPtrInput
}

func (StatuscakeMaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*statuscakeMaintenanceWindowState)(nil)).Elem()
}

type statuscakeMaintenanceWindowArgs struct {
	// End of the maintenance window (RFC3339 format)
	End string `pulumi:"end"`
	// Name of the maintenance window
	Name *string `pulumi:"name"`
	// How often the maintenance window should occur
	RepeatInterval *string `pulumi:"repeatInterval"`
	// Start of the maintenance window (RFC3339 format)
	Start string `pulumi:"start"`
	// List of tags used to include matching uptime checks in this maintenance window
	Tags []string `pulumi:"tags"`
	// List of uptime check IDs explicitly included in this maintenance window
	Tests []string `pulumi:"tests"`
	// Standard timezone associated with this maintenance window
	Timezone string `pulumi:"timezone"`
}

// The set of arguments for constructing a StatuscakeMaintenanceWindow resource.
type StatuscakeMaintenanceWindowArgs struct {
	// End of the maintenance window (RFC3339 format)
	End pulumi.StringInput
	// Name of the maintenance window
	Name pulumi.StringPtrInput
	// How often the maintenance window should occur
	RepeatInterval pulumi.StringPtrInput
	// Start of the maintenance window (RFC3339 format)
	Start pulumi.StringInput
	// List of tags used to include matching uptime checks in this maintenance window
	Tags pulumi.StringArrayInput
	// List of uptime check IDs explicitly included in this maintenance window
	Tests pulumi.StringArrayInput
	// Standard timezone associated with this maintenance window
	Timezone pulumi.StringInput
}

func (StatuscakeMaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statuscakeMaintenanceWindowArgs)(nil)).Elem()
}

type StatuscakeMaintenanceWindowInput interface {
	pulumi.Input

	ToStatuscakeMaintenanceWindowOutput() StatuscakeMaintenanceWindowOutput
	ToStatuscakeMaintenanceWindowOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowOutput
}

func (*StatuscakeMaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (i *StatuscakeMaintenanceWindow) ToStatuscakeMaintenanceWindowOutput() StatuscakeMaintenanceWindowOutput {
	return i.ToStatuscakeMaintenanceWindowOutputWithContext(context.Background())
}

func (i *StatuscakeMaintenanceWindow) ToStatuscakeMaintenanceWindowOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeMaintenanceWindowOutput)
}

// StatuscakeMaintenanceWindowArrayInput is an input type that accepts StatuscakeMaintenanceWindowArray and StatuscakeMaintenanceWindowArrayOutput values.
// You can construct a concrete instance of `StatuscakeMaintenanceWindowArrayInput` via:
//
//	StatuscakeMaintenanceWindowArray{ StatuscakeMaintenanceWindowArgs{...} }
type StatuscakeMaintenanceWindowArrayInput interface {
	pulumi.Input

	ToStatuscakeMaintenanceWindowArrayOutput() StatuscakeMaintenanceWindowArrayOutput
	ToStatuscakeMaintenanceWindowArrayOutputWithContext(context.Context) StatuscakeMaintenanceWindowArrayOutput
}

type StatuscakeMaintenanceWindowArray []StatuscakeMaintenanceWindowInput

func (StatuscakeMaintenanceWindowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (i StatuscakeMaintenanceWindowArray) ToStatuscakeMaintenanceWindowArrayOutput() StatuscakeMaintenanceWindowArrayOutput {
	return i.ToStatuscakeMaintenanceWindowArrayOutputWithContext(context.Background())
}

func (i StatuscakeMaintenanceWindowArray) ToStatuscakeMaintenanceWindowArrayOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeMaintenanceWindowArrayOutput)
}

// StatuscakeMaintenanceWindowMapInput is an input type that accepts StatuscakeMaintenanceWindowMap and StatuscakeMaintenanceWindowMapOutput values.
// You can construct a concrete instance of `StatuscakeMaintenanceWindowMapInput` via:
//
//	StatuscakeMaintenanceWindowMap{ "key": StatuscakeMaintenanceWindowArgs{...} }
type StatuscakeMaintenanceWindowMapInput interface {
	pulumi.Input

	ToStatuscakeMaintenanceWindowMapOutput() StatuscakeMaintenanceWindowMapOutput
	ToStatuscakeMaintenanceWindowMapOutputWithContext(context.Context) StatuscakeMaintenanceWindowMapOutput
}

type StatuscakeMaintenanceWindowMap map[string]StatuscakeMaintenanceWindowInput

func (StatuscakeMaintenanceWindowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (i StatuscakeMaintenanceWindowMap) ToStatuscakeMaintenanceWindowMapOutput() StatuscakeMaintenanceWindowMapOutput {
	return i.ToStatuscakeMaintenanceWindowMapOutputWithContext(context.Background())
}

func (i StatuscakeMaintenanceWindowMap) ToStatuscakeMaintenanceWindowMapOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeMaintenanceWindowMapOutput)
}

type StatuscakeMaintenanceWindowOutput struct{ *pulumi.OutputState }

func (StatuscakeMaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (o StatuscakeMaintenanceWindowOutput) ToStatuscakeMaintenanceWindowOutput() StatuscakeMaintenanceWindowOutput {
	return o
}

func (o StatuscakeMaintenanceWindowOutput) ToStatuscakeMaintenanceWindowOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowOutput {
	return o
}

// End of the maintenance window (RFC3339 format)
func (o StatuscakeMaintenanceWindowOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringOutput { return v.End }).(pulumi.StringOutput)
}

// Name of the maintenance window
func (o StatuscakeMaintenanceWindowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// How often the maintenance window should occur
func (o StatuscakeMaintenanceWindowOutput) RepeatInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringPtrOutput { return v.RepeatInterval }).(pulumi.StringPtrOutput)
}

// Start of the maintenance window (RFC3339 format)
func (o StatuscakeMaintenanceWindowOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// List of tags used to include matching uptime checks in this maintenance window
func (o StatuscakeMaintenanceWindowOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// List of uptime check IDs explicitly included in this maintenance window
func (o StatuscakeMaintenanceWindowOutput) Tests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringArrayOutput { return v.Tests }).(pulumi.StringArrayOutput)
}

// Standard timezone associated with this maintenance window
func (o StatuscakeMaintenanceWindowOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v *StatuscakeMaintenanceWindow) pulumi.StringOutput { return v.Timezone }).(pulumi.StringOutput)
}

type StatuscakeMaintenanceWindowArrayOutput struct{ *pulumi.OutputState }

func (StatuscakeMaintenanceWindowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (o StatuscakeMaintenanceWindowArrayOutput) ToStatuscakeMaintenanceWindowArrayOutput() StatuscakeMaintenanceWindowArrayOutput {
	return o
}

func (o StatuscakeMaintenanceWindowArrayOutput) ToStatuscakeMaintenanceWindowArrayOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowArrayOutput {
	return o
}

func (o StatuscakeMaintenanceWindowArrayOutput) Index(i pulumi.IntInput) StatuscakeMaintenanceWindowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatuscakeMaintenanceWindow {
		return vs[0].([]*StatuscakeMaintenanceWindow)[vs[1].(int)]
	}).(StatuscakeMaintenanceWindowOutput)
}

type StatuscakeMaintenanceWindowMapOutput struct{ *pulumi.OutputState }

func (StatuscakeMaintenanceWindowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatuscakeMaintenanceWindow)(nil)).Elem()
}

func (o StatuscakeMaintenanceWindowMapOutput) ToStatuscakeMaintenanceWindowMapOutput() StatuscakeMaintenanceWindowMapOutput {
	return o
}

func (o StatuscakeMaintenanceWindowMapOutput) ToStatuscakeMaintenanceWindowMapOutputWithContext(ctx context.Context) StatuscakeMaintenanceWindowMapOutput {
	return o
}

func (o StatuscakeMaintenanceWindowMapOutput) MapIndex(k pulumi.StringInput) StatuscakeMaintenanceWindowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatuscakeMaintenanceWindow {
		return vs[0].(map[string]*StatuscakeMaintenanceWindow)[vs[1].(string)]
	}).(StatuscakeMaintenanceWindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeMaintenanceWindowInput)(nil)).Elem(), &StatuscakeMaintenanceWindow{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeMaintenanceWindowArrayInput)(nil)).Elem(), StatuscakeMaintenanceWindowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeMaintenanceWindowMapInput)(nil)).Elem(), StatuscakeMaintenanceWindowMap{})
	pulumi.RegisterOutputType(StatuscakeMaintenanceWindowOutput{})
	pulumi.RegisterOutputType(StatuscakeMaintenanceWindowArrayOutput{})
	pulumi.RegisterOutputType(StatuscakeMaintenanceWindowMapOutput{})
}