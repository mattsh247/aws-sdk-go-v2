// Code generated by smithy-go-codegen DO NOT EDIT.


package glue

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
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Searches a set of tables based on properties in the table metadata as well as
// on the parent database. You can search against text or filter conditions. You
// can only get tables that you have access to based on the security policies
// defined in Lake Formation. You need at least a read-only access to the table for
// it to be returned. If you do not have access to all the columns in the table,
// these columns will not be searched against when returning the list of tables
// back to you. If you have access to the columns but not the data in the columns,
// those columns and the associated metadata for those columns will be included in
// the search.
func (c *Client) SearchTables(ctx context.Context, params *SearchTablesInput, optFns ...func(*Options)) (*SearchTablesOutput, error) {
	if params == nil { params = &SearchTablesInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "SearchTables", params, optFns, c.addOperationSearchTablesMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*SearchTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTablesInput struct {
	
	// A unique identifier, consisting of  account_id .
	CatalogId *string
	
	// A list of key-value pairs, and a comparator used to filter the search results.
	// Returns all entities matching the predicate. The Comparator member of the
	// PropertyPredicate struct is used only for time fields, and can be omitted for
	// other field types. Also, when comparing string values, such as when Key=Name , a
	// fuzzy match algorithm is used. The Key field (for example, the value of the Name
	// field) is split on certain punctuation characters, for example, -, :, #, etc.
	// into tokens. Then each token is exact-match compared with the Value member of
	// PropertyPredicate . For example, if Key=Name and Value=link , tables named
	// customer-link and xx-link-yy are returned, but xxlinkyy is not returned.
	Filters []types.PropertyPredicate
	
	// The maximum number of tables to return in a single response.
	MaxResults *int32
	
	// A continuation token, included if this is a continuation call.
	NextToken *string
	
	// Allows you to specify that you want to search the tables shared with your
	// account. The allowable values are FOREIGN or ALL .
	//   - If set to FOREIGN , will search the tables shared with your account.
	//   - If set to ALL , will search the tables shared with your account, as well as
	//   the tables in yor local account.
	ResourceShareType types.ResourceShareType
	
	// A string used for a text search. Specifying a value in quotes filters based on
	// an exact match to the value.
	SearchText *string
	
	// A list of criteria for sorting the results by a field name, in an ascending or
	// descending order.
	SortCriteria []types.SortCriterion
	
	noSmithyDocumentSerde
}

type SearchTablesOutput struct {
	
	// A continuation token, present if the current list segment is not the last.
	NextToken *string
	
	// A list of the requested Table objects. The SearchTables response returns only
	// the tables that you have access to.
	TableList []types.Table
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchTables{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchTables{}, middleware.After)
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
	if err = addSearchTablesResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTables(options.Region, ), middleware.Before); err != nil {
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

// SearchTablesAPIClient is a client that implements the SearchTables operation.
type SearchTablesAPIClient interface {
	SearchTables(context.Context, *SearchTablesInput, ...func(*Options)) (*SearchTablesOutput, error)
}

var _ SearchTablesAPIClient = (*Client)(nil)

// SearchTablesPaginatorOptions is the paginator options for SearchTables
type SearchTablesPaginatorOptions struct {
	// The maximum number of tables to return in a single response.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchTablesPaginator is a paginator for SearchTables
type SearchTablesPaginator struct {
    options SearchTablesPaginatorOptions
    client SearchTablesAPIClient
    params *SearchTablesInput
    nextToken *string
    firstPage bool
}

// NewSearchTablesPaginator returns a new SearchTablesPaginator
func NewSearchTablesPaginator(client SearchTablesAPIClient, params *SearchTablesInput, optFns ...func(*SearchTablesPaginatorOptions)) *SearchTablesPaginator {
	if params == nil {
	    params = &SearchTablesInput{}
	}
	
	options := SearchTablesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &SearchTablesPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchTablesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next SearchTables page.
func (p *SearchTablesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchTablesOutput, error) {
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
	
	result, err := p.client.SearchTables(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchTables(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "glue",
	OperationName: "SearchTables",
	}
}

type opSearchTablesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opSearchTablesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSearchTablesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "glue"
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
	        signingName = "glue"
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
	        v4aScheme.SigningName = aws.String("glue")
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

func addSearchTablesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opSearchTablesResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
