// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AccessControlRuleEffect string

// Enum values for AccessControlRuleEffect
const (
	AccessControlRuleEffectAllow AccessControlRuleEffect = "ALLOW"
	AccessControlRuleEffectDeny AccessControlRuleEffect = "DENY"
)

// Values returns all known values for AccessControlRuleEffect. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessControlRuleEffect) Values() []AccessControlRuleEffect {
	return []AccessControlRuleEffect{
		"ALLOW",
		"DENY",
	}
}

type AccessEffect string

// Enum values for AccessEffect
const (
	AccessEffectAllow AccessEffect = "ALLOW"
	AccessEffectDeny AccessEffect = "DENY"
)

// Values returns all known values for AccessEffect. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessEffect) Values() []AccessEffect {
	return []AccessEffect{
		"ALLOW",
		"DENY",
	}
}

type AvailabilityProviderType string

// Enum values for AvailabilityProviderType
const (
	AvailabilityProviderTypeEws AvailabilityProviderType = "EWS"
	AvailabilityProviderTypeLambda AvailabilityProviderType = "LAMBDA"
)

// Values returns all known values for AvailabilityProviderType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AvailabilityProviderType) Values() []AvailabilityProviderType {
	return []AvailabilityProviderType{
		"EWS",
		"LAMBDA",
	}
}

type DnsRecordVerificationStatus string

// Enum values for DnsRecordVerificationStatus
const (
	DnsRecordVerificationStatusPending DnsRecordVerificationStatus = "PENDING"
	DnsRecordVerificationStatusVerified DnsRecordVerificationStatus = "VERIFIED"
	DnsRecordVerificationStatusFailed DnsRecordVerificationStatus = "FAILED"
)

// Values returns all known values for DnsRecordVerificationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DnsRecordVerificationStatus) Values() []DnsRecordVerificationStatus {
	return []DnsRecordVerificationStatus{
		"PENDING",
		"VERIFIED",
		"FAILED",
	}
}

type EntityState string

// Enum values for EntityState
const (
	EntityStateEnabled EntityState = "ENABLED"
	EntityStateDisabled EntityState = "DISABLED"
	EntityStateDeleted EntityState = "DELETED"
)

// Values returns all known values for EntityState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EntityState) Values() []EntityState {
	return []EntityState{
		"ENABLED",
		"DISABLED",
		"DELETED",
	}
}

type FolderName string

// Enum values for FolderName
const (
	FolderNameInbox FolderName = "INBOX"
	FolderNameDeletedItems FolderName = "DELETED_ITEMS"
	FolderNameSentItems FolderName = "SENT_ITEMS"
	FolderNameDrafts FolderName = "DRAFTS"
	FolderNameJunkEmail FolderName = "JUNK_EMAIL"
)

// Values returns all known values for FolderName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FolderName) Values() []FolderName {
	return []FolderName{
		"INBOX",
		"DELETED_ITEMS",
		"SENT_ITEMS",
		"DRAFTS",
		"JUNK_EMAIL",
	}
}

type ImpersonationRoleType string

// Enum values for ImpersonationRoleType
const (
	ImpersonationRoleTypeFullAccess ImpersonationRoleType = "FULL_ACCESS"
	ImpersonationRoleTypeReadOnly ImpersonationRoleType = "READ_ONLY"
)

// Values returns all known values for ImpersonationRoleType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImpersonationRoleType) Values() []ImpersonationRoleType {
	return []ImpersonationRoleType{
		"FULL_ACCESS",
		"READ_ONLY",
	}
}

type MailboxExportJobState string

// Enum values for MailboxExportJobState
const (
	MailboxExportJobStateRunning MailboxExportJobState = "RUNNING"
	MailboxExportJobStateCompleted MailboxExportJobState = "COMPLETED"
	MailboxExportJobStateFailed MailboxExportJobState = "FAILED"
	MailboxExportJobStateCancelled MailboxExportJobState = "CANCELLED"
)

// Values returns all known values for MailboxExportJobState. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MailboxExportJobState) Values() []MailboxExportJobState {
	return []MailboxExportJobState{
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"CANCELLED",
	}
}

type MemberType string

// Enum values for MemberType
const (
	MemberTypeGroup MemberType = "GROUP"
	MemberTypeUser MemberType = "USER"
)

// Values returns all known values for MemberType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MemberType) Values() []MemberType {
	return []MemberType{
		"GROUP",
		"USER",
	}
}

type MobileDeviceAccessRuleEffect string

// Enum values for MobileDeviceAccessRuleEffect
const (
	MobileDeviceAccessRuleEffectAllow MobileDeviceAccessRuleEffect = "ALLOW"
	MobileDeviceAccessRuleEffectDeny MobileDeviceAccessRuleEffect = "DENY"
)

// Values returns all known values for MobileDeviceAccessRuleEffect. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (MobileDeviceAccessRuleEffect) Values() []MobileDeviceAccessRuleEffect {
	return []MobileDeviceAccessRuleEffect{
		"ALLOW",
		"DENY",
	}
}

type PermissionType string

// Enum values for PermissionType
const (
	PermissionTypeFullAccess PermissionType = "FULL_ACCESS"
	PermissionTypeSendAs PermissionType = "SEND_AS"
	PermissionTypeSendOnBehalf PermissionType = "SEND_ON_BEHALF"
)

// Values returns all known values for PermissionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionType) Values() []PermissionType {
	return []PermissionType{
		"FULL_ACCESS",
		"SEND_AS",
		"SEND_ON_BEHALF",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeRoom ResourceType = "ROOM"
	ResourceTypeEquipment ResourceType = "EQUIPMENT"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"ROOM",
		"EQUIPMENT",
	}
}

type RetentionAction string

// Enum values for RetentionAction
const (
	RetentionActionNone RetentionAction = "NONE"
	RetentionActionDelete RetentionAction = "DELETE"
	RetentionActionPermanentlyDelete RetentionAction = "PERMANENTLY_DELETE"
)

// Values returns all known values for RetentionAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RetentionAction) Values() []RetentionAction {
	return []RetentionAction{
		"NONE",
		"DELETE",
		"PERMANENTLY_DELETE",
	}
}

type UserRole string

// Enum values for UserRole
const (
	UserRoleUser UserRole = "USER"
	UserRoleResource UserRole = "RESOURCE"
	UserRoleSystemUser UserRole = "SYSTEM_USER"
)

// Values returns all known values for UserRole. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserRole) Values() []UserRole {
	return []UserRole{
		"USER",
		"RESOURCE",
		"SYSTEM_USER",
	}
}
