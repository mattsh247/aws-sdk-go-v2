// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type Service string

// Enum values for Service
const (
	ServiceTurn Service = "TURN"
)

// Values returns all known values for Service. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Service) Values() []Service {
	return []Service{
		"TURN",
	}
}
