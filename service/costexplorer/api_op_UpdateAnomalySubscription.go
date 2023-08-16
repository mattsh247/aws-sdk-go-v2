// Code generated by smithy-go-codegen DO NOT EDIT.


package costexplorer

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
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Updates an existing cost anomaly subscription. Specify the fields that you want
// to update. Omitted fields are unchanged. The JSON below describes the generic
// construct for each type. See Request Parameters (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_UpdateAnomalySubscription.html#API_UpdateAnomalySubscription_RequestParameters)
// for possible values as they apply to AnomalySubscription .
func (c *Client) UpdateAnomalySubscription(ctx context.Context, params *UpdateAnomalySubscriptionInput, optFns ...func(*Options)) (*UpdateAnomalySubscriptionOutput, error) {
	if params == nil { params = &UpdateAnomalySubscriptionInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateAnomalySubscription", params, optFns, c.addOperationUpdateAnomalySubscriptionMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateAnomalySubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnomalySubscriptionInput struct {
	
	// A cost anomaly subscription Amazon Resource Name (ARN).
	//
	// This member is required.
	SubscriptionArn *string
	
	// The update to the frequency value that subscribers receive notifications.
	Frequency types.AnomalySubscriptionFrequency
	
	// A list of cost anomaly monitor ARNs.
	MonitorArnList []string
	
	// The update to the subscriber list.
	Subscribers []types.Subscriber
	
	// The new name of the subscription.
	SubscriptionName *string
	
	// (deprecated) The update to the threshold value for receiving notifications.
	// This field has been deprecated. To update a threshold, use ThresholdExpression.
	// Continued use of Threshold will be treated as shorthand syntax for a
	// ThresholdExpression. You can specify either Threshold or ThresholdExpression,
	// but not both.
	//
	// Deprecated: Threshold has been deprecated in favor of ThresholdExpression
	Threshold *float64
	
	// The update to the Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object used to specify the anomalies that you want to generate alerts for. This
	// supports dimensions and nested expressions. The supported dimensions are
	// ANOMALY_TOTAL_IMPACT_ABSOLUTE and ANOMALY_TOTAL_IMPACT_PERCENTAGE ,
	// corresponding to an anomaly’s TotalImpact and TotalImpactPercentage,
	// respectively (see Impact (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Impact.html)
	// for more details). The supported nested expression types are AND and OR . The
	// match option GREATER_THAN_OR_EQUAL is required. Values must be numbers between
	// 0 and 10,000,000,000 in string format. You can specify either Threshold or
	// ThresholdExpression, but not both. The following are examples of valid
	// ThresholdExpressions:
	//   - Absolute threshold: { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } }
	//   - Percentage threshold: { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } }
	//   - AND two thresholds together: { "And": [ { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } }, { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } } ] }
	//   - OR two thresholds together: { "Or": [ { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_ABSOLUTE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } }, { "Dimensions": { "Key":
	//   "ANOMALY_TOTAL_IMPACT_PERCENTAGE", "MatchOptions": [ "GREATER_THAN_OR_EQUAL" ],
	//   "Values": [ "100" ] } } ] }
	ThresholdExpression *types.Expression
	
	noSmithyDocumentSerde
}

type UpdateAnomalySubscriptionOutput struct {
	
	// A cost anomaly subscription ARN.
	//
	// This member is required.
	SubscriptionArn *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnomalySubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAnomalySubscription{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAnomalySubscription{}, middleware.After)
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
	if err = addUpdateAnomalySubscriptionResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateAnomalySubscriptionValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnomalySubscription(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAnomalySubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ce",
	OperationName: "UpdateAnomalySubscription",
	}
}

type opUpdateAnomalySubscriptionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateAnomalySubscriptionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateAnomalySubscriptionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ce"
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
	        signingName = "ce"
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
	        v4aScheme.SigningName = aws.String("ce")
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

func addUpdateAnomalySubscriptionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateAnomalySubscriptionResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
