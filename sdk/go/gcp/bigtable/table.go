// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigtable

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a Google Cloud Bigtable table inside an instance. For more information see
// [the official documentation](https://cloud.google.com/bigtable/) and
// [API](https://cloud.google.com/bigtable/docs/go/reference).
type Table struct {
	pulumi.CustomResourceState

	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayOutput `pulumi:"columnFamilies"`
	// The name of the Bigtable instance.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The name of the table.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of predefined keys to split the table on.
	SplitKeys pulumi.StringArrayOutput `pulumi:"splitKeys"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil || args.InstanceName == nil {
		return nil, errors.New("missing required argument 'InstanceName'")
	}
	if args == nil {
		args = &TableArgs{}
	}
	var resource Table
	err := ctx.RegisterResource("gcp:bigtable/table:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("gcp:bigtable/table:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies []TableColumnFamily `pulumi:"columnFamilies"`
	// The name of the Bigtable instance.
	InstanceName *string `pulumi:"instanceName"`
	// The name of the table.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of predefined keys to split the table on.
	SplitKeys []string `pulumi:"splitKeys"`
}

type TableState struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringPtrInput
	// The name of the table.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of predefined keys to split the table on.
	SplitKeys pulumi.StringArrayInput
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies []TableColumnFamily `pulumi:"columnFamilies"`
	// The name of the Bigtable instance.
	InstanceName string `pulumi:"instanceName"`
	// The name of the table.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of predefined keys to split the table on.
	SplitKeys []string `pulumi:"splitKeys"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	// A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
	ColumnFamilies TableColumnFamilyArrayInput
	// The name of the Bigtable instance.
	InstanceName pulumi.StringInput
	// The name of the table.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of predefined keys to split the table on.
	SplitKeys pulumi.StringArrayInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}
