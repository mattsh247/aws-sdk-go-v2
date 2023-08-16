// Code generated by smithy-go-codegen DO NOT EDIT.


package machinelearning

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
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns a list of MLModel that match the search criteria in the request.
func (c *Client) DescribeMLModels(ctx context.Context, params *DescribeMLModelsInput, optFns ...func(*Options)) (*DescribeMLModelsOutput, error) {
	if params == nil { params = &DescribeMLModelsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeMLModels", params, optFns, c.addOperationDescribeMLModelsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeMLModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMLModelsInput struct {
	
	// The equal to operator. The MLModel results will have FilterVariable values that
	// exactly match the value specified with EQ .
	EQ *string
	
	// Use one of the following variables to filter a list of MLModel :
	//   - CreatedAt - Sets the search criteria to MLModel creation date.
	//   - Status - Sets the search criteria to MLModel status.
	//   - Name - Sets the search criteria to the contents of MLModel Name .
	//   - IAMUser - Sets the search criteria to the user account that invoked the
	//   MLModel creation.
	//   - TrainingDataSourceId - Sets the search criteria to the DataSource used to
	//   train one or more MLModel .
	//   - RealtimeEndpointStatus - Sets the search criteria to the MLModel real-time
	//   endpoint status.
	//   - MLModelType - Sets the search criteria to MLModel type: binary, regression,
	//   or multi-class.
	//   - Algorithm - Sets the search criteria to the algorithm that the MLModel uses.
	//   - TrainingDataURI - Sets the search criteria to the data file(s) used in
	//   training a MLModel . The URL can identify either a file or an Amazon Simple
	//   Storage Service (Amazon S3) bucket or directory.
	FilterVariable types.MLModelFilterVariable
	
	// The greater than or equal to operator. The MLModel results will have
	// FilterVariable values that are greater than or equal to the value specified with
	// GE .
	GE *string
	
	// The greater than operator. The MLModel results will have FilterVariable values
	// that are greater than the value specified with GT .
	GT *string
	
	// The less than or equal to operator. The MLModel results will have FilterVariable
	// values that are less than or equal to the value specified with LE .
	LE *string
	
	// The less than operator. The MLModel results will have FilterVariable values
	// that are less than the value specified with LT .
	LT *string
	
	// The number of pages of information to include in the result. The range of
	// acceptable values is 1 through 100 . The default value is 100 .
	Limit *int32
	
	// The not equal to operator. The MLModel results will have FilterVariable values
	// not equal to the value specified with NE .
	NE *string
	
	// The ID of the page in the paginated results.
	NextToken *string
	
	// A string that is found at the beginning of a variable, such as Name or Id . For
	// example, an MLModel could have the Name 2014-09-09-HolidayGiftMailer . To search
	// for this MLModel , select Name for the FilterVariable and any of the following
	// strings for the Prefix :
	//   - 2014-09
	//   - 2014-09-09
	//   - 2014-09-09-Holiday
	Prefix *string
	
	// A two-value parameter that determines the sequence of the resulting list of
	// MLModel .
	//   - asc - Arranges the list in ascending order (A-Z, 0-9).
	//   - dsc - Arranges the list in descending order (Z-A, 9-0).
	// Results are sorted by FilterVariable .
	SortOrder types.SortOrder
	
	noSmithyDocumentSerde
}

// Represents the output of a DescribeMLModels operation. The content is
// essentially a list of MLModel .
type DescribeMLModelsOutput struct {
	
	// The ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string
	
	// A list of MLModel that meet the search criteria.
	Results []types.MLModel
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMLModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMLModels{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMLModels{}, middleware.After)
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
	if err = addDescribeMLModelsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMLModels(options.Region, ), middleware.Before); err != nil {
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

// DescribeMLModelsAPIClient is a client that implements the DescribeMLModels
// operation.
type DescribeMLModelsAPIClient interface {
	DescribeMLModels(context.Context, *DescribeMLModelsInput, ...func(*Options)) (*DescribeMLModelsOutput, error)
}

var _ DescribeMLModelsAPIClient = (*Client)(nil)

// DescribeMLModelsPaginatorOptions is the paginator options for DescribeMLModels
type DescribeMLModelsPaginatorOptions struct {
	// The number of pages of information to include in the result. The range of
	// acceptable values is 1 through 100 . The default value is 100 .
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMLModelsPaginator is a paginator for DescribeMLModels
type DescribeMLModelsPaginator struct {
    options DescribeMLModelsPaginatorOptions
    client DescribeMLModelsAPIClient
    params *DescribeMLModelsInput
    nextToken *string
    firstPage bool
}

// NewDescribeMLModelsPaginator returns a new DescribeMLModelsPaginator
func NewDescribeMLModelsPaginator(client DescribeMLModelsAPIClient, params *DescribeMLModelsInput, optFns ...func(*DescribeMLModelsPaginatorOptions)) *DescribeMLModelsPaginator {
	if params == nil {
	    params = &DescribeMLModelsInput{}
	}
	
	options := DescribeMLModelsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeMLModelsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMLModelsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeMLModels page.
func (p *DescribeMLModelsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMLModelsOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.NextToken = p.nextToken
	
	var limit *int32
	if p.options.Limit > 0 {
	    limit = &p.options.Limit
	}
	params.Limit = limit
	
	result, err := p.client.DescribeMLModels(ctx, &params, optFns...)
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

// MLModelAvailableWaiterOptions are waiter options for MLModelAvailableWaiter
type MLModelAvailableWaiterOptions struct {
	
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error
	
	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// MLModelAvailableWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration
	
	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, MLModelAvailableWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeMLModelsInput, *DescribeMLModelsOutput, error) (bool, error)
}

// MLModelAvailableWaiter defines the waiters for MLModelAvailable
type MLModelAvailableWaiter struct {
	
	client DescribeMLModelsAPIClient
	
	options MLModelAvailableWaiterOptions
}

// NewMLModelAvailableWaiter constructs a MLModelAvailableWaiter.
func NewMLModelAvailableWaiter(client DescribeMLModelsAPIClient, optFns ...func(*MLModelAvailableWaiterOptions)) *MLModelAvailableWaiter {
	options := MLModelAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = mLModelAvailableStateRetryable
	
	for _, fn := range optFns {
		fn(&options)
	}
	return &MLModelAvailableWaiter {
		client: client,
		options: options,
	}
}

// Wait calls the waiter function for MLModelAvailable waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *MLModelAvailableWaiter) Wait(ctx context.Context, params *DescribeMLModelsInput, maxWaitDur time.Duration, optFns ...func(*MLModelAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for MLModelAvailable waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *MLModelAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeMLModelsInput, maxWaitDur time.Duration, optFns ...func(*MLModelAvailableWaiterOptions)) (*DescribeMLModelsOutput, error) {
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
		
		out, err := w.client.DescribeMLModels(ctx, params, func (o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for MLModelAvailable waiter")
}

func mLModelAvailableStateRetryable(ctx context.Context, input *DescribeMLModelsInput, output *DescribeMLModelsOutput, err error) (bool, error) {
	
	if err == nil {
		pathValue, err :=  jmespath.Search("Results[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "COMPLETED"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		 if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}
		
		if len(listOfValues) == 0 { match = false }
		for _, v := range listOfValues {
			value, ok := v.(types.EntityStatus)
			if !ok {
			return false, fmt.Errorf("waiter comparator expected types.EntityStatus value, got %T", pathValue)}
			
			if string(value) != expectedValue { match = false }
		}
		
		if match {
			return false, nil
		}
	}
	
	if err == nil {
		pathValue, err :=  jmespath.Search("Results[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}
		
		expectedValue := "FAILED"
		listOfValues, ok := pathValue.([]interface{})
		 if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}
		
		for _, v := range listOfValues {
			value, ok := v.(types.EntityStatus)
			if !ok {
			return false, fmt.Errorf("waiter comparator expected types.EntityStatus value, got %T", pathValue)}
			
			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}
	
	return true, nil
}

func newServiceMetadataMiddleware_opDescribeMLModels(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "machinelearning",
	OperationName: "DescribeMLModels",
	}
}

type opDescribeMLModelsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeMLModelsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeMLModelsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "machinelearning"
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
	        signingName = "machinelearning"
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
	        v4aScheme.SigningName = aws.String("machinelearning")
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

func addDescribeMLModelsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeMLModelsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
