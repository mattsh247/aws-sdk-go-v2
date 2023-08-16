// Code generated by smithy-go-codegen DO NOT EDIT.


package rekognition

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
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Detects Personal Protective Equipment (PPE) worn by people detected in an
// image. Amazon Rekognition can detect the following types of PPE.
//   - Face cover
//   - Hand cover
//   - Head cover
// You pass the input image as base64-encoded image bytes or as a reference to an
// image in an Amazon S3 bucket. The image must be either a PNG or JPG formatted
// file. DetectProtectiveEquipment detects PPE worn by up to 15 persons detected
// in an image. For each person detected in the image the API returns an array of
// body parts (face, head, left-hand, right-hand). For each body part, an array of
// detected items of PPE is returned, including an indicator of whether or not the
// PPE covers the body part. The API returns the confidence it has in each
// detection (person, PPE, body part and body part coverage). It also returns a
// bounding box ( BoundingBox ) for each detected person and each detected item of
// PPE. You can optionally request a summary of detected PPE items with the
// SummarizationAttributes input parameter. The summary provides the following
// information.
//   - The persons detected as wearing all of the types of PPE that you specify.
//   - The persons detected as not wearing all of the types PPE that you specify.
//   - The persons detected where PPE adornment could not be determined.
// This is a stateless API operation. That is, the operation does not persist any
// data. This operation requires permissions to perform the
// rekognition:DetectProtectiveEquipment action.
func (c *Client) DetectProtectiveEquipment(ctx context.Context, params *DetectProtectiveEquipmentInput, optFns ...func(*Options)) (*DetectProtectiveEquipmentOutput, error) {
	if params == nil { params = &DetectProtectiveEquipmentInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DetectProtectiveEquipment", params, optFns, c.addOperationDetectProtectiveEquipmentMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DetectProtectiveEquipmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectProtectiveEquipmentInput struct {
	
	// The image in which you want to detect PPE on detected persons. The image can be
	// passed as image bytes or you can reference an image stored in an Amazon S3
	// bucket.
	//
	// This member is required.
	Image *types.Image
	
	// An array of PPE types that you want to summarize.
	SummarizationAttributes *types.ProtectiveEquipmentSummarizationAttributes
	
	noSmithyDocumentSerde
}

type DetectProtectiveEquipmentOutput struct {
	
	// An array of persons detected in the image (including persons not wearing PPE).
	Persons []types.ProtectiveEquipmentPerson
	
	// The version number of the PPE detection model used to detect PPE in the image.
	ProtectiveEquipmentModelVersion *string
	
	// Summary information for the types of PPE specified in the
	// SummarizationAttributes input parameter.
	Summary *types.ProtectiveEquipmentSummary
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectProtectiveEquipmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectProtectiveEquipment{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectProtectiveEquipment{}, middleware.After)
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
	if err = addDetectProtectiveEquipmentResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDetectProtectiveEquipmentValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectProtectiveEquipment(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectProtectiveEquipment(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rekognition",
	OperationName: "DetectProtectiveEquipment",
	}
}

type opDetectProtectiveEquipmentResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDetectProtectiveEquipmentResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDetectProtectiveEquipmentResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDetectProtectiveEquipmentResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDetectProtectiveEquipmentResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
