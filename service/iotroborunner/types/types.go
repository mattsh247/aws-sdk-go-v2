// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Cartesian coordinates in 3D space relative to the RoboRunner origin.
type CartesianCoordinates struct {
	
	// X coordinate.
	//
	// This member is required.
	X *float64
	
	// Y coordinate.
	//
	// This member is required.
	Y *float64
	
	// Z coordinate.
	Z *float64
	
	noSmithyDocumentSerde
}

// Area within a facility where work can be performed.
type Destination struct {
	
	// Destination ARN.
	//
	// This member is required.
	Arn *string
	
	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// Filters access by the destination's identifier
	//
	// This member is required.
	Id *string
	
	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string
	
	// Site ARN.
	//
	// This member is required.
	Site *string
	
	// State of the destination.
	//
	// This member is required.
	State DestinationState
	
	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time
	
	// JSON document containing additional fixed properties regarding the destination
	AdditionalFixedProperties *string
	
	noSmithyDocumentSerde
}

// Worker orientation measured in units clockwise from north.
//
// The following types satisfy this interface:
//  OrientationMemberDegrees
type Orientation interface {
	isOrientation()
}

// Degrees, limited on [0, 360)
type OrientationMemberDegrees struct {
	Value float64
	
	noSmithyDocumentSerde
}
func (*OrientationMemberDegrees) isOrientation() {}

// Supported coordinates for worker position.
//
// The following types satisfy this interface:
//  PositionCoordinatesMemberCartesianCoordinates
type PositionCoordinates interface {
	isPositionCoordinates()
}

// Cartesian coordinates.
type PositionCoordinatesMemberCartesianCoordinates struct {
	Value CartesianCoordinates
	
	noSmithyDocumentSerde
}
func (*PositionCoordinatesMemberCartesianCoordinates) isPositionCoordinates() {}

// Facility containing destinations, workers, activities, and tasks.
type Site struct {
	
	// Site ARN.
	//
	// This member is required.
	Arn *string
	
	// A valid ISO 3166-1 alpha-2 code for the country in which the site resides.
	// e.g., US.
	//
	// This member is required.
	CountryCode *string
	
	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// The name of the site. Mutable after creation and unique within a given account.
	//
	// This member is required.
	Name *string
	
	noSmithyDocumentSerde
}

// Properties of the worker that are provided by the vendor FMS.
type VendorProperties struct {
	
	// The worker ID defined by the vendor FMS.
	//
	// This member is required.
	VendorWorkerId *string
	
	// JSON blob containing unstructured vendor properties that are fixed and won't
	// change during regular operation.
	VendorAdditionalFixedProperties *string
	
	// JSON blob containing unstructured vendor properties that are transient and may
	// change during regular operation.
	VendorAdditionalTransientProperties *string
	
	// The worker IP address defined by the vendor FMS.
	VendorWorkerIpAddress *string
	
	noSmithyDocumentSerde
}

// A unit capable of performing tasks.
type Worker struct {
	
	// Full ARN of the worker.
	//
	// This member is required.
	Arn *string
	
	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// Full ARN of the worker fleet.
	//
	// This member is required.
	Fleet *string
	
	// Filters access by the workers identifier
	//
	// This member is required.
	Id *string
	
	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string
	
	// Site ARN.
	//
	// This member is required.
	Site *string
	
	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time
	
	// JSON blob containing unstructured worker properties that are fixed and won't
	// change during regular operation.
	AdditionalFixedProperties *string
	
	// JSON blob containing unstructured worker properties that are transient and may
	// change during regular operation.
	AdditionalTransientProperties *string
	
	// Worker orientation measured in units clockwise from north.
	Orientation Orientation
	
	// Supported coordinates for worker position.
	Position PositionCoordinates
	
	// Properties of the worker that are provided by the vendor FMS.
	VendorProperties *VendorProperties
	
	noSmithyDocumentSerde
}

// A collection of workers organized within a facility.
type WorkerFleet struct {
	
	// Full ARN of the worker fleet.
	//
	// This member is required.
	Arn *string
	
	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time
	
	// Filters access by the worker fleet's identifier
	//
	// This member is required.
	Id *string
	
	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string
	
	// Site ARN.
	//
	// This member is required.
	Site *string
	
	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time
	
	// JSON blob containing additional fixed properties regarding the worker fleet
	AdditionalFixedProperties *string
	
	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag string
	Value []byte
	
	noSmithyDocumentSerde
}
func (*UnknownUnionMember) isOrientation() {}
func (*UnknownUnionMember) isPositionCoordinates() {}
