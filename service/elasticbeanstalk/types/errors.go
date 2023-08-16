// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// AWS CodeBuild is not available in the specified region.
type CodeBuildNotInServiceRegionException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *CodeBuildNotInServiceRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CodeBuildNotInServiceRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CodeBuildNotInServiceRegionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CodeBuildNotInServiceRegionException"
	}
	return *e.ErrorCodeOverride
}
func (e *CodeBuildNotInServiceRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A generic service exception has occurred.
type ElasticBeanstalkServiceException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ElasticBeanstalkServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ElasticBeanstalkServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ElasticBeanstalkServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ElasticBeanstalkServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *ElasticBeanstalkServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account does not have sufficient privileges for one or more AWS
// services.
type InsufficientPrivilegesException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InsufficientPrivilegesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientPrivilegesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientPrivilegesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InsufficientPrivilegesException"
	}
	return *e.ErrorCodeOverride
}
func (e *InsufficientPrivilegesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more input parameters is not valid. Please correct the input parameters
// and try the operation again.
type InvalidRequestException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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

// Cannot modify the managed action in its current state.
type ManagedActionInvalidStateException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ManagedActionInvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ManagedActionInvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ManagedActionInvalidStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ManagedActionInvalidStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *ManagedActionInvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to perform the specified operation because another operation that
// effects an element in this activity is already in progress.
type OperationInProgressException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *OperationInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationInProgressException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OperationInProgressFailure"
	}
	return *e.ErrorCodeOverride
}
func (e *OperationInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You cannot delete the platform version because there are still environments
// running on it.
type PlatformVersionStillReferencedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *PlatformVersionStillReferencedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PlatformVersionStillReferencedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PlatformVersionStillReferencedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PlatformVersionStillReferencedException"
	}
	return *e.ErrorCodeOverride
}
func (e *PlatformVersionStillReferencedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource doesn't exist for the specified Amazon Resource Name (ARN).
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

// The type of the specified Amazon Resource Name (ARN) isn't supported for this
// operation.
type ResourceTypeNotSupportedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceTypeNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceTypeNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceTypeNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceTypeNotSupportedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceTypeNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified S3 bucket does not belong to the S3 region in which the service
// is running. The following regions are supported:
//   - IAD/us-east-1
//   - PDX/us-west-2
//   - DUB/eu-west-1
type S3LocationNotInServiceRegionException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *S3LocationNotInServiceRegionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3LocationNotInServiceRegionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3LocationNotInServiceRegionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "S3LocationNotInServiceRegionException"
	}
	return *e.ErrorCodeOverride
}
func (e *S3LocationNotInServiceRegionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account does not have a subscription to Amazon S3.
type S3SubscriptionRequiredException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *S3SubscriptionRequiredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *S3SubscriptionRequiredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *S3SubscriptionRequiredException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "S3SubscriptionRequiredException"
	}
	return *e.ErrorCodeOverride
}
func (e *S3SubscriptionRequiredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to delete the Amazon S3 source bundle associated with the application
// version. The application version was deleted successfully.
type SourceBundleDeletionException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SourceBundleDeletionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SourceBundleDeletionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SourceBundleDeletionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SourceBundleDeletionFailure"
	}
	return *e.ErrorCodeOverride
}
func (e *SourceBundleDeletionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account has reached its limit of applications.
type TooManyApplicationsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyApplicationsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyApplicationsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyApplicationsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyApplicationsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyApplicationsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account has reached its limit of application versions.
type TooManyApplicationVersionsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyApplicationVersionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyApplicationVersionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyApplicationVersionsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyApplicationVersionsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyApplicationVersionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account has reached its limit of Amazon S3 buckets.
type TooManyBucketsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyBucketsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyBucketsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyBucketsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyBucketsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyBucketsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account has reached its limit of configuration templates.
type TooManyConfigurationTemplatesException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyConfigurationTemplatesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyConfigurationTemplatesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyConfigurationTemplatesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyConfigurationTemplatesException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyConfigurationTemplatesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified account has reached its limit of environments.
type TooManyEnvironmentsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyEnvironmentsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyEnvironmentsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyEnvironmentsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyEnvironmentsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyEnvironmentsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You have exceeded the maximum number of allowed platforms associated with the
// account.
type TooManyPlatformsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyPlatformsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyPlatformsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyPlatformsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyPlatformsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyPlatformsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The number of tags in the resource would exceed the number of tags that each
// resource can have. To calculate this, the operation considers both the number of
// tags the resource already has and the tags this operation would add if it
// succeeded.
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
