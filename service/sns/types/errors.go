// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Indicates that the user has been denied access to the requested resource.
type AuthorizationErrorException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AuthorizationErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthorizationErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthorizationErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AuthorizationError"
	}
	return *e.ErrorCodeOverride
}
func (e *AuthorizationErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Two or more batch entries in the request have the same Id .
type BatchEntryIdsNotDistinctException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *BatchEntryIdsNotDistinctException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchEntryIdsNotDistinctException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchEntryIdsNotDistinctException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BatchEntryIdsNotDistinct"
	}
	return *e.ErrorCodeOverride
}
func (e *BatchEntryIdsNotDistinctException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The length of all the batch messages put together is more than the limit.
type BatchRequestTooLongException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *BatchRequestTooLongException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BatchRequestTooLongException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BatchRequestTooLongException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "BatchRequestTooLong"
	}
	return *e.ErrorCodeOverride
}
func (e *BatchRequestTooLongException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Can't perform multiple operations on a tag simultaneously. Perform the
// operations sequentially.
type ConcurrentAccessException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ConcurrentAccessException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentAccessException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentAccessException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConcurrentAccess"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentAccessException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request doesn't contain any entries.
type EmptyBatchRequestException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EmptyBatchRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EmptyBatchRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EmptyBatchRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EmptyBatchRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *EmptyBatchRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception error indicating endpoint disabled.
type EndpointDisabledException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EndpointDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EndpointDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EndpointDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EndpointDisabled"
	}
	return *e.ErrorCodeOverride
}
func (e *EndpointDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the number of filter polices in your Amazon Web Services account
// exceeds the limit. To add more filter polices, submit an Amazon SNS Limit
// Increase case in the Amazon Web Services Support Center.
type FilterPolicyLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *FilterPolicyLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FilterPolicyLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FilterPolicyLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FilterPolicyLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *FilterPolicyLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates an internal service error.
type InternalErrorException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The Id of a batch entry in a batch request doesn't abide by the specification.
type InvalidBatchEntryIdException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidBatchEntryIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidBatchEntryIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidBatchEntryIdException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidBatchEntryId"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidBatchEntryIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a request parameter does not comply with the associated
// constraints.
type InvalidParameterException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameter"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a request parameter does not comply with the associated
// constraints.
type InvalidParameterValueException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ParameterValueInvalid"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The credential signature isn't valid. You must use an HTTPS endpoint and sign
// your request using Signature Version 4.
type InvalidSecurityException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSecurityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecurityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecurityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSecurity"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecurityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The ciphertext references a key that doesn't exist or that you don't have
// access to.
type KMSAccessDeniedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSAccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSAccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSAccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSAccessDenied"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSAccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified customer master key (CMK) isn't
// enabled.
type KMSDisabledException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSDisabled"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the state of the specified resource isn't
// valid for this request. For more information, see How Key State Affects Use of
// a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide.
type KMSInvalidStateException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSInvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSInvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSInvalidStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSInvalidState"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSInvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because the specified entity or resource can't be
// found.
type KMSNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The Amazon Web Services access key ID needs a subscription for the service.
type KMSOptInRequired struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSOptInRequired) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSOptInRequired) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSOptInRequired) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSOptInRequired"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSOptInRequired) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was denied due to request throttling. For more information about
// throttling, see Limits (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html#requests-per-second)
// in the Key Management Service Developer Guide.
type KMSThrottlingException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *KMSThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *KMSThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *KMSThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "KMSThrottling"
	}
	return *e.ErrorCodeOverride
}
func (e *KMSThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the requested resource does not exist.
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
		return "NotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the specified phone number opted out of receiving SMS messages
// from your Amazon Web Services account. You can't send SMS messages to phone
// numbers that opt out.
type OptedOutException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *OptedOutException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OptedOutException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OptedOutException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OptedOut"
	}
	return *e.ErrorCodeOverride
}
func (e *OptedOutException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception error indicating platform application disabled.
type PlatformApplicationDisabledException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *PlatformApplicationDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PlatformApplicationDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PlatformApplicationDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PlatformApplicationDisabled"
	}
	return *e.ErrorCodeOverride
}
func (e *PlatformApplicationDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Can’t perform the action on the specified resource. Make sure that the resource
// exists.
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
		return "ResourceNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A tag has been added to a resource with the same ARN as a deleted resource.
// Wait a short while and then retry the operation.
type StaleTagException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *StaleTagException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *StaleTagException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *StaleTagException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "StaleTag"
	}
	return *e.ErrorCodeOverride
}
func (e *StaleTagException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the customer already owns the maximum allowed number of
// subscriptions.
type SubscriptionLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SubscriptionLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubscriptionLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubscriptionLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubscriptionLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *SubscriptionLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Can't add more than 50 tags to a topic.
type TagLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TagLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *TagLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request doesn't comply with the IAM tag policy. Correct your request and
// then retry it.
type TagPolicyException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TagPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TagPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TagPolicyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TagPolicy"
	}
	return *e.ErrorCodeOverride
}
func (e *TagPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the rate at which requests have been submitted for this action
// exceeds the limit for your Amazon Web Services account.
type ThrottledException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ThrottledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "Throttled"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The batch request contains more entries than permissible.
type TooManyEntriesInBatchRequestException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyEntriesInBatchRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyEntriesInBatchRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyEntriesInBatchRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyEntriesInBatchRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyEntriesInBatchRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the customer already owns the maximum allowed number of topics.
type TopicLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TopicLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TopicLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TopicLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TopicLimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *TopicLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a request parameter does not comply with the associated
// constraints.
type UserErrorException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *UserErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UserErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UserErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UserError"
	}
	return *e.ErrorCodeOverride
}
func (e *UserErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a parameter in the request is invalid.
type ValidationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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

// Indicates that the one-time password (OTP) used for verification is invalid.
type VerificationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Status *string
	
	noSmithyDocumentSerde
}

func (e *VerificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *VerificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *VerificationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "VerificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *VerificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
