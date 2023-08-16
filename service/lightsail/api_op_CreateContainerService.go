// Code generated by smithy-go-codegen DO NOT EDIT.


package lightsail

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
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates an Amazon Lightsail container service. A Lightsail container service is
// a compute resource to which you can deploy containers. For more information, see
// Container services in Amazon Lightsail (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-container-services)
// in the Lightsail Dev Guide.
func (c *Client) CreateContainerService(ctx context.Context, params *CreateContainerServiceInput, optFns ...func(*Options)) (*CreateContainerServiceOutput, error) {
	if params == nil { params = &CreateContainerServiceInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateContainerService", params, optFns, c.addOperationCreateContainerServiceMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateContainerServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerServiceInput struct {
	
	// The power specification for the container service. The power specifies the
	// amount of memory, vCPUs, and base monthly cost of each node of the container
	// service. The power and scale of a container service makes up its configured
	// capacity. To determine the monthly price of your container service, multiply the
	// base price of the power with the scale (the number of nodes) of the service.
	// Use the GetContainerServicePowers action to get a list of power options that
	// you can specify using this parameter, and their base monthly cost.
	//
	// This member is required.
	Power types.ContainerServicePowerName
	
	// The scale specification for the container service. The scale specifies the
	// allocated compute nodes of the container service. The power and scale of a
	// container service makes up its configured capacity. To determine the monthly
	// price of your container service, multiply the base price of the power with the
	// scale (the number of nodes) of the service.
	//
	// This member is required.
	Scale *int32
	
	// The name for the container service. The name that you specify for your
	// container service will make up part of its default domain. The default domain of
	// a container service is typically https://...cs.amazonlightsail.com . If the name
	// of your container service is container-service-1 , and it's located in the US
	// East (Ohio) Amazon Web Services Region ( us-east-2 ), then the domain for your
	// container service will be like the following example:
	// https://container-service-1.ur4EXAMPLE2uq.us-east-2.cs.amazonlightsail.com The
	// following are the requirements for container service names:
	//   - Must be unique within each Amazon Web Services Region in your Lightsail
	//   account.
	//   - Must contain 1 to 63 characters.
	//   - Must contain only alphanumeric characters and hyphens.
	//   - A hyphen (-) can separate words but cannot be at the start or end of the
	//   name.
	//
	// This member is required.
	ServiceName *string
	
	// An object that describes a deployment for the container service. A deployment
	// specifies the containers that will be launched on the container service and
	// their settings, such as the ports to open, the environment variables to apply,
	// and the launch command to run. It also specifies the container that will serve
	// as the public endpoint of the deployment and its settings, such as the HTTP or
	// HTTPS port to use, and the health check configuration.
	Deployment *types.ContainerServiceDeploymentRequest
	
	// An object to describe the configuration for the container service to access
	// private container image repositories, such as Amazon Elastic Container Registry
	// (Amazon ECR) private repositories. For more information, see Configuring access
	// to an Amazon ECR private repository for an Amazon Lightsail container service (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-container-service-ecr-private-repo-access)
	// in the Amazon Lightsail Developer Guide.
	PrivateRegistryAccess *types.PrivateRegistryAccessRequest
	
	// The public domain names to use with the container service, such as example.com
	// and www.example.com . You can specify up to four public domain names for a
	// container service. The domain names that you specify are used when you create a
	// deployment with a container configured as the public endpoint of your container
	// service. If you don't specify public domain names, then you can use the default
	// domain of the container service. You must create and validate an SSL/TLS
	// certificate before you can use public domain names with your container service.
	// Use the CreateCertificate action to create a certificate for the public domain
	// names you want to use with your container service. You can specify public domain
	// names using a string to array map as shown in the example later on this page.
	PublicDomainNames map[string][]string
	
	// The tag keys and optional values to add to the container service during create.
	// Use the TagResource action to tag a resource after it's created. For more
	// information about tags in Lightsail, see the Amazon Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-tags)
	// .
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateContainerServiceOutput struct {
	
	// An object that describes a container service.
	ContainerService *types.ContainerService
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainerService{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainerService{}, middleware.After)
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
	if err = addCreateContainerServiceResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateContainerServiceValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerService(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContainerService(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "lightsail",
	OperationName: "CreateContainerService",
	}
}

type opCreateContainerServiceResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateContainerServiceResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateContainerServiceResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addCreateContainerServiceResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateContainerServiceResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
