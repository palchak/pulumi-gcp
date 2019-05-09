// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A VPC network is a virtual version of the traditional physical networks
// that exist within and between physical data centers. A VPC network
// provides connectivity for your Compute Engine virtual machine (VM)
// instances, Container Engine containers, App Engine Flex services, and
// other network-related resources.
// 
// Each GCP project contains one or more VPC networks. Each VPC network is a
// global entity spanning all GCP regions. This global VPC network allows VM
// instances and other resources to communicate with each other via internal,
// private IP addresses.
// 
// Each VPC network is subdivided into subnets, and each subnet is contained
// within a single region. You can have more than one subnet in a region for
// a given VPC network. Each subnet has a contiguous private RFC1918 IP
// space. You create instances, containers, and the like in these subnets.
// When you create an instance, you must create it in a subnet, and the
// instance draws its internal IP address from that subnet.
// 
// Virtual machine (VM) instances in a VPC network can communicate with
// instances in all other subnets of the same VPC network, regardless of
// region, using their RFC1918 private IP addresses. You can isolate portions
// of the network, even entire subnets, using firewall rules.
// 
// 
// To get more information about Subnetwork, see:
// 
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/subnetworks)
// * How-to Guides
//     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
//     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
type Subnetwork struct {
	s *pulumi.ResourceState
}

// NewSubnetwork registers a new resource with the given unique name, arguments, and options.
func NewSubnetwork(ctx *pulumi.Context,
	name string, args *SubnetworkArgs, opts ...pulumi.ResourceOpt) (*Subnetwork, error) {
	if args == nil || args.IpCidrRange == nil {
		return nil, errors.New("missing required argument 'IpCidrRange'")
	}
	if args == nil || args.Network == nil {
		return nil, errors.New("missing required argument 'Network'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["enableFlowLogs"] = nil
		inputs["ipCidrRange"] = nil
		inputs["logConfig"] = nil
		inputs["name"] = nil
		inputs["network"] = nil
		inputs["privateIpGoogleAccess"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["secondaryIpRanges"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["enableFlowLogs"] = args.EnableFlowLogs
		inputs["ipCidrRange"] = args.IpCidrRange
		inputs["logConfig"] = args.LogConfig
		inputs["name"] = args.Name
		inputs["network"] = args.Network
		inputs["privateIpGoogleAccess"] = args.PrivateIpGoogleAccess
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["secondaryIpRanges"] = args.SecondaryIpRanges
	}
	inputs["creationTimestamp"] = nil
	inputs["fingerprint"] = nil
	inputs["gatewayAddress"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:compute/subnetwork:Subnetwork", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnetwork{s: s}, nil
}

// GetSubnetwork gets an existing Subnetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetwork(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SubnetworkState, opts ...pulumi.ResourceOpt) (*Subnetwork, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["description"] = state.Description
		inputs["enableFlowLogs"] = state.EnableFlowLogs
		inputs["fingerprint"] = state.Fingerprint
		inputs["gatewayAddress"] = state.GatewayAddress
		inputs["ipCidrRange"] = state.IpCidrRange
		inputs["logConfig"] = state.LogConfig
		inputs["name"] = state.Name
		inputs["network"] = state.Network
		inputs["privateIpGoogleAccess"] = state.PrivateIpGoogleAccess
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["secondaryIpRanges"] = state.SecondaryIpRanges
		inputs["selfLink"] = state.SelfLink
	}
	s, err := ctx.ReadResource("gcp:compute/subnetwork:Subnetwork", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Subnetwork{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Subnetwork) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Subnetwork) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *Subnetwork) CreationTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

func (r *Subnetwork) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *Subnetwork) EnableFlowLogs() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enableFlowLogs"])
}

func (r *Subnetwork) Fingerprint() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fingerprint"])
}

func (r *Subnetwork) GatewayAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["gatewayAddress"])
}

func (r *Subnetwork) IpCidrRange() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipCidrRange"])
}

func (r *Subnetwork) LogConfig() *pulumi.Output {
	return r.s.State["logConfig"]
}

func (r *Subnetwork) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *Subnetwork) Network() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["network"])
}

func (r *Subnetwork) PrivateIpGoogleAccess() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["privateIpGoogleAccess"])
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *Subnetwork) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *Subnetwork) Region() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["region"])
}

func (r *Subnetwork) SecondaryIpRanges() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["secondaryIpRanges"])
}

// The URI of the created resource.
func (r *Subnetwork) SelfLink() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["selfLink"])
}

// Input properties used for looking up and filtering Subnetwork resources.
type SubnetworkState struct {
	CreationTimestamp interface{}
	Description interface{}
	EnableFlowLogs interface{}
	Fingerprint interface{}
	GatewayAddress interface{}
	IpCidrRange interface{}
	LogConfig interface{}
	Name interface{}
	Network interface{}
	PrivateIpGoogleAccess interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	SecondaryIpRanges interface{}
	// The URI of the created resource.
	SelfLink interface{}
}

// The set of arguments for constructing a Subnetwork resource.
type SubnetworkArgs struct {
	Description interface{}
	EnableFlowLogs interface{}
	IpCidrRange interface{}
	LogConfig interface{}
	Name interface{}
	Network interface{}
	PrivateIpGoogleAccess interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Region interface{}
	SecondaryIpRanges interface{}
}
