# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClientConfigResult:
    """
    A collection of values returned by getClientConfig.
    """
    def __init__(__self__, access_token=None, project=None, region=None, id=None):
        if access_token and not isinstance(access_token, str):
            raise TypeError('Expected argument access_token to be a str')
        __self__.access_token = access_token
        """
        The OAuth2 access token used by the client to authenticate against the Google Cloud API.
        """
        if project and not isinstance(project, str):
            raise TypeError('Expected argument project to be a str')
        __self__.project = project
        """
        The ID of the project to apply any resources to.
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        The region to operate under.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_client_config(opts=None):
    """
    Use this data source to access the configuration of the Google Cloud provider.
    """
    __args__ = dict()

    __ret__ = await pulumi.runtime.invoke('gcp:organizations/getClientConfig:getClientConfig', __args__, opts=opts)

    return GetClientConfigResult(
        access_token=__ret__.get('accessToken'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        id=__ret__.get('id'))
