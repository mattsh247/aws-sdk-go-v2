// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/document"
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A collection of error messages for any views that Amazon Web Services Resource
// Explorer couldn't retrieve details.
type BatchGetViewError struct {
	
	// The description of the error for the specified view.
	//
	// This member is required.
	ErrorMessage *string
	
	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view for which Resource Explorer failed to retrieve details.
	//
	// This member is required.
	ViewArn *string
	
	noSmithyDocumentSerde
}

// Information about an additional property that describes a resource, that you
// can optionally include in the view. This lets you view that property in search
// results, and filter your search results based on the value of the property.
type IncludedProperty struct {
	
	// The name of the property that is included in this view. You can specify the
	// following property names for this field:
	//   - Tags
	//
	// This member is required.
	Name *string
	
	noSmithyDocumentSerde
}

// An index is the data store used by Amazon Web Services Resource Explorer to
// hold information about your Amazon Web Services resources that the service
// discovers. Creating an index in an Amazon Web Services Region turns on Resource
// Explorer and lets it discover your resources. By default, an index is local,
// meaning that it contains information about resources in only the same Region as
// the index. However, you can promote the index of one Region in the account by
// calling UpdateIndexType to convert it into an aggregator index. The aggregator
// index receives a replicated copy of the index information from all other Regions
// where Resource Explorer is turned on. This allows search operations in that
// Region to return results from all Regions in the account.
type Index struct {
	
	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the index.
	Arn *string
	
	// The Amazon Web Services Region in which the index exists.
	Region *string
	
	// The type of index. It can be one of the following values:
	//   - LOCAL – The index contains information about resources from only the same
	//   Amazon Web Services Region.
	//   - AGGREGATOR – Resource Explorer replicates copies of the indexed information
	//   about resources in all other Amazon Web Services Regions to the aggregator
	//   index. This lets search results in the Region with the aggregator index to
	//   include resources from all Regions in the account where Resource Explorer is
	//   turned on.
	Type IndexType
	
	noSmithyDocumentSerde
}

// A resource in Amazon Web Services that Amazon Web Services Resource Explorer
// has discovered, and for which it has stored information in the index of the
// Amazon Web Services Region that contains the resource.
type Resource struct {
	
	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource.
	Arn *string
	
	// The date and time that Resource Explorer last queried this resource and updated
	// the index with the latest information about the resource.
	LastReportedAt *time.Time
	
	// The Amazon Web Services account that owns the resource.
	OwningAccountId *string
	
	// A structure with additional type-specific details about the resource. These
	// properties can be added by turning on integration between Resource Explorer and
	// other Amazon Web Services services.
	Properties []ResourceProperty
	
	// The Amazon Web Services Region in which the resource was created and exists.
	Region *string
	
	// The type of the resource.
	ResourceType *string
	
	// The Amazon Web Service that owns the resource and is responsible for creating
	// and updating it.
	Service *string
	
	noSmithyDocumentSerde
}

// Information about the number of results that match the query. At this time,
// Amazon Web Services Resource Explorer doesn't count more than 1,000 matches for
// any query. This structure provides information about whether the query exceeded
// this limit. This field is included in every page when you paginate the results.
type ResourceCount struct {
	
	// Indicates whether the TotalResources value represents an exhaustive count of
	// search results.
	//   - If True , it indicates that the search was exhaustive. Every resource that
	//   matches the query was counted.
	//   - If False , then the search reached the limit of 1,000 matching results, and
	//   stopped counting.
	Complete *bool
	
	// The number of resources that match the search query. This value can't exceed
	// 1,000. If there are more than 1,000 resources that match the query, then only
	// 1,000 are counted and the Complete field is set to false. We recommend that you
	// refine your query to return a smaller number of results.
	TotalResources *int64
	
	noSmithyDocumentSerde
}

// A structure that describes a property of a resource.
type ResourceProperty struct {
	
	// Details about this property. The content of this field is a JSON object that
	// varies based on the resource type.
	Data document.Interface
	
	// The date and time that the information about this resource property was last
	// updated.
	LastReportedAt *time.Time
	
	// The name of this property of the resource.
	Name *string
	
	noSmithyDocumentSerde
}

// A search filter defines which resources can be part of a search query result
// set.
type SearchFilter struct {
	
	// The string that contains the search keywords, prefixes, and operators to
	// control the results that can be returned by a Search operation. For more
	// details, see Search query syntax (https://docs.aws.amazon.com/resource-explorer/latest/APIReference/about-query-syntax.html)
	// .
	//
	// This member is required.
	FilterString *string
	
	noSmithyDocumentSerde
}

// A structure that describes a resource type supported by Amazon Web Services
// Resource Explorer.
type SupportedResourceType struct {
	
	// The unique identifier of the resource type.
	ResourceType *string
	
	// The Amazon Web Service that is associated with the resource type. This is the
	// primary service that lets you create and interact with resources of this type.
	Service *string
	
	noSmithyDocumentSerde
}

// A structure that describes a request field with a validation error.
type ValidationExceptionField struct {
	
	// The name of the request field that had a validation error.
	//
	// This member is required.
	Name *string
	
	// The validation error caused by the request field.
	//
	// This member is required.
	ValidationIssue *string
	
	noSmithyDocumentSerde
}

// A view is a structure that defines a set of filters that provide a view into
// the information in the Amazon Web Services Resource Explorer index. The filters
// specify which information from the index is visible to the users of the view.
// For example, you can specify filters that include only resources that are tagged
// with the key "ENV" and the value "DEVELOPMENT" in the results returned by this
// view. You could also create a second view that includes only resources that are
// tagged with "ENV" and "PRODUCTION".
type View struct {
	
	// An array of SearchFilter objects that specify which resources can be included
	// in the results of queries made using this view.
	Filters *SearchFilter
	
	// A structure that contains additional information about the view.
	IncludedProperties []IncludedProperty
	
	// The date and time when this view was last modified.
	LastUpdatedAt *time.Time
	
	// The Amazon Web Services account that owns this view.
	Owner *string
	
	// An Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an Amazon Web Services account, an organization, or an organizational unit
	// (OU) that specifies whether this view includes resources from only the specified
	// Amazon Web Services account, all accounts in the specified organization, or all
	// accounts in the specified OU. If not specified, the value defaults to the Amazon
	// Web Services account used to call this operation.
	Scope *string
	
	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view.
	ViewArn *string
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
