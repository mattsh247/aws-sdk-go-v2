// Code generated by smithy-go-codegen DO NOT EDIT.


package elastictranscoder

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
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// The ListJobsByPipeline operation gets a list of the jobs currently in a
// pipeline. Elastic Transcoder returns all of the jobs currently in the specified
// pipeline. The response body contains one element for each job that satisfies the
// search criteria.
func (c *Client) ListJobsByPipeline(ctx context.Context, params *ListJobsByPipelineInput, optFns ...func(*Options)) (*ListJobsByPipelineOutput, error) {
	if params == nil { params = &ListJobsByPipelineInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListJobsByPipeline", params, optFns, c.addOperationListJobsByPipelineMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListJobsByPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListJobsByPipelineRequest structure.
type ListJobsByPipelineInput struct {
	
	// The ID of the pipeline for which you want to get job information.
	//
	// This member is required.
	PipelineId *string
	
	// To list jobs in chronological order by the date and time that they were
	// submitted, enter true . To list jobs in reverse chronological order, enter false
	// .
	Ascending *string
	
	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
	
	noSmithyDocumentSerde
}

// The ListJobsByPipelineResponse structure.
type ListJobsByPipelineOutput struct {
	
	// An array of Job objects that are in the specified pipeline.
	Jobs []types.Job
	
	// A value that you use to access the second and subsequent pages of results, if
	// any. When the jobs in the specified pipeline fit on one page or when you've
	// reached the last page of results, the value of NextPageToken is null .
	NextPageToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobsByPipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobsByPipeline{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobsByPipeline{}, middleware.After)
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
	if err = addListJobsByPipelineResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListJobsByPipelineValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobsByPipeline(options.Region, ), middleware.Before); err != nil {
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

// ListJobsByPipelineAPIClient is a client that implements the ListJobsByPipeline
// operation.
type ListJobsByPipelineAPIClient interface {
	ListJobsByPipeline(context.Context, *ListJobsByPipelineInput, ...func(*Options)) (*ListJobsByPipelineOutput, error)
}

var _ ListJobsByPipelineAPIClient = (*Client)(nil)

// ListJobsByPipelinePaginatorOptions is the paginator options for
// ListJobsByPipeline
type ListJobsByPipelinePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobsByPipelinePaginator is a paginator for ListJobsByPipeline
type ListJobsByPipelinePaginator struct {
    options ListJobsByPipelinePaginatorOptions
    client ListJobsByPipelineAPIClient
    params *ListJobsByPipelineInput
    nextToken *string
    firstPage bool
}

// NewListJobsByPipelinePaginator returns a new ListJobsByPipelinePaginator
func NewListJobsByPipelinePaginator(client ListJobsByPipelineAPIClient, params *ListJobsByPipelineInput, optFns ...func(*ListJobsByPipelinePaginatorOptions)) *ListJobsByPipelinePaginator {
	if params == nil {
	    params = &ListJobsByPipelineInput{}
	}
	
	options := ListJobsByPipelinePaginatorOptions{}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListJobsByPipelinePaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobsByPipelinePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListJobsByPipeline page.
func (p *ListJobsByPipelinePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobsByPipelineOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.PageToken = p.nextToken
	
	result, err := p.client.ListJobsByPipeline(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobsByPipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "elastictranscoder",
	OperationName: "ListJobsByPipeline",
	}
}

type opListJobsByPipelineResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListJobsByPipelineResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListJobsByPipelineResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "elastictranscoder"
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
	        signingName = "elastictranscoder"
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
	        v4aScheme.SigningName = aws.String("elastictranscoder")
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

func addListJobsByPipelineResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListJobsByPipelineResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
