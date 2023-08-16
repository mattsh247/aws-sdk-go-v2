// Code generated by smithy-go-codegen DO NOT EDIT.


package globalaccelerator

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
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Provides a complete mapping from the public accelerator IP address and port to
// destination EC2 instance IP addresses and ports in the virtual public cloud
// (VPC) subnet endpoint for a custom routing accelerator. For each subnet endpoint
// that you add, Global Accelerator creates a new static port mapping for the
// accelerator. The port mappings don't change after Global Accelerator generates
// them, so you can retrieve and cache the full mapping on your servers. If you
// remove a subnet from your accelerator, Global Accelerator removes (reclaims) the
// port mappings. If you add a subnet to your accelerator, Global Accelerator
// creates new port mappings (the existing ones don't change). If you add or remove
// EC2 instances in your subnet, the port mappings don't change, because the
// mappings are created when you add the subnet to Global Accelerator. The mappings
// also include a flag for each destination denoting which destination IP addresses
// and ports are allowed or denied traffic.
func (c *Client) ListCustomRoutingPortMappings(ctx context.Context, params *ListCustomRoutingPortMappingsInput, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsOutput, error) {
	if params == nil { params = &ListCustomRoutingPortMappingsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListCustomRoutingPortMappings", params, optFns, c.addOperationListCustomRoutingPortMappingsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListCustomRoutingPortMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomRoutingPortMappingsInput struct {
	
	// The Amazon Resource Name (ARN) of the accelerator to list the custom routing
	// port mappings for.
	//
	// This member is required.
	AcceleratorArn *string
	
	// The Amazon Resource Name (ARN) of the endpoint group to list the custom routing
	// port mappings for.
	EndpointGroupArn *string
	
	// The number of destination port mappings that you want to return with this call.
	// The default value is 10.
	MaxResults *int32
	
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
	
	noSmithyDocumentSerde
}

type ListCustomRoutingPortMappingsOutput struct {
	
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
	
	// The port mappings for a custom routing accelerator.
	PortMappings []types.PortMapping
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomRoutingPortMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCustomRoutingPortMappings{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCustomRoutingPortMappings{}, middleware.After)
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
	if err = addListCustomRoutingPortMappingsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpListCustomRoutingPortMappingsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomRoutingPortMappings(options.Region, ), middleware.Before); err != nil {
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

// ListCustomRoutingPortMappingsAPIClient is a client that implements the
// ListCustomRoutingPortMappings operation.
type ListCustomRoutingPortMappingsAPIClient interface {
	ListCustomRoutingPortMappings(context.Context, *ListCustomRoutingPortMappingsInput, ...func(*Options)) (*ListCustomRoutingPortMappingsOutput, error)
}

var _ ListCustomRoutingPortMappingsAPIClient = (*Client)(nil)

// ListCustomRoutingPortMappingsPaginatorOptions is the paginator options for
// ListCustomRoutingPortMappings
type ListCustomRoutingPortMappingsPaginatorOptions struct {
	// The number of destination port mappings that you want to return with this call.
	// The default value is 10.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomRoutingPortMappingsPaginator is a paginator for
// ListCustomRoutingPortMappings
type ListCustomRoutingPortMappingsPaginator struct {
    options ListCustomRoutingPortMappingsPaginatorOptions
    client ListCustomRoutingPortMappingsAPIClient
    params *ListCustomRoutingPortMappingsInput
    nextToken *string
    firstPage bool
}

// NewListCustomRoutingPortMappingsPaginator returns a new
// ListCustomRoutingPortMappingsPaginator
func NewListCustomRoutingPortMappingsPaginator(client ListCustomRoutingPortMappingsAPIClient, params *ListCustomRoutingPortMappingsInput, optFns ...func(*ListCustomRoutingPortMappingsPaginatorOptions)) *ListCustomRoutingPortMappingsPaginator {
	if params == nil {
	    params = &ListCustomRoutingPortMappingsInput{}
	}
	
	options := ListCustomRoutingPortMappingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListCustomRoutingPortMappingsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomRoutingPortMappingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListCustomRoutingPortMappings page.
func (p *ListCustomRoutingPortMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomRoutingPortMappingsOutput, error) {
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
	
	result, err := p.client.ListCustomRoutingPortMappings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomRoutingPortMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "globalaccelerator",
	OperationName: "ListCustomRoutingPortMappings",
	}
}

type opListCustomRoutingPortMappingsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListCustomRoutingPortMappingsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListCustomRoutingPortMappingsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "globalaccelerator"
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
	        signingName = "globalaccelerator"
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
	        v4aScheme.SigningName = aws.String("globalaccelerator")
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

func addListCustomRoutingPortMappingsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListCustomRoutingPortMappingsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
