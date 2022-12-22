// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package statuscake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StatuscakeUptimeCheck struct {
	pulumi.CustomResourceState

	// Number of seconds between checks
	CheckInterval pulumi.IntOutput `pulumi:"checkInterval"`
	// Number of confirmation servers to confirm downtime before an alert is triggered
	Confirmation pulumi.IntPtrOutput `pulumi:"confirmation"`
	// List of contact group IDs
	ContactGroups pulumi.StringArrayOutput `pulumi:"contactGroups"`
	// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	DnsCheck StatuscakeUptimeCheckDnsCheckPtrOutput `pulumi:"dnsCheck"`
	// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	HttpCheck StatuscakeUptimeCheckHttpCheckPtrOutput `pulumi:"httpCheck"`
	// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	IcmpCheck StatuscakeUptimeCheckIcmpCheckPtrOutput `pulumi:"icmpCheck"`
	// List of assigned monitoring locations on which to run checks
	Locations StatuscakeUptimeCheckLocationArrayOutput `pulumi:"locations"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource StatuscakeUptimeCheckMonitoredResourceOutput `pulumi:"monitoredResource"`
	// Name of the check
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the check should be run
	Paused pulumi.BoolPtrOutput `pulumi:"paused"`
	// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
	// /v1/uptime-locations` endpoint
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// List of tags
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	TcpCheck StatuscakeUptimeCheckTcpCheckPtrOutput `pulumi:"tcpCheck"`
	// The number of minutes to wait before sending an alert
	TriggerRate pulumi.IntPtrOutput `pulumi:"triggerRate"`
}

// NewStatuscakeUptimeCheck registers a new resource with the given unique name, arguments, and options.
func NewStatuscakeUptimeCheck(ctx *pulumi.Context,
	name string, args *StatuscakeUptimeCheckArgs, opts ...pulumi.ResourceOption) (*StatuscakeUptimeCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CheckInterval == nil {
		return nil, errors.New("invalid value for required argument 'CheckInterval'")
	}
	if args.MonitoredResource == nil {
		return nil, errors.New("invalid value for required argument 'MonitoredResource'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StatuscakeUptimeCheck
	err := ctx.RegisterResource("statuscake:index/statuscakeUptimeCheck:StatuscakeUptimeCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStatuscakeUptimeCheck gets an existing StatuscakeUptimeCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStatuscakeUptimeCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StatuscakeUptimeCheckState, opts ...pulumi.ResourceOption) (*StatuscakeUptimeCheck, error) {
	var resource StatuscakeUptimeCheck
	err := ctx.ReadResource("statuscake:index/statuscakeUptimeCheck:StatuscakeUptimeCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StatuscakeUptimeCheck resources.
type statuscakeUptimeCheckState struct {
	// Number of seconds between checks
	CheckInterval *int `pulumi:"checkInterval"`
	// Number of confirmation servers to confirm downtime before an alert is triggered
	Confirmation *int `pulumi:"confirmation"`
	// List of contact group IDs
	ContactGroups []string `pulumi:"contactGroups"`
	// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	DnsCheck *StatuscakeUptimeCheckDnsCheck `pulumi:"dnsCheck"`
	// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	HttpCheck *StatuscakeUptimeCheckHttpCheck `pulumi:"httpCheck"`
	// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	IcmpCheck *StatuscakeUptimeCheckIcmpCheck `pulumi:"icmpCheck"`
	// List of assigned monitoring locations on which to run checks
	Locations []StatuscakeUptimeCheckLocation `pulumi:"locations"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource *StatuscakeUptimeCheckMonitoredResource `pulumi:"monitoredResource"`
	// Name of the check
	Name *string `pulumi:"name"`
	// Whether the check should be run
	Paused *bool `pulumi:"paused"`
	// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
	// /v1/uptime-locations` endpoint
	Regions []string `pulumi:"regions"`
	// List of tags
	Tags []string `pulumi:"tags"`
	// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	TcpCheck *StatuscakeUptimeCheckTcpCheck `pulumi:"tcpCheck"`
	// The number of minutes to wait before sending an alert
	TriggerRate *int `pulumi:"triggerRate"`
}

type StatuscakeUptimeCheckState struct {
	// Number of seconds between checks
	CheckInterval pulumi.IntPtrInput
	// Number of confirmation servers to confirm downtime before an alert is triggered
	Confirmation pulumi.IntPtrInput
	// List of contact group IDs
	ContactGroups pulumi.StringArrayInput
	// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	DnsCheck StatuscakeUptimeCheckDnsCheckPtrInput
	// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	HttpCheck StatuscakeUptimeCheckHttpCheckPtrInput
	// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	IcmpCheck StatuscakeUptimeCheckIcmpCheckPtrInput
	// List of assigned monitoring locations on which to run checks
	Locations StatuscakeUptimeCheckLocationArrayInput
	// Monitored resource configuration block. The describes server under test
	MonitoredResource StatuscakeUptimeCheckMonitoredResourcePtrInput
	// Name of the check
	Name pulumi.StringPtrInput
	// Whether the check should be run
	Paused pulumi.BoolPtrInput
	// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
	// /v1/uptime-locations` endpoint
	Regions pulumi.StringArrayInput
	// List of tags
	Tags pulumi.StringArrayInput
	// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	TcpCheck StatuscakeUptimeCheckTcpCheckPtrInput
	// The number of minutes to wait before sending an alert
	TriggerRate pulumi.IntPtrInput
}

func (StatuscakeUptimeCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*statuscakeUptimeCheckState)(nil)).Elem()
}

type statuscakeUptimeCheckArgs struct {
	// Number of seconds between checks
	CheckInterval int `pulumi:"checkInterval"`
	// Number of confirmation servers to confirm downtime before an alert is triggered
	Confirmation *int `pulumi:"confirmation"`
	// List of contact group IDs
	ContactGroups []string `pulumi:"contactGroups"`
	// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	DnsCheck *StatuscakeUptimeCheckDnsCheck `pulumi:"dnsCheck"`
	// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	HttpCheck *StatuscakeUptimeCheckHttpCheck `pulumi:"httpCheck"`
	// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	IcmpCheck *StatuscakeUptimeCheckIcmpCheck `pulumi:"icmpCheck"`
	// Monitored resource configuration block. The describes server under test
	MonitoredResource StatuscakeUptimeCheckMonitoredResource `pulumi:"monitoredResource"`
	// Name of the check
	Name *string `pulumi:"name"`
	// Whether the check should be run
	Paused *bool `pulumi:"paused"`
	// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
	// /v1/uptime-locations` endpoint
	Regions []string `pulumi:"regions"`
	// List of tags
	Tags []string `pulumi:"tags"`
	// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	TcpCheck *StatuscakeUptimeCheckTcpCheck `pulumi:"tcpCheck"`
	// The number of minutes to wait before sending an alert
	TriggerRate *int `pulumi:"triggerRate"`
}

// The set of arguments for constructing a StatuscakeUptimeCheck resource.
type StatuscakeUptimeCheckArgs struct {
	// Number of seconds between checks
	CheckInterval pulumi.IntInput
	// Number of confirmation servers to confirm downtime before an alert is triggered
	Confirmation pulumi.IntPtrInput
	// List of contact group IDs
	ContactGroups pulumi.StringArrayInput
	// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	DnsCheck StatuscakeUptimeCheckDnsCheckPtrInput
	// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	HttpCheck StatuscakeUptimeCheckHttpCheckPtrInput
	// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	IcmpCheck StatuscakeUptimeCheckIcmpCheckPtrInput
	// Monitored resource configuration block. The describes server under test
	MonitoredResource StatuscakeUptimeCheckMonitoredResourceInput
	// Name of the check
	Name pulumi.StringPtrInput
	// Whether the check should be run
	Paused pulumi.BoolPtrInput
	// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
	// /v1/uptime-locations` endpoint
	Regions pulumi.StringArrayInput
	// List of tags
	Tags pulumi.StringArrayInput
	// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
	TcpCheck StatuscakeUptimeCheckTcpCheckPtrInput
	// The number of minutes to wait before sending an alert
	TriggerRate pulumi.IntPtrInput
}

func (StatuscakeUptimeCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*statuscakeUptimeCheckArgs)(nil)).Elem()
}

type StatuscakeUptimeCheckInput interface {
	pulumi.Input

	ToStatuscakeUptimeCheckOutput() StatuscakeUptimeCheckOutput
	ToStatuscakeUptimeCheckOutputWithContext(ctx context.Context) StatuscakeUptimeCheckOutput
}

func (*StatuscakeUptimeCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**StatuscakeUptimeCheck)(nil)).Elem()
}

func (i *StatuscakeUptimeCheck) ToStatuscakeUptimeCheckOutput() StatuscakeUptimeCheckOutput {
	return i.ToStatuscakeUptimeCheckOutputWithContext(context.Background())
}

func (i *StatuscakeUptimeCheck) ToStatuscakeUptimeCheckOutputWithContext(ctx context.Context) StatuscakeUptimeCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeUptimeCheckOutput)
}

// StatuscakeUptimeCheckArrayInput is an input type that accepts StatuscakeUptimeCheckArray and StatuscakeUptimeCheckArrayOutput values.
// You can construct a concrete instance of `StatuscakeUptimeCheckArrayInput` via:
//
//	StatuscakeUptimeCheckArray{ StatuscakeUptimeCheckArgs{...} }
type StatuscakeUptimeCheckArrayInput interface {
	pulumi.Input

	ToStatuscakeUptimeCheckArrayOutput() StatuscakeUptimeCheckArrayOutput
	ToStatuscakeUptimeCheckArrayOutputWithContext(context.Context) StatuscakeUptimeCheckArrayOutput
}

type StatuscakeUptimeCheckArray []StatuscakeUptimeCheckInput

func (StatuscakeUptimeCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatuscakeUptimeCheck)(nil)).Elem()
}

func (i StatuscakeUptimeCheckArray) ToStatuscakeUptimeCheckArrayOutput() StatuscakeUptimeCheckArrayOutput {
	return i.ToStatuscakeUptimeCheckArrayOutputWithContext(context.Background())
}

func (i StatuscakeUptimeCheckArray) ToStatuscakeUptimeCheckArrayOutputWithContext(ctx context.Context) StatuscakeUptimeCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeUptimeCheckArrayOutput)
}

// StatuscakeUptimeCheckMapInput is an input type that accepts StatuscakeUptimeCheckMap and StatuscakeUptimeCheckMapOutput values.
// You can construct a concrete instance of `StatuscakeUptimeCheckMapInput` via:
//
//	StatuscakeUptimeCheckMap{ "key": StatuscakeUptimeCheckArgs{...} }
type StatuscakeUptimeCheckMapInput interface {
	pulumi.Input

	ToStatuscakeUptimeCheckMapOutput() StatuscakeUptimeCheckMapOutput
	ToStatuscakeUptimeCheckMapOutputWithContext(context.Context) StatuscakeUptimeCheckMapOutput
}

type StatuscakeUptimeCheckMap map[string]StatuscakeUptimeCheckInput

func (StatuscakeUptimeCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatuscakeUptimeCheck)(nil)).Elem()
}

func (i StatuscakeUptimeCheckMap) ToStatuscakeUptimeCheckMapOutput() StatuscakeUptimeCheckMapOutput {
	return i.ToStatuscakeUptimeCheckMapOutputWithContext(context.Background())
}

func (i StatuscakeUptimeCheckMap) ToStatuscakeUptimeCheckMapOutputWithContext(ctx context.Context) StatuscakeUptimeCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatuscakeUptimeCheckMapOutput)
}

type StatuscakeUptimeCheckOutput struct{ *pulumi.OutputState }

func (StatuscakeUptimeCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatuscakeUptimeCheck)(nil)).Elem()
}

func (o StatuscakeUptimeCheckOutput) ToStatuscakeUptimeCheckOutput() StatuscakeUptimeCheckOutput {
	return o
}

func (o StatuscakeUptimeCheckOutput) ToStatuscakeUptimeCheckOutputWithContext(ctx context.Context) StatuscakeUptimeCheckOutput {
	return o
}

// Number of seconds between checks
func (o StatuscakeUptimeCheckOutput) CheckInterval() pulumi.IntOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.IntOutput { return v.CheckInterval }).(pulumi.IntOutput)
}

// Number of confirmation servers to confirm downtime before an alert is triggered
func (o StatuscakeUptimeCheckOutput) Confirmation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.IntPtrOutput { return v.Confirmation }).(pulumi.IntPtrOutput)
}

// List of contact group IDs
func (o StatuscakeUptimeCheckOutput) ContactGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.StringArrayOutput { return v.ContactGroups }).(pulumi.StringArrayOutput)
}

// DNS check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
func (o StatuscakeUptimeCheckOutput) DnsCheck() StatuscakeUptimeCheckDnsCheckPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckDnsCheckPtrOutput { return v.DnsCheck }).(StatuscakeUptimeCheckDnsCheckPtrOutput)
}

// HTTP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
func (o StatuscakeUptimeCheckOutput) HttpCheck() StatuscakeUptimeCheckHttpCheckPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckHttpCheckPtrOutput { return v.HttpCheck }).(StatuscakeUptimeCheckHttpCheckPtrOutput)
}

// ICMP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
func (o StatuscakeUptimeCheckOutput) IcmpCheck() StatuscakeUptimeCheckIcmpCheckPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckIcmpCheckPtrOutput { return v.IcmpCheck }).(StatuscakeUptimeCheckIcmpCheckPtrOutput)
}

// List of assigned monitoring locations on which to run checks
func (o StatuscakeUptimeCheckOutput) Locations() StatuscakeUptimeCheckLocationArrayOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckLocationArrayOutput { return v.Locations }).(StatuscakeUptimeCheckLocationArrayOutput)
}

// Monitored resource configuration block. The describes server under test
func (o StatuscakeUptimeCheckOutput) MonitoredResource() StatuscakeUptimeCheckMonitoredResourceOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckMonitoredResourceOutput {
		return v.MonitoredResource
	}).(StatuscakeUptimeCheckMonitoredResourceOutput)
}

// Name of the check
func (o StatuscakeUptimeCheckOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the check should be run
func (o StatuscakeUptimeCheckOutput) Paused() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.BoolPtrOutput { return v.Paused }).(pulumi.BoolPtrOutput)
}

// List of regions on which to run checks. The values required for this parameter can be retrieved from the `GET
// /v1/uptime-locations` endpoint
func (o StatuscakeUptimeCheckOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// List of tags
func (o StatuscakeUptimeCheckOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// TCP check configuration block. Only one of `dns_check`, `http_check`, `icmp_check`, and `tcp_check` may be specified
func (o StatuscakeUptimeCheckOutput) TcpCheck() StatuscakeUptimeCheckTcpCheckPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) StatuscakeUptimeCheckTcpCheckPtrOutput { return v.TcpCheck }).(StatuscakeUptimeCheckTcpCheckPtrOutput)
}

// The number of minutes to wait before sending an alert
func (o StatuscakeUptimeCheckOutput) TriggerRate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StatuscakeUptimeCheck) pulumi.IntPtrOutput { return v.TriggerRate }).(pulumi.IntPtrOutput)
}

type StatuscakeUptimeCheckArrayOutput struct{ *pulumi.OutputState }

func (StatuscakeUptimeCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StatuscakeUptimeCheck)(nil)).Elem()
}

func (o StatuscakeUptimeCheckArrayOutput) ToStatuscakeUptimeCheckArrayOutput() StatuscakeUptimeCheckArrayOutput {
	return o
}

func (o StatuscakeUptimeCheckArrayOutput) ToStatuscakeUptimeCheckArrayOutputWithContext(ctx context.Context) StatuscakeUptimeCheckArrayOutput {
	return o
}

func (o StatuscakeUptimeCheckArrayOutput) Index(i pulumi.IntInput) StatuscakeUptimeCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StatuscakeUptimeCheck {
		return vs[0].([]*StatuscakeUptimeCheck)[vs[1].(int)]
	}).(StatuscakeUptimeCheckOutput)
}

type StatuscakeUptimeCheckMapOutput struct{ *pulumi.OutputState }

func (StatuscakeUptimeCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StatuscakeUptimeCheck)(nil)).Elem()
}

func (o StatuscakeUptimeCheckMapOutput) ToStatuscakeUptimeCheckMapOutput() StatuscakeUptimeCheckMapOutput {
	return o
}

func (o StatuscakeUptimeCheckMapOutput) ToStatuscakeUptimeCheckMapOutputWithContext(ctx context.Context) StatuscakeUptimeCheckMapOutput {
	return o
}

func (o StatuscakeUptimeCheckMapOutput) MapIndex(k pulumi.StringInput) StatuscakeUptimeCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StatuscakeUptimeCheck {
		return vs[0].(map[string]*StatuscakeUptimeCheck)[vs[1].(string)]
	}).(StatuscakeUptimeCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeUptimeCheckInput)(nil)).Elem(), &StatuscakeUptimeCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeUptimeCheckArrayInput)(nil)).Elem(), StatuscakeUptimeCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StatuscakeUptimeCheckMapInput)(nil)).Elem(), StatuscakeUptimeCheckMap{})
	pulumi.RegisterOutputType(StatuscakeUptimeCheckOutput{})
	pulumi.RegisterOutputType(StatuscakeUptimeCheckArrayOutput{})
	pulumi.RegisterOutputType(StatuscakeUptimeCheckMapOutput{})
}
