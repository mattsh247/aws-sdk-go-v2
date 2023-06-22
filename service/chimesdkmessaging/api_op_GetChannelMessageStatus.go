// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets message status for a specified messageId . Use this API to determine the
// intermediate status of messages going through channel flow processing. The API
// provides an alternative to retrieving message status if the event was not
// received because a client wasn't connected to a websocket. Messages can have any
// one of these statuses. SENT Message processed successfully PENDING Ongoing
// processing FAILED Processing failed DENIED Message denied by the processor
//   - This API does not return statuses for denied messages, because we don't
//     store them once the processor denies them.
//   - Only the message sender can invoke this API.
//   - The x-amz-chime-bearer request header is mandatory. Use the ARN of the
//     AppInstanceUser or AppInstanceBot that makes the API call as the value in the
//     header.
func (c *Client) GetChannelMessageStatus(ctx context.Context, params *GetChannelMessageStatusInput, optFns ...func(*Options)) (*GetChannelMessageStatusOutput, error) {
	if params == nil {
		params = &GetChannelMessageStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChannelMessageStatus", params, optFns, c.addOperationGetChannelMessageStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChannelMessageStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChannelMessageStatusInput struct {

	// The ARN of the channel
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user making the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The ID of the message.
	//
	// This member is required.
	MessageId *string

	// The ID of the SubChannel in the request. Only required when getting message
	// status in a SubChannel that the user belongs to.
	SubChannelId *string

	noSmithyDocumentSerde
}

type GetChannelMessageStatusOutput struct {

	// The message status and details.
	Status *types.ChannelMessageStatusStructure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChannelMessageStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetChannelMessageStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetChannelMessageStatus{}, middleware.After)
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
	if err = addOpGetChannelMessageStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChannelMessageStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetChannelMessageStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetChannelMessageStatus",
	}
}
