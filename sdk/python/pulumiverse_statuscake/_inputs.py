# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'StatuscakePagespeedCheckAlertConfigArgs',
    'StatuscakePagespeedCheckMonitoredResourceArgs',
    'StatuscakeSslCheckAlertConfigArgs',
    'StatuscakeSslCheckMonitoredResourceArgs',
    'StatuscakeUptimeCheckDnsCheckArgs',
    'StatuscakeUptimeCheckHttpCheckArgs',
    'StatuscakeUptimeCheckHttpCheckBasicAuthenticationArgs',
    'StatuscakeUptimeCheckHttpCheckContentMatchersArgs',
    'StatuscakeUptimeCheckIcmpCheckArgs',
    'StatuscakeUptimeCheckLocationArgs',
    'StatuscakeUptimeCheckMonitoredResourceArgs',
    'StatuscakeUptimeCheckTcpCheckArgs',
    'StatuscakeUptimeCheckTcpCheckAuthenticationArgs',
]

@pulumi.input_type
class StatuscakePagespeedCheckAlertConfigArgs:
    def __init__(__self__, *,
                 alert_bigger: Optional[pulumi.Input[int]] = None,
                 alert_slower: Optional[pulumi.Input[int]] = None,
                 alert_smaller: Optional[pulumi.Input[int]] = None):
        if alert_bigger is not None:
            pulumi.set(__self__, "alert_bigger", alert_bigger)
        if alert_slower is not None:
            pulumi.set(__self__, "alert_slower", alert_slower)
        if alert_smaller is not None:
            pulumi.set(__self__, "alert_smaller", alert_smaller)

    @property
    @pulumi.getter(name="alertBigger")
    def alert_bigger(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alert_bigger")

    @alert_bigger.setter
    def alert_bigger(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alert_bigger", value)

    @property
    @pulumi.getter(name="alertSlower")
    def alert_slower(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alert_slower")

    @alert_slower.setter
    def alert_slower(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alert_slower", value)

    @property
    @pulumi.getter(name="alertSmaller")
    def alert_smaller(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "alert_smaller")

    @alert_smaller.setter
    def alert_smaller(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alert_smaller", value)


@pulumi.input_type
class StatuscakePagespeedCheckMonitoredResourceArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str]):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)


@pulumi.input_type
class StatuscakeSslCheckAlertConfigArgs:
    def __init__(__self__, *,
                 alert_ats: pulumi.Input[Sequence[pulumi.Input[int]]],
                 on_broken: Optional[pulumi.Input[bool]] = None,
                 on_expiry: Optional[pulumi.Input[bool]] = None,
                 on_mixed: Optional[pulumi.Input[bool]] = None,
                 on_reminder: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "alert_ats", alert_ats)
        if on_broken is not None:
            pulumi.set(__self__, "on_broken", on_broken)
        if on_expiry is not None:
            pulumi.set(__self__, "on_expiry", on_expiry)
        if on_mixed is not None:
            pulumi.set(__self__, "on_mixed", on_mixed)
        if on_reminder is not None:
            pulumi.set(__self__, "on_reminder", on_reminder)

    @property
    @pulumi.getter(name="alertAts")
    def alert_ats(self) -> pulumi.Input[Sequence[pulumi.Input[int]]]:
        return pulumi.get(self, "alert_ats")

    @alert_ats.setter
    def alert_ats(self, value: pulumi.Input[Sequence[pulumi.Input[int]]]):
        pulumi.set(self, "alert_ats", value)

    @property
    @pulumi.getter(name="onBroken")
    def on_broken(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "on_broken")

    @on_broken.setter
    def on_broken(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_broken", value)

    @property
    @pulumi.getter(name="onExpiry")
    def on_expiry(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "on_expiry")

    @on_expiry.setter
    def on_expiry(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_expiry", value)

    @property
    @pulumi.getter(name="onMixed")
    def on_mixed(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "on_mixed")

    @on_mixed.setter
    def on_mixed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_mixed", value)

    @property
    @pulumi.getter(name="onReminder")
    def on_reminder(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "on_reminder")

    @on_reminder.setter
    def on_reminder(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "on_reminder", value)


@pulumi.input_type
class StatuscakeSslCheckMonitoredResourceArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 hostname: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "address", address)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def hostname(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hostname")

    @hostname.setter
    def hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hostname", value)


@pulumi.input_type
class StatuscakeUptimeCheckDnsCheckArgs:
    def __init__(__self__, *,
                 dns_ips: pulumi.Input[Sequence[pulumi.Input[str]]],
                 dns_server: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "dns_ips", dns_ips)
        if dns_server is not None:
            pulumi.set(__self__, "dns_server", dns_server)

    @property
    @pulumi.getter(name="dnsIps")
    def dns_ips(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "dns_ips")

    @dns_ips.setter
    def dns_ips(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "dns_ips", value)

    @property
    @pulumi.getter(name="dnsServer")
    def dns_server(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dns_server")

    @dns_server.setter
    def dns_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_server", value)


@pulumi.input_type
class StatuscakeUptimeCheckHttpCheckArgs:
    def __init__(__self__, *,
                 status_codes: pulumi.Input[Sequence[pulumi.Input[str]]],
                 basic_authentication: Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckBasicAuthenticationArgs']] = None,
                 content_matchers: Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckContentMatchersArgs']] = None,
                 enable_cookies: Optional[pulumi.Input[bool]] = None,
                 final_endpoint: Optional[pulumi.Input[str]] = None,
                 follow_redirects: Optional[pulumi.Input[bool]] = None,
                 request_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_method: Optional[pulumi.Input[str]] = None,
                 request_payload: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_payload_raw: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 validate_ssl: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "status_codes", status_codes)
        if basic_authentication is not None:
            pulumi.set(__self__, "basic_authentication", basic_authentication)
        if content_matchers is not None:
            pulumi.set(__self__, "content_matchers", content_matchers)
        if enable_cookies is not None:
            pulumi.set(__self__, "enable_cookies", enable_cookies)
        if final_endpoint is not None:
            pulumi.set(__self__, "final_endpoint", final_endpoint)
        if follow_redirects is not None:
            pulumi.set(__self__, "follow_redirects", follow_redirects)
        if request_headers is not None:
            pulumi.set(__self__, "request_headers", request_headers)
        if request_method is not None:
            pulumi.set(__self__, "request_method", request_method)
        if request_payload is not None:
            pulumi.set(__self__, "request_payload", request_payload)
        if request_payload_raw is not None:
            pulumi.set(__self__, "request_payload_raw", request_payload_raw)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)
        if validate_ssl is not None:
            pulumi.set(__self__, "validate_ssl", validate_ssl)

    @property
    @pulumi.getter(name="statusCodes")
    def status_codes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "status_codes")

    @status_codes.setter
    def status_codes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "status_codes", value)

    @property
    @pulumi.getter(name="basicAuthentication")
    def basic_authentication(self) -> Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckBasicAuthenticationArgs']]:
        return pulumi.get(self, "basic_authentication")

    @basic_authentication.setter
    def basic_authentication(self, value: Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckBasicAuthenticationArgs']]):
        pulumi.set(self, "basic_authentication", value)

    @property
    @pulumi.getter(name="contentMatchers")
    def content_matchers(self) -> Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckContentMatchersArgs']]:
        return pulumi.get(self, "content_matchers")

    @content_matchers.setter
    def content_matchers(self, value: Optional[pulumi.Input['StatuscakeUptimeCheckHttpCheckContentMatchersArgs']]):
        pulumi.set(self, "content_matchers", value)

    @property
    @pulumi.getter(name="enableCookies")
    def enable_cookies(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_cookies")

    @enable_cookies.setter
    def enable_cookies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_cookies", value)

    @property
    @pulumi.getter(name="finalEndpoint")
    def final_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "final_endpoint")

    @final_endpoint.setter
    def final_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "final_endpoint", value)

    @property
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "follow_redirects")

    @follow_redirects.setter
    def follow_redirects(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "follow_redirects", value)

    @property
    @pulumi.getter(name="requestHeaders")
    def request_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "request_headers")

    @request_headers.setter
    def request_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "request_headers", value)

    @property
    @pulumi.getter(name="requestMethod")
    def request_method(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_method")

    @request_method.setter
    def request_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_method", value)

    @property
    @pulumi.getter(name="requestPayload")
    def request_payload(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "request_payload")

    @request_payload.setter
    def request_payload(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "request_payload", value)

    @property
    @pulumi.getter(name="requestPayloadRaw")
    def request_payload_raw(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_payload_raw")

    @request_payload_raw.setter
    def request_payload_raw(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_payload_raw", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)

    @property
    @pulumi.getter(name="validateSsl")
    def validate_ssl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "validate_ssl")

    @validate_ssl.setter
    def validate_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_ssl", value)


@pulumi.input_type
class StatuscakeUptimeCheckHttpCheckBasicAuthenticationArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class StatuscakeUptimeCheckHttpCheckContentMatchersArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 include_headers: Optional[pulumi.Input[bool]] = None,
                 matcher: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "content", content)
        if include_headers is not None:
            pulumi.set(__self__, "include_headers", include_headers)
        if matcher is not None:
            pulumi.set(__self__, "matcher", matcher)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="includeHeaders")
    def include_headers(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "include_headers")

    @include_headers.setter
    def include_headers(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_headers", value)

    @property
    @pulumi.getter
    def matcher(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "matcher")

    @matcher.setter
    def matcher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "matcher", value)


@pulumi.input_type
class StatuscakeUptimeCheckIcmpCheckArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class StatuscakeUptimeCheckLocationArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ipv4: Optional[pulumi.Input[str]] = None,
                 ipv6: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 region_code: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ipv4 is not None:
            pulumi.set(__self__, "ipv4", ipv4)
        if ipv6 is not None:
            pulumi.set(__self__, "ipv6", ipv6)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if region_code is not None:
            pulumi.set(__self__, "region_code", region_code)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipv4")

    @ipv4.setter
    def ipv4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv4", value)

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipv6")

    @ipv6.setter
    def ipv6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="regionCode")
    def region_code(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "region_code")

    @region_code.setter
    def region_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_code", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class StatuscakeUptimeCheckMonitoredResourceArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 host: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "address", address)
        if host is not None:
            pulumi.set(__self__, "host", host)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def host(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host")

    @host.setter
    def host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host", value)


@pulumi.input_type
class StatuscakeUptimeCheckTcpCheckArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 authentication: Optional[pulumi.Input['StatuscakeUptimeCheckTcpCheckAuthenticationArgs']] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "port", port)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def authentication(self) -> Optional[pulumi.Input['StatuscakeUptimeCheckTcpCheckAuthenticationArgs']]:
        return pulumi.get(self, "authentication")

    @authentication.setter
    def authentication(self, value: Optional[pulumi.Input['StatuscakeUptimeCheckTcpCheckAuthenticationArgs']]):
        pulumi.set(self, "authentication", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout", value)


@pulumi.input_type
class StatuscakeUptimeCheckTcpCheckAuthenticationArgs:
    def __init__(__self__, *,
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


