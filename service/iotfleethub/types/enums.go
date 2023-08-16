// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type ApplicationState string

// Enum values for ApplicationState
const (
	ApplicationStateCreating ApplicationState = "CREATING"
	ApplicationStateDeleting ApplicationState = "DELETING"
	ApplicationStateActive ApplicationState = "ACTIVE"
	ApplicationStateCreateFailed ApplicationState = "CREATE_FAILED"
	ApplicationStateDeleteFailed ApplicationState = "DELETE_FAILED"
)

// Values returns all known values for ApplicationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationState) Values() []ApplicationState {
	return []ApplicationState{
		"CREATING",
		"DELETING",
		"ACTIVE",
		"CREATE_FAILED",
		"DELETE_FAILED",
	}
}
