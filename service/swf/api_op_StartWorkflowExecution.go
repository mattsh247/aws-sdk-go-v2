// Code generated by smithy-go-codegen DO NOT EDIT.


package swf

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
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Starts an execution of the workflow type in the specified domain using the
// provided workflowId and input data. This action returns the newly started
// workflow execution. Access Control You can use IAM policies to control this
// action's access to Amazon SWF resources as follows:
//   - Use a Resource element with the domain name to limit the action to only
//   specified domains.
//   - Use an Action element to allow or deny permission to call this action.
//   - Constrain the following parameters by using a Condition element with the
//   appropriate keys.
//   - tagList.member.0 : The key is swf:tagList.member.0 .
//   - tagList.member.1 : The key is swf:tagList.member.1 .
//   - tagList.member.2 : The key is swf:tagList.member.2 .
//   - tagList.member.3 : The key is swf:tagList.member.3 .
//   - tagList.member.4 : The key is swf:tagList.member.4 .
//   - taskList : String constraint. The key is swf:taskList.name .
//   - workflowType.name : String constraint. The key is swf:workflowType.name .
//   - workflowType.version : String constraint. The key is
//   swf:workflowType.version .
// If the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED .
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) StartWorkflowExecution(ctx context.Context, params *StartWorkflowExecutionInput, optFns ...func(*Options)) (*StartWorkflowExecutionOutput, error) {
	if params == nil { params = &StartWorkflowExecutionInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "StartWorkflowExecution", params, optFns, c.addOperationStartWorkflowExecutionMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*StartWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWorkflowExecutionInput struct {
	
	// The name of the domain in which the workflow execution is created. The
	// specified string must not contain a : (colon), / (slash), | (vertical bar), or
	// any control characters ( \u0000-\u001f | \u007f-\u009f ). Also, it must not be
	// the literal string arn .
	//
	// This member is required.
	Domain *string
	
	// The user defined identifier associated with the workflow execution. You can use
	// this to associate a custom identifier with the workflow execution. You may
	// specify the same identifier if a workflow execution is logically a restart of a
	// previous execution. You cannot have two open workflow executions with the same
	// workflowId at the same time within the same domain. The specified string must
	// not contain a : (colon), / (slash), | (vertical bar), or any control characters
	// ( \u0000-\u001f | \u007f-\u009f ). Also, it must not be the literal string arn .
	//
	// This member is required.
	WorkflowId *string
	
	// The type of the workflow to start.
	//
	// This member is required.
	WorkflowType *types.WorkflowType
	
	// If set, specifies the policy to use for the child workflow executions of this
	// workflow execution if it is terminated, by calling the
	// TerminateWorkflowExecution action explicitly or due to an expired timeout. This
	// policy overrides the default child policy specified when registering the
	// workflow type using RegisterWorkflowType . The supported child policies are:
	//   - TERMINATE – The child executions are terminated.
	//   - REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//   by recording a WorkflowExecutionCancelRequested event in its history. It is up
	//   to the decider to take appropriate actions when it receives an execution history
	//   with this event.
	//   - ABANDON – No action is taken. The child executions continue to run.
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration time
	// then a fault is returned.
	ChildPolicy types.ChildPolicy
	
	// The total duration for this workflow execution. This overrides the
	// defaultExecutionStartToCloseTimeout specified when registering the workflow
	// type. The duration is specified in seconds; an integer greater than or equal to
	// 0 . Exceeding this limit causes the workflow execution to time out. Unlike some
	// of the other timeout parameters in Amazon SWF, you cannot specify a value of
	// "NONE" for this timeout; there is a one-year max limit on the time that a
	// workflow execution can run. An execution start-to-close timeout must be
	// specified either through this parameter or as a default when the workflow type
	// is registered. If neither this parameter nor a default execution start-to-close
	// timeout is specified, a fault is returned.
	ExecutionStartToCloseTimeout *string
	
	// The input for the workflow execution. This is a free form string which should
	// be meaningful to the workflow you are starting. This input is made available to
	// the new workflow execution in the WorkflowExecutionStarted history event.
	Input *string
	
	// The IAM role to attach to this workflow execution. Executions of this workflow
	// type need IAM roles to invoke Lambda functions. If you don't attach an IAM role,
	// any attempt to schedule a Lambda task fails. This results in a
	// ScheduleLambdaFunctionFailed history event. For more information, see
	// https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html (https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	LambdaRole *string
	
	// The list of tags to associate with the workflow execution. You can specify a
	// maximum of 5 tags. You can list workflow executions with a specific tag by
	// calling ListOpenWorkflowExecutions or ListClosedWorkflowExecutions and
	// specifying a TagFilter .
	TagList []string
	
	// The task list to use for the decision tasks generated for this workflow
	// execution. This overrides the defaultTaskList specified when registering the
	// workflow type. A task list for this workflow execution must be specified either
	// as a default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default task list was specified at registration time then
	// a fault is returned. The specified string must not contain a : (colon), /
	// (slash), | (vertical bar), or any control characters ( \u0000-\u001f |
	// \u007f-\u009f ). Also, it must not be the literal string arn .
	TaskList *types.TaskList
	
	// The task priority to use for this workflow execution. This overrides any
	// default priority that was assigned when the workflow type was registered. If not
	// set, then the default task priority for the workflow type is used. Valid values
	// are integers that range from Java's Integer.MIN_VALUE (-2147483648) to
	// Integer.MAX_VALUE (2147483647). Higher numbers indicate higher priority. For
	// more information about setting task priority, see Setting Task Priority (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	TaskPriority *string
	
	// Specifies the maximum duration of decision tasks for this workflow execution.
	// This parameter overrides the defaultTaskStartToCloseTimout specified when
	// registering the workflow type using RegisterWorkflowType . The duration is
	// specified in seconds, an integer greater than or equal to 0 . You can use NONE
	// to specify unlimited duration. A task start-to-close timeout for this workflow
	// execution must be specified either as a default for the workflow type or through
	// this parameter. If neither this parameter is set nor a default task
	// start-to-close timeout was specified at registration time then a fault is
	// returned.
	TaskStartToCloseTimeout *string
	
	noSmithyDocumentSerde
}

// Specifies the runId of a workflow execution.
type StartWorkflowExecutionOutput struct {
	
	// The runId of a workflow execution. This ID is generated by the service and can
	// be used to uniquely identify the workflow execution within a domain.
	RunId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationStartWorkflowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartWorkflowExecution{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartWorkflowExecution{}, middleware.After)
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
	if err = addStartWorkflowExecutionResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpStartWorkflowExecutionValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartWorkflowExecution(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartWorkflowExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "swf",
	OperationName: "StartWorkflowExecution",
	}
}

type opStartWorkflowExecutionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opStartWorkflowExecutionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opStartWorkflowExecutionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "swf"
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
	        signingName = "swf"
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
	        v4aScheme.SigningName = aws.String("swf")
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

func addStartWorkflowExecutionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opStartWorkflowExecutionResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
