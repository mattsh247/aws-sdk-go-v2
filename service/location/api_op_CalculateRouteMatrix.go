// Code generated by smithy-go-codegen DO NOT EDIT.


package location

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
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

//  Calculates a route matrix (https://docs.aws.amazon.com/location/latest/developerguide/calculate-route-matrix.html)
// given the following required parameters: DeparturePositions and
// DestinationPositions . CalculateRouteMatrix calculates routes and returns the
// travel time and travel distance from each departure position to each destination
// position in the request. For example, given departure positions A and B, and
// destination positions X and Y, CalculateRouteMatrix will return time and
// distance for routes from A to X, A to Y, B to X, and B to Y (in that order). The
// number of results returned (and routes calculated) will be the number of
// DeparturePositions times the number of DestinationPositions . Your account is
// charged for each route calculated, not the number of requests. Requires that you
// first create a route calculator resource (https://docs.aws.amazon.com/location-routes/latest/APIReference/API_CreateRouteCalculator.html)
// . By default, a request that doesn't specify a departure time uses the best time
// of day to travel with the best traffic conditions when calculating routes.
// Additional options include:
//   - Specifying a departure time (https://docs.aws.amazon.com/location/latest/developerguide/departure-time.html)
//   using either DepartureTime or DepartNow . This calculates routes based on
//   predictive traffic data at the given time. You can't specify both
//   DepartureTime and DepartNow in a single request. Specifying both parameters
//   returns a validation error.
//   - Specifying a travel mode (https://docs.aws.amazon.com/location/latest/developerguide/travel-mode.html)
//   using TravelMode sets the transportation mode used to calculate the routes. This
//   also lets you specify additional route preferences in CarModeOptions if
//   traveling by Car , or TruckModeOptions if traveling by Truck .
func (c *Client) CalculateRouteMatrix(ctx context.Context, params *CalculateRouteMatrixInput, optFns ...func(*Options)) (*CalculateRouteMatrixOutput, error) {
	if params == nil { params = &CalculateRouteMatrixInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CalculateRouteMatrix", params, optFns, c.addOperationCalculateRouteMatrixMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CalculateRouteMatrixOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CalculateRouteMatrixInput struct {
	
	// The name of the route calculator resource that you want to use to calculate the
	// route matrix.
	//
	// This member is required.
	CalculatorName *string
	
	// The list of departure (origin) positions for the route matrix. An array of
	// points, each of which is itself a 2-value array defined in WGS 84 (https://earth-info.nga.mil/GandG/wgs84/index.html)
	// format: [longitude, latitude] . For example, [-123.115, 49.285] . Depending on
	// the data provider selected in the route calculator resource there may be
	// additional restrictions on the inputs you can choose. See Position restrictions (https://docs.aws.amazon.com/location/latest/developerguide/calculate-route-matrix.html#matrix-routing-position-limits)
	// in the Amazon Location Service Developer Guide. For route calculators that use
	// Esri as the data provider, if you specify a departure that's not located on a
	// road, Amazon Location moves the position to the nearest road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// . The snapped value is available in the result in SnappedDeparturePositions .
	// Valid Values: [-180 to 180,-90 to 90]
	//
	// This member is required.
	DeparturePositions [][]float64
	
	// The list of destination positions for the route matrix. An array of points,
	// each of which is itself a 2-value array defined in WGS 84 (https://earth-info.nga.mil/GandG/wgs84/index.html)
	// format: [longitude, latitude] . For example, [-122.339, 47.615] Depending on
	// the data provider selected in the route calculator resource there may be
	// additional restrictions on the inputs you can choose. See Position restrictions (https://docs.aws.amazon.com/location/latest/developerguide/calculate-route-matrix.html#matrix-routing-position-limits)
	// in the Amazon Location Service Developer Guide. For route calculators that use
	// Esri as the data provider, if you specify a destination that's not located on a
	// road, Amazon Location moves the position to the nearest road (https://docs.aws.amazon.com/location/latest/developerguide/snap-to-nearby-road.html)
	// . The snapped value is available in the result in SnappedDestinationPositions .
	// Valid Values: [-180 to 180,-90 to 90]
	//
	// This member is required.
	DestinationPositions [][]float64
	
	// Specifies route preferences when traveling by Car , such as avoiding routes that
	// use ferries or tolls. Requirements: TravelMode must be specified as Car .
	CarModeOptions *types.CalculateRouteCarModeOptions
	
	// Sets the time of departure as the current time. Uses the current time to
	// calculate the route matrix. You can't set both DepartureTime and DepartNow . If
	// neither is set, the best time of day to travel with the best traffic conditions
	// is used to calculate the route matrix. Default Value: false Valid Values: false
	// | true
	DepartNow *bool
	
	// Specifies the desired time of departure. Uses the given time to calculate the
	// route matrix. You can't set both DepartureTime and DepartNow . If neither is
	// set, the best time of day to travel with the best traffic conditions is used to
	// calculate the route matrix. Setting a departure time in the past returns a 400
	// ValidationException error.
	//   - In ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	//   YYYY-MM-DDThh:mm:ss.sssZ . For example, 2020–07-2T12:15:20.000Z+01:00
	DepartureTime *time.Time
	
	// Set the unit system to specify the distance. Default Value: Kilometers
	DistanceUnit types.DistanceUnit
	
	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string
	
	// Specifies the mode of transport when calculating a route. Used in estimating
	// the speed of travel and road compatibility. The TravelMode you specify also
	// determines how you specify route preferences:
	//   - If traveling by Car use the CarModeOptions parameter.
	//   - If traveling by Truck use the TruckModeOptions parameter.
	// Bicycle or Motorcycle are only valid when using Grab as a data provider, and
	// only within Southeast Asia. Truck is not available for Grab. For more
	// information about using Grab as a data provider, see GrabMaps (https://docs.aws.amazon.com/location/latest/developerguide/grab.html)
	// in the Amazon Location Service Developer Guide. Default Value: Car
	TravelMode types.TravelMode
	
	// Specifies route preferences when traveling by Truck , such as avoiding routes
	// that use ferries or tolls, and truck specifications to consider when choosing an
	// optimal road. Requirements: TravelMode must be specified as Truck .
	TruckModeOptions *types.CalculateRouteTruckModeOptions
	
	noSmithyDocumentSerde
}

// Returns the result of the route matrix calculation.
type CalculateRouteMatrixOutput struct {
	
	// The calculated route matrix containing the results for all pairs of
	// DeparturePositions to DestinationPositions . Each row corresponds to one entry
	// in DeparturePositions . Each entry in the row corresponds to the route from that
	// entry in DeparturePositions to an entry in DestinationPositions .
	//
	// This member is required.
	RouteMatrix [][]types.RouteMatrixEntry
	
	// Contains information about the route matrix, DataSource , DistanceUnit ,
	// RouteCount and ErrorCount .
	//
	// This member is required.
	Summary *types.CalculateRouteMatrixSummary
	
	// For routes calculated using an Esri route calculator resource, departure
	// positions are snapped to the closest road. For Esri route calculator resources,
	// this returns the list of departure/origin positions used for calculation of the
	// RouteMatrix .
	SnappedDeparturePositions [][]float64
	
	// The list of destination positions for the route matrix used for calculation of
	// the RouteMatrix .
	SnappedDestinationPositions [][]float64
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCalculateRouteMatrixMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCalculateRouteMatrix{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCalculateRouteMatrix{}, middleware.After)
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
	if err = addEndpointPrefix_opCalculateRouteMatrixMiddleware(stack); err != nil {
	return err
	}
	if err = addCalculateRouteMatrixResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCalculateRouteMatrixValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCalculateRouteMatrix(options.Region, ), middleware.Before); err != nil {
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

type endpointPrefix_opCalculateRouteMatrixMiddleware struct {
}

func (*endpointPrefix_opCalculateRouteMatrixMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCalculateRouteMatrixMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}
	
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}
	
	req.URL.Host = "routes." + req.URL.Host
	
	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCalculateRouteMatrixMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCalculateRouteMatrixMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCalculateRouteMatrix(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "geo",
	OperationName: "CalculateRouteMatrix",
	}
}

type opCalculateRouteMatrixResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCalculateRouteMatrixResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCalculateRouteMatrixResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "geo"
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
	        signingName = "geo"
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
	        v4aScheme.SigningName = aws.String("geo")
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

func addCalculateRouteMatrixResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCalculateRouteMatrixResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
