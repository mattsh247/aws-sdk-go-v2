// Code generated by smithy-go-codegen DO NOT EDIT.


package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/v4a"
)

// Lists the S3 Intelligent-Tiering configuration from the specified bucket. The
// S3 Intelligent-Tiering storage class is designed to optimize storage costs by
// automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities. The S3
// Intelligent-Tiering storage class is the ideal storage class for data with
// unknown, changing, or unpredictable access patterns, independent of object size
// or retention period. If the size of an object is less than 128 KB, it is not
// monitored and not eligible for auto-tiering. Smaller objects can be stored, but
// they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class. For more information, see Storage class for
// automatically optimizing frequently and infrequently accessed objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access)
// . Operations related to ListBucketIntelligentTieringConfigurations include:
//   - DeleteBucketIntelligentTieringConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html)
//   - PutBucketIntelligentTieringConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html)
//   - GetBucketIntelligentTieringConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html)
func (c *Client) ListBucketIntelligentTieringConfigurations(ctx context.Context, params *ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*Options)) (*ListBucketIntelligentTieringConfigurationsOutput, error) {
	if params == nil { params = &ListBucketIntelligentTieringConfigurationsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListBucketIntelligentTieringConfigurations", params, optFns, c.addOperationListBucketIntelligentTieringConfigurationsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListBucketIntelligentTieringConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBucketIntelligentTieringConfigurationsInput struct {
	
	// The name of the Amazon S3 bucket whose configuration you want to modify or
	// retrieve.
	//
	// This member is required.
	Bucket *string
	
	// The ContinuationToken that represents a placeholder from where this request
	// should begin.
	ContinuationToken *string
	
	noSmithyDocumentSerde
}

type ListBucketIntelligentTieringConfigurationsOutput struct {
	
	// The ContinuationToken that represents a placeholder from where this request
	// should begin.
	ContinuationToken *string
	
	// The list of S3 Intelligent-Tiering configurations for a bucket.
	IntelligentTieringConfigurationList []types.IntelligentTieringConfiguration
	
	// Indicates whether the returned list of analytics configurations is complete. A
	// value of true indicates that the list is not complete and the
	// NextContinuationToken will be provided for a subsequent request.
	IsTruncated bool
	
	// The marker used to continue this inventory configuration listing. Use the
	// NextContinuationToken from this response to continue the listing in a subsequent
	// request. The continuation token is an opaque value that Amazon S3 understands.
	NextContinuationToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListBucketIntelligentTieringConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListBucketIntelligentTieringConfigurations{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListBucketIntelligentTieringConfigurations{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
	return err
	}
	if err = addListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListBucketIntelligentTieringConfigurationsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBucketIntelligentTieringConfigurations(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
	return err
	}
	if err = addListBucketIntelligentTieringConfigurationsUpdateEndpoint(stack, options); err != nil {
	return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
	return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
	return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
	return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
	return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
	return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
	return err
	}
	return nil
}

func (v *ListBucketIntelligentTieringConfigurationsInput) bucket() (string, bool) {
    if v.Bucket == nil {
        return "", false
    }
    return *v.Bucket, true
}

func newServiceMetadataMiddleware_opListBucketIntelligentTieringConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "s3",
	OperationName: "ListBucketIntelligentTieringConfigurations",
	}
}

// getListBucketIntelligentTieringConfigurationsBucketMember returns a pointer to
// string denoting a provided bucket member valueand a boolean indicating if the
// input has a modeled bucket name,
func getListBucketIntelligentTieringConfigurationsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListBucketIntelligentTieringConfigurationsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListBucketIntelligentTieringConfigurationsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
	Accessor : s3cust.UpdateEndpointParameterAccessor{
	GetBucketFromInput: getListBucketIntelligentTieringConfigurationsBucketMember,
	},
	UsePathStyle: options.UsePathStyle,
	UseAccelerate: options.UseAccelerate,
	SupportsAccelerate: true,
	TargetS3ObjectLambda: false,
	EndpointResolver: options.EndpointResolver,
	EndpointResolverOptions: options.EndpointOptions,
	UseARNRegion: options.UseARNRegion,
	DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	input, ok := in.Parameters.(*ListBucketIntelligentTieringConfigurationsInput)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	if m.EndpointResolver == nil {
	        return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	    }
	
	params := EndpointParameters{}
	
	m.BuiltInResolver.ResolveBuiltIns(&params)
	
	params.Bucket = input.Bucket
	
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
	            signingName := "s3"
	            signingRegion := m.BuiltInResolver.(*builtInResolver).Region
	            ctx = awsmiddleware.SetSigningName(ctx, signingName)
	            ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
	            ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
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
	        signingName = "s3"
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
	    ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
	            break
	        case *internalauth.AuthenticationSchemeV4A:
	            v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
	    if v4aScheme.SigningName == nil {
	        v4aScheme.SigningName = aws.String("s3")
	    }
	    if v4aScheme.DisableDoubleEncoding != nil {
	        // The signer sets an equivalent value at client initialization time.
	        // Setting this context value will cause the signer to extract it
	        // and override the value set at client initialization time.
	        ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
	    }
	    ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
	    ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
	    ctx = s3cust.SetSignerVersion(ctx, v4a.Version)
	            break
	        case *internalauth.AuthenticationSchemeNone:
	            break
	    }
	}
	
	return next.HandleSerialize(ctx, in)
}

func addListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListBucketIntelligentTieringConfigurationsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
Endpoint: options.BaseEndpoint,
ForcePathStyle: options.UsePathStyle,
Accelerate: options.UseAccelerate,
DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
UseArnRegion: options.UseARNRegion,
    },

    }, "ResolveEndpoint", middleware.After)
}
