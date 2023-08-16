// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing virtual routers in a service mesh.
func (c *Client) ListVirtualRouters(ctx context.Context, params *ListVirtualRoutersInput, optFns ...func(*Options)) (*ListVirtualRoutersOutput, error) {
	if params == nil {
		params = &ListVirtualRoutersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualRouters", params, optFns, c.addOperationListVirtualRoutersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualRoutersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualRoutersInput struct {

	// The name of the service mesh to list virtual routers in.
	//
	// This member is required.
	MeshName *string

	// The maximum number of results returned by ListVirtualRouters in paginated
	// output. When you use this parameter, ListVirtualRouters returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualRouters request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualRouters returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The Amazon Web Services IAM account ID of the service mesh owner. If the
	// account ID is not your own, then it's the ID of the account that shared the mesh
	// with your account. For more information about mesh sharing, see Working with
	// shared meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html)
	// .
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListVirtualRouters
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVirtualRoutersOutput struct {

	// The list of existing virtual routers for the specified service mesh.
	//
	// This member is required.
	VirtualRouters []types.VirtualRouterRef

	// The nextToken value to include in a future ListVirtualRouters request. When the
	// results of a ListVirtualRouters request exceed limit , you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVirtualRoutersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualRouters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualRouters{}, middleware.After)
	if err != nil {
		return err
	}
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
	if err = addListVirtualRoutersResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListVirtualRoutersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualRouters(options.Region), middleware.Before); err != nil {
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

// ListVirtualRoutersAPIClient is a client that implements the ListVirtualRouters
// operation.
type ListVirtualRoutersAPIClient interface {
	ListVirtualRouters(context.Context, *ListVirtualRoutersInput, ...func(*Options)) (*ListVirtualRoutersOutput, error)
}

var _ ListVirtualRoutersAPIClient = (*Client)(nil)

// ListVirtualRoutersPaginatorOptions is the paginator options for
// ListVirtualRouters
type ListVirtualRoutersPaginatorOptions struct {
	// The maximum number of results returned by ListVirtualRouters in paginated
	// output. When you use this parameter, ListVirtualRouters returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualRouters request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualRouters returns
	// up to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualRoutersPaginator is a paginator for ListVirtualRouters
type ListVirtualRoutersPaginator struct {
	options   ListVirtualRoutersPaginatorOptions
	client    ListVirtualRoutersAPIClient
	params    *ListVirtualRoutersInput
	nextToken *string
	firstPage bool
}

// NewListVirtualRoutersPaginator returns a new ListVirtualRoutersPaginator
func NewListVirtualRoutersPaginator(client ListVirtualRoutersAPIClient, params *ListVirtualRoutersInput, optFns ...func(*ListVirtualRoutersPaginatorOptions)) *ListVirtualRoutersPaginator {
	if params == nil {
		params = &ListVirtualRoutersInput{}
	}

	options := ListVirtualRoutersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVirtualRoutersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualRoutersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVirtualRouters page.
func (p *ListVirtualRoutersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualRoutersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListVirtualRouters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVirtualRouters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "ListVirtualRouters",
	}
}

type opListVirtualRoutersResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListVirtualRoutersResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListVirtualRoutersResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "appmesh"
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
				signingName = "appmesh"
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
				v4aScheme.SigningName = aws.String("appmesh")
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

func addListVirtualRoutersResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListVirtualRoutersResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
