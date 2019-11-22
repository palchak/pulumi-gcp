// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of Organization policies for a Google Project. For more information see
// [the official
// documentation](https://cloud.google.com/resource-manager/docs/organization-policy/overview) and
// [API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/setOrgPolicy).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/project_organization_policy.html.markdown.
type OrganizationPolicy struct {
	s *pulumi.ResourceState
}

// NewOrganizationPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrganizationPolicy(ctx *pulumi.Context,
	name string, args *OrganizationPolicyArgs, opts ...pulumi.ResourceOpt) (*OrganizationPolicy, error) {
	if args == nil || args.Constraint == nil {
		return nil, errors.New("missing required argument 'Constraint'")
	}
	if args == nil || args.Project == nil {
		return nil, errors.New("missing required argument 'Project'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["booleanPolicy"] = nil
		inputs["constraint"] = nil
		inputs["listPolicy"] = nil
		inputs["project"] = nil
		inputs["restorePolicy"] = nil
		inputs["version"] = nil
	} else {
		inputs["booleanPolicy"] = args.BooleanPolicy
		inputs["constraint"] = args.Constraint
		inputs["listPolicy"] = args.ListPolicy
		inputs["project"] = args.Project
		inputs["restorePolicy"] = args.RestorePolicy
		inputs["version"] = args.Version
	}
	inputs["etag"] = nil
	inputs["updateTime"] = nil
	s, err := ctx.RegisterResource("gcp:projects/organizationPolicy:OrganizationPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationPolicy{s: s}, nil
}

// GetOrganizationPolicy gets an existing OrganizationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrganizationPolicyState, opts ...pulumi.ResourceOpt) (*OrganizationPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["booleanPolicy"] = state.BooleanPolicy
		inputs["constraint"] = state.Constraint
		inputs["etag"] = state.Etag
		inputs["listPolicy"] = state.ListPolicy
		inputs["project"] = state.Project
		inputs["restorePolicy"] = state.RestorePolicy
		inputs["updateTime"] = state.UpdateTime
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("gcp:projects/organizationPolicy:OrganizationPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OrganizationPolicy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OrganizationPolicy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
func (r *OrganizationPolicy) BooleanPolicy() pulumi.Output {
	return r.s.State["booleanPolicy"]
}

// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
func (r *OrganizationPolicy) Constraint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["constraint"])
}

// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
func (r *OrganizationPolicy) Etag() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["etag"])
}

// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
func (r *OrganizationPolicy) ListPolicy() pulumi.Output {
	return r.s.State["listPolicy"]
}

// The project id of the project to set the policy for.
func (r *OrganizationPolicy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A restore policy is a constraint to restore the default policy. Structure is documented below.
func (r *OrganizationPolicy) RestorePolicy() pulumi.Output {
	return r.s.State["restorePolicy"]
}

// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
func (r *OrganizationPolicy) UpdateTime() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["updateTime"])
}

// Version of the Policy. Default version is 0.
func (r *OrganizationPolicy) Version() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering OrganizationPolicy resources.
type OrganizationPolicyState struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy interface{}
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint interface{}
	// (Computed) The etag of the organization policy. `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	Etag interface{}
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy interface{}
	// The project id of the project to set the policy for.
	Project interface{}
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy interface{}
	// (Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	UpdateTime interface{}
	// Version of the Policy. Default version is 0.
	Version interface{}
}

// The set of arguments for constructing a OrganizationPolicy resource.
type OrganizationPolicyArgs struct {
	// A boolean policy is a constraint that is either enforced or not. Structure is documented below.
	BooleanPolicy interface{}
	// The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).
	Constraint interface{}
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy interface{}
	// The project id of the project to set the policy for.
	Project interface{}
	// A restore policy is a constraint to restore the default policy. Structure is documented below.
	RestorePolicy interface{}
	// Version of the Policy. Default version is 0.
	Version interface{}
}
