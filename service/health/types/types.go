// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an entity that is affected by a Health event.
type AffectedEntity struct {
	
	// The 12-digit Amazon Web Services account number that contains the affected
	// entity.
	AwsAccountId *string
	
	// The unique identifier for the entity. Format:
	// arn:aws:health:entity-region:aws-account:entity/entity-id . Example:
	// arn:aws:health:us-east-1:111222333444:entity/AVh5GGT7ul1arKr1sE1K
	EntityArn *string
	
	// The URL of the affected entity.
	EntityUrl *string
	
	// The ID of the affected entity.
	EntityValue *string
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
	
	// The most recent time that the entity was updated.
	LastUpdatedTime *time.Time
	
	// The most recent status of the entity affected by the event. The possible values
	// are IMPAIRED , UNIMPAIRED , and UNKNOWN .
	StatusCode EntityStatusCode
	
	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

// A range of dates and times that is used by the EventFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
// and EntityFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
// objects. If from is set and to is set: match items where the timestamp (
// startTime , endTime , or lastUpdatedTime ) is between from and to inclusive. If
// from is set and to is not set: match items where the timestamp value is equal
// to or after from . If from is not set and to is set: match items where the
// timestamp value is equal to or before to .
type DateTimeRange struct {
	
	// The starting date and time of a time range.
	From *time.Time
	
	// The ending date and time of a time range.
	To *time.Time
	
	noSmithyDocumentSerde
}

// The number of entities that are affected by one or more events. Returned by the
// DescribeEntityAggregates (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEntityAggregates.html)
// operation.
type EntityAggregate struct {
	
	// The number of entities that match the criteria for the specified events.
	Count int32
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
	
	noSmithyDocumentSerde
}

// The values to use to filter results from the DescribeAffectedEntities (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntities.html)
// operation.
type EntityFilter struct {
	
	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// This member is required.
	EventArns []string
	
	// A list of entity ARNs (unique identifiers).
	EntityArns []string
	
	// A list of IDs for affected entities.
	EntityValues []string
	
	// A list of the most recent dates and times that the entity was updated.
	LastUpdatedTimes []DateTimeRange
	
	// A list of entity status codes ( IMPAIRED , UNIMPAIRED , or UNKNOWN ).
	StatusCodes []EntityStatusCode
	
	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags []map[string]string
	
	noSmithyDocumentSerde
}

// Summary information about an Health event. Health events can be public or
// account-specific:
//   - Public events might be service events that are not specific to an Amazon
//   Web Services account. For example, if there is an issue with an Amazon Web
//   Services Region, Health provides information about the event, even if you don't
//   use services or resources in that Region.
//   - Account-specific events are specific to either your Amazon Web Services
//   account or an account in your organization. For example, if there's an issue
//   with Amazon Elastic Compute Cloud in a Region that you use, Health provides
//   information about the event and the affected resources in the account.
// You can determine if an event is public or account-specific by using the
// eventScopeCode parameter. For more information, see eventScopeCode (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html#AWSHealth-Type-Event-eventScopeCode)
// .
type Event struct {
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string
	
	// The Amazon Web Services Availability Zone of the event. For example, us-east-1a.
	AvailabilityZone *string
	
	// The date and time that the event ended.
	EndTime *time.Time
	
	// This parameter specifies if the Health event is a public Amazon Web Service
	// event or an account-specific event.
	//   - If the eventScopeCode value is PUBLIC , then the affectedAccounts value is
	//   always empty.
	//   - If the eventScopeCode value is ACCOUNT_SPECIFIC , then the affectedAccounts
	//   value lists the affected Amazon Web Services accounts in your organization. For
	//   example, if an event affects a service such as Amazon Elastic Compute Cloud and
	//   you have Amazon Web Services accounts that use that service, those account IDs
	//   appear in the response.
	//   - If the eventScopeCode value is NONE , then the eventArn that you specified
	//   in the request is invalid or doesn't exist.
	EventScopeCode EventScopeCode
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	EventTypeCategory EventTypeCategory
	
	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT .
	EventTypeCode *string
	
	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time
	
	// The Amazon Web Services Region name of the event.
	Region *string
	
	// The Amazon Web Service that is affected by the event. For example, EC2 , RDS .
	Service *string
	
	// The date and time that the event began.
	StartTime *time.Time
	
	// The most recent status of the event. Possible values are open , closed , and
	// upcoming .
	StatusCode EventStatusCode
	
	noSmithyDocumentSerde
}

// The values used to filter results from the DescribeEventDetailsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// and DescribeAffectedEntitiesForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operations.
type EventAccountFilter struct {
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	//
	// This member is required.
	EventArn *string
	
	// The 12-digit Amazon Web Services account numbers that contains the affected
	// entities.
	AwsAccountId *string
	
	noSmithyDocumentSerde
}

// The number of events of each issue type. Returned by the DescribeEventAggregates (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventAggregates.html)
// operation.
type EventAggregate struct {
	
	// The issue type for the associated count.
	AggregateValue *string
	
	// The number of events of the associated issue type.
	Count int32
	
	noSmithyDocumentSerde
}

// The detailed description of the event. Included in the information returned by
// the DescribeEventDetails (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation.
type EventDescription struct {
	
	// The most recent description of the event.
	LatestDescription *string
	
	noSmithyDocumentSerde
}

// Detailed information about an event. A combination of an Event (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html)
// object, an EventDescription (https://docs.aws.amazon.com/health/latest/APIReference/API_EventDescription.html)
// object, and additional metadata about the event. Returned by the
// DescribeEventDetails (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation.
type EventDetails struct {
	
	// Summary information about the event.
	Event *Event
	
	// The most recent description of the event.
	EventDescription *EventDescription
	
	// Additional metadata about the event.
	EventMetadata map[string]string
	
	noSmithyDocumentSerde
}

// Error information returned when a DescribeEventDetails (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
// operation can't find a specified event.
type EventDetailsErrorItem struct {
	
	// A message that describes the error.
	ErrorMessage *string
	
	// The name of the error.
	ErrorName *string
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
	
	noSmithyDocumentSerde
}

// The values to use to filter results from the DescribeEvents (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEvents.html)
// and DescribeEventAggregates (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventAggregates.html)
// operations.
type EventFilter struct {
	
	// A list of Amazon Web Services Availability Zones.
	AvailabilityZones []string
	
	// A list of dates and times that the event ended.
	EndTimes []DateTimeRange
	
	// A list of entity ARNs (unique identifiers).
	EntityArns []string
	
	// A list of entity identifiers, such as EC2 instance IDs ( i-34ab692e ) or EBS
	// volumes ( vol-426ab23e ).
	EntityValues []string
	
	// A list of event ARNs (unique identifiers). For example:
	// "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	EventArns []string
	
	// A list of event status codes.
	EventStatusCodes []EventStatusCode
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	EventTypeCategories []EventTypeCategory
	
	// A list of unique identifiers for event types. For example,
	// "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []string
	
	// A list of dates and times that the event was last updated.
	LastUpdatedTimes []DateTimeRange
	
	// A list of Amazon Web Services Regions.
	Regions []string
	
	// The Amazon Web Services associated with the event. For example, EC2 , RDS .
	Services []string
	
	// A list of dates and times that the event began.
	StartTimes []DateTimeRange
	
	// A map of entity tags attached to the affected entity. Currently, the tags
	// property isn't supported.
	Tags []map[string]string
	
	noSmithyDocumentSerde
}

// Contains the metadata about a type of event that is reported by Health. The
// EventType shows the category, service, and the event type code of the event. For
// example, an issue might be the category, EC2 the service, and
// AWS_EC2_SYSTEM_MAINTENANCE_EVENT the event type code. You can use the
// DescribeEventTypes (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventTypes.html)
// API operation to return this information about an event. You can also use the
// Amazon CloudWatch Events console to create a rule so that you can get notified
// or take action when Health delivers a specific event to your Amazon Web Services
// account. For more information, see Monitor for Health events with Amazon
// CloudWatch Events (https://docs.aws.amazon.com/health/latest/ug/cloudwatch-events-health.html)
// in the Health User Guide.
type EventType struct {
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	Category EventTypeCategory
	
	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// ; for example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT .
	Code *string
	
	// The Amazon Web Service that is affected by the event. For example, EC2 , RDS .
	Service *string
	
	noSmithyDocumentSerde
}

// The values to use to filter results from the DescribeEventTypes (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventTypes.html)
// operation.
type EventTypeFilter struct {
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	EventTypeCategories []EventTypeCategory
	
	// A list of event type codes.
	EventTypeCodes []string
	
	// The Amazon Web Services associated with the event. For example, EC2 , RDS .
	Services []string
	
	noSmithyDocumentSerde
}

// Error information returned when a DescribeAffectedEntitiesForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeAffectedEntitiesForOrganization.html)
// operation can't find or process a specific entity.
type OrganizationAffectedEntitiesErrorItem struct {
	
	// The 12-digit Amazon Web Services account numbers that contains the affected
	// entities.
	AwsAccountId *string
	
	// A message that describes the error. Follow the error message and retry your
	// request. For example, the InvalidAccountInputError error message appears if you
	// call the DescribeAffectedEntitiesForOrganization operation and specify the
	// AccountSpecific value for the EventScopeCode parameter, but don't specify an
	// Amazon Web Services account.
	ErrorMessage *string
	
	// The name of the error.
	ErrorName *string
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
	
	noSmithyDocumentSerde
}

// Summary information about an event, returned by the
// DescribeEventsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventsForOrganization.html)
// operation.
type OrganizationEvent struct {
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	Arn *string
	
	// The date and time that the event ended.
	EndTime *time.Time
	
	// This parameter specifies if the Health event is a public Amazon Web Service
	// event or an account-specific event.
	//   - If the eventScopeCode value is PUBLIC , then the affectedAccounts value is
	//   always empty.
	//   - If the eventScopeCode value is ACCOUNT_SPECIFIC , then the affectedAccounts
	//   value lists the affected Amazon Web Services accounts in your organization. For
	//   example, if an event affects a service such as Amazon Elastic Compute Cloud and
	//   you have Amazon Web Services accounts that use that service, those account IDs
	//   appear in the response.
	//   - If the eventScopeCode value is NONE , then the eventArn that you specified
	//   in the request is invalid or doesn't exist.
	EventScopeCode EventScopeCode
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	EventTypeCategory EventTypeCategory
	
	// The unique identifier for the event type. The format is AWS_SERVICE_DESCRIPTION
	// . For example, AWS_EC2_SYSTEM_MAINTENANCE_EVENT .
	EventTypeCode *string
	
	// The most recent date and time that the event was updated.
	LastUpdatedTime *time.Time
	
	// The Amazon Web Services Region name of the event.
	Region *string
	
	// The Amazon Web Service that is affected by the event, such as EC2 and RDS.
	Service *string
	
	// The date and time that the event began.
	StartTime *time.Time
	
	// The most recent status of the event. Possible values are open , closed , and
	// upcoming .
	StatusCode EventStatusCode
	
	noSmithyDocumentSerde
}

// Detailed information about an event. A combination of an Event (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html)
// object, an EventDescription (https://docs.aws.amazon.com/health/latest/APIReference/API_EventDescription.html)
// object, and additional metadata about the event. Returned by the
// DescribeEventDetailsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// operation.
type OrganizationEventDetails struct {
	
	// The 12-digit Amazon Web Services account numbers that contains the affected
	// entities.
	AwsAccountId *string
	
	// Summary information about an Health event. Health events can be public or
	// account-specific:
	//   - Public events might be service events that are not specific to an Amazon
	//   Web Services account. For example, if there is an issue with an Amazon Web
	//   Services Region, Health provides information about the event, even if you don't
	//   use services or resources in that Region.
	//   - Account-specific events are specific to either your Amazon Web Services
	//   account or an account in your organization. For example, if there's an issue
	//   with Amazon Elastic Compute Cloud in a Region that you use, Health provides
	//   information about the event and the affected resources in the account.
	// You can determine if an event is public or account-specific by using the
	// eventScopeCode parameter. For more information, see eventScopeCode (https://docs.aws.amazon.com/health/latest/APIReference/API_Event.html#AWSHealth-Type-Event-eventScopeCode)
	// .
	Event *Event
	
	// The detailed description of the event. Included in the information returned by
	// the DescribeEventDetails (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetails.html)
	// operation.
	EventDescription *EventDescription
	
	// Additional metadata about the event.
	EventMetadata map[string]string
	
	noSmithyDocumentSerde
}

// Error information returned when a DescribeEventDetailsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
// operation can't find a specified event.
type OrganizationEventDetailsErrorItem struct {
	
	// Error information returned when a DescribeEventDetailsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventDetailsForOrganization.html)
	// operation can't find a specified event.
	AwsAccountId *string
	
	// A message that describes the error. If you call the
	// DescribeEventDetailsForOrganization operation and receive one of the following
	// errors, follow the recommendations in the message:
	//   - We couldn't find a public event that matches your request. To find an event
	//   that is account specific, you must enter an Amazon Web Services account ID in
	//   the request.
	//   - We couldn't find an account specific event for the specified Amazon Web
	//   Services account. To find an event that is public, you must enter a null value
	//   for the Amazon Web Services account ID in the request.
	//   - Your Amazon Web Services account doesn't include the Amazon Web Services
	//   Support plan required to use the Health API. You must have either a Business,
	//   Enterprise On-Ramp, or Enterprise Support plan.
	ErrorMessage *string
	
	// The name of the error.
	ErrorName *string
	
	// The unique identifier for the event. The event ARN has the
	// arn:aws:health:event-region::event/SERVICE/EVENT_TYPE_CODE/EVENT_TYPE_PLUS_ID
	// format. For example, an event ARN might look like the following:
	// arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-DEF456
	EventArn *string
	
	noSmithyDocumentSerde
}

// The values to filter results from the DescribeEventsForOrganization (https://docs.aws.amazon.com/health/latest/APIReference/API_DescribeEventsForOrganization.html)
// operation.
type OrganizationEventFilter struct {
	
	// A list of 12-digit Amazon Web Services account numbers that contains the
	// affected entities.
	AwsAccountIds []string
	
	// A range of dates and times that is used by the EventFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp (
	// startTime , endTime , or lastUpdatedTime ) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal
	// to or after from . If from is not set and to is set: match items where the
	// timestamp value is equal to or before to .
	EndTime *DateTimeRange
	
	// A list of entity ARNs (unique identifiers).
	EntityArns []string
	
	// A list of entity identifiers, such as EC2 instance IDs (i-34ab692e) or EBS
	// volumes (vol-426ab23e).
	EntityValues []string
	
	// A list of event status codes.
	EventStatusCodes []EventStatusCode
	
	// A list of event type category codes. Possible values are issue ,
	// accountNotification , or scheduledChange . Currently, the investigation value
	// isn't supported at this time.
	EventTypeCategories []EventTypeCategory
	
	// A list of unique identifiers for event types. For example,
	// "AWS_EC2_SYSTEM_MAINTENANCE_EVENT","AWS_RDS_MAINTENANCE_SCHEDULED".
	EventTypeCodes []string
	
	// A range of dates and times that is used by the EventFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp (
	// startTime , endTime , or lastUpdatedTime ) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal
	// to or after from . If from is not set and to is set: match items where the
	// timestamp value is equal to or before to .
	LastUpdatedTime *DateTimeRange
	
	// A list of Amazon Web Services Regions.
	Regions []string
	
	// The Amazon Web Services associated with the event. For example, EC2 , RDS .
	Services []string
	
	// A range of dates and times that is used by the EventFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EventFilter.html)
	// and EntityFilter (https://docs.aws.amazon.com/health/latest/APIReference/API_EntityFilter.html)
	// objects. If from is set and to is set: match items where the timestamp (
	// startTime , endTime , or lastUpdatedTime ) is between from and to inclusive. If
	// from is set and to is not set: match items where the timestamp value is equal
	// to or after from . If from is not set and to is set: match items where the
	// timestamp value is equal to or before to .
	StartTime *DateTimeRange
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
