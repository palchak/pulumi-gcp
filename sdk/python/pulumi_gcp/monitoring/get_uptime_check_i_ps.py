# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetUptimeCheckIPsResult:
    """
    A collection of values returned by getUptimeCheckIPs.
    """
    def __init__(__self__, id=None, uptime_check_ips=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if uptime_check_ips and not isinstance(uptime_check_ips, list):
            raise TypeError("Expected argument 'uptime_check_ips' to be a list")
        __self__.uptime_check_ips = uptime_check_ips
class AwaitableGetUptimeCheckIPsResult(GetUptimeCheckIPsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUptimeCheckIPsResult(
            id=self.id,
            uptime_check_ips=self.uptime_check_ips)

def get_uptime_check_i_ps(opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:monitoring/getUptimeCheckIPs:getUptimeCheckIPs', __args__, opts=opts).value

    return AwaitableGetUptimeCheckIPsResult(
        id=__ret__.get('id'),
        uptime_check_ips=__ret__.get('uptimeCheckIps'))
