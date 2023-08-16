// Code generated by smithy-go-codegen DO NOT EDIT.


package query

import (
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

func (c *Client) EndpointWithHostLabelOperation(ctx context.Context, params *EndpointWithHostLabelOperationInput, optFns ...func(*Options)) (*EndpointWithHostLabelOperationOutput, error) {
	if params == nil { params = &EndpointWithHostLabelOperationInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "EndpointWithHostLabelOperation", params, optFns, c.addOperationEndpointWithHostLabelOperationMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*EndpointWithHostLabelOperationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EndpointWithHostLabelOperationInput struct {
	
	// This member is required.
	Label *string
	
	noSmithyDocumentSerde
}

type EndpointWithHostLabelOperationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationEndpointWithHostLabelOperationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEndpointWithHostLabelOperation{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEndpointWithHostLabelOperation{}, middleware.After)
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
	if err = addEndpointPrefix_opEndpointWithHostLabelOperationMiddleware(stack); err != nil {
	return err
	}
	if err = addOpEndpointWithHostLabelOperationValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEndpointWithHostLabelOperation(options.Region, ), middleware.Before); err != nil {
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

type endpointPrefix_opEndpointWithHostLabelOperationMiddleware struct {
}

func (*endpointPrefix_opEndpointWithHostLabelOperationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opEndpointWithHostLabelOperationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}
	
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}
	
	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}
	
	var prefix strings.Builder
	prefix.WriteString("foo.")
	if input.Label == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("Label forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.Label) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("Label forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.Label)}
	} else {
		prefix.WriteString(*input.Label)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host
	
	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opEndpointWithHostLabelOperationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opEndpointWithHostLabelOperationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opEndpointWithHostLabelOperation(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	OperationName: "EndpointWithHostLabelOperation",
	}
}
