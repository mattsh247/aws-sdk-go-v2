// Code generated by smithy-go-codegen DO NOT EDIT.


package awsrestjson

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Null and empty headers are not sent over the wire.
func (c *Client) NullAndEmptyHeadersServer(ctx context.Context, params *NullAndEmptyHeadersServerInput, optFns ...func(*Options)) (*NullAndEmptyHeadersServerOutput, error) {
	if params == nil { params = &NullAndEmptyHeadersServerInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "NullAndEmptyHeadersServer", params, optFns, c.addOperationNullAndEmptyHeadersServerMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*NullAndEmptyHeadersServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NullAndEmptyHeadersServerInput struct {
	
	A *string
	
	B *string
	
	C []string
	
	noSmithyDocumentSerde
}

type NullAndEmptyHeadersServerOutput struct {
	
	A *string
	
	B *string
	
	C []string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationNullAndEmptyHeadersServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpNullAndEmptyHeadersServer{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpNullAndEmptyHeadersServer{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNullAndEmptyHeadersServer(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNullAndEmptyHeadersServer(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "NullAndEmptyHeadersServer",
	}
}
