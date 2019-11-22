// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
//  For more information see 
// [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications) 
// and 
// [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
// 
// In order to enable notifications, a special Google Cloud Storage service account unique to the project
// must have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project. To get the service
// account's email address, use the `storage.getProjectServiceAccount` datasource's `emailAddress` value, and see below
// for an example of enabling notifications by granting the correct IAM permission. See
// [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/storage_notification.html.markdown.
type Notification struct {
	s *pulumi.ResourceState
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOpt) (*Notification, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.PayloadFormat == nil {
		return nil, errors.New("missing required argument 'PayloadFormat'")
	}
	if args == nil || args.Topic == nil {
		return nil, errors.New("missing required argument 'Topic'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["bucket"] = nil
		inputs["customAttributes"] = nil
		inputs["eventTypes"] = nil
		inputs["objectNamePrefix"] = nil
		inputs["payloadFormat"] = nil
		inputs["topic"] = nil
	} else {
		inputs["bucket"] = args.Bucket
		inputs["customAttributes"] = args.CustomAttributes
		inputs["eventTypes"] = args.EventTypes
		inputs["objectNamePrefix"] = args.ObjectNamePrefix
		inputs["payloadFormat"] = args.PayloadFormat
		inputs["topic"] = args.Topic
	}
	inputs["notificationId"] = nil
	inputs["selfLink"] = nil
	s, err := ctx.RegisterResource("gcp:storage/notification:Notification", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Notification{s: s}, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.ID, state *NotificationState, opts ...pulumi.ResourceOpt) (*Notification, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["bucket"] = state.Bucket
		inputs["customAttributes"] = state.CustomAttributes
		inputs["eventTypes"] = state.EventTypes
		inputs["notificationId"] = state.NotificationId
		inputs["objectNamePrefix"] = state.ObjectNamePrefix
		inputs["payloadFormat"] = state.PayloadFormat
		inputs["selfLink"] = state.SelfLink
		inputs["topic"] = state.Topic
	}
	s, err := ctx.ReadResource("gcp:storage/notification:Notification", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Notification{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Notification) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Notification) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The name of the bucket.
func (r *Notification) Bucket() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["bucket"])
}

// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
func (r *Notification) CustomAttributes() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["customAttributes"])
}

// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
func (r *Notification) EventTypes() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["eventTypes"])
}

// The ID of the created notification.
func (r *Notification) NotificationId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["notificationId"])
}

// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
func (r *Notification) ObjectNamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["objectNamePrefix"])
}

// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
func (r *Notification) PayloadFormat() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["payloadFormat"])
}

// The URI of the created resource.
func (r *Notification) SelfLink() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["selfLink"])
}

// The Cloud PubSub topic to which this subscription publishes. Expects either the 
// topic name, assumed to belong to the default GCP provider project, or the project-level name,
// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
func (r *Notification) Topic() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["topic"])
}

// Input properties used for looking up and filtering Notification resources.
type NotificationState struct {
	// The name of the bucket.
	Bucket interface{}
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes interface{}
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes interface{}
	// The ID of the created notification.
	NotificationId interface{}
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix interface{}
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat interface{}
	// The URI of the created resource.
	SelfLink interface{}
	// The Cloud PubSub topic to which this subscription publishes. Expects either the 
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
	Topic interface{}
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// The name of the bucket.
	Bucket interface{}
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes interface{}
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes interface{}
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix interface{}
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat interface{}
	// The Cloud PubSub topic to which this subscription publishes. Expects either the 
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`.
	Topic interface{}
}
