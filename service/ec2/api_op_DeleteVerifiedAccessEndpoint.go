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

// Delete an Amazon Web Services Verified Access endpoint.
func (c *Client) DeleteVerifiedAccessEndpoint(ctx context.Context, params *DeleteVerifiedAccessEndpointInput, optFns ...func(*Options)) (*DeleteVerifiedAccessEndpointOutput, error) {
	if params == nil { params = &DeleteVerifiedAccessEndpointInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DeleteVerifiedAccessEndpoint", params, optFns, c.addOperationDeleteVerifiedAccessEndpointMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DeleteVerifiedAccessEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVerifiedAccessEndpointInput struct {
	
	// The ID of the Verified Access endpoint.
	//
	// This member is required.
	VerifiedAccessEndpointId *string
	
	// A unique, case-sensitive token that you provide to ensure idempotency of your
	// modification request. For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// .
	ClientToken *string
	
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool
	
	noSmithyDocumentSerde
}

type DeleteVerifiedAccessEndpointOutput struct {
	
	// The ID of the Verified Access endpoint.
	VerifiedAccessEndpoint *types.VerifiedAccessEndpoint
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVerifiedAccessEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteVerifiedAccessEndpoint{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteVerifiedAccessEndpoint{}, middleware.After)
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
	if err = addDeleteVerifiedAccessEndpointResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opDeleteVerifiedAccessEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDeleteVerifiedAccessEndpointValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVerifiedAccessEndpoint(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteVerifiedAccessEndpoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteVerifiedAccessEndpoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteVerifiedAccessEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*DeleteVerifiedAccessEndpointInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteVerifiedAccessEndpointInput ")}
	
	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteVerifiedAccessEndpointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteVerifiedAccessEndpoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteVerifiedAccessEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ec2",
	OperationName: "DeleteVerifiedAccessEndpoint",
	}
}

type opDeleteVerifiedAccessEndpointResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDeleteVerifiedAccessEndpointResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteVerifiedAccessEndpointResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDeleteVerifiedAccessEndpointResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDeleteVerifiedAccessEndpointResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
