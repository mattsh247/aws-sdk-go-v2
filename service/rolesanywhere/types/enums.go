// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type NotificationChannel string

// Enum values for NotificationChannel
const (
	NotificationChannelAll NotificationChannel = "ALL"
)

// Values returns all known values for NotificationChannel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationChannel) Values() []NotificationChannel {
	return []NotificationChannel{
		"ALL",
	}
}

type NotificationEvent string

// Enum values for NotificationEvent
const (
	NotificationEventCaCertificateExpiry        NotificationEvent = "CA_CERTIFICATE_EXPIRY"
	NotificationEventEndEntityCertificateExpiry NotificationEvent = "END_ENTITY_CERTIFICATE_EXPIRY"
)

// Values returns all known values for NotificationEvent. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationEvent) Values() []NotificationEvent {
	return []NotificationEvent{
		"CA_CERTIFICATE_EXPIRY",
		"END_ENTITY_CERTIFICATE_EXPIRY",
	}
}

type TrustAnchorType string

// Enum values for TrustAnchorType
const (
	TrustAnchorTypeAwsAcmPca            TrustAnchorType = "AWS_ACM_PCA"
	TrustAnchorTypeCertificateBundle    TrustAnchorType = "CERTIFICATE_BUNDLE"
	TrustAnchorTypeSelfSignedRepository TrustAnchorType = "SELF_SIGNED_REPOSITORY"
)

// Values returns all known values for TrustAnchorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TrustAnchorType) Values() []TrustAnchorType {
	return []TrustAnchorType{
		"AWS_ACM_PCA",
		"CERTIFICATE_BUNDLE",
		"SELF_SIGNED_REPOSITORY",
	}
}
