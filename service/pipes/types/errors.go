// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An action you attempted resulted in an exception.
type ConflictException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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

// This exception occurs due to unexpected causes.
type InternalException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	RetryAfterSeconds *int32
	
	noSmithyDocumentSerde
}

func (e *InternalException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// An entity that you specified does not exist.
type NotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "NotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A quota has been exceeded.
type ServiceQuotaExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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

// An action was throttled.
type ThrottlingException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ServiceCode *string
	QuotaCode *string
	RetryAfterSeconds *int32
	
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

// Indicates that an error has occurred while performing a validate operation.
type ValidationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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
