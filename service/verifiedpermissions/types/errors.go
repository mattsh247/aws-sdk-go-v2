// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You don't have sufficient access to perform this action.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

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

// The request failed because another request to modify a resource occurred at the
// same.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	Resources []ResourceConflict

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

// The request failed because of an internal error. Try your request again later
type InternalServerException struct {
	Message *string

	ErrorCodeOverride *string

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

// The request failed because it references a resource that doesn't exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceId   *string
	ResourceType ResourceType

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

// The request failed because it would cause a service quota to be exceeded.
type ServiceQuotaExceededException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceId   *string
	ResourceType ResourceType
	ServiceCode  *string
	QuotaCode    *string

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

// The request failed because it exceeded a throttling quota.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	ServiceCode *string
	QuotaCode   *string

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

// The request failed because one or more input parameters don't satisfy their
// constraint requirements. The output is provided as a list of fields and a reason
// for each field that isn't valid. The possible reasons include the following:
//   - UnrecognizedEntityType The policy includes an entity type that isn't found
//     in the schema.
//   - UnrecognizedActionId The policy includes an action id that isn't found in
//     the schema.
//   - InvalidActionApplication The policy includes an action that, according to
//     the schema, doesn't support the specified principal and resource.
//   - UnexpectedType The policy included an operand that isn't a valid type for
//     the specified operation.
//   - IncompatibleTypes The types of elements included in a set , or the types of
//     expressions used in an if...then...else clause aren't compatible in this
//     context.
//   - MissingAttribute The policy attempts to access a record or entity attribute
//     that isn't specified in the schema. Test for the existence of the attribute
//     first before attempting to access its value. For more information, see the
//     has (presence of attribute test) operator (https://docs.cedarpolicy.com/syntax-operators.html#has-presence-of-attribute-test)
//     in the Cedar Policy Language Guide.
//   - UnsafeOptionalAttributeAccess The policy attempts to access a record or
//     entity attribute that is optional and isn't guaranteed to be present. Test for
//     the existence of the attribute first before attempting to access its value. For
//     more information, see the has (presence of attribute test) operator (https://docs.cedarpolicy.com/syntax-operators.html#has-presence-of-attribute-test)
//     in the Cedar Policy Language Guide.
//   - ImpossiblePolicy Cedar has determined that a policy condition always
//     evaluates to false. If the policy is always false, it can never apply to any
//     query, and so it can never affect an authorization decision.
//   - WrongNumberArguments The policy references an extension type with the wrong
//     number of arguments.
//   - FunctionArgumentValidationError Cedar couldn't parse the argument passed to
//     an extension type. For example, a string that is to be parsed as an IPv4 address
//     can contain only digits and the period character.
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
