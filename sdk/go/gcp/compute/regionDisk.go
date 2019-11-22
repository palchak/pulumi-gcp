// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_disk.html.markdown.
type RegionDisk struct {
	s *pulumi.ResourceState
}

// NewRegionDisk registers a new resource with the given unique name, arguments, and options.
func NewRegionDisk(ctx *pulumi.Context,
	name string, args *RegionDiskArgs, opts ...pulumi.ResourceOpt) (*RegionDisk, error) {
	if args == nil || args.ReplicaZones == nil {
		return nil, errors.New("missing required argument 'ReplicaZones'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["diskEncryptionKey"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["physicalBlockSizeBytes"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["replicaZones"] = nil
		inputs["size"] = nil
		inputs["snapshot"] = nil
		inputs["sourceSnapshotEncryptionKey"] = nil
		inputs["type"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["diskEncryptionKey"] = args.DiskEncryptionKey
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["physicalBlockSizeBytes"] = args.PhysicalBlockSizeBytes
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["replicaZones"] = args.ReplicaZones
		inputs["size"] = args.Size
		inputs["snapshot"] = args.Snapshot
		inputs["sourceSnapshotEncryptionKey"] = args.SourceSnapshotEncryptionKey
		inputs["type"] = args.Type
	}
	inputs["creationTimestamp"] = nil
	inputs["labelFingerprint"] = nil
	inputs["lastAttachTimestamp"] = nil
	inputs["lastDetachTimestamp"] = nil
	inputs["selfLink"] = nil
	inputs["sourceSnapshotId"] = nil
	inputs["users"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionDisk:RegionDisk", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionDisk{s: s}, nil
}

// GetRegionDisk gets an existing RegionDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionDisk(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionDiskState, opts ...pulumi.ResourceOpt) (*RegionDisk, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["diskEncryptionKey"] = state.DiskEncryptionKey
		inputs["labelFingerprint"] = state.LabelFingerprint
		inputs["labels"] = state.Labels
		inputs["lastAttachTimestamp"] = state.LastAttachTimestamp
		inputs["lastDetachTimestamp"] = state.LastDetachTimestamp
		inputs["name"] = state.Name
		inputs["physicalBlockSizeBytes"] = state.PhysicalBlockSizeBytes
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["replicaZones"] = state.ReplicaZones
		inputs["selfLink"] = state.SelfLink
		inputs["size"] = state.Size
		inputs["snapshot"] = state.Snapshot
		inputs["sourceSnapshotEncryptionKey"] = state.SourceSnapshotEncryptionKey
		inputs["sourceSnapshotId"] = state.SourceSnapshotId
		inputs["type"] = state.Type
		inputs["users"] = state.Users
	}
	s, err := ctx.ReadResource("gcp:compute/regionDisk:RegionDisk", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionDisk{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionDisk) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionDisk) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Creation timestamp in RFC3339 text format.
func (r *RegionDisk) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource. Provide this property when you create the resource.
func (r *RegionDisk) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
// and you do not need to provide a key to use the disk later.
func (r *RegionDisk) DiskEncryptionKey() pulumi.Output {
	return r.s.State["diskEncryptionKey"]
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (r *RegionDisk) LabelFingerprint() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["labelFingerprint"])
}

// Labels to apply to this disk. A list of key->value pairs.
func (r *RegionDisk) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// Last attach timestamp in RFC3339 text format.
func (r *RegionDisk) LastAttachTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastAttachTimestamp"])
}

// Last detach timestamp in RFC3339 text format.
func (r *RegionDisk) LastDetachTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["lastDetachTimestamp"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *RegionDisk) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
// error message will list the supported values for the caller's project.
func (r *RegionDisk) PhysicalBlockSizeBytes() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["physicalBlockSizeBytes"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *RegionDisk) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// A reference to the region where the disk resides.
func (r *RegionDisk) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// URLs of the zones where the disk should be replicated to.
func (r *RegionDisk) ReplicaZones() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["replicaZones"])
}

// The URI of the created resource.
func (r *RegionDisk) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
// sourceImage or sourceSnapshot parameter, or specify it alone to create an empty persistent disk. If you specify this
// field along with sourceImage or sourceSnapshot, the value of sizeGb must not be less than the size of the sourceImage or
// the size of the snapshot.
func (r *RegionDisk) Size() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["size"])
}

// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For
// example, the following are valid values: *
// 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
func (r *RegionDisk) Snapshot() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["snapshot"])
}

// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
// customer-supplied encryption key.
func (r *RegionDisk) SourceSnapshotEncryptionKey() pulumi.Output {
	return r.s.State["sourceSnapshotEncryptionKey"]
}

// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create
// this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
func (r *RegionDisk) SourceSnapshotId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["sourceSnapshotId"])
}

// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.
func (r *RegionDisk) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
func (r *RegionDisk) Users() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["users"])
}

// Input properties used for looking up and filtering RegionDisk resources.
type RegionDiskState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey interface{}
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint interface{}
	// Labels to apply to this disk. A list of key->value pairs.
	Labels interface{}
	// Last attach timestamp in RFC3339 text format.
	LastAttachTimestamp interface{}
	// Last detach timestamp in RFC3339 text format.
	LastDetachTimestamp interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the region where the disk resides.
	Region interface{}
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// sourceImage or sourceSnapshot parameter, or specify it alone to create an empty persistent disk. If you specify this
	// field along with sourceImage or sourceSnapshot, the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size interface{}
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For
	// example, the following are valid values: *
	// 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot interface{}
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey interface{}
	// The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to
	// create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and
	// recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.
	SourceSnapshotId interface{}
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the
	// disk.
	Type interface{}
	// Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance
	Users interface{}
}

// The set of arguments for constructing a RegionDisk resource.
type RegionDiskArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description interface{}
	// Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you
	// must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk
	// to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not
	// provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key
	// and you do not need to provide a key to use the disk later.
	DiskEncryptionKey interface{}
	// Labels to apply to this disk. A list of key->value pairs.
	Labels interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently
	// supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the
	// error message will list the supported values for the caller's project.
	PhysicalBlockSizeBytes interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// A reference to the region where the disk resides.
	Region interface{}
	// URLs of the zones where the disk should be replicated to.
	ReplicaZones interface{}
	// Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the
	// sourceImage or sourceSnapshot parameter, or specify it alone to create an empty persistent disk. If you specify this
	// field along with sourceImage or sourceSnapshot, the value of sizeGb must not be less than the size of the sourceImage
	// or the size of the snapshot.
	Size interface{}
	// The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For
	// example, the following are valid values: *
	// 'https://www.googleapis.com/compute/v1/projects/project/global/snapshots/snapshot' *
	// 'projects/project/global/snapshots/snapshot' * 'global/snapshots/snapshot' * 'snapshot'
	Snapshot interface{}
	// The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a
	// customer-supplied encryption key.
	SourceSnapshotEncryptionKey interface{}
	// URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the
	// disk.
	Type interface{}
}
