// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package serviceAccount

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Allows management of a [Google Cloud Platform service account](https://cloud.google.com/compute/docs/access/service-accounts)
// 
// > Creation of service accounts is eventually consistent, and that can lead to
// errors when you try to apply ACLs to service accounts immediately after
// creation.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/service_account.html.markdown.
type Account struct {
	s *pulumi.ResourceState
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOpt) (*Account, error) {
	if args == nil || args.AccountId == nil {
		return nil, errors.New("missing required argument 'AccountId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accountId"] = nil
		inputs["description"] = nil
		inputs["displayName"] = nil
		inputs["project"] = nil
	} else {
		inputs["accountId"] = args.AccountId
		inputs["description"] = args.Description
		inputs["displayName"] = args.DisplayName
		inputs["project"] = args.Project
	}
	inputs["email"] = nil
	inputs["name"] = nil
	inputs["uniqueId"] = nil
	s, err := ctx.RegisterResource("gcp:serviceAccount/account:Account", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Account{s: s}, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AccountState, opts ...pulumi.ResourceOpt) (*Account, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accountId"] = state.AccountId
		inputs["description"] = state.Description
		inputs["displayName"] = state.DisplayName
		inputs["email"] = state.Email
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["uniqueId"] = state.UniqueId
	}
	s, err := ctx.ReadResource("gcp:serviceAccount/account:Account", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Account{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Account) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Account) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The account id that is used to generate the service
// account email address and a stable unique id. It is unique within a project,
// must be 6-30 characters long, and match the regular expression `a-z`
// to comply with RFC1035. Changing this forces a new service account to be created.
func (r *Account) AccountId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["accountId"])
}

// A text description of the service account.
func (r *Account) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// The display name for the service account.
// Can be updated without creating a new resource.
func (r *Account) DisplayName() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["displayName"])
}

// The e-mail address of the service account. This value
// should be referenced from any `organizations.getIAMPolicy` data sources
// that would grant the service account privileges.
func (r *Account) Email() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["email"])
}

// The fully-qualified name of the service account.
func (r *Account) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project that the service account will be created in.
// Defaults to the provider project configuration.
func (r *Account) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The unique id of the service account.
func (r *Account) UniqueId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["uniqueId"])
}

// Input properties used for looking up and filtering Account resources.
type AccountState struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId interface{}
	// A text description of the service account.
	Description interface{}
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName interface{}
	// The e-mail address of the service account. This value
	// should be referenced from any `organizations.getIAMPolicy` data sources
	// that would grant the service account privileges.
	Email interface{}
	// The fully-qualified name of the service account.
	Name interface{}
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project interface{}
	// The unique id of the service account.
	UniqueId interface{}
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The account id that is used to generate the service
	// account email address and a stable unique id. It is unique within a project,
	// must be 6-30 characters long, and match the regular expression `a-z`
	// to comply with RFC1035. Changing this forces a new service account to be created.
	AccountId interface{}
	// A text description of the service account.
	Description interface{}
	// The display name for the service account.
	// Can be updated without creating a new resource.
	DisplayName interface{}
	// The ID of the project that the service account will be created in.
	// Defaults to the provider project configuration.
	Project interface{}
}
