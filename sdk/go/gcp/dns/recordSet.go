// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a set of DNS records within Google Cloud DNS. For more information see [the official documentation](https://cloud.google.com/dns/records/) and
// [API](https://cloud.google.com/dns/api/v1/resourceRecordSets).
// 
// > **Note:** The provider treats this resource as an authoritative record set. This means existing records (including the default records) for the given type will be overwritten when you create this resource with this provider. In addition, the Google Cloud DNS API requires NS records to be present at all times, so this provider will not actually remove NS records during destroy but will report that it did.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/dns_record_set.html.markdown.
type RecordSet struct {
	s *pulumi.ResourceState
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOpt) (*RecordSet, error) {
	if args == nil || args.ManagedZone == nil {
		return nil, errors.New("missing required argument 'ManagedZone'")
	}
	if args == nil || args.Rrdatas == nil {
		return nil, errors.New("missing required argument 'Rrdatas'")
	}
	if args == nil || args.Ttl == nil {
		return nil, errors.New("missing required argument 'Ttl'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["managedZone"] = nil
		inputs["name"] = nil
		inputs["project"] = nil
		inputs["rrdatas"] = nil
		inputs["ttl"] = nil
		inputs["type"] = nil
	} else {
		inputs["managedZone"] = args.ManagedZone
		inputs["name"] = args.Name
		inputs["project"] = args.Project
		inputs["rrdatas"] = args.Rrdatas
		inputs["ttl"] = args.Ttl
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("gcp:dns/recordSet:RecordSet", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RecordSet{s: s}, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.ID, state *RecordSetState, opts ...pulumi.ResourceOpt) (*RecordSet, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["managedZone"] = state.ManagedZone
		inputs["name"] = state.Name
		inputs["project"] = state.Project
		inputs["rrdatas"] = state.Rrdatas
		inputs["ttl"] = state.Ttl
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("gcp:dns/recordSet:RecordSet", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &RecordSet{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *RecordSet) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *RecordSet) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the zone in which this record set will
// reside.
func (r *RecordSet) ManagedZone() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["managedZone"])
}

// The DNS name this record set will apply to.
func (r *RecordSet) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The ID of the project in which the resource belongs. If it
// is not provided, the provider project is used.
func (r *RecordSet) Project() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["project"])
}

// The string data for the records in this record set
// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
func (r *RecordSet) Rrdatas() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["rrdatas"])
}

// The time-to-live of this record set (seconds).
func (r *RecordSet) Ttl() pulumi.IntOutput {
	return (pulumi.IntOutput)(r.s.State["ttl"])
}

// The DNS record set type.
func (r *RecordSet) Type() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering RecordSet resources.
type RecordSetState struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone interface{}
	// The DNS name this record set will apply to.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The string data for the records in this record set
	// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
	Rrdatas interface{}
	// The time-to-live of this record set (seconds).
	Ttl interface{}
	// The DNS record set type.
	Type interface{}
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// The name of the zone in which this record set will
	// reside.
	ManagedZone interface{}
	// The DNS name this record set will apply to.
	Name interface{}
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project interface{}
	// The string data for the records in this record set
	// whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding `\"` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside this provider's configuration string (e.g. `"first255characters\"\"morecharacters"`).
	Rrdatas interface{}
	// The time-to-live of this record set (seconds).
	Ttl interface{}
	// The DNS record set type.
	Type interface{}
}
