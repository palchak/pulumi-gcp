// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a group of dissimilar Compute Engine virtual machine instances.
// For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
type InstanceGroup struct {
	pulumi.CustomResourceState

	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of instances in the group. They should be given
	// as selfLink URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayOutput `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The number of instances in the group.
	Size pulumi.IntOutput `pulumi:"size"`
	// The zone that this instance group should be created in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroup registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroup(ctx *pulumi.Context,
	name string, args *InstanceGroupArgs, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	if args == nil {
		args = &InstanceGroupArgs{}
	}
	var resource InstanceGroup
	err := ctx.RegisterResource("gcp:compute/instanceGroup:InstanceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroup gets an existing InstanceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupState, opts ...pulumi.ResourceOption) (*InstanceGroup, error) {
	var resource InstanceGroup
	err := ctx.ReadResource("gcp:compute/instanceGroup:InstanceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroup resources.
type instanceGroupState struct {
	// An optional textual description of the instance
	// group.
	Description *string `pulumi:"description"`
	// List of instances in the group. They should be given
	// as selfLink URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances []string `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupNamedPortType `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The number of instances in the group.
	Size *int `pulumi:"size"`
	// The zone that this instance group should be created in.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupState struct {
	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrInput
	// List of instances in the group. They should be given
	// as selfLink URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayInput
	// The name which the port will be mapped to.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayInput
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The number of instances in the group.
	Size pulumi.IntPtrInput
	// The zone that this instance group should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupState)(nil)).Elem()
}

type instanceGroupArgs struct {
	// An optional textual description of the instance
	// group.
	Description *string `pulumi:"description"`
	// List of instances in the group. They should be given
	// as selfLink URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances []string `pulumi:"instances"`
	// The name which the port will be mapped to.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupNamedPortType `pulumi:"namedPorts"`
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone that this instance group should be created in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroup resource.
type InstanceGroupArgs struct {
	// An optional textual description of the instance
	// group.
	Description pulumi.StringPtrInput
	// List of instances in the group. They should be given
	// as selfLink URLs. When adding instances they must all be in the same
	// network and zone as the instance group.
	Instances pulumi.StringArrayInput
	// The name which the port will be mapped to.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupNamedPortTypeArrayInput
	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// `network` nor `instances` is specified, this field will be blank).
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone that this instance group should be created in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupArgs)(nil)).Elem()
}
