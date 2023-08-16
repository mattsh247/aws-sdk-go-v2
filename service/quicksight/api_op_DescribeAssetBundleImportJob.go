// Code generated by smithy-go-codegen DO NOT EDIT.


package quicksight

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
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Describes an existing import job. Poll job descriptions after starting a job to
// know when it has succeeded or failed. Job descriptions are available for 14 days
// after job starts.
func (c *Client) DescribeAssetBundleImportJob(ctx context.Context, params *DescribeAssetBundleImportJobInput, optFns ...func(*Options)) (*DescribeAssetBundleImportJobOutput, error) {
	if params == nil { params = &DescribeAssetBundleImportJobInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetBundleImportJob", params, optFns, c.addOperationDescribeAssetBundleImportJobMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeAssetBundleImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetBundleImportJobInput struct {
	
	// The ID of the job. The job ID is set when you start a new job with a
	// StartAssetBundleImportJob API call.
	//
	// This member is required.
	AssetBundleImportJobId *string
	
	// The ID of the Amazon Web Services account the import job was executed in.
	//
	// This member is required.
	AwsAccountId *string
	
	noSmithyDocumentSerde
}

type DescribeAssetBundleImportJobOutput struct {
	
	// The Amazon Resource Name (ARN) for the import job.
	Arn *string
	
	// The ID of the job. The job ID is set when you start a new job with a
	// StartAssetBundleImportJob API call.
	AssetBundleImportJobId *string
	
	// The source of the asset bundle zip file that contains the data that is imported
	// by the job.
	AssetBundleImportSource *types.AssetBundleImportSourceDescription
	
	// The ID of the Amazon Web Services account the import job was executed in.
	AwsAccountId *string
	
	// The time that the import job was created.
	CreatedTime *time.Time
	
	// An array of error records that describes any failures that occurred during the
	// export job processing. Error records accumulate while the job is still running.
	// The complete set of error records is available after the job has completed and
	// failed.
	Errors []types.AssetBundleImportJobError
	
	// The failure action for the import job.
	FailureAction types.AssetBundleImportFailureAction
	
	// Indicates the status of a job through its queuing and execution. Poll the
	// DescribeAssetBundleImport API until JobStatus returns one of the following
	// values:
	//   - SUCCESSFUL
	//   - FAILED
	//   - FAILED_ROLLBACK_COMPLETED
	//   - FAILED_ROLLBACK_ERROR
	JobStatus types.AssetBundleImportJobStatus
	
	// Optional overrides to be applied to the resource configuration before import.
	OverrideParameters *types.AssetBundleImportJobOverrideParameters
	
	// The Amazon Web Services request ID for this operation.
	RequestId *string
	
	// An array of error records that describes any failures that occurred while an
	// import job was attempting a rollback. Error records accumulate while the job is
	// still running. The complete set of error records is available after the job has
	// completed and failed.
	RollbackErrors []types.AssetBundleImportJobError
	
	// The HTTP status of the response.
	Status int32
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssetBundleImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetBundleImportJob{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetBundleImportJob{}, middleware.After)
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
	if err = addDescribeAssetBundleImportJobResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeAssetBundleImportJobValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetBundleImportJob(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssetBundleImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "quicksight",
	OperationName: "DescribeAssetBundleImportJob",
	}
}

type opDescribeAssetBundleImportJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeAssetBundleImportJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeAssetBundleImportJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "quicksight"
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
	        signingName = "quicksight"
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
	        v4aScheme.SigningName = aws.String("quicksight")
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

func addDescribeAssetBundleImportJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeAssetBundleImportJobResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
