// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// A zone is a subtree of the DNS namespace under one administrative
// responsibility. A ManagedZone is a resource that represents a DNS zone
// hosted by the Cloud DNS service.
// 
// 
// To get more information about ManagedZone, see:
// 
// * [API documentation](https://cloud.google.com/dns/api/v1/managedZones)
// * How-to Guides
//     * [Managing Zones](https://cloud.google.com/dns/zones/)
// 
// <div class = "oics-button" style="float: right; margin: 0 0 -15px">
//   <a href="https://console.cloud.google.com/cloudshell/open?cloudshell_git_repo=https%3A%2F%2Fgithub.com%2Fterraform-google-modules%2Fdocs-examples.git&cloudshell_working_dir=dns_managed_zone_basic&cloudshell_image=gcr.io%2Fgraphite-cloud-shell-images%2Fterraform%3Alatest&open_in_editor=main.tf&cloudshell_print=.%2Fmotd&cloudshell_tutorial=.%2Ftutorial.md" target="_blank">
//     <img alt="Open in Cloud Shell" src="//gstatic.com/cloudssh/images/open-btn.svg" style="max-height: 44px; margin: 32px auto; max-width: 100%;">
//   </a>
// </div>
type ManagedZone struct {
	s *pulumi.ResourceState
}

// NewManagedZone registers a new resource with the given unique name, arguments, and options.
func NewManagedZone(ctx *pulumi.Context,
	name string, args *ManagedZoneArgs, opts ...pulumi.ResourceOpt) (*ManagedZone, error) {
	if args == nil || args.DnsName == nil {
		return nil, errors.New("missing required argument 'DnsName'")
	}
	inputs := make(map[string]interface{})
	inputs["description"] = "Managed by Pulumi"
	if args == nil {
		inputs["dnsName"] = nil
		inputs["forwardingConfig"] = nil
		inputs["labels"] = nil
		inputs["name"] = nil
		inputs["privateVisibilityConfig"] = nil
		inputs["project"] = nil
		inputs["visibility"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["dnsName"] = args.DnsName
		inputs["forwardingConfig"] = args.ForwardingConfig
		inputs["labels"] = args.Labels
		inputs["name"] = args.Name
		inputs["privateVisibilityConfig"] = args.PrivateVisibilityConfig
		inputs["project"] = args.Project
		inputs["visibility"] = args.Visibility
	}
	inputs["nameServers"] = nil
	s, err := ctx.RegisterResource("gcp:dns/managedZone:ManagedZone", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagedZone{s: s}, nil
}

// GetManagedZone gets an existing ManagedZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedZone(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ManagedZoneState, opts ...pulumi.ResourceOpt) (*ManagedZone, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["description"] = state.Description
		inputs["dnsName"] = state.DnsName
		inputs["forwardingConfig"] = state.ForwardingConfig
		inputs["labels"] = state.Labels
		inputs["name"] = state.Name
		inputs["nameServers"] = state.NameServers
		inputs["privateVisibilityConfig"] = state.PrivateVisibilityConfig
		inputs["project"] = state.Project
		inputs["visibility"] = state.Visibility
	}
	s, err := ctx.ReadResource("gcp:dns/managedZone:ManagedZone", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ManagedZone{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ManagedZone) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ManagedZone) ID() *pulumi.IDOutput {
	return r.s.ID()
}

func (r *ManagedZone) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

func (r *ManagedZone) DnsName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dnsName"])
}

func (r *ManagedZone) ForwardingConfig() *pulumi.Output {
	return r.s.State["forwardingConfig"]
}

func (r *ManagedZone) Labels() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["labels"])
}

func (r *ManagedZone) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *ManagedZone) NameServers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nameServers"])
}

func (r *ManagedZone) PrivateVisibilityConfig() *pulumi.Output {
	return r.s.State["privateVisibilityConfig"]
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *ManagedZone) Project() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["project"])
}

func (r *ManagedZone) Visibility() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["visibility"])
}

// Input properties used for looking up and filtering ManagedZone resources.
type ManagedZoneState struct {
	Description interface{}
	DnsName interface{}
	ForwardingConfig interface{}
	Labels interface{}
	Name interface{}
	NameServers interface{}
	PrivateVisibilityConfig interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Visibility interface{}
}

// The set of arguments for constructing a ManagedZone resource.
type ManagedZoneArgs struct {
	Description interface{}
	DnsName interface{}
	ForwardingConfig interface{}
	Labels interface{}
	Name interface{}
	PrivateVisibilityConfig interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	Visibility interface{}
}
