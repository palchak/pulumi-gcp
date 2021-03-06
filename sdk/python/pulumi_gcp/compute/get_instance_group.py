# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetInstanceGroupResult:
    """
    A collection of values returned by getInstanceGroup.
    """
    def __init__(__self__, description=None, id=None, instances=None, name=None, named_ports=None, network=None, project=None, self_link=None, size=None, zone=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Textual description of the instance group.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        """
        List of instances in the group.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if named_ports and not isinstance(named_ports, list):
            raise TypeError("Expected argument 'named_ports' to be a list")
        __self__.named_ports = named_ports
        """
        List of named ports in the group.
        """
        if network and not isinstance(network, str):
            raise TypeError("Expected argument 'network' to be a str")
        __self__.network = network
        """
        The URL of the network the instance group is in.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        The URI of the resource.
        """
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The number of instances in the group.
        """
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        __self__.zone = zone
class AwaitableGetInstanceGroupResult(GetInstanceGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceGroupResult(
            description=self.description,
            id=self.id,
            instances=self.instances,
            name=self.name,
            named_ports=self.named_ports,
            network=self.network,
            project=self.project,
            self_link=self.self_link,
            size=self.size,
            zone=self.zone)

def get_instance_group(name=None,project=None,self_link=None,zone=None,opts=None):
    """
    Get a Compute Instance Group within GCE.
    For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
    and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)


    :param str name: The name of the instance group. Either `name` or `self_link` must be provided.
    :param str project: The ID of the project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str self_link: The self link of the instance group. Either `name` or `self_link` must be provided.
    :param str zone: The zone of the instance group. If referencing the instance group by name
           and `zone` is not provided, the provider zone is used.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    __args__['selfLink'] = self_link
    __args__['zone'] = zone
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:compute/getInstanceGroup:getInstanceGroup', __args__, opts=opts).value

    return AwaitableGetInstanceGroupResult(
        description=__ret__.get('description'),
        id=__ret__.get('id'),
        instances=__ret__.get('instances'),
        name=__ret__.get('name'),
        named_ports=__ret__.get('namedPorts'),
        network=__ret__.get('network'),
        project=__ret__.get('project'),
        self_link=__ret__.get('selfLink'),
        size=__ret__.get('size'),
        zone=__ret__.get('zone'))
