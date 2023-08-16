// Code generated by smithy-go-codegen DO NOT EDIT.


package macie2

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
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates and defines the criteria and other settings for a custom data
// identifier.
func (c *Client) CreateCustomDataIdentifier(ctx context.Context, params *CreateCustomDataIdentifierInput, optFns ...func(*Options)) (*CreateCustomDataIdentifierOutput, error) {
	if params == nil { params = &CreateCustomDataIdentifierInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateCustomDataIdentifier", params, optFns, c.addOperationCreateCustomDataIdentifierMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomDataIdentifierInput struct {
	
	// A custom name for the custom data identifier. The name can contain as many as
	// 128 characters. We strongly recommend that you avoid including any sensitive
	// data in the name of a custom data identifier. Other users of your account might
	// be able to see this name, depending on the actions that they're allowed to
	// perform in Amazon Macie.
	//
	// This member is required.
	Name *string
	
	// The regular expression (regex) that defines the pattern to match. The
	// expression can contain as many as 512 characters.
	//
	// This member is required.
	Regex *string
	
	// A unique, case-sensitive token that you provide to ensure the idempotency of
	// the request.
	ClientToken *string
	
	// A custom description of the custom data identifier. The description can contain
	// as many as 512 characters. We strongly recommend that you avoid including any
	// sensitive data in the description of a custom data identifier. Other users of
	// your account might be able to see this description, depending on the actions
	// that they're allowed to perform in Amazon Macie.
	Description *string
	
	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression contains any string
	// in this array, Amazon Macie ignores it. The array can contain as many as 10
	// ignore words. Each ignore word can contain 4-90 UTF-8 characters. Ignore words
	// are case sensitive.
	IgnoreWords []string
	
	// An array that lists specific character sequences (keywords), one of which must
	// precede and be within proximity (maximumMatchDistance) of the regular expression
	// to match. The array can contain as many as 50 keywords. Each keyword can contain
	// 3-90 UTF-8 characters. Keywords aren't case sensitive.
	Keywords []string
	
	// The maximum number of characters that can exist between the end of at least one
	// complete character sequence specified by the keywords array and the end of the
	// text that matches the regex pattern. If a complete keyword precedes all the text
	// that matches the pattern and the keyword is within the specified distance,
	// Amazon Macie includes the result. The distance can be 1-300 characters. The
	// default value is 50.
	MaximumMatchDistance int32
	
	// The severity to assign to findings that the custom data identifier produces,
	// based on the number of occurrences of text that match the custom data
	// identifier's detection criteria. You can specify as many as three SeverityLevel
	// objects in this array, one for each severity: LOW, MEDIUM, or HIGH. If you
	// specify more than one, the occurrences thresholds must be in ascending order by
	// severity, moving from LOW to HIGH. For example, 1 for LOW, 50 for MEDIUM, and
	// 100 for HIGH. If an S3 object contains fewer occurrences than the lowest
	// specified threshold, Amazon Macie doesn't create a finding. If you don't specify
	// any values for this array, Macie creates findings for S3 objects that contain at
	// least one occurrence of text that matches the detection criteria, and Macie
	// assigns the MEDIUM severity to those findings.
	SeverityLevels []types.SeverityLevel
	
	// A map of key-value pairs that specifies the tags to associate with the custom
	// data identifier. A custom data identifier can have a maximum of 50 tags. Each
	// tag consists of a tag key and an associated tag value. The maximum length of a
	// tag key is 128 characters. The maximum length of a tag value is 256 characters.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

type CreateCustomDataIdentifierOutput struct {
	
	// The unique identifier for the custom data identifier that was created.
	CustomDataIdentifierId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomDataIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomDataIdentifier{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomDataIdentifier{}, middleware.After)
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
	if err = addCreateCustomDataIdentifierResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opCreateCustomDataIdentifierMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateCustomDataIdentifierValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomDataIdentifier(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomDataIdentifier struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomDataIdentifier) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomDataIdentifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*CreateCustomDataIdentifierInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomDataIdentifierInput ")}
	
	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateCustomDataIdentifierMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomDataIdentifier{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomDataIdentifier(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "macie2",
	OperationName: "CreateCustomDataIdentifier",
	}
}

type opCreateCustomDataIdentifierResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateCustomDataIdentifierResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateCustomDataIdentifierResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "macie2"
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
	        signingName = "macie2"
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
	        v4aScheme.SigningName = aws.String("macie2")
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

func addCreateCustomDataIdentifierResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateCustomDataIdentifierResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
