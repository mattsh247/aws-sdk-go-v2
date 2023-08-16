// Code generated by smithy-go-codegen DO NOT EDIT.


package ecs

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
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Modifies the status of an Amazon ECS container instance. Once a container
// instance has reached an ACTIVE state, you can change the status of a container
// instance to DRAINING to manually remove an instance from a cluster, for example
// to perform system updates, update the Docker daemon, or scale down the cluster
// size. A container instance can't be changed to DRAINING until it has reached an
// ACTIVE status. If the instance is in any other status, an error will be
// received. When you set a container instance to DRAINING , Amazon ECS prevents
// new tasks from being scheduled for placement on the container instance and
// replacement service tasks are started on other container instances in the
// cluster if the resources are available. Service tasks on the container instance
// that are in the PENDING state are stopped immediately. Service tasks on the
// container instance that are in the RUNNING state are stopped and replaced
// according to the service's deployment configuration parameters,
// minimumHealthyPercent and maximumPercent . You can change the deployment
// configuration of your service using UpdateService .
//   - If minimumHealthyPercent is below 100%, the scheduler can ignore
//   desiredCount temporarily during task replacement. For example, desiredCount is
//   four tasks, a minimum of 50% allows the scheduler to stop two existing tasks
//   before starting two new tasks. If the minimum is 100%, the service scheduler
//   can't remove existing tasks until the replacement tasks are considered healthy.
//   Tasks for services that do not use a load balancer are considered healthy if
//   they're in the RUNNING state. Tasks for services that use a load balancer are
//   considered healthy if they're in the RUNNING state and are reported as healthy
//   by the load balancer.
//   - The maximumPercent parameter represents an upper limit on the number of
//   running tasks during task replacement. You can use this to define the
//   replacement batch size. For example, if desiredCount is four tasks, a maximum
//   of 200% starts four new tasks before stopping the four tasks to be drained,
//   provided that the cluster resources required to do this are available. If the
//   maximum is 100%, then replacement tasks can't start until the draining tasks
//   have stopped.
// Any PENDING or RUNNING tasks that do not belong to a service aren't affected.
// You must wait for them to finish or stop them manually. A container instance has
// completed draining when it has no more RUNNING tasks. You can verify this using
// ListTasks . When a container instance has been drained, you can set a container
// instance to ACTIVE status and once it has reached that status the Amazon ECS
// scheduler can begin scheduling tasks on the instance again.
func (c *Client) UpdateContainerInstancesState(ctx context.Context, params *UpdateContainerInstancesStateInput, optFns ...func(*Options)) (*UpdateContainerInstancesStateOutput, error) {
	if params == nil { params = &UpdateContainerInstancesStateInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateContainerInstancesState", params, optFns, c.addOperationUpdateContainerInstancesStateMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateContainerInstancesStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContainerInstancesStateInput struct {
	
	// A list of up to 10 container instance IDs or full ARN entries.
	//
	// This member is required.
	ContainerInstances []string
	
	// The container instance state to update the container instance with. The only
	// valid values for this action are ACTIVE and DRAINING . A container instance can
	// only be updated to DRAINING status once it has reached an ACTIVE state. If a
	// container instance is in REGISTERING , DEREGISTERING , or REGISTRATION_FAILED
	// state you can describe the container instance but can't update the container
	// instance state.
	//
	// This member is required.
	Status types.ContainerInstanceStatus
	
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instance to update. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string
	
	noSmithyDocumentSerde
}

type UpdateContainerInstancesStateOutput struct {
	
	// The list of container instances.
	ContainerInstances []types.ContainerInstance
	
	// Any failures associated with the call.
	Failures []types.Failure
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateContainerInstancesStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateContainerInstancesState{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateContainerInstancesState{}, middleware.After)
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
	if err = addUpdateContainerInstancesStateResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateContainerInstancesStateValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContainerInstancesState(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateContainerInstancesState(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ecs",
	OperationName: "UpdateContainerInstancesState",
	}
}

type opUpdateContainerInstancesStateResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateContainerInstancesStateResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateContainerInstancesStateResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ecs"
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
	        signingName = "ecs"
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
	        v4aScheme.SigningName = aws.String("ecs")
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

func addUpdateContainerInstancesStateResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateContainerInstancesStateResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
