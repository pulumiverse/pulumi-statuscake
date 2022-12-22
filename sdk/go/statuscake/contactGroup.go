// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ContactGroup struct {
	pulumi.CustomResourceState

	// List of email addresses
	EmailAddresses pulumi.StringArrayOutput `pulumi:"emailAddresses"`
	// List of integration IDs
	Integrations pulumi.StringArrayOutput `pulumi:"integrations"`
	// List of international format mobile phone numbers
	MobileNumbers pulumi.StringArrayOutput `pulumi:"mobileNumbers"`
	// Name of the contact group
	Name pulumi.StringOutput `pulumi:"name"`
	// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
	PingUrl pulumi.StringPtrOutput `pulumi:"pingUrl"`
}

// NewContactGroup registers a new resource with the given unique name, arguments, and options.
func NewContactGroup(ctx *pulumi.Context,
	name string, args *ContactGroupArgs, opts ...pulumi.ResourceOption) (*ContactGroup, error) {
	if args == nil {
		args = &ContactGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ContactGroup
	err := ctx.RegisterResource("statuscake:index/contactGroup:ContactGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactGroup gets an existing ContactGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactGroupState, opts ...pulumi.ResourceOption) (*ContactGroup, error) {
	var resource ContactGroup
	err := ctx.ReadResource("statuscake:index/contactGroup:ContactGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactGroup resources.
type contactGroupState struct {
	// List of email addresses
	EmailAddresses []string `pulumi:"emailAddresses"`
	// List of integration IDs
	Integrations []string `pulumi:"integrations"`
	// List of international format mobile phone numbers
	MobileNumbers []string `pulumi:"mobileNumbers"`
	// Name of the contact group
	Name *string `pulumi:"name"`
	// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
	PingUrl *string `pulumi:"pingUrl"`
}

type ContactGroupState struct {
	// List of email addresses
	EmailAddresses pulumi.StringArrayInput
	// List of integration IDs
	Integrations pulumi.StringArrayInput
	// List of international format mobile phone numbers
	MobileNumbers pulumi.StringArrayInput
	// Name of the contact group
	Name pulumi.StringPtrInput
	// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
	PingUrl pulumi.StringPtrInput
}

func (ContactGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactGroupState)(nil)).Elem()
}

type contactGroupArgs struct {
	// List of email addresses
	EmailAddresses []string `pulumi:"emailAddresses"`
	// List of integration IDs
	Integrations []string `pulumi:"integrations"`
	// List of international format mobile phone numbers
	MobileNumbers []string `pulumi:"mobileNumbers"`
	// Name of the contact group
	Name *string `pulumi:"name"`
	// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
	PingUrl *string `pulumi:"pingUrl"`
}

// The set of arguments for constructing a ContactGroup resource.
type ContactGroupArgs struct {
	// List of email addresses
	EmailAddresses pulumi.StringArrayInput
	// List of integration IDs
	Integrations pulumi.StringArrayInput
	// List of international format mobile phone numbers
	MobileNumbers pulumi.StringArrayInput
	// Name of the contact group
	Name pulumi.StringPtrInput
	// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
	PingUrl pulumi.StringPtrInput
}

func (ContactGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactGroupArgs)(nil)).Elem()
}

type ContactGroupInput interface {
	pulumi.Input

	ToContactGroupOutput() ContactGroupOutput
	ToContactGroupOutputWithContext(ctx context.Context) ContactGroupOutput
}

func (*ContactGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactGroup)(nil)).Elem()
}

func (i *ContactGroup) ToContactGroupOutput() ContactGroupOutput {
	return i.ToContactGroupOutputWithContext(context.Background())
}

func (i *ContactGroup) ToContactGroupOutputWithContext(ctx context.Context) ContactGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactGroupOutput)
}

// ContactGroupArrayInput is an input type that accepts ContactGroupArray and ContactGroupArrayOutput values.
// You can construct a concrete instance of `ContactGroupArrayInput` via:
//
//	ContactGroupArray{ ContactGroupArgs{...} }
type ContactGroupArrayInput interface {
	pulumi.Input

	ToContactGroupArrayOutput() ContactGroupArrayOutput
	ToContactGroupArrayOutputWithContext(context.Context) ContactGroupArrayOutput
}

type ContactGroupArray []ContactGroupInput

func (ContactGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactGroup)(nil)).Elem()
}

func (i ContactGroupArray) ToContactGroupArrayOutput() ContactGroupArrayOutput {
	return i.ToContactGroupArrayOutputWithContext(context.Background())
}

func (i ContactGroupArray) ToContactGroupArrayOutputWithContext(ctx context.Context) ContactGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactGroupArrayOutput)
}

// ContactGroupMapInput is an input type that accepts ContactGroupMap and ContactGroupMapOutput values.
// You can construct a concrete instance of `ContactGroupMapInput` via:
//
//	ContactGroupMap{ "key": ContactGroupArgs{...} }
type ContactGroupMapInput interface {
	pulumi.Input

	ToContactGroupMapOutput() ContactGroupMapOutput
	ToContactGroupMapOutputWithContext(context.Context) ContactGroupMapOutput
}

type ContactGroupMap map[string]ContactGroupInput

func (ContactGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactGroup)(nil)).Elem()
}

func (i ContactGroupMap) ToContactGroupMapOutput() ContactGroupMapOutput {
	return i.ToContactGroupMapOutputWithContext(context.Background())
}

func (i ContactGroupMap) ToContactGroupMapOutputWithContext(ctx context.Context) ContactGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactGroupMapOutput)
}

type ContactGroupOutput struct{ *pulumi.OutputState }

func (ContactGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactGroup)(nil)).Elem()
}

func (o ContactGroupOutput) ToContactGroupOutput() ContactGroupOutput {
	return o
}

func (o ContactGroupOutput) ToContactGroupOutputWithContext(ctx context.Context) ContactGroupOutput {
	return o
}

// List of email addresses
func (o ContactGroupOutput) EmailAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactGroup) pulumi.StringArrayOutput { return v.EmailAddresses }).(pulumi.StringArrayOutput)
}

// List of integration IDs
func (o ContactGroupOutput) Integrations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactGroup) pulumi.StringArrayOutput { return v.Integrations }).(pulumi.StringArrayOutput)
}

// List of international format mobile phone numbers
func (o ContactGroupOutput) MobileNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContactGroup) pulumi.StringArrayOutput { return v.MobileNumbers }).(pulumi.StringArrayOutput)
}

// Name of the contact group
func (o ContactGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
func (o ContactGroupOutput) PingUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactGroup) pulumi.StringPtrOutput { return v.PingUrl }).(pulumi.StringPtrOutput)
}

type ContactGroupArrayOutput struct{ *pulumi.OutputState }

func (ContactGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContactGroup)(nil)).Elem()
}

func (o ContactGroupArrayOutput) ToContactGroupArrayOutput() ContactGroupArrayOutput {
	return o
}

func (o ContactGroupArrayOutput) ToContactGroupArrayOutputWithContext(ctx context.Context) ContactGroupArrayOutput {
	return o
}

func (o ContactGroupArrayOutput) Index(i pulumi.IntInput) ContactGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContactGroup {
		return vs[0].([]*ContactGroup)[vs[1].(int)]
	}).(ContactGroupOutput)
}

type ContactGroupMapOutput struct{ *pulumi.OutputState }

func (ContactGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContactGroup)(nil)).Elem()
}

func (o ContactGroupMapOutput) ToContactGroupMapOutput() ContactGroupMapOutput {
	return o
}

func (o ContactGroupMapOutput) ToContactGroupMapOutputWithContext(ctx context.Context) ContactGroupMapOutput {
	return o
}

func (o ContactGroupMapOutput) MapIndex(k pulumi.StringInput) ContactGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContactGroup {
		return vs[0].(map[string]*ContactGroup)[vs[1].(string)]
	}).(ContactGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactGroupInput)(nil)).Elem(), &ContactGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactGroupArrayInput)(nil)).Elem(), ContactGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContactGroupMapInput)(nil)).Elem(), ContactGroupMap{})
	pulumi.RegisterOutputType(ContactGroupOutput{})
	pulumi.RegisterOutputType(ContactGroupArrayOutput{})
	pulumi.RegisterOutputType(ContactGroupMapOutput{})
}
