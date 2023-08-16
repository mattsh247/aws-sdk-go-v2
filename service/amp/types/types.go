// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the properties of an alert manager definition.
type AlertManagerDefinitionDescription struct {
	
	// The time when the alert manager definition was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The alert manager definition.
	//
	// This member is required.
	Data []byte
	
	// The time when the alert manager definition was modified.
	//
	// This member is required.
	ModifiedAt *time.Time
	
	// The status of alert manager definition.
	//
	// This member is required.
	Status *AlertManagerDefinitionStatus
	
	noSmithyDocumentSerde
}

// Represents the status of a definition.
type AlertManagerDefinitionStatus struct {
	
	// Status code of this definition.
	//
	// This member is required.
	StatusCode AlertManagerDefinitionStatusCode
	
	// The reason for failure if any.
	StatusReason *string
	
	noSmithyDocumentSerde
}

// Represents the properties of a logging configuration metadata.
type LoggingConfigurationMetadata struct {
	
	// The time when the logging configuration was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The ARN of the CW log group to which the vended log data will be published.
	//
	// This member is required.
	LogGroupArn *string
	
	// The time when the logging configuration was modified.
	//
	// This member is required.
	ModifiedAt *time.Time
	
	// The status of the logging configuration.
	//
	// This member is required.
	Status *LoggingConfigurationStatus
	
	// The workspace where the logging configuration exists.
	//
	// This member is required.
	Workspace *string
	
	noSmithyDocumentSerde
}

// Represents the status of a logging configuration.
type LoggingConfigurationStatus struct {
	
	// Status code of the logging configuration.
	//
	// This member is required.
	StatusCode LoggingConfigurationStatusCode
	
	// The reason for failure if any.
	StatusReason *string
	
	noSmithyDocumentSerde
}

// Represents a description of the rule groups namespace.
type RuleGroupsNamespaceDescription struct {
	
	// The Amazon Resource Name (ARN) of this rule groups namespace.
	//
	// This member is required.
	Arn *string
	
	// The time when the rule groups namespace was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The rule groups namespace data.
	//
	// This member is required.
	Data []byte
	
	// The time when the rule groups namespace was modified.
	//
	// This member is required.
	ModifiedAt *time.Time
	
	// The rule groups namespace name.
	//
	// This member is required.
	Name *string
	
	// The status of rule groups namespace.
	//
	// This member is required.
	Status *RuleGroupsNamespaceStatus
	
	// The tags of this rule groups namespace.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

// Represents the status of a namespace.
type RuleGroupsNamespaceStatus struct {
	
	// Status code of this namespace.
	//
	// This member is required.
	StatusCode RuleGroupsNamespaceStatusCode
	
	// The reason for failure if any.
	StatusReason *string
	
	noSmithyDocumentSerde
}

// Represents a summary of the rule groups namespace.
type RuleGroupsNamespaceSummary struct {
	
	// The Amazon Resource Name (ARN) of this rule groups namespace.
	//
	// This member is required.
	Arn *string
	
	// The time when the rule groups namespace was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The time when the rule groups namespace was modified.
	//
	// This member is required.
	ModifiedAt *time.Time
	
	// The rule groups namespace name.
	//
	// This member is required.
	Name *string
	
	// The status of rule groups namespace.
	//
	// This member is required.
	Status *RuleGroupsNamespaceStatus
	
	// The tags of this rule groups namespace.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

// Stores information about a field passed inside a request that resulted in an
// exception.
type ValidationExceptionField struct {
	
	// Message describing why the field failed validation.
	//
	// This member is required.
	Message *string
	
	// The field name.
	//
	// This member is required.
	Name *string
	
	noSmithyDocumentSerde
}

// Represents the properties of a workspace.
type WorkspaceDescription struct {
	
	// The Amazon Resource Name (ARN) of this workspace.
	//
	// This member is required.
	Arn *string
	
	// The time when the workspace was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The status of this workspace.
	//
	// This member is required.
	Status *WorkspaceStatus
	
	// Unique string identifying this workspace.
	//
	// This member is required.
	WorkspaceId *string
	
	// Alias of this workspace.
	Alias *string
	
	// Prometheus endpoint URI.
	PrometheusEndpoint *string
	
	// The tags of this workspace.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

// Represents the status of a workspace.
type WorkspaceStatus struct {
	
	// Status code of this workspace.
	//
	// This member is required.
	StatusCode WorkspaceStatusCode
	
	noSmithyDocumentSerde
}

// Represents a summary of the properties of a workspace.
type WorkspaceSummary struct {
	
	// The AmazonResourceName of this workspace.
	//
	// This member is required.
	Arn *string
	
	// The time when the workspace was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The status of this workspace.
	//
	// This member is required.
	Status *WorkspaceStatus
	
	// Unique string identifying this workspace.
	//
	// This member is required.
	WorkspaceId *string
	
	// Alias of this workspace.
	Alias *string
	
	// The tags of this workspace.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
