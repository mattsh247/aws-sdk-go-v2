// Code generated by smithy-go-codegen DO NOT EDIT.


package kendraranking

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/kendraranking/types"
)

type validateOpCreateRescoreExecutionPlan struct {
}

func (*validateOpCreateRescoreExecutionPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRescoreExecutionPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRescoreExecutionPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRescoreExecutionPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRescoreExecutionPlan struct {
}

func (*validateOpDeleteRescoreExecutionPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRescoreExecutionPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRescoreExecutionPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRescoreExecutionPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRescoreExecutionPlan struct {
}

func (*validateOpDescribeRescoreExecutionPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRescoreExecutionPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRescoreExecutionPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRescoreExecutionPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRescore struct {
}

func (*validateOpRescore) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRescore) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RescoreInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRescoreInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateRescoreExecutionPlan struct {
}

func (*validateOpUpdateRescoreExecutionPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateRescoreExecutionPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateRescoreExecutionPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateRescoreExecutionPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateRescoreExecutionPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRescoreExecutionPlan{}, middleware.After)
}

func addOpDeleteRescoreExecutionPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRescoreExecutionPlan{}, middleware.After)
}

func addOpDescribeRescoreExecutionPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRescoreExecutionPlan{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRescoreValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRescore{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateRescoreExecutionPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateRescoreExecutionPlan{}, middleware.After)
}

func validateCapacityUnitsConfiguration(v *types.CapacityUnitsConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapacityUnitsConfiguration"}
	if v.RescoreCapacityUnits == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RescoreCapacityUnits"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateDocument(v *types.Document) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Document"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.OriginalScore == nil {
	invalidParams.Add(smithy.NewErrParamRequired("OriginalScore"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateDocumentList(v []types.Document) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DocumentList"}
	for i := range v {
		if err := validateDocument(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateRescoreExecutionPlanInput(v *CreateRescoreExecutionPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRescoreExecutionPlanInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.CapacityUnits != nil {
		if err := validateCapacityUnitsConfiguration(v.CapacityUnits); err != nil {
			invalidParams.AddNested("CapacityUnits", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteRescoreExecutionPlanInput(v *DeleteRescoreExecutionPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRescoreExecutionPlanInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeRescoreExecutionPlanInput(v *DescribeRescoreExecutionPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRescoreExecutionPlanInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpRescoreInput(v *RescoreInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RescoreInput"}
	if v.RescoreExecutionPlanId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RescoreExecutionPlanId"))
	}
	if v.SearchQuery == nil {
	invalidParams.Add(smithy.NewErrParamRequired("SearchQuery"))
	}
	if v.Documents == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Documents"))
	} else if v.Documents != nil {
		if err := validateDocumentList(v.Documents); err != nil {
			invalidParams.AddNested("Documents", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateRescoreExecutionPlanInput(v *UpdateRescoreExecutionPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRescoreExecutionPlanInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.CapacityUnits != nil {
		if err := validateCapacityUnitsConfiguration(v.CapacityUnits); err != nil {
			invalidParams.AddNested("CapacityUnits", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
