// Code generated by smithy-go-codegen DO NOT EDIT.


package honeycode

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
)

type validateOpBatchCreateTableRows struct {
}

func (*validateOpBatchCreateTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchCreateTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchCreateTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchCreateTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchDeleteTableRows struct {
}

func (*validateOpBatchDeleteTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchDeleteTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchDeleteTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchDeleteTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchUpdateTableRows struct {
}

func (*validateOpBatchUpdateTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchUpdateTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchUpdateTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchUpdateTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchUpsertTableRows struct {
}

func (*validateOpBatchUpsertTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchUpsertTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchUpsertTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchUpsertTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTableDataImportJob struct {
}

func (*validateOpDescribeTableDataImportJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTableDataImportJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTableDataImportJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTableDataImportJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetScreenData struct {
}

func (*validateOpGetScreenData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetScreenData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetScreenDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetScreenDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpInvokeScreenAutomation struct {
}

func (*validateOpInvokeScreenAutomation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpInvokeScreenAutomation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*InvokeScreenAutomationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpInvokeScreenAutomationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTableColumns struct {
}

func (*validateOpListTableColumns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTableColumns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTableColumnsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTableColumnsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTableRows struct {
}

func (*validateOpListTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTables struct {
}

func (*validateOpListTables) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTables) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTablesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTablesInput(input); err != nil {
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

type validateOpQueryTableRows struct {
}

func (*validateOpQueryTableRows) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpQueryTableRows) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*QueryTableRowsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpQueryTableRowsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartTableDataImportJob struct {
}

func (*validateOpStartTableDataImportJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartTableDataImportJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartTableDataImportJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartTableDataImportJobInput(input); err != nil {
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

func addOpBatchCreateTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchCreateTableRows{}, middleware.After)
}

func addOpBatchDeleteTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchDeleteTableRows{}, middleware.After)
}

func addOpBatchUpdateTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchUpdateTableRows{}, middleware.After)
}

func addOpBatchUpsertTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchUpsertTableRows{}, middleware.After)
}

func addOpDescribeTableDataImportJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTableDataImportJob{}, middleware.After)
}

func addOpGetScreenDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetScreenData{}, middleware.After)
}

func addOpInvokeScreenAutomationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpInvokeScreenAutomation{}, middleware.After)
}

func addOpListTableColumnsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTableColumns{}, middleware.After)
}

func addOpListTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTableRows{}, middleware.After)
}

func addOpListTablesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTables{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpQueryTableRowsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpQueryTableRows{}, middleware.After)
}

func addOpStartTableDataImportJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartTableDataImportJob{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateCreateRowData(v *types.CreateRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRowData"}
	if v.BatchItemId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("BatchItemId"))
	}
	if v.CellsToCreate == nil {
	invalidParams.Add(smithy.NewErrParamRequired("CellsToCreate"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateCreateRowDataList(v []types.CreateRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRowDataList"}
	for i := range v {
		if err := validateCreateRowData(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateDelimitedTextImportOptions(v *types.DelimitedTextImportOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DelimitedTextImportOptions"}
	if v.Delimiter == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Delimiter"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateFilter(v *types.Filter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Filter"}
	if v.Formula == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Formula"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateImportDataSource(v *types.ImportDataSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportDataSource"}
	if v.DataSourceConfig == nil {
	invalidParams.Add(smithy.NewErrParamRequired("DataSourceConfig"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateImportOptions(v *types.ImportOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportOptions"}
	if v.DelimitedTextOptions != nil {
		if err := validateDelimitedTextImportOptions(v.DelimitedTextOptions); err != nil {
			invalidParams.AddNested("DelimitedTextOptions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateUpdateRowData(v *types.UpdateRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRowData"}
	if v.RowId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RowId"))
	}
	if v.CellsToUpdate == nil {
	invalidParams.Add(smithy.NewErrParamRequired("CellsToUpdate"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateUpdateRowDataList(v []types.UpdateRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRowDataList"}
	for i := range v {
		if err := validateUpdateRowData(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateUpsertRowData(v *types.UpsertRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpsertRowData"}
	if v.BatchItemId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("BatchItemId"))
	}
	if v.Filter == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Filter"))
	} else if v.Filter != nil {
		if err := validateFilter(v.Filter); err != nil {
			invalidParams.AddNested("Filter", err.(smithy.InvalidParamsError))
		}
	}
	if v.CellsToUpdate == nil {
	invalidParams.Add(smithy.NewErrParamRequired("CellsToUpdate"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateUpsertRowDataList(v []types.UpsertRowData) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpsertRowDataList"}
	for i := range v {
		if err := validateUpsertRowData(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateVariableValue(v *types.VariableValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableValue"}
	if v.RawValue == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RawValue"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateVariableValueMap(v map[string]types.VariableValue) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VariableValueMap"}
	for key := range v {
		value := v[key]
		if err := validateVariableValue(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpBatchCreateTableRowsInput(v *BatchCreateTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchCreateTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.RowsToCreate == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RowsToCreate"))
	} else if v.RowsToCreate != nil {
		if err := validateCreateRowDataList(v.RowsToCreate); err != nil {
			invalidParams.AddNested("RowsToCreate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpBatchDeleteTableRowsInput(v *BatchDeleteTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchDeleteTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.RowIds == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RowIds"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpBatchUpdateTableRowsInput(v *BatchUpdateTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchUpdateTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.RowsToUpdate == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RowsToUpdate"))
	} else if v.RowsToUpdate != nil {
		if err := validateUpdateRowDataList(v.RowsToUpdate); err != nil {
			invalidParams.AddNested("RowsToUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpBatchUpsertTableRowsInput(v *BatchUpsertTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchUpsertTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.RowsToUpsert == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RowsToUpsert"))
	} else if v.RowsToUpsert != nil {
		if err := validateUpsertRowDataList(v.RowsToUpsert); err != nil {
			invalidParams.AddNested("RowsToUpsert", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeTableDataImportJobInput(v *DescribeTableDataImportJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTableDataImportJobInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.JobId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("JobId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetScreenDataInput(v *GetScreenDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetScreenDataInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.AppId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.ScreenId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ScreenId"))
	}
	if v.Variables != nil {
		if err := validateVariableValueMap(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpInvokeScreenAutomationInput(v *InvokeScreenAutomationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InvokeScreenAutomationInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.AppId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("AppId"))
	}
	if v.ScreenId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ScreenId"))
	}
	if v.ScreenAutomationId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ScreenAutomationId"))
	}
	if v.Variables != nil {
		if err := validateVariableValueMap(v.Variables); err != nil {
			invalidParams.AddNested("Variables", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListTableColumnsInput(v *ListTableColumnsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTableColumnsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListTableRowsInput(v *ListTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListTablesInput(v *ListTablesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTablesInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
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
	if v.ResourceArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpQueryTableRowsInput(v *QueryTableRowsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "QueryTableRowsInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.TableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("TableId"))
	}
	if v.FilterFormula == nil {
	invalidParams.Add(smithy.NewErrParamRequired("FilterFormula"))
	} else if v.FilterFormula != nil {
		if err := validateFilter(v.FilterFormula); err != nil {
			invalidParams.AddNested("FilterFormula", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpStartTableDataImportJobInput(v *StartTableDataImportJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartTableDataImportJobInput"}
	if v.WorkbookId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("WorkbookId"))
	}
	if v.DataSource == nil {
	invalidParams.Add(smithy.NewErrParamRequired("DataSource"))
	} else if v.DataSource != nil {
		if err := validateImportDataSource(v.DataSource); err != nil {
			invalidParams.AddNested("DataSource", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.DataFormat) == 0 {
	invalidParams.Add(smithy.NewErrParamRequired("DataFormat"))
	}
	if v.DestinationTableId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("DestinationTableId"))
	}
	if v.ImportOptions == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ImportOptions"))
	} else if v.ImportOptions != nil {
		if err := validateImportOptions(v.ImportOptions); err != nil {
			invalidParams.AddNested("ImportOptions", err.(smithy.InvalidParamsError))
		}
	}
	if v.ClientRequestToken == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
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
	if v.ResourceArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Tags"))
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
	if v.ResourceArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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
