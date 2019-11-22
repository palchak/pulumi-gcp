// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfunctions

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/cloudfunctions_function_iam_policy.html.markdown.
type FunctionIamPolicy struct {
	s *pulumi.ResourceState
}

// NewFunctionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewFunctionIamPolicy(ctx *pulumi.Context,
	name string, args *FunctionIamPolicyArgs, opts ...pulumi.ResourceOpt) (*FunctionIamPolicy, error) {
	if args == nil || args.CloudFunction == nil {
		return nil, errors.New("missing required argument 'CloudFunction'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cloudFunction"] = nil
		inputs["policyData"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
	} else {
		inputs["cloudFunction"] = args.CloudFunction
		inputs["policyData"] = args.PolicyData
		inputs["project"] = args.Project
		inputs["region"] = args.Region
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FunctionIamPolicy{s: s}, nil
}

// GetFunctionIamPolicy gets an existing FunctionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FunctionIamPolicyState, opts ...pulumi.ResourceOpt) (*FunctionIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cloudFunction"] = state.CloudFunction
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
		inputs["project"] = state.Project
		inputs["region"] = state.Region
	}
	s, err := ctx.ReadResource("gcp:cloudfunctions/functionIamPolicy:FunctionIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FunctionIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FunctionIamPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FunctionIamPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Used to find the parent resource to bind the IAM policy to
func (r *FunctionIamPolicy) CloudFunction() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cloudFunction"])
}

// (Computed) The etag of the IAM policy.
func (r *FunctionIamPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `organizations.getIAMPolicy` data source.
func (r *FunctionIamPolicy) PolicyData() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["policyData"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *FunctionIamPolicy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
// region is specified, it is taken from the provider configuration.
func (r *FunctionIamPolicy) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// Input properties used for looking up and filtering FunctionIamPolicy resources.
type FunctionIamPolicyState struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region interface{}
}

// The set of arguments for constructing a FunctionIamPolicy resource.
type FunctionIamPolicyArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	CloudFunction interface{}
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The location of this cloud function. Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region interface{}
}
