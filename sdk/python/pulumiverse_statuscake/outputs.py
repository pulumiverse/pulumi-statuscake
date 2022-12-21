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

__all__ = [
    'StatuscakePagespeedCheckAlertConfig',
    'StatuscakePagespeedCheckMonitoredResource',
    'StatuscakeSslCheckAlertConfig',
    'StatuscakeSslCheckMonitoredResource',
    'StatuscakeUptimeCheckDnsCheck',
    'StatuscakeUptimeCheckHttpCheck',
    'StatuscakeUptimeCheckHttpCheckBasicAuthentication',
    'StatuscakeUptimeCheckHttpCheckContentMatchers',
    'StatuscakeUptimeCheckIcmpCheck',
    'StatuscakeUptimeCheckLocation',
    'StatuscakeUptimeCheckMonitoredResource',
    'StatuscakeUptimeCheckTcpCheck',
    'StatuscakeUptimeCheckTcpCheckAuthentication',
    'GetStatuscakePagespeedMonitoringLocationsLocationResult',
    'GetStatuscakeUptimeMonitoringLocationsLocationResult',
]

@pulumi.output_type
class StatuscakePagespeedCheckAlertConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alertBigger":
            suggest = "alert_bigger"
        elif key == "alertSlower":
            suggest = "alert_slower"
        elif key == "alertSmaller":
            suggest = "alert_smaller"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakePagespeedCheckAlertConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakePagespeedCheckAlertConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakePagespeedCheckAlertConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alert_bigger: Optional[int] = None,
                 alert_slower: Optional[int] = None,
                 alert_smaller: Optional[int] = None):
        if alert_bigger is not None:
            pulumi.set(__self__, "alert_bigger", alert_bigger)
        if alert_slower is not None:
            pulumi.set(__self__, "alert_slower", alert_slower)
        if alert_smaller is not None:
            pulumi.set(__self__, "alert_smaller", alert_smaller)

    @property
    @pulumi.getter(name="alertBigger")
    def alert_bigger(self) -> Optional[int]:
        return pulumi.get(self, "alert_bigger")

    @property
    @pulumi.getter(name="alertSlower")
    def alert_slower(self) -> Optional[int]:
        return pulumi.get(self, "alert_slower")

    @property
    @pulumi.getter(name="alertSmaller")
    def alert_smaller(self) -> Optional[int]:
        return pulumi.get(self, "alert_smaller")


@pulumi.output_type
class StatuscakePagespeedCheckMonitoredResource(dict):
    def __init__(__self__, *,
                 address: str):
        pulumi.set(__self__, "address", address)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")


@pulumi.output_type
class StatuscakeSslCheckAlertConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "alertAts":
            suggest = "alert_ats"
        elif key == "onBroken":
            suggest = "on_broken"
        elif key == "onExpiry":
            suggest = "on_expiry"
        elif key == "onMixed":
            suggest = "on_mixed"
        elif key == "onReminder":
            suggest = "on_reminder"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakeSslCheckAlertConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakeSslCheckAlertConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakeSslCheckAlertConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 alert_ats: Sequence[int],
                 on_broken: Optional[bool] = None,
                 on_expiry: Optional[bool] = None,
                 on_mixed: Optional[bool] = None,
                 on_reminder: Optional[bool] = None):
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
    def alert_ats(self) -> Sequence[int]:
        return pulumi.get(self, "alert_ats")

    @property
    @pulumi.getter(name="onBroken")
    def on_broken(self) -> Optional[bool]:
        return pulumi.get(self, "on_broken")

    @property
    @pulumi.getter(name="onExpiry")
    def on_expiry(self) -> Optional[bool]:
        return pulumi.get(self, "on_expiry")

    @property
    @pulumi.getter(name="onMixed")
    def on_mixed(self) -> Optional[bool]:
        return pulumi.get(self, "on_mixed")

    @property
    @pulumi.getter(name="onReminder")
    def on_reminder(self) -> Optional[bool]:
        return pulumi.get(self, "on_reminder")


@pulumi.output_type
class StatuscakeSslCheckMonitoredResource(dict):
    def __init__(__self__, *,
                 address: str,
                 hostname: Optional[str] = None):
        pulumi.set(__self__, "address", address)
        if hostname is not None:
            pulumi.set(__self__, "hostname", hostname)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def hostname(self) -> Optional[str]:
        return pulumi.get(self, "hostname")


@pulumi.output_type
class StatuscakeUptimeCheckDnsCheck(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "dnsIps":
            suggest = "dns_ips"
        elif key == "dnsServer":
            suggest = "dns_server"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakeUptimeCheckDnsCheck. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakeUptimeCheckDnsCheck.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakeUptimeCheckDnsCheck.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 dns_ips: Sequence[str],
                 dns_server: Optional[str] = None):
        pulumi.set(__self__, "dns_ips", dns_ips)
        if dns_server is not None:
            pulumi.set(__self__, "dns_server", dns_server)

    @property
    @pulumi.getter(name="dnsIps")
    def dns_ips(self) -> Sequence[str]:
        return pulumi.get(self, "dns_ips")

    @property
    @pulumi.getter(name="dnsServer")
    def dns_server(self) -> Optional[str]:
        return pulumi.get(self, "dns_server")


@pulumi.output_type
class StatuscakeUptimeCheckHttpCheck(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "statusCodes":
            suggest = "status_codes"
        elif key == "basicAuthentication":
            suggest = "basic_authentication"
        elif key == "contentMatchers":
            suggest = "content_matchers"
        elif key == "enableCookies":
            suggest = "enable_cookies"
        elif key == "finalEndpoint":
            suggest = "final_endpoint"
        elif key == "followRedirects":
            suggest = "follow_redirects"
        elif key == "requestHeaders":
            suggest = "request_headers"
        elif key == "requestMethod":
            suggest = "request_method"
        elif key == "requestPayload":
            suggest = "request_payload"
        elif key == "requestPayloadRaw":
            suggest = "request_payload_raw"
        elif key == "userAgent":
            suggest = "user_agent"
        elif key == "validateSsl":
            suggest = "validate_ssl"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakeUptimeCheckHttpCheck. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakeUptimeCheckHttpCheck.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakeUptimeCheckHttpCheck.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 status_codes: Sequence[str],
                 basic_authentication: Optional['outputs.StatuscakeUptimeCheckHttpCheckBasicAuthentication'] = None,
                 content_matchers: Optional['outputs.StatuscakeUptimeCheckHttpCheckContentMatchers'] = None,
                 enable_cookies: Optional[bool] = None,
                 final_endpoint: Optional[str] = None,
                 follow_redirects: Optional[bool] = None,
                 request_headers: Optional[Mapping[str, str]] = None,
                 request_method: Optional[str] = None,
                 request_payload: Optional[Mapping[str, str]] = None,
                 request_payload_raw: Optional[str] = None,
                 timeout: Optional[int] = None,
                 user_agent: Optional[str] = None,
                 validate_ssl: Optional[bool] = None):
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
    def status_codes(self) -> Sequence[str]:
        return pulumi.get(self, "status_codes")

    @property
    @pulumi.getter(name="basicAuthentication")
    def basic_authentication(self) -> Optional['outputs.StatuscakeUptimeCheckHttpCheckBasicAuthentication']:
        return pulumi.get(self, "basic_authentication")

    @property
    @pulumi.getter(name="contentMatchers")
    def content_matchers(self) -> Optional['outputs.StatuscakeUptimeCheckHttpCheckContentMatchers']:
        return pulumi.get(self, "content_matchers")

    @property
    @pulumi.getter(name="enableCookies")
    def enable_cookies(self) -> Optional[bool]:
        return pulumi.get(self, "enable_cookies")

    @property
    @pulumi.getter(name="finalEndpoint")
    def final_endpoint(self) -> Optional[str]:
        return pulumi.get(self, "final_endpoint")

    @property
    @pulumi.getter(name="followRedirects")
    def follow_redirects(self) -> Optional[bool]:
        return pulumi.get(self, "follow_redirects")

    @property
    @pulumi.getter(name="requestHeaders")
    def request_headers(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "request_headers")

    @property
    @pulumi.getter(name="requestMethod")
    def request_method(self) -> Optional[str]:
        return pulumi.get(self, "request_method")

    @property
    @pulumi.getter(name="requestPayload")
    def request_payload(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "request_payload")

    @property
    @pulumi.getter(name="requestPayloadRaw")
    def request_payload_raw(self) -> Optional[str]:
        return pulumi.get(self, "request_payload_raw")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[str]:
        return pulumi.get(self, "user_agent")

    @property
    @pulumi.getter(name="validateSsl")
    def validate_ssl(self) -> Optional[bool]:
        return pulumi.get(self, "validate_ssl")


@pulumi.output_type
class StatuscakeUptimeCheckHttpCheckBasicAuthentication(dict):
    def __init__(__self__, *,
                 password: str,
                 username: str):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


@pulumi.output_type
class StatuscakeUptimeCheckHttpCheckContentMatchers(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "includeHeaders":
            suggest = "include_headers"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakeUptimeCheckHttpCheckContentMatchers. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakeUptimeCheckHttpCheckContentMatchers.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakeUptimeCheckHttpCheckContentMatchers.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 content: str,
                 include_headers: Optional[bool] = None,
                 matcher: Optional[str] = None):
        pulumi.set(__self__, "content", content)
        if include_headers is not None:
            pulumi.set(__self__, "include_headers", include_headers)
        if matcher is not None:
            pulumi.set(__self__, "matcher", matcher)

    @property
    @pulumi.getter
    def content(self) -> str:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="includeHeaders")
    def include_headers(self) -> Optional[bool]:
        return pulumi.get(self, "include_headers")

    @property
    @pulumi.getter
    def matcher(self) -> Optional[str]:
        return pulumi.get(self, "matcher")


@pulumi.output_type
class StatuscakeUptimeCheckIcmpCheck(dict):
    def __init__(__self__, *,
                 enabled: Optional[bool] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        return pulumi.get(self, "enabled")


@pulumi.output_type
class StatuscakeUptimeCheckLocation(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "regionCode":
            suggest = "region_code"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatuscakeUptimeCheckLocation. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatuscakeUptimeCheckLocation.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatuscakeUptimeCheckLocation.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 ipv4: Optional[str] = None,
                 ipv6: Optional[str] = None,
                 region: Optional[str] = None,
                 region_code: Optional[str] = None,
                 status: Optional[str] = None):
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
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ipv4(self) -> Optional[str]:
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> Optional[str]:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionCode")
    def region_code(self) -> Optional[str]:
        return pulumi.get(self, "region_code")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        return pulumi.get(self, "status")


@pulumi.output_type
class StatuscakeUptimeCheckMonitoredResource(dict):
    def __init__(__self__, *,
                 address: str,
                 host: Optional[str] = None):
        pulumi.set(__self__, "address", address)
        if host is not None:
            pulumi.set(__self__, "host", host)

    @property
    @pulumi.getter
    def address(self) -> str:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")


@pulumi.output_type
class StatuscakeUptimeCheckTcpCheck(dict):
    def __init__(__self__, *,
                 port: int,
                 authentication: Optional['outputs.StatuscakeUptimeCheckTcpCheckAuthentication'] = None,
                 protocol: Optional[str] = None,
                 timeout: Optional[int] = None):
        pulumi.set(__self__, "port", port)
        if authentication is not None:
            pulumi.set(__self__, "authentication", authentication)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def authentication(self) -> Optional['outputs.StatuscakeUptimeCheckTcpCheckAuthentication']:
        return pulumi.get(self, "authentication")

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def timeout(self) -> Optional[int]:
        return pulumi.get(self, "timeout")


@pulumi.output_type
class StatuscakeUptimeCheckTcpCheckAuthentication(dict):
    def __init__(__self__, *,
                 password: str,
                 username: str):
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def password(self) -> str:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


@pulumi.output_type
class GetStatuscakePagespeedMonitoringLocationsLocationResult(dict):
    def __init__(__self__, *,
                 description: str,
                 ipv4: str,
                 ipv6: str,
                 region: str,
                 region_code: str,
                 status: str):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "ipv4", ipv4)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "region_code", region_code)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ipv4(self) -> str:
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionCode")
    def region_code(self) -> str:
        return pulumi.get(self, "region_code")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


@pulumi.output_type
class GetStatuscakeUptimeMonitoringLocationsLocationResult(dict):
    def __init__(__self__, *,
                 description: str,
                 ipv4: str,
                 ipv6: str,
                 region: str,
                 region_code: str,
                 status: str):
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "ipv4", ipv4)
        pulumi.set(__self__, "ipv6", ipv6)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "region_code", region_code)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ipv4(self) -> str:
        return pulumi.get(self, "ipv4")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="regionCode")
    def region_code(self) -> str:
        return pulumi.get(self, "region_code")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


