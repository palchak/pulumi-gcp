// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package healthcare

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > **Warning:** These resources are in beta, and should be used with the terraform-provider-google-beta provider.
// See [Provider Versions](https://terraform.io/docs/providers/google/provider_versions.html) for more details on beta resources.
// 
// Three different resources help you manage your IAM policy for Healthcare dataset. Each of these resources serves a different use case:
// 
// * `google_healthcare_dataset_iam_policy`: Authoritative. Sets the IAM policy for the dataset and replaces any existing policy already attached.
// * `google_healthcare_dataset_iam_binding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the dataset are preserved.
// * `google_healthcare_dataset_iam_member`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the dataset are preserved.
// 
// > **Note:** `google_healthcare_dataset_iam_policy` **cannot** be used in conjunction with `google_healthcare_dataset_iam_binding` and `google_healthcare_dataset_iam_member` or they will fight over what your policy should be.
// 
// > **Note:** `google_healthcare_dataset_iam_binding` resources **can be** used in conjunction with `google_healthcare_dataset_iam_member` resources **only if** they do not grant privilege to the same role.
type DatasetIamPolicy struct {
	s *pulumi.ResourceState
}

// NewDatasetIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDatasetIamPolicy(ctx *pulumi.Context,
	name string, args *DatasetIamPolicyArgs, opts ...pulumi.ResourceOpt) (*DatasetIamPolicy, error) {
	if args == nil || args.DatasetId == nil {
		return nil, errors.New("missing required argument 'DatasetId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["datasetId"] = nil
		inputs["policyData"] = nil
	} else {
		inputs["datasetId"] = args.DatasetId
		inputs["policyData"] = args.PolicyData
	}
	inputs["etag"] = nil
	s, err := ctx.RegisterResource("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatasetIamPolicy{s: s}, nil
}

// GetDatasetIamPolicy gets an existing DatasetIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DatasetIamPolicyState, opts ...pulumi.ResourceOpt) (*DatasetIamPolicy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["datasetId"] = state.DatasetId
		inputs["etag"] = state.Etag
		inputs["policyData"] = state.PolicyData
	}
	s, err := ctx.ReadResource("gcp:healthcare/datasetIamPolicy:DatasetIamPolicy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DatasetIamPolicy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DatasetIamPolicy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DatasetIamPolicy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The dataset ID, in the form
// `{project_id}/{location_name}/{dataset_name}` or
// `{location_name}/{dataset_name}`. In the second form, the provider's
// project setting will be used as a fallback.
func (r *DatasetIamPolicy) DatasetId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["datasetId"])
}

// (Computed) The etag of the dataset's IAM policy.
func (r *DatasetIamPolicy) Etag() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["etag"])
}

// The policy data generated by
// a `google_iam_policy` data source.
func (r *DatasetIamPolicy) PolicyData() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyData"])
}

// Input properties used for looking up and filtering DatasetIamPolicy resources.
type DatasetIamPolicyState struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId interface{}
	// (Computed) The etag of the dataset's IAM policy.
	Etag interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}

// The set of arguments for constructing a DatasetIamPolicy resource.
type DatasetIamPolicyArgs struct {
	// The dataset ID, in the form
	// `{project_id}/{location_name}/{dataset_name}` or
	// `{location_name}/{dataset_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	DatasetId interface{}
	// The policy data generated by
	// a `google_iam_policy` data source.
	PolicyData interface{}
}