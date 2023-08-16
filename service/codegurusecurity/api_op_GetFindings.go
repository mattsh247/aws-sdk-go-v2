// Code generated by smithy-go-codegen DO NOT EDIT.


package codegurusecurity

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
	"github.com/aws/aws-sdk-go-v2/service/codegurusecurity/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns a list of all findings generated by a particular scan.
func (c *Client) GetFindings(ctx context.Context, params *GetFindingsInput, optFns ...func(*Options)) (*GetFindingsOutput, error) {
	if params == nil { params = &GetFindingsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetFindings", params, optFns, c.addOperationGetFindingsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsInput struct {
	
	// The name of the scan you want to retrieve findings from.
	//
	// This member is required.
	ScanName *string
	
	// The maximum number of results to return in the response. Use this parameter
	// when paginating results. If additional results exist beyond the number you
	// specify, the nextToken element is returned in the response. Use nextToken in a
	// subsequent request to retrieve additional results.
	MaxResults *int32
	
	// A token to use for paginating results that are returned in the response. Set
	// the value of this parameter to null for the first request. For subsequent calls,
	// use the nextToken value returned from the previous request to continue listing
	// results after the first page.
	NextToken *string
	
	// The status of the findings you want to get. Pass either Open , Closed , or All .
	Status types.Status
	
	noSmithyDocumentSerde
}

type GetFindingsOutput struct {
	
	// A list of findings generated by the specified scan.
	Findings []types.Finding
	
	// A pagination token. You can use this in future calls to GetFindings to continue
	// listing results after the current page.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindings{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindings{}, middleware.After)
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
	if err = addGetFindingsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpGetFindingsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindings(options.Region, ), middleware.Before); err != nil {
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

// GetFindingsAPIClient is a client that implements the GetFindings operation.
type GetFindingsAPIClient interface {
	GetFindings(context.Context, *GetFindingsInput, ...func(*Options)) (*GetFindingsOutput, error)
}

var _ GetFindingsAPIClient = (*Client)(nil)

// GetFindingsPaginatorOptions is the paginator options for GetFindings
type GetFindingsPaginatorOptions struct {
	// The maximum number of results to return in the response. Use this parameter
	// when paginating results. If additional results exist beyond the number you
	// specify, the nextToken element is returned in the response. Use nextToken in a
	// subsequent request to retrieve additional results.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFindingsPaginator is a paginator for GetFindings
type GetFindingsPaginator struct {
    options GetFindingsPaginatorOptions
    client GetFindingsAPIClient
    params *GetFindingsInput
    nextToken *string
    firstPage bool
}

// NewGetFindingsPaginator returns a new GetFindingsPaginator
func NewGetFindingsPaginator(client GetFindingsAPIClient, params *GetFindingsInput, optFns ...func(*GetFindingsPaginatorOptions)) *GetFindingsPaginator {
	if params == nil {
	    params = &GetFindingsInput{}
	}
	
	options := GetFindingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &GetFindingsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next GetFindings page.
func (p *GetFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFindingsOutput, error) {
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
	
	result, err := p.client.GetFindings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "codeguru-security",
	OperationName: "GetFindings",
	}
}

type opGetFindingsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetFindingsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetFindingsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "codeguru-security"
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
	        signingName = "codeguru-security"
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
	        v4aScheme.SigningName = aws.String("codeguru-security")
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

func addGetFindingsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetFindingsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
