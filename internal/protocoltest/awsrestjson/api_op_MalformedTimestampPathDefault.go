// Code generated by smithy-go-codegen DO NOT EDIT.


package awsrestjson

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

func (c *Client) MalformedTimestampPathDefault(ctx context.Context, params *MalformedTimestampPathDefaultInput, optFns ...func(*Options)) (*MalformedTimestampPathDefaultOutput, error) {
	if params == nil { params = &MalformedTimestampPathDefaultInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "MalformedTimestampPathDefault", params, optFns, c.addOperationMalformedTimestampPathDefaultMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*MalformedTimestampPathDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MalformedTimestampPathDefaultInput struct {
	
	// This member is required.
	Timestamp *time.Time
	
	noSmithyDocumentSerde
}

type MalformedTimestampPathDefaultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationMalformedTimestampPathDefaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpMalformedTimestampPathDefault{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpMalformedTimestampPathDefault{}, middleware.After)
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
	if err = addOpMalformedTimestampPathDefaultValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMalformedTimestampPathDefault(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMalformedTimestampPathDefault(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "MalformedTimestampPathDefault",
	}
}
