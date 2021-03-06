# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetTransferProjectServieAccountResult:
    """
    A collection of values returned by getTransferProjectServieAccount.
    """
    def __init__(__self__, email=None, id=None, project=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        Email address of the default service account used by Storage Transfer Jobs running in this project
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
class AwaitableGetTransferProjectServieAccountResult(GetTransferProjectServieAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTransferProjectServieAccountResult(
            email=self.email,
            id=self.id,
            project=self.project)

def get_transfer_project_servie_account(project=None,opts=None):
    """
    Use this data source to retrieve Storage Transfer service account for this project




    :param str project: The project ID. If it is not provided, the provider project is used.
    """
    __args__ = dict()


    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getTransferProjectServieAccount:getTransferProjectServieAccount', __args__, opts=opts).value

    return AwaitableGetTransferProjectServieAccountResult(
        email=__ret__.get('email'),
        id=__ret__.get('id'),
        project=__ret__.get('project'))
