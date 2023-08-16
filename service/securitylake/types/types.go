// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The AWS identity.
type AwsIdentity struct {
	
	// The external ID used to estalish trust relationship with the AWS identity.
	//
	// This member is required.
	ExternalId *string
	
	// The AWS identity principal.
	//
	// This member is required.
	Principal *string
	
	noSmithyDocumentSerde
}

// The Security Lake logs source configuration file describes the information
// needed to generate Security Lake logs.
type AwsLogSourceConfiguration struct {
	
	// Specify the Regions where you want to enable Security Lake.
	//
	// This member is required.
	Regions []string
	
	// The name for a Amazon Web Services source. This must be a Regionally unique
	// value.
	//
	// This member is required.
	SourceName AwsLogSourceName
	
	// Specify the Amazon Web Services account information where you want to enable
	// Security Lake.
	Accounts []string
	
	// The version for a Amazon Web Services source. This must be a Regionally unique
	// value.
	SourceVersion *string
	
	noSmithyDocumentSerde
}

// Amazon Security Lake can collect logs and events from natively-supported Amazon
// Web Services services.
type AwsLogSourceResource struct {
	
	// The name for a Amazon Web Services source. This must be a Regionally unique
	// value.
	SourceName AwsLogSourceName
	
	// The version for a Amazon Web Services source. This must be a Regionally unique
	// value.
	SourceVersion *string
	
	noSmithyDocumentSerde
}

// The attributes of a third-party custom source.
type CustomLogSourceAttributes struct {
	
	// The ARN of the Glue crawler.
	CrawlerArn *string
	
	// The ARN of the Glue database where results are written, such as:
	// arn:aws:daylight:us-east-1::database/sometable/* .
	DatabaseArn *string
	
	// The ARN of the Glue table.
	TableArn *string
	
	noSmithyDocumentSerde
}

// The configuration for the third-party custom source.
type CustomLogSourceConfiguration struct {
	
	// The configuration for the Glue Crawler for the third-party custom source.
	//
	// This member is required.
	CrawlerConfiguration *CustomLogSourceCrawlerConfiguration
	
	// The identity of the log provider for the third-party custom source.
	//
	// This member is required.
	ProviderIdentity *AwsIdentity
	
	noSmithyDocumentSerde
}

// The configuration for the Glue Crawler for the third-party custom source.
type CustomLogSourceCrawlerConfiguration struct {
	
	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// to be used by the Glue crawler. The recommended IAM policies are:
	//   - The managed policy AWSGlueServiceRole
	//   - A custom policy granting access to your Amazon S3 Data Lake
	//
	// This member is required.
	RoleArn *string
	
	noSmithyDocumentSerde
}

// The details of the log provider for a third-party custom source.
type CustomLogSourceProvider struct {
	
	// The location of the partition in the Amazon S3 bucket for Security Lake.
	Location *string
	
	// The ARN of the IAM role to be used by the entity putting logs into your custom
	// source partition. Security Lake will apply the correct access policies to this
	// role, but you must first manually create the trust policy for this role. The IAM
	// role name must start with the text 'Security Lake'. The IAM role must trust the
	// logProviderAccountId to assume the role.
	RoleArn *string
	
	noSmithyDocumentSerde
}

// Amazon Security Lake can collect logs and events from third-party custom
// sources.
type CustomLogSourceResource struct {
	
	// The attributes of a third-party custom source.
	Attributes *CustomLogSourceAttributes
	
	// The details of the log provider for a third-party custom source.
	Provider *CustomLogSourceProvider
	
	// The name for a third-party custom source. This must be a Regionally unique
	// value.
	SourceName *string
	
	// The version for a third-party custom source. This must be a Regionally unique
	// value.
	SourceVersion *string
	
	noSmithyDocumentSerde
}

// Automatically enable new organization accounts as member accounts from an
// Amazon Security Lake administrator account.
type DataLakeAutoEnableNewAccountConfiguration struct {
	
	// The Amazon Web Services Regions where Security Lake is automatically enabled.
	//
	// This member is required.
	Region *string
	
	// The Amazon Web Services sources that are automatically enabled in Security Lake.
	//
	// This member is required.
	Sources []AwsLogSourceResource
	
	noSmithyDocumentSerde
}

// Provides details of Amazon Security Lake object.
type DataLakeConfiguration struct {
	
	// The Amazon Web Services Regions where Security Lake is automatically enabled.
	//
	// This member is required.
	Region *string
	
	// Provides encryption details of Amazon Security Lake object.
	EncryptionConfiguration *DataLakeEncryptionConfiguration
	
	// Provides lifecycle details of Amazon Security Lake object.
	LifecycleConfiguration *DataLakeLifecycleConfiguration
	
	// Provides replication details of Amazon Security Lake object.
	ReplicationConfiguration *DataLakeReplicationConfiguration
	
	noSmithyDocumentSerde
}

// Provides encryption details of Amazon Security Lake object.
type DataLakeEncryptionConfiguration struct {
	
	// The id of KMS encryption key used by Amazon Security Lake to encrypt the
	// Security Lake object.
	KmsKeyId *string
	
	noSmithyDocumentSerde
}

// The details for an Amazon Security Lake exception.
type DataLakeException struct {
	
	// The underlying exception of a Security Lake exception.
	Exception *string
	
	// The Amazon Web Services Regions where the exception occurred.
	Region *string
	
	// List of all remediation steps for a Security Lake exception.
	Remediation *string
	
	// This error can occur if you configure the wrong timestamp format, or if the
	// subset of entries used for validation had errors or missing values.
	Timestamp *time.Time
	
	noSmithyDocumentSerde
}

// Provides lifecycle details of Amazon Security Lake object.
type DataLakeLifecycleConfiguration struct {
	
	// Provides data expiration details of Amazon Security Lake object.
	Expiration *DataLakeLifecycleExpiration
	
	// Provides data storage transition details of Amazon Security Lake object.
	Transitions []DataLakeLifecycleTransition
	
	noSmithyDocumentSerde
}

// Provide expiration lifecycle details of Amazon Security Lake object.
type DataLakeLifecycleExpiration struct {
	
	// Number of days before data expires in the Amazon Security Lake object.
	Days *int32
	
	noSmithyDocumentSerde
}

// Provide transition lifecycle details of Amazon Security Lake object.
type DataLakeLifecycleTransition struct {
	
	// Number of days before data transitions to a different S3 Storage Class in the
	// Amazon Security Lake object.
	Days *int32
	
	// The range of storage classes that you can choose from based on the data access,
	// resiliency, and cost requirements of your workloads.
	StorageClass *string
	
	noSmithyDocumentSerde
}

// Provides replication details of Amazon Security Lake object.
type DataLakeReplicationConfiguration struct {
	
	// Replication enables automatic, asynchronous copying of objects across Amazon S3
	// buckets. Amazon S3 buckets that are configured for object replication can be
	// owned by the same Amazon Web Services account or by different accounts. You can
	// replicate objects to a single destination bucket or to multiple destination
	// buckets. The destination buckets can be in different Amazon Web Services Regions
	// or within the same Region as the source bucket. Set up one or more rollup
	// Regions by providing the Region or Regions that should contribute to the central
	// rollup Region.
	Regions []string
	
	// Replication settings for the Amazon S3 buckets. This parameter uses the
	// Identity and Access Management (IAM) role you created that is managed by
	// Security Lake, to ensure the replication setting is correct.
	RoleArn *string
	
	noSmithyDocumentSerde
}

// Provides details of Amazon Security Lake object.
type DataLakeResource struct {
	
	// The Amazon Resource Name (ARN) created by you to provide to the subscriber. For
	// more information about ARNs and how to use them in policies, see the Amazon
	// Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/subscriber-management.html)
	// .
	//
	// This member is required.
	DataLakeArn *string
	
	// The Amazon Web Services Regions where Security Lake is enabled.
	//
	// This member is required.
	Region *string
	
	// Retrieves the status of the configuration operation for an account in Amazon
	// Security Lake.
	CreateStatus DataLakeStatus
	
	// Provides encryption details of Amazon Security Lake object.
	EncryptionConfiguration *DataLakeEncryptionConfiguration
	
	// Provides lifecycle details of Amazon Security Lake object.
	LifecycleConfiguration *DataLakeLifecycleConfiguration
	
	// Provides replication details of Amazon Security Lake object.
	ReplicationConfiguration *DataLakeReplicationConfiguration
	
	// The ARN for the Amazon Security Lake Amazon S3 bucket.
	S3BucketArn *string
	
	// The status of the last UpdateDataLake or DeleteDataLake API request.
	UpdateStatus *DataLakeUpdateStatus
	
	noSmithyDocumentSerde
}

// Amazon Security Lake collects logs and events from supported Amazon Web
// Services and custom sources. For the list of supported Amazon Web Services, see
// the Amazon Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html)
// .
type DataLakeSource struct {
	
	// The ID of the Security Lake account for which logs are collected.
	Account *string
	
	// The Open Cybersecurity Schema Framework (OCSF) event classes which describes
	// the type of data that the custom source will send to Security Lake. The
	// supported event classes are:
	//   - ACCESS_ACTIVITY
	//   - FILE_ACTIVITY
	//   - KERNEL_ACTIVITY
	//   - KERNEL_EXTENSION
	//   - MEMORY_ACTIVITY
	//   - MODULE_ACTIVITY
	//   - PROCESS_ACTIVITY
	//   - REGISTRY_KEY_ACTIVITY
	//   - REGISTRY_VALUE_ACTIVITY
	//   - RESOURCE_ACTIVITY
	//   - SCHEDULED_JOB_ACTIVITY
	//   - SECURITY_FINDING
	//   - ACCOUNT_CHANGE
	//   - AUTHENTICATION
	//   - AUTHORIZATION
	//   - ENTITY_MANAGEMENT_AUDIT
	//   - DHCP_ACTIVITY
	//   - NETWORK_ACTIVITY
	//   - DNS_ACTIVITY
	//   - FTP_ACTIVITY
	//   - HTTP_ACTIVITY
	//   - RDP_ACTIVITY
	//   - SMB_ACTIVITY
	//   - SSH_ACTIVITY
	//   - CONFIG_STATE
	//   - INVENTORY_INFO
	//   - EMAIL_ACTIVITY
	//   - API_ACTIVITY
	//   - CLOUD_API
	EventClasses []string
	
	// The supported Amazon Web Services from which logs and events are collected.
	// Amazon Security Lake supports log and event collection for natively supported
	// Amazon Web Services.
	SourceName *string
	
	// The log status for the Security Lake account.
	SourceStatuses []DataLakeSourceStatus
	
	noSmithyDocumentSerde
}

// Retrieves the Logs status for the Amazon Security Lake account.
type DataLakeSourceStatus struct {
	
	// Defines path the stored logs are available which has information on your
	// systems, applications, and services.
	Resource *string
	
	// The health status of services, including error codes and patterns.
	Status SourceCollectionStatus
	
	noSmithyDocumentSerde
}

// The details of the last UpdateDataLake or DeleteDataLake API request which
// failed.
type DataLakeUpdateException struct {
	
	// The reason code for the exception of the last UpdateDataLake or DeleteDataLake
	// API request.
	Code *string
	
	// The reason for the exception of the last UpdateDataLake or DeleteDataLake API
	// request.
	Reason *string
	
	noSmithyDocumentSerde
}

// The status of the last UpdateDataLake or DeleteDataLake API request. This is
// set to Completed after the configuration is updated, or removed if deletion of
// the data lake is successful.
type DataLakeUpdateStatus struct {
	
	// The details of the last UpdateDataLake or DeleteDataLake API request which
	// failed.
	Exception *DataLakeUpdateException
	
	// The unique ID for the last UpdateDataLake or DeleteDataLake API request.
	RequestId *string
	
	// The status of the last UpdateDataLake or DeleteDataLake API request that was
	// requested.
	Status DataLakeStatus
	
	noSmithyDocumentSerde
}

// The configurations for HTTPS subscriber notification.
type HttpsNotificationConfiguration struct {
	
	// The subscription endpoint in Security Lake. If you prefer notification with an
	// HTTPs endpoint, populate this field.
	//
	// This member is required.
	Endpoint *string
	
	// The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role
	// that you created. For more information about ARNs and how to use them in
	// policies, see Managing data access (https://docs.aws.amazon.com//security-lake/latest/userguide/subscriber-data-access.html)
	// and Amazon Web Services Managed Policies (https://docs.aws.amazon.com/security-lake/latest/userguide/security-iam-awsmanpol.html)
	// in the Amazon Security Lake User Guide.
	//
	// This member is required.
	TargetRoleArn *string
	
	// The key name for the notification subscription.
	AuthorizationApiKeyName *string
	
	// The key value for the notification subscription.
	AuthorizationApiKeyValue *string
	
	// The HTTPS method used for the notification subscription.
	HttpMethod HttpMethod
	
	noSmithyDocumentSerde
}

// Amazon Security Lake can collect logs and events from natively-supported Amazon
// Web Services services and custom sources.
type LogSource struct {
	
	// Specify the account from which you want to collect logs.
	Account *string
	
	// Specify the Regions from which you want to collect logs.
	Region *string
	
	// Specify the sources from which you want to collect logs.
	Sources []LogSourceResource
	
	noSmithyDocumentSerde
}

// The supported source types from which logs and events are collected in Amazon
// Security Lake. For a list of supported Amazon Web Services, see the Amazon
// Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html)
// .
//
// The following types satisfy this interface:
//  LogSourceResourceMemberAwsLogSource
//  LogSourceResourceMemberCustomLogSource
type LogSourceResource interface {
	isLogSourceResource()
}

// Amazon Security Lake supports log and event collection for natively supported
// Amazon Web Services. For more information, see the Amazon Security Lake User
// Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html)
// .
type LogSourceResourceMemberAwsLogSource struct {
	Value AwsLogSourceResource
	
	noSmithyDocumentSerde
}
func (*LogSourceResourceMemberAwsLogSource) isLogSourceResource() {}
// Amazon Security Lake supports custom source types. For more information, see
// the Amazon Security Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/custom-sources.html)
// .
type LogSourceResourceMemberCustomLogSource struct {
	Value CustomLogSourceResource
	
	noSmithyDocumentSerde
}
func (*LogSourceResourceMemberCustomLogSource) isLogSourceResource() {}

// Specify the configurations you want to use for subscriber notification to
// notify the subscriber when new data is written to the data lake for sources that
// the subscriber consumes in Security Lake.
//
// The following types satisfy this interface:
//  NotificationConfigurationMemberHttpsNotificationConfiguration
//  NotificationConfigurationMemberSqsNotificationConfiguration
type NotificationConfiguration interface {
	isNotificationConfiguration()
}

// The configurations for HTTPS subscriber notification.
type NotificationConfigurationMemberHttpsNotificationConfiguration struct {
	Value HttpsNotificationConfiguration
	
	noSmithyDocumentSerde
}
func (*NotificationConfigurationMemberHttpsNotificationConfiguration) isNotificationConfiguration() {}
// The configurations for SQS subscriber notification.
type NotificationConfigurationMemberSqsNotificationConfiguration struct {
	Value SqsNotificationConfiguration
	
	noSmithyDocumentSerde
}
func (*NotificationConfigurationMemberSqsNotificationConfiguration) isNotificationConfiguration() {}

// The configurations for SQS subscriber notification.
type SqsNotificationConfiguration struct {
	
	noSmithyDocumentSerde
}

// Provides details about the Amazon Security Lake account subscription.
// Subscribers are notified of new objects for a source as the data is written to
// your Amazon S3 bucket for Security Lake.
type SubscriberResource struct {
	
	// Amazon Security Lake supports log and event collection for natively supported
	// Amazon Web Services. For more information, see the Amazon Security Lake User
	// Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/source-management.html)
	// .
	//
	// This member is required.
	Sources []LogSourceResource
	
	// The subscriber ARN of the Amazon Security Lake subscriber account.
	//
	// This member is required.
	SubscriberArn *string
	
	// The subscriber ID of the Amazon Security Lake subscriber account.
	//
	// This member is required.
	SubscriberId *string
	
	// The AWS identity used to access your data.
	//
	// This member is required.
	SubscriberIdentity *AwsIdentity
	
	// The name of your Amazon Security Lake subscriber account.
	//
	// This member is required.
	SubscriberName *string
	
	// You can choose to notify subscribers of new objects with an Amazon Simple Queue
	// Service (Amazon SQS) queue or through messaging to an HTTPS endpoint provided by
	// the subscriber. Subscribers can consume data by directly querying Lake Formation
	// tables in your Amazon S3 bucket through services like Amazon Athena. This
	// subscription type is defined as LAKEFORMATION .
	AccessTypes []AccessType
	
	// The date and time when the subscriber was created.
	CreatedAt *time.Time
	
	// The Amazon Resource Name (ARN) which uniquely defines the AWS RAM resource
	// share. Before accepting the RAM resource share invitation, you can view details
	// related to the RAM resource share. This field is available only for Lake
	// Formation subscribers created after March 8, 2023.
	ResourceShareArn *string
	
	// The name of the resource share.
	ResourceShareName *string
	
	// The Amazon Resource Name (ARN) specifying the role of the subscriber.
	RoleArn *string
	
	// The ARN for the Amazon S3 bucket.
	S3BucketArn *string
	
	// The subscriber descriptions for a subscriber account. The description for a
	// subscriber includes subscriberName , accountID , externalID , and subscriberId .
	SubscriberDescription *string
	
	// The subscriber endpoint to which exception messages are posted.
	SubscriberEndpoint *string
	
	// The subscriber status of the Amazon Security Lake subscriber account.
	SubscriberStatus SubscriberStatus
	
	// The date and time when the subscriber was last updated.
	UpdatedAt *time.Time
	
	noSmithyDocumentSerde
}

// A tag is a label that you can define and associate with Amazon Web Services
// resources, including certain types of Amazon Security Lake resources. Tags can
// help you identify, categorize, and manage resources in different ways, such as
// by owner, environment, or other criteria. You can associate tags with the
// following types of Security Lake resources: subscribers, and the data lake
// configuration for your Amazon Web Services account in individual Amazon Web
// Services Regions. A resource can have up to 50 tags. Each tag consists of a
// required tag key and an associated tag value. A tag key is a general label that
// acts as a category for a more specific tag value. Each tag key must be unique
// and it can have only one tag value. A tag value acts as a descriptor for a tag
// key. Tag keys and values are case sensitive. They can contain letters, numbers,
// spaces, or the following symbols: _ . : / = + @ - For more information, see
// Tagging Amazon Security Lake resources (https://docs.aws.amazon.com/security-lake/latest/userguide/tagging-resources.html)
// in the Amazon Security Lake User Guide.
type Tag struct {
	
	// The name of the tag. This is a general label that acts as a category for a more
	// specific tag value ( value ).
	//
	// This member is required.
	Key *string
	
	// The value that’s associated with the specified tag key ( key ). This value acts
	// as a descriptor for the tag key. A tag value cannot be null, but it can be an
	// empty string.
	//
	// This member is required.
	Value *string
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag string
	Value []byte
	
	noSmithyDocumentSerde
}
func (*UnknownUnionMember) isLogSourceResource() {}
func (*UnknownUnionMember) isNotificationConfiguration() {}
