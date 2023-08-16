// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the current status of a resource operation request. For more
// information, see Managing resource operation requests (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html)
// in the Amazon Web Services Cloud Control API User Guide.
type ProgressEvent struct {

	// For requests with a status of FAILED , the associated error code. For error code
	// definitions, see Handler error codes (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-test-contract-errors.html)
	// in the CloudFormation Command Line Interface User Guide for Extension
	// Development.
	ErrorCode HandlerErrorCode

	// When the resource operation request was initiated.
	EventTime *time.Time

	// The primary identifier for the resource. In some cases, the resource identifier
	// may be available before the resource operation has reached a status of SUCCESS .
	Identifier *string

	// The resource operation type.
	Operation Operation

	// The current status of the resource operation request.
	//   - PENDING : The resource operation hasn't yet started.
	//   - IN_PROGRESS : The resource operation is currently in progress.
	//   - SUCCESS : The resource operation has successfully completed.
	//   - FAILED : The resource operation has failed. Refer to the error code and
	//   status message for more information.
	//   - CANCEL_IN_PROGRESS : The resource operation is in the process of being
	//   canceled.
	//   - CANCEL_COMPLETE : The resource operation has been canceled.
	OperationStatus OperationStatus

	// The unique token representing this resource operation request. Use the
	// RequestToken with GetResourceRequestStatus (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// to return the current status of a resource operation request.
	RequestToken *string

	// A JSON string containing the resource model, consisting of each resource
	// property and its current value.
	ResourceModel *string

	// When to next request the status of this resource operation request.
	RetryAfter *time.Time

	// Any message explaining the current status.
	StatusMessage *string

	// The name of the resource type used in the operation.
	TypeName *string

	noSmithyDocumentSerde
}

// Represents information about a provisioned resource.
type ResourceDescription struct {

	// The primary identifier for the resource. For more information, see Identifying
	// resources (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-identifier.html)
	// in the Amazon Web Services Cloud Control API User Guide.
	Identifier *string

	// A list of the resource properties and their current values.
	Properties *string

	noSmithyDocumentSerde
}

// The filter criteria to use in determining the requests returned.
type ResourceRequestStatusFilter struct {

	// The operation statuses to include in the filter.
	//   - PENDING : The operation has been requested, but not yet initiated.
	//   - IN_PROGRESS : The operation is in progress.
	//   - SUCCESS : The operation completed.
	//   - FAILED : The operation failed.
	//   - CANCEL_IN_PROGRESS : The operation is in the process of being canceled.
	//   - CANCEL_COMPLETE : The operation has been canceled.
	OperationStatuses []OperationStatus

	// The operation types to include in the filter.
	Operations []Operation

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
