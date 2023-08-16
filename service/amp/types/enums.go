// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AlertManagerDefinitionStatusCode string

// Enum values for AlertManagerDefinitionStatusCode
const (
	// Definition is being created. Update/Deletion is disallowed until definition is
	// ACTIVE and workspace status is ACTIVE.
	AlertManagerDefinitionStatusCodeCreating AlertManagerDefinitionStatusCode = "CREATING"
	// Definition has been created/updated. Update/Deletion is disallowed until
	// definition is ACTIVE and workspace status is ACTIVE.
	AlertManagerDefinitionStatusCodeActive AlertManagerDefinitionStatusCode = "ACTIVE"
	// Definition is being updated. Update/Deletion is disallowed until definition is
	// ACTIVE and workspace status is ACTIVE.
	AlertManagerDefinitionStatusCodeUpdating AlertManagerDefinitionStatusCode = "UPDATING"
	// Definition is being deleting. Update/Deletion is disallowed until definition is
	// ACTIVE and workspace status is ACTIVE.
	AlertManagerDefinitionStatusCodeDeleting AlertManagerDefinitionStatusCode = "DELETING"
	// Definition creation failed.
	AlertManagerDefinitionStatusCodeCreationFailed AlertManagerDefinitionStatusCode = "CREATION_FAILED"
	// Definition update failed.
	AlertManagerDefinitionStatusCodeUpdateFailed AlertManagerDefinitionStatusCode = "UPDATE_FAILED"
)

// Values returns all known values for AlertManagerDefinitionStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AlertManagerDefinitionStatusCode) Values() []AlertManagerDefinitionStatusCode {
	return []AlertManagerDefinitionStatusCode{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATION_FAILED",
		"UPDATE_FAILED",
	}
}

type LoggingConfigurationStatusCode string

// Enum values for LoggingConfigurationStatusCode
const (
	// Logging configuration is being created. Update/Deletion is disallowed until
	// logging configuration is ACTIVE and workspace status is ACTIVE.
	LoggingConfigurationStatusCodeCreating LoggingConfigurationStatusCode = "CREATING"
	// Logging configuration has been created/updated. Update/Deletion is disallowed
	// until logging configuration is ACTIVE and workspace status is ACTIVE.
	LoggingConfigurationStatusCodeActive LoggingConfigurationStatusCode = "ACTIVE"
	// Logging configuration is being updated. Update/Deletion is disallowed until
	// logging configuration is ACTIVE and workspace status is ACTIVE.
	LoggingConfigurationStatusCodeUpdating LoggingConfigurationStatusCode = "UPDATING"
	// Logging configuration is being deleting. Update/Deletion is disallowed until
	// logging configuration is ACTIVE and workspace status is ACTIVE.
	LoggingConfigurationStatusCodeDeleting LoggingConfigurationStatusCode = "DELETING"
	// Logging configuration creation failed.
	LoggingConfigurationStatusCodeCreationFailed LoggingConfigurationStatusCode = "CREATION_FAILED"
	// Logging configuration update failed.
	LoggingConfigurationStatusCodeUpdateFailed LoggingConfigurationStatusCode = "UPDATE_FAILED"
)

// Values returns all known values for LoggingConfigurationStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (LoggingConfigurationStatusCode) Values() []LoggingConfigurationStatusCode {
	return []LoggingConfigurationStatusCode{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATION_FAILED",
		"UPDATE_FAILED",
	}
}

type RuleGroupsNamespaceStatusCode string

// Enum values for RuleGroupsNamespaceStatusCode
const (
	// Namespace is being created. Update/Deletion is disallowed until namespace is
	// ACTIVE and workspace status is ACTIVE.
	RuleGroupsNamespaceStatusCodeCreating RuleGroupsNamespaceStatusCode = "CREATING"
	// Namespace has been created/updated. Update/Deletion is disallowed until
	// namespace is ACTIVE and workspace status is ACTIVE.
	RuleGroupsNamespaceStatusCodeActive RuleGroupsNamespaceStatusCode = "ACTIVE"
	// Namespace is being updated. Update/Deletion is disallowed until namespace is
	// ACTIVE and workspace status is ACTIVE.
	RuleGroupsNamespaceStatusCodeUpdating RuleGroupsNamespaceStatusCode = "UPDATING"
	// Namespace is being deleting. Update/Deletion is disallowed until namespace is
	// ACTIVE and workspace status is ACTIVE.
	RuleGroupsNamespaceStatusCodeDeleting RuleGroupsNamespaceStatusCode = "DELETING"
	// Namespace creation failed.
	RuleGroupsNamespaceStatusCodeCreationFailed RuleGroupsNamespaceStatusCode = "CREATION_FAILED"
	// Namespace update failed.
	RuleGroupsNamespaceStatusCodeUpdateFailed RuleGroupsNamespaceStatusCode = "UPDATE_FAILED"
)

// Values returns all known values for RuleGroupsNamespaceStatusCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RuleGroupsNamespaceStatusCode) Values() []RuleGroupsNamespaceStatusCode {
	return []RuleGroupsNamespaceStatusCode{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATION_FAILED",
		"UPDATE_FAILED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation ValidationExceptionReason = "UNKNOWN_OPERATION"
	ValidationExceptionReasonCannotParse ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReasonOther ValidationExceptionReason = "OTHER"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"UNKNOWN_OPERATION",
		"CANNOT_PARSE",
		"FIELD_VALIDATION_FAILED",
		"OTHER",
	}
}

type WorkspaceStatusCode string

// Enum values for WorkspaceStatusCode
const (
	// Workspace is being created. Deletion is disallowed until status is ACTIVE.
	WorkspaceStatusCodeCreating WorkspaceStatusCode = "CREATING"
	// Workspace has been created and is usable.
	WorkspaceStatusCodeActive WorkspaceStatusCode = "ACTIVE"
	// Workspace is being updated. Updates are allowed only when status is ACTIVE.
	WorkspaceStatusCodeUpdating WorkspaceStatusCode = "UPDATING"
	// Workspace is being deleted. Deletions are allowed only when status is ACTIVE.
	WorkspaceStatusCodeDeleting WorkspaceStatusCode = "DELETING"
	// Workspace creation failed. Refer to WorkspaceStatus.failureReason for more
	// details.
	WorkspaceStatusCodeCreationFailed WorkspaceStatusCode = "CREATION_FAILED"
)

// Values returns all known values for WorkspaceStatusCode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WorkspaceStatusCode) Values() []WorkspaceStatusCode {
	return []WorkspaceStatusCode{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATION_FAILED",
	}
}
