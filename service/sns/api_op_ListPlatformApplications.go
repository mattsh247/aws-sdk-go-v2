// Code generated by smithy-go-codegen DO NOT EDIT.


package sns

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
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). The results for
// ListPlatformApplications are paginated and return a limited list of
// applications, up to 100. If additional records are available after the first
// page results, then a NextToken string will be returned. To receive the next
// page, you call ListPlatformApplications using the NextToken string received
// from the previous call. When there are no more records to return, NextToken
// will be null. For more information, see Using Amazon SNS Mobile Push
// Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html) .
// This action is throttled at 15 transactions per second (TPS).
func (c *Client) ListPlatformApplications(ctx context.Context, params *ListPlatformApplicationsInput, optFns ...func(*Options)) (*ListPlatformApplicationsOutput, error) {
	if params == nil { params = &ListPlatformApplicationsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ListPlatformApplications", params, optFns, c.addOperationListPlatformApplicationsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ListPlatformApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListPlatformApplications action.
type ListPlatformApplicationsInput struct {
	
	// NextToken string is used when calling ListPlatformApplications action to
	// retrieve additional records that are available after the first page results.
	NextToken *string
	
	noSmithyDocumentSerde
}

// Response for ListPlatformApplications action.
type ListPlatformApplicationsOutput struct {
	
	// NextToken string is returned when calling ListPlatformApplications action if
	// additional records are available after the first page results.
	NextToken *string
	
	// Platform applications returned when calling ListPlatformApplications action.
	PlatformApplications []types.PlatformApplication
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationListPlatformApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformApplications{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformApplications{}, middleware.After)
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
	if err = addListPlatformApplicationsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformApplications(options.Region, ), middleware.Before); err != nil {
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

// ListPlatformApplicationsAPIClient is a client that implements the
// ListPlatformApplications operation.
type ListPlatformApplicationsAPIClient interface {
	ListPlatformApplications(context.Context, *ListPlatformApplicationsInput, ...func(*Options)) (*ListPlatformApplicationsOutput, error)
}

var _ ListPlatformApplicationsAPIClient = (*Client)(nil)

// ListPlatformApplicationsPaginatorOptions is the paginator options for
// ListPlatformApplications
type ListPlatformApplicationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPlatformApplicationsPaginator is a paginator for ListPlatformApplications
type ListPlatformApplicationsPaginator struct {
    options ListPlatformApplicationsPaginatorOptions
    client ListPlatformApplicationsAPIClient
    params *ListPlatformApplicationsInput
    nextToken *string
    firstPage bool
}

// NewListPlatformApplicationsPaginator returns a new
// ListPlatformApplicationsPaginator
func NewListPlatformApplicationsPaginator(client ListPlatformApplicationsAPIClient, params *ListPlatformApplicationsInput, optFns ...func(*ListPlatformApplicationsPaginatorOptions)) *ListPlatformApplicationsPaginator {
	if params == nil {
	    params = &ListPlatformApplicationsInput{}
	}
	
	options := ListPlatformApplicationsPaginatorOptions{}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ListPlatformApplicationsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPlatformApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ListPlatformApplications page.
func (p *ListPlatformApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPlatformApplicationsOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.NextToken = p.nextToken
	
	result, err := p.client.ListPlatformApplications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPlatformApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "sns",
	OperationName: "ListPlatformApplications",
	}
}

type opListPlatformApplicationsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opListPlatformApplicationsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListPlatformApplicationsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "sns"
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
	        signingName = "sns"
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
	        v4aScheme.SigningName = aws.String("sns")
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

func addListPlatformApplicationsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opListPlatformApplicationsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
