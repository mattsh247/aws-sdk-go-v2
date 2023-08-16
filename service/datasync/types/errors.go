// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This exception is thrown when an error occurs in the DataSync service.
type InternalException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ErrorCode_ *string
	
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

// This exception is thrown when the client submits a malformed request.
type InvalidRequestException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ErrorCode_ *string
	DatasyncErrorCode *string
	
	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
