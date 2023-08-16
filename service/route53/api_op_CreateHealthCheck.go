// Code generated by smithy-go-codegen DO NOT EDIT.


package route53

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
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a new health check. For information about adding health checks to
// resource record sets, see HealthCheckId (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId)
// in ChangeResourceRecordSets (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html)
// . ELB Load Balancers If you're registering EC2 instances with an Elastic Load
// Balancing (ELB) load balancer, do not create Amazon Route 53 health checks for
// the EC2 instances. When you register an EC2 instance with a load balancer, you
// configure settings for an ELB health check, which performs a similar function to
// a Route 53 health check. Private Hosted Zones You can associate health checks
// with failover resource record sets in a private hosted zone. Note the following:
//
//   - Route 53 health checkers are outside the VPC. To check the health of an
//   endpoint within a VPC by IP address, you must assign a public IP address to the
//   instance in the VPC.
//   - You can configure a health checker to check the health of an external
//   resource that the instance relies on, such as a database server.
//   - You can create a CloudWatch metric, associate an alarm with the metric, and
//   then create a health check that is based on the state of the alarm. For example,
//   you might create a CloudWatch metric that checks the status of the Amazon EC2
//   StatusCheckFailed metric, add an alarm to the metric, and then create a health
//   check that is based on the state of the alarm. For information about creating
//   CloudWatch metrics and alarms by using the CloudWatch console, see the Amazon
//   CloudWatch User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html)
//   .
func (c *Client) CreateHealthCheck(ctx context.Context, params *CreateHealthCheckInput, optFns ...func(*Options)) (*CreateHealthCheckOutput, error) {
	if params == nil { params = &CreateHealthCheckInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateHealthCheck", params, optFns, c.addOperationCreateHealthCheckMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains the health check request information.
type CreateHealthCheckInput struct {
	
	// A unique string that identifies the request and that allows you to retry a
	// failed CreateHealthCheck request without the risk of creating two identical
	// health checks:
	//   - If you send a CreateHealthCheck request with the same CallerReference and
	//   settings as a previous request, and if the health check doesn't exist, Amazon
	//   Route 53 creates the health check. If the health check does exist, Route 53
	//   returns the settings for the existing health check.
	//   - If you send a CreateHealthCheck request with the same CallerReference as a
	//   deleted health check, regardless of the settings, Route 53 returns a
	//   HealthCheckAlreadyExists error.
	//   - If you send a CreateHealthCheck request with the same CallerReference as an
	//   existing health check but with different settings, Route 53 returns a
	//   HealthCheckAlreadyExists error.
	//   - If you send a CreateHealthCheck request with a unique CallerReference but
	//   settings identical to an existing health check, Route 53 creates the health
	//   check.
	//
	// This member is required.
	CallerReference *string
	
	// A complex type that contains settings for a new health check.
	//
	// This member is required.
	HealthCheckConfig *types.HealthCheckConfig
	
	noSmithyDocumentSerde
}

// A complex type containing the response information for the new health check.
type CreateHealthCheckOutput struct {
	
	// A complex type that contains identifying information about the health check.
	//
	// This member is required.
	HealthCheck *types.HealthCheck
	
	// The unique URL representing the new health check.
	//
	// This member is required.
	Location *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateHealthCheck{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateHealthCheck{}, middleware.After)
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
	if err = addCreateHealthCheckResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateHealthCheckValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHealthCheck(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "route53",
	OperationName: "CreateHealthCheck",
	}
}

type opCreateHealthCheckResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateHealthCheckResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateHealthCheckResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "route53"
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
	        signingName = "route53"
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
	        v4aScheme.SigningName = aws.String("route53")
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

func addCreateHealthCheckResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateHealthCheckResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
