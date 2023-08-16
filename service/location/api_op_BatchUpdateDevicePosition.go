// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads position update data for one or more devices to a tracker resource (up
// to 10 devices per batch). Amazon Location uses the data when it reports the last
// known device position and position history. Amazon Location retains location
// data for 30 days. Position updates are handled based on the PositionFiltering
// property of the tracker. When PositionFiltering is set to TimeBased , updates
// are evaluated against linked geofence collections, and location data is stored
// at a maximum of one position per 30 second interval. If your update frequency is
// more often than every 30 seconds, only one update per 30 seconds is stored for
// each unique device ID. When PositionFiltering is set to DistanceBased
// filtering, location data is stored and evaluated against linked geofence
// collections only if the device has moved more than 30 m (98.4 ft). When
// PositionFiltering is set to AccuracyBased filtering, location data is stored
// and evaluated against linked geofence collections only if the device has moved
// more than the measured accuracy. For example, if two consecutive updates from a
// device have a horizontal accuracy of 5 m and 10 m, the second update is neither
// stored or evaluated if the device has moved less than 15 m. If PositionFiltering
// is set to AccuracyBased filtering, Amazon Location uses the default value {
// "Horizontal": 0} when accuracy is not provided on a DevicePositionUpdate .
func (c *Client) BatchUpdateDevicePosition(ctx context.Context, params *BatchUpdateDevicePositionInput, optFns ...func(*Options)) (*BatchUpdateDevicePositionOutput, error) {
	if params == nil {
		params = &BatchUpdateDevicePositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateDevicePosition", params, optFns, c.addOperationBatchUpdateDevicePositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateDevicePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateDevicePositionInput struct {

	// The name of the tracker resource to update.
	//
	// This member is required.
	TrackerName *string

	// Contains the position update details for each device, up to 10 devices.
	//
	// This member is required.
	Updates []types.DevicePositionUpdate

	noSmithyDocumentSerde
}

type BatchUpdateDevicePositionOutput struct {

	// Contains error details for each device that failed to update its position.
	//
	// This member is required.
	Errors []types.BatchUpdateDevicePositionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateDevicePositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateDevicePosition{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchUpdateDevicePositionMiddleware(stack); err != nil {
		return err
	}
	if err = addBatchUpdateDevicePositionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBatchUpdateDevicePositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateDevicePosition(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchUpdateDevicePositionMiddleware struct {
}

func (*endpointPrefix_opBatchUpdateDevicePositionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchUpdateDevicePositionMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchUpdateDevicePositionMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchUpdateDevicePositionMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchUpdateDevicePosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchUpdateDevicePosition",
	}
}

type opBatchUpdateDevicePositionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opBatchUpdateDevicePositionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opBatchUpdateDevicePositionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addBatchUpdateDevicePositionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opBatchUpdateDevicePositionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
