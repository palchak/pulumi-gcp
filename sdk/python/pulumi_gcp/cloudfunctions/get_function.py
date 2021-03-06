# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetFunctionResult:
    """
    A collection of values returned by getFunction.
    """
    def __init__(__self__, available_memory_mb=None, description=None, entry_point=None, environment_variables=None, event_triggers=None, https_trigger_url=None, id=None, ingress_settings=None, labels=None, max_instances=None, name=None, project=None, region=None, runtime=None, service_account_email=None, source_archive_bucket=None, source_archive_object=None, source_repositories=None, timeout=None, trigger_http=None, vpc_connector=None, vpc_connector_egress_settings=None):
        if available_memory_mb and not isinstance(available_memory_mb, float):
            raise TypeError("Expected argument 'available_memory_mb' to be a float")
        __self__.available_memory_mb = available_memory_mb
        """
        Available memory (in MB) to the function.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        Description of the function.
        """
        if entry_point and not isinstance(entry_point, str):
            raise TypeError("Expected argument 'entry_point' to be a str")
        __self__.entry_point = entry_point
        """
        Name of a JavaScript function that will be executed when the Google Cloud Function is triggered.
        """
        if environment_variables and not isinstance(environment_variables, dict):
            raise TypeError("Expected argument 'environment_variables' to be a dict")
        __self__.environment_variables = environment_variables
        if event_triggers and not isinstance(event_triggers, list):
            raise TypeError("Expected argument 'event_triggers' to be a list")
        __self__.event_triggers = event_triggers
        """
        A source that fires events in response to a condition in another service. Structure is documented below.
        """
        if https_trigger_url and not isinstance(https_trigger_url, str):
            raise TypeError("Expected argument 'https_trigger_url' to be a str")
        __self__.https_trigger_url = https_trigger_url
        """
        If function is triggered by HTTP, trigger URL is set here.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if ingress_settings and not isinstance(ingress_settings, str):
            raise TypeError("Expected argument 'ingress_settings' to be a str")
        __self__.ingress_settings = ingress_settings
        """
        Controls what traffic can reach the function.
        """
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        __self__.labels = labels
        """
        A map of labels applied to this function.
        """
        if max_instances and not isinstance(max_instances, float):
            raise TypeError("Expected argument 'max_instances' to be a float")
        __self__.max_instances = max_instances
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the Cloud Function.
        """
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        __self__.project = project
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if runtime and not isinstance(runtime, str):
            raise TypeError("Expected argument 'runtime' to be a str")
        __self__.runtime = runtime
        """
        The runtime in which the function is running.
        """
        if service_account_email and not isinstance(service_account_email, str):
            raise TypeError("Expected argument 'service_account_email' to be a str")
        __self__.service_account_email = service_account_email
        """
        The service account email to be assumed by the cloud function.
        """
        if source_archive_bucket and not isinstance(source_archive_bucket, str):
            raise TypeError("Expected argument 'source_archive_bucket' to be a str")
        __self__.source_archive_bucket = source_archive_bucket
        """
        The GCS bucket containing the zip archive which contains the function.
        """
        if source_archive_object and not isinstance(source_archive_object, str):
            raise TypeError("Expected argument 'source_archive_object' to be a str")
        __self__.source_archive_object = source_archive_object
        """
        The source archive object (file) in archive bucket.
        """
        if source_repositories and not isinstance(source_repositories, list):
            raise TypeError("Expected argument 'source_repositories' to be a list")
        __self__.source_repositories = source_repositories
        if timeout and not isinstance(timeout, float):
            raise TypeError("Expected argument 'timeout' to be a float")
        __self__.timeout = timeout
        """
        Function execution timeout (in seconds).
        """
        if trigger_http and not isinstance(trigger_http, bool):
            raise TypeError("Expected argument 'trigger_http' to be a bool")
        __self__.trigger_http = trigger_http
        """
        If function is triggered by HTTP, this boolean is set.
        """
        if vpc_connector and not isinstance(vpc_connector, str):
            raise TypeError("Expected argument 'vpc_connector' to be a str")
        __self__.vpc_connector = vpc_connector
        """
        The VPC Network Connector that this cloud function can connect to. 
        """
        if vpc_connector_egress_settings and not isinstance(vpc_connector_egress_settings, str):
            raise TypeError("Expected argument 'vpc_connector_egress_settings' to be a str")
        __self__.vpc_connector_egress_settings = vpc_connector_egress_settings
        """
        The egress settings for the connector, controlling what traffic is diverted through it.
        """
class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            available_memory_mb=self.available_memory_mb,
            description=self.description,
            entry_point=self.entry_point,
            environment_variables=self.environment_variables,
            event_triggers=self.event_triggers,
            https_trigger_url=self.https_trigger_url,
            id=self.id,
            ingress_settings=self.ingress_settings,
            labels=self.labels,
            max_instances=self.max_instances,
            name=self.name,
            project=self.project,
            region=self.region,
            runtime=self.runtime,
            service_account_email=self.service_account_email,
            source_archive_bucket=self.source_archive_bucket,
            source_archive_object=self.source_archive_object,
            source_repositories=self.source_repositories,
            timeout=self.timeout,
            trigger_http=self.trigger_http,
            vpc_connector=self.vpc_connector,
            vpc_connector_egress_settings=self.vpc_connector_egress_settings)

def get_function(name=None,project=None,region=None,opts=None):
    """
    Get information about a Google Cloud Function. For more information see
    the [official documentation](https://cloud.google.com/functions/docs/)
    and [API](https://cloud.google.com/functions/docs/apis).




    :param str name: The name of a Cloud Function.
    :param str project: The project in which the resource belongs. If it
           is not provided, the provider project is used.
    :param str region: The region in which the resource belongs. If it
           is not provided, the provider region is used.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:cloudfunctions/getFunction:getFunction', __args__, opts=opts).value

    return AwaitableGetFunctionResult(
        available_memory_mb=__ret__.get('availableMemoryMb'),
        description=__ret__.get('description'),
        entry_point=__ret__.get('entryPoint'),
        environment_variables=__ret__.get('environmentVariables'),
        event_triggers=__ret__.get('eventTriggers'),
        https_trigger_url=__ret__.get('httpsTriggerUrl'),
        id=__ret__.get('id'),
        ingress_settings=__ret__.get('ingressSettings'),
        labels=__ret__.get('labels'),
        max_instances=__ret__.get('maxInstances'),
        name=__ret__.get('name'),
        project=__ret__.get('project'),
        region=__ret__.get('region'),
        runtime=__ret__.get('runtime'),
        service_account_email=__ret__.get('serviceAccountEmail'),
        source_archive_bucket=__ret__.get('sourceArchiveBucket'),
        source_archive_object=__ret__.get('sourceArchiveObject'),
        source_repositories=__ret__.get('sourceRepositories'),
        timeout=__ret__.get('timeout'),
        trigger_http=__ret__.get('triggerHttp'),
        vpc_connector=__ret__.get('vpcConnector'),
        vpc_connector_egress_settings=__ret__.get('vpcConnectorEgressSettings'))
