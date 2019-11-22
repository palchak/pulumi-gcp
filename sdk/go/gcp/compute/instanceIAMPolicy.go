// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_instance_iam_policy.html.markdown.
type InstanceIAMPolicy struct {
	s *pulumi.ResourceState
}

// NewInstanceIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceIAMPolicy(ctx *pulumi.Context,
	name string, args *InstanceIAMPolicyArgs, opts ...pulumi.ResourceOpt) (*InstanceIAMPolicy, error) {
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["instanceName"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["zone"] = nil
	} else {
		inputs["instanceName"] = args.InstanceName
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceIAMPolicy{s: s}, nil
}

// GetInstanceIAMPolicy gets an existing InstanceIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InstanceIAMPolicyState, opts ...pulumi.ResourceOpt) (*InstanceIAMPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["etag"] = state.Etag
		inputs["instanceName"] = state.InstanceName
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/instanceIAMPolicy:InstanceIAMPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InstanceIAMPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InstanceIAMPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InstanceIAMPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// (Computed) The etag of the IAM policy.
func (r *InstanceIAMPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// Used to find the parent resource to bind the IAM policy to
func (r *InstanceIAMPolicy) InstanceName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["instanceName"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *InstanceIAMPolicy) PolicyData() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *InstanceIAMPolicy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
// zone is specified, it is taken from the provider configuration.
func (r *InstanceIAMPolicy) Zone() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering InstanceIAMPolicy resources.
type InstanceIAMPolicyState struct {
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// Used to find the parent resource to bind the IAM policy to
	InstanceName interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone interface{}
}

// The set of arguments for constructing a InstanceIAMPolicy resource.
type InstanceIAMPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	InstanceName interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no zone is provided in the parent identifier and no
	// zone is specified, it is taken from the provider configuration.
	Zone interface{}
}
