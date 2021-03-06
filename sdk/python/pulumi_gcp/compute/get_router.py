# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRouterResult:
    """
    A collection of values returned by getRouter.
    """
    def __init__(__self__, bgps=None, creation_timestamp=None, description=None, id=None, name=None, network=None, project=None, region=None, self_link=None):
        if bgps and not isinstance(bgps, list):
            raise TypeError("Expected argument 'bgps' to be a list")
        __self__.bgps = bgps
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        __self__.creation_timestamp = creation_timestamp
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        __self__.network = network
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
class AwaitableGetRouterResult(GetRouterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouterResult(
            bgps=self.bgps,
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            id=self.id,
            name=self.name,
            network=self.network,
            project=self.project,
            region=self.region,
            self_link=self.self_link)

def get_router(name=None,network=None,project=None,region=None,opts=None):
    """
    Get a router within GCE from its name and VPC.




    :param str name: The name of the router.
    :param str network: The VPC network on which this router lives.
    :param str project: The ID of the project in which the resource
           belongs. If it is not provided, the provider project is used.
    :param str region: The region this router has been created in. If
           unspecified, this defaults to the region configured in the provider.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['network'] = network
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRouter:getRouter', __args__, opts=opts).value

    return AwaitableGetRouterResult(
        bgps=__ret__.get('bgps'),
        creation_timestamp=__ret__.get('creationTimestamp'),
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        network=__ret__.get('network'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        self_link=__ret__.get('selfLink'))
