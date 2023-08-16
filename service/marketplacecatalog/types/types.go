// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// An object that contains the ChangeType , Details , and Entity .
type Change struct {
	
	// Change types are single string values that describe your intention for the
	// change. Each change type is unique for each EntityType provided in the change's
	// scope. For more information on change types available for single-AMI products,
	// see Working with single-AMI products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products)
	// . Also, for more information on change types available for container-based
	// products, see Working with container products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products)
	// .
	//
	// This member is required.
	ChangeType *string
	
	// This object contains details specific to the change type of the requested
	// change. For more information on change types available for single-AMI products,
	// see Working with single-AMI products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products)
	// . Also, for more information on change types available for container-based
	// products, see Working with container products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products)
	// .
	//
	// This member is required.
	Details *string
	
	// The entity to be changed.
	//
	// This member is required.
	Entity *Entity
	
	// Optional name for the change.
	ChangeName *string
	
	// The tags associated with the change.
	EntityTags []Tag
	
	noSmithyDocumentSerde
}

// A summary of a change set returned in a list of change sets when the
// ListChangeSets action is called.
type ChangeSetSummaryListItem struct {
	
	// The ARN associated with the unique identifier for the change set referenced in
	// this request.
	ChangeSetArn *string
	
	// The unique identifier for a change set.
	ChangeSetId *string
	
	// The non-unique name for the change set.
	ChangeSetName *string
	
	// The time, in ISO 8601 format (2018-02-27T13:45:22Z), when the change set was
	// finished.
	EndTime *string
	
	// This object is a list of entity IDs (string) that are a part of a change set.
	// The entity ID list is a maximum of 20 entities. It must contain at least one
	// entity.
	EntityIdList []string
	
	// Returned if the change set is in FAILED status. Can be either CLIENT_ERROR ,
	// which means that there are issues with the request (see the ErrorDetailList of
	// DescribeChangeSet ), or SERVER_FAULT , which means that there is a problem in
	// the system, and you should retry your request.
	FailureCode FailureCode
	
	// The time, in ISO 8601 format (2018-02-27T13:45:22Z), when the change set was
	// started.
	StartTime *string
	
	// The current status of the change set.
	Status ChangeStatus
	
	noSmithyDocumentSerde
}

// This object is a container for common summary information about the change. The
// summary doesn't contain the whole change structure.
type ChangeSummary struct {
	
	// Optional name for the change.
	ChangeName *string
	
	// The type of the change.
	ChangeType *string
	
	// This object contains details specific to the change type of the requested
	// change.
	Details *string
	
	// The entity to be changed.
	Entity *Entity
	
	// An array of ErrorDetail objects associated with the change.
	ErrorDetailList []ErrorDetail
	
	noSmithyDocumentSerde
}

// An entity contains data that describes your product, its supported features,
// and how it can be used or launched by your customer.
type Entity struct {
	
	// The type of entity.
	//
	// This member is required.
	Type *string
	
	// The identifier for the entity.
	Identifier *string
	
	noSmithyDocumentSerde
}

// This object is a container for common summary information about the entity. The
// summary doesn't contain the whole entity structure, but it does contain
// information common across all entities.
type EntitySummary struct {
	
	// The ARN associated with the unique identifier for the entity.
	EntityArn *string
	
	// The unique identifier for the entity.
	EntityId *string
	
	// The type of the entity.
	EntityType *string
	
	// The last time the entity was published, using ISO 8601 format
	// (2018-02-27T13:45:22Z).
	LastModifiedDate *string
	
	// The name for the entity. This value is not unique. It is defined by the seller.
	Name *string
	
	// The visibility status of the entity to buyers. This value can be Public
	// (everyone can view the entity), Limited (the entity is visible to limited
	// accounts only), or Restricted (the entity was published and then unpublished
	// and only existing buyers can view it).
	Visibility *string
	
	noSmithyDocumentSerde
}

// Details about the error.
type ErrorDetail struct {
	
	// The error code that identifies the type of error.
	ErrorCode *string
	
	// The message for the error.
	ErrorMessage *string
	
	noSmithyDocumentSerde
}

// A filter object, used to optionally filter results from calls to the
// ListEntities and ListChangeSets actions.
type Filter struct {
	
	// For ListEntities , the supported value for this is an EntityId . For
	// ListChangeSets , the supported values are as follows:
	Name *string
	
	// ListEntities - This is a list of unique EntityId s. ListChangeSets - The
	// supported filter names and associated ValueList s is as follows:
	//   - ChangeSetName - The supported ValueList is a list of non-unique
	//   ChangeSetName s. These are defined when you call the StartChangeSet action.
	//   - Status - The supported ValueList is a list of statuses for all change set
	//   requests.
	//   - EntityId - The supported ValueList is a list of unique EntityId s.
	//   - BeforeStartTime - The supported ValueList is a list of all change sets that
	//   started before the filter value.
	//   - AfterStartTime - The supported ValueList is a list of all change sets that
	//   started after the filter value.
	//   - BeforeEndTime - The supported ValueList is a list of all change sets that
	//   ended before the filter value.
	//   - AfterEndTime - The supported ValueList is a list of all change sets that
	//   ended after the filter value.
	ValueList []string
	
	noSmithyDocumentSerde
}

// An object that contains two attributes, SortBy and SortOrder .
type Sort struct {
	
	// For ListEntities , supported attributes include LastModifiedDate (default),
	// Visibility , EntityId , and Name . For ListChangeSets , supported attributes
	// include StartTime and EndTime .
	SortBy *string
	
	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder
	
	noSmithyDocumentSerde
}

// A list of objects specifying each key name and value.
type Tag struct {
	
	// The key associated with the tag.
	//
	// This member is required.
	Key *string
	
	// The value associated with the tag.
	//
	// This member is required.
	Value *string
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
