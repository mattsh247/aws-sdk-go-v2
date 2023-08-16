// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// The forecast value for a specific date. Part of the Forecast object.
type DataPoint struct {
	
	// The timestamp of the specific forecast.
	Timestamp *string
	
	// The forecast value.
	Value *float64
	
	noSmithyDocumentSerde
}

// Provides information about a forecast. Returned as part of the QueryForecast
// response.
type Forecast struct {
	
	// The forecast. The string of the string-to-array map is one of the following
	// values:
	//   - p10
	//   - p50
	//   - p90
	// The default setting is ["0.1", "0.5", "0.9"] . Use the optional ForecastTypes
	// parameter of the CreateForecast (https://docs.aws.amazon.com/forecast/latest/dg/API_CreateForecast.html)
	// operation to change the values. The values will vary depending on how this is
	// set, with a minimum of 1 and a maximum of 5.
	Predictions map[string][]DataPoint
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
