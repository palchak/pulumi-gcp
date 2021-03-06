# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    clusters: pulumi.Output[list]
    """
    A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.

      * `cluster_id` (`str`) - The ID of the Cloud Bigtable cluster.
      * `num_nodes` (`float`) - The number of nodes in your Cloud Bigtable cluster.
        Required, with a minimum of `3` for a `PRODUCTION` instance. Must be left unset
        for a `DEVELOPMENT` instance.
      * `storageType` (`str`) - The storage type to use. One of `"SSD"` or
        `"HDD"`. Defaults to `"SSD"`.
      * `zone` (`str`) - The zone to create the Cloud Bigtable cluster in. Each
        cluster must have a different zone in the same region. Zones that support
        Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
    """
    display_name: pulumi.Output[str]
    """
    The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
    """
    instance_type: pulumi.Output[str]
    """
    The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
    """
    name: pulumi.Output[str]
    """
    The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs. If it
    is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, clusters=None, display_name=None, instance_type=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Google Bigtable instance. For more information see
        [the official documentation](https://cloud.google.com/bigtable/) and
        [API](https://cloud.google.com/bigtable/docs/go/reference).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] clusters: A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.

        The **clusters** object supports the following:

          * `cluster_id` (`pulumi.Input[str]`) - The ID of the Cloud Bigtable cluster.
          * `num_nodes` (`pulumi.Input[float]`) - The number of nodes in your Cloud Bigtable cluster.
            Required, with a minimum of `3` for a `PRODUCTION` instance. Must be left unset
            for a `DEVELOPMENT` instance.
          * `storageType` (`pulumi.Input[str]`) - The storage type to use. One of `"SSD"` or
            `"HDD"`. Defaults to `"SSD"`.
          * `zone` (`pulumi.Input[str]`) - The zone to create the Cloud Bigtable cluster in. Each
            cluster must have a different zone in the same region. Zones that support
            Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
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

            __props__['clusters'] = clusters
            __props__['display_name'] = display_name
            __props__['instance_type'] = instance_type
            __props__['name'] = name
            __props__['project'] = project
        super(Instance, __self__).__init__(
            'gcp:bigtable/instance:Instance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, clusters=None, display_name=None, instance_type=None, name=None, project=None):
        """
        Get an existing Instance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] clusters: A block of cluster configuration options. This can be specified 1 or 2 times. See structure below.
        :param pulumi.Input[str] display_name: The human-readable display name of the Bigtable instance. Defaults to the instance `name`.
        :param pulumi.Input[str] instance_type: The instance type to create. One of `"DEVELOPMENT"` or `"PRODUCTION"`. Defaults to `"PRODUCTION"`.
        :param pulumi.Input[str] name: The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs. If it
               is not provided, the provider project is used.

        The **clusters** object supports the following:

          * `cluster_id` (`pulumi.Input[str]`) - The ID of the Cloud Bigtable cluster.
          * `num_nodes` (`pulumi.Input[float]`) - The number of nodes in your Cloud Bigtable cluster.
            Required, with a minimum of `3` for a `PRODUCTION` instance. Must be left unset
            for a `DEVELOPMENT` instance.
          * `storageType` (`pulumi.Input[str]`) - The storage type to use. One of `"SSD"` or
            `"HDD"`. Defaults to `"SSD"`.
          * `zone` (`pulumi.Input[str]`) - The zone to create the Cloud Bigtable cluster in. Each
            cluster must have a different zone in the same region. Zones that support
            Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["clusters"] = clusters
        __props__["display_name"] = display_name
        __props__["instance_type"] = instance_type
        __props__["name"] = name
        __props__["project"] = project
        return Instance(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

