# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IAMCustomRole(pulumi.CustomResource):
    deleted: pulumi.Output[bool]
    description: pulumi.Output[str]
    """
    A human-readable description for the role.
    """
    permissions: pulumi.Output[list]
    """
    The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
    """
    project: pulumi.Output[str]
    """
    The project that the service account will be created in.
    Defaults to the provider project configuration.
    """
    role_id: pulumi.Output[str]
    """
    The role id to use for this role.
    """
    stage: pulumi.Output[str]
    """
    The current launch stage of the role.
    Defaults to `GA`.
    List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
    """
    title: pulumi.Output[str]
    """
    A human-readable title for the role.
    """
    def __init__(__self__, __name__, __opts__=None, deleted=None, description=None, permissions=None, project=None, role_id=None, stage=None, title=None):
        """
        Allows management of a customized Cloud IAM project role. For more information see
        [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
        and
        [API](https://cloud.google.com/iam/reference/rest/v1/projects.roles).
        
        > **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
         from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
         same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
         after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
         made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
         by Terraform, and new roles cannot share that name.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] deleted
        :param pulumi.Input[str] description: A human-readable description for the role.
        :param pulumi.Input[list] permissions: The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
        :param pulumi.Input[str] project: The project that the service account will be created in.
               Defaults to the provider project configuration.
        :param pulumi.Input[str] role_id: The role id to use for this role.
        :param pulumi.Input[str] stage: The current launch stage of the role.
               Defaults to `GA`.
               List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
        :param pulumi.Input[str] title: A human-readable title for the role.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['deleted'] = deleted

        __props__['description'] = description

        if not permissions:
            raise TypeError('Missing required property permissions')
        __props__['permissions'] = permissions

        __props__['project'] = project

        if not role_id:
            raise TypeError('Missing required property role_id')
        __props__['role_id'] = role_id

        __props__['stage'] = stage

        if not title:
            raise TypeError('Missing required property title')
        __props__['title'] = title

        super(IAMCustomRole, __self__).__init__(
            'gcp:projects/iAMCustomRole:IAMCustomRole',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

