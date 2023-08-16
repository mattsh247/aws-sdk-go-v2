// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type BlockerStatus string

// Enum values for BlockerStatus
const (
	BlockerStatusActive BlockerStatus = "ACTIVE"
	BlockerStatusResolved BlockerStatus = "RESOLVED"
)

// Values returns all known values for BlockerStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BlockerStatus) Values() []BlockerStatus {
	return []BlockerStatus{
		"ACTIVE",
		"RESOLVED",
	}
}

type BlockerType string

// Enum values for BlockerType
const (
	BlockerTypeAutomated BlockerType = "AUTOMATED"
)

// Values returns all known values for BlockerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BlockerType) Values() []BlockerType {
	return []BlockerType{
		"AUTOMATED",
	}
}

type ComponentDeploymentUpdateType string

// Enum values for ComponentDeploymentUpdateType
const (
	ComponentDeploymentUpdateTypeNone ComponentDeploymentUpdateType = "NONE"
	ComponentDeploymentUpdateTypeCurrentVersion ComponentDeploymentUpdateType = "CURRENT_VERSION"
)

// Values returns all known values for ComponentDeploymentUpdateType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ComponentDeploymentUpdateType) Values() []ComponentDeploymentUpdateType {
	return []ComponentDeploymentUpdateType{
		"NONE",
		"CURRENT_VERSION",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusInProgress DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusFailed DeploymentStatus = "FAILED"
	DeploymentStatusSucceeded DeploymentStatus = "SUCCEEDED"
	DeploymentStatusDeleteInProgress DeploymentStatus = "DELETE_IN_PROGRESS"
	DeploymentStatusDeleteFailed DeploymentStatus = "DELETE_FAILED"
	DeploymentStatusDeleteComplete DeploymentStatus = "DELETE_COMPLETE"
	DeploymentStatusCancelling DeploymentStatus = "CANCELLING"
	DeploymentStatusCancelled DeploymentStatus = "CANCELLED"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETE_COMPLETE",
		"CANCELLING",
		"CANCELLED",
	}
}

type DeploymentTargetResourceType string

// Enum values for DeploymentTargetResourceType
const (
	DeploymentTargetResourceTypeEnvironment DeploymentTargetResourceType = "ENVIRONMENT"
	DeploymentTargetResourceTypeServicePipeline DeploymentTargetResourceType = "SERVICE_PIPELINE"
	DeploymentTargetResourceTypeServiceInstance DeploymentTargetResourceType = "SERVICE_INSTANCE"
	DeploymentTargetResourceTypeComponent DeploymentTargetResourceType = "COMPONENT"
)

// Values returns all known values for DeploymentTargetResourceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DeploymentTargetResourceType) Values() []DeploymentTargetResourceType {
	return []DeploymentTargetResourceType{
		"ENVIRONMENT",
		"SERVICE_PIPELINE",
		"SERVICE_INSTANCE",
		"COMPONENT",
	}
}

type DeploymentUpdateType string

// Enum values for DeploymentUpdateType
const (
	DeploymentUpdateTypeNone DeploymentUpdateType = "NONE"
	DeploymentUpdateTypeCurrentVersion DeploymentUpdateType = "CURRENT_VERSION"
	DeploymentUpdateTypeMinorVersion DeploymentUpdateType = "MINOR_VERSION"
	DeploymentUpdateTypeMajorVersion DeploymentUpdateType = "MAJOR_VERSION"
)

// Values returns all known values for DeploymentUpdateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentUpdateType) Values() []DeploymentUpdateType {
	return []DeploymentUpdateType{
		"NONE",
		"CURRENT_VERSION",
		"MINOR_VERSION",
		"MAJOR_VERSION",
	}
}

type EnvironmentAccountConnectionRequesterAccountType string

// Enum values for EnvironmentAccountConnectionRequesterAccountType
const (
	EnvironmentAccountConnectionRequesterAccountTypeManagementAccount EnvironmentAccountConnectionRequesterAccountType = "MANAGEMENT_ACCOUNT"
	EnvironmentAccountConnectionRequesterAccountTypeEnvironmentAccount EnvironmentAccountConnectionRequesterAccountType = "ENVIRONMENT_ACCOUNT"
)

// Values returns all known values for
// EnvironmentAccountConnectionRequesterAccountType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EnvironmentAccountConnectionRequesterAccountType) Values() []EnvironmentAccountConnectionRequesterAccountType {
	return []EnvironmentAccountConnectionRequesterAccountType{
		"MANAGEMENT_ACCOUNT",
		"ENVIRONMENT_ACCOUNT",
	}
}

type EnvironmentAccountConnectionStatus string

// Enum values for EnvironmentAccountConnectionStatus
const (
	EnvironmentAccountConnectionStatusPending EnvironmentAccountConnectionStatus = "PENDING"
	EnvironmentAccountConnectionStatusConnected EnvironmentAccountConnectionStatus = "CONNECTED"
	EnvironmentAccountConnectionStatusRejected EnvironmentAccountConnectionStatus = "REJECTED"
)

// Values returns all known values for EnvironmentAccountConnectionStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EnvironmentAccountConnectionStatus) Values() []EnvironmentAccountConnectionStatus {
	return []EnvironmentAccountConnectionStatus{
		"PENDING",
		"CONNECTED",
		"REJECTED",
	}
}

type ListServiceInstancesFilterBy string

// Enum values for ListServiceInstancesFilterBy
const (
	ListServiceInstancesFilterByName ListServiceInstancesFilterBy = "name"
	ListServiceInstancesFilterByDeploymentStatus ListServiceInstancesFilterBy = "deploymentStatus"
	ListServiceInstancesFilterByTemplateName ListServiceInstancesFilterBy = "templateName"
	ListServiceInstancesFilterByServiceName ListServiceInstancesFilterBy = "serviceName"
	ListServiceInstancesFilterByDeployedTemplateVersionStatus ListServiceInstancesFilterBy = "deployedTemplateVersionStatus"
	ListServiceInstancesFilterByEnvironmentName ListServiceInstancesFilterBy = "environmentName"
	ListServiceInstancesFilterByLastDeploymentAttemptedAtBefore ListServiceInstancesFilterBy = "lastDeploymentAttemptedAtBefore"
	ListServiceInstancesFilterByLastDeploymentAttemptedAtAfter ListServiceInstancesFilterBy = "lastDeploymentAttemptedAtAfter"
	ListServiceInstancesFilterByCreatedAtBefore ListServiceInstancesFilterBy = "createdAtBefore"
	ListServiceInstancesFilterByCreatedAtAfter ListServiceInstancesFilterBy = "createdAtAfter"
)

// Values returns all known values for ListServiceInstancesFilterBy. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ListServiceInstancesFilterBy) Values() []ListServiceInstancesFilterBy {
	return []ListServiceInstancesFilterBy{
		"name",
		"deploymentStatus",
		"templateName",
		"serviceName",
		"deployedTemplateVersionStatus",
		"environmentName",
		"lastDeploymentAttemptedAtBefore",
		"lastDeploymentAttemptedAtAfter",
		"createdAtBefore",
		"createdAtAfter",
	}
}

type ListServiceInstancesSortBy string

// Enum values for ListServiceInstancesSortBy
const (
	ListServiceInstancesSortByName ListServiceInstancesSortBy = "name"
	ListServiceInstancesSortByDeploymentStatus ListServiceInstancesSortBy = "deploymentStatus"
	ListServiceInstancesSortByTemplateName ListServiceInstancesSortBy = "templateName"
	ListServiceInstancesSortByServiceName ListServiceInstancesSortBy = "serviceName"
	ListServiceInstancesSortByEnvironmentName ListServiceInstancesSortBy = "environmentName"
	ListServiceInstancesSortByLastDeploymentAttemptedAt ListServiceInstancesSortBy = "lastDeploymentAttemptedAt"
	ListServiceInstancesSortByCreatedAt ListServiceInstancesSortBy = "createdAt"
)

// Values returns all known values for ListServiceInstancesSortBy. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListServiceInstancesSortBy) Values() []ListServiceInstancesSortBy {
	return []ListServiceInstancesSortBy{
		"name",
		"deploymentStatus",
		"templateName",
		"serviceName",
		"environmentName",
		"lastDeploymentAttemptedAt",
		"createdAt",
	}
}

type ProvisionedResourceEngine string

// Enum values for ProvisionedResourceEngine
const (
	ProvisionedResourceEngineCloudformation ProvisionedResourceEngine = "CLOUDFORMATION"
	ProvisionedResourceEngineTerraform ProvisionedResourceEngine = "TERRAFORM"
)

// Values returns all known values for ProvisionedResourceEngine. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ProvisionedResourceEngine) Values() []ProvisionedResourceEngine {
	return []ProvisionedResourceEngine{
		"CLOUDFORMATION",
		"TERRAFORM",
	}
}

type Provisioning string

// Enum values for Provisioning
const (
	ProvisioningCustomerManaged Provisioning = "CUSTOMER_MANAGED"
)

// Values returns all known values for Provisioning. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (Provisioning) Values() []Provisioning {
	return []Provisioning{
		"CUSTOMER_MANAGED",
	}
}

type RepositoryProvider string

// Enum values for RepositoryProvider
const (
	RepositoryProviderGithub RepositoryProvider = "GITHUB"
	RepositoryProviderGithubEnterprise RepositoryProvider = "GITHUB_ENTERPRISE"
	RepositoryProviderBitbucket RepositoryProvider = "BITBUCKET"
)

// Values returns all known values for RepositoryProvider. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RepositoryProvider) Values() []RepositoryProvider {
	return []RepositoryProvider{
		"GITHUB",
		"GITHUB_ENTERPRISE",
		"BITBUCKET",
	}
}

type RepositorySyncStatus string

// Enum values for RepositorySyncStatus
const (
	// A repository sync attempt has been created and will begin soon.
	RepositorySyncStatusInitiated RepositorySyncStatus = "INITIATED"
	// A repository sync attempt has started and work is being done to reconcile the
	// branch.
	RepositorySyncStatusInProgress RepositorySyncStatus = "IN_PROGRESS"
	// The repository sync attempt has completed successfully.
	RepositorySyncStatusSucceeded RepositorySyncStatus = "SUCCEEDED"
	// The repository sync attempt has failed.
	RepositorySyncStatusFailed RepositorySyncStatus = "FAILED"
	// The repository sync attempt didn't execute and was queued.
	RepositorySyncStatusQueued RepositorySyncStatus = "QUEUED"
)

// Values returns all known values for RepositorySyncStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RepositorySyncStatus) Values() []RepositorySyncStatus {
	return []RepositorySyncStatus{
		"INITIATED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"QUEUED",
	}
}

type ResourceDeploymentStatus string

// Enum values for ResourceDeploymentStatus
const (
	ResourceDeploymentStatusInProgress ResourceDeploymentStatus = "IN_PROGRESS"
	ResourceDeploymentStatusFailed ResourceDeploymentStatus = "FAILED"
	ResourceDeploymentStatusSucceeded ResourceDeploymentStatus = "SUCCEEDED"
)

// Values returns all known values for ResourceDeploymentStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResourceDeploymentStatus) Values() []ResourceDeploymentStatus {
	return []ResourceDeploymentStatus{
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
	}
}

type ResourceSyncStatus string

// Enum values for ResourceSyncStatus
const (
	// A sync attempt has been created and will begin soon.
	ResourceSyncStatusInitiated ResourceSyncStatus = "INITIATED"
	// Syncing has started and work is being done to reconcile state.
	ResourceSyncStatusInProgress ResourceSyncStatus = "IN_PROGRESS"
	// Syncing has completed successfully.
	ResourceSyncStatusSucceeded ResourceSyncStatus = "SUCCEEDED"
	// Syncing has failed.
	ResourceSyncStatusFailed ResourceSyncStatus = "FAILED"
)

// Values returns all known values for ResourceSyncStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceSyncStatus) Values() []ResourceSyncStatus {
	return []ResourceSyncStatus{
		"INITIATED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type ServiceStatus string

// Enum values for ServiceStatus
const (
	ServiceStatusCreateInProgress ServiceStatus = "CREATE_IN_PROGRESS"
	ServiceStatusCreateFailedCleanupInProgress ServiceStatus = "CREATE_FAILED_CLEANUP_IN_PROGRESS"
	ServiceStatusCreateFailedCleanupComplete ServiceStatus = "CREATE_FAILED_CLEANUP_COMPLETE"
	ServiceStatusCreateFailedCleanupFailed ServiceStatus = "CREATE_FAILED_CLEANUP_FAILED"
	ServiceStatusCreateFailed ServiceStatus = "CREATE_FAILED"
	ServiceStatusActive ServiceStatus = "ACTIVE"
	ServiceStatusDeleteInProgress ServiceStatus = "DELETE_IN_PROGRESS"
	ServiceStatusDeleteFailed ServiceStatus = "DELETE_FAILED"
	ServiceStatusUpdateInProgress ServiceStatus = "UPDATE_IN_PROGRESS"
	ServiceStatusUpdateFailedCleanupInProgress ServiceStatus = "UPDATE_FAILED_CLEANUP_IN_PROGRESS"
	ServiceStatusUpdateFailedCleanupComplete ServiceStatus = "UPDATE_FAILED_CLEANUP_COMPLETE"
	ServiceStatusUpdateFailedCleanupFailed ServiceStatus = "UPDATE_FAILED_CLEANUP_FAILED"
	ServiceStatusUpdateFailed ServiceStatus = "UPDATE_FAILED"
	ServiceStatusUpdateCompleteCleanupFailed ServiceStatus = "UPDATE_COMPLETE_CLEANUP_FAILED"
)

// Values returns all known values for ServiceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceStatus) Values() []ServiceStatus {
	return []ServiceStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED_CLEANUP_IN_PROGRESS",
		"CREATE_FAILED_CLEANUP_COMPLETE",
		"CREATE_FAILED_CLEANUP_FAILED",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED_CLEANUP_IN_PROGRESS",
		"UPDATE_FAILED_CLEANUP_COMPLETE",
		"UPDATE_FAILED_CLEANUP_FAILED",
		"UPDATE_FAILED",
		"UPDATE_COMPLETE_CLEANUP_FAILED",
	}
}

type ServiceTemplateSupportedComponentSourceType string

// Enum values for ServiceTemplateSupportedComponentSourceType
const (
	ServiceTemplateSupportedComponentSourceTypeDirectlyDefined ServiceTemplateSupportedComponentSourceType = "DIRECTLY_DEFINED"
)

// Values returns all known values for
// ServiceTemplateSupportedComponentSourceType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ServiceTemplateSupportedComponentSourceType) Values() []ServiceTemplateSupportedComponentSourceType {
	return []ServiceTemplateSupportedComponentSourceType{
		"DIRECTLY_DEFINED",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type SyncType string

// Enum values for SyncType
const (
	// Syncs environment and service templates to Proton.
	SyncTypeTemplateSync SyncType = "TEMPLATE_SYNC"
	// Syncs services and service instances to Proton.
	SyncTypeServiceSync SyncType = "SERVICE_SYNC"
)

// Values returns all known values for SyncType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (SyncType) Values() []SyncType {
	return []SyncType{
		"TEMPLATE_SYNC",
		"SERVICE_SYNC",
	}
}

type TemplateType string

// Enum values for TemplateType
const (
	TemplateTypeEnvironment TemplateType = "ENVIRONMENT"
	TemplateTypeService TemplateType = "SERVICE"
)

// Values returns all known values for TemplateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateType) Values() []TemplateType {
	return []TemplateType{
		"ENVIRONMENT",
		"SERVICE",
	}
}

type TemplateVersionStatus string

// Enum values for TemplateVersionStatus
const (
	TemplateVersionStatusRegistrationInProgress TemplateVersionStatus = "REGISTRATION_IN_PROGRESS"
	TemplateVersionStatusRegistrationFailed TemplateVersionStatus = "REGISTRATION_FAILED"
	TemplateVersionStatusDraft TemplateVersionStatus = "DRAFT"
	TemplateVersionStatusPublished TemplateVersionStatus = "PUBLISHED"
)

// Values returns all known values for TemplateVersionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateVersionStatus) Values() []TemplateVersionStatus {
	return []TemplateVersionStatus{
		"REGISTRATION_IN_PROGRESS",
		"REGISTRATION_FAILED",
		"DRAFT",
		"PUBLISHED",
	}
}
