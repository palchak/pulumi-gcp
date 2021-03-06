// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Persistent disks can be attached to a compute instance using the `attachedDisk`
// section within the compute instance configuration.
// However there may be situations where managing the attached disks via the compute
// instance config isn't preferable or possible, such as attaching dynamic
// numbers of disks using the `count` variable.
//
//
// To get more information about attaching disks, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
// * How-to Guides
//     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
//
// **Note:** When using `compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attachedDisk"]` on the `compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
type AttachedDisk struct {
	pulumi.CustomResourceState

	DeviceName pulumi.StringOutput    `pulumi:"deviceName"`
	Disk       pulumi.StringOutput    `pulumi:"disk"`
	Instance   pulumi.StringOutput    `pulumi:"instance"`
	Mode       pulumi.StringPtrOutput `pulumi:"mode"`
	Project    pulumi.StringOutput    `pulumi:"project"`
	Zone       pulumi.StringOutput    `pulumi:"zone"`
}

// NewAttachedDisk registers a new resource with the given unique name, arguments, and options.
func NewAttachedDisk(ctx *pulumi.Context,
	name string, args *AttachedDiskArgs, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	if args == nil || args.Disk == nil {
		return nil, errors.New("missing required argument 'Disk'")
	}
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &AttachedDiskArgs{}
	}
	var resource AttachedDisk
	err := ctx.RegisterResource("gcp:compute/attachedDisk:AttachedDisk", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDisk gets an existing AttachedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDiskState, opts ...pulumi.ResourceOption) (*AttachedDisk, error) {
	var resource AttachedDisk
	err := ctx.ReadResource("gcp:compute/attachedDisk:AttachedDisk", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDisk resources.
type attachedDiskState struct {
	DeviceName *string `pulumi:"deviceName"`
	Disk       *string `pulumi:"disk"`
	Instance   *string `pulumi:"instance"`
	Mode       *string `pulumi:"mode"`
	Project    *string `pulumi:"project"`
	Zone       *string `pulumi:"zone"`
}

type AttachedDiskState struct {
	DeviceName pulumi.StringPtrInput
	Disk       pulumi.StringPtrInput
	Instance   pulumi.StringPtrInput
	Mode       pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	Zone       pulumi.StringPtrInput
}

func (AttachedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskState)(nil)).Elem()
}

type attachedDiskArgs struct {
	DeviceName *string `pulumi:"deviceName"`
	Disk       string  `pulumi:"disk"`
	Instance   string  `pulumi:"instance"`
	Mode       *string `pulumi:"mode"`
	Project    *string `pulumi:"project"`
	Zone       *string `pulumi:"zone"`
}

// The set of arguments for constructing a AttachedDisk resource.
type AttachedDiskArgs struct {
	DeviceName pulumi.StringPtrInput
	Disk       pulumi.StringInput
	Instance   pulumi.StringInput
	Mode       pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	Zone       pulumi.StringPtrInput
}

func (AttachedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDiskArgs)(nil)).Elem()
}
