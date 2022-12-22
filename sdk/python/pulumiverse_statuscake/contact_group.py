# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ContactGroupArgs', 'ContactGroup']

@pulumi.input_type
class ContactGroupArgs:
    def __init__(__self__, *,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integrations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mobile_numbers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ping_url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ContactGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: List of email addresses
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integrations: List of integration IDs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mobile_numbers: List of international format mobile phone numbers
        :param pulumi.Input[str] name: Name of the contact group
        :param pulumi.Input[str] ping_url: URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        if email_addresses is not None:
            pulumi.set(__self__, "email_addresses", email_addresses)
        if integrations is not None:
            pulumi.set(__self__, "integrations", integrations)
        if mobile_numbers is not None:
            pulumi.set(__self__, "mobile_numbers", mobile_numbers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ping_url is not None:
            pulumi.set(__self__, "ping_url", ping_url)

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of email addresses
        """
        return pulumi.get(self, "email_addresses")

    @email_addresses.setter
    def email_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "email_addresses", value)

    @property
    @pulumi.getter
    def integrations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of integration IDs
        """
        return pulumi.get(self, "integrations")

    @integrations.setter
    def integrations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "integrations", value)

    @property
    @pulumi.getter(name="mobileNumbers")
    def mobile_numbers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of international format mobile phone numbers
        """
        return pulumi.get(self, "mobile_numbers")

    @mobile_numbers.setter
    def mobile_numbers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "mobile_numbers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the contact group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pingUrl")
    def ping_url(self) -> Optional[pulumi.Input[str]]:
        """
        URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        return pulumi.get(self, "ping_url")

    @ping_url.setter
    def ping_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ping_url", value)


@pulumi.input_type
class _ContactGroupState:
    def __init__(__self__, *,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integrations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mobile_numbers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ping_url: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ContactGroup resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: List of email addresses
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integrations: List of integration IDs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mobile_numbers: List of international format mobile phone numbers
        :param pulumi.Input[str] name: Name of the contact group
        :param pulumi.Input[str] ping_url: URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        if email_addresses is not None:
            pulumi.set(__self__, "email_addresses", email_addresses)
        if integrations is not None:
            pulumi.set(__self__, "integrations", integrations)
        if mobile_numbers is not None:
            pulumi.set(__self__, "mobile_numbers", mobile_numbers)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ping_url is not None:
            pulumi.set(__self__, "ping_url", ping_url)

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of email addresses
        """
        return pulumi.get(self, "email_addresses")

    @email_addresses.setter
    def email_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "email_addresses", value)

    @property
    @pulumi.getter
    def integrations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of integration IDs
        """
        return pulumi.get(self, "integrations")

    @integrations.setter
    def integrations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "integrations", value)

    @property
    @pulumi.getter(name="mobileNumbers")
    def mobile_numbers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of international format mobile phone numbers
        """
        return pulumi.get(self, "mobile_numbers")

    @mobile_numbers.setter
    def mobile_numbers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "mobile_numbers", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the contact group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="pingUrl")
    def ping_url(self) -> Optional[pulumi.Input[str]]:
        """
        URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        return pulumi.get(self, "ping_url")

    @ping_url.setter
    def ping_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ping_url", value)


class ContactGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integrations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mobile_numbers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ping_url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ContactGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: List of email addresses
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integrations: List of integration IDs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mobile_numbers: List of international format mobile phone numbers
        :param pulumi.Input[str] name: Name of the contact group
        :param pulumi.Input[str] ping_url: URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ContactGroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ContactGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ContactGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ContactGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 integrations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 mobile_numbers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ping_url: Optional[pulumi.Input[str]] = None,
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
            __props__ = ContactGroupArgs.__new__(ContactGroupArgs)

            __props__.__dict__["email_addresses"] = email_addresses
            __props__.__dict__["integrations"] = integrations
            __props__.__dict__["mobile_numbers"] = mobile_numbers
            __props__.__dict__["name"] = name
            __props__.__dict__["ping_url"] = ping_url
        super(ContactGroup, __self__).__init__(
            'statuscake:index/contactGroup:ContactGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            email_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            integrations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            mobile_numbers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ping_url: Optional[pulumi.Input[str]] = None) -> 'ContactGroup':
        """
        Get an existing ContactGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] email_addresses: List of email addresses
        :param pulumi.Input[Sequence[pulumi.Input[str]]] integrations: List of integration IDs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] mobile_numbers: List of international format mobile phone numbers
        :param pulumi.Input[str] name: Name of the contact group
        :param pulumi.Input[str] ping_url: URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ContactGroupState.__new__(_ContactGroupState)

        __props__.__dict__["email_addresses"] = email_addresses
        __props__.__dict__["integrations"] = integrations
        __props__.__dict__["mobile_numbers"] = mobile_numbers
        __props__.__dict__["name"] = name
        __props__.__dict__["ping_url"] = ping_url
        return ContactGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of email addresses
        """
        return pulumi.get(self, "email_addresses")

    @property
    @pulumi.getter
    def integrations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of integration IDs
        """
        return pulumi.get(self, "integrations")

    @property
    @pulumi.getter(name="mobileNumbers")
    def mobile_numbers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of international format mobile phone numbers
        """
        return pulumi.get(self, "mobile_numbers")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the contact group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pingUrl")
    def ping_url(self) -> pulumi.Output[Optional[str]]:
        """
        URL or IP address of an endpoint to push uptime events. Currently this only supports HTTP GET endpoints
        """
        return pulumi.get(self, "ping_url")
