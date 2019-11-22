// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_interconnect_attachment.html.markdown.
type InterconnectAttachment struct {
	s *pulumi.ResourceState
}

// NewInterconnectAttachment registers a new resource with the given unique name, arguments, and options.
func NewInterconnectAttachment(ctx *pulumi.Context,
	name string, args *InterconnectAttachmentArgs, opts ...pulumi.ResourceOpt) (*InterconnectAttachment, error) {
	if args == nil || args.Router == nil {
		return nil, errors.New("missing required argument 'Router'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adminEnabled"] = nil
		inputs["bandwidth"] = nil
		inputs["candidateSubnets"] = nil
		inputs["description"] = nil
		inputs["edgeAvailabilityDomain"] = nil
		inputs["interconnect"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["region"] = nil
		inputs["router"] = nil
		inputs["type"] = nil
		inputs["vlanTag8021q"] = nil
	} else {
		inputs["adminEnabled"] = args.AdminEnabled
		inputs["bandwidth"] = args.Bandwidth
		inputs["candidateSubnets"] = args.CandidateSubnets
		inputs["description"] = args.Description
		inputs["edgeAvailabilityDomain"] = args.EdgeAvailabilityDomain
		inputs["interconnect"] = args.Interconnect
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["region"] = args.Region
		inputs["router"] = args.Router
		inputs["type"] = args.Type
		inputs["vlanTag8021q"] = args.VlanTag8021q
	}
	inputs["cloudRouterIpAddress"] = nil
	inputs["creationTimestamp"] = nil
	inputs["customerRouterIpAddress"] = nil
	inputs["googleReferenceId"] = nil
	inputs["pairingKey"] = nil
	inputs["partnerAsn"] = nil
	inputs["privateInterconnectInfo"] = nil
	inputs["selfLink"] = nil
	inputs["state"] = nil
	s, err := ctx.RegisterResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachment{s: s}, nil
}

// GetInterconnectAttachment gets an existing InterconnectAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterconnectAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *InterconnectAttachmentState, opts ...pulumi.ResourceOpt) (*InterconnectAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adminEnabled"] = state.AdminEnabled
		inputs["bandwidth"] = state.Bandwidth
		inputs["candidateSubnets"] = state.CandidateSubnets
		inputs["cloudRouterIpAddress"] = state.CloudRouterIpAddress
		inputs["creationTimestamp"] = state.CreationTimestamp
		inputs["customerRouterIpAddress"] = state.CustomerRouterIpAddress
		inputs["description"] = state.Description
		inputs["edgeAvailabilityDomain"] = state.EdgeAvailabilityDomain
		inputs["googleReferenceId"] = state.GoogleReferenceId
		inputs["interconnect"] = state.Interconnect
		inputs["name"] = state.Name
		inputs["pairingKey"] = state.PairingKey
		inputs["partnerAsn"] = state.PartnerAsn
		inputs["privateInterconnectInfo"] = state.PrivateInterconnectInfo
		inputs["project"] = state.Project
		inputs["region"] = state.Region
		inputs["router"] = state.Router
		inputs["selfLink"] = state.SelfLink
		inputs["state"] = state.State
		inputs["type"] = state.Type
		inputs["vlanTag8021q"] = state.VlanTag8021q
	}
	s, err := ctx.ReadResource("gcp:compute/interconnectAttachment:InterconnectAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &InterconnectAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *InterconnectAttachment) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *InterconnectAttachment) ID() pulumi.IDOutput {
	return r.s.ID()
}

// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the interconnect
// attachment
func (r *InterconnectAttachment) AdminEnabled() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["adminEnabled"])
}

// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the
// bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the
// bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, Defaults to BPS_10G
func (r *InterconnectAttachment) Bandwidth() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bandwidth"])
}

// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and
// must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate
// prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will
// randomly select an unused /29 from all of link-local space.
func (r *InterconnectAttachment) CandidateSubnets() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["candidateSubnets"])
}

// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
func (r *InterconnectAttachment) CloudRouterIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["cloudRouterIpAddress"])
}

// Creation timestamp in RFC3339 text format.
func (r *InterconnectAttachment) CreationTimestamp() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["creationTimestamp"])
}

// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
func (r *InterconnectAttachment) CustomerRouterIpAddress() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["customerRouterIpAddress"])
}

// An optional description of this resource.
func (r *InterconnectAttachment) Description() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["description"])
}

// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For improved
// reliability, customers should configure a pair of attachments with one per availability domain. The selected
// availability domain will be provided to the Partner via the pairing key so that the provisioned circuit will lie in the
// specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
func (r *InterconnectAttachment) EdgeAvailabilityDomain() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["edgeAvailabilityDomain"])
}

// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
// issues.
func (r *InterconnectAttachment) GoogleReferenceId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["googleReferenceId"])
}

// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if type is
// DEDICATED, must not be set if type is PARTNER.
func (r *InterconnectAttachment) Interconnect() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["interconnect"])
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (r *InterconnectAttachment) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
// initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
func (r *InterconnectAttachment) PairingKey() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["pairingKey"])
}

// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
// layer 3 Partner if they configured BGP on behalf of the customer.
func (r *InterconnectAttachment) PartnerAsn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["partnerAsn"])
}

// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
// to is of type DEDICATED.
func (r *InterconnectAttachment) PrivateInterconnectInfo() pulumi.Output {
	return r.s.State["privateInterconnectInfo"]
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (r *InterconnectAttachment) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// Region where the regional interconnect attachment resides.
func (r *InterconnectAttachment) Region() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["region"])
}

// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region
// within which the Cloud Router is configured.
func (r *InterconnectAttachment) Router() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["router"])
}

// The URI of the created resource.
func (r *InterconnectAttachment) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// [Output Only] The current state of this attachment's functionality.
func (r *InterconnectAttachment) State() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["state"])
}

// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
func (r *InterconnectAttachment) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be managed
// upstream.
func (r *InterconnectAttachment) VlanTag8021q() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["vlanTag8021q"])
}

// Input properties used for looking up and filtering InterconnectAttachment resources.
type InterconnectAttachmentState struct {
	// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the interconnect
	// attachment
	AdminEnabled interface{}
	// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the
	// bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the
	// bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, Defaults to BPS_10G
	Bandwidth interface{}
	// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
	// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and
	// must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate
	// prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will
	// randomly select an unused /29 from all of link-local space.
	CandidateSubnets interface{}
	// IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
	CloudRouterIpAddress interface{}
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp interface{}
	// IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
	CustomerRouterIpAddress interface{}
	// An optional description of this resource.
	Description interface{}
	// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For improved
	// reliability, customers should configure a pair of attachments with one per availability domain. The selected
	// availability domain will be provided to the Partner via the pairing key so that the provisioned circuit will lie in the
	// specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain interface{}
	// Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity
	// issues.
	GoogleReferenceId interface{}
	// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if type is
	// DEDICATED, must not be set if type is PARTNER.
	Interconnect interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// [Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to
	// initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
	PairingKey interface{}
	// [Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a
	// layer 3 Partner if they configured BGP on behalf of the customer.
	PartnerAsn interface{}
	// Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached
	// to is of type DEDICATED.
	PrivateInterconnectInfo interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Region where the regional interconnect attachment resides.
	Region interface{}
	// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
	// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region
	// within which the Cloud Router is configured.
	Router interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// [Output Only] The current state of this attachment's functionality.
	State interface{}
	// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
	Type interface{}
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be managed
	// upstream.
	VlanTag8021q interface{}
}

// The set of arguments for constructing a InterconnectAttachment resource.
type InterconnectAttachmentArgs struct {
	// Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the interconnect
	// attachment
	AdminEnabled interface{}
	// Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the
	// bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the
	// bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, Defaults to BPS_10G
	Bandwidth interface{}
	// Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and
	// customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and
	// must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate
	// prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will
	// randomly select an unused /29 from all of link-local space.
	CandidateSubnets interface{}
	// An optional description of this resource.
	Description interface{}
	// Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For improved
	// reliability, customers should configure a pair of attachments with one per availability domain. The selected
	// availability domain will be provided to the Partner via the pairing key so that the provisioned circuit will lie in the
	// specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
	EdgeAvailabilityDomain interface{}
	// URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if type is
	// DEDICATED, must not be set if type is PARTNER.
	Interconnect interface{}
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression
	// '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character must be a lowercase letter, and all following characters
	// must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name interface{}
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project interface{}
	// Region where the regional interconnect attachment resides.
	Region interface{}
	// URL of the cloud router to be used for dynamic routing. This router must be in the same region as this
	// InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region
	// within which the Cloud Router is configured.
	Router interface{}
	// The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.
	Type interface{}
	// The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be managed
	// upstream.
	VlanTag8021q interface{}
}
