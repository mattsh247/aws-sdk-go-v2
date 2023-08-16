// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type ItemType string

// Enum values for ItemType
const (
	ItemTypeObject ItemType = "OBJECT"
	ItemTypeFolder ItemType = "FOLDER"
)

// Values returns all known values for ItemType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ItemType) Values() []ItemType {
	return []ItemType{
		"OBJECT",
		"FOLDER",
	}
}

type StorageClass string

// Enum values for StorageClass
const (
	StorageClassTemporal StorageClass = "TEMPORAL"
)

// Values returns all known values for StorageClass. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StorageClass) Values() []StorageClass {
	return []StorageClass{
		"TEMPORAL",
	}
}

type UploadAvailability string

// Enum values for UploadAvailability
const (
	UploadAvailabilityStandard UploadAvailability = "STANDARD"
	UploadAvailabilityStreaming UploadAvailability = "STREAMING"
)

// Values returns all known values for UploadAvailability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UploadAvailability) Values() []UploadAvailability {
	return []UploadAvailability{
		"STANDARD",
		"STREAMING",
	}
}
