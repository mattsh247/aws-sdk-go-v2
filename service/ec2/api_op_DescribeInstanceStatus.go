// Code generated by smithy-go-codegen DO NOT EDIT.


package ec2

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
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Describes the status of the specified instances or all of your instances. By
// default, only running instances are described, unless you specifically indicate
// to return the status of all instances. Instance status includes the following
// components:
//   - Status checks - Amazon EC2 performs status checks on running EC2 instances
//   to identify hardware and software issues. For more information, see Status
//   checks for your instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)
//   and Troubleshoot instances with failed status checks (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html)
//   in the Amazon EC2 User Guide.
//   - Scheduled events - Amazon EC2 can schedule events (such as reboot, stop, or
//   terminate) for your instances related to hardware issues, software updates, or
//   system maintenance. For more information, see Scheduled events for your
//   instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html)
//   in the Amazon EC2 User Guide.
//   - Instance state - You can manage your instances from the moment you launch
//   them through their termination. For more information, see Instance lifecycle (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
//   in the Amazon EC2 User Guide.
func (c *Client) DescribeInstanceStatus(ctx context.Context, params *DescribeInstanceStatusInput, optFns ...func(*Options)) (*DescribeInstanceStatusOutput, error) {
	if params == nil { params = &DescribeInstanceStatusInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceStatus", params, optFns, c.addOperationDescribeInstanceStatusMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeInstanceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceStatusInput struct {
	
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool
	
	// The filters.
	//   - availability-zone - The Availability Zone of the instance.
	//   - event.code - The code for the scheduled event ( instance-reboot |
	//   system-reboot | system-maintenance | instance-retirement | instance-stop ).
	//   - event.description - A description of the event.
	//   - event.instance-event-id - The ID of the event whose date and time you are
	//   modifying.
	//   - event.not-after - The latest end time for the scheduled event (for example,
	//   2014-09-15T17:15:20.000Z ).
	//   - event.not-before - The earliest start time for the scheduled event (for
	//   example, 2014-09-15T17:15:20.000Z ).
	//   - event.not-before-deadline - The deadline for starting the event (for
	//   example, 2014-09-15T17:15:20.000Z ).
	//   - instance-state-code - The code for the instance state, as a 16-bit unsigned
	//   integer. The high byte is used for internal purposes and should be ignored. The
	//   low byte is set based on the state represented. The valid values are 0
	//   (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping), and
	//   80 (stopped).
	//   - instance-state-name - The state of the instance ( pending | running |
	//   shutting-down | terminated | stopping | stopped ).
	//   - instance-status.reachability - Filters on instance status where the name is
	//   reachability ( passed | failed | initializing | insufficient-data ).
	//   - instance-status.status - The status of the instance ( ok | impaired |
	//   initializing | insufficient-data | not-applicable ).
	//   - system-status.reachability - Filters on system status where the name is
	//   reachability ( passed | failed | initializing | insufficient-data ).
	//   - system-status.status - The system status of the instance ( ok | impaired |
	//   initializing | insufficient-data | not-applicable ).
	Filters []types.Filter
	
	// When true , includes the health status for all instances. When false , includes
	// the health status for running instances only. Default: false
	IncludeAllInstances *bool
	
	// The instance IDs. Default: Describes all your instances. Constraints: Maximum
	// 100 explicitly specified instance IDs.
	InstanceIds []string
	
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the instance IDs parameter in the same
	// request.
	MaxResults *int32
	
	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string
	
	noSmithyDocumentSerde
}

type DescribeInstanceStatusOutput struct {
	
	// Information about the status of the instances.
	InstanceStatuses []types.InstanceStatus
	
	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceStatus{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceStatus{}, middleware.After)
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
	if err = addDescribeInstanceStatusResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceStatus(options.Region, ), middleware.Before); err != nil {
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

// DescribeInstanceStatusAPIClient is a client that implements the
// DescribeInstanceStatus operation.
type DescribeInstanceStatusAPIClient interface {
	DescribeInstanceStatus(context.Context, *DescribeInstanceStatusInput, ...func(*Options)) (*DescribeInstanceStatusOutput, error)
}

var _ DescribeInstanceStatusAPIClient = (*Client)(nil)

// DescribeInstanceStatusPaginatorOptions is the paginator options for
// DescribeInstanceStatus
type DescribeInstanceStatusPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the instance IDs parameter in the same
	// request.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceStatusPaginator is a paginator for DescribeInstanceStatus
type DescribeInstanceStatusPaginator struct {
    options DescribeInstanceStatusPaginatorOptions
    client DescribeInstanceStatusAPIClient
    params *DescribeInstanceStatusInput
    nextToken *string
    firstPage bool
}

// NewDescribeInstanceStatusPaginator returns a new DescribeInstanceStatusPaginator
func NewDescribeInstanceStatusPaginator(client DescribeInstanceStatusAPIClient, params *DescribeInstanceStatusInput, optFns ...func(*DescribeInstanceStatusPaginatorOptions)) *DescribeInstanceStatusPaginator {
	if params == nil {
	    params = &DescribeInstanceStatusInput{}
	}
	
	options := DescribeInstanceStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeInstanceStatusPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeInstanceStatus page.
func (p *DescribeInstanceStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceStatusOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.NextToken = p.nextToken
	
	var limit *int32
	if p.options.Limit > 0 {
	    limit = &p.options.Limit
	}
	params.MaxResults = limit
	
	result, err := p.client.DescribeInstanceStatus(ctx, &params, optFns...)
	if err != nil {
	    return nil, err
	}
	p.firstPage = false
	
	prevToken := p.nextToken
	p.nextToken = result.NextToken
	
	if p.options.StopOnDuplicateToken &&
	    prevToken != nil &&
	    p.nextToken != nil &&
	    *prevToken == *p.nextToken {
	    p.nextToken = nil
	}
	
	return result, nil
}

// InstanceStatusOkWaiterOptions are waiter options for InstanceStatusOkWaiter
type InstanceStatusOkWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceStatusOkWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, InstanceStatusOkWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeInstanceStatusInput, *DescribeInstanceStatusOutput, error) (bool, error)
}

// InstanceStatusOkWaiter defines the waiters for InstanceStatusOk
type InstanceStatusOkWaiter struct {
	
	client DescribeInstanceStatusAPIClient
	
	options InstanceStatusOkWaiterOptions
}

// NewInstanceStatusOkWaiter constructs a InstanceStatusOkWaiter.
func NewInstanceStatusOkWaiter(client DescribeInstanceStatusAPIClient, optFns ...func(*InstanceStatusOkWaiterOptions)) *InstanceStatusOkWaiter {
	options := InstanceStatusOkWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceStatusOkStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceStatusOkWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceStatusOk waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *InstanceStatusOkWaiter) Wait(ctx context.Context, params *DescribeInstanceStatusInput, maxWaitDur time.Duration, optFns ...func(*InstanceStatusOkWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for InstanceStatusOk waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *InstanceStatusOkWaiter) WaitForOutput(ctx context.Context, params *DescribeInstanceStatusInput, maxWaitDur time.Duration, optFns ...func(*InstanceStatusOkWaiterOptions)) (*DescribeInstanceStatusOutput, error) {
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
		
		out, err := w.client.DescribeInstanceStatus(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for InstanceStatusOk waiter")
}

func instanceStatusOkStateRetryable(ctx context.Context, input *DescribeInstanceStatusInput, output *DescribeInstanceStatusOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("InstanceStatuses[].InstanceStatus.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "ok"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		 if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}
		
		if len(listOfValues) == 0 { match = false }
		for _, v := range listOfValues {
			value, ok := v.(types.SummaryStatus)
			if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SummaryStatus value, got %T", pathValue)}
			
			if string(value) != expectedValue { match = false }
		}
		
		if match {
			return false, nil
		}
	}
	
	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}
		
		if "InvalidInstanceID.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}
	
	return true, nil
}

// SystemStatusOkWaiterOptions are waiter options for SystemStatusOkWaiter
type SystemStatusOkWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SystemStatusOkWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, SystemStatusOkWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeInstanceStatusInput, *DescribeInstanceStatusOutput, error) (bool, error)
}

// SystemStatusOkWaiter defines the waiters for SystemStatusOk
type SystemStatusOkWaiter struct {
	
	client DescribeInstanceStatusAPIClient
	
	options SystemStatusOkWaiterOptions
}

// NewSystemStatusOkWaiter constructs a SystemStatusOkWaiter.
func NewSystemStatusOkWaiter(client DescribeInstanceStatusAPIClient, optFns ...func(*SystemStatusOkWaiterOptions)) *SystemStatusOkWaiter {
	options := SystemStatusOkWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = systemStatusOkStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &SystemStatusOkWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for SystemStatusOk waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *SystemStatusOkWaiter) Wait(ctx context.Context, params *DescribeInstanceStatusInput, maxWaitDur time.Duration, optFns ...func(*SystemStatusOkWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SystemStatusOk waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *SystemStatusOkWaiter) WaitForOutput(ctx context.Context, params *DescribeInstanceStatusInput, maxWaitDur time.Duration, optFns ...func(*SystemStatusOkWaiterOptions)) (*DescribeInstanceStatusOutput, error) {
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
		
		out, err := w.client.DescribeInstanceStatus(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for SystemStatusOk waiter")
}

func systemStatusOkStateRetryable(ctx context.Context, input *DescribeInstanceStatusInput, output *DescribeInstanceStatusOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("InstanceStatuses[].SystemStatus.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "ok"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		 if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}
		
		if len(listOfValues) == 0 { match = false }
		for _, v := range listOfValues {
			value, ok := v.(types.SummaryStatus)
			if !ok {
			return false, fmt.Errorf("waiter comparator expected types.SummaryStatus value, got %T", pathValue)}
			
			if string(value) != expectedValue { match = false }
		}
		
		if match {
			return false, nil
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opDescribeInstanceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ec2",
	OperationName: "DescribeInstanceStatus",
	}
}

type opDescribeInstanceStatusResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeInstanceStatusResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeInstanceStatusResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ec2"
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
	        signingName = "ec2"
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
	        v4aScheme.SigningName = aws.String("ec2")
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

func addDescribeInstanceStatusResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeInstanceStatusResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
