// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// User-provided application code (query) is invalid. This can be a simple syntax
// error.
type CodeValidationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *CodeValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CodeValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *CodeValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exception thrown as a result of concurrent modification to an application. For
// example, two individuals attempting to edit the same application at the same
// time.
type ConcurrentModificationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConcurrentModificationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// User-provided application configuration is not valid.
type InvalidApplicationConfigurationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidApplicationConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidApplicationConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidApplicationConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidApplicationConfigurationException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidApplicationConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Specified input parameter value is invalid.
type InvalidArgumentException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArgumentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Exceeded the number of applications allowed.
type LimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Application is not available for this operation.
type ResourceInUseException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Specified application can't be found.
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

// Discovery failed to get a record from the streaming source because of the
// Amazon Kinesis Streams ProvisionedThroughputExceededException. For more
// information, see GetRecords (https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetRecords.html)
// in the Amazon Kinesis Streams API Reference.
type ResourceProvisionedThroughputExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceProvisionedThroughputExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceProvisionedThroughputExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceProvisionedThroughputExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceProvisionedThroughputExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceProvisionedThroughputExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is unavailable. Back off and retry the operation.
type ServiceUnavailableException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Application created with too many tags, or too many tags added to an
// application. Note that the maximum number of application tags includes system
// tags. The maximum number of user-defined application tags is 50.
type TooManyTagsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTagsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Data format is not valid. Amazon Kinesis Analytics is not able to detect schema
// for the given streaming source.
type UnableToDetectSchemaException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	RawInputRecords []string
	ProcessedInputRecords []string
	
	noSmithyDocumentSerde
}

func (e *UnableToDetectSchemaException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnableToDetectSchemaException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnableToDetectSchemaException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnableToDetectSchemaException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnableToDetectSchemaException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request was rejected because a specified parameter is not supported or a
// specified resource is not valid for this operation.
type UnsupportedOperationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedOperationException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
