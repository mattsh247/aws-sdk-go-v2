// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AccessorStatus string

// Enum values for AccessorStatus
const (
	AccessorStatusAvailable AccessorStatus = "AVAILABLE"
	AccessorStatusPendingDeletion AccessorStatus = "PENDING_DELETION"
	AccessorStatusDeleted AccessorStatus = "DELETED"
)

// Values returns all known values for AccessorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessorStatus) Values() []AccessorStatus {
	return []AccessorStatus{
		"AVAILABLE",
		"PENDING_DELETION",
		"DELETED",
	}
}

type AccessorType string

// Enum values for AccessorType
const (
	AccessorTypeBillingToken AccessorType = "BILLING_TOKEN"
)

// Values returns all known values for AccessorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessorType) Values() []AccessorType {
	return []AccessorType{
		"BILLING_TOKEN",
	}
}

type Edition string

// Enum values for Edition
const (
	EditionStarter Edition = "STARTER"
	EditionStandard Edition = "STANDARD"
)

// Values returns all known values for Edition. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Edition) Values() []Edition {
	return []Edition{
		"STARTER",
		"STANDARD",
	}
}

type Framework string

// Enum values for Framework
const (
	FrameworkHyperledgerFabric Framework = "HYPERLEDGER_FABRIC"
	FrameworkEthereum Framework = "ETHEREUM"
)

// Values returns all known values for Framework. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Framework) Values() []Framework {
	return []Framework{
		"HYPERLEDGER_FABRIC",
		"ETHEREUM",
	}
}

type InvitationStatus string

// Enum values for InvitationStatus
const (
	InvitationStatusPending InvitationStatus = "PENDING"
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	InvitationStatusAccepting InvitationStatus = "ACCEPTING"
	InvitationStatusRejected InvitationStatus = "REJECTED"
	InvitationStatusExpired InvitationStatus = "EXPIRED"
)

// Values returns all known values for InvitationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InvitationStatus) Values() []InvitationStatus {
	return []InvitationStatus{
		"PENDING",
		"ACCEPTED",
		"ACCEPTING",
		"REJECTED",
		"EXPIRED",
	}
}

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusCreating MemberStatus = "CREATING"
	MemberStatusAvailable MemberStatus = "AVAILABLE"
	MemberStatusCreateFailed MemberStatus = "CREATE_FAILED"
	MemberStatusUpdating MemberStatus = "UPDATING"
	MemberStatusDeleting MemberStatus = "DELETING"
	MemberStatusDeleted MemberStatus = "DELETED"
	MemberStatusInaccessibleEncryptionKey MemberStatus = "INACCESSIBLE_ENCRYPTION_KEY"
)

// Values returns all known values for MemberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MemberStatus) Values() []MemberStatus {
	return []MemberStatus{
		"CREATING",
		"AVAILABLE",
		"CREATE_FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"INACCESSIBLE_ENCRYPTION_KEY",
	}
}

type NetworkStatus string

// Enum values for NetworkStatus
const (
	NetworkStatusCreating NetworkStatus = "CREATING"
	NetworkStatusAvailable NetworkStatus = "AVAILABLE"
	NetworkStatusCreateFailed NetworkStatus = "CREATE_FAILED"
	NetworkStatusDeleting NetworkStatus = "DELETING"
	NetworkStatusDeleted NetworkStatus = "DELETED"
)

// Values returns all known values for NetworkStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NetworkStatus) Values() []NetworkStatus {
	return []NetworkStatus{
		"CREATING",
		"AVAILABLE",
		"CREATE_FAILED",
		"DELETING",
		"DELETED",
	}
}

type NodeStatus string

// Enum values for NodeStatus
const (
	NodeStatusCreating NodeStatus = "CREATING"
	NodeStatusAvailable NodeStatus = "AVAILABLE"
	NodeStatusUnhealthy NodeStatus = "UNHEALTHY"
	NodeStatusCreateFailed NodeStatus = "CREATE_FAILED"
	NodeStatusUpdating NodeStatus = "UPDATING"
	NodeStatusDeleting NodeStatus = "DELETING"
	NodeStatusDeleted NodeStatus = "DELETED"
	NodeStatusFailed NodeStatus = "FAILED"
	NodeStatusInaccessibleEncryptionKey NodeStatus = "INACCESSIBLE_ENCRYPTION_KEY"
)

// Values returns all known values for NodeStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (NodeStatus) Values() []NodeStatus {
	return []NodeStatus{
		"CREATING",
		"AVAILABLE",
		"UNHEALTHY",
		"CREATE_FAILED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"INACCESSIBLE_ENCRYPTION_KEY",
	}
}

type ProposalStatus string

// Enum values for ProposalStatus
const (
	ProposalStatusInProgress ProposalStatus = "IN_PROGRESS"
	ProposalStatusApproved ProposalStatus = "APPROVED"
	ProposalStatusRejected ProposalStatus = "REJECTED"
	ProposalStatusExpired ProposalStatus = "EXPIRED"
	ProposalStatusActionFailed ProposalStatus = "ACTION_FAILED"
)

// Values returns all known values for ProposalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProposalStatus) Values() []ProposalStatus {
	return []ProposalStatus{
		"IN_PROGRESS",
		"APPROVED",
		"REJECTED",
		"EXPIRED",
		"ACTION_FAILED",
	}
}

type StateDBType string

// Enum values for StateDBType
const (
	StateDBTypeLevelDB StateDBType = "LevelDB"
	StateDBTypeCouchDB StateDBType = "CouchDB"
)

// Values returns all known values for StateDBType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StateDBType) Values() []StateDBType {
	return []StateDBType{
		"LevelDB",
		"CouchDB",
	}
}

type ThresholdComparator string

// Enum values for ThresholdComparator
const (
	ThresholdComparatorGreaterThan ThresholdComparator = "GREATER_THAN"
	ThresholdComparatorGreaterThanOrEqualTo ThresholdComparator = "GREATER_THAN_OR_EQUAL_TO"
)

// Values returns all known values for ThresholdComparator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ThresholdComparator) Values() []ThresholdComparator {
	return []ThresholdComparator{
		"GREATER_THAN",
		"GREATER_THAN_OR_EQUAL_TO",
	}
}

type VoteValue string

// Enum values for VoteValue
const (
	VoteValueYes VoteValue = "YES"
	VoteValueNo VoteValue = "NO"
)

// Values returns all known values for VoteValue. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (VoteValue) Values() []VoteValue {
	return []VoteValue{
		"YES",
		"NO",
	}
}
