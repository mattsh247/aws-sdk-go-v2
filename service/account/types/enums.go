// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AlternateContactType string

// Enum values for AlternateContactType
const (
	AlternateContactTypeBilling AlternateContactType = "BILLING"
	AlternateContactTypeOperations AlternateContactType = "OPERATIONS"
	AlternateContactTypeSecurity AlternateContactType = "SECURITY"
)

// Values returns all known values for AlternateContactType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlternateContactType) Values() []AlternateContactType {
	return []AlternateContactType{
		"BILLING",
		"OPERATIONS",
		"SECURITY",
	}
}

type RegionOptStatus string

// Enum values for RegionOptStatus
const (
	RegionOptStatusEnabled RegionOptStatus = "ENABLED"
	RegionOptStatusEnabling RegionOptStatus = "ENABLING"
	RegionOptStatusDisabling RegionOptStatus = "DISABLING"
	RegionOptStatusDisabled RegionOptStatus = "DISABLED"
	RegionOptStatusEnabledByDefault RegionOptStatus = "ENABLED_BY_DEFAULT"
)

// Values returns all known values for RegionOptStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RegionOptStatus) Values() []RegionOptStatus {
	return []RegionOptStatus{
		"ENABLED",
		"ENABLING",
		"DISABLING",
		"DISABLED",
		"ENABLED_BY_DEFAULT",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonInvalidRegionOptTarget ValidationExceptionReason = "invalidRegionOptTarget"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"invalidRegionOptTarget",
		"fieldValidationFailed",
	}
}
