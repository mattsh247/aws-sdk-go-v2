// Code generated by smithy-go-codegen DO NOT EDIT.


package kendra

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
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Searches an index given an input query. You can configure boosting or relevance
// tuning at the query level to override boosting at the index level, filter based
// on document fields/attributes and faceted search, and filter based on the user
// or their group access to documents. You can also include certain fields in the
// response that might provide useful additional information. A query response
// contains three types of results.
//   - Relevant suggested answers. The answers can be either a text excerpt or
//   table excerpt. The answer can be highlighted in the excerpt.
//   - Matching FAQs or questions-answer from your FAQ file.
//   - Relevant documents. This result type includes an excerpt of the document
//   with the document title. The searched terms can be highlighted in the excerpt.
// You can specify that the query return only one type of result using the
// QueryResultTypeFilter parameter. Each query returns the 100 most relevant
// results. If you filter result type to only question-answers, a maximum of four
// results are returned. If you filter result type to only answers, a maximum of
// three results are returned.
func (c *Client) Query(ctx context.Context, params *QueryInput, optFns ...func(*Options)) (*QueryOutput, error) {
	if params == nil { params = &QueryInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "Query", params, optFns, c.addOperationQueryMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*QueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryInput struct {
	
	// The identifier of the index for the search.
	//
	// This member is required.
	IndexId *string
	
	// Filters search results by document fields/attributes. You can only provide one
	// attribute filter; however, the AndAllFilters , NotFilter , and OrAllFilters
	// parameters contain a list of other filters. The AttributeFilter parameter means
	// you can create a set of filtering rules that a document must satisfy to be
	// included in the query results.
	AttributeFilter *types.AttributeFilter
	
	// Overrides relevance tuning configurations of fields/attributes set at the index
	// level. If you use this API to override the relevance tuning configured at the
	// index level, but there is no relevance tuning configured at the index level,
	// then Amazon Kendra does not apply any relevance tuning. If there is relevance
	// tuning configured for fields at the index level, and you use this API to
	// override only some of these fields, then for the fields you did not override,
	// the importance is set to 1.
	DocumentRelevanceOverrideConfigurations []types.DocumentRelevanceConfiguration
	
	// An array of documents fields/attributes for faceted search. Amazon Kendra
	// returns a count for each field key specified. This helps your users narrow their
	// search.
	Facets []types.Facet
	
	// Query results are returned in pages the size of the PageSize parameter. By
	// default, Amazon Kendra returns the first page of results. Use this parameter to
	// get result pages after the first one.
	PageNumber *int32
	
	// Sets the number of results that are returned in each page of results. The
	// default page size is 10. The maximum number of results returned is 100. If you
	// ask for more than 100 results, only 100 are returned.
	PageSize *int32
	
	// Sets the type of query result or response. Only results for the specified type
	// are returned.
	QueryResultTypeFilter types.QueryResultType
	
	// The input query text for the search. Amazon Kendra truncates queries at 30
	// token words, which excludes punctuation and stop words. Truncation still applies
	// if you use Boolean or more advanced, complex queries.
	QueryText *string
	
	// An array of document fields/attributes to include in the response. You can
	// limit the response to include certain document fields. By default, all document
	// attributes are included in the response.
	RequestedDocumentAttributes []string
	
	// Provides information that determines how the results of the query are sorted.
	// You can set the field that Amazon Kendra should sort the results on, and specify
	// whether the results should be sorted in ascending or descending order. In the
	// case of ties in sorting the results, the results are sorted by relevance. If you
	// don't provide sorting configuration, the results are sorted by the relevance
	// that Amazon Kendra determines for the result.
	SortingConfiguration *types.SortingConfiguration
	
	// Enables suggested spell corrections for queries.
	SpellCorrectionConfiguration *types.SpellCorrectionConfiguration
	
	// The user context token or user and group information.
	UserContext *types.UserContext
	
	// Provides an identifier for a specific user. The VisitorId should be a unique
	// identifier, such as a GUID. Don't use personally identifiable information, such
	// as the user's email address, as the VisitorId .
	VisitorId *string
	
	noSmithyDocumentSerde
}

type QueryOutput struct {
	
	// Contains the facet results. A FacetResult contains the counts for each
	// field/attribute key that was specified in the Facets input parameter.
	FacetResults []types.FacetResult
	
	// The list of featured result items. Featured results are displayed at the top of
	// the search results page, placed above all other results for certain queries. If
	// there's an exact match of a query, then certain documents are featured in the
	// search results.
	FeaturedResultsItems []types.FeaturedResultsItem
	
	// The identifier for the search. You also use QueryId to identify the search when
	// using the SubmitFeedback (https://docs.aws.amazon.com/kendra/latest/APIReference/API_SubmitFeedback.html)
	// API.
	QueryId *string
	
	// The results of the search.
	ResultItems []types.QueryResultItem
	
	// A list of information related to suggested spell corrections for a query.
	SpellCorrectedQueries []types.SpellCorrectedQuery
	
	// The total number of items found by the search. However, you can only retrieve
	// up to 100 items. For example, if the search found 192 items, you can only
	// retrieve the first 100 of the items.
	TotalNumberOfResults *int32
	
	// A list of warning codes and their messages on problems with your query. Amazon
	// Kendra currently only supports one type of warning, which is a warning on
	// invalid syntax used in the query. For examples of invalid query syntax, see
	// Searching with advanced query syntax (https://docs.aws.amazon.com/kendra/latest/dg/searching-example.html#searching-index-query-syntax)
	// .
	Warnings []types.Warning
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpQuery{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpQuery{}, middleware.After)
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
	if err = addQueryResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpQueryValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opQuery(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "kendra",
	OperationName: "Query",
	}
}

type opQueryResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opQueryResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opQueryResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "kendra"
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
	        signingName = "kendra"
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
	        v4aScheme.SigningName = aws.String("kendra")
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

func addQueryResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opQueryResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
