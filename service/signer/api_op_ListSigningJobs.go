// Code generated by smithy-go-codegen DO NOT EDIT.


package signer

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
	"time"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Lists all your signing jobs. You can use the maxResults parameter to limit the
// number of signing jobs that are returned in the response. If additional jobs
// remain to be listed, code signing returns a nextToken value. Use this value in
// subsequent calls to ListSigningJobs to fetch the remaining values. You can
// continue calling ListSigningJobs with your maxResults parameter and with new
// values that code signing returns in the nextToken parameter until all of your
// signing jobs have been returned.
func (c *Client) ListSigningJobs(ctx context.Context, params *ListSigningJobsInput, optFns ...func(*Options)) (*ListSigningJobsOutput, error) {
	if params == nil { params = &ListSigningJobsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListSigningJobs", params, optFns, c.addOperationListSigningJobsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListSigningJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningJobsInput struct {
	
	// Filters results to return only signing jobs with revoked signatures.
	IsRevoked bool
	
	// Filters results to return only signing jobs initiated by a specified IAM entity.
	JobInvoker *string
	
	// Specifies the maximum number of items to return in the response. Use this
	// parameter when paginating results. If additional items exist beyond the number
	// you specify, the nextToken element is set in the response. Use the nextToken
	// value in a subsequent request to retrieve additional items.
	MaxResults *int32
	
	// String for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string
	
	// The ID of microcontroller platform that you specified for the distribution of
	// your code image.
	PlatformId *string
	
	// The IAM principal that requested the signing job.
	RequestedBy *string
	
	// Filters results to return only signing jobs with signatures expiring after a
	// specified timestamp.
	SignatureExpiresAfter *time.Time
	
	// Filters results to return only signing jobs with signatures expiring before a
	// specified timestamp.
	SignatureExpiresBefore *time.Time
	
	// A status value with which to filter your results.
	Status types.SigningStatus
	
	noSmithyDocumentSerde
}

type ListSigningJobsOutput struct {
	
	// A list of your signing jobs.
	Jobs []types.SigningJob
	
	// String for specifying the next set of paginated results.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListSigningJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSigningJobs{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSigningJobs{}, middleware.After)
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
	if err = addListSigningJobsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningJobs(options.Region, ), middleware.Before); err != nil {
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

// ListSigningJobsAPIClient is a client that implements the ListSigningJobs
// operation.
type ListSigningJobsAPIClient interface {
	ListSigningJobs(context.Context, *ListSigningJobsInput, ...func(*Options)) (*ListSigningJobsOutput, error)
}

var _ ListSigningJobsAPIClient = (*Client)(nil)

// ListSigningJobsPaginatorOptions is the paginator options for ListSigningJobs
type ListSigningJobsPaginatorOptions struct {
	// Specifies the maximum number of items to return in the response. Use this
	// parameter when paginating results. If additional items exist beyond the number
	// you specify, the nextToken element is set in the response. Use the nextToken
	// value in a subsequent request to retrieve additional items.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSigningJobsPaginator is a paginator for ListSigningJobs
type ListSigningJobsPaginator struct {
    options ListSigningJobsPaginatorOptions
    client ListSigningJobsAPIClient
    params *ListSigningJobsInput
    nextToken *string
    firstPage bool
}

// NewListSigningJobsPaginator returns a new ListSigningJobsPaginator
func NewListSigningJobsPaginator(client ListSigningJobsAPIClient, params *ListSigningJobsInput, optFns ...func(*ListSigningJobsPaginatorOptions)) *ListSigningJobsPaginator {
	if params == nil {
	    params = &ListSigningJobsInput{}
	}
	
	options := ListSigningJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListSigningJobsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSigningJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListSigningJobs page.
func (p *ListSigningJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSigningJobsOutput, error) {
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
	
	result, err := p.client.ListSigningJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSigningJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "signer",
	OperationName: "ListSigningJobs",
	}
}

type opListSigningJobsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListSigningJobsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListSigningJobsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "signer"
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
	        signingName = "signer"
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
	        v4aScheme.SigningName = aws.String("signer")
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

func addListSigningJobsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListSigningJobsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
