// Code generated by smithy-go-codegen DO NOT EDIT.


package serverlessapplicationrepository

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
)

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateApplicationVersion struct {
}

func (*validateOpCreateApplicationVersion) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplicationVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationVersionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCloudFormationChangeSet struct {
}

func (*validateOpCreateCloudFormationChangeSet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCloudFormationChangeSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCloudFormationChangeSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCloudFormationChangeSetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCloudFormationTemplate struct {
}

func (*validateOpCreateCloudFormationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCloudFormationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCloudFormationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCloudFormationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplication struct {
}

func (*validateOpDeleteApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetApplication struct {
}

func (*validateOpGetApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetApplicationPolicy struct {
}

func (*validateOpGetApplicationPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetApplicationPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetApplicationPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetApplicationPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCloudFormationTemplate struct {
}

func (*validateOpGetCloudFormationTemplate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCloudFormationTemplate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCloudFormationTemplateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCloudFormationTemplateInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListApplicationDependencies struct {
}

func (*validateOpListApplicationDependencies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListApplicationDependencies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListApplicationDependenciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListApplicationDependenciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListApplicationVersions struct {
}

func (*validateOpListApplicationVersions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListApplicationVersions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListApplicationVersionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListApplicationVersionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutApplicationPolicy struct {
}

func (*validateOpPutApplicationPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutApplicationPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutApplicationPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutApplicationPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUnshareApplication struct {
}

func (*validateOpUnshareApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUnshareApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UnshareApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUnshareApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplication struct {
}

func (*validateOpUpdateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplication{}, middleware.After)
}

func addOpCreateApplicationVersionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplicationVersion{}, middleware.After)
}

func addOpCreateCloudFormationChangeSetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCloudFormationChangeSet{}, middleware.After)
}

func addOpCreateCloudFormationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCloudFormationTemplate{}, middleware.After)
}

func addOpDeleteApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplication{}, middleware.After)
}

func addOpGetApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetApplication{}, middleware.After)
}

func addOpGetApplicationPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetApplicationPolicy{}, middleware.After)
}

func addOpGetCloudFormationTemplateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCloudFormationTemplate{}, middleware.After)
}

func addOpListApplicationDependenciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListApplicationDependencies{}, middleware.After)
}

func addOpListApplicationVersionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListApplicationVersions{}, middleware.After)
}

func addOpPutApplicationPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutApplicationPolicy{}, middleware.After)
}

func addOpUnshareApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUnshareApplication{}, middleware.After)
}

func addOpUpdateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplication{}, middleware.After)
}

func validate__listOfApplicationPolicyStatement(v []types.ApplicationPolicyStatement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfApplicationPolicyStatement"}
	for i := range v {
		if err := validateApplicationPolicyStatement(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validate__listOfParameterValue(v []types.ParameterValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfParameterValue"}
	for i := range v {
		if err := validateParameterValue(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validate__listOfRollbackTrigger(v []types.RollbackTrigger) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfRollbackTrigger"}
	for i := range v {
		if err := validateRollbackTrigger(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validate__listOfTag(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOfTag"}
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

func validateApplicationPolicyStatement(v *types.ApplicationPolicyStatement) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ApplicationPolicyStatement"}
	if v.Actions == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Actions"))
	}
	if v.Principals == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Principals"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateParameterValue(v *types.ParameterValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ParameterValue"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
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

func validateRollbackConfiguration(v *types.RollbackConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RollbackConfiguration"}
	if v.RollbackTriggers != nil {
		if err := validate__listOfRollbackTrigger(v.RollbackTriggers); err != nil {
			invalidParams.AddNested("RollbackTriggers", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateRollbackTrigger(v *types.RollbackTrigger) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RollbackTrigger"}
	if v.Arn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.Type == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Type"))
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

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.Author == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Author"))
	}
	if v.Description == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateApplicationVersionInput(v *CreateApplicationVersionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationVersionInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.SemanticVersion == nil {
	invalidParams.Add(smithy.NewErrParamRequired("SemanticVersion"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateCloudFormationChangeSetInput(v *CreateCloudFormationChangeSetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCloudFormationChangeSetInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.ParameterOverrides != nil {
		if err := validate__listOfParameterValue(v.ParameterOverrides); err != nil {
			invalidParams.AddNested("ParameterOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if v.RollbackConfiguration != nil {
		if err := validateRollbackConfiguration(v.RollbackConfiguration); err != nil {
			invalidParams.AddNested("RollbackConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.StackName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("StackName"))
	}
	if v.Tags != nil {
		if err := validate__listOfTag(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateCloudFormationTemplateInput(v *CreateCloudFormationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCloudFormationTemplateInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteApplicationInput(v *DeleteApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetApplicationInput(v *GetApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetApplicationInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetApplicationPolicyInput(v *GetApplicationPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetApplicationPolicyInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetCloudFormationTemplateInput(v *GetCloudFormationTemplateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCloudFormationTemplateInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.TemplateId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TemplateId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListApplicationDependenciesInput(v *ListApplicationDependenciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListApplicationDependenciesInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListApplicationVersionsInput(v *ListApplicationVersionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListApplicationVersionsInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpPutApplicationPolicyInput(v *PutApplicationPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutApplicationPolicyInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.Statements == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Statements"))
	} else if v.Statements != nil {
		if err := validate__listOfApplicationPolicyStatement(v.Statements); err != nil {
			invalidParams.AddNested("Statements", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUnshareApplicationInput(v *UnshareApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UnshareApplicationInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.OrganizationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("OrganizationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateApplicationInput(v *UpdateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationInput"}
	if v.ApplicationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
