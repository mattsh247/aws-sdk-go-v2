// Code generated by smithy-go-codegen DO NOT EDIT.


package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// The GetItem operation returns a set of attributes for the item with the given
// primary key. If there is no matching item, GetItem does not return any data and
// there will be no Item element in the response. GetItem provides an eventually
// consistent read by default. If your application requires a strongly consistent
// read, set ConsistentRead to true . Although a strongly consistent read might
// take more time than an eventually consistent read, it always returns the last
// updated value.
func (c *Client) GetItem(ctx context.Context, params *GetItemInput, optFns ...func(*Options)) (*GetItemOutput, error) {
	if params == nil { params = &GetItemInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetItem", params, optFns, c.addOperationGetItemMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetItem operation.
type GetItemInput struct {
	
	// A map of attribute names to AttributeValue objects, representing the primary
	// key of the item to retrieve. For the primary key, you must provide all of the
	// attributes. For example, with a simple primary key, you only need to provide a
	// value for the partition key. For a composite primary key, you must provide
	// values for both the partition key and the sort key.
	//
	// This member is required.
	Key map[string]types.AttributeValue
	
	// The name of the table containing the requested item.
	//
	// This member is required.
	TableName *string
	
	// This is a legacy parameter. Use ProjectionExpression instead. For more
	// information, see AttributesToGet (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html)
	// in the Amazon DynamoDB Developer Guide.
	AttributesToGet []string
	
	// Determines the read consistency model: If set to true , then the operation uses
	// strongly consistent reads; otherwise, the operation uses eventually consistent
	// reads.
	ConsistentRead *bool
	
	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames :
	//   - To access an attribute whose name conflicts with a DynamoDB reserved word.
	//   - To create a placeholder for repeating occurrences of an attribute name in
	//   an expression.
	//   - To prevent special characters in an attribute name from being
	//   misinterpreted in an expression.
	// Use the # character in an expression to dereference an attribute name. For
	// example, consider the following attribute name:
	//   - Percentile
	// The name of this attribute conflicts with a reserved word, so it cannot be used
	// directly in an expression. (For the complete list of reserved words, see
	// Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames :
	//   - {"#P":"Percentile"}
	// You could then use this substitution in an expression, as in this example:
	//   - #P = :val
	// Tokens that begin with the : character are expression attribute values, which
	// are placeholders for the actual value at runtime. For more information on
	// expression attribute names, see Specifying Item Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]string
	
	// A string that identifies one or more attributes to retrieve from the table.
	// These attributes can include scalars, sets, or elements of a JSON document. The
	// attributes in the expression must be separated by commas. If no attribute names
	// are specified, then all attributes are returned. If any of the requested
	// attributes are not found, they do not appear in the result. For more
	// information, see Specifying Item Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ProjectionExpression *string
	
	// Determines the level of detail about either provisioned or on-demand throughput
	// consumption that is returned in the response:
	//   - INDEXES - The response includes the aggregate ConsumedCapacity for the
	//   operation, together with ConsumedCapacity for each table and secondary index
	//   that was accessed. Note that some operations, such as GetItem and BatchGetItem
	//   , do not access any indexes at all. In these cases, specifying INDEXES will
	//   only return ConsumedCapacity information for table(s).
	//   - TOTAL - The response includes only the aggregate ConsumedCapacity for the
	//   operation.
	//   - NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity
	
	noSmithyDocumentSerde
}

// Represents the output of a GetItem operation.
type GetItemOutput struct {
	
	// The capacity units consumed by the GetItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics for
	// the table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity
	
	// A map of attribute names to AttributeValue objects, as specified by
	// ProjectionExpression .
	Item map[string]types.AttributeValue
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetItem{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetItem{}, middleware.After)
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
	if err = addOpGetItemDiscoverEndpointMiddleware(stack, options, c); err != nil {
	return err
	}
	if err = addGetItemResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpGetItemValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetItem(options.Region, ), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
	return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
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

func addOpGetItemDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func (opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation: c.fetchOpGetItemDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired: false,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpGetItemDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*GetItemInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in
	
	identifierMap := make(map[string]string, 0)
	
	key := fmt.Sprintf("DynamoDB.%v", identifierMap)
	
	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}
	
	discoveryOperationInput := &DescribeEndpointsInput{
	}
	
	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}
	
	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opGetItem(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "dynamodb",
	OperationName: "GetItem",
	}
}

type opGetItemResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetItemResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetItemResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "dynamodb"
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
	        signingName = "dynamodb"
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
	        v4aScheme.SigningName = aws.String("dynamodb")
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

func addGetItemResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetItemResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
