// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package runtimeconfig

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/runtimeconfig_config_iam_binding.html.markdown.
type ConfigIamBinding struct {
	s *pulumi.ResourceState
}

// NewConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewConfigIamBinding(ctx *pulumi.Context,
	name string, args *ConfigIamBindingArgs, opts ...pulumi.ResourceOpt) (*ConfigIamBinding, error) {
	if args == nil || args.Config == nil {
		return nil, errors.New("missing required argument 'Config'")
	}
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["config"] = nil
		inputs["members"] = nil
		inputs["project"] = nil
		inputs["role"] = nil
	} else {
		inputs["config"] = args.Config
		inputs["members"] = args.Members
		inputs["project"] = args.Project
		inputs["role"] = args.Role
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:runtimeconfig/configIamBinding:ConfigIamBinding", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamBinding{s: s}, nil
}

// GetConfigIamBinding gets an existing ConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ConfigIamBindingState, opts ...pulumi.ResourceOpt) (*ConfigIamBinding, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["config"] = state.Config
		inputs["etag"] = state.Etag
		inputs["members"] = state.Members
		inputs["project"] = state.Project
		inputs["role"] = state.Role
	}
	s, err := ctx.ReadResource("gcp:runtimeconfig/configIamBinding:ConfigIamBinding", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ConfigIamBinding{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ConfigIamBinding) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ConfigIamBinding) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Used to find the parent resource to bind the IAM policy to
func (r *ConfigIamBinding) Config() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["config"])
}

// (Computed) The etag of the IAM policy.
func (r *ConfigIamBinding) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

func (r *ConfigIamBinding) Members() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["members"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
func (r *ConfigIamBinding) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The role that should be applied. Only one
// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
// `[projects|organizations]/{parent-name}/roles/{role-name}`.
func (r *ConfigIamBinding) Role() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["role"])
}

// Input properties used for looking up and filtering ConfigIamBinding resources.
type ConfigIamBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	// (Computed) The etag of the IAM policy.
	Etag interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}

// The set of arguments for constructing a ConfigIamBinding resource.
type ConfigIamBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Config interface{}
	Members interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project interface{}
	// The role that should be applied. Only one
	// `runtimeconfig.ConfigIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role interface{}
}