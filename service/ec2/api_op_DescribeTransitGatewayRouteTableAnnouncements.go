// Code generated by smithy-go-codegen DO NOT EDIT.


package ec2

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
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Describes one or more transit gateway route table advertisements.
func (c *Client) DescribeTransitGatewayRouteTableAnnouncements(ctx context.Context, params *DescribeTransitGatewayRouteTableAnnouncementsInput, optFns ...func(*Options)) (*DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	if params == nil { params = &DescribeTransitGatewayRouteTableAnnouncementsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayRouteTableAnnouncements", params, optFns, c.addOperationDescribeTransitGatewayRouteTableAnnouncementsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeTransitGatewayRouteTableAnnouncementsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayRouteTableAnnouncementsInput struct {
	
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool
	
	// The filters associated with the transit gateway policy table.
	Filters []types.Filter
	
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	
	// The token for the next page of results.
	NextToken *string
	
	// The IDs of the transit gateway route tables that are being advertised.
	TransitGatewayRouteTableAnnouncementIds []string
	
	noSmithyDocumentSerde
}

type DescribeTransitGatewayRouteTableAnnouncementsOutput struct {
	
	// The token for the next page of results.
	NextToken *string
	
	// Describes the transit gateway route table announcement.
	TransitGatewayRouteTableAnnouncements []types.TransitGatewayRouteTableAnnouncement
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewayRouteTableAnnouncementsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayRouteTableAnnouncements{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayRouteTableAnnouncements{}, middleware.After)
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
	if err = addDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTableAnnouncements(options.Region, ), middleware.Before); err != nil {
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

// DescribeTransitGatewayRouteTableAnnouncementsAPIClient is a client that
// implements the DescribeTransitGatewayRouteTableAnnouncements operation.
type DescribeTransitGatewayRouteTableAnnouncementsAPIClient interface {
	DescribeTransitGatewayRouteTableAnnouncements(context.Context, *DescribeTransitGatewayRouteTableAnnouncementsInput, ...func(*Options)) (*DescribeTransitGatewayRouteTableAnnouncementsOutput, error)
}

var _ DescribeTransitGatewayRouteTableAnnouncementsAPIClient = (*Client)(nil)

// DescribeTransitGatewayRouteTableAnnouncementsPaginatorOptions is the paginator
// options for DescribeTransitGatewayRouteTableAnnouncements
type DescribeTransitGatewayRouteTableAnnouncementsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayRouteTableAnnouncementsPaginator is a paginator for
// DescribeTransitGatewayRouteTableAnnouncements
type DescribeTransitGatewayRouteTableAnnouncementsPaginator struct {
    options DescribeTransitGatewayRouteTableAnnouncementsPaginatorOptions
    client DescribeTransitGatewayRouteTableAnnouncementsAPIClient
    params *DescribeTransitGatewayRouteTableAnnouncementsInput
    nextToken *string
    firstPage bool
}

// NewDescribeTransitGatewayRouteTableAnnouncementsPaginator returns a new
// DescribeTransitGatewayRouteTableAnnouncementsPaginator
func NewDescribeTransitGatewayRouteTableAnnouncementsPaginator(client DescribeTransitGatewayRouteTableAnnouncementsAPIClient, params *DescribeTransitGatewayRouteTableAnnouncementsInput, optFns ...func(*DescribeTransitGatewayRouteTableAnnouncementsPaginatorOptions)) *DescribeTransitGatewayRouteTableAnnouncementsPaginator {
	if params == nil {
	    params = &DescribeTransitGatewayRouteTableAnnouncementsInput{}
	}
	
	options := DescribeTransitGatewayRouteTableAnnouncementsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeTransitGatewayRouteTableAnnouncementsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayRouteTableAnnouncementsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeTransitGatewayRouteTableAnnouncements page.
func (p *DescribeTransitGatewayRouteTableAnnouncementsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
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
	
	result, err := p.client.DescribeTransitGatewayRouteTableAnnouncements(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTableAnnouncements(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ec2",
	OperationName: "DescribeTransitGatewayRouteTableAnnouncements",
	}
}

type opDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ec2"
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
	        signingName = "ec2"
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
	        v4aScheme.SigningName = aws.String("ec2")
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

func addDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeTransitGatewayRouteTableAnnouncementsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
