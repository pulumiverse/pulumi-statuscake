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

__all__ = ['SslCheckArgs', 'SslCheck']

@pulumi.input_type
class SslCheckArgs:
    def __init__(__self__, *,
                 alert_config: pulumi.Input['SslCheckAlertConfigArgs'],
                 check_interval: pulumi.Input[int],
                 monitored_resource: pulumi.Input['SslCheckMonitoredResourceArgs'],
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslCheck resource.
        :param pulumi.Input['SslCheckAlertConfigArgs'] alert_config: Alert configuration block
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input['SslCheckMonitoredResourceArgs'] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[bool] follow_redirects: Whether to follow redirects when testing. Disabled by default
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] user_agent: Custom user agent string set when testing
        """
        pulumi.set(__self__, "alert_config", alert_config)
        pulumi.set(__self__, "check_interval", check_interval)
        pulumi.set(__self__, "monitored_resource", monitored_resource)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if follow_redirects is not None:
            pulumi.set(__self__, "follow_redirects", follow_redirects)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> pulumi.Input['SslCheckAlertConfigArgs']:
        """
        Alert configuration block
        """
        return pulumi.get(self, "alert_config")

    @alert_config.setter
    def alert_config(self, value: pulumi.Input['SslCheckAlertConfigArgs']):
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
    def monitored_resource(self) -> pulumi.Input['SslCheckMonitoredResourceArgs']:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @monitored_resource.setter
    def monitored_resource(self, value: pulumi.Input['SslCheckMonitoredResourceArgs']):
        pulumi.set(self, "monitored_resource", value)

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
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to follow redirects when testing. Disabled by default
        """
        return pulumi.get(self, "follow_redirects")

    @follow_redirects.setter
    def follow_redirects(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "follow_redirects", value)

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
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Custom user agent string set when testing
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)


@pulumi.input_type
class _SslCheckState:
    def __init__(__self__, *,
                 alert_config: Optional[pulumi.Input['SslCheckAlertConfigArgs']] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 monitored_resource: Optional[pulumi.Input['SslCheckMonitoredResourceArgs']] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SslCheck resources.
        :param pulumi.Input['SslCheckAlertConfigArgs'] alert_config: Alert configuration block
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[bool] follow_redirects: Whether to follow redirects when testing. Disabled by default
        :param pulumi.Input['SslCheckMonitoredResourceArgs'] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] user_agent: Custom user agent string set when testing
        """
        if alert_config is not None:
            pulumi.set(__self__, "alert_config", alert_config)
        if check_interval is not None:
            pulumi.set(__self__, "check_interval", check_interval)
        if contact_groups is not None:
            pulumi.set(__self__, "contact_groups", contact_groups)
        if follow_redirects is not None:
            pulumi.set(__self__, "follow_redirects", follow_redirects)
        if monitored_resource is not None:
            pulumi.set(__self__, "monitored_resource", monitored_resource)
        if paused is not None:
            pulumi.set(__self__, "paused", paused)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> Optional[pulumi.Input['SslCheckAlertConfigArgs']]:
        """
        Alert configuration block
        """
        return pulumi.get(self, "alert_config")

    @alert_config.setter
    def alert_config(self, value: Optional[pulumi.Input['SslCheckAlertConfigArgs']]):
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
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to follow redirects when testing. Disabled by default
        """
        return pulumi.get(self, "follow_redirects")

    @follow_redirects.setter
    def follow_redirects(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "follow_redirects", value)

    @property
    @pulumi.getter(name="monitoredResource")
    def monitored_resource(self) -> Optional[pulumi.Input['SslCheckMonitoredResourceArgs']]:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @monitored_resource.setter
    def monitored_resource(self, value: Optional[pulumi.Input['SslCheckMonitoredResourceArgs']]):
        pulumi.set(self, "monitored_resource", value)

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
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        """
        Custom user agent string set when testing
        """
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)


class SslCheck(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_config: Optional[pulumi.Input[pulumi.InputType['SslCheckAlertConfigArgs']]] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 monitored_resource: Optional[pulumi.Input[pulumi.InputType['SslCheckMonitoredResourceArgs']]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SslCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SslCheckAlertConfigArgs']] alert_config: Alert configuration block
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[bool] follow_redirects: Whether to follow redirects when testing. Disabled by default
        :param pulumi.Input[pulumi.InputType['SslCheckMonitoredResourceArgs']] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] user_agent: Custom user agent string set when testing
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SslCheckArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SslCheck resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SslCheckArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslCheckArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_config: Optional[pulumi.Input[pulumi.InputType['SslCheckAlertConfigArgs']]] = None,
                 check_interval: Optional[pulumi.Input[int]] = None,
                 contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 monitored_resource: Optional[pulumi.Input[pulumi.InputType['SslCheckMonitoredResourceArgs']]] = None,
                 paused: Optional[pulumi.Input[bool]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
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
            __props__ = SslCheckArgs.__new__(SslCheckArgs)

            if alert_config is None and not opts.urn:
                raise TypeError("Missing required property 'alert_config'")
            __props__.__dict__["alert_config"] = alert_config
            if check_interval is None and not opts.urn:
                raise TypeError("Missing required property 'check_interval'")
            __props__.__dict__["check_interval"] = check_interval
            __props__.__dict__["contact_groups"] = contact_groups
            __props__.__dict__["follow_redirects"] = follow_redirects
            if monitored_resource is None and not opts.urn:
                raise TypeError("Missing required property 'monitored_resource'")
            __props__.__dict__["monitored_resource"] = monitored_resource
            __props__.__dict__["paused"] = paused
            __props__.__dict__["user_agent"] = user_agent
        super(SslCheck, __self__).__init__(
            'statuscake:index/sslCheck:SslCheck',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alert_config: Optional[pulumi.Input[pulumi.InputType['SslCheckAlertConfigArgs']]] = None,
            check_interval: Optional[pulumi.Input[int]] = None,
            contact_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            follow_redirects: Optional[pulumi.Input[bool]] = None,
            monitored_resource: Optional[pulumi.Input[pulumi.InputType['SslCheckMonitoredResourceArgs']]] = None,
            paused: Optional[pulumi.Input[bool]] = None,
            user_agent: Optional[pulumi.Input[str]] = None) -> 'SslCheck':
        """
        Get an existing SslCheck resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SslCheckAlertConfigArgs']] alert_config: Alert configuration block
        :param pulumi.Input[int] check_interval: Number of seconds between checks
        :param pulumi.Input[Sequence[pulumi.Input[str]]] contact_groups: List of contact group IDs
        :param pulumi.Input[bool] follow_redirects: Whether to follow redirects when testing. Disabled by default
        :param pulumi.Input[pulumi.InputType['SslCheckMonitoredResourceArgs']] monitored_resource: Monitored resource configuration block. The describes server under test
        :param pulumi.Input[bool] paused: Whether the check should be run
        :param pulumi.Input[str] user_agent: Custom user agent string set when testing
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SslCheckState.__new__(_SslCheckState)

        __props__.__dict__["alert_config"] = alert_config
        __props__.__dict__["check_interval"] = check_interval
        __props__.__dict__["contact_groups"] = contact_groups
        __props__.__dict__["follow_redirects"] = follow_redirects
        __props__.__dict__["monitored_resource"] = monitored_resource
        __props__.__dict__["paused"] = paused
        __props__.__dict__["user_agent"] = user_agent
        return SslCheck(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertConfig")
    def alert_config(self) -> pulumi.Output['outputs.SslCheckAlertConfig']:
        """
        Alert configuration block
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
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to follow redirects when testing. Disabled by default
        """
        return pulumi.get(self, "follow_redirects")

    @property
    @pulumi.getter(name="monitoredResource")
    def monitored_resource(self) -> pulumi.Output['outputs.SslCheckMonitoredResource']:
        """
        Monitored resource configuration block. The describes server under test
        """
        return pulumi.get(self, "monitored_resource")

    @property
    @pulumi.getter
    def paused(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the check should be run
        """
        return pulumi.get(self, "paused")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> pulumi.Output[Optional[str]]:
        """
        Custom user agent string set when testing
        """
        return pulumi.get(self, "user_agent")

