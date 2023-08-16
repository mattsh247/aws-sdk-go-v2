// Code generated by smithy-go-codegen DO NOT EDIT.


package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalChecksum "github.com/aws/aws-sdk-go-v2/service/internal/checksum"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/internal/v4a"
)

// This action enables you to delete multiple objects from a bucket using a single
// HTTP request. If you know the object keys that you want to delete, then this
// action provides a suitable alternative to sending individual delete requests,
// reducing per-request overhead. The request contains a list of up to 1000 keys
// that you want to delete. In the XML, you provide the object key names, and
// optionally, version IDs if you want to delete a specific version of the object
// from a versioning-enabled bucket. For each key, Amazon S3 performs a delete
// action and returns the result of that delete, success, or failure, in the
// response. Note that if the object specified in the request is not found, Amazon
// S3 returns the result as deleted. The action supports two modes for the
// response: verbose and quiet. By default, the action uses verbose mode in which
// the response includes the result of deletion of each key in your request. In
// quiet mode the response includes only keys where the delete action encountered
// an error. For a successful deletion, the action does not return any information
// about the delete in the response body. When performing this action on an MFA
// Delete enabled bucket, that attempts to delete any versioned objects, you must
// include an MFA token. If you do not provide one, the entire request will fail,
// even if there are non-versioned objects you are trying to delete. If you provide
// an invalid token, whether there are versioned keys in the request or not, the
// entire Multi-Object Delete request will fail. For information about MFA Delete,
// see MFA Delete (https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete)
// . Finally, the Content-MD5 header is required for all Multi-Object Delete
// requests. Amazon S3 uses the header value to ensure that your request body has
// not been altered in transit. The following operations are related to
// DeleteObjects :
//   - CreateMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//   - UploadPart (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//   - CompleteMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//   - ListParts (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html)
//   - AbortMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
func (c *Client) DeleteObjects(ctx context.Context, params *DeleteObjectsInput, optFns ...func(*Options)) (*DeleteObjectsOutput, error) {
	if params == nil { params = &DeleteObjectsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DeleteObjects", params, optFns, c.addOperationDeleteObjectsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DeleteObjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteObjectsInput struct {
	
	// The bucket name containing the objects to delete. When using this action with
	// an access point, you must direct requests to the access point hostname. The
	// access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts? (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string
	
	// Container for the request.
	//
	// This member is required.
	Delete *types.Delete
	
	// Specifies whether you want to delete this object even if it has a
	// Governance-type Object Lock in place. To use this header, you must have the
	// s3:BypassGovernanceRetention permission.
	BypassGovernanceRetention bool
	
	// Indicates the algorithm used to create the checksum for the object when using
	// the SDK. This header will not provide any additional functionality if not using
	// the SDK. When sending this header, there must be a corresponding x-amz-checksum
	// or x-amz-trailer header sent. Otherwise, Amazon S3 fails the request with the
	// HTTP status code 400 Bad Request . For more information, see Checking object
	// integrity (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide. If you provide an individual checksum, Amazon S3
	// ignores any provided ChecksumAlgorithm parameter. This checksum algorithm must
	// be the same for all parts and it match the checksum value supplied in the
	// CreateMultipartUpload request.
	ChecksumAlgorithm types.ChecksumAlgorithm
	
	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string
	
	// The concatenation of the authentication device's serial number, a space, and
	// the value that is displayed on your authentication device. Required to
	// permanently delete a versioned object if versioning is configured with MFA
	// delete enabled.
	MFA *string
	
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer
	
	noSmithyDocumentSerde
}

type DeleteObjectsOutput struct {
	
	// Container element for a successful delete. It identifies the object that was
	// successfully deleted.
	Deleted []types.DeletedObject
	
	// Container for a failed delete action that describes the object that Amazon S3
	// attempted to delete and the error it encountered.
	Errors []types.Error
	
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteObjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteObjects{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteObjects{}, middleware.After)
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
	if err = addDeleteObjectsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDeleteObjectsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteObjects(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
	return err
	}
	if err = addDeleteObjectsInputChecksumMiddlewares(stack, options); err != nil {
	return err
	}
	if err = addDeleteObjectsUpdateEndpoint(stack, options); err != nil {
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

func (v *DeleteObjectsInput) bucket() (string, bool) {
    if v.Bucket == nil {
        return "", false
    }
    return *v.Bucket, true
}

func newServiceMetadataMiddleware_opDeleteObjects(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "s3",
	OperationName: "DeleteObjects",
	}
}

// getDeleteObjectsRequestAlgorithmMember gets the request checksum algorithm
// value provided as input.
func getDeleteObjectsRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*DeleteObjectsInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addDeleteObjectsInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
	    GetAlgorithm: getDeleteObjectsRequestAlgorithmMember,
	    RequireChecksum: true,
	    EnableTrailingChecksum: false,
	    EnableComputeSHA256PayloadHash: true,
	    EnableDecodedContentLengthHeader: true,
	})
}

// getDeleteObjectsBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getDeleteObjectsBucketMember(input interface{}) (*string, bool) {
	in := input.(*DeleteObjectsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addDeleteObjectsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
	Accessor : s3cust.UpdateEndpointParameterAccessor{
	GetBucketFromInput: getDeleteObjectsBucketMember,
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

type opDeleteObjectsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDeleteObjectsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteObjectsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	input, ok := in.Parameters.(*DeleteObjectsInput)
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

func addDeleteObjectsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDeleteObjectsResolveEndpointMiddleware{
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
