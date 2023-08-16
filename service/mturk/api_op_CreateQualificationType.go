// Code generated by smithy-go-codegen DO NOT EDIT.


package mturk

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
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// The CreateQualificationType operation creates a new Qualification type, which
// is represented by a QualificationType data structure.
func (c *Client) CreateQualificationType(ctx context.Context, params *CreateQualificationTypeInput, optFns ...func(*Options)) (*CreateQualificationTypeOutput, error) {
	if params == nil { params = &CreateQualificationTypeInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateQualificationType", params, optFns, c.addOperationCreateQualificationTypeMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQualificationTypeInput struct {
	
	// A long description for the Qualification type. On the Amazon Mechanical Turk
	// website, the long description is displayed when a Worker examines a
	// Qualification type.
	//
	// This member is required.
	Description *string
	
	// The name you give to the Qualification type. The type name is used to represent
	// the Qualification to Workers, and to find the type using a Qualification type
	// search. It must be unique across all of your Qualification types.
	//
	// This member is required.
	Name *string
	
	// The initial status of the Qualification type. Constraints: Valid values are:
	// Active | Inactive
	//
	// This member is required.
	QualificationTypeStatus types.QualificationTypeStatus
	
	// The answers to the Qualification test specified in the Test parameter, in the
	// form of an AnswerKey data structure. Constraints: Must not be longer than 65535
	// bytes. Constraints: None. If not specified, you must process Qualification
	// requests manually.
	AnswerKey *string
	
	// Specifies whether requests for the Qualification type are granted immediately,
	// without prompting the Worker with a Qualification test. Constraints: If the Test
	// parameter is specified, this parameter cannot be true.
	AutoGranted *bool
	
	// The Qualification value to use for automatically granted Qualifications. This
	// parameter is used only if the AutoGranted parameter is true.
	AutoGrantedValue *int32
	
	// One or more words or phrases that describe the Qualification type, separated by
	// commas. The keywords of a type make the type easier to find during a search.
	Keywords *string
	
	// The number of seconds that a Worker must wait after requesting a Qualification
	// of the Qualification type before the worker can retry the Qualification request.
	// Constraints: None. If not specified, retries are disabled and Workers can
	// request a Qualification of this type only once, even if the Worker has not been
	// granted the Qualification. It is not possible to disable retries for a
	// Qualification type after it has been created with retries enabled. If you want
	// to disable retries, you must delete existing retry-enabled Qualification type
	// and then create a new Qualification type with retries disabled.
	RetryDelayInSeconds *int64
	
	// The questions for the Qualification test a Worker must answer correctly to
	// obtain a Qualification of this type. If this parameter is specified,
	// TestDurationInSeconds must also be specified. Constraints: Must not be longer
	// than 65535 bytes. Must be a QuestionForm data structure. This parameter cannot
	// be specified if AutoGranted is true. Constraints: None. If not specified, the
	// Worker may request the Qualification without answering any questions.
	Test *string
	
	// The number of seconds the Worker has to complete the Qualification test,
	// starting from the time the Worker requests the Qualification.
	TestDurationInSeconds *int64
	
	noSmithyDocumentSerde
}

type CreateQualificationTypeOutput struct {
	
	// The created Qualification type, returned as a QualificationType data structure.
	QualificationType *types.QualificationType
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateQualificationTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateQualificationType{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateQualificationType{}, middleware.After)
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
	if err = addCreateQualificationTypeResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateQualificationTypeValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQualificationType(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateQualificationType(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "mturk-requester",
	OperationName: "CreateQualificationType",
	}
}

type opCreateQualificationTypeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateQualificationTypeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateQualificationTypeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "mturk-requester"
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
	        signingName = "mturk-requester"
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
	        v4aScheme.SigningName = aws.String("mturk-requester")
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

func addCreateQualificationTypeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateQualificationTypeResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
