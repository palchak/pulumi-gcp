// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_web_backend_service_iam_binding.html.markdown.
type WebBackendServiceIamBinding struct {
	s *pulumi.ResourceState
}

// NewWebBackendServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWebBackendServiceIamBinding(ctx *pulumi.Context,
	name string, args *WebBackendServiceIamBindingArgs, opts ...pulumi.ResourceOpt) (*WebBackendServiceIamBinding, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.WebBackendService == nil {
		return nil, errors.New("missing required argument 'WebBackendService'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["condition"] = nil
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
		inputs["webBackendService"] = nil
	} else {
		inputs["condition"] = args.Condition
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
		inputs["webBackendService"] = args.WebBackendService
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebBackendServiceIamBinding{s: s}, nil
}

// GetWebBackendServiceIamBinding gets an existing WebBackendServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebBackendServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WebBackendServiceIamBindingState, opts ...pulumi.ResourceOpt) (*WebBackendServiceIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["condition"] = state.Condition
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
		inputs["webBackendService"] = state.WebBackendService
	}
	s, err := ctx.ReadResource("gcp:iap/webBackendServiceIamBinding:WebBackendServiceIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &WebBackendServiceIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *WebBackendServiceIamBinding) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *WebBackendServiceIamBinding) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *WebBackendServiceIamBinding) Condition() pulumi.Output {
	return r.s.State["condition"]
}

// (Computed) The etag of the IAM policy.
func (r *WebBackendServiceIamBinding) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

func (r *WebBackendServiceIamBinding) Members() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["members"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *WebBackendServiceIamBinding) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *WebBackendServiceIamBinding) Role() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["role"])
}

// Used to find the parent resource to bind the IAM policy to
func (r *WebBackendServiceIamBinding) WebBackendService() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["webBackendService"])
}

// Input properties used for looking up and filtering WebBackendServiceIamBinding resources.
type WebBackendServiceIamBindingState struct {
	Condition interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService interface{}
}

// The set of arguments for constructing a WebBackendServiceIamBinding resource.
type WebBackendServiceIamBindingArgs struct {
	Condition interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `iap.WebBackendServiceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
	// Used to find the parent resource to bind the IAM policy to
	WebBackendService interface{}
}
