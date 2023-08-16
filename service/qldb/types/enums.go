// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EncryptionStatus string

// Enum values for EncryptionStatus
const (
	EncryptionStatusEnabled            EncryptionStatus = "ENABLED"
	EncryptionStatusUpdating           EncryptionStatus = "UPDATING"
	EncryptionStatusKmsKeyInaccessible EncryptionStatus = "KMS_KEY_INACCESSIBLE"
)

// Values returns all known values for EncryptionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionStatus) Values() []EncryptionStatus {
	return []EncryptionStatus{
		"ENABLED",
		"UPDATING",
		"KMS_KEY_INACCESSIBLE",
	}
}

type ErrorCause string

// Enum values for ErrorCause
const (
	ErrorCauseKinesisStreamNotFound ErrorCause = "KINESIS_STREAM_NOT_FOUND"
	ErrorCauseIamPermissionRevoked  ErrorCause = "IAM_PERMISSION_REVOKED"
)

// Values returns all known values for ErrorCause. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ErrorCause) Values() []ErrorCause {
	return []ErrorCause{
		"KINESIS_STREAM_NOT_FOUND",
		"IAM_PERMISSION_REVOKED",
	}
}

type ExportStatus string

// Enum values for ExportStatus
const (
	ExportStatusInProgress ExportStatus = "IN_PROGRESS"
	ExportStatusCompleted  ExportStatus = "COMPLETED"
	ExportStatusCancelled  ExportStatus = "CANCELLED"
)

// Values returns all known values for ExportStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportStatus) Values() []ExportStatus {
	return []ExportStatus{
		"IN_PROGRESS",
		"COMPLETED",
		"CANCELLED",
	}
}

type LedgerState string

// Enum values for LedgerState
const (
	LedgerStateCreating LedgerState = "CREATING"
	LedgerStateActive   LedgerState = "ACTIVE"
	LedgerStateDeleting LedgerState = "DELETING"
	LedgerStateDeleted  LedgerState = "DELETED"
)

// Values returns all known values for LedgerState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LedgerState) Values() []LedgerState {
	return []LedgerState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatIonBinary OutputFormat = "ION_BINARY"
	OutputFormatIonText   OutputFormat = "ION_TEXT"
	OutputFormatJson      OutputFormat = "JSON"
)

// Values returns all known values for OutputFormat. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"ION_BINARY",
		"ION_TEXT",
		"JSON",
	}
}

type PermissionsMode string

// Enum values for PermissionsMode
const (
	PermissionsModeAllowAll PermissionsMode = "ALLOW_ALL"
	PermissionsModeStandard PermissionsMode = "STANDARD"
)

// Values returns all known values for PermissionsMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionsMode) Values() []PermissionsMode {
	return []PermissionsMode{
		"ALLOW_ALL",
		"STANDARD",
	}
}

type S3ObjectEncryptionType string

// Enum values for S3ObjectEncryptionType
const (
	S3ObjectEncryptionTypeSseKms       S3ObjectEncryptionType = "SSE_KMS"
	S3ObjectEncryptionTypeSseS3        S3ObjectEncryptionType = "SSE_S3"
	S3ObjectEncryptionTypeNoEncryption S3ObjectEncryptionType = "NO_ENCRYPTION"
)

// Values returns all known values for S3ObjectEncryptionType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (S3ObjectEncryptionType) Values() []S3ObjectEncryptionType {
	return []S3ObjectEncryptionType{
		"SSE_KMS",
		"SSE_S3",
		"NO_ENCRYPTION",
	}
}

type StreamStatus string

// Enum values for StreamStatus
const (
	StreamStatusActive    StreamStatus = "ACTIVE"
	StreamStatusCompleted StreamStatus = "COMPLETED"
	StreamStatusCanceled  StreamStatus = "CANCELED"
	StreamStatusFailed    StreamStatus = "FAILED"
	StreamStatusImpaired  StreamStatus = "IMPAIRED"
)

// Values returns all known values for StreamStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StreamStatus) Values() []StreamStatus {
	return []StreamStatus{
		"ACTIVE",
		"COMPLETED",
		"CANCELED",
		"FAILED",
		"IMPAIRED",
	}
}
