// Code generated by smithy-go-codegen DO NOT EDIT.


package servicecatalog

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
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Provisions the specified product. A provisioned product is a resourced instance
// of a product. For example, provisioning a product that's based on an
// CloudFormation template launches an CloudFormation stack and its underlying
// resources. You can check the status of this request using DescribeRecord . If
// the request contains a tag key with an empty list of values, there's a tag
// conflict for that key. Don't include conflicted keys as tags, or this will cause
// the error "Parameter validation failed: Missing required parameter in
// Tags[N]:Value". When provisioning a product that's been added to a portfolio,
// you must grant your user, group, or role access to the portfolio. For more
// information, see Granting users access (https://docs.aws.amazon.com/servicecatalog/latest/adminguide/catalogs_portfolios_users.html)
// in the Service Catalog User Guide.
func (c *Client) ProvisionProduct(ctx context.Context, params *ProvisionProductInput, optFns ...func(*Options)) (*ProvisionProductOutput, error) {
	if params == nil { params = &ProvisionProductInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ProvisionProduct", params, optFns, c.addOperationProvisionProductMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ProvisionProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionProductInput struct {
	
	// An idempotency token that uniquely identifies the provisioning request.
	//
	// This member is required.
	ProvisionToken *string
	
	// A user-friendly name for the provisioned product. This value must be unique for
	// the Amazon Web Services account and cannot be updated after the product is
	// provisioned.
	//
	// This member is required.
	ProvisionedProductName *string
	
	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string
	
	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []string
	
	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths . You must provide the name or ID, but
	// not both.
	PathId *string
	
	// The name of the path. You must provide the name or ID, but not both.
	PathName *string
	
	// The product identifier. You must provide the name or ID, but not both.
	ProductId *string
	
	// The name of the product. You must provide the name or ID, but not both.
	ProductName *string
	
	// The identifier of the provisioning artifact. You must provide the name or ID,
	// but not both.
	ProvisioningArtifactId *string
	
	// The name of the provisioning artifact. You must provide the name or ID, but not
	// both.
	ProvisioningArtifactName *string
	
	// Parameters specified by the administrator that are required for provisioning
	// the product.
	ProvisioningParameters []types.ProvisioningParameter
	
	// An object that contains information about the provisioning preferences for a
	// stack set.
	ProvisioningPreferences *types.ProvisioningPreferences
	
	// One or more tags.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type ProvisionProductOutput struct {
	
	// Information about the result of provisioning the product.
	RecordDetail *types.RecordDetail
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationProvisionProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpProvisionProduct{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpProvisionProduct{}, middleware.After)
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
	if err = addProvisionProductResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opProvisionProductMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpProvisionProductValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionProduct(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpProvisionProduct struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpProvisionProduct) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpProvisionProduct) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*ProvisionProductInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *ProvisionProductInput ")}
	
	if input.ProvisionToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ProvisionToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opProvisionProductMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpProvisionProduct{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opProvisionProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "servicecatalog",
	OperationName: "ProvisionProduct",
	}
}

type opProvisionProductResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opProvisionProductResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opProvisionProductResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "servicecatalog"
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
	        signingName = "servicecatalog"
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
	        v4aScheme.SigningName = aws.String("servicecatalog")
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

func addProvisionProductResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opProvisionProductResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
