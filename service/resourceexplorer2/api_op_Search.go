// Code generated by smithy-go-codegen DO NOT EDIT.


package resourceexplorer2

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
	"github.com/aws/aws-sdk-go-v2/service/resourceexplorer2/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Searches for resources and displays details about all resources that match the
// specified criteria. You must specify a query string. All search queries must use
// a view. If you don't explicitly specify a view, then Amazon Web Services
// Resource Explorer uses the default view for the Amazon Web Services Region in
// which you call this operation. The results are the logical intersection of the
// results that match both the QueryString parameter supplied to this operation
// and the SearchFilter parameter attached to the view. For the complete syntax
// supported by the QueryString parameter, see Search query syntax reference for
// Resource Explorer (https://docs.aws.amazon.com/resource-explorer/latest/APIReference/about-query-syntax.html)
// . If your search results are empty, or are missing results that you think should
// be there, see Troubleshooting Resource Explorer search (https://docs.aws.amazon.com/resource-explorer/latest/userguide/troubleshooting_search.html)
// .
func (c *Client) Search(ctx context.Context, params *SearchInput, optFns ...func(*Options)) (*SearchOutput, error) {
	if params == nil { params = &SearchInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "Search", params, optFns, c.addOperationSearchMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*SearchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchInput struct {
	
	// A string that includes keywords and filters that specify the resources that you
	// want to include in the results. For the complete syntax supported by the
	// QueryString parameter, see Search query syntax reference for Resource Explorer (https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html)
	// . The search is completely case insensitive. You can specify an empty string to
	// return all results up to the limit of 1,000 total results. The operation can
	// return only the first 1,000 results. If the resource you want is not included,
	// then use a different value for QueryString to refine the results.
	//
	// This member is required.
	QueryString *string
	
	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32
	
	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
	
	// Specifies the Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view to use for the query. If you don't specify a value for this
	// parameter, then the operation automatically uses the default view for the Amazon
	// Web Services Region in which you called this operation. If the Region either
	// doesn't have a default view or if you don't have permission to use the default
	// view, then the operation fails with a 401 Unauthorized exception.
	ViewArn *string
	
	noSmithyDocumentSerde
}

type SearchOutput struct {
	
	// The number of resources that match the query.
	Count *types.ResourceCount
	
	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string
	
	// The list of structures that describe the resources that match the query.
	Resources []types.Resource
	
	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view that this operation used to perform the search.
	ViewArn *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearch{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearch{}, middleware.After)
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
	if err = addSearchResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpSearchValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearch(options.Region, ), middleware.Before); err != nil {
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

// SearchAPIClient is a client that implements the Search operation.
type SearchAPIClient interface {
	Search(context.Context, *SearchInput, ...func(*Options)) (*SearchOutput, error)
}

var _ SearchAPIClient = (*Client)(nil)

// SearchPaginatorOptions is the paginator options for Search
type SearchPaginatorOptions struct {
	// The maximum number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value
	// appropriate to the operation. If additional items exist beyond those included in
	// the current response, the NextToken response element is present and has a value
	// (is not null). Include that value as the NextToken request parameter in the
	// next call to the operation to get the next part of the results. An API operation
	// can return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchPaginator is a paginator for Search
type SearchPaginator struct {
    options SearchPaginatorOptions
    client SearchAPIClient
    params *SearchInput
    nextToken *string
    firstPage bool
}

// NewSearchPaginator returns a new SearchPaginator
func NewSearchPaginator(client SearchAPIClient, params *SearchInput, optFns ...func(*SearchPaginatorOptions)) *SearchPaginator {
	if params == nil {
	    params = &SearchInput{}
	}
	
	options := SearchPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &SearchPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next Search page.
func (p *SearchPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.NextToken = p.nextToken
	
	var limit *int32
	if p.options.Limit > 0 {
	    limit = &p.options.Limit
	}
	params.MaxResults = limit
	
	result, err := p.client.Search(ctx, &params, optFns...)
	if err != nil {
	    return nil, err
	}
	p.firstPage = false
	
	prevToken := p.nextToken
	p.nextToken = result.NextToken
	
	if p.options.StopOnDuplicateToken &&
	    prevToken != nil &&
	    p.nextToken != nil &&
	    *prevToken == *p.nextToken {
	    p.nextToken = nil
	}
	
	return result, nil
}

func newServiceMetadataMiddleware_opSearch(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "resource-explorer-2",
	OperationName: "Search",
	}
}

type opSearchResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opSearchResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSearchResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "resource-explorer-2"
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
	        signingName = "resource-explorer-2"
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
	        v4aScheme.SigningName = aws.String("resource-explorer-2")
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

func addSearchResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opSearchResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
