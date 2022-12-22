// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MaintenanceWindow struct {
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

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
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
	var resource MaintenanceWindow
	err := ctx.RegisterResource("statuscake:index/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("statuscake:index/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
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

type MaintenanceWindowState struct {
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

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
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

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
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

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (*MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

// MaintenanceWindowArrayInput is an input type that accepts MaintenanceWindowArray and MaintenanceWindowArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowArrayInput` via:
//
//	MaintenanceWindowArray{ MaintenanceWindowArgs{...} }
type MaintenanceWindowArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput
	ToMaintenanceWindowArrayOutputWithContext(context.Context) MaintenanceWindowArrayOutput
}

type MaintenanceWindowArray []MaintenanceWindowInput

func (MaintenanceWindowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return i.ToMaintenanceWindowArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowArrayOutput)
}

// MaintenanceWindowMapInput is an input type that accepts MaintenanceWindowMap and MaintenanceWindowMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowMapInput` via:
//
//	MaintenanceWindowMap{ "key": MaintenanceWindowArgs{...} }
type MaintenanceWindowMapInput interface {
	pulumi.Input

	ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput
	ToMaintenanceWindowMapOutputWithContext(context.Context) MaintenanceWindowMapOutput
}

type MaintenanceWindowMap map[string]MaintenanceWindowInput

func (MaintenanceWindowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return i.ToMaintenanceWindowMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowMapOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

// End of the maintenance window (RFC3339 format)
func (o MaintenanceWindowOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.End }).(pulumi.StringOutput)
}

// Name of the maintenance window
func (o MaintenanceWindowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// How often the maintenance window should occur
func (o MaintenanceWindowOutput) RepeatInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.RepeatInterval }).(pulumi.StringPtrOutput)
}

// Start of the maintenance window (RFC3339 format)
func (o MaintenanceWindowOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

// List of tags used to include matching uptime checks in this maintenance window
func (o MaintenanceWindowOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// List of uptime check IDs explicitly included in this maintenance window
func (o MaintenanceWindowOutput) Tests() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringArrayOutput { return v.Tests }).(pulumi.StringArrayOutput)
}

// Standard timezone associated with this maintenance window
func (o MaintenanceWindowOutput) Timezone() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Timezone }).(pulumi.StringOutput)
}

type MaintenanceWindowArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].([]*MaintenanceWindow)[vs[1].(int)]
	}).(MaintenanceWindowOutput)
}

type MaintenanceWindowMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].(map[string]*MaintenanceWindow)[vs[1].(string)]
	}).(MaintenanceWindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowInput)(nil)).Elem(), &MaintenanceWindow{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowArrayInput)(nil)).Elem(), MaintenanceWindowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowMapInput)(nil)).Elem(), MaintenanceWindowMap{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowMapOutput{})
}
