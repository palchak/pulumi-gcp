// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Contains the data that describes an Identity Aware Proxy owned client.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_client.html.markdown.
type Client struct {
	pulumi.CustomResourceState

	// Identifier of the brand to which this client is attached to. The format is
	// 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	Brand pulumi.StringOutput `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil || args.Brand == nil {
		return nil, errors.New("missing required argument 'Brand'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil {
		args = &ClientArgs{}
	}
	var resource Client
	err := ctx.RegisterResource("gcp:iap/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("gcp:iap/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	// Identifier of the brand to which this client is attached to. The format is
	// 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	Brand *string `pulumi:"brand"`
	// Output only. Unique identifier of the OAuth client.
	ClientId *string `pulumi:"clientId"`
	// Human-friendly name given to the OAuth client.
	DisplayName *string `pulumi:"displayName"`
	// Output only. Client secret of the OAuth client.
	Secret *string `pulumi:"secret"`
}

type ClientState struct {
	// Identifier of the brand to which this client is attached to. The format is
	// 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	Brand pulumi.StringPtrInput
	// Output only. Unique identifier of the OAuth client.
	ClientId pulumi.StringPtrInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringPtrInput
	// Output only. Client secret of the OAuth client.
	Secret pulumi.StringPtrInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	// Identifier of the brand to which this client is attached to. The format is
	// 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	Brand string `pulumi:"brand"`
	// Human-friendly name given to the OAuth client.
	DisplayName string `pulumi:"displayName"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	// Identifier of the brand to which this client is attached to. The format is
	// 'projects/{project_number}/brands/{brand_id}/identityAwareProxyClients/{client_id}'.
	Brand pulumi.StringInput
	// Human-friendly name given to the OAuth client.
	DisplayName pulumi.StringInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}