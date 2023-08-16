// Code generated by smithy-go-codegen DO NOT EDIT.


package iotroborunner

import (
	"context"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
)

type validateOpCreateDestination struct {
}

func (*validateOpCreateDestination) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDestinationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSite struct {
}

func (*validateOpCreateSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWorkerFleet struct {
}

func (*validateOpCreateWorkerFleet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWorkerFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWorkerFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWorkerFleetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWorker struct {
}

func (*validateOpCreateWorker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWorker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWorkerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWorkerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDestination struct {
}

func (*validateOpDeleteDestination) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDestinationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSite struct {
}

func (*validateOpDeleteSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWorkerFleet struct {
}

func (*validateOpDeleteWorkerFleet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWorkerFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWorkerFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWorkerFleetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWorker struct {
}

func (*validateOpDeleteWorker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWorker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWorkerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWorkerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDestination struct {
}

func (*validateOpGetDestination) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDestinationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSite struct {
}

func (*validateOpGetSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorkerFleet struct {
}

func (*validateOpGetWorkerFleet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorkerFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkerFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkerFleetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetWorker struct {
}

func (*validateOpGetWorker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetWorker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetWorkerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetWorkerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDestinations struct {
}

func (*validateOpListDestinations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDestinations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDestinationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDestinationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListWorkerFleets struct {
}

func (*validateOpListWorkerFleets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListWorkerFleets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListWorkerFleetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListWorkerFleetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListWorkers struct {
}

func (*validateOpListWorkers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListWorkers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListWorkersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListWorkersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDestination struct {
}

func (*validateOpUpdateDestination) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDestinationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSite struct {
}

func (*validateOpUpdateSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkerFleet struct {
}

func (*validateOpUpdateWorkerFleet) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkerFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkerFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkerFleetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorker struct {
}

func (*validateOpUpdateWorker) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorker) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateDestinationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDestination{}, middleware.After)
}

func addOpCreateSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSite{}, middleware.After)
}

func addOpCreateWorkerFleetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWorkerFleet{}, middleware.After)
}

func addOpCreateWorkerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWorker{}, middleware.After)
}

func addOpDeleteDestinationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDestination{}, middleware.After)
}

func addOpDeleteSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSite{}, middleware.After)
}

func addOpDeleteWorkerFleetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorkerFleet{}, middleware.After)
}

func addOpDeleteWorkerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorker{}, middleware.After)
}

func addOpGetDestinationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDestination{}, middleware.After)
}

func addOpGetSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSite{}, middleware.After)
}

func addOpGetWorkerFleetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorkerFleet{}, middleware.After)
}

func addOpGetWorkerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetWorker{}, middleware.After)
}

func addOpListDestinationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDestinations{}, middleware.After)
}

func addOpListWorkerFleetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListWorkerFleets{}, middleware.After)
}

func addOpListWorkersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListWorkers{}, middleware.After)
}

func addOpUpdateDestinationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDestination{}, middleware.After)
}

func addOpUpdateSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSite{}, middleware.After)
}

func addOpUpdateWorkerFleetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkerFleet{}, middleware.After)
}

func addOpUpdateWorkerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorker{}, middleware.After)
}

func validateCartesianCoordinates(v *types.CartesianCoordinates) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CartesianCoordinates"}
	if v.X == nil {
	invalidParams.Add(smithy.NewErrParamRequired("X"))
	}
	if v.Y == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Y"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validatePositionCoordinates(v types.PositionCoordinates) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PositionCoordinates"}
	switch uv := v.(type) {
		case *types.PositionCoordinatesMemberCartesianCoordinates:
			if err := validateCartesianCoordinates(&uv.Value); err != nil {
				invalidParams.AddNested("[cartesianCoordinates]", err.(smithy.InvalidParamsError))
			}
		
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateVendorProperties(v *types.VendorProperties) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VendorProperties"}
	if v.VendorWorkerId == nil {
	invalidParams.Add(smithy.NewErrParamRequired("VendorWorkerId"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateDestinationInput(v *CreateDestinationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDestinationInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Site == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Site"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateSiteInput(v *CreateSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSiteInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.CountryCode == nil {
	invalidParams.Add(smithy.NewErrParamRequired("CountryCode"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateWorkerFleetInput(v *CreateWorkerFleetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWorkerFleetInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Site == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Site"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpCreateWorkerInput(v *CreateWorkerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWorkerInput"}
	if v.Name == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Fleet == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Fleet"))
	}
	if v.VendorProperties != nil {
		if err := validateVendorProperties(v.VendorProperties); err != nil {
			invalidParams.AddNested("VendorProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Position != nil {
		if err := validatePositionCoordinates(v.Position); err != nil {
			invalidParams.AddNested("Position", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteDestinationInput(v *DeleteDestinationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDestinationInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteSiteInput(v *DeleteSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSiteInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteWorkerFleetInput(v *DeleteWorkerFleetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWorkerFleetInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpDeleteWorkerInput(v *DeleteWorkerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWorkerInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetDestinationInput(v *GetDestinationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDestinationInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetSiteInput(v *GetSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSiteInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetWorkerFleetInput(v *GetWorkerFleetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkerFleetInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpGetWorkerInput(v *GetWorkerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetWorkerInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListDestinationsInput(v *ListDestinationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDestinationsInput"}
	if v.Site == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Site"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListWorkerFleetsInput(v *ListWorkerFleetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListWorkerFleetsInput"}
	if v.Site == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Site"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpListWorkersInput(v *ListWorkersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListWorkersInput"}
	if v.Site == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Site"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateDestinationInput(v *UpdateDestinationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDestinationInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateSiteInput(v *UpdateSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSiteInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateWorkerFleetInput(v *UpdateWorkerFleetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkerFleetInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}

func validateOpUpdateWorkerInput(v *UpdateWorkerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkerInput"}
	if v.Id == nil {
	invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if v.VendorProperties != nil {
		if err := validateVendorProperties(v.VendorProperties); err != nil {
			invalidParams.AddNested("VendorProperties", err.(smithy.InvalidParamsError))
		}
	}
	if v.Position != nil {
		if err := validatePositionCoordinates(v.Position); err != nil {
			invalidParams.AddNested("Position", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
	return invalidParams
	} else {
	return nil
	}
}
