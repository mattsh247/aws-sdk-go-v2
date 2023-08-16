// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An error that occured when putting the metric data.
type BatchPutMetricsError struct {

	// The error code of an error that occured when attempting to put metrics.
	//   - METRIC_LIMIT_EXCEEDED : The maximum amount of metrics per resource is
	//   exceeded.
	//   - INTERNAL_ERROR : An internal error occured.
	//   - VALIDATION_ERROR : The metric data failed validation.
	//   - CONFLICT_ERROR : Multiple requests attempted to modify the same data
	//   simultaneously.
	Code PutMetricsErrorCode

	// An index that corresponds to the metric in the request.
	MetricIndex int32

	noSmithyDocumentSerde
}

// The raw metric data to associate with the resource.
type RawMetricData struct {

	// The name of the metric.
	//
	// This member is required.
	MetricName *string

	// The time that the metric was recorded.
	//
	// This member is required.
	Timestamp *time.Time

	// The metric value.
	//
	// This member is required.
	Value float64

	// The metric step (epoch).
	Step *int32

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
