// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an existing state machine by modifying its definition , roleArn , or
// loggingConfiguration . Running executions will continue to use the previous
// definition and roleArn . You must include at least one of definition or roleArn
// or you will receive a MissingRequiredParameter error. A qualified state machine
// ARN refers to a Distributed Map state defined within a state machine. For
// example, the qualified state machine ARN
// arn:partition:states:region:account-id:stateMachine:stateMachineName/mapStateLabel
// refers to a Distributed Map state with a label mapStateLabel in the state
// machine named stateMachineName . A qualified state machine ARN can either refer
// to a Distributed Map state defined within a state machine, a version ARN, or an
// alias ARN. The following are some examples of qualified and unqualified state
// machine ARNs:
//   - The following qualified state machine ARN refers to a Distributed Map state
//     with a label mapStateLabel in a state machine named myStateMachine .
//     arn:partition:states:region:account-id:stateMachine:myStateMachine/mapStateLabel
//     If you provide a qualified state machine ARN that refers to a Distributed Map
//     state, the request fails with ValidationException .
//   - The following qualified state machine ARN refers to an alias named PROD .
//     arn::states:::stateMachine: If you provide a qualified state machine ARN that
//     refers to a version ARN or an alias ARN, the request starts execution for that
//     version or alias.
//   - The following unqualified state machine ARN refers to a state machine named
//     myStateMachine . arn::states:::stateMachine:
//
// After you update your state machine, you can set the publish parameter to true
// in the same action to publish a new version (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html)
// . This way, you can opt-in to strict versioning of your state machine. Step
// Functions assigns monotonically increasing integers for state machine versions,
// starting at version number 1. All StartExecution calls within a few seconds use
// the updated definition and roleArn . Executions started immediately after you
// call UpdateStateMachine may use the previous state machine definition and
// roleArn .
func (c *Client) UpdateStateMachine(ctx context.Context, params *UpdateStateMachineInput, optFns ...func(*Options)) (*UpdateStateMachineOutput, error) {
	if params == nil {
		params = &UpdateStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStateMachine", params, optFns, c.addOperationUpdateStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)
	// .
	Definition *string

	// Use the LoggingConfiguration data type to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// Specifies whether the state machine version is published. The default is false .
	// To publish a version after updating the state machine, set publish to true .
	Publish bool

	// The Amazon Resource Name (ARN) of the IAM role of the state machine.
	RoleArn *string

	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// An optional description of the state machine version to publish. You can only
	// specify the versionDescription parameter if you've set publish to true .
	VersionDescription *string

	noSmithyDocumentSerde
}

type UpdateStateMachineOutput struct {

	// The date and time the state machine was updated.
	//
	// This member is required.
	UpdateDate *time.Time

	// The revision identifier for the updated state machine.
	RevisionId *string

	// The Amazon Resource Name (ARN) of the published state machine version. If the
	// publish parameter isn't set to true , this field returns null.
	StateMachineVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateStateMachine{}, middleware.After)
	if err != nil {
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStateMachine(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "UpdateStateMachine",
	}
}
