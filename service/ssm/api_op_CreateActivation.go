// Code generated by smithy-go-codegen DO NOT EDIT.


package ssm

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
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Generates an activation code and activation ID you can use to register your
// on-premises servers, edge devices, or virtual machine (VM) with Amazon Web
// Services Systems Manager. Registering these machines with Systems Manager makes
// it possible to manage them using Systems Manager capabilities. You use the
// activation code and ID when installing SSM Agent on machines in your hybrid
// environment. For more information about requirements for managing on-premises
// machines using Systems Manager, see Setting up Amazon Web Services Systems
// Manager for hybrid environments (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-managedinstances.html)
// in the Amazon Web Services Systems Manager User Guide. Amazon Elastic Compute
// Cloud (Amazon EC2) instances, edge devices, and on-premises servers and VMs that
// are configured for Systems Manager are all called managed nodes.
func (c *Client) CreateActivation(ctx context.Context, params *CreateActivationInput, optFns ...func(*Options)) (*CreateActivationOutput, error) {
	if params == nil { params = &CreateActivationInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateActivation", params, optFns, c.addOperationCreateActivationMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateActivationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateActivationInput struct {
	
	// The name of the Identity and Access Management (IAM) role that you want to
	// assign to the managed node. This IAM role must provide AssumeRole permissions
	// for the Amazon Web Services Systems Manager service principal ssm.amazonaws.com
	// . For more information, see Create an IAM service role for a hybrid environment (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-service-role.html)
	// in the Amazon Web Services Systems Manager User Guide. You can't specify an IAM
	// service-linked role for this parameter. You must create a unique role.
	//
	// This member is required.
	IamRole *string
	
	// The name of the registered, managed node as it will appear in the Amazon Web
	// Services Systems Manager console or when you use the Amazon Web Services command
	// line tools to list Systems Manager resources. Don't enter personally
	// identifiable information in this field.
	DefaultInstanceName *string
	
	// A user-defined description of the resource that you want to register with
	// Systems Manager. Don't enter personally identifiable information in this field.
	Description *string
	
	// The date by which this activation request should expire, in timestamp format,
	// such as "2021-07-07T00:00:00". You can specify a date up to 30 days in advance.
	// If you don't provide an expiration date, the activation code expires in 24
	// hours.
	ExpirationDate *time.Time
	
	// Specify the maximum number of managed nodes you want to register. The default
	// value is 1 .
	RegistrationLimit *int32
	
	// Reserved for internal use.
	RegistrationMetadata []types.RegistrationMetadataItem
	
	// Optional metadata that you assign to a resource. Tags enable you to categorize
	// a resource in different ways, such as by purpose, owner, or environment. For
	// example, you might want to tag an activation to identify which servers or
	// virtual machines (VMs) in your on-premises environment you intend to activate.
	// In this case, you could specify the following key-value pairs:
	//   - Key=OS,Value=Windows
	//   - Key=Environment,Value=Production
	// When you install SSM Agent on your on-premises servers and VMs, you specify an
	// activation ID and code. When you specify the activation ID and code, tags
	// assigned to the activation are automatically applied to the on-premises servers
	// or VMs. You can't add tags to or delete tags from an existing activation. You
	// can tag your on-premises servers, edge devices, and VMs after they connect to
	// Systems Manager for the first time and are assigned a managed node ID. This
	// means they are listed in the Amazon Web Services Systems Manager console with an
	// ID that is prefixed with "mi-". For information about how to add tags to your
	// managed nodes, see AddTagsToResource . For information about how to remove tags
	// from your managed nodes, see RemoveTagsFromResource .
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateActivationOutput struct {
	
	// The code the system generates when it processes the activation. The activation
	// code functions like a password to validate the activation ID.
	ActivationCode *string
	
	// The ID number generated by the system when it processed the activation. The
	// activation ID functions like a user name.
	ActivationId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateActivationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateActivation{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateActivation{}, middleware.After)
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
	if err = addCreateActivationResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateActivationValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateActivation(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateActivation(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ssm",
	OperationName: "CreateActivation",
	}
}

type opCreateActivationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateActivationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateActivationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ssm"
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
	        signingName = "ssm"
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
	        v4aScheme.SigningName = aws.String("ssm")
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

func addCreateActivationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateActivationResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
