// Code generated by smithy-go-codegen DO NOT EDIT.


package neptune

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
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns a list of orderable DB instance options for the specified engine.
func (c *Client) DescribeOrderableDBInstanceOptions(ctx context.Context, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
	if params == nil { params = &DescribeOrderableDBInstanceOptionsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableDBInstanceOptions", params, optFns, c.addOperationDescribeOrderableDBInstanceOptionsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeOrderableDBInstanceOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrderableDBInstanceOptionsInput struct {
	
	// The name of the engine to retrieve DB instance options for.
	//
	// This member is required.
	Engine *string
	
	// The DB instance class filter value. Specify this parameter to show only the
	// available offerings matching the specified DB instance class.
	DBInstanceClass *string
	
	// The engine version filter value. Specify this parameter to show only the
	// available offerings matching the specified engine version.
	EngineVersion *string
	
	// This parameter is not currently supported.
	Filters []types.Filter
	
	// The license model filter value. Specify this parameter to show only the
	// available offerings matching the specified license model.
	LicenseModel *string
	
	// An optional pagination token provided by a previous
	// DescribeOrderableDBInstanceOptions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string
	
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	
	// The VPC filter value. Specify this parameter to show only the available VPC or
	// non-VPC offerings.
	Vpc *bool
	
	noSmithyDocumentSerde
}

type DescribeOrderableDBInstanceOptionsOutput struct {
	
	// An optional pagination token provided by a previous OrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string
	
	// An OrderableDBInstanceOption structure containing information about orderable
	// options for the DB instance.
	OrderableDBInstanceOptions []types.OrderableDBInstanceOption
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrderableDBInstanceOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeOrderableDBInstanceOptions{}, middleware.After)
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
	if err = addDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeOrderableDBInstanceOptionsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(options.Region, ), middleware.Before); err != nil {
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

// DescribeOrderableDBInstanceOptionsAPIClient is a client that implements the
// DescribeOrderableDBInstanceOptions operation.
type DescribeOrderableDBInstanceOptionsAPIClient interface {
	DescribeOrderableDBInstanceOptions(context.Context, *DescribeOrderableDBInstanceOptionsInput, ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error)
}

var _ DescribeOrderableDBInstanceOptionsAPIClient = (*Client)(nil)

// DescribeOrderableDBInstanceOptionsPaginatorOptions is the paginator options for
// DescribeOrderableDBInstanceOptions
type DescribeOrderableDBInstanceOptionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrderableDBInstanceOptionsPaginator is a paginator for
// DescribeOrderableDBInstanceOptions
type DescribeOrderableDBInstanceOptionsPaginator struct {
    options DescribeOrderableDBInstanceOptionsPaginatorOptions
    client DescribeOrderableDBInstanceOptionsAPIClient
    params *DescribeOrderableDBInstanceOptionsInput
    nextToken *string
    firstPage bool
}

// NewDescribeOrderableDBInstanceOptionsPaginator returns a new
// DescribeOrderableDBInstanceOptionsPaginator
func NewDescribeOrderableDBInstanceOptionsPaginator(client DescribeOrderableDBInstanceOptionsAPIClient, params *DescribeOrderableDBInstanceOptionsInput, optFns ...func(*DescribeOrderableDBInstanceOptionsPaginatorOptions)) *DescribeOrderableDBInstanceOptionsPaginator {
	if params == nil {
	    params = &DescribeOrderableDBInstanceOptionsInput{}
	}
	
	options := DescribeOrderableDBInstanceOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeOrderableDBInstanceOptionsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrderableDBInstanceOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeOrderableDBInstanceOptions page.
func (p *DescribeOrderableDBInstanceOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrderableDBInstanceOptionsOutput, error) {
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
	
	result, err := p.client.DescribeOrderableDBInstanceOptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOrderableDBInstanceOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rds",
	OperationName: "DescribeOrderableDBInstanceOptions",
	}
}

type opDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "rds"
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
	        signingName = "rds"
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
	        v4aScheme.SigningName = aws.String("rds")
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

func addDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeOrderableDBInstanceOptionsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
