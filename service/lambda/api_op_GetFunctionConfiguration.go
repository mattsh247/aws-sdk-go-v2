// Code generated by smithy-go-codegen DO NOT EDIT.


package lambda

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
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns the version-specific settings of a Lambda function or version. The
// output includes only options that can vary between versions of a function. To
// modify these settings, use UpdateFunctionConfiguration . To get all of a
// function's details, including function-level settings, use GetFunction .
func (c *Client) GetFunctionConfiguration(ctx context.Context, params *GetFunctionConfigurationInput, optFns ...func(*Options)) (*GetFunctionConfigurationOutput, error) {
	if params == nil { params = &GetFunctionConfigurationInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetFunctionConfiguration", params, optFns, c.addOperationGetFunctionConfigurationMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetFunctionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionConfigurationInput struct {
	
	// The name of the Lambda function, version, or alias. Name formats
	//   - Function name – my-function (name-only), my-function:v1 (with alias).
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function .
	//   - Partial ARN – 123456789012:function:my-function .
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function name,
	// it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string
	
	// Specify a version or alias to get details about a published version of the
	// function.
	Qualifier *string
	
	noSmithyDocumentSerde
}

// Details about a function's configuration.
type GetFunctionConfigurationOutput struct {
	
	// The instruction set architecture that the function supports. Architecture is a
	// string array with one of the valid values. The default architecture value is
	// x86_64 .
	Architectures []types.Architecture
	
	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string
	
	// The size of the function's deployment package, in bytes.
	CodeSize int64
	
	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig
	
	// The function's description.
	Description *string
	
	// The function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html)
	// . Omitted from CloudTrail logs.
	Environment *types.EnvironmentResponse
	
	// The size of the function’s /tmp directory in MB. The default value is 512, but
	// it can be any whole number between 512 and 10,240 MB.
	EphemeralStorage *types.EphemeralStorage
	
	// Connection settings for an Amazon EFS file system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html)
	// .
	FileSystemConfigs []types.FileSystemConfig
	
	// The function's Amazon Resource Name (ARN).
	FunctionArn *string
	
	// The name of the function.
	FunctionName *string
	
	// The function that Lambda calls to begin running your function.
	Handler *string
	
	// The function's image configuration values.
	ImageConfigResponse *types.ImageConfigResponse
	
	// The KMS key that's used to encrypt the function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption)
	// . When Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html)
	// is activated, this key is also used to encrypt the function's snapshot. This key
	// is returned only if you've configured a customer managed key.
	KMSKeyArn *string
	
	// The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string
	
	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus
	
	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string
	
	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode
	
	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	// .
	Layers []types.Layer
	
	// For Lambda@Edge functions, the ARN of the main function.
	MasterArn *string
	
	// The amount of memory available to the function at runtime.
	MemorySize *int32
	
	// The type of deployment package. Set to Image for container image and set Zip
	// for .zip file archive.
	PackageType types.PackageType
	
	// The latest updated revision of the function or alias.
	RevisionId *string
	
	// The function's execution role.
	Role *string
	
	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html)
	// . Runtime is required if the deployment package is a .zip file archive. The
	// following list includes deprecated runtimes. For more information, see Runtime
	// deprecation policy (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy)
	// .
	Runtime types.Runtime
	
	// The ARN of the runtime and any errors that occured.
	RuntimeVersionConfig *types.RuntimeVersionConfig
	
	// The ARN of the signing job.
	SigningJobArn *string
	
	// The ARN of the signing profile version.
	SigningProfileVersionArn *string
	
	// Set ApplyOn to PublishedVersions to create a snapshot of the initialized
	// execution environment when you publish a function version. For more information,
	// see Improving startup performance with Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
	// .
	SnapStart *types.SnapStartResponse
	
	// The current state of the function. When the state is Inactive , you can
	// reactivate the function by invoking it.
	State types.State
	
	// The reason for the function's current state.
	StateReason *string
	
	// The reason code for the function's current state. When the code is Creating ,
	// you can't invoke or modify the function.
	StateReasonCode types.StateReasonCode
	
	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32
	
	// The function's X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse
	
	// The version of the Lambda function.
	Version *string
	
	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFunctionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionConfiguration{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionConfiguration{}, middleware.After)
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
	if err = addGetFunctionConfigurationResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpGetFunctionConfigurationValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionConfiguration(options.Region, ), middleware.Before); err != nil {
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

// GetFunctionConfigurationAPIClient is a client that implements the
// GetFunctionConfiguration operation.
type GetFunctionConfigurationAPIClient interface {
	GetFunctionConfiguration(context.Context, *GetFunctionConfigurationInput, ...func(*Options)) (*GetFunctionConfigurationOutput, error)
}

var _ GetFunctionConfigurationAPIClient = (*Client)(nil)

// FunctionActiveWaiterOptions are waiter options for FunctionActiveWaiter
type FunctionActiveWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// FunctionActiveWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FunctionActiveWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *GetFunctionConfigurationInput, *GetFunctionConfigurationOutput, error) (bool, error)
}

// FunctionActiveWaiter defines the waiters for FunctionActive
type FunctionActiveWaiter struct {
	
	client GetFunctionConfigurationAPIClient
	
	options FunctionActiveWaiterOptions
}

// NewFunctionActiveWaiter constructs a FunctionActiveWaiter.
func NewFunctionActiveWaiter(client GetFunctionConfigurationAPIClient, optFns ...func(*FunctionActiveWaiterOptions)) *FunctionActiveWaiter {
	options := FunctionActiveWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = functionActiveStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &FunctionActiveWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for FunctionActive waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *FunctionActiveWaiter) Wait(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*FunctionActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FunctionActive waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *FunctionActiveWaiter) WaitForOutput(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*FunctionActiveWaiterOptions)) (*GetFunctionConfigurationOutput, error) {
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
		
		out, err := w.client.GetFunctionConfiguration(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FunctionActive waiter")
}

func functionActiveStateRetryable(ctx context.Context, input *GetFunctionConfigurationInput, output *GetFunctionConfigurationOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Active"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Failed"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Pending"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	return true, nil
}

// FunctionUpdatedWaiterOptions are waiter options for FunctionUpdatedWaiter
type FunctionUpdatedWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// FunctionUpdatedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FunctionUpdatedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *GetFunctionConfigurationInput, *GetFunctionConfigurationOutput, error) (bool, error)
}

// FunctionUpdatedWaiter defines the waiters for FunctionUpdated
type FunctionUpdatedWaiter struct {
	
	client GetFunctionConfigurationAPIClient
	
	options FunctionUpdatedWaiterOptions
}

// NewFunctionUpdatedWaiter constructs a FunctionUpdatedWaiter.
func NewFunctionUpdatedWaiter(client GetFunctionConfigurationAPIClient, optFns ...func(*FunctionUpdatedWaiterOptions)) *FunctionUpdatedWaiter {
	options := FunctionUpdatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = functionUpdatedStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &FunctionUpdatedWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for FunctionUpdated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *FunctionUpdatedWaiter) Wait(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*FunctionUpdatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FunctionUpdated waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *FunctionUpdatedWaiter) WaitForOutput(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*FunctionUpdatedWaiterOptions)) (*GetFunctionConfigurationOutput, error) {
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
		
		out, err := w.client.GetFunctionConfiguration(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FunctionUpdated waiter")
}

func functionUpdatedStateRetryable(ctx context.Context, input *GetFunctionConfigurationInput, output *GetFunctionConfigurationOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("LastUpdateStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Successful"
		value, ok := pathValue.(types.LastUpdateStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.LastUpdateStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("LastUpdateStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Failed"
		value, ok := pathValue.(types.LastUpdateStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.LastUpdateStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("LastUpdateStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "InProgress"
		value, ok := pathValue.(types.LastUpdateStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.LastUpdateStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	return true, nil
}

// PublishedVersionActiveWaiterOptions are waiter options for
// PublishedVersionActiveWaiter
type PublishedVersionActiveWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// PublishedVersionActiveWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, PublishedVersionActiveWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
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
	Retryable func(context.Context, *GetFunctionConfigurationInput, *GetFunctionConfigurationOutput, error) (bool, error)
}

// PublishedVersionActiveWaiter defines the waiters for PublishedVersionActive
type PublishedVersionActiveWaiter struct {
	
	client GetFunctionConfigurationAPIClient
	
	options PublishedVersionActiveWaiterOptions
}

// NewPublishedVersionActiveWaiter constructs a PublishedVersionActiveWaiter.
func NewPublishedVersionActiveWaiter(client GetFunctionConfigurationAPIClient, optFns ...func(*PublishedVersionActiveWaiterOptions)) *PublishedVersionActiveWaiter {
	options := PublishedVersionActiveWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = publishedVersionActiveStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &PublishedVersionActiveWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for PublishedVersionActive waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *PublishedVersionActiveWaiter) Wait(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*PublishedVersionActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PublishedVersionActive waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *PublishedVersionActiveWaiter) WaitForOutput(ctx context.Context, params *GetFunctionConfigurationInput, maxWaitDur time.Duration, optFns ...func(*PublishedVersionActiveWaiterOptions)) (*GetFunctionConfigurationOutput, error) {
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
		
		out, err := w.client.GetFunctionConfiguration(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PublishedVersionActive waiter")
}

func publishedVersionActiveStateRetryable(ctx context.Context, input *GetFunctionConfigurationInput, output *GetFunctionConfigurationOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Active"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Failed"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "Pending"
		value, ok := pathValue.(types.State)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.State value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opGetFunctionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "lambda",
	OperationName: "GetFunctionConfiguration",
	}
}

type opGetFunctionConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetFunctionConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetFunctionConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "lambda"
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
	        signingName = "lambda"
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
	        v4aScheme.SigningName = aws.String("lambda")
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

func addGetFunctionConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetFunctionConfigurationResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
