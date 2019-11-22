// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_external_vpn_gateway.html.markdown.
type ExternalVpnGateway struct {
	s *pulumi.ResourceState
}

// NewExternalVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewExternalVpnGateway(ctx *pulumi.Context,
	name string, args *ExternalVpnGatewayArgs, opts ...pulumi.ResourceOpt) (*ExternalVpnGateway, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["interfaces"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["redundancyType"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["interfaces"] = args.Interfaces
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["redundancyType"] = args.RedundancyType
	}
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalVpnGateway{s: s}, nil
}

// GetExternalVpnGateway gets an existing ExternalVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ExternalVpnGatewayState, opts ...pulumi.ResourceOpt) (*ExternalVpnGateway, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["interfaces"] = state.Interfaces
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["redundancyType"] = state.RedundancyType
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:compute/externalVpnGateway:ExternalVpnGateway", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ExternalVpnGateway{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ExternalVpnGateway) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ExternalVpnGateway) ID() pulumi.IDOutput {
	return r.s.ID()
}

// An optional description of this resource.
func (r *ExternalVpnGateway) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// A list of interfaces on this external VPN gateway.
func (r *ExternalVpnGateway) Interfaces() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["interfaces"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *ExternalVpnGateway) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *ExternalVpnGateway) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Indicates the redundancy type of this external VPN gateway
func (r *ExternalVpnGateway) RedundancyType() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["redundancyType"])
}

func (r *ExternalVpnGateway) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering ExternalVpnGateway resources.
type ExternalVpnGatewayState struct {
	// An optional description of this resource.
	Description interface{}
	// A list of interfaces on this external VPN gateway.
	Interfaces interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType interface{}
	SelfLink interface{}
}

// The set of arguments for constructing a ExternalVpnGateway resource.
type ExternalVpnGatewayArgs struct {
	// An optional description of this resource.
	Description interface{}
	// A list of interfaces on this external VPN gateway.
	Interfaces interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Indicates the redundancy type of this external VPN gateway
	RedundancyType interface{}
}
