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

// Deletes a single item in a table by primary key. You can perform a conditional
// delete operation that deletes the item if it exists, or if it has an expected
// attribute value. In addition to deleting an item, you can also return the item's
// attribute values in the same operation, using the ReturnValues parameter.
// Unless you specify conditions, the DeleteItem is an idempotent operation;
// running it multiple times on the same item or attribute does not result in an
// error response. Conditional deletes are useful for deleting items only if
// specific conditions are met. If those conditions are met, DynamoDB performs the
// delete. Otherwise, the item is not deleted.
func (c *Client) DeleteItem(ctx context.Context, params *DeleteItemInput, optFns ...func(*Options)) (*DeleteItemOutput, error) {
	if params == nil { params = &DeleteItemInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DeleteItem", params, optFns, c.addOperationDeleteItemMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DeleteItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteItem operation.
type DeleteItemInput struct {
	
	// A map of attribute names to AttributeValue objects, representing the primary
	// key of the item to delete. For the primary key, you must provide all of the key
	// attributes. For example, with a simple primary key, you only need to provide a
	// value for the partition key. For a composite primary key, you must provide
	// values for both the partition key and the sort key.
	//
	// This member is required.
	Key map[string]types.AttributeValue
	
	// The name of the table from which to delete the item.
	//
	// This member is required.
	TableName *string
	
	// A condition that must be satisfied in order for a conditional DeleteItem to
	// succeed. An expression can contain any of the following:
	//   - Functions: attribute_exists | attribute_not_exists | attribute_type |
	//   contains | begins_with | size These function names are case-sensitive.
	//   - Comparison operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//   - Logical operators: AND | OR | NOT
	// For more information about condition expressions, see Condition Expressions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string
	
	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see ConditionalOperator (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator types.ConditionalOperator
	
	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see Expected (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]types.ExpectedAttributeValue
	
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
	
	// One or more values that can be substituted in an expression. Use the : (colon)
	// character in an expression to dereference an attribute value. For example,
	// suppose that you wanted to check whether the value of the ProductStatus
	// attribute was one of the following: Available | Backordered | Discontinued You
	// would first need to specify ExpressionAttributeValues as follows: {
	// ":avail":{"S":"Available"}, ":back":{"S":"Backordered"},
	// ":disc":{"S":"Discontinued"} } You could then use these values in an expression,
	// such as this: ProductStatus IN (:avail, :back, :disc) For more information on
	// expression attribute values, see Condition Expressions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]types.AttributeValue
	
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
	
	// Determines whether item collection metrics are returned. If set to SIZE , the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the
	// default), no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics
	
	// Use ReturnValues if you want to get the item attributes as they appeared before
	// they were deleted. For DeleteItem , the valid values are:
	//   - NONE - If ReturnValues is not specified, or if its value is NONE , then
	//   nothing is returned. (This setting is the default for ReturnValues .)
	//   - ALL_OLD - The content of the old item is returned.
	// There is no additional cost associated with requesting a return value aside
	// from the small network and processing overhead of receiving a larger response.
	// No read capacity units are consumed. The ReturnValues parameter is used by
	// several DynamoDB operations; however, DeleteItem does not recognize any values
	// other than NONE or ALL_OLD .
	ReturnValues types.ReturnValue
	
	// An optional parameter that returns the item attributes for a DeleteItem
	// operation that failed a condition check. There is no additional cost associated
	// with requesting a return value aside from the small network and processing
	// overhead of receiving a larger response. No read capacity units are consumed.
	ReturnValuesOnConditionCheckFailure types.ReturnValuesOnConditionCheckFailure
	
	noSmithyDocumentSerde
}

// Represents the output of a DeleteItem operation.
type DeleteItemOutput struct {
	
	// A map of attribute names to AttributeValue objects, representing the item as it
	// appeared before the DeleteItem operation. This map appears in the response only
	// if ReturnValues was specified as ALL_OLD in the request.
	Attributes map[string]types.AttributeValue
	
	// The capacity units consumed by the DeleteItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics for
	// the table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity
	
	// Information about item collections, if any, that were affected by the DeleteItem
	// operation. ItemCollectionMetrics is only returned if the
	// ReturnItemCollectionMetrics parameter was specified. If the table does not have
	// any local secondary indexes, this information is not returned in the response.
	// Each ItemCollectionMetrics element consists of:
	//   - ItemCollectionKey - The partition key value of the item collection. This is
	//   the same as the partition key value of the item itself.
	//   - SizeEstimateRangeGB - An estimate of item collection size, in gigabytes.
	//   This value is a two-element array containing a lower bound and an upper bound
	//   for the estimate. The estimate includes the size of all the items in the table,
	//   plus the size of all attributes projected into all of the local secondary
	//   indexes on that table. Use this estimate to measure whether a local secondary
	//   index is approaching its size limit. The estimate is subject to change over
	//   time; therefore, do not rely on the precision or accuracy of the estimate.
	ItemCollectionMetrics *types.ItemCollectionMetrics
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteItem{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteItem{}, middleware.After)
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
	if err = addOpDeleteItemDiscoverEndpointMiddleware(stack, options, c); err != nil {
	return err
	}
	if err = addDeleteItemResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDeleteItemValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteItem(options.Region, ), middleware.Before); err != nil {
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

func addOpDeleteItemDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func (opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation: c.fetchOpDeleteItemDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired: false,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpDeleteItemDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*DeleteItemInput)
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

func newServiceMetadataMiddleware_opDeleteItem(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "dynamodb",
	OperationName: "DeleteItem",
	}
}

type opDeleteItemResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDeleteItemResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteItemResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDeleteItemResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDeleteItemResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
