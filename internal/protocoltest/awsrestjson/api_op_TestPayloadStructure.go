// Code generated by smithy-go-codegen DO NOT EDIT.


package awsrestjson

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
)

// This example operation serializes a payload targeting a structure. This
// enforces the same requirements as TestBodyStructure but with the body specified
// by the @httpPayload trait.
func (c *Client) TestPayloadStructure(ctx context.Context, params *TestPayloadStructureInput, optFns ...func(*Options)) (*TestPayloadStructureOutput, error) {
	if params == nil { params = &TestPayloadStructureInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "TestPayloadStructure", params, optFns, c.addOperationTestPayloadStructureMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*TestPayloadStructureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestPayloadStructureInput struct {
	
	PayloadConfig *types.PayloadConfig
	
	TestId *string
	
	noSmithyDocumentSerde
}

type TestPayloadStructureOutput struct {
	
	PayloadConfig *types.PayloadConfig
	
	TestId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationTestPayloadStructureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestPayloadStructure{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestPayloadStructure{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestPayloadStructure(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestPayloadStructure(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "TestPayloadStructure",
	}
}
