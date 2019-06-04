# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RouterInterface(pulumi.CustomResource):
    interconnect_attachment: pulumi.Output[str]
    """
    The name or resource link to the
    VLAN interconnect for this interface. Changing this forces a new interface to
    be created. Only one of `vpn_tunnel` and `interconnect_attachment` can be
    specified.
    """
    ip_range: pulumi.Output[str]
    """
    IP address and range of the interface. The IP range must be
    in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
    """
    name: pulumi.Output[str]
    """
    A unique name for the interface, required by GCE. Changing
    this forces a new interface to be created.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which this interface's router belongs. If it
    is not provided, the provider project is used. Changing this forces a new interface to be created.
    """
    region: pulumi.Output[str]
    """
    The region this interface's router sits in. If not specified,
    the project region will be used. Changing this forces a new interface to be
    created.
    """
    router: pulumi.Output[str]
    """
    The name of the router this interface will be attached to.
    Changing this forces a new interface to be created.
    """
    vpn_tunnel: pulumi.Output[str]
    """
    The name or resource link to the VPN tunnel this
    interface will be linked to. Changing this forces a new interface to be created. Only
    one of `vpn_tunnel` and `interconnect_attachment` can be specified.
    """
    def __init__(__self__, resource_name, opts=None, interconnect_attachment=None, ip_range=None, name=None, project=None, region=None, router=None, vpn_tunnel=None, __name__=None, __opts__=None):
        """
        Manages a Cloud Router interface. For more information see
        [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
        and
        [API](https://cloud.google.com/compute/docs/reference/latest/routers).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] interconnect_attachment: The name or resource link to the
               VLAN interconnect for this interface. Changing this forces a new interface to
               be created. Only one of `vpn_tunnel` and `interconnect_attachment` can be
               specified.
        :param pulumi.Input[str] ip_range: IP address and range of the interface. The IP range must be
               in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
        :param pulumi.Input[str] name: A unique name for the interface, required by GCE. Changing
               this forces a new interface to be created.
        :param pulumi.Input[str] project: The ID of the project in which this interface's router belongs. If it
               is not provided, the provider project is used. Changing this forces a new interface to be created.
        :param pulumi.Input[str] region: The region this interface's router sits in. If not specified,
               the project region will be used. Changing this forces a new interface to be
               created.
        :param pulumi.Input[str] router: The name of the router this interface will be attached to.
               Changing this forces a new interface to be created.
        :param pulumi.Input[str] vpn_tunnel: The name or resource link to the VPN tunnel this
               interface will be linked to. Changing this forces a new interface to be created. Only
               one of `vpn_tunnel` and `interconnect_attachment` can be specified.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['interconnect_attachment'] = interconnect_attachment

        __props__['ip_range'] = ip_range

        __props__['name'] = name

        __props__['project'] = project

        __props__['region'] = region

        if router is None:
            raise TypeError("Missing required property 'router'")
        __props__['router'] = router

        __props__['vpn_tunnel'] = vpn_tunnel

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RouterInterface, __self__).__init__(
            'gcp:compute/routerInterface:RouterInterface',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

