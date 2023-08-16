// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type BadRequestReason string

// Enum values for BadRequestReason
const (
	// Indicates there was a problem with one or more of the parameters. See
	// InvalidParameters in the BadRequestDetails for more information.
	BadRequestReasonInvalidParameters BadRequestReason = "InvalidParameters"
)

// Values returns all known values for BadRequestReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BadRequestReason) Values() []BadRequestReason {
	return []BadRequestReason{
		"InvalidParameters",
	}
}

type InvalidParameterProblem string

// Enum values for InvalidParameterProblem
const (
	// The parameter was corrupted and could not be understood by the service.
	InvalidParameterProblemCorrupted InvalidParameterProblem = "Corrupted"
	// The parameter was expired and can no longer be used.
	InvalidParameterProblemExpired InvalidParameterProblem = "Expired"
	// The client called the service before the time specified in the poll interval.
	InvalidParameterProblemPollIntervalNotSatisfied InvalidParameterProblem = "PollIntervalNotSatisfied"
)

// Values returns all known values for InvalidParameterProblem. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InvalidParameterProblem) Values() []InvalidParameterProblem {
	return []InvalidParameterProblem{
		"Corrupted",
		"Expired",
		"PollIntervalNotSatisfied",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	// Resource type value for the Application resource.
	ResourceTypeApplication ResourceType = "Application"
	// Resource type value for the ConfigurationProfile resource.
	ResourceTypeConfigurationProfile ResourceType = "ConfigurationProfile"
	// Resource type value for the Deployment resource.
	ResourceTypeDeployment ResourceType = "Deployment"
	// Resource type value for the Environment resource.
	ResourceTypeEnvironment ResourceType = "Environment"
	// Resource type value for the Configuration resource.
	ResourceTypeConfiguration ResourceType = "Configuration"
)

// Values returns all known values for ResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"Application",
		"ConfigurationProfile",
		"Deployment",
		"Environment",
		"Configuration",
	}
}
