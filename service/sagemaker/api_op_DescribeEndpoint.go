// Code generated by smithy-go-codegen DO NOT EDIT.


package sagemaker

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	"github.com/jmespath/go-jmespath"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithytime "github.com/aws/smithy-go/time"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns the description of an endpoint.
func (c *Client) DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) {
	if params == nil { params = &DescribeEndpointInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoint", params, optFns, c.addOperationDescribeEndpointMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointInput struct {
	
	// The name of the endpoint.
	//
	// This member is required.
	EndpointName *string
	
	noSmithyDocumentSerde
}

type DescribeEndpointOutput struct {
	
	// A timestamp that shows when the endpoint was created.
	//
	// This member is required.
	CreationTime *time.Time
	
	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string
	
	// The name of the endpoint configuration associated with this endpoint.
	//
	// This member is required.
	EndpointConfigName *string
	
	// Name of the endpoint.
	//
	// This member is required.
	EndpointName *string
	
	// The status of the endpoint.
	//   - OutOfService : Endpoint is not available to take incoming requests.
	//   - Creating : CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
	//   is executing.
	//   - Updating : UpdateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpoint.html)
	//   or UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//   is executing.
	//   - SystemUpdating : Endpoint is undergoing maintenance and cannot be updated or
	//   deleted or re-scaled until it has completed. This maintenance operation does not
	//   change any customer-specified values such as VPC config, KMS encryption, model,
	//   instance type, or instance count.
	//   - RollingBack : Endpoint fails to scale up or down or change its variant
	//   weight and is in the process of rolling back to its previous configuration. Once
	//   the rollback completes, endpoint returns to an InService status. This
	//   transitional status only applies to an endpoint that has autoscaling enabled and
	//   is undergoing variant weight or capacity changes as part of an
	//   UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//   call or when the UpdateEndpointWeightsAndCapacities (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_UpdateEndpointWeightsAndCapacities.html)
	//   operation is called explicitly.
	//   - InService : Endpoint is available to process incoming requests.
	//   - Deleting : DeleteEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteEndpoint.html)
	//   is executing.
	//   - Failed : Endpoint could not be created, updated, or re-scaled. Use the
	//   FailureReason value returned by DescribeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html)
	//   for information about the failure. DeleteEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DeleteEndpoint.html)
	//   is the only operation that can be performed on a failed endpoint.
	//   - UpdateRollbackFailed : Both the rolling deployment and auto-rollback failed.
	//   Your endpoint is in service with a mix of the old and new endpoint
	//   configurations. For information about how to remedy this issue and restore the
	//   endpoint's status to InService , see Rolling Deployments (https://docs.aws.amazon.com/sagemaker/latest/dg/deployment-guardrails-rolling.html)
	//   .
	//
	// This member is required.
	EndpointStatus types.EndpointStatus
	
	// A timestamp that shows when the endpoint was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time
	
	// Returns the description of an endpoint configuration created using the
	// CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html)
	// API.
	AsyncInferenceConfig *types.AsyncInferenceConfig
	
	// The currently active data capture configuration used by your Endpoint.
	DataCaptureConfig *types.DataCaptureConfigSummary
	
	// The configuration parameters for an explainer.
	ExplainerConfig *types.ExplainerConfig
	
	// If the status of the endpoint is Failed , the reason why it failed.
	FailureReason *string
	
	// The most recent deployment configuration for the endpoint.
	LastDeploymentConfig *types.DeploymentConfig
	
	// Returns the summary of an in-progress deployment. This field is only returned
	// when the endpoint is creating or updating with a new endpoint configuration.
	PendingDeploymentSummary *types.PendingDeploymentSummary
	
	// An array of ProductionVariantSummary (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariantSummary.html)
	// objects, one for each model hosted behind this endpoint.
	ProductionVariants []types.ProductionVariantSummary
	
	// An array of ProductionVariantSummary (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariantSummary.html)
	// objects, one for each model that you want to host at this endpoint in shadow
	// mode with production traffic replicated from the model specified on
	// ProductionVariants .
	ShadowProductionVariants []types.ProductionVariantSummary
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpoint{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpoint{}, middleware.After)
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
	if err = addDescribeEndpointResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeEndpointValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoint(options.Region, ), middleware.Before); err != nil {
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

// DescribeEndpointAPIClient is a client that implements the DescribeEndpoint
// operation.
type DescribeEndpointAPIClient interface {
	DescribeEndpoint(context.Context, *DescribeEndpointInput, ...func(*Options)) (*DescribeEndpointOutput, error)
}

var _ DescribeEndpointAPIClient = (*Client)(nil)

// EndpointDeletedWaiterOptions are waiter options for EndpointDeletedWaiter
type EndpointDeletedWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// EndpointDeletedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EndpointDeletedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeEndpointInput, *DescribeEndpointOutput, error) (bool, error)
}

// EndpointDeletedWaiter defines the waiters for EndpointDeleted
type EndpointDeletedWaiter struct {
	
	client DescribeEndpointAPIClient
	
	options EndpointDeletedWaiterOptions
}

// NewEndpointDeletedWaiter constructs a EndpointDeletedWaiter.
func NewEndpointDeletedWaiter(client DescribeEndpointAPIClient, optFns ...func(*EndpointDeletedWaiterOptions)) *EndpointDeletedWaiter {
	options := EndpointDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = endpointDeletedStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &EndpointDeletedWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for EndpointDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *EndpointDeletedWaiter) Wait(ctx context.Context, params *DescribeEndpointInput, maxWaitDur time.Duration, optFns ...func(*EndpointDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EndpointDeleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *EndpointDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeEndpointInput, maxWaitDur time.Duration, optFns ...func(*EndpointDeletedWaiterOptions)) (*DescribeEndpointOutput, error) {
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
		
		out, err := w.client.DescribeEndpoint(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EndpointDeleted waiter")
}

func endpointDeletedStateRetryable(ctx context.Context, input *DescribeEndpointInput, output *DescribeEndpointOutput, err error) (bool, error) {
	
	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}
		
		if "ValidationException" == apiErr.ErrorCode() {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("EndpointStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Failed"
		value, ok := pathValue.(types.EndpointStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.EndpointStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	return true, nil
}

// EndpointInServiceWaiterOptions are waiter options for EndpointInServiceWaiter
type EndpointInServiceWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// EndpointInServiceWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, EndpointInServiceWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeEndpointInput, *DescribeEndpointOutput, error) (bool, error)
}

// EndpointInServiceWaiter defines the waiters for EndpointInService
type EndpointInServiceWaiter struct {
	
	client DescribeEndpointAPIClient
	
	options EndpointInServiceWaiterOptions
}

// NewEndpointInServiceWaiter constructs a EndpointInServiceWaiter.
func NewEndpointInServiceWaiter(client DescribeEndpointAPIClient, optFns ...func(*EndpointInServiceWaiterOptions)) *EndpointInServiceWaiter {
	options := EndpointInServiceWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = endpointInServiceStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &EndpointInServiceWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for EndpointInService waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *EndpointInServiceWaiter) Wait(ctx context.Context, params *DescribeEndpointInput, maxWaitDur time.Duration, optFns ...func(*EndpointInServiceWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for EndpointInService waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *EndpointInServiceWaiter) WaitForOutput(ctx context.Context, params *DescribeEndpointInput, maxWaitDur time.Duration, optFns ...func(*EndpointInServiceWaiterOptions)) (*DescribeEndpointOutput, error) {
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
		
		out, err := w.client.DescribeEndpoint(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for EndpointInService waiter")
}

func endpointInServiceStateRetryable(ctx context.Context, input *DescribeEndpointInput, output *DescribeEndpointOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("EndpointStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "InService"
		value, ok := pathValue.(types.EndpointStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.EndpointStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("EndpointStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Failed"
		value, ok := pathValue.(types.EndpointStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.EndpointStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}
		
		if "ValidationException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opDescribeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "sagemaker",
	OperationName: "DescribeEndpoint",
	}
}

type opDescribeEndpointResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeEndpointResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeEndpointResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "sagemaker"
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
	        signingName = "sagemaker"
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
	        v4aScheme.SigningName = aws.String("sagemaker")
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

func addDescribeEndpointResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeEndpointResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
