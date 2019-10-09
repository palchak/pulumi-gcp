// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package binaryauthorization

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/binary_authorization_attestor_iam_binding.html.markdown.
type AttestorIamBinding struct {
	s *pulumi.ResourceState
}

// NewAttestorIamBinding registers a new resource with the given unique name, arguments, and options.
func NewAttestorIamBinding(ctx *pulumi.Context,
	name string, args *AttestorIamBindingArgs, opts ...pulumi.ResourceOpt) (*AttestorIamBinding, error) {
	if args == nil || args.Attestor == nil {
		return nil, errors.New("missing required argument 'Attestor'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["attestor"] = nil
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
	} else {
		inputs["attestor"] = args.Attestor
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AttestorIamBinding{s: s}, nil
}

// GetAttestorIamBinding gets an existing AttestorIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttestorIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AttestorIamBindingState, opts ...pulumi.ResourceOpt) (*AttestorIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["attestor"] = state.Attestor
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:binaryauthorization/attestorIamBinding:AttestorIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AttestorIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AttestorIamBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AttestorIamBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Used to find the parent resource to bind the IAM policy to
func (r *AttestorIamBinding) Attestor() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["attestor"])
}

// (Computed) The etag of the IAM policy.
func (r *AttestorIamBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *AttestorIamBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *AttestorIamBinding) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *AttestorIamBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering AttestorIamBinding resources.
type AttestorIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a AttestorIamBinding resource.
type AttestorIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Attestor interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}
