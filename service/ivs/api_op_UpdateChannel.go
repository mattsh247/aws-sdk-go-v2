// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a channel's configuration. Live channels cannot be updated. You must
// stop the ongoing stream, update the channel, and restart the stream for the
// changes to take effect.
func (c *Client) UpdateChannel(ctx context.Context, params *UpdateChannelInput, optFns ...func(*Options)) (*UpdateChannelOutput, error) {
	if params == nil {
		params = &UpdateChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChannel", params, optFns, c.addOperationUpdateChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChannelInput struct {

	// ARN of the channel to be updated.
	//
	// This member is required.
	Arn *string

	// Whether the channel is private (enabled for playback authorization).
	Authorized bool

	// Whether the channel allows insecure RTMP ingest. Default: false .
	InsecureIngest bool

	// Channel latency mode. Use NORMAL to broadcast and deliver live video up to Full
	// HD. Use LOW for near-real-time interaction with viewers. (Note: In the Amazon
	// IVS console, LOW and NORMAL correspond to Ultra-low and Standard, respectively.)
	LatencyMode types.ChannelLatencyMode

	// Channel name.
	Name *string

	// Optional transcode preset for the channel. This is selectable only for
	// ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default
	// preset is HIGHER_BANDWIDTH_DELIVERY . For other channel types ( BASIC and
	// STANDARD ), preset is the empty string ( "" ).
	Preset types.TranscodePreset

	// Recording-configuration ARN. If this is set to an empty string, recording is
	// disabled. A value other than an empty string indicates that recording is enabled
	RecordingConfigurationArn *string

	// Channel type, which determines the allowable resolution and bitrate. If you
	// exceed the allowable input resolution or bitrate, the stream probably will
	// disconnect immediately. Some types generate multiple qualities (renditions) from
	// the original input; this automatically gives viewers the best experience for
	// their devices and network conditions. Some types provide transcoded video;
	// transcoding allows higher playback quality across a range of download speeds.
	// Default: STANDARD . Valid values:
	//   - BASIC : Video is transmuxed: Amazon IVS delivers the original input quality
	//   to viewers. The viewer’s video-quality choice is limited to the original input.
	//   Input resolution can be up to 1080p and bitrate can be up to 1.5 Mbps for 480p
	//   and up to 3.5 Mbps for resolutions between 480p and 1080p. Original audio is
	//   passed through.
	//   - STANDARD : Video is transcoded: multiple qualities are generated from the
	//   original input, to automatically give viewers the best experience for their
	//   devices and network conditions. Transcoding allows higher playback quality
	//   across a range of download speeds. Resolution can be up to 1080p and bitrate can
	//   be up to 8.5 Mbps. Audio is transcoded only for renditions 360p and below; above
	//   that, audio is passed through. This is the default when you create a channel.
	//   - ADVANCED_SD : Video is transcoded; multiple qualities are generated from the
	//   original input, to automatically give viewers the best experience for their
	//   devices and network conditions. Input resolution can be up to 1080p and bitrate
	//   can be up to 8.5 Mbps; output is capped at SD quality (480p). You can select an
	//   optional transcode preset (see below). Audio for all renditions is transcoded,
	//   and an audio-only rendition is available.
	//   - ADVANCED_HD : Video is transcoded; multiple qualities are generated from the
	//   original input, to automatically give viewers the best experience for their
	//   devices and network conditions. Input resolution can be up to 1080p and bitrate
	//   can be up to 8.5 Mbps; output is capped at HD quality (720p). You can select an
	//   optional transcode preset (see below). Audio for all renditions is transcoded,
	//   and an audio-only rendition is available.
	// Optional transcode presets (available for the ADVANCED types) allow you to
	// trade off available download bandwidth and video quality, to optimize the
	// viewing experience. There are two presets:
	//   - Constrained bandwidth delivery uses a lower bitrate for each quality level.
	//   Use it if you have low download bandwidth and/or simple video content (e.g.,
	//   talking heads)
	//   - Higher bandwidth delivery uses a higher bitrate for each quality level. Use
	//   it if you have high download bandwidth and/or complex video content (e.g.,
	//   flashes and quick scene changes).
	Type types.ChannelType

	noSmithyDocumentSerde
}

type UpdateChannelOutput struct {

	// Object specifying a channel.
	Channel *types.Channel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannel{}, middleware.After)
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
	if err = addUpdateChannelResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "UpdateChannel",
	}
}

type opUpdateChannelResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateChannelResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateChannelResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ivs"
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
				signingName = "ivs"
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
				v4aScheme.SigningName = aws.String("ivs")
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

func addUpdateChannelResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateChannelResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
