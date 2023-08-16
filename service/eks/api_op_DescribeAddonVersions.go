// Code generated by smithy-go-codegen DO NOT EDIT.


package eks

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
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Describes the versions for an add-on. Information such as the Kubernetes
// versions that you can use the add-on with, the owner , publisher , and the type
// of the add-on are returned.
func (c *Client) DescribeAddonVersions(ctx context.Context, params *DescribeAddonVersionsInput, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
	if params == nil { params = &DescribeAddonVersionsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeAddonVersions", params, optFns, c.addOperationDescribeAddonVersionsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeAddonVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddonVersionsInput struct {
	
	// The name of the add-on. The name must match one of the names returned by
	// ListAddons (https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html)
	// .
	AddonName *string
	
	// The Kubernetes versions that you can use the add-on with.
	KubernetesVersion *string
	
	// The maximum number of results to return.
	MaxResults *int32
	
	// The nextToken value returned from a previous paginated
	// DescribeAddonVersionsRequest where maxResults was used and the results exceeded
	// the value of that parameter. Pagination continues from the end of the previous
	// results that returned the nextToken value. This token should be treated as an
	// opaque identifier that is used only to retrieve the next items in a list and not
	// for other programmatic purposes.
	NextToken *string
	
	// The owner of the add-on. For valid owners , don't specify a value for this
	// property.
	Owners []string
	
	// The publisher of the add-on. For valid publishers , don't specify a value for
	// this property.
	Publishers []string
	
	// The type of the add-on. For valid types , don't specify a value for this
	// property.
	Types []string
	
	noSmithyDocumentSerde
}

type DescribeAddonVersionsOutput struct {
	
	// The list of available versions with Kubernetes version compatibility and other
	// properties.
	Addons []types.AddonInfo
	
	// The nextToken value returned from a previous paginated
	// DescribeAddonVersionsResponse where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. This token should be
	// treated as an opaque identifier that is used only to retrieve the next items in
	// a list and not for other programmatic purposes.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAddonVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAddonVersions{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAddonVersions{}, middleware.After)
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
	if err = addDescribeAddonVersionsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddonVersions(options.Region, ), middleware.Before); err != nil {
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

// DescribeAddonVersionsAPIClient is a client that implements the
// DescribeAddonVersions operation.
type DescribeAddonVersionsAPIClient interface {
	DescribeAddonVersions(context.Context, *DescribeAddonVersionsInput, ...func(*Options)) (*DescribeAddonVersionsOutput, error)
}

var _ DescribeAddonVersionsAPIClient = (*Client)(nil)

// DescribeAddonVersionsPaginatorOptions is the paginator options for
// DescribeAddonVersions
type DescribeAddonVersionsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddonVersionsPaginator is a paginator for DescribeAddonVersions
type DescribeAddonVersionsPaginator struct {
    options DescribeAddonVersionsPaginatorOptions
    client DescribeAddonVersionsAPIClient
    params *DescribeAddonVersionsInput
    nextToken *string
    firstPage bool
}

// NewDescribeAddonVersionsPaginator returns a new DescribeAddonVersionsPaginator
func NewDescribeAddonVersionsPaginator(client DescribeAddonVersionsAPIClient, params *DescribeAddonVersionsInput, optFns ...func(*DescribeAddonVersionsPaginatorOptions)) *DescribeAddonVersionsPaginator {
	if params == nil {
	    params = &DescribeAddonVersionsInput{}
	}
	
	options := DescribeAddonVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeAddonVersionsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddonVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeAddonVersions page.
func (p *DescribeAddonVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
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
	
	result, err := p.client.DescribeAddonVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAddonVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "eks",
	OperationName: "DescribeAddonVersions",
	}
}

type opDescribeAddonVersionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeAddonVersionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeAddonVersionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "eks"
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
	        signingName = "eks"
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
	        v4aScheme.SigningName = aws.String("eks")
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

func addDescribeAddonVersionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeAddonVersionsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
