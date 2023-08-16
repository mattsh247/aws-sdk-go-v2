// Code generated by smithy-go-codegen DO NOT EDIT.


package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a cluster. All nodes in the cluster run the same protocol-compliant
// cache engine software, either Memcached or Redis. This operation is not
// supported for Redis (cluster mode enabled) clusters.
func (c *Client) CreateCacheCluster(ctx context.Context, params *CreateCacheClusterInput, optFns ...func(*Options)) (*CreateCacheClusterOutput, error) {
	if params == nil { params = &CreateCacheClusterInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateCacheCluster", params, optFns, c.addOperationCreateCacheClusterMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateCacheClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCacheCluster operation.
type CreateCacheClusterInput struct {
	
	// The node group (shard) identifier. This parameter is stored as a lowercase
	// string. Constraints:
	//   - A name must contain from 1 to 50 alphanumeric characters or hyphens.
	//   - The first character must be a letter.
	//   - A name cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	CacheClusterId *string
	
	// Specifies whether the nodes in this Memcached cluster are created in a single
	// Availability Zone or created across multiple Availability Zones in the cluster's
	// region. This parameter is only supported for Memcached clusters. If the AZMode
	// and PreferredAvailabilityZones are not specified, ElastiCache assumes single-az
	// mode.
	AZMode types.AZMode
	
	// Reserved parameter. The password used to access a password protected server.
	// Password constraints:
	//   - Must be only printable ASCII characters.
	//   - Must be at least 16 characters and no more than 128 characters in length.
	//   - The only permitted printable special characters are !, &, #, $, ^, <, >,
	//   and -. Other printable special characters cannot be used in the AUTH token.
	// For more information, see AUTH password (http://redis.io/commands/AUTH) at
	// http://redis.io/commands/AUTH.
	AuthToken *string
	
	//  If you are running Redis engine version 6.0 or later, set this parameter to
	// yes if you want to opt-in to the next auto minor version upgrade campaign. This
	// parameter is disabled for previous versions.
	AutoMinorVersionUpgrade *bool
	
	// The compute and memory capacity of the nodes in the node group (shard). The
	// following node types are supported by ElastiCache. Generally speaking, the
	// current generation types provide more memory and computational power at lower
	// cost when compared to their equivalent previous generation counterparts.
	//   - General purpose:
	//   - Current generation: M6g node types (available only for Redis engine version
	//   5.0.6 onward and for Memcached engine version 1.5.16 onward): cache.m6g.large
	//   , cache.m6g.xlarge , cache.m6g.2xlarge , cache.m6g.4xlarge , cache.m6g.8xlarge
	//   , cache.m6g.12xlarge , cache.m6g.16xlarge For region availability, see
	//   Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	//   M5 node types: cache.m5.large , cache.m5.xlarge , cache.m5.2xlarge ,
	//   cache.m5.4xlarge , cache.m5.12xlarge , cache.m5.24xlarge M4 node types:
	//   cache.m4.large , cache.m4.xlarge , cache.m4.2xlarge , cache.m4.4xlarge ,
	//   cache.m4.10xlarge T4g node types (available only for Redis engine version
	//   5.0.6 onward and Memcached engine version 1.5.16 onward): cache.t4g.micro ,
	//   cache.t4g.small , cache.t4g.medium T3 node types: cache.t3.micro ,
	//   cache.t3.small , cache.t3.medium T2 node types: cache.t2.micro ,
	//   cache.t2.small , cache.t2.medium
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.) T1
	//   node types: cache.t1.micro M1 node types: cache.m1.small , cache.m1.medium ,
	//   cache.m1.large , cache.m1.xlarge M3 node types: cache.m3.medium ,
	//   cache.m3.large , cache.m3.xlarge , cache.m3.2xlarge
	//   - Compute optimized:
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.) C1
	//   node types: cache.c1.xlarge
	//   - Memory optimized:
	//   - Current generation: R6g node types (available only for Redis engine version
	//   5.0.6 onward and for Memcached engine version 1.5.16 onward). cache.r6g.large
	//   , cache.r6g.xlarge , cache.r6g.2xlarge , cache.r6g.4xlarge , cache.r6g.8xlarge
	//   , cache.r6g.12xlarge , cache.r6g.16xlarge For region availability, see
	//   Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion)
	//   R5 node types: cache.r5.large , cache.r5.xlarge , cache.r5.2xlarge ,
	//   cache.r5.4xlarge , cache.r5.12xlarge , cache.r5.24xlarge R4 node types:
	//   cache.r4.large , cache.r4.xlarge , cache.r4.2xlarge , cache.r4.4xlarge ,
	//   cache.r4.8xlarge , cache.r4.16xlarge
	//   - Previous generation: (not recommended. Existing clusters are still
	//   supported but creation of new clusters is not supported for these types.) M2
	//   node types: cache.m2.xlarge , cache.m2.2xlarge , cache.m2.4xlarge R3 node
	//   types: cache.r3.large , cache.r3.xlarge , cache.r3.2xlarge ,
	//
	// cache.r3.4xlarge , cache.r3.8xlarge
	// Additional node type info
	//   - All current generation instance types are created in Amazon VPC by default.
	//   - Redis append-only files (AOF) are not supported for T1 or T2 instances.
	//   - Redis Multi-AZ with automatic failover is not supported on T1 instances.
	//   - Redis configuration variables appendonly and appendfsync are not supported
	//   on Redis version 2.8.22 and later.
	CacheNodeType *string
	
	// The name of the parameter group to associate with this cluster. If this
	// argument is omitted, the default parameter group for the specified engine is
	// used. You cannot use any parameter group which has cluster-enabled='yes' when
	// creating a cluster.
	CacheParameterGroupName *string
	
	// A list of security group names to associate with this cluster. Use this
	// parameter only when you are creating a cluster outside of an Amazon Virtual
	// Private Cloud (Amazon VPC).
	CacheSecurityGroupNames []string
	
	// The name of the subnet group to be used for the cluster. Use this parameter
	// only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon
	// VPC). If you're going to launch your cluster in an Amazon VPC, you need to
	// create a subnet group before you start creating a cluster. For more information,
	// see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html)
	// .
	CacheSubnetGroupName *string
	
	// The name of the cache engine to be used for this cluster. Valid values for this
	// parameter are: memcached | redis
	Engine *string
	
	// The version number of the cache engine to be used for this cluster. To view the
	// supported cache engine versions, use the DescribeCacheEngineVersions operation.
	// Important: You can upgrade to a newer engine version (see Selecting a Cache
	// Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)
	// ), but you cannot downgrade to an earlier engine version. If you want to use an
	// earlier engine version, you must delete the existing cluster or replication
	// group and create it anew with the earlier engine version.
	EngineVersion *string
	
	// The network type you choose when modifying a cluster, either ipv4 | ipv6 . IPv6
	// is supported for workloads using Redis engine version 6.2 onward or Memcached
	// engine version 1.6.6 on all instances built on the Nitro system (http://aws.amazon.com/ec2/nitro/)
	// .
	IpDiscovery types.IpDiscovery
	
	// Specifies the destination, format and type of the logs.
	LogDeliveryConfigurations []types.LogDeliveryConfigurationRequest
	
	// Must be either ipv4 | ipv6 | dual_stack . IPv6 is supported for workloads using
	// Redis engine version 6.2 onward or Memcached engine version 1.6.6 on all
	// instances built on the Nitro system (http://aws.amazon.com/ec2/nitro/) .
	NetworkType types.NetworkType
	
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS)
	// topic to which notifications are sent. The Amazon SNS topic owner must be the
	// same as the cluster owner.
	NotificationTopicArn *string
	
	// The initial number of cache nodes that the cluster has. For clusters running
	// Redis, this value must be 1. For clusters running Memcached, this value must be
	// between 1 and 40. If you need more than 40 nodes for your Memcached cluster,
	// please fill out the ElastiCache Limit Increase Request form at
	// http://aws.amazon.com/contact-us/elasticache-node-limit-request/ (http://aws.amazon.com/contact-us/elasticache-node-limit-request/)
	// .
	NumCacheNodes *int32
	
	// Specifies whether the nodes in the cluster are created in a single outpost or
	// across multiple outposts.
	OutpostMode types.OutpostMode
	
	// The port number on which each of the cache nodes accepts connections.
	Port *int32
	
	// The EC2 Availability Zone in which the cluster is created. All nodes belonging
	// to this cluster are placed in the preferred Availability Zone. If you want to
	// create your nodes across multiple Availability Zones, use
	// PreferredAvailabilityZones . Default: System chosen Availability Zone.
	PreferredAvailabilityZone *string
	
	// A list of the Availability Zones in which cache nodes are created. The order of
	// the zones in the list is not important. This option is only supported on
	// Memcached. If you are creating your cluster in an Amazon VPC (recommended) you
	// can only locate nodes in Availability Zones that are associated with the subnets
	// in the selected subnet group. The number of Availability Zones listed must equal
	// the value of NumCacheNodes . If you want all the nodes in the same Availability
	// Zone, use PreferredAvailabilityZone instead, or repeat the Availability Zone
	// multiple times in the list. Default: System chosen Availability Zones.
	PreferredAvailabilityZones []string
	
	// Specifies the weekly time range during which maintenance on the cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period.
	PreferredMaintenanceWindow *string
	
	// The outpost ARN in which the cache cluster is created.
	PreferredOutpostArn *string
	
	// The outpost ARNs in which the cache cluster is created.
	PreferredOutpostArns []string
	
	// The ID of the replication group to which this cluster should belong. If this
	// parameter is specified, the cluster is added to the specified replication group
	// as a read replica; otherwise, the cluster is a standalone primary that is not
	// part of any replication group. If the specified replication group is Multi-AZ
	// enabled and the Availability Zone is not specified, the cluster is created in
	// Availability Zones that provide the best spread of read replicas across
	// Availability Zones. This parameter is only valid if the Engine parameter is
	// redis .
	ReplicationGroupId *string
	
	// One or more VPC security groups associated with the cluster. Use this parameter
	// only when you are creating a cluster in an Amazon Virtual Private Cloud (Amazon
	// VPC).
	SecurityGroupIds []string
	
	// A single-element string list containing an Amazon Resource Name (ARN) that
	// uniquely identifies a Redis RDB snapshot file stored in Amazon S3. The snapshot
	// file is used to populate the node group (shard). The Amazon S3 object name in
	// the ARN cannot contain any commas. This parameter is only valid if the Engine
	// parameter is redis . Example of an Amazon S3 ARN:
	// arn:aws:s3:::my_bucket/snapshot1.rdb
	SnapshotArns []string
	
	// The name of a Redis snapshot from which to restore data into the new node group
	// (shard). The snapshot status changes to restoring while the new node group
	// (shard) is being created. This parameter is only valid if the Engine parameter
	// is redis .
	SnapshotName *string
	
	// The number of days for which ElastiCache retains automatic snapshots before
	// deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot
	// taken today is retained for 5 days before being deleted. This parameter is only
	// valid if the Engine parameter is redis . Default: 0 (i.e., automatic backups are
	// disabled for this cache cluster).
	SnapshotRetentionLimit *int32
	
	// The daily time range (in UTC) during which ElastiCache begins taking a daily
	// snapshot of your node group (shard). Example: 05:00-09:00 If you do not specify
	// this parameter, ElastiCache automatically chooses an appropriate time range.
	// This parameter is only valid if the Engine parameter is redis .
	SnapshotWindow *string
	
	// A list of tags to be added to this resource.
	Tags []types.Tag
	
	// A flag that enables in-transit encryption when set to true.
	TransitEncryptionEnabled *bool
	
	noSmithyDocumentSerde
}

type CreateCacheClusterOutput struct {
	
	// Contains all of the attributes of a specific cluster.
	CacheCluster *types.CacheCluster
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCacheClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCacheCluster{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCacheCluster{}, middleware.After)
	if err != nil { return err }
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
	return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
	return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
	return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
	return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
	return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
	return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
	return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
	return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
	return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
	return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
	return err
	}
	if err = addCreateCacheClusterResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateCacheClusterValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCacheCluster(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
	return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
	return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
	return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
	return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateCacheCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "elasticache",
	OperationName: "CreateCacheCluster",
	}
}

type opCreateCacheClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateCacheClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateCacheClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	if m.EndpointResolver == nil {
	        return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	    }
	
	params := EndpointParameters{}
	
	m.BuiltInResolver.ResolveBuiltIns(&params)
	
	var resolvedEndpoint smithyendpoints.Endpoint
	    resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	    if err != nil {
	        return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	    }
	
	    req.URL = &resolvedEndpoint.URI
	
	    for k := range resolvedEndpoint.Headers {
	        req.Header.Set(
	            k,
	            resolvedEndpoint.Headers.Get(k),
	        )
	    }
	
	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	    if err != nil {
	        var nfe *internalauth.NoAuthenticationSchemesFoundError
	        if errors.As(err, &nfe) {
	            // if no auth scheme is found, default to sigv4
	            signingName := "elasticache"
	            signingRegion := m.BuiltInResolver.(*builtInResolver).Region
	            ctx = awsmiddleware.SetSigningName(ctx, signingName)
	            ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
	
	        }
	        var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
	        if errors.As(err, &ue) {
	            return out, metadata, fmt.Errorf(
	                "This operation requests signer version(s) %v but the client only supports %v",
	                ue.UnsupportedSchemes,
	                internalauth.SupportedSchemes,
	            )
	        }
	    }
	
	for _, authScheme := range authSchemes {
	    switch authScheme.(type) {
	        case *internalauth.AuthenticationSchemeV4:
	            v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
	    var signingName, signingRegion string
	    if v4Scheme.SigningName == nil {
	        signingName = "elasticache"
	    } else {
	        signingName = *v4Scheme.SigningName
	    }
	    if v4Scheme.SigningRegion == nil {
	        signingRegion = m.BuiltInResolver.(*builtInResolver).Region
	    } else {
	        signingRegion = *v4Scheme.SigningRegion
	    }
	    if v4Scheme.DisableDoubleEncoding != nil {
	        // The signer sets an equivalent value at client initialization time.
	        // Setting this context value will cause the signer to extract it
	        // and override the value set at client initialization time.
	        ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
	    }
	    ctx = awsmiddleware.SetSigningName(ctx, signingName)
	    ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
	            break
	        case *internalauth.AuthenticationSchemeV4A:
	            v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
	    if v4aScheme.SigningName == nil {
	        v4aScheme.SigningName = aws.String("elasticache")
	    }
	    if v4aScheme.DisableDoubleEncoding != nil {
	        // The signer sets an equivalent value at client initialization time.
	        // Setting this context value will cause the signer to extract it
	        // and override the value set at client initialization time.
	        ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
	    }
	    ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
	    ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
	            break
	        case *internalauth.AuthenticationSchemeNone:
	            break
	    }
	}
	
	return next.HandleSerialize(ctx, in)
}

func addCreateCacheClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateCacheClusterResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
