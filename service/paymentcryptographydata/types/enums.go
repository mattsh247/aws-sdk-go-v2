// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type DukptDerivationType string

// Enum values for DukptDerivationType
const (
	DukptDerivationTypeTdes2key DukptDerivationType = "TDES_2KEY"
	DukptDerivationTypeTdes3key DukptDerivationType = "TDES_3KEY"
	DukptDerivationTypeAes128 DukptDerivationType = "AES_128"
	DukptDerivationTypeAes192 DukptDerivationType = "AES_192"
	DukptDerivationTypeAes256 DukptDerivationType = "AES_256"
)

// Values returns all known values for DukptDerivationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DukptDerivationType) Values() []DukptDerivationType {
	return []DukptDerivationType{
		"TDES_2KEY",
		"TDES_3KEY",
		"AES_128",
		"AES_192",
		"AES_256",
	}
}

type DukptEncryptionMode string

// Enum values for DukptEncryptionMode
const (
	DukptEncryptionModeEcb DukptEncryptionMode = "ECB"
	DukptEncryptionModeCbc DukptEncryptionMode = "CBC"
)

// Values returns all known values for DukptEncryptionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DukptEncryptionMode) Values() []DukptEncryptionMode {
	return []DukptEncryptionMode{
		"ECB",
		"CBC",
	}
}

type DukptKeyVariant string

// Enum values for DukptKeyVariant
const (
	DukptKeyVariantBidirectional DukptKeyVariant = "BIDIRECTIONAL"
	DukptKeyVariantRequest DukptKeyVariant = "REQUEST"
	DukptKeyVariantResponse DukptKeyVariant = "RESPONSE"
)

// Values returns all known values for DukptKeyVariant. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DukptKeyVariant) Values() []DukptKeyVariant {
	return []DukptKeyVariant{
		"BIDIRECTIONAL",
		"REQUEST",
		"RESPONSE",
	}
}

type EncryptionMode string

// Enum values for EncryptionMode
const (
	EncryptionModeEcb EncryptionMode = "ECB"
	EncryptionModeCbc EncryptionMode = "CBC"
	EncryptionModeCfb EncryptionMode = "CFB"
	EncryptionModeCfb1 EncryptionMode = "CFB1"
	EncryptionModeCfb8 EncryptionMode = "CFB8"
	EncryptionModeCfb64 EncryptionMode = "CFB64"
	EncryptionModeCfb128 EncryptionMode = "CFB128"
	EncryptionModeOfb EncryptionMode = "OFB"
)

// Values returns all known values for EncryptionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionMode) Values() []EncryptionMode {
	return []EncryptionMode{
		"ECB",
		"CBC",
		"CFB",
		"CFB1",
		"CFB8",
		"CFB64",
		"CFB128",
		"OFB",
	}
}

type MacAlgorithm string

// Enum values for MacAlgorithm
const (
	MacAlgorithmIso9797Algorithm1 MacAlgorithm = "ISO9797_ALGORITHM1"
	MacAlgorithmIso9797Algorithm3 MacAlgorithm = "ISO9797_ALGORITHM3"
	MacAlgorithmCmac MacAlgorithm = "CMAC"
	MacAlgorithmHmacSha224 MacAlgorithm = "HMAC_SHA224"
	MacAlgorithmHmacSha256 MacAlgorithm = "HMAC_SHA256"
	MacAlgorithmHmacSha384 MacAlgorithm = "HMAC_SHA384"
	MacAlgorithmHmacSha512 MacAlgorithm = "HMAC_SHA512"
)

// Values returns all known values for MacAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MacAlgorithm) Values() []MacAlgorithm {
	return []MacAlgorithm{
		"ISO9797_ALGORITHM1",
		"ISO9797_ALGORITHM3",
		"CMAC",
		"HMAC_SHA224",
		"HMAC_SHA256",
		"HMAC_SHA384",
		"HMAC_SHA512",
	}
}

type MajorKeyDerivationMode string

// Enum values for MajorKeyDerivationMode
const (
	MajorKeyDerivationModeEmvOptionA MajorKeyDerivationMode = "EMV_OPTION_A"
	MajorKeyDerivationModeEmvOptionB MajorKeyDerivationMode = "EMV_OPTION_B"
)

// Values returns all known values for MajorKeyDerivationMode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MajorKeyDerivationMode) Values() []MajorKeyDerivationMode {
	return []MajorKeyDerivationMode{
		"EMV_OPTION_A",
		"EMV_OPTION_B",
	}
}

type PaddingType string

// Enum values for PaddingType
const (
	PaddingTypePkcs1 PaddingType = "PKCS1"
	PaddingTypeOaepSha1 PaddingType = "OAEP_SHA1"
	PaddingTypeOaepSha256 PaddingType = "OAEP_SHA256"
	PaddingTypeOaepSha512 PaddingType = "OAEP_SHA512"
)

// Values returns all known values for PaddingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PaddingType) Values() []PaddingType {
	return []PaddingType{
		"PKCS1",
		"OAEP_SHA1",
		"OAEP_SHA256",
		"OAEP_SHA512",
	}
}

type PinBlockFormatForPinData string

// Enum values for PinBlockFormatForPinData
const (
	PinBlockFormatForPinDataIsoFormat0 PinBlockFormatForPinData = "ISO_FORMAT_0"
	PinBlockFormatForPinDataIsoFormat3 PinBlockFormatForPinData = "ISO_FORMAT_3"
)

// Values returns all known values for PinBlockFormatForPinData. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (PinBlockFormatForPinData) Values() []PinBlockFormatForPinData {
	return []PinBlockFormatForPinData{
		"ISO_FORMAT_0",
		"ISO_FORMAT_3",
	}
}

type SessionKeyDerivationMode string

// Enum values for SessionKeyDerivationMode
const (
	SessionKeyDerivationModeEmvCommonSessionKey SessionKeyDerivationMode = "EMV_COMMON_SESSION_KEY"
	SessionKeyDerivationModeEmv2000 SessionKeyDerivationMode = "EMV2000"
	SessionKeyDerivationModeAmex SessionKeyDerivationMode = "AMEX"
	SessionKeyDerivationModeMastercardSessionKey SessionKeyDerivationMode = "MASTERCARD_SESSION_KEY"
	SessionKeyDerivationModeVisa SessionKeyDerivationMode = "VISA"
)

// Values returns all known values for SessionKeyDerivationMode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SessionKeyDerivationMode) Values() []SessionKeyDerivationMode {
	return []SessionKeyDerivationMode{
		"EMV_COMMON_SESSION_KEY",
		"EMV2000",
		"AMEX",
		"MASTERCARD_SESSION_KEY",
		"VISA",
	}
}

type VerificationFailedReason string

// Enum values for VerificationFailedReason
const (
	VerificationFailedReasonInvalidMac VerificationFailedReason = "INVALID_MAC"
	VerificationFailedReasonInvalidPin VerificationFailedReason = "INVALID_PIN"
	VerificationFailedReasonInvalidValidationData VerificationFailedReason = "INVALID_VALIDATION_DATA"
	VerificationFailedReasonInvalidAuthRequestCryptogram VerificationFailedReason = "INVALID_AUTH_REQUEST_CRYPTOGRAM"
)

// Values returns all known values for VerificationFailedReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (VerificationFailedReason) Values() []VerificationFailedReason {
	return []VerificationFailedReason{
		"INVALID_MAC",
		"INVALID_PIN",
		"INVALID_VALIDATION_DATA",
		"INVALID_AUTH_REQUEST_CRYPTOGRAM",
	}
}
