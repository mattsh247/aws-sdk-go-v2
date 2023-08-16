// Code generated by smithy-go-codegen DO NOT EDIT.


package sagemakerruntime

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
)

type validateOpInvokeEndpointAsync struct {
}

func (*validateOpInvokeEndpointAsync) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeEndpointAsync) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeEndpointAsyncInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeEndpointAsyncInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInvokeEndpoint struct {
}

func (*validateOpInvokeEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpInvokeEndpointAsyncValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeEndpointAsync{}, middleware.After)
}

func addOpInvokeEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeEndpoint{}, middleware.After)
}

func validateOpInvokeEndpointAsyncInput(v *InvokeEndpointAsyncInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeEndpointAsyncInput"}
	if v.EndpointName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("EndpointName"))
	}
	if v.InputLocation == nil {
	invalidParams.Add(smithy.NewErrParamRequired("InputLocation"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpInvokeEndpointInput(v *InvokeEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeEndpointInput"}
	if v.EndpointName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("EndpointName"))
	}
	if v.Body == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Body"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
