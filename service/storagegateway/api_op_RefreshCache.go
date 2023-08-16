// Code generated by smithy-go-codegen DO NOT EDIT.


package storagegateway

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
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Refreshes the cached inventory of objects for the specified file share. This
// operation finds objects in the Amazon S3 bucket that were added, removed, or
// replaced since the gateway last listed the bucket's contents and cached the
// results. This operation does not import files into the S3 File Gateway cache
// storage. It only updates the cached inventory to reflect changes in the
// inventory of the objects in the S3 bucket. This operation is only supported in
// the S3 File Gateway types. You can subscribe to be notified through an Amazon
// CloudWatch event when your RefreshCache operation completes. For more
// information, see Getting notified about file operations (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification)
// in the Storage Gateway User Guide. This operation is Only supported for S3 File
// Gateways. When this API is called, it only initiates the refresh operation. When
// the API call completes and returns a success code, it doesn't necessarily mean
// that the file refresh has completed. You should use the refresh-complete
// notification to determine that the operation has completed before you check for
// new files on the gateway file share. You can subscribe to be notified through a
// CloudWatch event when your RefreshCache operation completes. Throttle limit:
// This API is asynchronous, so the gateway will accept no more than two refreshes
// at any time. We recommend using the refresh-complete CloudWatch event
// notification before issuing additional requests. For more information, see
// Getting notified about file operations (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification)
// in the Storage Gateway User Guide.
//   - Wait at least 60 seconds between consecutive RefreshCache API requests.
//   - RefreshCache does not evict cache entries if invoked consecutively within
//   60 seconds of a previous RefreshCache request.
//   - If you invoke the RefreshCache API when two requests are already being
//   processed, any new request will cause an InvalidGatewayRequestException error
//   because too many requests were sent to the server.
// The S3 bucket name does not need to be included when entering the list of
// folders in the FolderList parameter. For more information, see Getting notified
// about file operations (https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html#get-notification)
// in the Storage Gateway User Guide.
func (c *Client) RefreshCache(ctx context.Context, params *RefreshCacheInput, optFns ...func(*Options)) (*RefreshCacheOutput, error) {
	if params == nil { params = &RefreshCacheInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "RefreshCache", params, optFns, c.addOperationRefreshCacheMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*RefreshCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// RefreshCacheInput
type RefreshCacheInput struct {
	
	// The Amazon Resource Name (ARN) of the file share you want to refresh.
	//
	// This member is required.
	FileShareARN *string
	
	// A comma-separated list of the paths of folders to refresh in the cache. The
	// default is [ "/" ]. The default refreshes objects and folders at the root of the
	// Amazon S3 bucket. If Recursive is set to true , the entire S3 bucket that the
	// file share has access to is refreshed.
	FolderList []string
	
	// A value that specifies whether to recursively refresh folders in the cache. The
	// refresh includes folders that were in the cache the last time the gateway listed
	// the folder's contents. If this value set to true , each folder that is listed in
	// FolderList is recursively updated. Otherwise, subfolders listed in FolderList
	// are not refreshed. Only objects that are in folders listed directly under
	// FolderList are found and used for the update. The default is true . Valid
	// Values: true | false
	Recursive *bool
	
	noSmithyDocumentSerde
}

// RefreshCacheOutput
type RefreshCacheOutput struct {
	
	// The Amazon Resource Name (ARN) of the file share.
	FileShareARN *string
	
	// The randomly generated ID of the notification that was sent. This ID is in UUID
	// format.
	NotificationId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationRefreshCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRefreshCache{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRefreshCache{}, middleware.After)
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
	if err = addRefreshCacheResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpRefreshCacheValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRefreshCache(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRefreshCache(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "storagegateway",
	OperationName: "RefreshCache",
	}
}

type opRefreshCacheResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opRefreshCacheResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRefreshCacheResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "storagegateway"
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
	        signingName = "storagegateway"
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
	        v4aScheme.SigningName = aws.String("storagegateway")
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

func addRefreshCacheResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opRefreshCacheResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
