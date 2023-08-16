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

// Returns metadata about all versions of the objects in a bucket. You can also
// use request parameters as selection criteria to return metadata about a subset
// of all the object versions. To use this operation, you must have permission to
// perform the s3:ListBucketVersions action. Be aware of the name difference. A
// 200 OK response can contain valid or invalid XML. Make sure to design your
// application to parse the contents of the response and handle it appropriately.
// To use this operation, you must have READ access to the bucket. This action is
// not supported by Amazon S3 on Outposts. The following operations are related to
// ListObjectVersions :
//   - ListObjectsV2 (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html)
//   - GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//   - PutObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//   - DeleteObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html)
func (c *Client) ListObjectVersions(ctx context.Context, params *ListObjectVersionsInput, optFns ...func(*Options)) (*ListObjectVersionsOutput, error) {
	if params == nil { params = &ListObjectVersionsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListObjectVersions", params, optFns, c.addOperationListObjectVersionsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListObjectVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectVersionsInput struct {
	
	// The bucket name that contains the objects.
	//
	// This member is required.
	Bucket *string
	
	// A delimiter is a character that you specify to group keys. All keys that
	// contain the same string between the prefix and the first occurrence of the
	// delimiter are grouped under a single result element in CommonPrefixes . These
	// groups are counted as one result against the max-keys limitation. These keys
	// are not returned elsewhere in the response.
	Delimiter *string
	
	// Requests Amazon S3 to encode the object keys in the response and specifies the
	// encoding method to use. An object key can contain any Unicode character;
	// however, the XML 1.0 parser cannot parse some characters, such as characters
	// with an ASCII value from 0 to 10. For characters that are not supported in XML
	// 1.0, you can add this parameter to request that Amazon S3 encode the keys in the
	// response.
	EncodingType types.EncodingType
	
	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string
	
	// Specifies the key to start with when listing objects in a bucket.
	KeyMarker *string
	
	// Sets the maximum number of keys returned in the response. By default, the
	// action returns up to 1,000 key names. The response might contain fewer keys but
	// will never contain more. If additional keys satisfy the search criteria, but
	// were not returned because max-keys was exceeded, the response contains true . To
	// return the additional keys, see key-marker and version-id-marker .
	MaxKeys int32
	
	// Specifies the optional fields that you want returned in the response. Fields
	// that you do not specify are not returned.
	OptionalObjectAttributes []types.OptionalObjectAttributes
	
	// Use this parameter to select only those keys that begin with the specified
	// prefix. You can use prefixes to separate a bucket into different groupings of
	// keys. (You can think of using prefix to make groups in the same way that you'd
	// use a folder in a file system.) You can use prefix with delimiter to roll up
	// numerous objects into a single result under CommonPrefixes .
	Prefix *string
	
	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer
	
	// Specifies the object version you want to start listing from.
	VersionIdMarker *string
	
	noSmithyDocumentSerde
}

type ListObjectVersionsOutput struct {
	
	// All of the keys rolled up into a common prefix count as a single return when
	// calculating the number of returns.
	CommonPrefixes []types.CommonPrefix
	
	// Container for an object that is a delete marker.
	DeleteMarkers []types.DeleteMarkerEntry
	
	// The delimiter grouping the included keys. A delimiter is a character that you
	// specify to group keys. All keys that contain the same string between the prefix
	// and the first occurrence of the delimiter are grouped under a single result
	// element in CommonPrefixes . These groups are counted as one result against the
	// max-keys limitation. These keys are not returned elsewhere in the response.
	Delimiter *string
	
	// Encoding type used by Amazon S3 to encode object key names in the XML response.
	// If you specify the encoding-type request parameter, Amazon S3 includes this
	// element in the response, and returns encoded key name values in the following
	// response elements: KeyMarker, NextKeyMarker, Prefix, Key , and Delimiter .
	EncodingType types.EncodingType
	
	// A flag that indicates whether Amazon S3 returned all of the results that
	// satisfied the search criteria. If your results were truncated, you can make a
	// follow-up paginated request by using the NextKeyMarker and NextVersionIdMarker
	// response parameters as a starting place in another request to return the rest of
	// the results.
	IsTruncated bool
	
	// Marks the last key returned in a truncated response.
	KeyMarker *string
	
	// Specifies the maximum number of objects to return.
	MaxKeys int32
	
	// The bucket name.
	Name *string
	
	// When the number of responses exceeds the value of MaxKeys , NextKeyMarker
	// specifies the first key not returned that satisfies the search criteria. Use
	// this value for the key-marker request parameter in a subsequent request.
	NextKeyMarker *string
	
	// When the number of responses exceeds the value of MaxKeys , NextVersionIdMarker
	// specifies the first object version not returned that satisfies the search
	// criteria. Use this value for the version-id-marker request parameter in a
	// subsequent request.
	NextVersionIdMarker *string
	
	// Selects objects that start with the value supplied by this parameter.
	Prefix *string
	
	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged
	
	// Marks the last version of the key returned in a truncated response.
	VersionIdMarker *string
	
	// Container for version information.
	Versions []types.ObjectVersion
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListObjectVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListObjectVersions{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListObjectVersions{}, middleware.After)
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
	if err = addListObjectVersionsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListObjectVersionsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectVersions(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
	return err
	}
	if err = addListObjectVersionsUpdateEndpoint(stack, options); err != nil {
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

func (v *ListObjectVersionsInput) bucket() (string, bool) {
    if v.Bucket == nil {
        return "", false
    }
    return *v.Bucket, true
}

func newServiceMetadataMiddleware_opListObjectVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "s3",
	OperationName: "ListObjectVersions",
	}
}

// getListObjectVersionsBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getListObjectVersionsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListObjectVersionsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListObjectVersionsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
	Accessor : s3cust.UpdateEndpointParameterAccessor{
	GetBucketFromInput: getListObjectVersionsBucketMember,
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

type opListObjectVersionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListObjectVersionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListObjectVersionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	input, ok := in.Parameters.(*ListObjectVersionsInput)
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

func addListObjectVersionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListObjectVersionsResolveEndpointMiddleware{
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
