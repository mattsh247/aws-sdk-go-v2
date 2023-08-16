// Code generated by smithy-go-codegen DO NOT EDIT.


package databasemigrationservice

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
	"time"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Starts the replication task. For more information about DMS tasks, see Working
// with Migration Tasks  (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.html)
// in the Database Migration Service User Guide.
func (c *Client) StartReplicationTask(ctx context.Context, params *StartReplicationTaskInput, optFns ...func(*Options)) (*StartReplicationTaskOutput, error) {
	if params == nil { params = &StartReplicationTaskInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "StartReplicationTask", params, optFns, c.addOperationStartReplicationTaskMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*StartReplicationTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskInput struct {
	
	// The Amazon Resource Name (ARN) of the replication task to be started.
	//
	// This member is required.
	ReplicationTaskArn *string
	
	// The type of replication task to start. When the migration type is full-load or
	// full-load-and-cdc , the only valid value for the first run of the task is
	// start-replication . This option will start the migration. You can also use
	// ReloadTables to reload specific tables that failed during migration instead of
	// restarting the task. The resume-processing option isn't applicable for a
	// full-load task, because you can't resume partially loaded tables during the full
	// load phase. For a full-load-and-cdc task, DMS migrates table data, and then
	// applies data changes that occur on the source. To load all the tables again, and
	// start capturing source changes, use reload-target . Otherwise use
	// resume-processing , to replicate the changes from the last stop position.
	//
	// This member is required.
	StartReplicationTaskType types.StartReplicationTaskTypeValue
	
	// Indicates when you want a change data capture (CDC) operation to start. Use
	// either CdcStartPosition or CdcStartTime to specify when you want a CDC operation
	// to start. Specifying both values results in an error. The value can be in date,
	// checkpoint, or LSN/SCN format. Date Example: --cdc-start-position
	// “2018-03-08T12:12:12” Checkpoint Example: --cdc-start-position
	// "checkpoint:V1#27#mysql-bin-changelog.157832:1975:-1:2002:677883278264080:mysql-bin-changelog.157832:1876#0#0#*#0#93"
	// LSN Example: --cdc-start-position “mysql-bin-changelog.000024:373” When you use
	// this task setting with a source PostgreSQL database, a logical replication slot
	// should already be created and associated with the source endpoint. You can
	// verify this by setting the slotName extra connection attribute to the name of
	// this logical replication slot. For more information, see Extra Connection
	// Attributes When Using PostgreSQL as a Source for DMS (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.PostgreSQL.html#CHAP_Source.PostgreSQL.ConnectionAttrib)
	// .
	CdcStartPosition *string
	
	// Indicates the start time for a change data capture (CDC) operation. Use either
	// CdcStartTime or CdcStartPosition to specify when you want a CDC operation to
	// start. Specifying both values results in an error. Timestamp Example:
	// --cdc-start-time “2018-03-08T12:12:12”
	CdcStartTime *time.Time
	
	// Indicates when you want a change data capture (CDC) operation to stop. The
	// value can be either server time or commit time. Server time example:
	// --cdc-stop-position “server_time:2018-02-09T12:12:12” Commit time example:
	// --cdc-stop-position “commit_time: 2018-02-09T12:12:12“
	CdcStopPosition *string
	
	noSmithyDocumentSerde
}

//
type StartReplicationTaskOutput struct {
	
	// The replication task started.
	ReplicationTask *types.ReplicationTask
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationStartReplicationTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTask{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTask{}, middleware.After)
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
	if err = addStartReplicationTaskResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpStartReplicationTaskValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTask(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartReplicationTask(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "dms",
	OperationName: "StartReplicationTask",
	}
}

type opStartReplicationTaskResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opStartReplicationTaskResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartReplicationTaskResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "dms"
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
	        signingName = "dms"
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
	        v4aScheme.SigningName = aws.String("dms")
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

func addStartReplicationTaskResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opStartReplicationTaskResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
