// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about health metrics for your Lightsail load balancer.
// Metrics report the utilization of your resources, and the error counts generated
// by them. Monitor and collect metric data regularly to maintain the reliability,
// availability, and performance of your resources.
func (c *Client) GetLoadBalancerMetricData(ctx context.Context, params *GetLoadBalancerMetricDataInput, optFns ...func(*Options)) (*GetLoadBalancerMetricDataOutput, error) {
	if params == nil {
		params = &GetLoadBalancerMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoadBalancerMetricData", params, optFns, c.addOperationGetLoadBalancerMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoadBalancerMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoadBalancerMetricDataInput struct {

	// The end time of the period.
	//
	// This member is required.
	EndTime *time.Time

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The metric for which you want to return information. Valid load balancer metric
	// names are listed below, along with the most useful statistics to include in
	// your request, and the published unit value.
	//   - ClientTLSNegotiationErrorCount - The number of TLS connections initiated by
	//   the client that did not establish a session with the load balancer due to a TLS
	//   error generated by the load balancer. Possible causes include a mismatch of
	//   ciphers or protocols. Statistics : The most useful statistic is Sum . Unit :
	//   The published unit is Count .
	//   - HealthyHostCount - The number of target instances that are considered
	//   healthy. Statistics : The most useful statistic are Average , Minimum , and
	//   Maximum . Unit : The published unit is Count .
	//   - HTTPCode_Instance_2XX_Count - The number of HTTP 2XX response codes
	//   generated by the target instances. This does not include any response codes
	//   generated by the load balancer. Statistics : The most useful statistic is Sum
	//   . Note that Minimum , Maximum , and Average all return 1 . Unit : The
	//   published unit is Count .
	//   - HTTPCode_Instance_3XX_Count - The number of HTTP 3XX response codes
	//   generated by the target instances. This does not include any response codes
	//   generated by the load balancer. Statistics : The most useful statistic is Sum
	//   . Note that Minimum , Maximum , and Average all return 1 . Unit : The
	//   published unit is Count .
	//   - HTTPCode_Instance_4XX_Count - The number of HTTP 4XX response codes
	//   generated by the target instances. This does not include any response codes
	//   generated by the load balancer. Statistics : The most useful statistic is Sum
	//   . Note that Minimum , Maximum , and Average all return 1 . Unit : The
	//   published unit is Count .
	//   - HTTPCode_Instance_5XX_Count - The number of HTTP 5XX response codes
	//   generated by the target instances. This does not include any response codes
	//   generated by the load balancer. Statistics : The most useful statistic is Sum
	//   . Note that Minimum , Maximum , and Average all return 1 . Unit : The
	//   published unit is Count .
	//   - HTTPCode_LB_4XX_Count - The number of HTTP 4XX client error codes that
	//   originated from the load balancer. Client errors are generated when requests are
	//   malformed or incomplete. These requests were not received by the target
	//   instance. This count does not include response codes generated by the target
	//   instances. Statistics : The most useful statistic is Sum . Note that Minimum ,
	//   Maximum , and Average all return 1 . Unit : The published unit is Count .
	//   - HTTPCode_LB_5XX_Count - The number of HTTP 5XX server error codes that
	//   originated from the load balancer. This does not include any response codes
	//   generated by the target instance. This metric is reported if there are no
	//   healthy instances attached to the load balancer, or if the request rate exceeds
	//   the capacity of the instances (spillover) or the load balancer. Statistics :
	//   The most useful statistic is Sum . Note that Minimum , Maximum , and Average
	//   all return 1 . Unit : The published unit is Count .
	//   - InstanceResponseTime - The time elapsed, in seconds, after the request
	//   leaves the load balancer until a response from the target instance is received.
	//   Statistics : The most useful statistic is Average . Unit : The published unit
	//   is Seconds .
	//   - RejectedConnectionCount - The number of connections that were rejected
	//   because the load balancer had reached its maximum number of connections.
	//   Statistics : The most useful statistic is Sum . Unit : The published unit is
	//   Count .
	//   - RequestCount - The number of requests processed over IPv4. This count
	//   includes only the requests with a response generated by a target instance of the
	//   load balancer. Statistics : The most useful statistic is Sum . Note that
	//   Minimum , Maximum , and Average all return 1 . Unit : The published unit is
	//   Count .
	//   - UnhealthyHostCount - The number of target instances that are considered
	//   unhealthy. Statistics : The most useful statistic are Average , Minimum , and
	//   Maximum . Unit : The published unit is Count .
	//
	// This member is required.
	MetricName types.LoadBalancerMetricName

	// The granularity, in seconds, of the returned data points.
	//
	// This member is required.
	Period int32

	// The start time of the period.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//   - Minimum - The lowest value observed during the specified period. Use this
	//   value to determine low volumes of activity for your application.
	//   - Maximum - The highest value observed during the specified period. Use this
	//   value to determine high volumes of activity for your application.
	//   - Sum - All values submitted for the matching metric added together. You can
	//   use this statistic to determine the total volume of a metric.
	//   - Average - The value of Sum / SampleCount during the specified period. By
	//   comparing this statistic with the Minimum and Maximum values, you can determine
	//   the full scope of a metric and how close the average use is to the Minimum and
	//   Maximum values. This comparison helps you to know when to increase or decrease
	//   your resources.
	//   - SampleCount - The count, or number, of data points used for the statistical
	//   calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The unit for the metric data request. Valid units depend on the metric data
	// being requested. For the valid units with each available metric, see the
	// metricName parameter.
	//
	// This member is required.
	Unit types.MetricUnit

	noSmithyDocumentSerde
}

type GetLoadBalancerMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.LoadBalancerMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLoadBalancerMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLoadBalancerMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLoadBalancerMetricData{}, middleware.After)
	if err != nil {
		return err
	}
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
	if err = addGetLoadBalancerMetricDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetLoadBalancerMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoadBalancerMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLoadBalancerMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetLoadBalancerMetricData",
	}
}

type opGetLoadBalancerMetricDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetLoadBalancerMetricDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetLoadBalancerMetricDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "lightsail"
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
				signingName = "lightsail"
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
				v4aScheme.SigningName = aws.String("lightsail")
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

func addGetLoadBalancerMetricDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetLoadBalancerMetricDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
