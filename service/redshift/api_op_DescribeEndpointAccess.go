// Code generated by smithy-go-codegen DO NOT EDIT.


package redshift

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
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Describes a Redshift-managed VPC endpoint.
func (c *Client) DescribeEndpointAccess(ctx context.Context, params *DescribeEndpointAccessInput, optFns ...func(*Options)) (*DescribeEndpointAccessOutput, error) {
	if params == nil { params = &DescribeEndpointAccessInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointAccess", params, optFns, c.addOperationDescribeEndpointAccessMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointAccessInput struct {
	
	// The cluster identifier associated with the described endpoint.
	ClusterIdentifier *string
	
	// The name of the endpoint to be described.
	EndpointName *string
	
	// An optional pagination token provided by a previous DescribeEndpointAccess
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string
	
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32
	
	// The Amazon Web Services account ID of the owner of the cluster.
	ResourceOwner *string
	
	// The virtual private cloud (VPC) identifier with access to the cluster.
	VpcId *string
	
	noSmithyDocumentSerde
}

type DescribeEndpointAccessOutput struct {
	
	// The list of endpoints with access to the cluster.
	EndpointAccessList []types.EndpointAccess
	
	// An optional pagination token provided by a previous DescribeEndpointAccess
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEndpointAccess{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEndpointAccess{}, middleware.After)
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
	if err = addDescribeEndpointAccessResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointAccess(options.Region, ), middleware.Before); err != nil {
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

// DescribeEndpointAccessAPIClient is a client that implements the
// DescribeEndpointAccess operation.
type DescribeEndpointAccessAPIClient interface {
	DescribeEndpointAccess(context.Context, *DescribeEndpointAccessInput, ...func(*Options)) (*DescribeEndpointAccessOutput, error)
}

var _ DescribeEndpointAccessAPIClient = (*Client)(nil)

// DescribeEndpointAccessPaginatorOptions is the paginator options for
// DescribeEndpointAccess
type DescribeEndpointAccessPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointAccessPaginator is a paginator for DescribeEndpointAccess
type DescribeEndpointAccessPaginator struct {
    options DescribeEndpointAccessPaginatorOptions
    client DescribeEndpointAccessAPIClient
    params *DescribeEndpointAccessInput
    nextToken *string
    firstPage bool
}

// NewDescribeEndpointAccessPaginator returns a new DescribeEndpointAccessPaginator
func NewDescribeEndpointAccessPaginator(client DescribeEndpointAccessAPIClient, params *DescribeEndpointAccessInput, optFns ...func(*DescribeEndpointAccessPaginatorOptions)) *DescribeEndpointAccessPaginator {
	if params == nil {
	    params = &DescribeEndpointAccessInput{}
	}
	
	options := DescribeEndpointAccessPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeEndpointAccessPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointAccessPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeEndpointAccess page.
func (p *DescribeEndpointAccessPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointAccessOutput, error) {
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
	
	result, err := p.client.DescribeEndpointAccess(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "redshift",
	OperationName: "DescribeEndpointAccess",
	}
}

type opDescribeEndpointAccessResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeEndpointAccessResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeEndpointAccessResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDescribeEndpointAccessResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeEndpointAccessResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
