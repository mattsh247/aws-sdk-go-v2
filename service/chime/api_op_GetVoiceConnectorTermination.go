// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves termination setting details for the specified Amazon Chime Voice
// Connector. This API is is no longer supported and will not be updated. We
// recommend using the latest version, GetVoiceConnectorTermination (https://docs.aws.amazon.com/chime-sdk/latest/APIReference/API_voice-chime_GetVoiceConnectorTermination.html)
// , in the Amazon Chime SDK. Using the latest version requires migrating to a
// dedicated namespace. For more information, refer to Migrating from the Amazon
// Chime namespace (https://docs.aws.amazon.com/chime-sdk/latest/dg/migrate-from-chm-namespace.html)
// in the Amazon Chime SDK Developer Guide.
//
// Deprecated: Replaced by GetVoiceConnectorTermination in the Amazon Chime SDK
// Voice Namespace
func (c *Client) GetVoiceConnectorTermination(ctx context.Context, params *GetVoiceConnectorTerminationInput, optFns ...func(*Options)) (*GetVoiceConnectorTerminationOutput, error) {
	if params == nil {
		params = &GetVoiceConnectorTerminationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVoiceConnectorTermination", params, optFns, c.addOperationGetVoiceConnectorTerminationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVoiceConnectorTerminationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetVoiceConnectorTerminationInput struct {

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	noSmithyDocumentSerde
}

type GetVoiceConnectorTerminationOutput struct {

	// The termination setting details.
	Termination *types.Termination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetVoiceConnectorTerminationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVoiceConnectorTermination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVoiceConnectorTermination{}, middleware.After)
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
	if err = addOpGetVoiceConnectorTerminationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetVoiceConnectorTermination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetVoiceConnectorTermination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetVoiceConnectorTermination",
	}
}
