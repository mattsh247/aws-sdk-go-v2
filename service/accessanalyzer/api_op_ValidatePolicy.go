// Code generated by smithy-go-codegen DO NOT EDIT.


package accessanalyzer

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
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Requests the validation of a policy and returns a list of findings. The
// findings help you identify issues and provide actionable recommendations to
// resolve the issue and enable you to author functional policies that meet
// security best practices.
func (c *Client) ValidatePolicy(ctx context.Context, params *ValidatePolicyInput, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
	if params == nil { params = &ValidatePolicyInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ValidatePolicy", params, optFns, c.addOperationValidatePolicyMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ValidatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ValidatePolicyInput struct {
	
	// The JSON policy document to use as the content for the policy.
	//
	// This member is required.
	PolicyDocument *string
	
	// The type of policy to validate. Identity policies grant permissions to IAM
	// principals. Identity policies include managed and inline policies for IAM roles,
	// users, and groups. They also include service-control policies (SCPs) that are
	// attached to an Amazon Web Services organization, organizational unit (OU), or an
	// account. Resource policies grant permissions on Amazon Web Services resources.
	// Resource policies include trust policies for IAM roles and bucket policies for
	// Amazon S3 buckets. You can provide a generic input such as identity policy or
	// resource policy or a specific input such as managed policy or Amazon S3 bucket
	// policy.
	//
	// This member is required.
	PolicyType types.PolicyType
	
	// The locale to use for localizing the findings.
	Locale types.Locale
	
	// The maximum number of results to return in the response.
	MaxResults *int32
	
	// A token used for pagination of results returned.
	NextToken *string
	
	// The type of resource to attach to your resource policy. Specify a value for the
	// policy validation resource type only if the policy type is RESOURCE_POLICY . For
	// example, to validate a resource policy to attach to an Amazon S3 bucket, you can
	// choose AWS::S3::Bucket for the policy validation resource type. For resource
	// types not supported as valid values, IAM Access Analyzer runs policy checks that
	// apply to all resource policies. For example, to validate a resource policy to
	// attach to a KMS key, do not specify a value for the policy validation resource
	// type and IAM Access Analyzer will run policy checks that apply to all resource
	// policies.
	ValidatePolicyResourceType types.ValidatePolicyResourceType
	
	noSmithyDocumentSerde
}

type ValidatePolicyOutput struct {
	
	// The list of findings in a policy returned by IAM Access Analyzer based on its
	// suite of policy checks.
	//
	// This member is required.
	Findings []types.ValidatePolicyFinding
	
	// A token used for pagination of results returned.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationValidatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpValidatePolicy{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpValidatePolicy{}, middleware.After)
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
	if err = addValidatePolicyResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpValidatePolicyValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opValidatePolicy(options.Region, ), middleware.Before); err != nil {
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

// ValidatePolicyAPIClient is a client that implements the ValidatePolicy
// operation.
type ValidatePolicyAPIClient interface {
	ValidatePolicy(context.Context, *ValidatePolicyInput, ...func(*Options)) (*ValidatePolicyOutput, error)
}

var _ ValidatePolicyAPIClient = (*Client)(nil)

// ValidatePolicyPaginatorOptions is the paginator options for ValidatePolicy
type ValidatePolicyPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ValidatePolicyPaginator is a paginator for ValidatePolicy
type ValidatePolicyPaginator struct {
    options ValidatePolicyPaginatorOptions
    client ValidatePolicyAPIClient
    params *ValidatePolicyInput
    nextToken *string
    firstPage bool
}

// NewValidatePolicyPaginator returns a new ValidatePolicyPaginator
func NewValidatePolicyPaginator(client ValidatePolicyAPIClient, params *ValidatePolicyInput, optFns ...func(*ValidatePolicyPaginatorOptions)) *ValidatePolicyPaginator {
	if params == nil {
	    params = &ValidatePolicyInput{}
	}
	
	options := ValidatePolicyPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &ValidatePolicyPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ValidatePolicyPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next ValidatePolicy page.
func (p *ValidatePolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ValidatePolicyOutput, error) {
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
	
	result, err := p.client.ValidatePolicy(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opValidatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "access-analyzer",
	OperationName: "ValidatePolicy",
	}
}

type opValidatePolicyResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opValidatePolicyResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opValidatePolicyResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "access-analyzer"
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
	        signingName = "access-analyzer"
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
	        v4aScheme.SigningName = aws.String("access-analyzer")
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

func addValidatePolicyResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opValidatePolicyResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
