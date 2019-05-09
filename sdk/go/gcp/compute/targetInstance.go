// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Represents a TargetInstance resource which defines an endpoint instance
// that terminates traffic of certain protocols. In particular, they are used
// in Protocol Forwarding, where forwarding rules can send packets to a
// non-NAT’ed target instance. Each target instance contains a single
// virtual machine instance that receives and handles traffic from the
// corresponding forwarding rules.
// 
// 
// To get more information about TargetInstance, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetInstances)
// * How-to Guides
//     * [Using Protocol Forwarding](https://cloud.google.com/compute/docs/protocol-forwarding)
type TargetInstance struct {
	s *pulumi.ResourceState
}

// NewTargetInstance registers a new resource with the given unique name, arguments, and options.
func NewTargetInstance(ctx *pulumi.Context,
	name string, args *TargetInstanceArgs, opts ...pulumi.ResourceOpt) (*TargetInstance, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["instance"] = nil
		inputs["name"] = nil
		inputs["natPolicy"] = nil
		inputs["project"] = nil
		inputs["zone"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["instance"] = args.Instance
		inputs["name"] = args.Name
		inputs["natPolicy"] = args.NatPolicy
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	inputs["creationTimestamp"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/targetInstance:TargetInstance", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetInstance{s: s}, nil
}

// GetTargetInstance gets an existing TargetInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetInstance(ctx *pulumi.Context,
	name string, id pulumi.ID, state *TargetInstanceState, opts ...pulumi.ResourceOpt) (*TargetInstance, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["instance"] = state.Instance
		inputs["name"] = state.Name
		inputs["natPolicy"] = state.NatPolicy
		inputs["project"] = state.Project
		inputs["selfLink"] = state.SelfLink
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/targetInstance:TargetInstance", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &TargetInstance{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *TargetInstance) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *TargetInstance) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *TargetInstance) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *TargetInstance) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *TargetInstance) Instance() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["instance"])
}

func (r *TargetInstance) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *TargetInstance) NatPolicy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["natPolicy"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *TargetInstance) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

// The URI of the created resource.
func (r *TargetInstance) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

func (r *TargetInstance) Zone() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering TargetInstance resources.
type TargetInstanceState struct {
	CreationTimestamp interface{}
	Description interface{}
	Instance interface{}
	Name interface{}
	NatPolicy interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The URI of the created resource.
	SelfLink interface{}
	Zone interface{}
}

// The set of arguments for constructing a TargetInstance resource.
type TargetInstanceArgs struct {
	Description interface{}
	Instance interface{}
	Name interface{}
	NatPolicy interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Zone interface{}
}
