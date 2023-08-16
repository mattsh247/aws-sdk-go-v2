// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type ActionStatus string

// Enum values for ActionStatus
const (
	ActionStatusStandby ActionStatus = "STANDBY"
	ActionStatusPending ActionStatus = "PENDING"
	ActionStatusExecutionInProgress ActionStatus = "EXECUTION_IN_PROGRESS"
	ActionStatusExecutionSuccess ActionStatus = "EXECUTION_SUCCESS"
	ActionStatusExecutionFailure ActionStatus = "EXECUTION_FAILURE"
	ActionStatusReverseInProgress ActionStatus = "REVERSE_IN_PROGRESS"
	ActionStatusReverseSuccess ActionStatus = "REVERSE_SUCCESS"
	ActionStatusReverseFailure ActionStatus = "REVERSE_FAILURE"
	ActionStatusResetInProgress ActionStatus = "RESET_IN_PROGRESS"
	ActionStatusResetFailure ActionStatus = "RESET_FAILURE"
)

// Values returns all known values for ActionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionStatus) Values() []ActionStatus {
	return []ActionStatus{
		"STANDBY",
		"PENDING",
		"EXECUTION_IN_PROGRESS",
		"EXECUTION_SUCCESS",
		"EXECUTION_FAILURE",
		"REVERSE_IN_PROGRESS",
		"REVERSE_SUCCESS",
		"REVERSE_FAILURE",
		"RESET_IN_PROGRESS",
		"RESET_FAILURE",
	}
}

type ActionSubType string

// Enum values for ActionSubType
const (
	ActionSubTypeStopEc2 ActionSubType = "STOP_EC2_INSTANCES"
	ActionSubTypeStopRds ActionSubType = "STOP_RDS_INSTANCES"
)

// Values returns all known values for ActionSubType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionSubType) Values() []ActionSubType {
	return []ActionSubType{
		"STOP_EC2_INSTANCES",
		"STOP_RDS_INSTANCES",
	}
}

type ActionType string

// Enum values for ActionType
const (
	ActionTypeIam ActionType = "APPLY_IAM_POLICY"
	ActionTypeScp ActionType = "APPLY_SCP_POLICY"
	ActionTypeSsm ActionType = "RUN_SSM_DOCUMENTS"
)

// Values returns all known values for ActionType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionType) Values() []ActionType {
	return []ActionType{
		"APPLY_IAM_POLICY",
		"APPLY_SCP_POLICY",
		"RUN_SSM_DOCUMENTS",
	}
}

type ApprovalModel string

// Enum values for ApprovalModel
const (
	ApprovalModelAuto ApprovalModel = "AUTOMATIC"
	ApprovalModelManual ApprovalModel = "MANUAL"
)

// Values returns all known values for ApprovalModel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApprovalModel) Values() []ApprovalModel {
	return []ApprovalModel{
		"AUTOMATIC",
		"MANUAL",
	}
}

type AutoAdjustType string

// Enum values for AutoAdjustType
const (
	AutoAdjustTypeHistorical AutoAdjustType = "HISTORICAL"
	AutoAdjustTypeForecast AutoAdjustType = "FORECAST"
)

// Values returns all known values for AutoAdjustType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoAdjustType) Values() []AutoAdjustType {
	return []AutoAdjustType{
		"HISTORICAL",
		"FORECAST",
	}
}

type BudgetType string

// Enum values for BudgetType
const (
	BudgetTypeUsage BudgetType = "USAGE"
	BudgetTypeCost BudgetType = "COST"
	BudgetTypeRIUtilization BudgetType = "RI_UTILIZATION"
	BudgetTypeRICoverage BudgetType = "RI_COVERAGE"
	BudgetTypeSPUtilization BudgetType = "SAVINGS_PLANS_UTILIZATION"
	BudgetTypeSPCoverage BudgetType = "SAVINGS_PLANS_COVERAGE"
)

// Values returns all known values for BudgetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BudgetType) Values() []BudgetType {
	return []BudgetType{
		"USAGE",
		"COST",
		"RI_UTILIZATION",
		"RI_COVERAGE",
		"SAVINGS_PLANS_UTILIZATION",
		"SAVINGS_PLANS_COVERAGE",
	}
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorGreaterThan ComparisonOperator = "GREATER_THAN"
	ComparisonOperatorLessThan ComparisonOperator = "LESS_THAN"
	ComparisonOperatorEqualTo ComparisonOperator = "EQUAL_TO"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"GREATER_THAN",
		"LESS_THAN",
		"EQUAL_TO",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeSystem EventType = "SYSTEM"
	EventTypeCreateAction EventType = "CREATE_ACTION"
	EventTypeDeleteAction EventType = "DELETE_ACTION"
	EventTypeUpdateAction EventType = "UPDATE_ACTION"
	EventTypeExecuteAction EventType = "EXECUTE_ACTION"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"SYSTEM",
		"CREATE_ACTION",
		"DELETE_ACTION",
		"UPDATE_ACTION",
		"EXECUTE_ACTION",
	}
}

type ExecutionType string

// Enum values for ExecutionType
const (
	ExecutionTypeApproveBudgetAction ExecutionType = "APPROVE_BUDGET_ACTION"
	ExecutionTypeRetryBudgetAction ExecutionType = "RETRY_BUDGET_ACTION"
	ExecutionTypeReverseBudgetAction ExecutionType = "REVERSE_BUDGET_ACTION"
	ExecutionTypeResetBudgetAction ExecutionType = "RESET_BUDGET_ACTION"
)

// Values returns all known values for ExecutionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutionType) Values() []ExecutionType {
	return []ExecutionType{
		"APPROVE_BUDGET_ACTION",
		"RETRY_BUDGET_ACTION",
		"REVERSE_BUDGET_ACTION",
		"RESET_BUDGET_ACTION",
	}
}

type NotificationState string

// Enum values for NotificationState
const (
	NotificationStateOk NotificationState = "OK"
	NotificationStateAlarm NotificationState = "ALARM"
)

// Values returns all known values for NotificationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationState) Values() []NotificationState {
	return []NotificationState{
		"OK",
		"ALARM",
	}
}

type NotificationType string

// Enum values for NotificationType
const (
	NotificationTypeActual NotificationType = "ACTUAL"
	NotificationTypeForecasted NotificationType = "FORECASTED"
)

// Values returns all known values for NotificationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationType) Values() []NotificationType {
	return []NotificationType{
		"ACTUAL",
		"FORECASTED",
	}
}

type SubscriptionType string

// Enum values for SubscriptionType
const (
	SubscriptionTypeSns SubscriptionType = "SNS"
	SubscriptionTypeEmail SubscriptionType = "EMAIL"
)

// Values returns all known values for SubscriptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SubscriptionType) Values() []SubscriptionType {
	return []SubscriptionType{
		"SNS",
		"EMAIL",
	}
}

type ThresholdType string

// Enum values for ThresholdType
const (
	ThresholdTypePercentage ThresholdType = "PERCENTAGE"
	ThresholdTypeAbsoluteValue ThresholdType = "ABSOLUTE_VALUE"
)

// Values returns all known values for ThresholdType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThresholdType) Values() []ThresholdType {
	return []ThresholdType{
		"PERCENTAGE",
		"ABSOLUTE_VALUE",
	}
}

type TimeUnit string

// Enum values for TimeUnit
const (
	TimeUnitDaily TimeUnit = "DAILY"
	TimeUnitMonthly TimeUnit = "MONTHLY"
	TimeUnitQuarterly TimeUnit = "QUARTERLY"
	TimeUnitAnnually TimeUnit = "ANNUALLY"
)

// Values returns all known values for TimeUnit. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TimeUnit) Values() []TimeUnit {
	return []TimeUnit{
		"DAILY",
		"MONTHLY",
		"QUARTERLY",
		"ANNUALLY",
	}
}
