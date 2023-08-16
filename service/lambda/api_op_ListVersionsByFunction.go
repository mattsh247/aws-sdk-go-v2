// Code generated by smithy-go-codegen DO NOT EDIT.


package lambda

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
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns a list of versions (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html)
// , with the version-specific configuration of each. Lambda returns up to 50
// versions per call.
func (c *Client) ListVersionsByFunction(ctx context.Context, params *ListVersionsByFunctionInput, optFns ...func(*Options)) (*ListVersionsByFunctionOutput, error) {
	if params == nil { params = &ListVersionsByFunctionInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListVersionsByFunction", params, optFns, c.addOperationListVersionsByFunctionMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListVersionsByFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVersionsByFunctionInput struct {
	
	// The name of the Lambda function. Name formats
	//   - Function name - MyFunction .
	//   - Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction .
	//   - Partial ARN - 123456789012:function:MyFunction .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string
	
	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string
	
	// The maximum number of versions to return. Note that ListVersionsByFunction
	// returns a maximum of 50 items in each response, even if you set the number
	// higher.
	MaxItems *int32
	
	noSmithyDocumentSerde
}

type ListVersionsByFunctionOutput struct {
	
	// The pagination token that's included if more results are available.
	NextMarker *string
	
	// A list of Lambda function versions.
	Versions []types.FunctionConfiguration
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListVersionsByFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVersionsByFunction{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVersionsByFunction{}, middleware.After)
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
	if err = addListVersionsByFunctionResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListVersionsByFunctionValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVersionsByFunction(options.Region, ), middleware.Before); err != nil {
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

// ListVersionsByFunctionAPIClient is a client that implements the
// ListVersionsByFunction operation.
type ListVersionsByFunctionAPIClient interface {
	ListVersionsByFunction(context.Context, *ListVersionsByFunctionInput, ...func(*Options)) (*ListVersionsByFunctionOutput, error)
}

var _ ListVersionsByFunctionAPIClient = (*Client)(nil)

// ListVersionsByFunctionPaginatorOptions is the paginator options for
// ListVersionsByFunction
type ListVersionsByFunctionPaginatorOptions struct {
	// The maximum number of versions to return. Note that ListVersionsByFunction
	// returns a maximum of 50 items in each response, even if you set the number
	// higher.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVersionsByFunctionPaginator is a paginator for ListVersionsByFunction
type ListVersionsByFunctionPaginator struct {
    options ListVersionsByFunctionPaginatorOptions
    client ListVersionsByFunctionAPIClient
    params *ListVersionsByFunctionInput
    nextToken *string
    firstPage bool
}

// NewListVersionsByFunctionPaginator returns a new ListVersionsByFunctionPaginator
func NewListVersionsByFunctionPaginator(client ListVersionsByFunctionAPIClient, params *ListVersionsByFunctionInput, optFns ...func(*ListVersionsByFunctionPaginatorOptions)) *ListVersionsByFunctionPaginator {
	if params == nil {
	    params = &ListVersionsByFunctionInput{}
	}
	
	options := ListVersionsByFunctionPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListVersionsByFunctionPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVersionsByFunctionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListVersionsByFunction page.
func (p *ListVersionsByFunctionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVersionsByFunctionOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.Marker = p.nextToken
	
	var limit *int32
	if p.options.Limit > 0 {
	    limit = &p.options.Limit
	}
	params.MaxItems = limit
	
	result, err := p.client.ListVersionsByFunction(ctx, &params, optFns...)
	if err != nil {
	    return nil, err
	}
	p.firstPage = false
	
	prevToken := p.nextToken
	p.nextToken = result.NextMarker
	
	if p.options.StopOnDuplicateToken &&
	    prevToken != nil &&
	    p.nextToken != nil &&
	    *prevToken == *p.nextToken {
	    p.nextToken = nil
	}
	
	return result, nil
}

func newServiceMetadataMiddleware_opListVersionsByFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "lambda",
	OperationName: "ListVersionsByFunction",
	}
}

type opListVersionsByFunctionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListVersionsByFunctionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListVersionsByFunctionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "lambda"
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
	        signingName = "lambda"
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
	        v4aScheme.SigningName = aws.String("lambda")
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

func addListVersionsByFunctionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListVersionsByFunctionResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
