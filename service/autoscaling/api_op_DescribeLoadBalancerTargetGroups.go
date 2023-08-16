// Code generated by smithy-go-codegen DO NOT EDIT.


package autoscaling

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
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// This API operation is superseded by DescribeTrafficSources , which can describe
// multiple traffic sources types. We recommend using DetachTrafficSources to
// simplify how you manage traffic sources. However, we continue to support
// DescribeLoadBalancerTargetGroups . You can use both the original
// DescribeLoadBalancerTargetGroups API operation and DescribeTrafficSources on
// the same Auto Scaling group. Gets information about the Elastic Load Balancing
// target groups for the specified Auto Scaling group. To determine the attachment
// status of the target group, use the State element in the response. When you
// attach a target group to an Auto Scaling group, the initial State value is
// Adding . The state transitions to Added after all Auto Scaling instances are
// registered with the target group. If Elastic Load Balancing health checks are
// enabled for the Auto Scaling group, the state transitions to InService after at
// least one Auto Scaling instance passes the health check. When the target group
// is in the InService state, Amazon EC2 Auto Scaling can terminate and replace
// any instances that are reported as unhealthy. If no registered instances pass
// the health checks, the target group doesn't enter the InService state. Target
// groups also have an InService state if you attach them in the
// CreateAutoScalingGroup API call. If your target group state is InService , but
// it is not working properly, check the scaling activities by calling
// DescribeScalingActivities and take any corrective actions necessary. For help
// with failed health checks, see Troubleshooting Amazon EC2 Auto Scaling: Health
// checks (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Use
// Elastic Load Balancing to distribute traffic across the instances in your Auto
// Scaling group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
// in the Amazon EC2 Auto Scaling User Guide. You can use this operation to
// describe target groups that were attached by using
// AttachLoadBalancerTargetGroups , but not for target groups that were attached by
// using AttachTrafficSources .
func (c *Client) DescribeLoadBalancerTargetGroups(ctx context.Context, params *DescribeLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*DescribeLoadBalancerTargetGroupsOutput, error) {
	if params == nil { params = &DescribeLoadBalancerTargetGroupsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancerTargetGroups", params, optFns, c.addOperationDescribeLoadBalancerTargetGroupsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancerTargetGroupsInput struct {
	
	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string
	
	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100 .
	MaxRecords *int32
	
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	
	noSmithyDocumentSerde
}

type DescribeLoadBalancerTargetGroupsOutput struct {
	
	// Information about the target groups.
	LoadBalancerTargetGroups []types.LoadBalancerTargetGroupState
	
	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
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
	if err = addDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeLoadBalancerTargetGroupsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "autoscaling",
	OperationName: "DescribeLoadBalancerTargetGroups",
	}
}

type opDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "autoscaling"
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
	        signingName = "autoscaling"
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
	        v4aScheme.SigningName = aws.String("autoscaling")
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

func addDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeLoadBalancerTargetGroupsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
