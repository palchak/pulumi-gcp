// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store.html.markdown.
type FhirStore struct {
	s *pulumi.ResourceState
}

// NewFhirStore registers a new resource with the given unique name, arguments, and options.
func NewFhirStore(ctx *pulumi.Context,
	name string, args *FhirStoreArgs, opts ...pulumi.ResourceOpt) (*FhirStore, error) {
	if args == nil || args.Dataset == nil {
		return nil, errors.New("missing required argument 'Dataset'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dataset"] = nil
		inputs["disableReferentialIntegrity"] = nil
		inputs["disableResourceVersioning"] = nil
		inputs["enableHistoryImport"] = nil
		inputs["enableUpdateCreate"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["notificationConfig"] = nil
	} else {
		inputs["dataset"] = args.Dataset
		inputs["disableReferentialIntegrity"] = args.DisableReferentialIntegrity
		inputs["disableResourceVersioning"] = args.DisableResourceVersioning
		inputs["enableHistoryImport"] = args.EnableHistoryImport
		inputs["enableUpdateCreate"] = args.EnableUpdateCreate
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["notificationConfig"] = args.NotificationConfig
	}
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/fhirStore:FhirStore", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FhirStore{s: s}, nil
}

// GetFhirStore gets an existing FhirStore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStore(ctx *pulumi.Context,
	name string, id pulumi.ID, state *FhirStoreState, opts ...pulumi.ResourceOpt) (*FhirStore, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["dataset"] = state.Dataset
		inputs["disableReferentialIntegrity"] = state.DisableReferentialIntegrity
		inputs["disableResourceVersioning"] = state.DisableResourceVersioning
		inputs["enableHistoryImport"] = state.EnableHistoryImport
		inputs["enableUpdateCreate"] = state.EnableUpdateCreate
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["notificationConfig"] = state.NotificationConfig
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:healthcare/fhirStore:FhirStore", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &FhirStore{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *FhirStore) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *FhirStore) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Identifies the dataset addressed by this request. Must be in the format
// 'projects/{project}/locations/{location}/datasets/{dataset}'
func (r *FhirStore) Dataset() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["dataset"])
}

// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
// default value is false, meaning that the API will enforce referential integrity and fail the requests that will result
// in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential integrity check.
// Consequently, operations that rely on references, such as Patient.get$everything, will not return all the results if
// broken references exist. ** Changing this property may recreate the FHIR store (removing all data) **
func (r *FhirStore) DisableReferentialIntegrity() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disableReferentialIntegrity"])
}

// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
// store. If set to false, which is the default behavior, all write operations will cause historical versions to be
// recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set
// to true, no historical versions will be kept. The server will send back errors for attempts to read the historical
// versions. ** Changing this property may recreate the FHIR store (removing all data) **
func (r *FhirStore) DisableResourceVersioning() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["disableResourceVersioning"])
}

// Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into the
// FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past, which
// clients may not want to allow. If set to false, history bundles within an import will fail with an error. ** Changing
// this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually in the Google
// Cloud Healthcare admin console without recreating the FHIR store **
func (r *FhirStore) EnableHistoryImport() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableHistoryImport"])
}

// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation to
// create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and
// attempts to Update a non-existent resource will return errors. Please treat the audit logs with appropriate levels of
// care if client-specified resource IDs contain sensitive data such as patient identifiers, those IDs will be part of the
// FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
func (r *FhirStore) EnableUpdateCreate() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["enableUpdateCreate"])
}

// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a
// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
func (r *FhirStore) Labels() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["labels"])
}

// The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
func (r *FhirStore) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// A nested object resource
func (r *FhirStore) NotificationConfig() pulumi.Output {
	return r.s.State["notificationConfig"]
}

// The fully qualified name of this dataset
func (r *FhirStore) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering FhirStore resources.
type FhirStoreState struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset interface{}
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
	// default value is false, meaning that the API will enforce referential integrity and fail the requests that will result
	// in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential integrity check.
	// Consequently, operations that rely on references, such as Patient.get$everything, will not return all the results if
	// broken references exist. ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity interface{}
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
	// store. If set to false, which is the default behavior, all write operations will cause historical versions to be
	// recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set
	// to true, no historical versions will be kept. The server will send back errors for attempts to read the historical
	// versions. ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning interface{}
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into
	// the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past,
	// which clients may not want to allow. If set to false, history bundles within an import will fail with an error. **
	// Changing this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually in
	// the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport interface{}
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation to
	// create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation
	// and attempts to Update a non-existent resource will return errors. Please treat the audit logs with appropriate levels
	// of care if client-specified resource IDs contain sensitive data such as patient identifiers, those IDs will be part of
	// the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
	EnableUpdateCreate interface{}
	// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels interface{}
	// The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
	Name interface{}
	// A nested object resource
	NotificationConfig interface{}
	// The fully qualified name of this dataset
	SelfLink interface{}
}

// The set of arguments for constructing a FhirStore resource.
type FhirStoreArgs struct {
	// Identifies the dataset addressed by this request. Must be in the format
	// 'projects/{project}/locations/{location}/datasets/{dataset}'
	Dataset interface{}
	// Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The
	// default value is false, meaning that the API will enforce referential integrity and fail the requests that will result
	// in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential integrity check.
	// Consequently, operations that rely on references, such as Patient.get$everything, will not return all the results if
	// broken references exist. ** Changing this property may recreate the FHIR store (removing all data) **
	DisableReferentialIntegrity interface{}
	// Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR
	// store. If set to false, which is the default behavior, all write operations will cause historical versions to be
	// recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set
	// to true, no historical versions will be kept. The server will send back errors for attempts to read the historical
	// versions. ** Changing this property may recreate the FHIR store (removing all data) **
	DisableResourceVersioning interface{}
	// Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into
	// the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past,
	// which clients may not want to allow. If set to false, history bundles within an import will fail with an error. **
	// Changing this property may recreate the FHIR store (removing all data) ** ** This property can be changed manually in
	// the Google Cloud Healthcare admin console without recreating the FHIR store **
	EnableHistoryImport interface{}
	// Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation to
	// create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation
	// and attempts to Update a non-existent resource will return errors. Please treat the audit logs with appropriate levels
	// of care if client-specified resource IDs contain sensitive data such as patient identifiers, those IDs will be part of
	// the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.
	EnableUpdateCreate interface{}
	// User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a
	// UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression:
	// [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of
	// "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels interface{}
	// The resource name for the FhirStore. ** Changing this property may recreate the FHIR store (removing all data) **
	Name interface{}
	// A nested object resource
	NotificationConfig interface{}
}
