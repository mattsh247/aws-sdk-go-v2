// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CampaignState string

// Enum values for CampaignState
const (
	// Campaign is in initialized state
	CampaignStateInitialized CampaignState = "Initialized"
	// Campaign is in running state
	CampaignStateRunning CampaignState = "Running"
	// Campaign is in paused state
	CampaignStatePaused CampaignState = "Paused"
	// Campaign is in stopped state
	CampaignStateStopped CampaignState = "Stopped"
	// Campaign is in failed state
	CampaignStateFailed CampaignState = "Failed"
)

// Values returns all known values for CampaignState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CampaignState) Values() []CampaignState {
	return []CampaignState{
		"Initialized",
		"Running",
		"Paused",
		"Stopped",
		"Failed",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeKms EncryptionType = "KMS"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"KMS",
	}
}

type FailureCode string

// Enum values for FailureCode
const (
	// The request failed to satisfy the constraints specified by the service
	FailureCodeInvalidInput FailureCode = "InvalidInput"
	// Request throttled due to large number of pending dial requests
	FailureCodeRequestThrottled FailureCode = "RequestThrottled"
	// Unexpected error during processing of request
	FailureCodeUnknownError FailureCode = "UnknownError"
)

// Values returns all known values for FailureCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FailureCode) Values() []FailureCode {
	return []FailureCode{
		"InvalidInput",
		"RequestThrottled",
		"UnknownError",
	}
}

type GetCampaignStateBatchFailureCode string

// Enum values for GetCampaignStateBatchFailureCode
const (
	// The specified resource was not found
	GetCampaignStateBatchFailureCodeResourceNotFound GetCampaignStateBatchFailureCode = "ResourceNotFound"
	// Unexpected error during processing of request
	GetCampaignStateBatchFailureCodeUnknownError GetCampaignStateBatchFailureCode = "UnknownError"
)

// Values returns all known values for GetCampaignStateBatchFailureCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (GetCampaignStateBatchFailureCode) Values() []GetCampaignStateBatchFailureCode {
	return []GetCampaignStateBatchFailureCode{
		"ResourceNotFound",
		"UnknownError",
	}
}

type InstanceIdFilterOperator string

// Enum values for InstanceIdFilterOperator
const (
	// Equals operator
	InstanceIdFilterOperatorEq InstanceIdFilterOperator = "Eq"
)

// Values returns all known values for InstanceIdFilterOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (InstanceIdFilterOperator) Values() []InstanceIdFilterOperator {
	return []InstanceIdFilterOperator{
		"Eq",
	}
}

type InstanceOnboardingJobFailureCode string

// Enum values for InstanceOnboardingJobFailureCode
const (
	InstanceOnboardingJobFailureCodeEventBridgeAccessDenied             InstanceOnboardingJobFailureCode = "EVENT_BRIDGE_ACCESS_DENIED"
	InstanceOnboardingJobFailureCodeEventBridgeManagedRuleLimitExceeded InstanceOnboardingJobFailureCode = "EVENT_BRIDGE_MANAGED_RULE_LIMIT_EXCEEDED"
	InstanceOnboardingJobFailureCodeIamAccessDenied                     InstanceOnboardingJobFailureCode = "IAM_ACCESS_DENIED"
	InstanceOnboardingJobFailureCodeKmsAccessDenied                     InstanceOnboardingJobFailureCode = "KMS_ACCESS_DENIED"
	InstanceOnboardingJobFailureCodeKmsKeyNotFound                      InstanceOnboardingJobFailureCode = "KMS_KEY_NOT_FOUND"
	InstanceOnboardingJobFailureCodeInternalFailure                     InstanceOnboardingJobFailureCode = "INTERNAL_FAILURE"
)

// Values returns all known values for InstanceOnboardingJobFailureCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (InstanceOnboardingJobFailureCode) Values() []InstanceOnboardingJobFailureCode {
	return []InstanceOnboardingJobFailureCode{
		"EVENT_BRIDGE_ACCESS_DENIED",
		"EVENT_BRIDGE_MANAGED_RULE_LIMIT_EXCEEDED",
		"IAM_ACCESS_DENIED",
		"KMS_ACCESS_DENIED",
		"KMS_KEY_NOT_FOUND",
		"INTERNAL_FAILURE",
	}
}

type InstanceOnboardingJobStatusCode string

// Enum values for InstanceOnboardingJobStatusCode
const (
	InstanceOnboardingJobStatusCodeInProgress InstanceOnboardingJobStatusCode = "IN_PROGRESS"
	InstanceOnboardingJobStatusCodeSucceeded  InstanceOnboardingJobStatusCode = "SUCCEEDED"
	InstanceOnboardingJobStatusCodeFailed     InstanceOnboardingJobStatusCode = "FAILED"
)

// Values returns all known values for InstanceOnboardingJobStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (InstanceOnboardingJobStatusCode) Values() []InstanceOnboardingJobStatusCode {
	return []InstanceOnboardingJobStatusCode{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}
