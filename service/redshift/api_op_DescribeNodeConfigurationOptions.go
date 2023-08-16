// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns properties of possible node configurations such as node type, number of
// nodes, and disk usage for the specified action type.
func (c *Client) DescribeNodeConfigurationOptions(ctx context.Context, params *DescribeNodeConfigurationOptionsInput, optFns ...func(*Options)) (*DescribeNodeConfigurationOptionsOutput, error) {
	if params == nil {
		params = &DescribeNodeConfigurationOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNodeConfigurationOptions", params, optFns, c.addOperationDescribeNodeConfigurationOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNodeConfigurationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNodeConfigurationOptionsInput struct {

	// The action type to evaluate for possible node configurations. Specify
	// "restore-cluster" to get configuration combinations based on an existing
	// snapshot. Specify "recommend-node-config" to get configuration recommendations
	// based on an existing cluster or snapshot. Specify "resize-cluster" to get
	// configuration combinations for elastic resize based on an existing cluster.
	//
	// This member is required.
	ActionType types.ActionType

	// The identifier of the cluster to evaluate for possible node configurations.
	ClusterIdentifier *string

	// A set of name, operator, and value items to filter the results.
	Filters []types.NodeConfigurationOptionsFilter

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeNodeConfigurationOptions
	// request exceed the value specified in MaxRecords , Amazon Web Services returns a
	// value in the Marker field of the response. You can retrieve the next set of
	// response records by providing the returned marker value in the Marker parameter
	// and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 500
	// Constraints: minimum 100, maximum 500.
	MaxRecords *int32

	// The Amazon Web Services account used to create or copy the snapshot. Required
	// if you are restoring a snapshot you do not own, optional if you own the
	// snapshot.
	OwnerAccount *string

	// The Amazon Resource Name (ARN) of the snapshot associated with the message to
	// describe node configuration.
	SnapshotArn *string

	// The identifier of the snapshot to evaluate for possible node configurations.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

type DescribeNodeConfigurationOptionsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of valid node configurations.
	NodeConfigurationOptionList []types.NodeConfigurationOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNodeConfigurationOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeNodeConfigurationOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeNodeConfigurationOptions{}, middleware.After)
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
	if err = addDescribeNodeConfigurationOptionsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDescribeNodeConfigurationOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNodeConfigurationOptions(options.Region), middleware.Before); err != nil {
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

// DescribeNodeConfigurationOptionsAPIClient is a client that implements the
// DescribeNodeConfigurationOptions operation.
type DescribeNodeConfigurationOptionsAPIClient interface {
	DescribeNodeConfigurationOptions(context.Context, *DescribeNodeConfigurationOptionsInput, ...func(*Options)) (*DescribeNodeConfigurationOptionsOutput, error)
}

var _ DescribeNodeConfigurationOptionsAPIClient = (*Client)(nil)

// DescribeNodeConfigurationOptionsPaginatorOptions is the paginator options for
// DescribeNodeConfigurationOptions
type DescribeNodeConfigurationOptionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 500
	// Constraints: minimum 100, maximum 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNodeConfigurationOptionsPaginator is a paginator for
// DescribeNodeConfigurationOptions
type DescribeNodeConfigurationOptionsPaginator struct {
	options   DescribeNodeConfigurationOptionsPaginatorOptions
	client    DescribeNodeConfigurationOptionsAPIClient
	params    *DescribeNodeConfigurationOptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNodeConfigurationOptionsPaginator returns a new
// DescribeNodeConfigurationOptionsPaginator
func NewDescribeNodeConfigurationOptionsPaginator(client DescribeNodeConfigurationOptionsAPIClient, params *DescribeNodeConfigurationOptionsInput, optFns ...func(*DescribeNodeConfigurationOptionsPaginatorOptions)) *DescribeNodeConfigurationOptionsPaginator {
	if params == nil {
		params = &DescribeNodeConfigurationOptionsInput{}
	}

	options := DescribeNodeConfigurationOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNodeConfigurationOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNodeConfigurationOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeNodeConfigurationOptions page.
func (p *DescribeNodeConfigurationOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNodeConfigurationOptionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeNodeConfigurationOptions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeNodeConfigurationOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeNodeConfigurationOptions",
	}
}

type opDescribeNodeConfigurationOptionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDescribeNodeConfigurationOptionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeNodeConfigurationOptionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "redshift"
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
				signingName = "redshift"
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
				v4aScheme.SigningName = aws.String("redshift")
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

func addDescribeNodeConfigurationOptionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDescribeNodeConfigurationOptionsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
