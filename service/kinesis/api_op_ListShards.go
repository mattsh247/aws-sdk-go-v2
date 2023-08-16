// Code generated by smithy-go-codegen DO NOT EDIT.


package kinesis

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Lists the shards in a stream and provides information about each shard. This
// operation has a limit of 1000 transactions per second per data stream. When
// invoking this API, it is recommended you use the StreamARN input parameter
// rather than the StreamName input parameter. This action does not list expired
// shards. For information about expired shards, see Data Routing, Data
// Persistence, and Shard State after a Reshard (https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing)
// . This API is a new operation that is used by the Amazon Kinesis Client Library
// (KCL). If you have a fine-grained IAM policy that only allows specific
// operations, you must update your policy to allow calls to this API. For more
// information, see Controlling Access to Amazon Kinesis Data Streams Resources
// Using IAM (https://docs.aws.amazon.com/streams/latest/dev/controlling-access.html)
// .
func (c *Client) ListShards(ctx context.Context, params *ListShardsInput, optFns ...func(*Options)) (*ListShardsOutput, error) {
	if params == nil { params = &ListShardsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListShards", params, optFns, c.addOperationListShardsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListShardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListShardsInput struct {
	
	// Specify this parameter to indicate that you want to list the shards starting
	// with the shard whose ID immediately follows ExclusiveStartShardId . If you don't
	// specify this parameter, the default behavior is for ListShards to list the
	// shards starting with the first one in the stream. You cannot specify this
	// parameter if you specify NextToken .
	ExclusiveStartShardId *string
	
	// The maximum number of shards to return in a single call to ListShards . The
	// maximum number of shards to return in a single call. The default value is 1000.
	// If you specify a value greater than 1000, at most 1000 results are returned.
	// When the number of shards to be listed is greater than the value of MaxResults ,
	// the response contains a NextToken value that you can use in a subsequent call
	// to ListShards to list the next set of shards.
	MaxResults *int32
	
	// When the number of shards in the data stream is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of shards in the data stream, the
	// response includes a pagination token named NextToken . You can specify this
	// NextToken value in a subsequent call to ListShards to list the next set of
	// shards. Don't specify StreamName or StreamCreationTimestamp if you specify
	// NextToken because the latter unambiguously identifies the stream. You can
	// optionally specify a value for the MaxResults parameter when you specify
	// NextToken . If you specify a MaxResults value that is less than the number of
	// shards that the operation returns if you don't specify MaxResults , the response
	// will contain a new NextToken value. You can use the new NextToken value in a
	// subsequent call to the ListShards operation. Tokens expire after 300 seconds.
	// When you obtain a value for NextToken in the response to a call to ListShards ,
	// you have 300 seconds to use that value. If you specify an expired token in a
	// call to ListShards , you get ExpiredNextTokenException .
	NextToken *string
	
	// Enables you to filter out the response of the ListShards API. You can only
	// specify one filter at a time. If you use the ShardFilter parameter when
	// invoking the ListShards API, the Type is the required property and must be
	// specified. If you specify the AT_TRIM_HORIZON , FROM_TRIM_HORIZON , or AT_LATEST
	// types, you do not need to specify either the ShardId or the Timestamp optional
	// properties. If you specify the AFTER_SHARD_ID type, you must also provide the
	// value for the optional ShardId property. The ShardId property is identical in
	// fuctionality to the ExclusiveStartShardId parameter of the ListShards API. When
	// ShardId property is specified, the response includes the shards starting with
	// the shard whose ID immediately follows the ShardId that you provided. If you
	// specify the AT_TIMESTAMP or FROM_TIMESTAMP_ID type, you must also provide the
	// value for the optional Timestamp property. If you specify the AT_TIMESTAMP
	// type, then all shards that were open at the provided timestamp are returned. If
	// you specify the FROM_TIMESTAMP type, then all shards starting from the provided
	// timestamp to TIP are returned.
	ShardFilter *types.ShardFilter
	
	// The ARN of the stream.
	StreamARN *string
	
	// Specify this input parameter to distinguish data streams that have the same
	// name. For example, if you create a data stream and then delete it, and you later
	// create another data stream with the same name, you can use this input parameter
	// to specify which of the two streams you want to list the shards for. You cannot
	// specify this parameter if you specify the NextToken parameter.
	StreamCreationTimestamp *time.Time
	
	// The name of the data stream whose shards you want to list. You cannot specify
	// this parameter if you specify the NextToken parameter.
	StreamName *string
	
	noSmithyDocumentSerde
}

type ListShardsOutput struct {
	
	// When the number of shards in the data stream is greater than the default value
	// for the MaxResults parameter, or if you explicitly specify a value for
	// MaxResults that is less than the number of shards in the data stream, the
	// response includes a pagination token named NextToken . You can specify this
	// NextToken value in a subsequent call to ListShards to list the next set of
	// shards. For more information about the use of this pagination token when calling
	// the ListShards operation, see ListShardsInput$NextToken . Tokens expire after
	// 300 seconds. When you obtain a value for NextToken in the response to a call to
	// ListShards , you have 300 seconds to use that value. If you specify an expired
	// token in a call to ListShards , you get ExpiredNextTokenException .
	NextToken *string
	
	// An array of JSON objects. Each object represents one shard and specifies the
	// IDs of the shard, the shard's parent, and the shard that's adjacent to the
	// shard's parent. Each object also contains the starting and ending hash keys and
	// the starting and ending sequence numbers for the shard.
	Shards []types.Shard
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListShardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListShards{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListShards{}, middleware.After)
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
	if err = addListShardsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListShardsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListShards(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListShards(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "kinesis",
	OperationName: "ListShards",
	}
}

type opListShardsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListShardsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListShardsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	input, ok := in.Parameters.(*ListShardsInput)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	if m.EndpointResolver == nil {
	        return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	    }
	
	params := EndpointParameters{}
	
	m.BuiltInResolver.ResolveBuiltIns(&params)
	
	params.StreamARN = input.StreamARN
	
	params.OperationType = ptr.String("control")
	
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
	            signingName := "kinesis"
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
	        signingName = "kinesis"
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
	        v4aScheme.SigningName = aws.String("kinesis")
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

func addListShardsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListShardsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
