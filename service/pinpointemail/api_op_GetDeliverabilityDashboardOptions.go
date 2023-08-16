// Code generated by smithy-go-codegen DO NOT EDIT.


package pinpointemail

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
	"time"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Retrieve information about the status of the Deliverability dashboard for your
// Amazon Pinpoint account. When the Deliverability dashboard is enabled, you gain
// access to reputation, deliverability, and other metrics for the domains that you
// use to send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests. When you use the Deliverability dashboard, you
// pay a monthly subscription charge, in addition to any other fees that you accrue
// by using Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing (http://aws.amazon.com/pinpoint/pricing/)
// .
func (c *Client) GetDeliverabilityDashboardOptions(ctx context.Context, params *GetDeliverabilityDashboardOptionsInput, optFns ...func(*Options)) (*GetDeliverabilityDashboardOptionsOutput, error) {
	if params == nil { params = &GetDeliverabilityDashboardOptionsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetDeliverabilityDashboardOptions", params, optFns, c.addOperationGetDeliverabilityDashboardOptionsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetDeliverabilityDashboardOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve information about the status of the Deliverability dashboard for your
// Amazon Pinpoint account. When the Deliverability dashboard is enabled, you gain
// access to reputation, deliverability, and other metrics for the domains that you
// use to send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests. When you use the Deliverability dashboard, you
// pay a monthly subscription charge, in addition to any other fees that you accrue
// by using Amazon Pinpoint. For more information about the features and cost of a
// Deliverability dashboard subscription, see Amazon Pinpoint Pricing (http://aws.amazon.com/pinpoint/pricing/)
// .
type GetDeliverabilityDashboardOptionsInput struct {
	
	noSmithyDocumentSerde
}

// An object that shows the status of the Deliverability dashboard for your Amazon
// Pinpoint account.
type GetDeliverabilityDashboardOptionsOutput struct {
	
	// Specifies whether the Deliverability dashboard is enabled for your Amazon
	// Pinpoint account. If this value is true , the dashboard is enabled.
	//
	// This member is required.
	DashboardEnabled bool
	
	// The current status of your Deliverability dashboard subscription. If this value
	// is PENDING_EXPIRATION , your subscription is scheduled to expire at the end of
	// the current calendar month.
	AccountStatus types.DeliverabilityDashboardAccountStatus
	
	// An array of objects, one for each verified domain that you use to send email
	// and currently has an active Deliverability dashboard subscription that isn’t
	// scheduled to expire at the end of the current calendar month.
	ActiveSubscribedDomains []types.DomainDeliverabilityTrackingOption
	
	// An array of objects, one for each verified domain that you use to send email
	// and currently has an active Deliverability dashboard subscription that's
	// scheduled to expire at the end of the current calendar month.
	PendingExpirationSubscribedDomains []types.DomainDeliverabilityTrackingOption
	
	// The date, in Unix time format, when your current subscription to the
	// Deliverability dashboard is scheduled to expire, if your subscription is
	// scheduled to expire at the end of the current calendar month. This value is null
	// if you have an active subscription that isn’t due to expire at the end of the
	// month.
	SubscriptionExpiryDate *time.Time
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeliverabilityDashboardOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeliverabilityDashboardOptions{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeliverabilityDashboardOptions{}, middleware.After)
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
	if err = addGetDeliverabilityDashboardOptionsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeliverabilityDashboardOptions(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDeliverabilityDashboardOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ses",
	OperationName: "GetDeliverabilityDashboardOptions",
	}
}

type opGetDeliverabilityDashboardOptionsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetDeliverabilityDashboardOptionsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetDeliverabilityDashboardOptionsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ses"
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
	        signingName = "ses"
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
	        v4aScheme.SigningName = aws.String("ses")
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

func addGetDeliverabilityDashboardOptionsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetDeliverabilityDashboardOptionsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
