// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A reservation is a mechanism used to guarantee BigQuery slots to users.
//
// To get more information about Reservation, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/reservations/rest/v1beta1/projects.locations.reservations/create)
// * How-to Guides
//     * [Introduction to Reservations](https://cloud.google.com/bigquery/docs/reservations-intro)
type Reservation struct {
	pulumi.CustomResourceState

	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrOutput `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntOutput `pulumi:"slotCapacity"`
}

// NewReservation registers a new resource with the given unique name, arguments, and options.
func NewReservation(ctx *pulumi.Context,
	name string, args *ReservationArgs, opts ...pulumi.ResourceOption) (*Reservation, error) {
	if args == nil || args.SlotCapacity == nil {
		return nil, errors.New("missing required argument 'SlotCapacity'")
	}
	if args == nil {
		args = &ReservationArgs{}
	}
	var resource Reservation
	err := ctx.RegisterResource("gcp:bigquery/reservation:Reservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReservation gets an existing Reservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReservationState, opts ...pulumi.ResourceOption) (*Reservation, error) {
	var resource Reservation
	err := ctx.ReadResource("gcp:bigquery/reservation:Reservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Reservation resources.
type reservationState struct {
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity *int `pulumi:"slotCapacity"`
}

type ReservationState struct {
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntPtrInput
}

func (ReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationState)(nil)).Elem()
}

type reservationArgs struct {
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots *bool `pulumi:"ignoreIdleSlots"`
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location *string `pulumi:"location"`
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity int `pulumi:"slotCapacity"`
}

// The set of arguments for constructing a Reservation resource.
type ReservationArgs struct {
	// If false, any query using this reservation will use idle slots from other reservations within the same admin project. If
	// true, a query using this reservation will execute with the slot capacity specified above at most.
	IgnoreIdleSlots pulumi.BoolPtrInput
	// The geographic location where the transfer config should reside. Examples: US, EU, asia-northeast1. The default value is
	// US.
	Location pulumi.StringPtrInput
	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit
	// of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	SlotCapacity pulumi.IntInput
}

func (ReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*reservationArgs)(nil)).Elem()
}
