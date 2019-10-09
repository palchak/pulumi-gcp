# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Dataset(pulumi.CustomResource):
    accesses: pulumi.Output[list]
    creation_time: pulumi.Output[float]
    dataset_id: pulumi.Output[str]
    default_encryption_configuration: pulumi.Output[dict]
    default_partition_expiration_ms: pulumi.Output[float]
    default_table_expiration_ms: pulumi.Output[float]
    delete_contents_on_destroy: pulumi.Output[bool]
    """
    If set to `true`, delete all the tables in the
    dataset when destroying the resource; otherwise,
    destroying the resource will fail if tables are present.
    """
    description: pulumi.Output[str]
    etag: pulumi.Output[str]
    friendly_name: pulumi.Output[str]
    labels: pulumi.Output[dict]
    last_modified_time: pulumi.Output[float]
    location: pulumi.Output[str]
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    self_link: pulumi.Output[str]
    """
    The URI of the created resource.
    """
    def __init__(__self__, resource_name, opts=None, accesses=None, dataset_id=None, default_encryption_configuration=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, friendly_name=None, labels=None, location=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Dataset resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        
        The **accesses** object supports the following:
        
          * `domain` (`pulumi.Input[str]`)
          * `groupByEmail` (`pulumi.Input[str]`)
          * `role` (`pulumi.Input[str]`)
          * `specialGroup` (`pulumi.Input[str]`)
          * `userByEmail` (`pulumi.Input[str]`)
          * `view` (`pulumi.Input[dict]`)
        
            * `dataset_id` (`pulumi.Input[str]`)
            * `projectId` (`pulumi.Input[str]`)
            * `table_id` (`pulumi.Input[str]`)
        
        The **default_encryption_configuration** object supports the following:
        
          * `kmsKeyName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_dataset.html.markdown.
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

            __props__['accesses'] = accesses
            if dataset_id is None:
                raise TypeError("Missing required property 'dataset_id'")
            __props__['dataset_id'] = dataset_id
            __props__['default_encryption_configuration'] = default_encryption_configuration
            __props__['default_partition_expiration_ms'] = default_partition_expiration_ms
            __props__['default_table_expiration_ms'] = default_table_expiration_ms
            __props__['delete_contents_on_destroy'] = delete_contents_on_destroy
            __props__['description'] = description
            __props__['friendly_name'] = friendly_name
            __props__['labels'] = labels
            __props__['location'] = location
            __props__['project'] = project
            __props__['creation_time'] = None
            __props__['etag'] = None
            __props__['last_modified_time'] = None
            __props__['self_link'] = None
        super(Dataset, __self__).__init__(
            'gcp:bigquery/dataset:Dataset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accesses=None, creation_time=None, dataset_id=None, default_encryption_configuration=None, default_partition_expiration_ms=None, default_table_expiration_ms=None, delete_contents_on_destroy=None, description=None, etag=None, friendly_name=None, labels=None, last_modified_time=None, location=None, project=None, self_link=None):
        """
        Get an existing Dataset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_contents_on_destroy: If set to `true`, delete all the tables in the
               dataset when destroying the resource; otherwise,
               destroying the resource will fail if tables are present.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] self_link: The URI of the created resource.
        
        The **accesses** object supports the following:
        
          * `domain` (`pulumi.Input[str]`)
          * `groupByEmail` (`pulumi.Input[str]`)
          * `role` (`pulumi.Input[str]`)
          * `specialGroup` (`pulumi.Input[str]`)
          * `userByEmail` (`pulumi.Input[str]`)
          * `view` (`pulumi.Input[dict]`)
        
            * `dataset_id` (`pulumi.Input[str]`)
            * `projectId` (`pulumi.Input[str]`)
            * `table_id` (`pulumi.Input[str]`)
        
        The **default_encryption_configuration** object supports the following:
        
          * `kmsKeyName` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/bigquery_dataset.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["accesses"] = accesses
        __props__["creation_time"] = creation_time
        __props__["dataset_id"] = dataset_id
        __props__["default_encryption_configuration"] = default_encryption_configuration
        __props__["default_partition_expiration_ms"] = default_partition_expiration_ms
        __props__["default_table_expiration_ms"] = default_table_expiration_ms
        __props__["delete_contents_on_destroy"] = delete_contents_on_destroy
        __props__["description"] = description
        __props__["etag"] = etag
        __props__["friendly_name"] = friendly_name
        __props__["labels"] = labels
        __props__["last_modified_time"] = last_modified_time
        __props__["location"] = location
        __props__["project"] = project
        __props__["self_link"] = self_link
        return Dataset(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

