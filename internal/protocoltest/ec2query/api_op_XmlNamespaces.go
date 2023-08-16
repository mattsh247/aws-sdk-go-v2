// Code generated by smithy-go-codegen DO NOT EDIT.


package ec2query

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
)

func (c *Client) XmlNamespaces(ctx context.Context, params *XmlNamespacesInput, optFns ...func(*Options)) (*XmlNamespacesOutput, error) {
	if params == nil { params = &XmlNamespacesInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "XmlNamespaces", params, optFns, c.addOperationXmlNamespacesMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*XmlNamespacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlNamespacesInput struct {
	
	noSmithyDocumentSerde
}

type XmlNamespacesOutput struct {
	
	Nested *types.XmlNamespaceNested
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationXmlNamespacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpXmlNamespaces{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpXmlNamespaces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opXmlNamespaces(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opXmlNamespaces(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "XmlNamespaces",
	}
}
