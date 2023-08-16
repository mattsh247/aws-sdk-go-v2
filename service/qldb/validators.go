// Code generated by smithy-go-codegen DO NOT EDIT.


package qldb

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
)

type validateOpCancelJournalKinesisStream struct {
}

func (*validateOpCancelJournalKinesisStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelJournalKinesisStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelJournalKinesisStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelJournalKinesisStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLedger struct {
}

func (*validateOpCreateLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLedgerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLedger struct {
}

func (*validateOpDeleteLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLedgerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJournalKinesisStream struct {
}

func (*validateOpDescribeJournalKinesisStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJournalKinesisStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJournalKinesisStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJournalKinesisStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeJournalS3Export struct {
}

func (*validateOpDescribeJournalS3Export) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeJournalS3Export) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeJournalS3ExportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeJournalS3ExportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLedger struct {
}

func (*validateOpDescribeLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLedgerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportJournalToS3 struct {
}

func (*validateOpExportJournalToS3) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportJournalToS3) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportJournalToS3Input)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportJournalToS3Input(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBlock struct {
}

func (*validateOpGetBlock) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBlock) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBlockInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBlockInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDigest struct {
}

func (*validateOpGetDigest) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDigest) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDigestInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDigestInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRevision struct {
}

func (*validateOpGetRevision) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRevision) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRevisionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRevisionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListJournalKinesisStreamsForLedger struct {
}

func (*validateOpListJournalKinesisStreamsForLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListJournalKinesisStreamsForLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListJournalKinesisStreamsForLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListJournalKinesisStreamsForLedgerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListJournalS3ExportsForLedger struct {
}

func (*validateOpListJournalS3ExportsForLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListJournalS3ExportsForLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListJournalS3ExportsForLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListJournalS3ExportsForLedgerInput(input); err != nil {
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

type validateOpStreamJournalToKinesis struct {
}

func (*validateOpStreamJournalToKinesis) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStreamJournalToKinesis) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StreamJournalToKinesisInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStreamJournalToKinesisInput(input); err != nil {
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

type validateOpUpdateLedger struct {
}

func (*validateOpUpdateLedger) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLedger) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLedgerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLedgerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateLedgerPermissionsMode struct {
}

func (*validateOpUpdateLedgerPermissionsMode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLedgerPermissionsMode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLedgerPermissionsModeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLedgerPermissionsModeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelJournalKinesisStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelJournalKinesisStream{}, middleware.After)
}

func addOpCreateLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLedger{}, middleware.After)
}

func addOpDeleteLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLedger{}, middleware.After)
}

func addOpDescribeJournalKinesisStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJournalKinesisStream{}, middleware.After)
}

func addOpDescribeJournalS3ExportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeJournalS3Export{}, middleware.After)
}

func addOpDescribeLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLedger{}, middleware.After)
}

func addOpExportJournalToS3ValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportJournalToS3{}, middleware.After)
}

func addOpGetBlockValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBlock{}, middleware.After)
}

func addOpGetDigestValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDigest{}, middleware.After)
}

func addOpGetRevisionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRevision{}, middleware.After)
}

func addOpListJournalKinesisStreamsForLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListJournalKinesisStreamsForLedger{}, middleware.After)
}

func addOpListJournalS3ExportsForLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListJournalS3ExportsForLedger{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStreamJournalToKinesisValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStreamJournalToKinesis{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLedgerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLedger{}, middleware.After)
}

func addOpUpdateLedgerPermissionsModeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLedgerPermissionsMode{}, middleware.After)
}

func validateKinesisConfiguration(v *types.KinesisConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KinesisConfiguration"}
	if v.StreamArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("StreamArn"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateS3EncryptionConfiguration(v *types.S3EncryptionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3EncryptionConfiguration"}
	if len(v.ObjectEncryptionType) == 0 {
	invalidParams.Add(smithy.NewErrParamRequired("ObjectEncryptionType"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateS3ExportConfiguration(v *types.S3ExportConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3ExportConfiguration"}
	if v.Bucket == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Bucket"))
	}
	if v.Prefix == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Prefix"))
	}
	if v.EncryptionConfiguration == nil {
	invalidParams.Add(smithy.NewErrParamRequired("EncryptionConfiguration"))
	} else if v.EncryptionConfiguration != nil {
		if err := validateS3EncryptionConfiguration(v.EncryptionConfiguration); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCancelJournalKinesisStreamInput(v *CancelJournalKinesisStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelJournalKinesisStreamInput"}
	if v.LedgerName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("LedgerName"))
	}
	if v.StreamId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("StreamId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateLedgerInput(v *CreateLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLedgerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.PermissionsMode) == 0 {
	invalidParams.Add(smithy.NewErrParamRequired("PermissionsMode"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteLedgerInput(v *DeleteLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLedgerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeJournalKinesisStreamInput(v *DescribeJournalKinesisStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJournalKinesisStreamInput"}
	if v.LedgerName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("LedgerName"))
	}
	if v.StreamId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("StreamId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeJournalS3ExportInput(v *DescribeJournalS3ExportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeJournalS3ExportInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ExportId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ExportId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDescribeLedgerInput(v *DescribeLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLedgerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpExportJournalToS3Input(v *ExportJournalToS3Input) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportJournalToS3Input"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.InclusiveStartTime == nil {
	invalidParams.Add(smithy.NewErrParamRequired("InclusiveStartTime"))
	}
	if v.ExclusiveEndTime == nil {
	invalidParams.Add(smithy.NewErrParamRequired("ExclusiveEndTime"))
	}
	if v.S3ExportConfiguration == nil {
	invalidParams.Add(smithy.NewErrParamRequired("S3ExportConfiguration"))
	} else if v.S3ExportConfiguration != nil {
		if err := validateS3ExportConfiguration(v.S3ExportConfiguration); err != nil {
			invalidParams.AddNested("S3ExportConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.RoleArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetBlockInput(v *GetBlockInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBlockInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.BlockAddress == nil {
	invalidParams.Add(smithy.NewErrParamRequired("BlockAddress"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetDigestInput(v *GetDigestInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDigestInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetRevisionInput(v *GetRevisionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRevisionInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.BlockAddress == nil {
	invalidParams.Add(smithy.NewErrParamRequired("BlockAddress"))
	}
	if v.DocumentId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("DocumentId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListJournalKinesisStreamsForLedgerInput(v *ListJournalKinesisStreamsForLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListJournalKinesisStreamsForLedgerInput"}
	if v.LedgerName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("LedgerName"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListJournalS3ExportsForLedgerInput(v *ListJournalS3ExportsForLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListJournalS3ExportsForLedgerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
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

func validateOpStreamJournalToKinesisInput(v *StreamJournalToKinesisInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StreamJournalToKinesisInput"}
	if v.LedgerName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("LedgerName"))
	}
	if v.RoleArn == nil {
	invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.InclusiveStartTime == nil {
	invalidParams.Add(smithy.NewErrParamRequired("InclusiveStartTime"))
	}
	if v.KinesisConfiguration == nil {
	invalidParams.Add(smithy.NewErrParamRequired("KinesisConfiguration"))
	} else if v.KinesisConfiguration != nil {
		if err := validateKinesisConfiguration(v.KinesisConfiguration); err != nil {
			invalidParams.AddNested("KinesisConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.StreamName == nil {
	invalidParams.Add(smithy.NewErrParamRequired("StreamName"))
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

func validateOpUpdateLedgerInput(v *UpdateLedgerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLedgerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateLedgerPermissionsModeInput(v *UpdateLedgerPermissionsModeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLedgerPermissionsModeInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.PermissionsMode) == 0 {
	invalidParams.Add(smithy.NewErrParamRequired("PermissionsMode"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
