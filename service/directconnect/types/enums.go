// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AddressFamily string

// Enum values for AddressFamily
const (
	AddressFamilyIPv4 AddressFamily = "ipv4"
	AddressFamilyIPv6 AddressFamily = "ipv6"
)

// Values returns all known values for AddressFamily. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AddressFamily) Values() []AddressFamily {
	return []AddressFamily{
		"ipv4",
		"ipv6",
	}
}

type BGPPeerState string

// Enum values for BGPPeerState
const (
	BGPPeerStateVerifying BGPPeerState = "verifying"
	BGPPeerStatePending BGPPeerState = "pending"
	BGPPeerStateAvailable BGPPeerState = "available"
	BGPPeerStateDeleting BGPPeerState = "deleting"
	BGPPeerStateDeleted BGPPeerState = "deleted"
)

// Values returns all known values for BGPPeerState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BGPPeerState) Values() []BGPPeerState {
	return []BGPPeerState{
		"verifying",
		"pending",
		"available",
		"deleting",
		"deleted",
	}
}

type BGPStatus string

// Enum values for BGPStatus
const (
	BGPStatusUp BGPStatus = "up"
	BGPStatusDown BGPStatus = "down"
	BGPStatusUnknown BGPStatus = "unknown"
)

// Values returns all known values for BGPStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BGPStatus) Values() []BGPStatus {
	return []BGPStatus{
		"up",
		"down",
		"unknown",
	}
}

type ConnectionState string

// Enum values for ConnectionState
const (
	ConnectionStateOrdering ConnectionState = "ordering"
	ConnectionStateRequested ConnectionState = "requested"
	ConnectionStatePending ConnectionState = "pending"
	ConnectionStateAvailable ConnectionState = "available"
	ConnectionStateDown ConnectionState = "down"
	ConnectionStateDeleting ConnectionState = "deleting"
	ConnectionStateDeleted ConnectionState = "deleted"
	ConnectionStateRejected ConnectionState = "rejected"
	ConnectionStateUnknown ConnectionState = "unknown"
)

// Values returns all known values for ConnectionState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionState) Values() []ConnectionState {
	return []ConnectionState{
		"ordering",
		"requested",
		"pending",
		"available",
		"down",
		"deleting",
		"deleted",
		"rejected",
		"unknown",
	}
}

type DirectConnectGatewayAssociationProposalState string

// Enum values for DirectConnectGatewayAssociationProposalState
const (
	DirectConnectGatewayAssociationProposalStateRequested DirectConnectGatewayAssociationProposalState = "requested"
	DirectConnectGatewayAssociationProposalStateAccepted DirectConnectGatewayAssociationProposalState = "accepted"
	DirectConnectGatewayAssociationProposalStateDeleted DirectConnectGatewayAssociationProposalState = "deleted"
)

// Values returns all known values for
// DirectConnectGatewayAssociationProposalState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DirectConnectGatewayAssociationProposalState) Values() []DirectConnectGatewayAssociationProposalState {
	return []DirectConnectGatewayAssociationProposalState{
		"requested",
		"accepted",
		"deleted",
	}
}

type DirectConnectGatewayAssociationState string

// Enum values for DirectConnectGatewayAssociationState
const (
	DirectConnectGatewayAssociationStateAssociating DirectConnectGatewayAssociationState = "associating"
	DirectConnectGatewayAssociationStateAssociated DirectConnectGatewayAssociationState = "associated"
	DirectConnectGatewayAssociationStateDisassociating DirectConnectGatewayAssociationState = "disassociating"
	DirectConnectGatewayAssociationStateDisassociated DirectConnectGatewayAssociationState = "disassociated"
	DirectConnectGatewayAssociationStateUpdating DirectConnectGatewayAssociationState = "updating"
)

// Values returns all known values for DirectConnectGatewayAssociationState. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DirectConnectGatewayAssociationState) Values() []DirectConnectGatewayAssociationState {
	return []DirectConnectGatewayAssociationState{
		"associating",
		"associated",
		"disassociating",
		"disassociated",
		"updating",
	}
}

type DirectConnectGatewayAttachmentState string

// Enum values for DirectConnectGatewayAttachmentState
const (
	DirectConnectGatewayAttachmentStateAttaching DirectConnectGatewayAttachmentState = "attaching"
	DirectConnectGatewayAttachmentStateAttached DirectConnectGatewayAttachmentState = "attached"
	DirectConnectGatewayAttachmentStateDetaching DirectConnectGatewayAttachmentState = "detaching"
	DirectConnectGatewayAttachmentStateDetached DirectConnectGatewayAttachmentState = "detached"
)

// Values returns all known values for DirectConnectGatewayAttachmentState. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DirectConnectGatewayAttachmentState) Values() []DirectConnectGatewayAttachmentState {
	return []DirectConnectGatewayAttachmentState{
		"attaching",
		"attached",
		"detaching",
		"detached",
	}
}

type DirectConnectGatewayAttachmentType string

// Enum values for DirectConnectGatewayAttachmentType
const (
	DirectConnectGatewayAttachmentTypeTransitVirtualInterface DirectConnectGatewayAttachmentType = "TransitVirtualInterface"
	DirectConnectGatewayAttachmentTypePrivateVirtualInterface DirectConnectGatewayAttachmentType = "PrivateVirtualInterface"
)

// Values returns all known values for DirectConnectGatewayAttachmentType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (DirectConnectGatewayAttachmentType) Values() []DirectConnectGatewayAttachmentType {
	return []DirectConnectGatewayAttachmentType{
		"TransitVirtualInterface",
		"PrivateVirtualInterface",
	}
}

type DirectConnectGatewayState string

// Enum values for DirectConnectGatewayState
const (
	DirectConnectGatewayStatePending DirectConnectGatewayState = "pending"
	DirectConnectGatewayStateAvailable DirectConnectGatewayState = "available"
	DirectConnectGatewayStateDeleting DirectConnectGatewayState = "deleting"
	DirectConnectGatewayStateDeleted DirectConnectGatewayState = "deleted"
)

// Values returns all known values for DirectConnectGatewayState. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DirectConnectGatewayState) Values() []DirectConnectGatewayState {
	return []DirectConnectGatewayState{
		"pending",
		"available",
		"deleting",
		"deleted",
	}
}

type GatewayType string

// Enum values for GatewayType
const (
	GatewayTypeVirtualPrivateGateway GatewayType = "virtualPrivateGateway"
	GatewayTypeTransitGateway GatewayType = "transitGateway"
)

// Values returns all known values for GatewayType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (GatewayType) Values() []GatewayType {
	return []GatewayType{
		"virtualPrivateGateway",
		"transitGateway",
	}
}

type HasLogicalRedundancy string

// Enum values for HasLogicalRedundancy
const (
	HasLogicalRedundancyUnknown HasLogicalRedundancy = "unknown"
	HasLogicalRedundancyYes HasLogicalRedundancy = "yes"
	HasLogicalRedundancyNo HasLogicalRedundancy = "no"
)

// Values returns all known values for HasLogicalRedundancy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HasLogicalRedundancy) Values() []HasLogicalRedundancy {
	return []HasLogicalRedundancy{
		"unknown",
		"yes",
		"no",
	}
}

type InterconnectState string

// Enum values for InterconnectState
const (
	InterconnectStateRequested InterconnectState = "requested"
	InterconnectStatePending InterconnectState = "pending"
	InterconnectStateAvailable InterconnectState = "available"
	InterconnectStateDown InterconnectState = "down"
	InterconnectStateDeleting InterconnectState = "deleting"
	InterconnectStateDeleted InterconnectState = "deleted"
	InterconnectStateUnknown InterconnectState = "unknown"
)

// Values returns all known values for InterconnectState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InterconnectState) Values() []InterconnectState {
	return []InterconnectState{
		"requested",
		"pending",
		"available",
		"down",
		"deleting",
		"deleted",
		"unknown",
	}
}

type LagState string

// Enum values for LagState
const (
	LagStateRequested LagState = "requested"
	LagStatePending LagState = "pending"
	LagStateAvailable LagState = "available"
	LagStateDown LagState = "down"
	LagStateDeleting LagState = "deleting"
	LagStateDeleted LagState = "deleted"
	LagStateUnknown LagState = "unknown"
)

// Values returns all known values for LagState. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (LagState) Values() []LagState {
	return []LagState{
		"requested",
		"pending",
		"available",
		"down",
		"deleting",
		"deleted",
		"unknown",
	}
}

type LoaContentType string

// Enum values for LoaContentType
const (
	LoaContentTypePdf LoaContentType = "application/pdf"
)

// Values returns all known values for LoaContentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LoaContentType) Values() []LoaContentType {
	return []LoaContentType{
		"application/pdf",
	}
}

type NniPartnerType string

// Enum values for NniPartnerType
const (
	NniPartnerTypeV1 NniPartnerType = "v1"
	NniPartnerTypeV2 NniPartnerType = "v2"
	NniPartnerTypeNonPartner NniPartnerType = "nonPartner"
)

// Values returns all known values for NniPartnerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NniPartnerType) Values() []NniPartnerType {
	return []NniPartnerType{
		"v1",
		"v2",
		"nonPartner",
	}
}

type VirtualInterfaceState string

// Enum values for VirtualInterfaceState
const (
	VirtualInterfaceStateConfirming VirtualInterfaceState = "confirming"
	VirtualInterfaceStateVerifying VirtualInterfaceState = "verifying"
	VirtualInterfaceStatePending VirtualInterfaceState = "pending"
	VirtualInterfaceStateAvailable VirtualInterfaceState = "available"
	VirtualInterfaceStateDown VirtualInterfaceState = "down"
	VirtualInterfaceStateDeleting VirtualInterfaceState = "deleting"
	VirtualInterfaceStateDeleted VirtualInterfaceState = "deleted"
	VirtualInterfaceStateRejected VirtualInterfaceState = "rejected"
	VirtualInterfaceStateUnknown VirtualInterfaceState = "unknown"
)

// Values returns all known values for VirtualInterfaceState. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VirtualInterfaceState) Values() []VirtualInterfaceState {
	return []VirtualInterfaceState{
		"confirming",
		"verifying",
		"pending",
		"available",
		"down",
		"deleting",
		"deleted",
		"rejected",
		"unknown",
	}
}
