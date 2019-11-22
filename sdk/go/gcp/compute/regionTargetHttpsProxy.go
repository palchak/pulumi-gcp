// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_region_target_https_proxy.html.markdown.
type RegionTargetHttpsProxy struct {
	s *pulumi.ResourceState
}

// NewRegionTargetHttpsProxy registers a new resource with the given unique name, arguments, and options.
func NewRegionTargetHttpsProxy(ctx *pulumi.Context,
	name string, args *RegionTargetHttpsProxyArgs, opts ...pulumi.ResourceOpt) (*RegionTargetHttpsProxy, error) {
	if args == nil || args.SslCertificates == nil {
		return nil, errors.New("missing required argument 'SslCertificates'")
	}
	if args == nil || args.UrlMap == nil {
		return nil, errors.New("missing required argument 'UrlMap'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["sslCertificates"] = nil
		inputs["urlMap"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["sslCertificates"] = args.SslCertificates
		inputs["urlMap"] = args.UrlMap
	}
	inputs["creationTimestamp"] = nil
	inputs["proxyId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionTargetHttpsProxy{s: s}, nil
}

// GetRegionTargetHttpsProxy gets an existing RegionTargetHttpsProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionTargetHttpsProxy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RegionTargetHttpsProxyState, opts ...pulumi.ResourceOpt) (*RegionTargetHttpsProxy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["proxyId"] = state.ProxyId
		inputs["region"] = state.Region
		inputs["selfLink"] = state.SelfLink
		inputs["sslCertificates"] = state.SslCertificates
		inputs["urlMap"] = state.UrlMap
	}
	s, err := ctx.ReadResource("gcp:compute/regionTargetHttpsProxy:RegionTargetHttpsProxy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RegionTargetHttpsProxy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RegionTargetHttpsProxy) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RegionTargetHttpsProxy) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Creation timestamp in RFC3339 text format.
func (r *RegionTargetHttpsProxy) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// An optional description of this resource.
func (r *RegionTargetHttpsProxy) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *RegionTargetHttpsProxy) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *RegionTargetHttpsProxy) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The unique identifier for the resource.
func (r *RegionTargetHttpsProxy) ProxyId() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["proxyId"])
}

// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
func (r *RegionTargetHttpsProxy) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// The URI of the created resource.
func (r *RegionTargetHttpsProxy) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
// Currently, exactly one SSL certificate must be specified.
func (r *RegionTargetHttpsProxy) SslCertificates() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["sslCertificates"])
}

// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
func (r *RegionTargetHttpsProxy) UrlMap() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["urlMap"])
}

// Input properties used for looking up and filtering RegionTargetHttpsProxy resources.
type RegionTargetHttpsProxyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// An optional description of this resource.
	Description interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The unique identifier for the resource.
	ProxyId interface{}
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates interface{}
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap interface{}
}

// The set of arguments for constructing a RegionTargetHttpsProxy resource.
type RegionTargetHttpsProxyArgs struct {
	// An optional description of this resource.
	Description interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// The Region in which the created target https proxy should reside. If it is not provided, the provider region is used.
	Region interface{}
	// A list of RegionSslCertificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, exactly one SSL certificate must be specified.
	SslCertificates interface{}
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the RegionBackendService.
	UrlMap interface{}
}
