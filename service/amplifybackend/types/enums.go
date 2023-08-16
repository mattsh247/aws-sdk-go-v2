// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdditionalConstraintsElement string

// Enum values for AdditionalConstraintsElement
const (
	AdditionalConstraintsElementRequireDigit     AdditionalConstraintsElement = "REQUIRE_DIGIT"
	AdditionalConstraintsElementRequireLowercase AdditionalConstraintsElement = "REQUIRE_LOWERCASE"
	AdditionalConstraintsElementRequireSymbol    AdditionalConstraintsElement = "REQUIRE_SYMBOL"
	AdditionalConstraintsElementRequireUppercase AdditionalConstraintsElement = "REQUIRE_UPPERCASE"
)

// Values returns all known values for AdditionalConstraintsElement. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AdditionalConstraintsElement) Values() []AdditionalConstraintsElement {
	return []AdditionalConstraintsElement{
		"REQUIRE_DIGIT",
		"REQUIRE_LOWERCASE",
		"REQUIRE_SYMBOL",
		"REQUIRE_UPPERCASE",
	}
}

type AuthenticatedElement string

// Enum values for AuthenticatedElement
const (
	AuthenticatedElementRead            AuthenticatedElement = "READ"
	AuthenticatedElementCreateAndUpdate AuthenticatedElement = "CREATE_AND_UPDATE"
	AuthenticatedElementDelete          AuthenticatedElement = "DELETE"
)

// Values returns all known values for AuthenticatedElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthenticatedElement) Values() []AuthenticatedElement {
	return []AuthenticatedElement{
		"READ",
		"CREATE_AND_UPDATE",
		"DELETE",
	}
}

type AuthResources string

// Enum values for AuthResources
const (
	AuthResourcesUserPoolOnly            AuthResources = "USER_POOL_ONLY"
	AuthResourcesIdentityPoolAndUserPool AuthResources = "IDENTITY_POOL_AND_USER_POOL"
)

// Values returns all known values for AuthResources. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuthResources) Values() []AuthResources {
	return []AuthResources{
		"USER_POOL_ONLY",
		"IDENTITY_POOL_AND_USER_POOL",
	}
}

type DeliveryMethod string

// Enum values for DeliveryMethod
const (
	DeliveryMethodEmail DeliveryMethod = "EMAIL"
	DeliveryMethodSms   DeliveryMethod = "SMS"
)

// Values returns all known values for DeliveryMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeliveryMethod) Values() []DeliveryMethod {
	return []DeliveryMethod{
		"EMAIL",
		"SMS",
	}
}

type MFAMode string

// Enum values for MFAMode
const (
	MFAModeOn       MFAMode = "ON"
	MFAModeOff      MFAMode = "OFF"
	MFAModeOptional MFAMode = "OPTIONAL"
)

// Values returns all known values for MFAMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (MFAMode) Values() []MFAMode {
	return []MFAMode{
		"ON",
		"OFF",
		"OPTIONAL",
	}
}

type MfaTypesElement string

// Enum values for MfaTypesElement
const (
	MfaTypesElementSms  MfaTypesElement = "SMS"
	MfaTypesElementTotp MfaTypesElement = "TOTP"
)

// Values returns all known values for MfaTypesElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MfaTypesElement) Values() []MfaTypesElement {
	return []MfaTypesElement{
		"SMS",
		"TOTP",
	}
}

type Mode string

// Enum values for Mode
const (
	ModeApiKey                 Mode = "API_KEY"
	ModeAwsIam                 Mode = "AWS_IAM"
	ModeAmazonCognitoUserPools Mode = "AMAZON_COGNITO_USER_POOLS"
	ModeOpenidConnect          Mode = "OPENID_CONNECT"
)

// Values returns all known values for Mode. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Mode) Values() []Mode {
	return []Mode{
		"API_KEY",
		"AWS_IAM",
		"AMAZON_COGNITO_USER_POOLS",
		"OPENID_CONNECT",
	}
}

type OAuthGrantType string

// Enum values for OAuthGrantType
const (
	OAuthGrantTypeCode     OAuthGrantType = "CODE"
	OAuthGrantTypeImplicit OAuthGrantType = "IMPLICIT"
)

// Values returns all known values for OAuthGrantType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OAuthGrantType) Values() []OAuthGrantType {
	return []OAuthGrantType{
		"CODE",
		"IMPLICIT",
	}
}

type OAuthScopesElement string

// Enum values for OAuthScopesElement
const (
	OAuthScopesElementPhone                     OAuthScopesElement = "PHONE"
	OAuthScopesElementEmail                     OAuthScopesElement = "EMAIL"
	OAuthScopesElementOpenid                    OAuthScopesElement = "OPENID"
	OAuthScopesElementProfile                   OAuthScopesElement = "PROFILE"
	OAuthScopesElementAwsCognitoSigninUserAdmin OAuthScopesElement = "AWS_COGNITO_SIGNIN_USER_ADMIN"
)

// Values returns all known values for OAuthScopesElement. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OAuthScopesElement) Values() []OAuthScopesElement {
	return []OAuthScopesElement{
		"PHONE",
		"EMAIL",
		"OPENID",
		"PROFILE",
		"AWS_COGNITO_SIGNIN_USER_ADMIN",
	}
}

type RequiredSignUpAttributesElement string

// Enum values for RequiredSignUpAttributesElement
const (
	RequiredSignUpAttributesElementAddress           RequiredSignUpAttributesElement = "ADDRESS"
	RequiredSignUpAttributesElementBirthdate         RequiredSignUpAttributesElement = "BIRTHDATE"
	RequiredSignUpAttributesElementEmail             RequiredSignUpAttributesElement = "EMAIL"
	RequiredSignUpAttributesElementFamilyName        RequiredSignUpAttributesElement = "FAMILY_NAME"
	RequiredSignUpAttributesElementGender            RequiredSignUpAttributesElement = "GENDER"
	RequiredSignUpAttributesElementGivenName         RequiredSignUpAttributesElement = "GIVEN_NAME"
	RequiredSignUpAttributesElementLocale            RequiredSignUpAttributesElement = "LOCALE"
	RequiredSignUpAttributesElementMiddleName        RequiredSignUpAttributesElement = "MIDDLE_NAME"
	RequiredSignUpAttributesElementName              RequiredSignUpAttributesElement = "NAME"
	RequiredSignUpAttributesElementNickname          RequiredSignUpAttributesElement = "NICKNAME"
	RequiredSignUpAttributesElementPhoneNumber       RequiredSignUpAttributesElement = "PHONE_NUMBER"
	RequiredSignUpAttributesElementPicture           RequiredSignUpAttributesElement = "PICTURE"
	RequiredSignUpAttributesElementPreferredUsername RequiredSignUpAttributesElement = "PREFERRED_USERNAME"
	RequiredSignUpAttributesElementProfile           RequiredSignUpAttributesElement = "PROFILE"
	RequiredSignUpAttributesElementUpdatedAt         RequiredSignUpAttributesElement = "UPDATED_AT"
	RequiredSignUpAttributesElementWebsite           RequiredSignUpAttributesElement = "WEBSITE"
	RequiredSignUpAttributesElementZoneInfo          RequiredSignUpAttributesElement = "ZONE_INFO"
)

// Values returns all known values for RequiredSignUpAttributesElement. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RequiredSignUpAttributesElement) Values() []RequiredSignUpAttributesElement {
	return []RequiredSignUpAttributesElement{
		"ADDRESS",
		"BIRTHDATE",
		"EMAIL",
		"FAMILY_NAME",
		"GENDER",
		"GIVEN_NAME",
		"LOCALE",
		"MIDDLE_NAME",
		"NAME",
		"NICKNAME",
		"PHONE_NUMBER",
		"PICTURE",
		"PREFERRED_USERNAME",
		"PROFILE",
		"UPDATED_AT",
		"WEBSITE",
		"ZONE_INFO",
	}
}

type ResolutionStrategy string

// Enum values for ResolutionStrategy
const (
	ResolutionStrategyOptimisticConcurrency ResolutionStrategy = "OPTIMISTIC_CONCURRENCY"
	ResolutionStrategyLambda                ResolutionStrategy = "LAMBDA"
	ResolutionStrategyAutomerge             ResolutionStrategy = "AUTOMERGE"
	ResolutionStrategyNone                  ResolutionStrategy = "NONE"
)

// Values returns all known values for ResolutionStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResolutionStrategy) Values() []ResolutionStrategy {
	return []ResolutionStrategy{
		"OPTIMISTIC_CONCURRENCY",
		"LAMBDA",
		"AUTOMERGE",
		"NONE",
	}
}

type Service string

// Enum values for Service
const (
	ServiceCognito Service = "COGNITO"
)

// Values returns all known values for Service. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Service) Values() []Service {
	return []Service{
		"COGNITO",
	}
}

type ServiceName string

// Enum values for ServiceName
const (
	ServiceNameS3 ServiceName = "S3"
)

// Values returns all known values for ServiceName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServiceName) Values() []ServiceName {
	return []ServiceName{
		"S3",
	}
}

type SignInMethod string

// Enum values for SignInMethod
const (
	SignInMethodEmail               SignInMethod = "EMAIL"
	SignInMethodEmailAndPhoneNumber SignInMethod = "EMAIL_AND_PHONE_NUMBER"
	SignInMethodPhoneNumber         SignInMethod = "PHONE_NUMBER"
	SignInMethodUsername            SignInMethod = "USERNAME"
)

// Values returns all known values for SignInMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SignInMethod) Values() []SignInMethod {
	return []SignInMethod{
		"EMAIL",
		"EMAIL_AND_PHONE_NUMBER",
		"PHONE_NUMBER",
		"USERNAME",
	}
}

type Status string

// Enum values for Status
const (
	StatusLatest Status = "LATEST"
	StatusStale  Status = "STALE"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"LATEST",
		"STALE",
	}
}

type UnAuthenticatedElement string

// Enum values for UnAuthenticatedElement
const (
	UnAuthenticatedElementRead            UnAuthenticatedElement = "READ"
	UnAuthenticatedElementCreateAndUpdate UnAuthenticatedElement = "CREATE_AND_UPDATE"
	UnAuthenticatedElementDelete          UnAuthenticatedElement = "DELETE"
)

// Values returns all known values for UnAuthenticatedElement. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UnAuthenticatedElement) Values() []UnAuthenticatedElement {
	return []UnAuthenticatedElement{
		"READ",
		"CREATE_AND_UPDATE",
		"DELETE",
	}
}
