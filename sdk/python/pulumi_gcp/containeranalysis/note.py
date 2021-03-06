# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Note(pulumi.CustomResource):
    attestation_authority: pulumi.Output[dict]
    """
    Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
    AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for the
    attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the
    ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find all
    attached Attestation Occurrences, even if they don't all live in the same project.

      * `hint` (`dict`)
        * `humanReadableName` (`str`)
    """
    name: pulumi.Output[str]
    """
    The name of the note.
    """
    project: pulumi.Output[str]
    """
    The ID of the project in which the resource belongs.
    If it is not provided, the provider project is used.
    """
    def __init__(__self__, resource_name, opts=None, attestation_authority=None, name=None, project=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a detailed description of a Note.


        To get more information about Note, see:

        * [API documentation](https://cloud.google.com/container-analysis/api/reference/rest/)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/container-analysis/)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attestation_authority: Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
               AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for the
               attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the
               ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find all
               attached Attestation Occurrences, even if they don't all live in the same project.
        :param pulumi.Input[str] name: The name of the note.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **attestation_authority** object supports the following:

          * `hint` (`pulumi.Input[dict]`)
            * `humanReadableName` (`pulumi.Input[str]`)
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

            if attestation_authority is None:
                raise TypeError("Missing required property 'attestation_authority'")
            __props__['attestation_authority'] = attestation_authority
            __props__['name'] = name
            __props__['project'] = project
        super(Note, __self__).__init__(
            'gcp:containeranalysis/note:Note',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attestation_authority=None, name=None, project=None):
        """
        Get an existing Note resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attestation_authority: Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one
               AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for the
               attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the
               ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find all
               attached Attestation Occurrences, even if they don't all live in the same project.
        :param pulumi.Input[str] name: The name of the note.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.

        The **attestation_authority** object supports the following:

          * `hint` (`pulumi.Input[dict]`)
            * `humanReadableName` (`pulumi.Input[str]`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attestation_authority"] = attestation_authority
        __props__["name"] = name
        __props__["project"] = project
        return Note(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

