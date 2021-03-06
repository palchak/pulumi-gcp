// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Secret Manager Secret. Each of these resources serves a different use case:
//
// * `secretmanager.SecretIamPolicy`: Authoritative. Sets the IAM policy for the secret and replaces any existing policy already attached.
// * `secretmanager.SecretIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the secret are preserved.
// * `secretmanager.SecretIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the secret are preserved.
//
// > **Note:** `secretmanager.SecretIamPolicy` **cannot** be used in conjunction with `secretmanager.SecretIamBinding` and `secretmanager.SecretIamMember` or they will fight over what your policy should be.
//
// > **Note:** `secretmanager.SecretIamBinding` resources **can be** used in conjunction with `secretmanager.SecretIamMember` resources **only if** they do not grant privilege to the same role.
type SecretIamMember struct {
	pulumi.CustomResourceState

	Condition SecretIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringOutput `pulumi:"role"`
	SecretId pulumi.StringOutput `pulumi:"secretId"`
}

// NewSecretIamMember registers a new resource with the given unique name, arguments, and options.
func NewSecretIamMember(ctx *pulumi.Context,
	name string, args *SecretIamMemberArgs, opts ...pulumi.ResourceOption) (*SecretIamMember, error) {
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	if args == nil {
		args = &SecretIamMemberArgs{}
	}
	var resource SecretIamMember
	err := ctx.RegisterResource("gcp:secretmanager/secretIamMember:SecretIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretIamMember gets an existing SecretIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretIamMemberState, opts ...pulumi.ResourceOption) (*SecretIamMember, error) {
	var resource SecretIamMember
	err := ctx.ReadResource("gcp:secretmanager/secretIamMember:SecretIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretIamMember resources.
type secretIamMemberState struct {
	Condition *SecretIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     *string `pulumi:"role"`
	SecretId *string `pulumi:"secretId"`
}

type SecretIamMemberState struct {
	Condition SecretIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringPtrInput
	SecretId pulumi.StringPtrInput
}

func (SecretIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamMemberState)(nil)).Elem()
}

type secretIamMemberArgs struct {
	Condition *SecretIamMemberCondition `pulumi:"condition"`
	Member    string                    `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     string `pulumi:"role"`
	SecretId string `pulumi:"secretId"`
}

// The set of arguments for constructing a SecretIamMember resource.
type SecretIamMemberArgs struct {
	Condition SecretIamMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `secretmanager.SecretIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role     pulumi.StringInput
	SecretId pulumi.StringInput
}

func (SecretIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretIamMemberArgs)(nil)).Elem()
}
