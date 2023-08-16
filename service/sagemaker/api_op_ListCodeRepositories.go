// Code generated by smithy-go-codegen DO NOT EDIT.


package sagemaker

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
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Gets a list of the Git repositories in your account.
func (c *Client) ListCodeRepositories(ctx context.Context, params *ListCodeRepositoriesInput, optFns ...func(*Options)) (*ListCodeRepositoriesOutput, error) {
	if params == nil { params = &ListCodeRepositoriesInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListCodeRepositories", params, optFns, c.addOperationListCodeRepositoriesMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListCodeRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCodeRepositoriesInput struct {
	
	// A filter that returns only Git repositories that were created after the
	// specified time.
	CreationTimeAfter *time.Time
	
	// A filter that returns only Git repositories that were created before the
	// specified time.
	CreationTimeBefore *time.Time
	
	// A filter that returns only Git repositories that were last modified after the
	// specified time.
	LastModifiedTimeAfter *time.Time
	
	// A filter that returns only Git repositories that were last modified before the
	// specified time.
	LastModifiedTimeBefore *time.Time
	
	// The maximum number of Git repositories to return in the response.
	MaxResults *int32
	
	// A string in the Git repositories name. This filter returns only repositories
	// whose name contains the specified string.
	NameContains *string
	
	// If the result of a ListCodeRepositoriesOutput request was truncated, the
	// response includes a NextToken . To get the next set of Git repositories, use the
	// token in the next request.
	NextToken *string
	
	// The field to sort results by. The default is Name .
	SortBy types.CodeRepositorySortBy
	
	// The sort order for results. The default is Ascending .
	SortOrder types.CodeRepositorySortOrder
	
	noSmithyDocumentSerde
}

type ListCodeRepositoriesOutput struct {
	
	// Gets a list of summaries of the Git repositories. Each summary specifies the
	// following values for the repository:
	//   - Name
	//   - Amazon Resource Name (ARN)
	//   - Creation time
	//   - Last modified time
	//   - Configuration information, including the URL location of the repository and
	//   the ARN of the Amazon Web Services Secrets Manager secret that contains the
	//   credentials used to access the repository.
	//
	// This member is required.
	CodeRepositorySummaryList []types.CodeRepositorySummary
	
	// If the result of a ListCodeRepositoriesOutput request was truncated, the
	// response includes a NextToken . To get the next set of Git repositories, use the
	// token in the next request.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListCodeRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCodeRepositories{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCodeRepositories{}, middleware.After)
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
	if err = addListCodeRepositoriesResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCodeRepositories(options.Region, ), middleware.Before); err != nil {
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

// ListCodeRepositoriesAPIClient is a client that implements the
// ListCodeRepositories operation.
type ListCodeRepositoriesAPIClient interface {
	ListCodeRepositories(context.Context, *ListCodeRepositoriesInput, ...func(*Options)) (*ListCodeRepositoriesOutput, error)
}

var _ ListCodeRepositoriesAPIClient = (*Client)(nil)

// ListCodeRepositoriesPaginatorOptions is the paginator options for
// ListCodeRepositories
type ListCodeRepositoriesPaginatorOptions struct {
	// The maximum number of Git repositories to return in the response.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCodeRepositoriesPaginator is a paginator for ListCodeRepositories
type ListCodeRepositoriesPaginator struct {
    options ListCodeRepositoriesPaginatorOptions
    client ListCodeRepositoriesAPIClient
    params *ListCodeRepositoriesInput
    nextToken *string
    firstPage bool
}

// NewListCodeRepositoriesPaginator returns a new ListCodeRepositoriesPaginator
func NewListCodeRepositoriesPaginator(client ListCodeRepositoriesAPIClient, params *ListCodeRepositoriesInput, optFns ...func(*ListCodeRepositoriesPaginatorOptions)) *ListCodeRepositoriesPaginator {
	if params == nil {
	    params = &ListCodeRepositoriesInput{}
	}
	
	options := ListCodeRepositoriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListCodeRepositoriesPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCodeRepositoriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListCodeRepositories page.
func (p *ListCodeRepositoriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCodeRepositoriesOutput, error) {
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
	
	result, err := p.client.ListCodeRepositories(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCodeRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "sagemaker",
	OperationName: "ListCodeRepositories",
	}
}

type opListCodeRepositoriesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListCodeRepositoriesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListCodeRepositoriesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "sagemaker"
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
	        signingName = "sagemaker"
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
	        v4aScheme.SigningName = aws.String("sagemaker")
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

func addListCodeRepositoriesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListCodeRepositoriesResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
