// Code generated by smithy-go-codegen DO NOT EDIT.


package ec2query

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
)

type validateOpEndpointWithHostLabelOperation struct {
}

func (*validateOpEndpointWithHostLabelOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEndpointWithHostLabelOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEndpointWithHostLabelOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpEndpointWithHostLabelOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEndpointWithHostLabelOperation{}, middleware.After)
}

func validateOpEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointWithHostLabelOperationInput"}
	if v.Label == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
