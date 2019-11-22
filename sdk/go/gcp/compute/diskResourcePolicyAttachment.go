// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_disk_resource_policy_attachment.html.markdown.
type DiskResourcePolicyAttachment struct {
	s *pulumi.ResourceState
}

// NewDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *DiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOpt) (*DiskResourcePolicyAttachment, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["disk"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["zone"] = nil
	} else {
		inputs["disk"] = args.Disk
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["zone"] = args.Zone
	}
	s, err := ctx.RegisterResource("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DiskResourcePolicyAttachment{s: s}, nil
}

// GetDiskResourcePolicyAttachment gets an existing DiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOpt) (*DiskResourcePolicyAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["disk"] = state.Disk
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["zone"] = state.Zone
	}
	s, err := ctx.ReadResource("gcp:compute/diskResourcePolicyAttachment:DiskResourcePolicyAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DiskResourcePolicyAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DiskResourcePolicyAttachment) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DiskResourcePolicyAttachment) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the disk in which the resource policies are attached to.
func (r *DiskResourcePolicyAttachment) Disk() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["disk"])
}

// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
func (r *DiskResourcePolicyAttachment) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *DiskResourcePolicyAttachment) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A reference to the zone where the disk resides.
func (r *DiskResourcePolicyAttachment) Zone() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["zone"])
}

// Input properties used for looking up and filtering DiskResourcePolicyAttachment resources.
type DiskResourcePolicyAttachmentState struct {
	// The name of the disk in which the resource policies are attached to.
	Disk interface{}
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the zone where the disk resides.
	Zone interface{}
}

// The set of arguments for constructing a DiskResourcePolicyAttachment resource.
type DiskResourcePolicyAttachmentArgs struct {
	// The name of the disk in which the resource policies are attached to.
	Disk interface{}
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the zone where the disk resides.
	Zone interface{}
}
