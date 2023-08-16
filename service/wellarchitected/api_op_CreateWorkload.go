// Code generated by smithy-go-codegen DO NOT EDIT.


package wellarchitected

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
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Create a new workload. The owner of a workload can share the workload with
// other Amazon Web Services accounts, users, an organization, and organizational
// units (OUs) in the same Amazon Web Services Region. Only the owner of a workload
// can delete it. For more information, see Defining a Workload (https://docs.aws.amazon.com/wellarchitected/latest/userguide/define-workload.html)
// in the Well-Architected Tool User Guide. Either AwsRegions , NonAwsRegions , or
// both must be specified when creating a workload. You also must specify
// ReviewOwner , even though the parameter is listed as not being required in the
// following section.
func (c *Client) CreateWorkload(ctx context.Context, params *CreateWorkloadInput, optFns ...func(*Options)) (*CreateWorkloadOutput, error) {
	if params == nil { params = &CreateWorkloadInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateWorkload", params, optFns, c.addOperationCreateWorkloadMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateWorkloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for workload creation.
type CreateWorkloadInput struct {
	
	// A unique case-sensitive string used to ensure that this request is idempotent
	// (executes only once). You should not reuse the same token for other requests. If
	// you retry a request with the same client request token and the same parameters
	// after the original request has completed successfully, the result of the
	// original request is returned. This token is listed as required, however, if you
	// do not specify it, the Amazon Web Services SDKs automatically generate one for
	// you. If you are not using the Amazon Web Services SDK or the CLI, you must
	// provide this token or the request will fail.
	//
	// This member is required.
	ClientRequestToken *string
	
	// The description for the workload.
	//
	// This member is required.
	Description *string
	
	// The environment for the workload.
	//
	// This member is required.
	Environment types.WorkloadEnvironment
	
	// The list of lenses associated with the workload. Each lens is identified by its
	// LensSummary$LensAlias .
	//
	// This member is required.
	Lenses []string
	
	// The name of the workload. The name must be unique within an account within an
	// Amazon Web Services Region. Spaces and capitalization are ignored when checking
	// for uniqueness.
	//
	// This member is required.
	WorkloadName *string
	
	// The list of Amazon Web Services account IDs associated with the workload.
	AccountIds []string
	
	// List of AppRegistry application ARNs associated to the workload.
	Applications []string
	
	// The URL of the architectural design for the workload.
	ArchitecturalDesign *string
	
	// The list of Amazon Web Services Regions associated with the workload, for
	// example, us-east-2 , or ca-central-1 .
	AwsRegions []string
	
	// Well-Architected discovery configuration settings associated to the workload.
	DiscoveryConfig *types.WorkloadDiscoveryConfig
	
	// The industry for the workload.
	Industry *string
	
	// The industry type for the workload. If specified, must be one of the following:
	//   - Agriculture
	//   - Automobile
	//   - Defense
	//   - Design and Engineering
	//   - Digital Advertising
	//   - Education
	//   - Environmental Protection
	//   - Financial Services
	//   - Gaming
	//   - General Public Services
	//   - Healthcare
	//   - Hospitality
	//   - InfoTech
	//   - Justice and Public Safety
	//   - Life Sciences
	//   - Manufacturing
	//   - Media & Entertainment
	//   - Mining & Resources
	//   - Oil & Gas
	//   - Power & Utilities
	//   - Professional Services
	//   - Real Estate & Construction
	//   - Retail & Wholesale
	//   - Social Protection
	//   - Telecommunications
	//   - Travel, Transportation & Logistics
	//   - Other
	IndustryType *string
	
	// The list of non-Amazon Web Services Regions associated with the workload.
	NonAwsRegions []string
	
	// The notes associated with the workload.
	Notes *string
	
	// The priorities of the pillars, which are used to order items in the improvement
	// plan. Each pillar is represented by its PillarReviewSummary$PillarId .
	PillarPriorities []string
	
	// The list of profile ARNs associated with the workload.
	ProfileArns []string
	
	// The review owner of the workload. The name, email address, or identifier for
	// the primary group or individual that owns the workload review process.
	ReviewOwner *string
	
	// The tags to be associated with the workload.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

// Output of a create workload call.
type CreateWorkloadOutput struct {
	
	// The ARN for the workload.
	WorkloadArn *string
	
	// The ID assigned to the workload. This ID is unique within an Amazon Web
	// Services Region.
	WorkloadId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkload{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkload{}, middleware.After)
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
	if err = addCreateWorkloadResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opCreateWorkloadMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateWorkloadValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkload(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorkload struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorkload) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorkload) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*CreateWorkloadInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkloadInput ")}
	
	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWorkloadMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorkload{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorkload(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "wellarchitected",
	OperationName: "CreateWorkload",
	}
}

type opCreateWorkloadResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateWorkloadResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateWorkloadResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "wellarchitected"
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
	        signingName = "wellarchitected"
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
	        v4aScheme.SigningName = aws.String("wellarchitected")
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

func addCreateWorkloadResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateWorkloadResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
