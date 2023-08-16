// Code generated by smithy-go-codegen DO NOT EDIT.


package query

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Flattened maps with @xmlNamespace and @xmlName
func (c *Client) FlattenedXmlMapWithXmlNamespace(ctx context.Context, params *FlattenedXmlMapWithXmlNamespaceInput, optFns ...func(*Options)) (*FlattenedXmlMapWithXmlNamespaceOutput, error) {
	if params == nil { params = &FlattenedXmlMapWithXmlNamespaceInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "FlattenedXmlMapWithXmlNamespace", params, optFns, c.addOperationFlattenedXmlMapWithXmlNamespaceMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*FlattenedXmlMapWithXmlNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FlattenedXmlMapWithXmlNamespaceInput struct {
	
	noSmithyDocumentSerde
}

type FlattenedXmlMapWithXmlNamespaceOutput struct {
	
	MyMap map[string]string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationFlattenedXmlMapWithXmlNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFlattenedXmlMapWithXmlNamespace{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFlattenedXmlMapWithXmlNamespace{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlNamespace(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFlattenedXmlMapWithXmlNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "FlattenedXmlMapWithXmlNamespace",
	}
}
