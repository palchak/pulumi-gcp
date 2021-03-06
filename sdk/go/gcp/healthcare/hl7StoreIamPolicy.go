// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
//
// * `healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
// * `healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
// * `healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
//
// > **Note:** `healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `healthcare.Hl7StoreIamBinding` and `healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
type Hl7StoreIamPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringOutput `pulumi:"hl7V2StoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewHl7StoreIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewHl7StoreIamPolicy(ctx *pulumi.Context,
	name string, args *Hl7StoreIamPolicyArgs, opts ...pulumi.ResourceOption) (*Hl7StoreIamPolicy, error) {
	if args == nil || args.Hl7V2StoreId == nil {
		return nil, errors.New("missing required argument 'Hl7V2StoreId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &Hl7StoreIamPolicyArgs{}
	}
	var resource Hl7StoreIamPolicy
	err := ctx.RegisterResource("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHl7StoreIamPolicy gets an existing Hl7StoreIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7StoreIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Hl7StoreIamPolicyState, opts ...pulumi.ResourceOption) (*Hl7StoreIamPolicy, error) {
	var resource Hl7StoreIamPolicy
	err := ctx.ReadResource("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hl7StoreIamPolicy resources.
type hl7StoreIamPolicyState struct {
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId *string `pulumi:"hl7V2StoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type Hl7StoreIamPolicyState struct {
	// (Computed) The etag of the HL7v2 store's IAM policy.
	Etag pulumi.StringPtrInput
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (Hl7StoreIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamPolicyState)(nil)).Elem()
}

type hl7StoreIamPolicyArgs struct {
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId string `pulumi:"hl7V2StoreId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a Hl7StoreIamPolicy resource.
type Hl7StoreIamPolicyArgs struct {
	// The HL7v2 store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
	// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	Hl7V2StoreId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (Hl7StoreIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7StoreIamPolicyArgs)(nil)).Elem()
}
