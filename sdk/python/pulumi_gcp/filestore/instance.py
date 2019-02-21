# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    create_time: pulumi.Output[str]
    description: pulumi.Output[str]
    etag: pulumi.Output[str]
    file_shares: pulumi.Output[dict]
    labels: pulumi.Output[dict]
    name: pulumi.Output[str]
    networks: pulumi.Output[list]
    project: pulumi.Output[str]
    tier: pulumi.Output[str]
    zone: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, description=None, file_shares=None, labels=None, name=None, networks=None, project=None, tier=None, zone=None, __name__=None, __opts__=None):
        """
        A Google Cloud Filestore instance.
        
        > **Warning:** This resource is in beta, and should be used with the terraform-provider-google-beta provider.
        See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
        
        To get more information about Instance, see:
        
        * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1beta1/projects.locations.instances/create)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/filestore/docs/creating-instances)
            * [Use with Kubernetes](https://cloud.google.com/filestore/docs/accessing-fileshares)
            * [Copying Data In/Out](https://cloud.google.com/filestore/docs/copying-data)
        
        <div class = "oics-button" style="float: right; margin: 0 0 -15px">
          <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=filestore_instance_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
            <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
          </a>
        </div>
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

        __props__['description'] = description

        if file_shares is None:
            raise TypeError('Missing required property file_shares')
        __props__['file_shares'] = file_shares

        __props__['labels'] = labels

        __props__['name'] = name

        if networks is None:
            raise TypeError('Missing required property networks')
        __props__['networks'] = networks

        __props__['project'] = project

        if tier is None:
            raise TypeError('Missing required property tier')
        __props__['tier'] = tier

        if zone is None:
            raise TypeError('Missing required property zone')
        __props__['zone'] = zone

        __props__['create_time'] = None
        __props__['etag'] = None

        super(Instance, __self__).__init__(
            'gcp:filestore/instance:Instance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

