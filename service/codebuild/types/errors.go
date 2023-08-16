// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// An Amazon Web Services service limit was exceeded for the calling Amazon Web
// Services account.
type AccountLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AccountLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccountLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccountLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input value that was provided is not valid.
type InvalidInputException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidInputException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was a problem with the underlying OAuth provider.
type OAuthProviderException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *OAuthProviderException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OAuthProviderException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OAuthProviderException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OAuthProviderException"
	}
	return *e.ErrorCodeOverride
}
func (e *OAuthProviderException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified Amazon Web Services resource cannot be created, because an Amazon
// Web Services resource with the same settings already exists.
type ResourceAlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified Amazon Web Services resource cannot be found.
type ResourceNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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
