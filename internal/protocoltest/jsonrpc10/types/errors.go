// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This error is thrown when a request is invalid.
type ComplexError struct {
	Message *string
	
	ErrorCodeOverride *string
	
	TopLevel *string
	Nested *ComplexNestedErrorData
	
	noSmithyDocumentSerde
}

func (e *ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ComplexError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ComplexError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ComplexError"
	}
	return *e.ErrorCodeOverride
}
func (e *ComplexError) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error has test cases that test some of the dark corners of Amazon service
// framework history. It should only be implemented by clients.
type FooError struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *FooError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FooError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FooError) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FooError"
	}
	return *e.ErrorCodeOverride
}
func (e *FooError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// This error is thrown when an invalid greeting value is provided.
type InvalidGreeting struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidGreeting) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidGreeting) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidGreeting) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidGreeting"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidGreeting) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
