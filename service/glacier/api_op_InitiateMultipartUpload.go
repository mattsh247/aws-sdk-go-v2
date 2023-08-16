// Code generated by smithy-go-codegen DO NOT EDIT.


package glacier

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// This operation initiates a multipart upload. Amazon S3 Glacier creates a
// multipart upload resource and returns its ID in the response. The multipart
// upload ID is used in subsequent requests to upload parts of an archive (see
// UploadMultipartPart ). When you initiate a multipart upload, you specify the
// part size in number of bytes. The part size must be a megabyte (1024 KB)
// multiplied by a power of 2-for example, 1048576 (1 MB), 2097152 (2 MB), 4194304
// (4 MB), 8388608 (8 MB), and so on. The minimum allowable part size is 1 MB, and
// the maximum is 4 GB. Every part you upload to this resource (see
// UploadMultipartPart ), except the last one, must have the same size. The last
// one can be the same size or smaller. For example, suppose you want to upload a
// 16.2 MB file. If you initiate the multipart upload with a part size of 4 MB, you
// will upload four parts of 4 MB each and one part of 0.2 MB. You don't need to
// know the size of the archive when you start a multipart upload because Amazon S3
// Glacier does not require you to specify the overall archive size. After you
// complete the multipart upload, Amazon S3 Glacier (Glacier) removes the multipart
// upload resource referenced by the ID. Glacier also removes the multipart upload
// resource if you cancel the multipart upload or it may be removed if there is no
// activity for a period of 24 hours. An AWS account has full permission to perform
// all operations (actions). However, AWS Identity and Access Management (IAM)
// users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see Access
// Control Using AWS Identity and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and underlying REST API, see Uploading Large
// Archives in Parts (Multipart Upload) (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html)
// and Initiate Multipart Upload (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-initiate-upload.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) InitiateMultipartUpload(ctx context.Context, params *InitiateMultipartUploadInput, optFns ...func(*Options)) (*InitiateMultipartUploadOutput, error) {
	if params == nil { params = &InitiateMultipartUploadInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "InitiateMultipartUpload", params, optFns, c.addOperationInitiateMultipartUploadMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*InitiateMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for initiating a multipart upload to an Amazon S3 Glacier
// vault.
type InitiateMultipartUploadInput struct {
	
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string
	
	// The name of the vault.
	//
	// This member is required.
	VaultName *string
	
	// The archive description that you are uploading in parts. The part size must be
	// a megabyte (1024 KB) multiplied by a power of 2, for example 1048576 (1 MB),
	// 2097152 (2 MB), 4194304 (4 MB), 8388608 (8 MB), and so on. The minimum allowable
	// part size is 1 MB, and the maximum is 4 GB (4096 MB).
	ArchiveDescription *string
	
	// The size of each part except the last, in bytes. The last part can be smaller
	// than this part size.
	PartSize *string
	
	noSmithyDocumentSerde
}

// The Amazon S3 Glacier response to your request.
type InitiateMultipartUploadOutput struct {
	
	// The relative URI path of the multipart upload ID Amazon S3 Glacier created.
	Location *string
	
	// The ID of the multipart upload. This value is also included as part of the
	// location.
	UploadId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationInitiateMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInitiateMultipartUpload{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInitiateMultipartUpload{}, middleware.After)
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
	if err = addInitiateMultipartUploadResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpInitiateMultipartUploadValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInitiateMultipartUpload(options.Region, ), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
	return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
	return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opInitiateMultipartUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "glacier",
	OperationName: "InitiateMultipartUpload",
	}
}

type opInitiateMultipartUploadResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opInitiateMultipartUploadResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opInitiateMultipartUploadResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "glacier"
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
	        signingName = "glacier"
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
	        v4aScheme.SigningName = aws.String("glacier")
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

func addInitiateMultipartUploadResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opInitiateMultipartUploadResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
