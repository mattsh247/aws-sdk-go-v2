// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Classic Load Balancer. You can add listeners, security groups,
// subnets, and tags when you create your load balancer, or you can add them later
// using CreateLoadBalancerListeners , ApplySecurityGroupsToLoadBalancer ,
// AttachLoadBalancerToSubnets , and AddTags . To describe your current load
// balancers, see DescribeLoadBalancers . When you are finished with a load
// balancer, you can delete it using DeleteLoadBalancer . You can create up to 20
// load balancers per region per account. You can request an increase for the
// number of load balancers for your account. For more information, see Limits for
// Your Classic Load Balancer (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-limits.html)
// in the Classic Load Balancers Guide.
func (c *Client) CreateLoadBalancer(ctx context.Context, params *CreateLoadBalancerInput, optFns ...func(*Options)) (*CreateLoadBalancerOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancer", params, optFns, c.addOperationCreateLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancer.
type CreateLoadBalancerInput struct {

	// The listeners. For more information, see Listeners for Your Classic Load
	// Balancer (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
	// in the Classic Load Balancers Guide.
	//
	// This member is required.
	Listeners []types.Listener

	// The name of the load balancer. This name must be unique within your set of load
	// balancers for the region, must have a maximum of 32 characters, must contain
	// only alphanumeric characters or hyphens, and cannot begin or end with a hyphen.
	//
	// This member is required.
	LoadBalancerName *string

	// One or more Availability Zones from the same region as the load balancer. You
	// must specify at least one Availability Zone. You can add more Availability Zones
	// after you create the load balancer using EnableAvailabilityZonesForLoadBalancer .
	AvailabilityZones []string

	// The type of a load balancer. Valid only for load balancers in a VPC. By
	// default, Elastic Load Balancing creates an Internet-facing load balancer with a
	// DNS name that resolves to public IP addresses. For more information about
	// Internet-facing and Internal load balancers, see Load Balancer Scheme (https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html#load-balancer-scheme)
	// in the Elastic Load Balancing User Guide. Specify internal to create a load
	// balancer with a DNS name that resolves to private IP addresses.
	Scheme *string

	// The IDs of the security groups to assign to the load balancer.
	SecurityGroups []string

	// The IDs of the subnets in your VPC to attach to the load balancer. Specify one
	// subnet per Availability Zone specified in AvailabilityZones .
	Subnets []string

	// A list of tags to assign to the load balancer. For more information about
	// tagging your load balancer, see Tag Your Classic Load Balancer (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/add-remove-tags.html)
	// in the Classic Load Balancers Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the output for CreateLoadBalancer.
type CreateLoadBalancerOutput struct {

	// The DNS name of the load balancer.
	DNSName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancer{}, middleware.After)
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
	if err = addCreateLoadBalancerResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancer",
	}
}

type opCreateLoadBalancerResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateLoadBalancerResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateLoadBalancerResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "elasticloadbalancing"
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
				signingName = "elasticloadbalancing"
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
				v4aScheme.SigningName = aws.String("elasticloadbalancing")
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

func addCreateLoadBalancerResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateLoadBalancerResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
