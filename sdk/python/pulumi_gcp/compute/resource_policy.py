# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ResourcePolicy(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63
    characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
    expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all following
    characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    region: pulumi.Output[str]
    """
    Region where resource policy resides.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    snapshot_schedule_policy: pulumi.Output[dict]
    """
    Policy for creating snapshots of persistent disks.

      * `retention_policy` (`dict`)
        * `maxRetentionDays` (`float`)
        * `onSourceDiskDelete` (`str`)

      * `schedule` (`dict`)
        * `dailySchedule` (`dict`)
          * `daysInCycle` (`float`)
          * `startTime` (`str`)

        * `hourlySchedule` (`dict`)
          * `hoursInCycle` (`float`)
          * `startTime` (`str`)

        * `weeklySchedule` (`dict`)
          * `dayOfWeeks` (`list`)
            * `day` (`str`)
            * `startTime` (`str`)

      * `snapshotProperties` (`dict`)
        * `guestFlush` (`bool`)
        * `labels` (`dict`)
        * `storageLocations` (`str`)
    """
    def __init__(__self__, resource_name, opts=None, name=None, project=None, region=None, snapshot_schedule_policy=None, __props__=None, __name__=None, __opts__=None):
        """
        A policy that can be attached to a resource to specify or schedule actions on that resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63
               characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
               expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where resource policy resides.
        :param pulumi.Input[dict] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.

        The **snapshot_schedule_policy** object supports the following:

          * `retention_policy` (`pulumi.Input[dict]`)
            * `maxRetentionDays` (`pulumi.Input[float]`)
            * `onSourceDiskDelete` (`pulumi.Input[str]`)

          * `schedule` (`pulumi.Input[dict]`)
            * `dailySchedule` (`pulumi.Input[dict]`)
              * `daysInCycle` (`pulumi.Input[float]`)
              * `startTime` (`pulumi.Input[str]`)

            * `hourlySchedule` (`pulumi.Input[dict]`)
              * `hoursInCycle` (`pulumi.Input[float]`)
              * `startTime` (`pulumi.Input[str]`)

            * `weeklySchedule` (`pulumi.Input[dict]`)
              * `dayOfWeeks` (`pulumi.Input[list]`)
                * `day` (`pulumi.Input[str]`)
                * `startTime` (`pulumi.Input[str]`)

          * `snapshotProperties` (`pulumi.Input[dict]`)
            * `guestFlush` (`pulumi.Input[bool]`)
            * `labels` (`pulumi.Input[dict]`)
            * `storageLocations` (`pulumi.Input[str]`)
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

            __props__['name'] = name
            __props__['project'] = project
            __props__['region'] = region
            __props__['snapshot_schedule_policy'] = snapshot_schedule_policy
            __props__['self_link'] = None
        super(ResourcePolicy, __self__).__init__(
            'gcp:compute/resourcePolicy:ResourcePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, project=None, region=None, self_link=None, snapshot_schedule_policy=None):
        """
        Get an existing ResourcePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63
               characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular
               expression '[a-z]([-a-z0-9]*[a-z0-9])'? which means the first character must be a lowercase letter, and all following
               characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] region: Region where resource policy resides.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        :param pulumi.Input[dict] snapshot_schedule_policy: Policy for creating snapshots of persistent disks.

        The **snapshot_schedule_policy** object supports the following:

          * `retention_policy` (`pulumi.Input[dict]`)
            * `maxRetentionDays` (`pulumi.Input[float]`)
            * `onSourceDiskDelete` (`pulumi.Input[str]`)

          * `schedule` (`pulumi.Input[dict]`)
            * `dailySchedule` (`pulumi.Input[dict]`)
              * `daysInCycle` (`pulumi.Input[float]`)
              * `startTime` (`pulumi.Input[str]`)

            * `hourlySchedule` (`pulumi.Input[dict]`)
              * `hoursInCycle` (`pulumi.Input[float]`)
              * `startTime` (`pulumi.Input[str]`)

            * `weeklySchedule` (`pulumi.Input[dict]`)
              * `dayOfWeeks` (`pulumi.Input[list]`)
                * `day` (`pulumi.Input[str]`)
                * `startTime` (`pulumi.Input[str]`)

          * `snapshotProperties` (`pulumi.Input[dict]`)
            * `guestFlush` (`pulumi.Input[bool]`)
            * `labels` (`pulumi.Input[dict]`)
            * `storageLocations` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["project"] = project
        __props__["region"] = region
        __props__["self_link"] = self_link
        __props__["snapshot_schedule_policy"] = snapshot_schedule_policy
        return ResourcePolicy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

