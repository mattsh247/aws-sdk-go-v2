// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AuthenticationDecision string

// Enum values for AuthenticationDecision
const (
	AuthenticationDecisionAccept AuthenticationDecision = "ACCEPT"
	AuthenticationDecisionReject AuthenticationDecision = "REJECT"
	AuthenticationDecisionNotEnoughSpeech AuthenticationDecision = "NOT_ENOUGH_SPEECH"
	AuthenticationDecisionSpeakerNotEnrolled AuthenticationDecision = "SPEAKER_NOT_ENROLLED"
	AuthenticationDecisionSpeakerOptedOut AuthenticationDecision = "SPEAKER_OPTED_OUT"
	AuthenticationDecisionSpeakerIdNotProvided AuthenticationDecision = "SPEAKER_ID_NOT_PROVIDED"
	AuthenticationDecisionSpeakerExpired AuthenticationDecision = "SPEAKER_EXPIRED"
)

// Values returns all known values for AuthenticationDecision. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticationDecision) Values() []AuthenticationDecision {
	return []AuthenticationDecision{
		"ACCEPT",
		"REJECT",
		"NOT_ENOUGH_SPEECH",
		"SPEAKER_NOT_ENROLLED",
		"SPEAKER_OPTED_OUT",
		"SPEAKER_ID_NOT_PROVIDED",
		"SPEAKER_EXPIRED",
	}
}

type ConflictType string

// Enum values for ConflictType
const (
	ConflictTypeAnotherActiveStream ConflictType = "ANOTHER_ACTIVE_STREAM"
	ConflictTypeDomainNotActive ConflictType = "DOMAIN_NOT_ACTIVE"
	ConflictTypeCannotChangeSpeakerAfterEnrollment ConflictType = "CANNOT_CHANGE_SPEAKER_AFTER_ENROLLMENT"
	ConflictTypeEnrollmentAlreadyExists ConflictType = "ENROLLMENT_ALREADY_EXISTS"
	ConflictTypeSpeakerNotSet ConflictType = "SPEAKER_NOT_SET"
	ConflictTypeSpeakerOptedOut ConflictType = "SPEAKER_OPTED_OUT"
	ConflictTypeConcurrentChanges ConflictType = "CONCURRENT_CHANGES"
	ConflictTypeDomainLockedFromEncryptionUpdates ConflictType = "DOMAIN_LOCKED_FROM_ENCRYPTION_UPDATES"
	ConflictTypeCannotDeleteNonEmptyWatchlist ConflictType = "CANNOT_DELETE_NON_EMPTY_WATCHLIST"
	ConflictTypeFraudsterMustBelongToAtLeastOneWatchlist ConflictType = "FRAUDSTER_MUST_BELONG_TO_AT_LEAST_ONE_WATCHLIST"
)

// Values returns all known values for ConflictType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConflictType) Values() []ConflictType {
	return []ConflictType{
		"ANOTHER_ACTIVE_STREAM",
		"DOMAIN_NOT_ACTIVE",
		"CANNOT_CHANGE_SPEAKER_AFTER_ENROLLMENT",
		"ENROLLMENT_ALREADY_EXISTS",
		"SPEAKER_NOT_SET",
		"SPEAKER_OPTED_OUT",
		"CONCURRENT_CHANGES",
		"DOMAIN_LOCKED_FROM_ENCRYPTION_UPDATES",
		"CANNOT_DELETE_NON_EMPTY_WATCHLIST",
		"FRAUDSTER_MUST_BELONG_TO_AT_LEAST_ONE_WATCHLIST",
	}
}

type DomainStatus string

// Enum values for DomainStatus
const (
	DomainStatusActive DomainStatus = "ACTIVE"
	DomainStatusPending DomainStatus = "PENDING"
	DomainStatusSuspended DomainStatus = "SUSPENDED"
)

// Values returns all known values for DomainStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"ACTIVE",
		"PENDING",
		"SUSPENDED",
	}
}

type DuplicateRegistrationAction string

// Enum values for DuplicateRegistrationAction
const (
	DuplicateRegistrationActionSkip DuplicateRegistrationAction = "SKIP"
	DuplicateRegistrationActionRegisterAsNew DuplicateRegistrationAction = "REGISTER_AS_NEW"
)

// Values returns all known values for DuplicateRegistrationAction. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DuplicateRegistrationAction) Values() []DuplicateRegistrationAction {
	return []DuplicateRegistrationAction{
		"SKIP",
		"REGISTER_AS_NEW",
	}
}

type ExistingEnrollmentAction string

// Enum values for ExistingEnrollmentAction
const (
	ExistingEnrollmentActionSkip ExistingEnrollmentAction = "SKIP"
	ExistingEnrollmentActionOverwrite ExistingEnrollmentAction = "OVERWRITE"
)

// Values returns all known values for ExistingEnrollmentAction. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ExistingEnrollmentAction) Values() []ExistingEnrollmentAction {
	return []ExistingEnrollmentAction{
		"SKIP",
		"OVERWRITE",
	}
}

type FraudDetectionAction string

// Enum values for FraudDetectionAction
const (
	FraudDetectionActionIgnore FraudDetectionAction = "IGNORE"
	FraudDetectionActionFail FraudDetectionAction = "FAIL"
)

// Values returns all known values for FraudDetectionAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FraudDetectionAction) Values() []FraudDetectionAction {
	return []FraudDetectionAction{
		"IGNORE",
		"FAIL",
	}
}

type FraudDetectionDecision string

// Enum values for FraudDetectionDecision
const (
	FraudDetectionDecisionHighRisk FraudDetectionDecision = "HIGH_RISK"
	FraudDetectionDecisionLowRisk FraudDetectionDecision = "LOW_RISK"
	FraudDetectionDecisionNotEnoughSpeech FraudDetectionDecision = "NOT_ENOUGH_SPEECH"
)

// Values returns all known values for FraudDetectionDecision. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FraudDetectionDecision) Values() []FraudDetectionDecision {
	return []FraudDetectionDecision{
		"HIGH_RISK",
		"LOW_RISK",
		"NOT_ENOUGH_SPEECH",
	}
}

type FraudDetectionReason string

// Enum values for FraudDetectionReason
const (
	FraudDetectionReasonKnownFraudster FraudDetectionReason = "KNOWN_FRAUDSTER"
	FraudDetectionReasonVoiceSpoofing FraudDetectionReason = "VOICE_SPOOFING"
)

// Values returns all known values for FraudDetectionReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FraudDetectionReason) Values() []FraudDetectionReason {
	return []FraudDetectionReason{
		"KNOWN_FRAUDSTER",
		"VOICE_SPOOFING",
	}
}

type FraudsterRegistrationJobStatus string

// Enum values for FraudsterRegistrationJobStatus
const (
	FraudsterRegistrationJobStatusSubmitted FraudsterRegistrationJobStatus = "SUBMITTED"
	FraudsterRegistrationJobStatusInProgress FraudsterRegistrationJobStatus = "IN_PROGRESS"
	FraudsterRegistrationJobStatusCompleted FraudsterRegistrationJobStatus = "COMPLETED"
	FraudsterRegistrationJobStatusCompletedWithErrors FraudsterRegistrationJobStatus = "COMPLETED_WITH_ERRORS"
	FraudsterRegistrationJobStatusFailed FraudsterRegistrationJobStatus = "FAILED"
)

// Values returns all known values for FraudsterRegistrationJobStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FraudsterRegistrationJobStatus) Values() []FraudsterRegistrationJobStatus {
	return []FraudsterRegistrationJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED",
		"COMPLETED_WITH_ERRORS",
		"FAILED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeBatchJob ResourceType = "BATCH_JOB"
	ResourceTypeComplianceConsent ResourceType = "COMPLIANCE_CONSENT"
	ResourceTypeDomain ResourceType = "DOMAIN"
	ResourceTypeFraudster ResourceType = "FRAUDSTER"
	ResourceTypeSession ResourceType = "SESSION"
	ResourceTypeSpeaker ResourceType = "SPEAKER"
	ResourceTypeWatchlist ResourceType = "WATCHLIST"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"BATCH_JOB",
		"COMPLIANCE_CONSENT",
		"DOMAIN",
		"FRAUDSTER",
		"SESSION",
		"SPEAKER",
		"WATCHLIST",
	}
}

type ServerSideEncryptionUpdateStatus string

// Enum values for ServerSideEncryptionUpdateStatus
const (
	ServerSideEncryptionUpdateStatusInProgress ServerSideEncryptionUpdateStatus = "IN_PROGRESS"
	ServerSideEncryptionUpdateStatusCompleted ServerSideEncryptionUpdateStatus = "COMPLETED"
	ServerSideEncryptionUpdateStatusFailed ServerSideEncryptionUpdateStatus = "FAILED"
)

// Values returns all known values for ServerSideEncryptionUpdateStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ServerSideEncryptionUpdateStatus) Values() []ServerSideEncryptionUpdateStatus {
	return []ServerSideEncryptionUpdateStatus{
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

type SpeakerEnrollmentJobStatus string

// Enum values for SpeakerEnrollmentJobStatus
const (
	SpeakerEnrollmentJobStatusSubmitted SpeakerEnrollmentJobStatus = "SUBMITTED"
	SpeakerEnrollmentJobStatusInProgress SpeakerEnrollmentJobStatus = "IN_PROGRESS"
	SpeakerEnrollmentJobStatusCompleted SpeakerEnrollmentJobStatus = "COMPLETED"
	SpeakerEnrollmentJobStatusCompletedWithErrors SpeakerEnrollmentJobStatus = "COMPLETED_WITH_ERRORS"
	SpeakerEnrollmentJobStatusFailed SpeakerEnrollmentJobStatus = "FAILED"
)

// Values returns all known values for SpeakerEnrollmentJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SpeakerEnrollmentJobStatus) Values() []SpeakerEnrollmentJobStatus {
	return []SpeakerEnrollmentJobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED",
		"COMPLETED_WITH_ERRORS",
		"FAILED",
	}
}

type SpeakerStatus string

// Enum values for SpeakerStatus
const (
	SpeakerStatusEnrolled SpeakerStatus = "ENROLLED"
	SpeakerStatusExpired SpeakerStatus = "EXPIRED"
	SpeakerStatusOptedOut SpeakerStatus = "OPTED_OUT"
	SpeakerStatusPending SpeakerStatus = "PENDING"
)

// Values returns all known values for SpeakerStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SpeakerStatus) Values() []SpeakerStatus {
	return []SpeakerStatus{
		"ENROLLED",
		"EXPIRED",
		"OPTED_OUT",
		"PENDING",
	}
}

type StreamingStatus string

// Enum values for StreamingStatus
const (
	StreamingStatusPendingConfiguration StreamingStatus = "PENDING_CONFIGURATION"
	StreamingStatusOngoing StreamingStatus = "ONGOING"
	StreamingStatusEnded StreamingStatus = "ENDED"
)

// Values returns all known values for StreamingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StreamingStatus) Values() []StreamingStatus {
	return []StreamingStatus{
		"PENDING_CONFIGURATION",
		"ONGOING",
		"ENDED",
	}
}
