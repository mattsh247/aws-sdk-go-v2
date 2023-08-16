// Code generated by smithy-go-codegen DO NOT EDIT.


package sagemakera2iruntime

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
)

type validateOpDeleteHumanLoop struct {
}

func (*validateOpDeleteHumanLoop) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteHumanLoop) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteHumanLoopInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteHumanLoopInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeHumanLoop struct {
}

func (*validateOpDescribeHumanLoop) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeHumanLoop) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeHumanLoopInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeHumanLoopInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListHumanLoops struct {
}

func (*validateOpListHumanLoops) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListHumanLoops) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListHumanLoopsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListHumanLoopsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartHumanLoop struct {
}

func (*validateOpStartHumanLoop) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartHumanLoop) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartHumanLoopInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartHumanLoopInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopHumanLoop struct {
}

func (*validateOpStopHumanLoop) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopHumanLoop) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopHumanLoopInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopHumanLoopInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteHumanLoopValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteHumanLoop{}, middleware.After)
}

func addOpDescribeHumanLoopValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeHumanLoop{}, middleware.After)
}

func addOpListHumanLoopsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListHumanLoops{}, middleware.After)
}

func addOpStartHumanLoopValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartHumanLoop{}, middleware.After)
}

func addOpStopHumanLoopValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopHumanLoop{}, middleware.After)
}

func validateHumanLoopDataAttributes(v *types.HumanLoopDataAttributes) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HumanLoopDataAttributes"}
	if v.ContentClassifiers == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ContentClassifiers"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateHumanLoopInput(v *types.HumanLoopInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HumanLoopInput"}
	if v.InputContent == nil {
	invalidParams.Add(smithy.NewErrParamRequired("InputContent"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteHumanLoopInput(v *DeleteHumanLoopInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteHumanLoopInput"}
	if v.HumanLoopName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeHumanLoopInput(v *DescribeHumanLoopInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeHumanLoopInput"}
	if v.HumanLoopName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListHumanLoopsInput(v *ListHumanLoopsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListHumanLoopsInput"}
	if v.FlowDefinitionArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("FlowDefinitionArn"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpStartHumanLoopInput(v *StartHumanLoopInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartHumanLoopInput"}
	if v.HumanLoopName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if v.FlowDefinitionArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("FlowDefinitionArn"))
	}
	if v.HumanLoopInput == nil {
	invalidParams.Add(smithy.NewErrParamRequired("HumanLoopInput"))
	} else if v.HumanLoopInput != nil {
		if err := validateHumanLoopInput(v.HumanLoopInput); err != nil {
			invalidParams.AddNested("HumanLoopInput", err.(smithy.InvalidParamsError))
		}
	}
	if v.DataAttributes != nil {
		if err := validateHumanLoopDataAttributes(v.DataAttributes); err != nil {
			invalidParams.AddNested("DataAttributes", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpStopHumanLoopInput(v *StopHumanLoopInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopHumanLoopInput"}
	if v.HumanLoopName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("HumanLoopName"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
