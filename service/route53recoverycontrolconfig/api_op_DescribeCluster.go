// Code generated by smithy-go-codegen DO NOT EDIT.


package route53recoverycontrolconfig

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/jmespath/go-jmespath"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithytime "github.com/aws/smithy-go/time"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Display the details about a cluster. The response includes the cluster name,
// endpoints, status, and Amazon Resource Name (ARN).
func (c *Client) DescribeCluster(ctx context.Context, params *DescribeClusterInput, optFns ...func(*Options)) (*DescribeClusterOutput, error) {
	if params == nil { params = &DescribeClusterInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeCluster", params, optFns, c.addOperationDescribeClusterMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterInput struct {
	
	// The Amazon Resource Name (ARN) of the cluster.
	//
	// This member is required.
	ClusterArn *string
	
	noSmithyDocumentSerde
}

type DescribeClusterOutput struct {
	
	// The cluster for the DescribeCluster request.
	Cluster *types.Cluster
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCluster{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCluster{}, middleware.After)
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
	if err = addDescribeClusterResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeClusterValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCluster(options.Region, ), middleware.Before); err != nil {
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

// DescribeClusterAPIClient is a client that implements the DescribeCluster
// operation.
type DescribeClusterAPIClient interface {
	DescribeCluster(context.Context, *DescribeClusterInput, ...func(*Options)) (*DescribeClusterOutput, error)
}

var _ DescribeClusterAPIClient = (*Client)(nil)

// ClusterCreatedWaiterOptions are waiter options for ClusterCreatedWaiter
type ClusterCreatedWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ClusterCreatedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ClusterCreatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration
	
	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool
	
	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeClusterInput, *DescribeClusterOutput, error) (bool, error)
}

// ClusterCreatedWaiter defines the waiters for ClusterCreated
type ClusterCreatedWaiter struct {
	
	client DescribeClusterAPIClient
	
	options ClusterCreatedWaiterOptions
}

// NewClusterCreatedWaiter constructs a ClusterCreatedWaiter.
func NewClusterCreatedWaiter(client DescribeClusterAPIClient, optFns ...func(*ClusterCreatedWaiterOptions)) *ClusterCreatedWaiter {
	options := ClusterCreatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterCreatedStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterCreatedWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterCreated waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterCreatedWaiter) Wait(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterCreated waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ClusterCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterCreatedWaiterOptions)) (*DescribeClusterOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}
	
	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}
	
	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}
	
	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}
	
	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()
	
	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur
	
	var attempt int64
	for {
		
		attempt++
		apiOptions := options.APIOptions
		start := time.Now()
		
		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}
		
		out, err := w.client.DescribeCluster(ctx, params, func (o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})
		
		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil { return nil, err }
		if !retryable { return out, nil }
		
		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}
		
		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil { return nil, fmt.Errorf("error computing waiter delay, %w", err)}
		
		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ClusterCreated waiter")
}

func clusterCreatedStateRetryable(ctx context.Context, input *DescribeClusterInput, output *DescribeClusterOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "DEPLOYED"
		value, ok := pathValue.(types.Status)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "PENDING"
		value, ok := pathValue.(types.Status)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}
	
	return true, nil
}

// ClusterDeletedWaiterOptions are waiter options for ClusterDeletedWaiter
type ClusterDeletedWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ClusterDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ClusterDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration
	
	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool
	
	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeClusterInput, *DescribeClusterOutput, error) (bool, error)
}

// ClusterDeletedWaiter defines the waiters for ClusterDeleted
type ClusterDeletedWaiter struct {
	
	client DescribeClusterAPIClient
	
	options ClusterDeletedWaiterOptions
}

// NewClusterDeletedWaiter constructs a ClusterDeletedWaiter.
func NewClusterDeletedWaiter(client DescribeClusterAPIClient, optFns ...func(*ClusterDeletedWaiterOptions)) *ClusterDeletedWaiter {
	options := ClusterDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterDeletedStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterDeletedWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterDeletedWaiter) Wait(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ClusterDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *ClusterDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeClusterInput, maxWaitDur time.Duration, optFns ...func(*ClusterDeletedWaiterOptions)) (*DescribeClusterOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}
	
	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}
	
	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}
	
	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}
	
	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()
	
	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur
	
	var attempt int64
	for {
		
		attempt++
		apiOptions := options.APIOptions
		start := time.Now()
		
		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}
		
		out, err := w.client.DescribeCluster(ctx, params, func (o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})
		
		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil { return nil, err }
		if !retryable { return out, nil }
		
		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}
		
		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil { return nil, fmt.Errorf("error computing waiter delay, %w", err)}
		
		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for ClusterDeleted waiter")
}

func clusterDeletedStateRetryable(ctx context.Context, input *DescribeClusterInput, output *DescribeClusterOutput, err error) (bool, error) {
	
	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("Cluster.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "PENDING_DELETION"
		value, ok := pathValue.(types.Status)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.Status value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	if err != nil {
		var errorType *types.InternalServerException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "route53-recovery-control-config",
	OperationName: "DescribeCluster",
	}
}

type opDescribeClusterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeClusterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeClusterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "route53-recovery-control-config"
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
	        signingName = "route53-recovery-control-config"
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
	        v4aScheme.SigningName = aws.String("route53-recovery-control-config")
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

func addDescribeClusterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeClusterResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
