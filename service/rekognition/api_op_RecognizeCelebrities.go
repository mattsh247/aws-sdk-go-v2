// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of celebrities recognized in the input image. For more
// information, see Recognizing celebrities in the Amazon Rekognition Developer
// Guide. RecognizeCelebrities returns the 64 largest faces in the image. It lists
// the recognized celebrities in the CelebrityFaces array and any unrecognized
// faces in the UnrecognizedFaces array. RecognizeCelebrities doesn't return
// celebrities whose faces aren't among the largest 64 faces in the image. For each
// celebrity recognized, RecognizeCelebrities returns a Celebrity object. The
// Celebrity object contains the celebrity name, ID, URL links to additional
// information, match confidence, and a ComparedFace object that you can use to
// locate the celebrity's face on the image. Amazon Rekognition doesn't retain
// information about which images a celebrity has been recognized in. Your
// application must store this information and use the Celebrity ID property as a
// unique identifier for the celebrity. If you don't store the celebrity name or
// additional information URLs returned by RecognizeCelebrities , you will need the
// ID to identify the celebrity in a call to the GetCelebrityInfo operation. You
// pass the input image either as base64-encoded image bytes or as a reference to
// an image in an Amazon S3 bucket. If you use the AWS CLI to call Amazon
// Rekognition operations, passing image bytes is not supported. The image must be
// either a PNG or JPEG formatted file. For an example, see Recognizing celebrities
// in an image in the Amazon Rekognition Developer Guide. This operation requires
// permissions to perform the rekognition:RecognizeCelebrities operation.
func (c *Client) RecognizeCelebrities(ctx context.Context, params *RecognizeCelebritiesInput, optFns ...func(*Options)) (*RecognizeCelebritiesOutput, error) {
	if params == nil {
		params = &RecognizeCelebritiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecognizeCelebrities", params, optFns, c.addOperationRecognizeCelebritiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RecognizeCelebritiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecognizeCelebritiesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	Image *types.Image

	noSmithyDocumentSerde
}

type RecognizeCelebritiesOutput struct {

	// Details about each celebrity found in the image. Amazon Rekognition can detect
	// a maximum of 64 celebrities in an image. Each celebrity object includes the
	// following attributes: Face , Confidence , Emotions , Landmarks , Pose , Quality
	// , Smile , Id , KnownGender , MatchConfidence , Name , Urls .
	CelebrityFaces []types.Celebrity

	// Support for estimating image orientation using the the OrientationCorrection
	// field has ceased as of August 2021. Any returned values for this field included
	// in an API response will always be NULL. The orientation of the input image
	// (counterclockwise direction). If your application displays the image, you can
	// use this value to correct the orientation. The bounding box coordinates returned
	// in CelebrityFaces and UnrecognizedFaces represent face locations before the
	// image orientation is corrected. If the input image is in .jpeg format, it might
	// contain exchangeable image (Exif) metadata that includes the image's
	// orientation. If so, and the Exif metadata for the input image populates the
	// orientation field, the value of OrientationCorrection is null. The
	// CelebrityFaces and UnrecognizedFaces bounding box coordinates represent face
	// locations after Exif metadata is used to correct the image orientation. Images
	// in .png format don't contain Exif metadata.
	OrientationCorrection types.OrientationCorrection

	// Details about each unrecognized face in the image.
	UnrecognizedFaces []types.ComparedFace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRecognizeCelebritiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRecognizeCelebrities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRecognizeCelebrities{}, middleware.After)
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
	if err = addRecognizeCelebritiesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpRecognizeCelebritiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRecognizeCelebrities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRecognizeCelebrities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "RecognizeCelebrities",
	}
}

type opRecognizeCelebritiesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opRecognizeCelebritiesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opRecognizeCelebritiesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "rekognition"
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
				signingName = "rekognition"
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
				v4aScheme.SigningName = aws.String("rekognition")
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

func addRecognizeCelebritiesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opRecognizeCelebritiesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
