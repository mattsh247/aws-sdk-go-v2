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

// Retrieves a forecast for how much Amazon Web Services predicts that you will
// spend over the forecast time period that you select, based on your past costs.
func (c *Client) GetCostForecast(ctx context.Context, params *GetCostForecastInput, optFns ...func(*Options)) (*GetCostForecastOutput, error) {
	if params == nil { params = &GetCostForecastInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetCostForecast", params, optFns, c.addOperationGetCostForecastMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetCostForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostForecastInput struct {
	
	// How granular you want the forecast to be. You can get 3 months of DAILY
	// forecasts or 12 months of MONTHLY forecasts. The GetCostForecast operation
	// supports only DAILY and MONTHLY granularities.
	//
	// This member is required.
	Granularity types.Granularity
	
	// Which metric Cost Explorer uses to create your forecast. For more information
	// about blended and unblended rates, see Why does the "blended" annotation appear
	// on some line items in my bill? (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/)
	// . Valid values for a GetCostForecast call are the following:
	//   - AMORTIZED_COST
	//   - BLENDED_COST
	//   - NET_AMORTIZED_COST
	//   - NET_UNBLENDED_COST
	//   - UNBLENDED_COST
	//
	// This member is required.
	Metric types.Metric
	
	// The period of time that you want the forecast to cover. The start date must be
	// equal to or no later than the current date to avoid a validation error.
	//
	// This member is required.
	TimePeriod *types.DateInterval
	
	// The filters that you want to use to filter your forecast. The GetCostForecast
	// API supports filtering by the following dimensions:
	//   - AZ
	//   - INSTANCE_TYPE
	//   - LINKED_ACCOUNT
	//   - LINKED_ACCOUNT_NAME
	//   - OPERATION
	//   - PURCHASE_TYPE
	//   - REGION
	//   - SERVICE
	//   - USAGE_TYPE
	//   - USAGE_TYPE_GROUP
	//   - RECORD_TYPE
	//   - OPERATING_SYSTEM
	//   - TENANCY
	//   - SCOPE
	//   - PLATFORM
	//   - SUBSCRIPTION_ID
	//   - LEGAL_ENTITY_NAME
	//   - DEPLOYMENT_OPTION
	//   - DATABASE_ENGINE
	//   - INSTANCE_TYPE_FAMILY
	//   - BILLING_ENTITY
	//   - RESERVATION_ID
	//   - SAVINGS_PLAN_ARN
	Filter *types.Expression
	
	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence level.
	// The higher the confidence level, the more confident Cost Explorer is about the
	// actual value falling in the prediction interval. Higher confidence levels result
	// in wider prediction intervals.
	PredictionIntervalLevel *int32
	
	noSmithyDocumentSerde
}

type GetCostForecastOutput struct {
	
	// The forecasts for your query, in order. For DAILY forecasts, this is a list of
	// days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []types.ForecastResult
	
	// How much you are forecasted to spend over the forecast period, in USD .
	Total *types.MetricValue
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostForecast{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostForecast{}, middleware.After)
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
	if err = addGetCostForecastResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpGetCostForecastValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostForecast(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ce",
	OperationName: "GetCostForecast",
	}
}

type opGetCostForecastResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetCostForecastResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetCostForecastResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addGetCostForecastResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetCostForecastResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
