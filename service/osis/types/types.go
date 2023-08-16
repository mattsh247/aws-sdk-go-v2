// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Progress details for a specific stage of a pipeline configuration change.
type ChangeProgressStage struct {

	// A description of the stage.
	Description *string

	// The most recent updated timestamp of the stage.
	LastUpdatedAt *time.Time

	// The name of the stage.
	Name *string

	// The current status of the stage that the change is in.
	Status ChangeProgressStageStatuses

	noSmithyDocumentSerde
}

// The progress details of a pipeline configuration change.
type ChangeProgressStatus struct {

	// Information about the stages that the pipeline is going through to perform the
	// configuration change.
	ChangeProgressStages []ChangeProgressStage

	// The time at which the configuration change is made on the pipeline.
	StartTime *time.Time

	// The overall status of the pipeline configuration change.
	Status ChangeProgressStatuses

	// The total number of stages required for the pipeline configuration change.
	TotalNumberOfStages int32

	noSmithyDocumentSerde
}

// The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch.
type CloudWatchLogDestination struct {

	// The name of the CloudWatch Logs group to send pipeline logs to. You can specify
	// an existing log group or create a new one. For example,
	// /aws/OpenSearchService/IngestionService/my-pipeline .
	//
	// This member is required.
	LogGroup *string

	noSmithyDocumentSerde
}

// Container for the values required to configure logging for the pipeline. If you
// don't specify these values, OpenSearch Ingestion will not publish logs from your
// application to CloudWatch Logs.
type LogPublishingOptions struct {

	// The destination for OpenSearch Ingestion logs sent to Amazon CloudWatch Logs.
	// This parameter is required if IsLoggingEnabled is set to true .
	CloudWatchLogDestination *CloudWatchLogDestination

	// Whether logs should be published.
	IsLoggingEnabled *bool

	noSmithyDocumentSerde
}

// Information about an existing OpenSearch Ingestion pipeline.
type Pipeline struct {

	// The date and time when the pipeline was created.
	CreatedAt *time.Time

	// The ingestion endpoints for the pipeline, which you can send data to.
	IngestEndpointUrls []string

	// The date and time when the pipeline was last updated.
	LastUpdatedAt *time.Time

	// Key-value pairs that represent log publishing settings.
	LogPublishingOptions *LogPublishingOptions

	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	MaxUnits int32

	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	MinUnits int32

	// The Amazon Resource Name (ARN) of the pipeline.
	PipelineArn *string

	// The Data Prepper pipeline configuration in YAML format.
	PipelineConfigurationBody *string

	// The name of the pipeline.
	PipelineName *string

	// The current status of the pipeline.
	Status PipelineStatus

	// The reason for the current status of the pipeline.
	StatusReason *PipelineStatusReason

	// The VPC interface endpoints that have access to the pipeline.
	VpcEndpoints []VpcEndpoint

	noSmithyDocumentSerde
}

// Container for information about an OpenSearch Ingestion blueprint.
type PipelineBlueprint struct {

	// The name of the blueprint.
	BlueprintName *string

	// The YAML configuration of the blueprint.
	PipelineConfigurationBody *string

	noSmithyDocumentSerde
}

// A summary of an OpenSearch Ingestion blueprint.
type PipelineBlueprintSummary struct {

	// The name of the blueprint.
	BlueprintName *string

	noSmithyDocumentSerde
}

// Information about a pipeline's current status.
type PipelineStatusReason struct {

	// A description of why a pipeline has a certain status.
	Description *string

	noSmithyDocumentSerde
}

// Summary information for an OpenSearch Ingestion pipeline.
type PipelineSummary struct {

	// The date and time when the pipeline was created.
	CreatedAt *time.Time

	// The date and time when the pipeline was last updated.
	LastUpdatedAt *time.Time

	// The maximum pipeline capacity, in Ingestion Compute Units (ICUs).
	MaxUnits *int32

	// The minimum pipeline capacity, in Ingestion Compute Units (ICUs).
	MinUnits *int32

	// The Amazon Resource Name (ARN) of the pipeline.
	PipelineArn *string

	// The name of the pipeline.
	PipelineName *string

	// The current status of the pipeline.
	Status PipelineStatus

	// Information about a pipeline's current status.
	StatusReason *PipelineStatusReason

	noSmithyDocumentSerde
}

// A tag (key-value pair) for an OpenSearch Ingestion pipeline.
type Tag struct {

	// The tag key. Tag keys must be unique for the pipeline to which they are
	// attached.
	//
	// This member is required.
	Key *string

	// The value assigned to the corresponding tag key. Tag values can be null and
	// don't have to be unique in a tag set. For example, you can have a key value pair
	// in a tag set of project : Trinity and cost-center : Trinity
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// A validation message associated with a ValidatePipeline request in OpenSearch
// Ingestion.
type ValidationMessage struct {

	// The validation message.
	Message *string

	noSmithyDocumentSerde
}

// An OpenSearch Ingestion-managed VPC endpoint that will access one or more
// pipelines.
type VpcEndpoint struct {

	// The unique identifier of the endpoint.
	VpcEndpointId *string

	// The ID for your VPC. Amazon Web Services PrivateLink generates this value when
	// you create a VPC.
	VpcId *string

	// Information about the VPC, including associated subnets and security groups.
	VpcOptions *VpcOptions

	noSmithyDocumentSerde
}

// Options that specify the subnets and security groups for an OpenSearch
// Ingestion VPC endpoint.
type VpcOptions struct {

	// A list of subnet IDs associated with the VPC endpoint.
	//
	// This member is required.
	SubnetIds []string

	// A list of security groups associated with the VPC endpoint.
	SecurityGroupIds []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
