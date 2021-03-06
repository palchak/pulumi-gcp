# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Client(pulumi.CustomResource):
    brand: pulumi.Output[str]
    """
    Identifier of the brand to which this client is attached to. The format is
    'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
    """
    client_id: pulumi.Output[str]
    """
    Output only. Unique identifier of the OAuth client.
    """
    display_name: pulumi.Output[str]
    """
    Human-friendly name given to the OAuth client.
    """
    secret: pulumi.Output[str]
    """
    Output only. Client secret of the OAuth client.
    """
    def __init__(__self__, resource_name, opts=None, brand=None, display_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Contains the data that describes an Identity Aware Proxy owned client.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] brand: Identifier of the brand to which this client is attached to. The format is
               'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
        :param pulumi.Input[str] display_name: Human-friendly name given to the OAuth client.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if brand is None:
                raise TypeError("Missing required property 'brand'")
            __props__['brand'] = brand
            if display_name is None:
                raise TypeError("Missing required property 'display_name'")
            __props__['display_name'] = display_name
            __props__['client_id'] = None
            __props__['secret'] = None
        super(Client, __self__).__init__(
            'gcp:iap/client:Client',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, brand=None, client_id=None, display_name=None, secret=None):
        """
        Get an existing Client resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] brand: Identifier of the brand to which this client is attached to. The format is
               'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
        :param pulumi.Input[str] client_id: Output only. Unique identifier of the OAuth client.
        :param pulumi.Input[str] display_name: Human-friendly name given to the OAuth client.
        :param pulumi.Input[str] secret: Output only. Client secret of the OAuth client.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["brand"] = brand
        __props__["client_id"] = client_id
        __props__["display_name"] = display_name
        __props__["secret"] = secret
        return Client(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

