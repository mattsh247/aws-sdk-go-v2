// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AccountStatus string

// Enum values for AccountStatus
const (
	AccountStatusSuspended AccountStatus = "Suspended"
	AccountStatusActive AccountStatus = "Active"
)

// Values returns all known values for AccountStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccountStatus) Values() []AccountStatus {
	return []AccountStatus{
		"Suspended",
		"Active",
	}
}

type AccountType string

// Enum values for AccountType
const (
	AccountTypeTeam AccountType = "Team"
	AccountTypeEnterpriseDirectory AccountType = "EnterpriseDirectory"
	AccountTypeEnterpriseLWA AccountType = "EnterpriseLWA"
	AccountTypeEnterpriseOIDC AccountType = "EnterpriseOIDC"
)

// Values returns all known values for AccountType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AccountType) Values() []AccountType {
	return []AccountType{
		"Team",
		"EnterpriseDirectory",
		"EnterpriseLWA",
		"EnterpriseOIDC",
	}
}

type AppInstanceDataType string

// Enum values for AppInstanceDataType
const (
	AppInstanceDataTypeChannel AppInstanceDataType = "Channel"
	AppInstanceDataTypeChannelMessage AppInstanceDataType = "ChannelMessage"
)

// Values returns all known values for AppInstanceDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppInstanceDataType) Values() []AppInstanceDataType {
	return []AppInstanceDataType{
		"Channel",
		"ChannelMessage",
	}
}

type ArtifactsState string

// Enum values for ArtifactsState
const (
	ArtifactsStateEnabled ArtifactsState = "Enabled"
	ArtifactsStateDisabled ArtifactsState = "Disabled"
)

// Values returns all known values for ArtifactsState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactsState) Values() []ArtifactsState {
	return []ArtifactsState{
		"Enabled",
		"Disabled",
	}
}

type AudioMuxType string

// Enum values for AudioMuxType
const (
	AudioMuxTypeAudioOnly AudioMuxType = "AudioOnly"
	AudioMuxTypeAudioWithActiveSpeakerVideo AudioMuxType = "AudioWithActiveSpeakerVideo"
)

// Values returns all known values for AudioMuxType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AudioMuxType) Values() []AudioMuxType {
	return []AudioMuxType{
		"AudioOnly",
		"AudioWithActiveSpeakerVideo",
	}
}

type BotType string

// Enum values for BotType
const (
	BotTypeChatBot BotType = "ChatBot"
)

// Values returns all known values for BotType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (BotType) Values() []BotType {
	return []BotType{
		"ChatBot",
	}
}

type CallingNameStatus string

// Enum values for CallingNameStatus
const (
	CallingNameStatusUnassigned CallingNameStatus = "Unassigned"
	CallingNameStatusUpdateInProgress CallingNameStatus = "UpdateInProgress"
	CallingNameStatusUpdateSucceeded CallingNameStatus = "UpdateSucceeded"
	CallingNameStatusUpdateFailed CallingNameStatus = "UpdateFailed"
)

// Values returns all known values for CallingNameStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CallingNameStatus) Values() []CallingNameStatus {
	return []CallingNameStatus{
		"Unassigned",
		"UpdateInProgress",
		"UpdateSucceeded",
		"UpdateFailed",
	}
}

type Capability string

// Enum values for Capability
const (
	CapabilityVoice Capability = "Voice"
	CapabilitySms Capability = "SMS"
)

// Values returns all known values for Capability. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Capability) Values() []Capability {
	return []Capability{
		"Voice",
		"SMS",
	}
}

type ChannelMembershipType string

// Enum values for ChannelMembershipType
const (
	ChannelMembershipTypeDefault ChannelMembershipType = "DEFAULT"
	ChannelMembershipTypeHidden ChannelMembershipType = "HIDDEN"
)

// Values returns all known values for ChannelMembershipType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelMembershipType) Values() []ChannelMembershipType {
	return []ChannelMembershipType{
		"DEFAULT",
		"HIDDEN",
	}
}

type ChannelMessagePersistenceType string

// Enum values for ChannelMessagePersistenceType
const (
	ChannelMessagePersistenceTypePersistent ChannelMessagePersistenceType = "PERSISTENT"
	ChannelMessagePersistenceTypeNonPersistent ChannelMessagePersistenceType = "NON_PERSISTENT"
)

// Values returns all known values for ChannelMessagePersistenceType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ChannelMessagePersistenceType) Values() []ChannelMessagePersistenceType {
	return []ChannelMessagePersistenceType{
		"PERSISTENT",
		"NON_PERSISTENT",
	}
}

type ChannelMessageType string

// Enum values for ChannelMessageType
const (
	ChannelMessageTypeStandard ChannelMessageType = "STANDARD"
	ChannelMessageTypeControl ChannelMessageType = "CONTROL"
)

// Values returns all known values for ChannelMessageType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelMessageType) Values() []ChannelMessageType {
	return []ChannelMessageType{
		"STANDARD",
		"CONTROL",
	}
}

type ChannelMode string

// Enum values for ChannelMode
const (
	ChannelModeUnrestricted ChannelMode = "UNRESTRICTED"
	ChannelModeRestricted ChannelMode = "RESTRICTED"
)

// Values returns all known values for ChannelMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChannelMode) Values() []ChannelMode {
	return []ChannelMode{
		"UNRESTRICTED",
		"RESTRICTED",
	}
}

type ChannelPrivacy string

// Enum values for ChannelPrivacy
const (
	ChannelPrivacyPublic ChannelPrivacy = "PUBLIC"
	ChannelPrivacyPrivate ChannelPrivacy = "PRIVATE"
)

// Values returns all known values for ChannelPrivacy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChannelPrivacy) Values() []ChannelPrivacy {
	return []ChannelPrivacy{
		"PUBLIC",
		"PRIVATE",
	}
}

type ContentMuxType string

// Enum values for ContentMuxType
const (
	ContentMuxTypeContentOnly ContentMuxType = "ContentOnly"
)

// Values returns all known values for ContentMuxType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ContentMuxType) Values() []ContentMuxType {
	return []ContentMuxType{
		"ContentOnly",
	}
}

type EmailStatus string

// Enum values for EmailStatus
const (
	EmailStatusNotSent EmailStatus = "NotSent"
	EmailStatusSent EmailStatus = "Sent"
	EmailStatusFailed EmailStatus = "Failed"
)

// Values returns all known values for EmailStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EmailStatus) Values() []EmailStatus {
	return []EmailStatus{
		"NotSent",
		"Sent",
		"Failed",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest ErrorCode = "BadRequest"
	ErrorCodeConflict ErrorCode = "Conflict"
	ErrorCodeForbidden ErrorCode = "Forbidden"
	ErrorCodeNotFound ErrorCode = "NotFound"
	ErrorCodePreconditionFailed ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled ErrorCode = "Throttled"
	ErrorCodeThrottling ErrorCode = "Throttling"
	ErrorCodeUnauthorized ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}

type GeoMatchLevel string

// Enum values for GeoMatchLevel
const (
	GeoMatchLevelCountry GeoMatchLevel = "Country"
	GeoMatchLevelAreaCode GeoMatchLevel = "AreaCode"
)

// Values returns all known values for GeoMatchLevel. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoMatchLevel) Values() []GeoMatchLevel {
	return []GeoMatchLevel{
		"Country",
		"AreaCode",
	}
}

type InviteStatus string

// Enum values for InviteStatus
const (
	InviteStatusPending InviteStatus = "Pending"
	InviteStatusAccepted InviteStatus = "Accepted"
	InviteStatusFailed InviteStatus = "Failed"
)

// Values returns all known values for InviteStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InviteStatus) Values() []InviteStatus {
	return []InviteStatus{
		"Pending",
		"Accepted",
		"Failed",
	}
}

type License string

// Enum values for License
const (
	LicenseBasic License = "Basic"
	LicensePlus License = "Plus"
	LicensePro License = "Pro"
	LicenseProTrial License = "ProTrial"
)

// Values returns all known values for License. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (License) Values() []License {
	return []License{
		"Basic",
		"Plus",
		"Pro",
		"ProTrial",
	}
}

type MediaPipelineSinkType string

// Enum values for MediaPipelineSinkType
const (
	MediaPipelineSinkTypeS3Bucket MediaPipelineSinkType = "S3Bucket"
)

// Values returns all known values for MediaPipelineSinkType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MediaPipelineSinkType) Values() []MediaPipelineSinkType {
	return []MediaPipelineSinkType{
		"S3Bucket",
	}
}

type MediaPipelineSourceType string

// Enum values for MediaPipelineSourceType
const (
	MediaPipelineSourceTypeChimeSdkMeeting MediaPipelineSourceType = "ChimeSdkMeeting"
)

// Values returns all known values for MediaPipelineSourceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MediaPipelineSourceType) Values() []MediaPipelineSourceType {
	return []MediaPipelineSourceType{
		"ChimeSdkMeeting",
	}
}

type MediaPipelineStatus string

// Enum values for MediaPipelineStatus
const (
	MediaPipelineStatusInitializing MediaPipelineStatus = "Initializing"
	MediaPipelineStatusInProgress MediaPipelineStatus = "InProgress"
	MediaPipelineStatusFailed MediaPipelineStatus = "Failed"
	MediaPipelineStatusStopping MediaPipelineStatus = "Stopping"
	MediaPipelineStatusStopped MediaPipelineStatus = "Stopped"
)

// Values returns all known values for MediaPipelineStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MediaPipelineStatus) Values() []MediaPipelineStatus {
	return []MediaPipelineStatus{
		"Initializing",
		"InProgress",
		"Failed",
		"Stopping",
		"Stopped",
	}
}

type MemberType string

// Enum values for MemberType
const (
	MemberTypeUser MemberType = "User"
	MemberTypeBot MemberType = "Bot"
	MemberTypeWebhook MemberType = "Webhook"
)

// Values returns all known values for MemberType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MemberType) Values() []MemberType {
	return []MemberType{
		"User",
		"Bot",
		"Webhook",
	}
}

type NotificationTarget string

// Enum values for NotificationTarget
const (
	NotificationTargetEventBridge NotificationTarget = "EventBridge"
	NotificationTargetSns NotificationTarget = "SNS"
	NotificationTargetSqs NotificationTarget = "SQS"
)

// Values returns all known values for NotificationTarget. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NotificationTarget) Values() []NotificationTarget {
	return []NotificationTarget{
		"EventBridge",
		"SNS",
		"SQS",
	}
}

type NumberSelectionBehavior string

// Enum values for NumberSelectionBehavior
const (
	NumberSelectionBehaviorPreferSticky NumberSelectionBehavior = "PreferSticky"
	NumberSelectionBehaviorAvoidSticky NumberSelectionBehavior = "AvoidSticky"
)

// Values returns all known values for NumberSelectionBehavior. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NumberSelectionBehavior) Values() []NumberSelectionBehavior {
	return []NumberSelectionBehavior{
		"PreferSticky",
		"AvoidSticky",
	}
}

type OrderedPhoneNumberStatus string

// Enum values for OrderedPhoneNumberStatus
const (
	OrderedPhoneNumberStatusProcessing OrderedPhoneNumberStatus = "Processing"
	OrderedPhoneNumberStatusAcquired OrderedPhoneNumberStatus = "Acquired"
	OrderedPhoneNumberStatusFailed OrderedPhoneNumberStatus = "Failed"
)

// Values returns all known values for OrderedPhoneNumberStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrderedPhoneNumberStatus) Values() []OrderedPhoneNumberStatus {
	return []OrderedPhoneNumberStatus{
		"Processing",
		"Acquired",
		"Failed",
	}
}

type OriginationRouteProtocol string

// Enum values for OriginationRouteProtocol
const (
	OriginationRouteProtocolTcp OriginationRouteProtocol = "TCP"
	OriginationRouteProtocolUdp OriginationRouteProtocol = "UDP"
)

// Values returns all known values for OriginationRouteProtocol. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (OriginationRouteProtocol) Values() []OriginationRouteProtocol {
	return []OriginationRouteProtocol{
		"TCP",
		"UDP",
	}
}

type PhoneNumberAssociationName string

// Enum values for PhoneNumberAssociationName
const (
	PhoneNumberAssociationNameAccountId PhoneNumberAssociationName = "AccountId"
	PhoneNumberAssociationNameUserId PhoneNumberAssociationName = "UserId"
	PhoneNumberAssociationNameVoiceConnectorId PhoneNumberAssociationName = "VoiceConnectorId"
	PhoneNumberAssociationNameVoiceConnectorGroupId PhoneNumberAssociationName = "VoiceConnectorGroupId"
	PhoneNumberAssociationNameSipRuleId PhoneNumberAssociationName = "SipRuleId"
)

// Values returns all known values for PhoneNumberAssociationName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberAssociationName) Values() []PhoneNumberAssociationName {
	return []PhoneNumberAssociationName{
		"AccountId",
		"UserId",
		"VoiceConnectorId",
		"VoiceConnectorGroupId",
		"SipRuleId",
	}
}

type PhoneNumberOrderStatus string

// Enum values for PhoneNumberOrderStatus
const (
	PhoneNumberOrderStatusProcessing PhoneNumberOrderStatus = "Processing"
	PhoneNumberOrderStatusSuccessful PhoneNumberOrderStatus = "Successful"
	PhoneNumberOrderStatusFailed PhoneNumberOrderStatus = "Failed"
	PhoneNumberOrderStatusPartial PhoneNumberOrderStatus = "Partial"
)

// Values returns all known values for PhoneNumberOrderStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberOrderStatus) Values() []PhoneNumberOrderStatus {
	return []PhoneNumberOrderStatus{
		"Processing",
		"Successful",
		"Failed",
		"Partial",
	}
}

type PhoneNumberProductType string

// Enum values for PhoneNumberProductType
const (
	PhoneNumberProductTypeBusinessCalling PhoneNumberProductType = "BusinessCalling"
	PhoneNumberProductTypeVoiceConnector PhoneNumberProductType = "VoiceConnector"
	PhoneNumberProductTypeSipMediaApplicationDialIn PhoneNumberProductType = "SipMediaApplicationDialIn"
)

// Values returns all known values for PhoneNumberProductType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberProductType) Values() []PhoneNumberProductType {
	return []PhoneNumberProductType{
		"BusinessCalling",
		"VoiceConnector",
		"SipMediaApplicationDialIn",
	}
}

type PhoneNumberStatus string

// Enum values for PhoneNumberStatus
const (
	PhoneNumberStatusAcquireInProgress PhoneNumberStatus = "AcquireInProgress"
	PhoneNumberStatusAcquireFailed PhoneNumberStatus = "AcquireFailed"
	PhoneNumberStatusUnassigned PhoneNumberStatus = "Unassigned"
	PhoneNumberStatusAssigned PhoneNumberStatus = "Assigned"
	PhoneNumberStatusReleaseInProgress PhoneNumberStatus = "ReleaseInProgress"
	PhoneNumberStatusDeleteInProgress PhoneNumberStatus = "DeleteInProgress"
	PhoneNumberStatusReleaseFailed PhoneNumberStatus = "ReleaseFailed"
	PhoneNumberStatusDeleteFailed PhoneNumberStatus = "DeleteFailed"
)

// Values returns all known values for PhoneNumberStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberStatus) Values() []PhoneNumberStatus {
	return []PhoneNumberStatus{
		"AcquireInProgress",
		"AcquireFailed",
		"Unassigned",
		"Assigned",
		"ReleaseInProgress",
		"DeleteInProgress",
		"ReleaseFailed",
		"DeleteFailed",
	}
}

type PhoneNumberType string

// Enum values for PhoneNumberType
const (
	PhoneNumberTypeLocal PhoneNumberType = "Local"
	PhoneNumberTypeTollFree PhoneNumberType = "TollFree"
)

// Values returns all known values for PhoneNumberType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PhoneNumberType) Values() []PhoneNumberType {
	return []PhoneNumberType{
		"Local",
		"TollFree",
	}
}

type ProxySessionStatus string

// Enum values for ProxySessionStatus
const (
	ProxySessionStatusOpen ProxySessionStatus = "Open"
	ProxySessionStatusInProgress ProxySessionStatus = "InProgress"
	ProxySessionStatusClosed ProxySessionStatus = "Closed"
)

// Values returns all known values for ProxySessionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProxySessionStatus) Values() []ProxySessionStatus {
	return []ProxySessionStatus{
		"Open",
		"InProgress",
		"Closed",
	}
}

type RegistrationStatus string

// Enum values for RegistrationStatus
const (
	RegistrationStatusUnregistered RegistrationStatus = "Unregistered"
	RegistrationStatusRegistered RegistrationStatus = "Registered"
	RegistrationStatusSuspended RegistrationStatus = "Suspended"
)

// Values returns all known values for RegistrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RegistrationStatus) Values() []RegistrationStatus {
	return []RegistrationStatus{
		"Unregistered",
		"Registered",
		"Suspended",
	}
}

type RoomMembershipRole string

// Enum values for RoomMembershipRole
const (
	RoomMembershipRoleAdministrator RoomMembershipRole = "Administrator"
	RoomMembershipRoleMember RoomMembershipRole = "Member"
)

// Values returns all known values for RoomMembershipRole. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RoomMembershipRole) Values() []RoomMembershipRole {
	return []RoomMembershipRole{
		"Administrator",
		"Member",
	}
}

type SipRuleTriggerType string

// Enum values for SipRuleTriggerType
const (
	SipRuleTriggerTypeToPhoneNumber SipRuleTriggerType = "ToPhoneNumber"
	SipRuleTriggerTypeRequestUriHostname SipRuleTriggerType = "RequestUriHostname"
)

// Values returns all known values for SipRuleTriggerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SipRuleTriggerType) Values() []SipRuleTriggerType {
	return []SipRuleTriggerType{
		"ToPhoneNumber",
		"RequestUriHostname",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type TranscribeContentIdentificationType string

// Enum values for TranscribeContentIdentificationType
const (
	TranscribeContentIdentificationTypePii TranscribeContentIdentificationType = "PII"
)

// Values returns all known values for TranscribeContentIdentificationType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeContentIdentificationType) Values() []TranscribeContentIdentificationType {
	return []TranscribeContentIdentificationType{
		"PII",
	}
}

type TranscribeContentRedactionType string

// Enum values for TranscribeContentRedactionType
const (
	TranscribeContentRedactionTypePii TranscribeContentRedactionType = "PII"
)

// Values returns all known values for TranscribeContentRedactionType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeContentRedactionType) Values() []TranscribeContentRedactionType {
	return []TranscribeContentRedactionType{
		"PII",
	}
}

type TranscribeLanguageCode string

// Enum values for TranscribeLanguageCode
const (
	TranscribeLanguageCodeEnUs TranscribeLanguageCode = "en-US"
	TranscribeLanguageCodeEnGb TranscribeLanguageCode = "en-GB"
	TranscribeLanguageCodeEsUs TranscribeLanguageCode = "es-US"
	TranscribeLanguageCodeFrCa TranscribeLanguageCode = "fr-CA"
	TranscribeLanguageCodeFrFr TranscribeLanguageCode = "fr-FR"
	TranscribeLanguageCodeEnAu TranscribeLanguageCode = "en-AU"
	TranscribeLanguageCodeItIt TranscribeLanguageCode = "it-IT"
	TranscribeLanguageCodeDeDe TranscribeLanguageCode = "de-DE"
	TranscribeLanguageCodePtBr TranscribeLanguageCode = "pt-BR"
	TranscribeLanguageCodeJaJp TranscribeLanguageCode = "ja-JP"
	TranscribeLanguageCodeKoKr TranscribeLanguageCode = "ko-KR"
	TranscribeLanguageCodeZhCn TranscribeLanguageCode = "zh-CN"
	TranscribeLanguageCodeThTh TranscribeLanguageCode = "th-TH"
	TranscribeLanguageCodeHiIn TranscribeLanguageCode = "hi-IN"
)

// Values returns all known values for TranscribeLanguageCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeLanguageCode) Values() []TranscribeLanguageCode {
	return []TranscribeLanguageCode{
		"en-US",
		"en-GB",
		"es-US",
		"fr-CA",
		"fr-FR",
		"en-AU",
		"it-IT",
		"de-DE",
		"pt-BR",
		"ja-JP",
		"ko-KR",
		"zh-CN",
		"th-TH",
		"hi-IN",
	}
}

type TranscribeMedicalContentIdentificationType string

// Enum values for TranscribeMedicalContentIdentificationType
const (
	TranscribeMedicalContentIdentificationTypePhi TranscribeMedicalContentIdentificationType = "PHI"
)

// Values returns all known values for TranscribeMedicalContentIdentificationType.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeMedicalContentIdentificationType) Values() []TranscribeMedicalContentIdentificationType {
	return []TranscribeMedicalContentIdentificationType{
		"PHI",
	}
}

type TranscribeMedicalLanguageCode string

// Enum values for TranscribeMedicalLanguageCode
const (
	TranscribeMedicalLanguageCodeEnUs TranscribeMedicalLanguageCode = "en-US"
)

// Values returns all known values for TranscribeMedicalLanguageCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeMedicalLanguageCode) Values() []TranscribeMedicalLanguageCode {
	return []TranscribeMedicalLanguageCode{
		"en-US",
	}
}

type TranscribeMedicalRegion string

// Enum values for TranscribeMedicalRegion
const (
	TranscribeMedicalRegionUsEast1 TranscribeMedicalRegion = "us-east-1"
	TranscribeMedicalRegionUsEast2 TranscribeMedicalRegion = "us-east-2"
	TranscribeMedicalRegionUsWest2 TranscribeMedicalRegion = "us-west-2"
	TranscribeMedicalRegionApSoutheast2 TranscribeMedicalRegion = "ap-southeast-2"
	TranscribeMedicalRegionCaCentral1 TranscribeMedicalRegion = "ca-central-1"
	TranscribeMedicalRegionEuWest1 TranscribeMedicalRegion = "eu-west-1"
	TranscribeMedicalRegionAuto TranscribeMedicalRegion = "auto"
)

// Values returns all known values for TranscribeMedicalRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalRegion) Values() []TranscribeMedicalRegion {
	return []TranscribeMedicalRegion{
		"us-east-1",
		"us-east-2",
		"us-west-2",
		"ap-southeast-2",
		"ca-central-1",
		"eu-west-1",
		"auto",
	}
}

type TranscribeMedicalSpecialty string

// Enum values for TranscribeMedicalSpecialty
const (
	TranscribeMedicalSpecialtyPrimaryCare TranscribeMedicalSpecialty = "PRIMARYCARE"
	TranscribeMedicalSpecialtyCardiology TranscribeMedicalSpecialty = "CARDIOLOGY"
	TranscribeMedicalSpecialtyNeurology TranscribeMedicalSpecialty = "NEUROLOGY"
	TranscribeMedicalSpecialtyOncology TranscribeMedicalSpecialty = "ONCOLOGY"
	TranscribeMedicalSpecialtyRadiology TranscribeMedicalSpecialty = "RADIOLOGY"
	TranscribeMedicalSpecialtyUrology TranscribeMedicalSpecialty = "UROLOGY"
)

// Values returns all known values for TranscribeMedicalSpecialty. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalSpecialty) Values() []TranscribeMedicalSpecialty {
	return []TranscribeMedicalSpecialty{
		"PRIMARYCARE",
		"CARDIOLOGY",
		"NEUROLOGY",
		"ONCOLOGY",
		"RADIOLOGY",
		"UROLOGY",
	}
}

type TranscribeMedicalType string

// Enum values for TranscribeMedicalType
const (
	TranscribeMedicalTypeConversation TranscribeMedicalType = "CONVERSATION"
	TranscribeMedicalTypeDictation TranscribeMedicalType = "DICTATION"
)

// Values returns all known values for TranscribeMedicalType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeMedicalType) Values() []TranscribeMedicalType {
	return []TranscribeMedicalType{
		"CONVERSATION",
		"DICTATION",
	}
}

type TranscribePartialResultsStability string

// Enum values for TranscribePartialResultsStability
const (
	TranscribePartialResultsStabilityLow TranscribePartialResultsStability = "low"
	TranscribePartialResultsStabilityMedium TranscribePartialResultsStability = "medium"
	TranscribePartialResultsStabilityHigh TranscribePartialResultsStability = "high"
)

// Values returns all known values for TranscribePartialResultsStability. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribePartialResultsStability) Values() []TranscribePartialResultsStability {
	return []TranscribePartialResultsStability{
		"low",
		"medium",
		"high",
	}
}

type TranscribeRegion string

// Enum values for TranscribeRegion
const (
	TranscribeRegionUsEast2 TranscribeRegion = "us-east-2"
	TranscribeRegionUsEast1 TranscribeRegion = "us-east-1"
	TranscribeRegionUsWest2 TranscribeRegion = "us-west-2"
	TranscribeRegionApNortheast2 TranscribeRegion = "ap-northeast-2"
	TranscribeRegionApSoutheast2 TranscribeRegion = "ap-southeast-2"
	TranscribeRegionApNortheast1 TranscribeRegion = "ap-northeast-1"
	TranscribeRegionCaCentral1 TranscribeRegion = "ca-central-1"
	TranscribeRegionEuCentral1 TranscribeRegion = "eu-central-1"
	TranscribeRegionEuWest1 TranscribeRegion = "eu-west-1"
	TranscribeRegionEuWest2 TranscribeRegion = "eu-west-2"
	TranscribeRegionSaEast1 TranscribeRegion = "sa-east-1"
	TranscribeRegionAuto TranscribeRegion = "auto"
)

// Values returns all known values for TranscribeRegion. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TranscribeRegion) Values() []TranscribeRegion {
	return []TranscribeRegion{
		"us-east-2",
		"us-east-1",
		"us-west-2",
		"ap-northeast-2",
		"ap-southeast-2",
		"ap-northeast-1",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"sa-east-1",
		"auto",
	}
}

type TranscribeVocabularyFilterMethod string

// Enum values for TranscribeVocabularyFilterMethod
const (
	TranscribeVocabularyFilterMethodRemove TranscribeVocabularyFilterMethod = "remove"
	TranscribeVocabularyFilterMethodMask TranscribeVocabularyFilterMethod = "mask"
	TranscribeVocabularyFilterMethodTag TranscribeVocabularyFilterMethod = "tag"
)

// Values returns all known values for TranscribeVocabularyFilterMethod. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (TranscribeVocabularyFilterMethod) Values() []TranscribeVocabularyFilterMethod {
	return []TranscribeVocabularyFilterMethod{
		"remove",
		"mask",
		"tag",
	}
}

type UserType string

// Enum values for UserType
const (
	UserTypePrivateUser UserType = "PrivateUser"
	UserTypeSharedDevice UserType = "SharedDevice"
)

// Values returns all known values for UserType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserType) Values() []UserType {
	return []UserType{
		"PrivateUser",
		"SharedDevice",
	}
}

type VideoMuxType string

// Enum values for VideoMuxType
const (
	VideoMuxTypeVideoOnly VideoMuxType = "VideoOnly"
)

// Values returns all known values for VideoMuxType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VideoMuxType) Values() []VideoMuxType {
	return []VideoMuxType{
		"VideoOnly",
	}
}

type VoiceConnectorAwsRegion string

// Enum values for VoiceConnectorAwsRegion
const (
	VoiceConnectorAwsRegionUsEast1 VoiceConnectorAwsRegion = "us-east-1"
	VoiceConnectorAwsRegionUsWest2 VoiceConnectorAwsRegion = "us-west-2"
)

// Values returns all known values for VoiceConnectorAwsRegion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VoiceConnectorAwsRegion) Values() []VoiceConnectorAwsRegion {
	return []VoiceConnectorAwsRegion{
		"us-east-1",
		"us-west-2",
	}
}
