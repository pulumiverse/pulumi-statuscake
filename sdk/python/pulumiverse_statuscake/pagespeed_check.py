# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PagespeedCheckArgs', 'PagespeedCheck']

@pulumi.input_type
class PagespeedCheckArgs:
    def __init__(__self__, *,
                 alert_config: pulumi.Input['PagespeedCheckAlertConfigArgs'],
                 check_interval: pulumi.Input[int],
                 monitored_resource: pulumi.Input['PagespeedCheckMonitoredResourceArgs'],
                 region: pulumi.Input[str],
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paused: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PagespeedCheck resource.
        :param pulumi.Input['PagespeedCheckAlertConfigArgs'] alert_config: Alert configuration block. Omitting this block disabled all alerts
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input['PagespeedCheckMonitoredResourceArgs'] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[str] region: Region on which to run checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[str] name: Name of the check
        :param pulumi.Input[bool] paused: Whether the check should be run
        """
        pulumi.set(__self__, "alert_config", alert_config)
        pulumi.set(__self__, "check_interval", check_interval)
        pulumi.set(__self__, "monitored_resource", monitored_resource)
        pulumi.set(__self__, "region", region)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> pulumi.Input['PagespeedCheckAlertConfigArgs']:
        """
        Alert configuration block. Omitting this block disabled all alerts
        """
        return pulumi.get(self, "alert_config")

    @alert_config.setter
    def alert_config(self, value: pulumi.Input['PagespeedCheckAlertConfigArgs']):
        pulumi.set(self, "alert_config", value)

    @property
    @pulumi.getter(name="checkInterval")
    def check_interval(self) -> pulumi.Input[int]:
        """
        Number of seconds between checks
        """
        return pulumi.get(self, "check_interval")

    @check_interval.setter
    def check_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "check_interval", value)

    @property
    @pulumi.getter(name="monitoredResource")
    def monitored_resource(self) -> pulumi.Input['PagespeedCheckMonitoredResourceArgs']:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @monitored_resource.setter
    def monitored_resource(self, value: pulumi.Input['PagespeedCheckMonitoredResourceArgs']):
        pulumi.set(self, "monitored_resource", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        Region on which to run checks
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of contact group IDs
        """
        return pulumi.get(self, "contact_groups")

    @contact_groups.setter
    def contact_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the check
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def paused(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the check should be run
        """
        return pulumi.get(self, "paused")

    @paused.setter
    def paused(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "paused", value)


@pulumi.input_type
class _PagespeedCheckState:
    def __init__(__self__, *,
                 alert_config: Optional[pulumi.Input['PagespeedCheckAlertConfigArgs']] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 monitored_resource: Optional[pulumi.Input['PagespeedCheckMonitoredResourceArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PagespeedCheck resources.
        :param pulumi.Input['PagespeedCheckAlertConfigArgs'] alert_config: Alert configuration block. Omitting this block disabled all alerts
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[str] location: Assigned monitoring location on which checks will be run
        :param pulumi.Input['PagespeedCheckMonitoredResourceArgs'] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[str] name: Name of the check
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] region: Region on which to run checks
        """
        if alert_config is not None:
            pulumi.set(__self__, "alert_config", alert_config)
        if check_interval is not None:
            pulumi.set(__self__, "check_interval", check_interval)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if monitored_resource is not None:
            pulumi.set(__self__, "monitored_resource", monitored_resource)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> Optional[pulumi.Input['PagespeedCheckAlertConfigArgs']]:
        """
        Alert configuration block. Omitting this block disabled all alerts
        """
        return pulumi.get(self, "alert_config")

    @alert_config.setter
    def alert_config(self, value: Optional[pulumi.Input['PagespeedCheckAlertConfigArgs']]):
        pulumi.set(self, "alert_config", value)

    @property
    @pulumi.getter(name="checkInterval")
    def check_interval(self) -> Optional[pulumi.Input[int]]:
        """
        Number of seconds between checks
        """
        return pulumi.get(self, "check_interval")

    @check_interval.setter
    def check_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "check_interval", value)

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of contact group IDs
        """
        return pulumi.get(self, "contact_groups")

    @contact_groups.setter
    def contact_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "contact_groups", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Assigned monitoring location on which checks will be run
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="monitoredResource")
    def monitored_resource(self) -> Optional[pulumi.Input['PagespeedCheckMonitoredResourceArgs']]:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @monitored_resource.setter
    def monitored_resource(self, value: Optional[pulumi.Input['PagespeedCheckMonitoredResourceArgs']]):
        pulumi.set(self, "monitored_resource", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the check
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def paused(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the check should be run
        """
        return pulumi.get(self, "paused")

    @paused.setter
    def paused(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "paused", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Region on which to run checks
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


class PagespeedCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_config: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckAlertConfigArgs']]] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 monitored_resource: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckMonitoredResourceArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PagespeedCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PagespeedCheckAlertConfigArgs']] alert_config: Alert configuration block. Omitting this block disabled all alerts
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[pulumi.InputType['PagespeedCheckMonitoredResourceArgs']] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[str] name: Name of the check
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] region: Region on which to run checks
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PagespeedCheckArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PagespeedCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PagespeedCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PagespeedCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_config: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckAlertConfigArgs']]] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 monitored_resource: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckMonitoredResourceArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        else:
            opts = copy.copy(opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PagespeedCheckArgs.__new__(PagespeedCheckArgs)

            if alert_config is None and not opts.urn:
                raise TypeError("Missing required property 'alert_config'")
            __props__.__dict__["alert_config"] = alert_config
            if check_interval is None and not opts.urn:
                raise TypeError("Missing required property 'check_interval'")
            __props__.__dict__["check_interval"] = check_interval
            __props__.__dict__["contact_groups"] = contact_groups
            if monitored_resource is None and not opts.urn:
                raise TypeError("Missing required property 'monitored_resource'")
            __props__.__dict__["monitored_resource"] = monitored_resource
            __props__.__dict__["name"] = name
            __props__.__dict__["paused"] = paused
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["location"] = None
        super(PagespeedCheck, __self__).__init__(
            'statuscake:index/pagespeedCheck:PagespeedCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_config: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckAlertConfigArgs']]] = None,
            check_interval: Optional[pulumi.Input[int]] = None,
            contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            monitored_resource: Optional[pulumi.Input[pulumi.InputType['PagespeedCheckMonitoredResourceArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            paused: Optional[pulumi.Input[bool]] = None,
            region: Optional[pulumi.Input[str]] = None) -> 'PagespeedCheck':
        """
        Get an existing PagespeedCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PagespeedCheckAlertConfigArgs']] alert_config: Alert configuration block. Omitting this block disabled all alerts
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[str] location: Assigned monitoring location on which checks will be run
        :param pulumi.Input[pulumi.InputType['PagespeedCheckMonitoredResourceArgs']] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[str] name: Name of the check
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] region: Region on which to run checks
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PagespeedCheckState.__new__(_PagespeedCheckState)

        __props__.__dict__["alert_config"] = alert_config
        __props__.__dict__["check_interval"] = check_interval
        __props__.__dict__["contact_groups"] = contact_groups
        __props__.__dict__["location"] = location
        __props__.__dict__["monitored_resource"] = monitored_resource
        __props__.__dict__["name"] = name
        __props__.__dict__["paused"] = paused
        __props__.__dict__["region"] = region
        return PagespeedCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> pulumi.Output['outputs.PagespeedCheckAlertConfig']:
        """
        Alert configuration block. Omitting this block disabled all alerts
        """
        return pulumi.get(self, "alert_config")

    @property
    @pulumi.getter(name="checkInterval")
    def check_interval(self) -> pulumi.Output[int]:
        """
        Number of seconds between checks
        """
        return pulumi.get(self, "check_interval")

    @property
    @pulumi.getter(name="contactGroups")
    def contact_groups(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of contact group IDs
        """
        return pulumi.get(self, "contact_groups")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Assigned monitoring location on which checks will be run
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="monitoredResource")
    def monitored_resource(self) -> pulumi.Output['outputs.PagespeedCheckMonitoredResource']:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the check
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def paused(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the check should be run
        """
        return pulumi.get(self, "paused")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Region on which to run checks
        """
        return pulumi.get(self, "region")

