// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type BulkPublishStatus string

// Enum values for BulkPublishStatus
const (
	BulkPublishStatusNotStarted BulkPublishStatus = "NOT_STARTED"
	BulkPublishStatusInProgress BulkPublishStatus = "IN_PROGRESS"
	BulkPublishStatusFailed BulkPublishStatus = "FAILED"
	BulkPublishStatusSucceeded BulkPublishStatus = "SUCCEEDED"
)

// Values returns all known values for BulkPublishStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BulkPublishStatus) Values() []BulkPublishStatus {
	return []BulkPublishStatus{
		"NOT_STARTED",
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

type Operation string

// Enum values for Operation
const (
	OperationReplace Operation = "replace"
	OperationRemove Operation = "remove"
)

// Values returns all known values for Operation. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Operation) Values() []Operation {
	return []Operation{
		"replace",
		"remove",
	}
}

type Platform string

// Enum values for Platform
const (
	PlatformApns Platform = "APNS"
	PlatformApnsSandbox Platform = "APNS_SANDBOX"
	PlatformGcm Platform = "GCM"
	PlatformAdm Platform = "ADM"
)

// Values returns all known values for Platform. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Platform) Values() []Platform {
	return []Platform{
		"APNS",
		"APNS_SANDBOX",
		"GCM",
		"ADM",
	}
}

type StreamingStatus string

// Enum values for StreamingStatus
const (
	StreamingStatusEnabled StreamingStatus = "ENABLED"
	StreamingStatusDisabled StreamingStatus = "DISABLED"
)

// Values returns all known values for StreamingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StreamingStatus) Values() []StreamingStatus {
	return []StreamingStatus{
		"ENABLED",
		"DISABLED",
	}
}
