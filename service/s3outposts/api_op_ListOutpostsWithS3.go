// Code generated by smithy-go-codegen DO NOT EDIT.

package s3outposts

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/s3outposts/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Outposts with S3 on Outposts capacity for your Amazon Web Services
// account. Includes S3 on Outposts that you have access to as the Outposts owner,
// or as a shared user from Resource Access Manager (RAM).
func (c *Client) ListOutpostsWithS3(ctx context.Context, params *ListOutpostsWithS3Input, optFns ...func(*Options)) (*ListOutpostsWithS3Output, error) {
	if params == nil {
		params = &ListOutpostsWithS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOutpostsWithS3", params, optFns, c.addOperationListOutpostsWithS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOutpostsWithS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOutpostsWithS3Input struct {

	// The maximum number of Outposts to return. The limit is 100.
	MaxResults int32

	// When you can get additional results from the ListOutpostsWithS3 call, a
	// NextToken parameter is returned in the output. You can then pass in a subsequent
	// command to the NextToken parameter to continue listing additional Outposts.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOutpostsWithS3Output struct {

	// Returns a token that you can use to call ListOutpostsWithS3 again and receive
	// additional results, if there are any.
	NextToken *string

	// Returns the list of Outposts that have the following characteristics:
	//   - outposts that have S3 provisioned
	//   - outposts that are Active (not pending any provisioning nor decommissioned)
	//   - outposts to which the the calling Amazon Web Services account has access
	Outposts []types.Outpost

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOutpostsWithS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOutpostsWithS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOutpostsWithS3{}, middleware.After)
	if err != nil {
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
	if err = addListOutpostsWithS3ResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOutpostsWithS3(options.Region), middleware.Before); err != nil {
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
	if err = addListOutpostsWithS3EndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListOutpostsWithS3APIClient is a client that implements the ListOutpostsWithS3
// operation.
type ListOutpostsWithS3APIClient interface {
	ListOutpostsWithS3(context.Context, *ListOutpostsWithS3Input, ...func(*Options)) (*ListOutpostsWithS3Output, error)
}

var _ ListOutpostsWithS3APIClient = (*Client)(nil)

// ListOutpostsWithS3PaginatorOptions is the paginator options for
// ListOutpostsWithS3
type ListOutpostsWithS3PaginatorOptions struct {
	// The maximum number of Outposts to return. The limit is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOutpostsWithS3Paginator is a paginator for ListOutpostsWithS3
type ListOutpostsWithS3Paginator struct {
	options   ListOutpostsWithS3PaginatorOptions
	client    ListOutpostsWithS3APIClient
	params    *ListOutpostsWithS3Input
	nextToken *string
	firstPage bool
}

// NewListOutpostsWithS3Paginator returns a new ListOutpostsWithS3Paginator
func NewListOutpostsWithS3Paginator(client ListOutpostsWithS3APIClient, params *ListOutpostsWithS3Input, optFns ...func(*ListOutpostsWithS3PaginatorOptions)) *ListOutpostsWithS3Paginator {
	if params == nil {
		params = &ListOutpostsWithS3Input{}
	}

	options := ListOutpostsWithS3PaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOutpostsWithS3Paginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOutpostsWithS3Paginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOutpostsWithS3 page.
func (p *ListOutpostsWithS3Paginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOutpostsWithS3Output, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListOutpostsWithS3(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOutpostsWithS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3-outposts",
		OperationName: "ListOutpostsWithS3",
	}
}

type opListOutpostsWithS3EndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opListOutpostsWithS3EndpointDisableHTTPSMiddleware) ID() string {
	return "opListOutpostsWithS3EndpointDisableHTTPSMiddleware"
}

func (m *opListOutpostsWithS3EndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addListOutpostsWithS3EndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opListOutpostsWithS3EndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opListOutpostsWithS3ResolveEndpointMiddleware", middleware.After)
}

type opListOutpostsWithS3ResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opListOutpostsWithS3ResolveEndpointMiddleware) ID() string {
	return "opListOutpostsWithS3ResolveEndpointMiddleware"
}

func (m *opListOutpostsWithS3ResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
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
			signingName := "s3-outposts"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
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
				signingName = "s3-outposts"
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
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
			var signingName string
			if v4aScheme.SigningName == nil {
				signingName = "s3-outposts"
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListOutpostsWithS3ResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListOutpostsWithS3ResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
