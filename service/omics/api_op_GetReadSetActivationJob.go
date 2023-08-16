// Code generated by smithy-go-codegen DO NOT EDIT.


package omics

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
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Gets information about a read set activation job.
func (c *Client) GetReadSetActivationJob(ctx context.Context, params *GetReadSetActivationJobInput, optFns ...func(*Options)) (*GetReadSetActivationJobOutput, error) {
	if params == nil { params = &GetReadSetActivationJobInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "GetReadSetActivationJob", params, optFns, c.addOperationGetReadSetActivationJobMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*GetReadSetActivationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReadSetActivationJobInput struct {
	
	// The job's ID.
	//
	// This member is required.
	Id *string
	
	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string
	
	noSmithyDocumentSerde
}

type GetReadSetActivationJobOutput struct {
	
	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time
	
	// The job's ID.
	//
	// This member is required.
	Id *string
	
	// The job's sequence store ID.
	//
	// This member is required.
	SequenceStoreId *string
	
	// The job's status.
	//
	// This member is required.
	Status types.ReadSetActivationJobStatus
	
	// When the job completed.
	CompletionTime *time.Time
	
	// The job's source files.
	Sources []types.ActivateReadSetSourceItem
	
	// The job's status message.
	StatusMessage *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReadSetActivationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReadSetActivationJob{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReadSetActivationJob{}, middleware.After)
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
	if err = addEndpointPrefix_opGetReadSetActivationJobMiddleware(stack); err != nil {
	return err
	}
	if err = addGetReadSetActivationJobResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpGetReadSetActivationJobValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReadSetActivationJob(options.Region, ), middleware.Before); err != nil {
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

type endpointPrefix_opGetReadSetActivationJobMiddleware struct {
}

func (*endpointPrefix_opGetReadSetActivationJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetReadSetActivationJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}
	
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}
	
	req.URL.Host = "control-storage-" + req.URL.Host
	
	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetReadSetActivationJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetReadSetActivationJobMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetReadSetActivationJobAPIClient is a client that implements the
// GetReadSetActivationJob operation.
type GetReadSetActivationJobAPIClient interface {
	GetReadSetActivationJob(context.Context, *GetReadSetActivationJobInput, ...func(*Options)) (*GetReadSetActivationJobOutput, error)
}

var _ GetReadSetActivationJobAPIClient = (*Client)(nil)

// ReadSetActivationJobCompletedWaiterOptions are waiter options for
// ReadSetActivationJobCompletedWaiter
type ReadSetActivationJobCompletedWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ReadSetActivationJobCompletedWaiter will use default minimum delay of 30
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, ReadSetActivationJobCompletedWaiter will use default max delay of
	// 600 seconds. Note that MaxDelay must resolve to value greater than or equal to
	// the MinDelay.
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
	Retryable func(context.Context, *GetReadSetActivationJobInput, *GetReadSetActivationJobOutput, error) (bool, error)
}

// ReadSetActivationJobCompletedWaiter defines the waiters for
// ReadSetActivationJobCompleted
type ReadSetActivationJobCompletedWaiter struct {
	
	client GetReadSetActivationJobAPIClient
	
	options ReadSetActivationJobCompletedWaiterOptions
}

// NewReadSetActivationJobCompletedWaiter constructs a
// ReadSetActivationJobCompletedWaiter.
func NewReadSetActivationJobCompletedWaiter(client GetReadSetActivationJobAPIClient, optFns ...func(*ReadSetActivationJobCompletedWaiterOptions)) *ReadSetActivationJobCompletedWaiter {
	options := ReadSetActivationJobCompletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = readSetActivationJobCompletedStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &ReadSetActivationJobCompletedWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for ReadSetActivationJobCompleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ReadSetActivationJobCompletedWaiter) Wait(ctx context.Context, params *GetReadSetActivationJobInput, maxWaitDur time.Duration, optFns ...func(*ReadSetActivationJobCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ReadSetActivationJobCompleted
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ReadSetActivationJobCompletedWaiter) WaitForOutput(ctx context.Context, params *GetReadSetActivationJobInput, maxWaitDur time.Duration, optFns ...func(*ReadSetActivationJobCompletedWaiterOptions)) (*GetReadSetActivationJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}
	
	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}
	
	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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
		
		out, err := w.client.GetReadSetActivationJob(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ReadSetActivationJobCompleted waiter")
}

func readSetActivationJobCompletedStateRetryable(ctx context.Context, input *GetReadSetActivationJobInput, output *GetReadSetActivationJobOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "SUBMITTED"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "IN_PROGRESS"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "CANCELLING"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return true, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "FAILED"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "COMPLETED_WITH_FAILURES"
		value, ok := pathValue.(types.ReadSetActivationJobStatus)
		if !ok {
		return false, fmt.Errorf("waiter comparator expected types.ReadSetActivationJobStatus value, got %T", pathValue)}
		
		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opGetReadSetActivationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "omics",
	OperationName: "GetReadSetActivationJob",
	}
}

type opGetReadSetActivationJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opGetReadSetActivationJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetReadSetActivationJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "omics"
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
	        signingName = "omics"
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
	        v4aScheme.SigningName = aws.String("omics")
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

func addGetReadSetActivationJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opGetReadSetActivationJobResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
