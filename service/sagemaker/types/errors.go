// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// There was a conflict when you attempted to modify a SageMaker entity such as an
// Experiment or Artifact .
type ConflictException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Resource being accessed is in use.
type ResourceInUse struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUse) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded an SageMaker resource limit. For example, you might have too
// many training jobs created.
type ResourceLimitExceeded struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceLimitExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceeded) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceLimitExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Resource being access is not found.
type ResourceNotFound struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFound) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
