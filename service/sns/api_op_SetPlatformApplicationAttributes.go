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
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Sets the attributes of the platform application object for the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
// information, see Using Amazon SNS Mobile Push Notifications (https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html)
// . For information on configuring attributes for message delivery status, see
// Using Amazon SNS Application Attributes for Message Delivery Status (https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html)
// .
func (c *Client) SetPlatformApplicationAttributes(ctx context.Context, params *SetPlatformApplicationAttributesInput, optFns ...func(*Options)) (*SetPlatformApplicationAttributesOutput, error) {
	if params == nil { params = &SetPlatformApplicationAttributesInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "SetPlatformApplicationAttributes", params, optFns, c.addOperationSetPlatformApplicationAttributesMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*SetPlatformApplicationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for SetPlatformApplicationAttributes action.
type SetPlatformApplicationAttributesInput struct {
	
	// A map of the platform application attributes. Attributes in this map include
	// the following:
	//   - PlatformCredential – The credential received from the notification service.
	//   - For ADM, PlatformCredential is client secret.
	//   - For Apple Services using certificate credentials, PlatformCredential is
	//   private key.
	//   - For Apple Services using token credentials, PlatformCredential is signing
	//   key.
	//   - For GCM (Firebase Cloud Messaging), PlatformCredential is API key.
	//
	//   - PlatformPrincipal – The principal received from the notification service.
	//   - For ADM, PlatformPrincipal is client id.
	//   - For Apple Services using certificate credentials, PlatformPrincipal is SSL
	//   certificate.
	//   - For Apple Services using token credentials, PlatformPrincipal is signing key
	//   ID.
	//   - For GCM (Firebase Cloud Messaging), there is no PlatformPrincipal .
	//
	//   - EventEndpointCreated – Topic ARN to which EndpointCreated event
	//   notifications are sent.
	//   - EventEndpointDeleted – Topic ARN to which EndpointDeleted event
	//   notifications are sent.
	//   - EventEndpointUpdated – Topic ARN to which EndpointUpdate event notifications
	//   are sent.
	//   - EventDeliveryFailure – Topic ARN to which DeliveryFailure event
	//   notifications are sent upon Direct Publish delivery failure (permanent) to one
	//   of the application's endpoints.
	//   - SuccessFeedbackRoleArn – IAM role ARN used to give Amazon SNS write access
	//   to use CloudWatch Logs on your behalf.
	//   - FailureFeedbackRoleArn – IAM role ARN used to give Amazon SNS write access
	//   to use CloudWatch Logs on your behalf.
	//   - SuccessFeedbackSampleRate – Sample rate percentage (0-100) of successfully
	//   delivered messages.
	// The following attributes only apply to APNs token-based authentication:
	//   - ApplePlatformTeamID – The identifier that's assigned to your Apple developer
	//   account team.
	//   - ApplePlatformBundleID – The bundle identifier that's assigned to your iOS
	//   app.
	//
	// This member is required.
	Attributes map[string]string
	
	// PlatformApplicationArn for SetPlatformApplicationAttributes action.
	//
	// This member is required.
	PlatformApplicationArn *string
	
	noSmithyDocumentSerde
}

type SetPlatformApplicationAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationSetPlatformApplicationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetPlatformApplicationAttributes{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetPlatformApplicationAttributes{}, middleware.After)
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
	if err = addSetPlatformApplicationAttributesResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpSetPlatformApplicationAttributesValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetPlatformApplicationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "sns",
	OperationName: "SetPlatformApplicationAttributes",
	}
}

type opSetPlatformApplicationAttributesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opSetPlatformApplicationAttributesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opSetPlatformApplicationAttributesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addSetPlatformApplicationAttributesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opSetPlatformApplicationAttributesResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
