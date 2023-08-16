// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You do not have sufficient access to perform this action.
type AccessDeniedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	
	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be completed due to a conflict with the current state of
// the target resource.
type ConflictException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	ResourceId *string
	ResourceType *string
	
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

// The request processing has failed because of an unknown error, exception or
// failure.
type InternalServerException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	RetryAfterSeconds int64
	
	noSmithyDocumentSerde
}

func (e *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServerException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The resource for this operation was not found.
type ResourceNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	ResourceId *string
	ResourceType *string
	
	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request could not be completed because its exceeded the service quota.
type ServiceQuotaExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	ResourceId *string
	ResourceType *string
	ServiceCode *string
	QuotaCode *string
	
	noSmithyDocumentSerde
}

func (e *ServiceQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceQuotaExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceQuotaExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling.
type ThrottlingException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ServiceCode *string
	QuotaCode *string
	RetryAfterSeconds *string
	
	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The account performing the request has not been initialized.
type UninitializedAccountException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	
	noSmithyDocumentSerde
}

func (e *UninitializedAccountException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UninitializedAccountException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UninitializedAccountException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UninitializedAccountException"
	}
	return *e.ErrorCodeOverride
}
func (e *UninitializedAccountException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input fails to satisfy the constraints specified by the AWS service.
type ValidationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Code *string
	Reason ValidationExceptionReason
	FieldList []ValidationExceptionField
	
	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
