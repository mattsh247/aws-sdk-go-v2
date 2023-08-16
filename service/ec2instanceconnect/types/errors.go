// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Either your AWS credentials are not valid or you do not have access to the EC2
// instance.
type AuthException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AuthException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AuthException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AuthException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AuthException"
	}
	return *e.ErrorCodeOverride
}
func (e *AuthException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified instance was not found.
type EC2InstanceNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EC2InstanceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2InstanceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2InstanceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2InstanceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2InstanceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to connect because the instance is not in a valid state. Connecting to a
// stopped or terminated instance is not supported. If the instance is stopped,
// start your instance, and try to connect again.
type EC2InstanceStateInvalidException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EC2InstanceStateInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2InstanceStateInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2InstanceStateInvalidException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2InstanceStateInvalidException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2InstanceStateInvalidException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The instance type is not supported for connecting via the serial console. Only
// Nitro instance types are currently supported.
type EC2InstanceTypeInvalidException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EC2InstanceTypeInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2InstanceTypeInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2InstanceTypeInvalidException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2InstanceTypeInvalidException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2InstanceTypeInvalidException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The instance is currently unavailable. Wait a few minutes and try again.
type EC2InstanceUnavailableException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *EC2InstanceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EC2InstanceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EC2InstanceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EC2InstanceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *EC2InstanceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// One of the parameters is not valid.
type InvalidArgsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidArgsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArgsException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArgsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your account is not authorized to use the EC2 Serial Console. To authorize your
// account, run the EnableSerialConsoleAccess API. For more information, see
// EnableSerialConsoleAccess (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnableSerialConsoleAccess.html)
// in the Amazon EC2 API Reference.
type SerialConsoleAccessDisabledException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SerialConsoleAccessDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SerialConsoleAccessDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SerialConsoleAccessDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SerialConsoleAccessDisabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *SerialConsoleAccessDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The instance currently has 1 active serial console session. Only 1 session is
// supported at a time.
type SerialConsoleSessionLimitExceededException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SerialConsoleSessionLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SerialConsoleSessionLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SerialConsoleSessionLimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SerialConsoleSessionLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *SerialConsoleSessionLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Unable to start a serial console session. Please try again.
type SerialConsoleSessionUnavailableException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SerialConsoleSessionUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SerialConsoleSessionUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SerialConsoleSessionUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SerialConsoleSessionUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *SerialConsoleSessionUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The service encountered an error. Follow the instructions in the error message
// and try again.
type ServiceException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The requests were made too frequently and have been throttled. Wait a while and
// try again. To increase the limit on your request frequency, contact AWS Support.
type ThrottlingException struct {
	Message *string
	
	ErrorCodeOverride *string
	
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
