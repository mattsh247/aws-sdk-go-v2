// Code generated by smithy-go-codegen DO NOT EDIT.


package iotfleetwise

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
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// A list of information about signal decoders specified in a decoder manifest.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func (c *Client) ListDecoderManifestSignals(ctx context.Context, params *ListDecoderManifestSignalsInput, optFns ...func(*Options)) (*ListDecoderManifestSignalsOutput, error) {
	if params == nil { params = &ListDecoderManifestSignalsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListDecoderManifestSignals", params, optFns, c.addOperationListDecoderManifestSignalsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListDecoderManifestSignalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDecoderManifestSignalsInput struct {
	
	// The name of the decoder manifest to list information about.
	//
	// This member is required.
	Name *string
	
	// The maximum number of items to return, between 1 and 100, inclusive.
	MaxResults *int32
	
	// A pagination token for the next set of results. If the results of a search are
	// large, only a portion of the results are returned, and a nextToken pagination
	// token is returned in the response. To retrieve the next set of results, reissue
	// the search request and include the returned token. When all results have been
	// returned, the response does not contain a pagination token value.
	NextToken *string
	
	noSmithyDocumentSerde
}

type ListDecoderManifestSignalsOutput struct {
	
	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string
	
	// Information about a list of signals to decode.
	SignalDecoders []types.SignalDecoder
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListDecoderManifestSignalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDecoderManifestSignals{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDecoderManifestSignals{}, middleware.After)
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
	if err = addListDecoderManifestSignalsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListDecoderManifestSignalsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDecoderManifestSignals(options.Region, ), middleware.Before); err != nil {
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

// ListDecoderManifestSignalsAPIClient is a client that implements the
// ListDecoderManifestSignals operation.
type ListDecoderManifestSignalsAPIClient interface {
	ListDecoderManifestSignals(context.Context, *ListDecoderManifestSignalsInput, ...func(*Options)) (*ListDecoderManifestSignalsOutput, error)
}

var _ ListDecoderManifestSignalsAPIClient = (*Client)(nil)

// ListDecoderManifestSignalsPaginatorOptions is the paginator options for
// ListDecoderManifestSignals
type ListDecoderManifestSignalsPaginatorOptions struct {
	// The maximum number of items to return, between 1 and 100, inclusive.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDecoderManifestSignalsPaginator is a paginator for
// ListDecoderManifestSignals
type ListDecoderManifestSignalsPaginator struct {
    options ListDecoderManifestSignalsPaginatorOptions
    client ListDecoderManifestSignalsAPIClient
    params *ListDecoderManifestSignalsInput
    nextToken *string
    firstPage bool
}

// NewListDecoderManifestSignalsPaginator returns a new
// ListDecoderManifestSignalsPaginator
func NewListDecoderManifestSignalsPaginator(client ListDecoderManifestSignalsAPIClient, params *ListDecoderManifestSignalsInput, optFns ...func(*ListDecoderManifestSignalsPaginatorOptions)) *ListDecoderManifestSignalsPaginator {
	if params == nil {
	    params = &ListDecoderManifestSignalsInput{}
	}
	
	options := ListDecoderManifestSignalsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListDecoderManifestSignalsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDecoderManifestSignalsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListDecoderManifestSignals page.
func (p *ListDecoderManifestSignalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDecoderManifestSignalsOutput, error) {
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
	
	result, err := p.client.ListDecoderManifestSignals(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDecoderManifestSignals(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "iotfleetwise",
	OperationName: "ListDecoderManifestSignals",
	}
}

type opListDecoderManifestSignalsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListDecoderManifestSignalsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListDecoderManifestSignalsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "iotfleetwise"
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
	        signingName = "iotfleetwise"
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
	        v4aScheme.SigningName = aws.String("iotfleetwise")
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

func addListDecoderManifestSignalsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListDecoderManifestSignalsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
