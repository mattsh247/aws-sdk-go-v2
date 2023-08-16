// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type DecimalReturnType string

// Enum values for DecimalReturnType
const (
	DecimalReturnTypeString DecimalReturnType = "STRING"
	DecimalReturnTypeDoubleOrLong DecimalReturnType = "DOUBLE_OR_LONG"
)

// Values returns all known values for DecimalReturnType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DecimalReturnType) Values() []DecimalReturnType {
	return []DecimalReturnType{
		"STRING",
		"DOUBLE_OR_LONG",
	}
}

type LongReturnType string

// Enum values for LongReturnType
const (
	LongReturnTypeString LongReturnType = "STRING"
	LongReturnTypeLong LongReturnType = "LONG"
)

// Values returns all known values for LongReturnType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LongReturnType) Values() []LongReturnType {
	return []LongReturnType{
		"STRING",
		"LONG",
	}
}

type RecordsFormatType string

// Enum values for RecordsFormatType
const (
	RecordsFormatTypeNone RecordsFormatType = "NONE"
	RecordsFormatTypeJson RecordsFormatType = "JSON"
)

// Values returns all known values for RecordsFormatType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RecordsFormatType) Values() []RecordsFormatType {
	return []RecordsFormatType{
		"NONE",
		"JSON",
	}
}

type TypeHint string

// Enum values for TypeHint
const (
	TypeHintJson TypeHint = "JSON"
	TypeHintUuid TypeHint = "UUID"
	TypeHintTimestamp TypeHint = "TIMESTAMP"
	TypeHintDate TypeHint = "DATE"
	TypeHintTime TypeHint = "TIME"
	TypeHintDecimal TypeHint = "DECIMAL"
)

// Values returns all known values for TypeHint. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TypeHint) Values() []TypeHint {
	return []TypeHint{
		"JSON",
		"UUID",
		"TIMESTAMP",
		"DATE",
		"TIME",
		"DECIMAL",
	}
}
