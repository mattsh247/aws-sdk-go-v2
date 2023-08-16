// Code generated by smithy-go-codegen DO NOT EDIT.


package servicecatalog

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
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Lists all PrincipalARN s and corresponding PrincipalType s associated with the
// specified portfolio.
func (c *Client) ListPrincipalsForPortfolio(ctx context.Context, params *ListPrincipalsForPortfolioInput, optFns ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error) {
	if params == nil { params = &ListPrincipalsForPortfolioInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListPrincipalsForPortfolio", params, optFns, c.addOperationListPrincipalsForPortfolioMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListPrincipalsForPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsForPortfolioInput struct {
	
	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string
	
	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string
	
	// The maximum number of items to return with this call.
	PageSize int32
	
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
	
	noSmithyDocumentSerde
}

type ListPrincipalsForPortfolioOutput struct {
	
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string
	
	// The PrincipalARN s and corresponding PrincipalType s associated with the
	// portfolio.
	Principals []types.Principal
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrincipalsForPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPrincipalsForPortfolio{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPrincipalsForPortfolio{}, middleware.After)
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
	if err = addListPrincipalsForPortfolioResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListPrincipalsForPortfolioValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalsForPortfolio(options.Region, ), middleware.Before); err != nil {
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

// ListPrincipalsForPortfolioAPIClient is a client that implements the
// ListPrincipalsForPortfolio operation.
type ListPrincipalsForPortfolioAPIClient interface {
	ListPrincipalsForPortfolio(context.Context, *ListPrincipalsForPortfolioInput, ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error)
}

var _ ListPrincipalsForPortfolioAPIClient = (*Client)(nil)

// ListPrincipalsForPortfolioPaginatorOptions is the paginator options for
// ListPrincipalsForPortfolio
type ListPrincipalsForPortfolioPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrincipalsForPortfolioPaginator is a paginator for
// ListPrincipalsForPortfolio
type ListPrincipalsForPortfolioPaginator struct {
    options ListPrincipalsForPortfolioPaginatorOptions
    client ListPrincipalsForPortfolioAPIClient
    params *ListPrincipalsForPortfolioInput
    nextToken *string
    firstPage bool
}

// NewListPrincipalsForPortfolioPaginator returns a new
// ListPrincipalsForPortfolioPaginator
func NewListPrincipalsForPortfolioPaginator(client ListPrincipalsForPortfolioAPIClient, params *ListPrincipalsForPortfolioInput, optFns ...func(*ListPrincipalsForPortfolioPaginatorOptions)) *ListPrincipalsForPortfolioPaginator {
	if params == nil {
	    params = &ListPrincipalsForPortfolioInput{}
	}
	
	options := ListPrincipalsForPortfolioPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListPrincipalsForPortfolioPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrincipalsForPortfolioPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListPrincipalsForPortfolio page.
func (p *ListPrincipalsForPortfolioPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.PageToken = p.nextToken
	
	params.PageSize = p.options.Limit
	
	result, err := p.client.ListPrincipalsForPortfolio(ctx, &params, optFns...)
	if err != nil {
	    return nil, err
	}
	p.firstPage = false
	
	prevToken := p.nextToken
	p.nextToken = result.NextPageToken
	
	if p.options.StopOnDuplicateToken &&
	    prevToken != nil &&
	    p.nextToken != nil &&
	    *prevToken == *p.nextToken {
	    p.nextToken = nil
	}
	
	return result, nil
}

func newServiceMetadataMiddleware_opListPrincipalsForPortfolio(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "servicecatalog",
	OperationName: "ListPrincipalsForPortfolio",
	}
}

type opListPrincipalsForPortfolioResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListPrincipalsForPortfolioResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListPrincipalsForPortfolioResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "servicecatalog"
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
	        signingName = "servicecatalog"
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
	        v4aScheme.SigningName = aws.String("servicecatalog")
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

func addListPrincipalsForPortfolioResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListPrincipalsForPortfolioResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
