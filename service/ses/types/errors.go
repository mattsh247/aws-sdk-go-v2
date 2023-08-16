// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Indicates that email sending is disabled for your entire Amazon SES account.
// You can enable or disable email sending for your Amazon SES account using
// UpdateAccountSendingEnabled .
type AccountSendingPausedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AccountSendingPausedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountSendingPausedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountSendingPausedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccountSendingPausedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccountSendingPausedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a resource could not be created because of a naming conflict.
type AlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Name *string
	
	noSmithyDocumentSerde
}

func (e *AlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *AlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the delete operation could not be completed.
type CannotDeleteException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Name *string
	
	noSmithyDocumentSerde
}

func (e *CannotDeleteException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CannotDeleteException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CannotDeleteException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CannotDelete"
	}
	return *e.ErrorCodeOverride
}
func (e *CannotDeleteException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set could not be created because of a naming
// conflict.
type ConfigurationSetAlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	
	noSmithyDocumentSerde
}

func (e *ConfigurationSetAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConfigurationSetAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *ConfigurationSetAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set does not exist.
type ConfigurationSetDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	
	noSmithyDocumentSerde
}

func (e *ConfigurationSetDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConfigurationSetDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *ConfigurationSetDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that email sending is disabled for the configuration set. You can
// enable or disable email sending for a configuration set using
// UpdateConfigurationSetSendingEnabled .
type ConfigurationSetSendingPausedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	
	noSmithyDocumentSerde
}

func (e *ConfigurationSetSendingPausedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConfigurationSetSendingPausedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConfigurationSetSendingPausedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConfigurationSetSendingPausedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConfigurationSetSendingPausedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that custom verification email template provided content is invalid.
type CustomVerificationEmailInvalidContentException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *CustomVerificationEmailInvalidContentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailInvalidContentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailInvalidContentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CustomVerificationEmailInvalidContent"
	}
	return *e.ErrorCodeOverride
}
func (e *CustomVerificationEmailInvalidContentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a custom verification email template with the name you specified
// already exists.
type CustomVerificationEmailTemplateAlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	CustomVerificationEmailTemplateName *string
	
	noSmithyDocumentSerde
}

func (e *CustomVerificationEmailTemplateAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CustomVerificationEmailTemplateAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *CustomVerificationEmailTemplateAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a custom verification email template with the name you specified
// does not exist.
type CustomVerificationEmailTemplateDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	CustomVerificationEmailTemplateName *string
	
	noSmithyDocumentSerde
}

func (e *CustomVerificationEmailTemplateDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CustomVerificationEmailTemplateDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *CustomVerificationEmailTemplateDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the event destination could not be created because of a naming
// conflict.
type EventDestinationAlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	EventDestinationName *string
	
	noSmithyDocumentSerde
}

func (e *EventDestinationAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventDestinationAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventDestinationAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EventDestinationAlreadyExists"
	}
	return *e.ErrorCodeOverride
}
func (e *EventDestinationAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the event destination does not exist.
type EventDestinationDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	EventDestinationName *string
	
	noSmithyDocumentSerde
}

func (e *EventDestinationDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EventDestinationDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EventDestinationDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EventDestinationDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *EventDestinationDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the sender address specified for a custom verification email is
// not verified, and is therefore not eligible to send the custom verification
// email.
type FromEmailAddressNotVerifiedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	FromEmailAddress *string
	
	noSmithyDocumentSerde
}

func (e *FromEmailAddressNotVerifiedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *FromEmailAddressNotVerifiedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *FromEmailAddressNotVerifiedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "FromEmailAddressNotVerified"
	}
	return *e.ErrorCodeOverride
}
func (e *FromEmailAddressNotVerifiedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Amazon CloudWatch destination is invalid. See the error
// message for details.
type InvalidCloudWatchDestinationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	EventDestinationName *string
	
	noSmithyDocumentSerde
}

func (e *InvalidCloudWatchDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidCloudWatchDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidCloudWatchDestinationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidCloudWatchDestination"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidCloudWatchDestinationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set is invalid. See the error message for
// details.
type InvalidConfigurationSetException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidConfigurationSetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationSetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationSetException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidConfigurationSet"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidConfigurationSetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that provided delivery option is invalid.
type InvalidDeliveryOptionsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidDeliveryOptionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDeliveryOptionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDeliveryOptionsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidDeliveryOptions"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidDeliveryOptionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Amazon Kinesis Firehose destination is invalid. See the
// error message for details.
type InvalidFirehoseDestinationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	EventDestinationName *string
	
	noSmithyDocumentSerde
}

func (e *InvalidFirehoseDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidFirehoseDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidFirehoseDestinationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidFirehoseDestination"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidFirehoseDestinationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided Amazon Web Services Lambda function is invalid, or
// that Amazon SES could not execute the provided function, possibly due to
// permissions issues. For information about giving permissions, see the Amazon
// SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html)
// .
type InvalidLambdaFunctionException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	FunctionArn *string
	
	noSmithyDocumentSerde
}

func (e *InvalidLambdaFunctionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLambdaFunctionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLambdaFunctionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidLambdaFunction"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidLambdaFunctionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided policy is invalid. Check the error stack for more
// information about what caused the error.
type InvalidPolicyException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPolicyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidPolicy"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the replacement values you provided is invalid.
// This error may occur when the TemplateData object contains invalid JSON.
type InvalidRenderingParameterException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	TemplateName *string
	
	noSmithyDocumentSerde
}

func (e *InvalidRenderingParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRenderingParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRenderingParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRenderingParameter"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRenderingParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided Amazon S3 bucket or Amazon Web Services KMS
// encryption key is invalid, or that Amazon SES could not publish to the bucket,
// possibly due to permissions issues. For information about giving permissions,
// see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html)
// .
type InvalidS3ConfigurationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Bucket *string
	
	noSmithyDocumentSerde
}

func (e *InvalidS3ConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3ConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3ConfigurationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidS3Configuration"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidS3ConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Amazon Simple Notification Service (Amazon SNS) destination
// is invalid. See the error message for details.
type InvalidSNSDestinationException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	EventDestinationName *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSNSDestinationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSNSDestinationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSNSDestinationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSNSDestination"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSNSDestinationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided Amazon SNS topic is invalid, or that Amazon SES
// could not publish to the topic, possibly due to permissions issues. For
// information about giving permissions, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-permissions.html)
// .
type InvalidSnsTopicException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Topic *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSnsTopicException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnsTopicException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnsTopicException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSnsTopic"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSnsTopicException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the template that you specified could not be rendered. This
// issue may occur when a template refers to a partial that does not exist.
type InvalidTemplateException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	TemplateName *string
	
	noSmithyDocumentSerde
}

func (e *InvalidTemplateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTemplateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTemplateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTemplate"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTemplateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the custom domain to be used for open and click tracking
// redirects is invalid. This error appears most often in the following situations:
//
//   - When the tracking domain you specified is not verified in Amazon SES.
//   - When the tracking domain you specified is not a valid domain or subdomain.
type InvalidTrackingOptionsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidTrackingOptionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTrackingOptionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTrackingOptionsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTrackingOptions"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTrackingOptionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that a resource could not be created because of service limits. For a
// list of Amazon SES limits, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/limits.html)
// .
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
		return "LimitExceeded"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the message could not be sent because Amazon SES could not read
// the MX record required to use the specified MAIL FROM domain. For information
// about editing the custom MAIL FROM domain settings for an identity, see the
// Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from-edit.html)
// .
type MailFromDomainNotVerifiedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *MailFromDomainNotVerifiedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MailFromDomainNotVerifiedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MailFromDomainNotVerifiedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MailFromDomainNotVerifiedException"
	}
	return *e.ErrorCodeOverride
}
func (e *MailFromDomainNotVerifiedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the action failed, and the message could not be sent. Check the
// error stack for more information about what caused the error.
type MessageRejected struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *MessageRejected) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MessageRejected) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MessageRejected) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MessageRejected"
	}
	return *e.ErrorCodeOverride
}
func (e *MessageRejected) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that one or more of the replacement values for the specified template
// was not specified. Ensure that the TemplateData object contains references to
// all of the replacement tags in the specified template.
type MissingRenderingAttributeException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	TemplateName *string
	
	noSmithyDocumentSerde
}

func (e *MissingRenderingAttributeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MissingRenderingAttributeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MissingRenderingAttributeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MissingRenderingAttribute"
	}
	return *e.ErrorCodeOverride
}
func (e *MissingRenderingAttributeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the account has not been granted production access.
type ProductionAccessNotGrantedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ProductionAccessNotGrantedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProductionAccessNotGrantedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProductionAccessNotGrantedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ProductionAccessNotGranted"
	}
	return *e.ErrorCodeOverride
}
func (e *ProductionAccessNotGrantedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided receipt rule does not exist.
type RuleDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Name *string
	
	noSmithyDocumentSerde
}

func (e *RuleDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RuleDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *RuleDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the provided receipt rule set does not exist.
type RuleSetDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	Name *string
	
	noSmithyDocumentSerde
}

func (e *RuleSetDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleSetDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleSetDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RuleSetDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *RuleSetDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the Template object you specified does not exist in your Amazon
// SES account.
type TemplateDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	TemplateName *string
	
	noSmithyDocumentSerde
}

func (e *TemplateDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TemplateDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TemplateDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TemplateDoesNotExist"
	}
	return *e.ErrorCodeOverride
}
func (e *TemplateDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the configuration set you specified already contains a
// TrackingOptions object.
type TrackingOptionsAlreadyExistsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	
	noSmithyDocumentSerde
}

func (e *TrackingOptionsAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrackingOptionsAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrackingOptionsAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TrackingOptionsAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *TrackingOptionsAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Indicates that the TrackingOptions object you specified does not exist.
type TrackingOptionsDoesNotExistException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	ConfigurationSetName *string
	
	noSmithyDocumentSerde
}

func (e *TrackingOptionsDoesNotExistException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrackingOptionsDoesNotExistException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrackingOptionsDoesNotExistException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TrackingOptionsDoesNotExistException"
	}
	return *e.ErrorCodeOverride
}
func (e *TrackingOptionsDoesNotExistException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
