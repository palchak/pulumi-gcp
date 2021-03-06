# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetManagedZoneResult:
    """
    A collection of values returned by getManagedZone.
    """
    def __init__(__self__, description=None, dns_name=None, id=None, name=None, name_servers=None, project=None, visibility=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A textual description field.
        """
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        __self__.dns_name = dns_name
        """
        The fully qualified DNS name of this zone, e.g. `example.io.`.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        __self__.name_servers = name_servers
        """
        The list of nameservers that will be authoritative for this
        domain. Use NS records to redirect from your DNS provider to these names,
        thus making Google Cloud DNS authoritative for this zone.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        __self__.visibility = visibility
        """
        The zone's visibility: public zones are exposed to the Internet,
        while private zones are visible only to Virtual Private Cloud resources.
        """
class AwaitableGetManagedZoneResult(GetManagedZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedZoneResult(
            description=self.description,
            dns_name=self.dns_name,
            id=self.id,
            name=self.name,
            name_servers=self.name_servers,
            project=self.project,
            visibility=self.visibility)

def get_managed_zone(name=None,project=None,opts=None):
    """
    Provides access to a zone's attributes within Google Cloud DNS.
    For more information see
    [the official documentation](https://cloud.google.com/dns/zones/)
    and
    [API](https://cloud.google.com/dns/api/v1/managedZones).


    :param str name: A unique name for the resource.
    :param str project: The ID of the project for the Google Cloud DNS zone.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:dns/getManagedZone:getManagedZone', __args__, opts=opts).value

    return AwaitableGetManagedZoneResult(
        description=__ret__.get('description'),
        dns_name=__ret__.get('dnsName'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        name_servers=__ret__.get('nameServers'),
        project=__ret__.get('project'),
        visibility=__ret__.get('visibility'))
