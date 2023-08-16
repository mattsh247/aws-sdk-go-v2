// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type CachePolicyCookieBehavior string

// Enum values for CachePolicyCookieBehavior
const (
	CachePolicyCookieBehaviorNone CachePolicyCookieBehavior = "none"
	CachePolicyCookieBehaviorWhitelist CachePolicyCookieBehavior = "whitelist"
	CachePolicyCookieBehaviorAllExcept CachePolicyCookieBehavior = "allExcept"
	CachePolicyCookieBehaviorAll CachePolicyCookieBehavior = "all"
)

// Values returns all known values for CachePolicyCookieBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CachePolicyCookieBehavior) Values() []CachePolicyCookieBehavior {
	return []CachePolicyCookieBehavior{
		"none",
		"whitelist",
		"allExcept",
		"all",
	}
}

type CachePolicyHeaderBehavior string

// Enum values for CachePolicyHeaderBehavior
const (
	CachePolicyHeaderBehaviorNone CachePolicyHeaderBehavior = "none"
	CachePolicyHeaderBehaviorWhitelist CachePolicyHeaderBehavior = "whitelist"
)

// Values returns all known values for CachePolicyHeaderBehavior. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CachePolicyHeaderBehavior) Values() []CachePolicyHeaderBehavior {
	return []CachePolicyHeaderBehavior{
		"none",
		"whitelist",
	}
}

type CachePolicyQueryStringBehavior string

// Enum values for CachePolicyQueryStringBehavior
const (
	CachePolicyQueryStringBehaviorNone CachePolicyQueryStringBehavior = "none"
	CachePolicyQueryStringBehaviorWhitelist CachePolicyQueryStringBehavior = "whitelist"
	CachePolicyQueryStringBehaviorAllExcept CachePolicyQueryStringBehavior = "allExcept"
	CachePolicyQueryStringBehaviorAll CachePolicyQueryStringBehavior = "all"
)

// Values returns all known values for CachePolicyQueryStringBehavior. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CachePolicyQueryStringBehavior) Values() []CachePolicyQueryStringBehavior {
	return []CachePolicyQueryStringBehavior{
		"none",
		"whitelist",
		"allExcept",
		"all",
	}
}

type CachePolicyType string

// Enum values for CachePolicyType
const (
	CachePolicyTypeManaged CachePolicyType = "managed"
	CachePolicyTypeCustom CachePolicyType = "custom"
)

// Values returns all known values for CachePolicyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CachePolicyType) Values() []CachePolicyType {
	return []CachePolicyType{
		"managed",
		"custom",
	}
}

type CertificateSource string

// Enum values for CertificateSource
const (
	CertificateSourceCloudfront CertificateSource = "cloudfront"
	CertificateSourceIam CertificateSource = "iam"
	CertificateSourceAcm CertificateSource = "acm"
)

// Values returns all known values for CertificateSource. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CertificateSource) Values() []CertificateSource {
	return []CertificateSource{
		"cloudfront",
		"iam",
		"acm",
	}
}

type ContinuousDeploymentPolicyType string

// Enum values for ContinuousDeploymentPolicyType
const (
	ContinuousDeploymentPolicyTypeSingleWeight ContinuousDeploymentPolicyType = "SingleWeight"
	ContinuousDeploymentPolicyTypeSingleHeader ContinuousDeploymentPolicyType = "SingleHeader"
)

// Values returns all known values for ContinuousDeploymentPolicyType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ContinuousDeploymentPolicyType) Values() []ContinuousDeploymentPolicyType {
	return []ContinuousDeploymentPolicyType{
		"SingleWeight",
		"SingleHeader",
	}
}

type EventType string

// Enum values for EventType
const (
	EventTypeViewerRequest EventType = "viewer-request"
	EventTypeViewerResponse EventType = "viewer-response"
	EventTypeOriginRequest EventType = "origin-request"
	EventTypeOriginResponse EventType = "origin-response"
)

// Values returns all known values for EventType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EventType) Values() []EventType {
	return []EventType{
		"viewer-request",
		"viewer-response",
		"origin-request",
		"origin-response",
	}
}

type Format string

// Enum values for Format
const (
	FormatURLEncoded Format = "URLEncoded"
)

// Values returns all known values for Format. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Format) Values() []Format {
	return []Format{
		"URLEncoded",
	}
}

type FrameOptionsList string

// Enum values for FrameOptionsList
const (
	FrameOptionsListDeny FrameOptionsList = "DENY"
	FrameOptionsListSameorigin FrameOptionsList = "SAMEORIGIN"
)

// Values returns all known values for FrameOptionsList. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FrameOptionsList) Values() []FrameOptionsList {
	return []FrameOptionsList{
		"DENY",
		"SAMEORIGIN",
	}
}

type FunctionRuntime string

// Enum values for FunctionRuntime
const (
	FunctionRuntimeCloudfrontJs10 FunctionRuntime = "cloudfront-js-1.0"
	FunctionRuntimeCloudfrontJs20 FunctionRuntime = "cloudfront-js-2.0"
)

// Values returns all known values for FunctionRuntime. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FunctionRuntime) Values() []FunctionRuntime {
	return []FunctionRuntime{
		"cloudfront-js-1.0",
		"cloudfront-js-2.0",
	}
}

type FunctionStage string

// Enum values for FunctionStage
const (
	FunctionStageDevelopment FunctionStage = "DEVELOPMENT"
	FunctionStageLive FunctionStage = "LIVE"
)

// Values returns all known values for FunctionStage. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FunctionStage) Values() []FunctionStage {
	return []FunctionStage{
		"DEVELOPMENT",
		"LIVE",
	}
}

type GeoRestrictionType string

// Enum values for GeoRestrictionType
const (
	GeoRestrictionTypeBlacklist GeoRestrictionType = "blacklist"
	GeoRestrictionTypeWhitelist GeoRestrictionType = "whitelist"
	GeoRestrictionTypeNone GeoRestrictionType = "none"
)

// Values returns all known values for GeoRestrictionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (GeoRestrictionType) Values() []GeoRestrictionType {
	return []GeoRestrictionType{
		"blacklist",
		"whitelist",
		"none",
	}
}

type HttpVersion string

// Enum values for HttpVersion
const (
	HttpVersionHttp11 HttpVersion = "http1.1"
	HttpVersionHttp2 HttpVersion = "http2"
	HttpVersionHttp3 HttpVersion = "http3"
	HttpVersionHttp2and3 HttpVersion = "http2and3"
)

// Values returns all known values for HttpVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HttpVersion) Values() []HttpVersion {
	return []HttpVersion{
		"http1.1",
		"http2",
		"http3",
		"http2and3",
	}
}

type ICPRecordalStatus string

// Enum values for ICPRecordalStatus
const (
	ICPRecordalStatusApproved ICPRecordalStatus = "APPROVED"
	ICPRecordalStatusSuspended ICPRecordalStatus = "SUSPENDED"
	ICPRecordalStatusPending ICPRecordalStatus = "PENDING"
)

// Values returns all known values for ICPRecordalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ICPRecordalStatus) Values() []ICPRecordalStatus {
	return []ICPRecordalStatus{
		"APPROVED",
		"SUSPENDED",
		"PENDING",
	}
}

type ItemSelection string

// Enum values for ItemSelection
const (
	ItemSelectionNone ItemSelection = "none"
	ItemSelectionWhitelist ItemSelection = "whitelist"
	ItemSelectionAll ItemSelection = "all"
)

// Values returns all known values for ItemSelection. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ItemSelection) Values() []ItemSelection {
	return []ItemSelection{
		"none",
		"whitelist",
		"all",
	}
}

type Method string

// Enum values for Method
const (
	MethodGet Method = "GET"
	MethodHead Method = "HEAD"
	MethodPost Method = "POST"
	MethodPut Method = "PUT"
	MethodPatch Method = "PATCH"
	MethodOptions Method = "OPTIONS"
	MethodDelete Method = "DELETE"
)

// Values returns all known values for Method. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Method) Values() []Method {
	return []Method{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"PATCH",
		"OPTIONS",
		"DELETE",
	}
}

type MinimumProtocolVersion string

// Enum values for MinimumProtocolVersion
const (
	MinimumProtocolVersionSSLv3 MinimumProtocolVersion = "SSLv3"
	MinimumProtocolVersionTLSv1 MinimumProtocolVersion = "TLSv1"
	MinimumProtocolVersionTLSv12016 MinimumProtocolVersion = "TLSv1_2016"
	MinimumProtocolVersionTLSv112016 MinimumProtocolVersion = "TLSv1.1_2016"
	MinimumProtocolVersionTLSv122018 MinimumProtocolVersion = "TLSv1.2_2018"
	MinimumProtocolVersionTLSv122019 MinimumProtocolVersion = "TLSv1.2_2019"
	MinimumProtocolVersionTLSv122021 MinimumProtocolVersion = "TLSv1.2_2021"
)

// Values returns all known values for MinimumProtocolVersion. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MinimumProtocolVersion) Values() []MinimumProtocolVersion {
	return []MinimumProtocolVersion{
		"SSLv3",
		"TLSv1",
		"TLSv1_2016",
		"TLSv1.1_2016",
		"TLSv1.2_2018",
		"TLSv1.2_2019",
		"TLSv1.2_2021",
	}
}

type OriginAccessControlOriginTypes string

// Enum values for OriginAccessControlOriginTypes
const (
	OriginAccessControlOriginTypesS3 OriginAccessControlOriginTypes = "s3"
	OriginAccessControlOriginTypesMediastore OriginAccessControlOriginTypes = "mediastore"
)

// Values returns all known values for OriginAccessControlOriginTypes. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginAccessControlOriginTypes) Values() []OriginAccessControlOriginTypes {
	return []OriginAccessControlOriginTypes{
		"s3",
		"mediastore",
	}
}

type OriginAccessControlSigningBehaviors string

// Enum values for OriginAccessControlSigningBehaviors
const (
	OriginAccessControlSigningBehaviorsNever OriginAccessControlSigningBehaviors = "never"
	OriginAccessControlSigningBehaviorsAlways OriginAccessControlSigningBehaviors = "always"
	OriginAccessControlSigningBehaviorsNoOverride OriginAccessControlSigningBehaviors = "no-override"
)

// Values returns all known values for OriginAccessControlSigningBehaviors. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginAccessControlSigningBehaviors) Values() []OriginAccessControlSigningBehaviors {
	return []OriginAccessControlSigningBehaviors{
		"never",
		"always",
		"no-override",
	}
}

type OriginAccessControlSigningProtocols string

// Enum values for OriginAccessControlSigningProtocols
const (
	OriginAccessControlSigningProtocolsSigv4 OriginAccessControlSigningProtocols = "sigv4"
)

// Values returns all known values for OriginAccessControlSigningProtocols. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginAccessControlSigningProtocols) Values() []OriginAccessControlSigningProtocols {
	return []OriginAccessControlSigningProtocols{
		"sigv4",
	}
}

type OriginProtocolPolicy string

// Enum values for OriginProtocolPolicy
const (
	OriginProtocolPolicyHttpOnly OriginProtocolPolicy = "http-only"
	OriginProtocolPolicyMatchViewer OriginProtocolPolicy = "match-viewer"
	OriginProtocolPolicyHttpsOnly OriginProtocolPolicy = "https-only"
)

// Values returns all known values for OriginProtocolPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginProtocolPolicy) Values() []OriginProtocolPolicy {
	return []OriginProtocolPolicy{
		"http-only",
		"match-viewer",
		"https-only",
	}
}

type OriginRequestPolicyCookieBehavior string

// Enum values for OriginRequestPolicyCookieBehavior
const (
	OriginRequestPolicyCookieBehaviorNone OriginRequestPolicyCookieBehavior = "none"
	OriginRequestPolicyCookieBehaviorWhitelist OriginRequestPolicyCookieBehavior = "whitelist"
	OriginRequestPolicyCookieBehaviorAll OriginRequestPolicyCookieBehavior = "all"
	OriginRequestPolicyCookieBehaviorAllExcept OriginRequestPolicyCookieBehavior = "allExcept"
)

// Values returns all known values for OriginRequestPolicyCookieBehavior. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginRequestPolicyCookieBehavior) Values() []OriginRequestPolicyCookieBehavior {
	return []OriginRequestPolicyCookieBehavior{
		"none",
		"whitelist",
		"all",
		"allExcept",
	}
}

type OriginRequestPolicyHeaderBehavior string

// Enum values for OriginRequestPolicyHeaderBehavior
const (
	OriginRequestPolicyHeaderBehaviorNone OriginRequestPolicyHeaderBehavior = "none"
	OriginRequestPolicyHeaderBehaviorWhitelist OriginRequestPolicyHeaderBehavior = "whitelist"
	OriginRequestPolicyHeaderBehaviorAllViewer OriginRequestPolicyHeaderBehavior = "allViewer"
	OriginRequestPolicyHeaderBehaviorAllViewerAndWhitelistCloudFront OriginRequestPolicyHeaderBehavior = "allViewerAndWhitelistCloudFront"
	OriginRequestPolicyHeaderBehaviorAllExcept OriginRequestPolicyHeaderBehavior = "allExcept"
)

// Values returns all known values for OriginRequestPolicyHeaderBehavior. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginRequestPolicyHeaderBehavior) Values() []OriginRequestPolicyHeaderBehavior {
	return []OriginRequestPolicyHeaderBehavior{
		"none",
		"whitelist",
		"allViewer",
		"allViewerAndWhitelistCloudFront",
		"allExcept",
	}
}

type OriginRequestPolicyQueryStringBehavior string

// Enum values for OriginRequestPolicyQueryStringBehavior
const (
	OriginRequestPolicyQueryStringBehaviorNone OriginRequestPolicyQueryStringBehavior = "none"
	OriginRequestPolicyQueryStringBehaviorWhitelist OriginRequestPolicyQueryStringBehavior = "whitelist"
	OriginRequestPolicyQueryStringBehaviorAll OriginRequestPolicyQueryStringBehavior = "all"
	OriginRequestPolicyQueryStringBehaviorAllExcept OriginRequestPolicyQueryStringBehavior = "allExcept"
)

// Values returns all known values for OriginRequestPolicyQueryStringBehavior.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (OriginRequestPolicyQueryStringBehavior) Values() []OriginRequestPolicyQueryStringBehavior {
	return []OriginRequestPolicyQueryStringBehavior{
		"none",
		"whitelist",
		"all",
		"allExcept",
	}
}

type OriginRequestPolicyType string

// Enum values for OriginRequestPolicyType
const (
	OriginRequestPolicyTypeManaged OriginRequestPolicyType = "managed"
	OriginRequestPolicyTypeCustom OriginRequestPolicyType = "custom"
)

// Values returns all known values for OriginRequestPolicyType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OriginRequestPolicyType) Values() []OriginRequestPolicyType {
	return []OriginRequestPolicyType{
		"managed",
		"custom",
	}
}

type PriceClass string

// Enum values for PriceClass
const (
	PriceClassPriceClass100 PriceClass = "PriceClass_100"
	PriceClassPriceClass200 PriceClass = "PriceClass_200"
	PriceClassPriceClassAll PriceClass = "PriceClass_All"
)

// Values returns all known values for PriceClass. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PriceClass) Values() []PriceClass {
	return []PriceClass{
		"PriceClass_100",
		"PriceClass_200",
		"PriceClass_All",
	}
}

type RealtimeMetricsSubscriptionStatus string

// Enum values for RealtimeMetricsSubscriptionStatus
const (
	RealtimeMetricsSubscriptionStatusEnabled RealtimeMetricsSubscriptionStatus = "Enabled"
	RealtimeMetricsSubscriptionStatusDisabled RealtimeMetricsSubscriptionStatus = "Disabled"
)

// Values returns all known values for RealtimeMetricsSubscriptionStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RealtimeMetricsSubscriptionStatus) Values() []RealtimeMetricsSubscriptionStatus {
	return []RealtimeMetricsSubscriptionStatus{
		"Enabled",
		"Disabled",
	}
}

type ReferrerPolicyList string

// Enum values for ReferrerPolicyList
const (
	ReferrerPolicyListNoReferrer ReferrerPolicyList = "no-referrer"
	ReferrerPolicyListNoReferrerWhenDowngrade ReferrerPolicyList = "no-referrer-when-downgrade"
	ReferrerPolicyListOrigin ReferrerPolicyList = "origin"
	ReferrerPolicyListOriginWhenCrossOrigin ReferrerPolicyList = "origin-when-cross-origin"
	ReferrerPolicyListSameOrigin ReferrerPolicyList = "same-origin"
	ReferrerPolicyListStrictOrigin ReferrerPolicyList = "strict-origin"
	ReferrerPolicyListStrictOriginWhenCrossOrigin ReferrerPolicyList = "strict-origin-when-cross-origin"
	ReferrerPolicyListUnsafeUrl ReferrerPolicyList = "unsafe-url"
)

// Values returns all known values for ReferrerPolicyList. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReferrerPolicyList) Values() []ReferrerPolicyList {
	return []ReferrerPolicyList{
		"no-referrer",
		"no-referrer-when-downgrade",
		"origin",
		"origin-when-cross-origin",
		"same-origin",
		"strict-origin",
		"strict-origin-when-cross-origin",
		"unsafe-url",
	}
}

type ResponseHeadersPolicyAccessControlAllowMethodsValues string

// Enum values for ResponseHeadersPolicyAccessControlAllowMethodsValues
const (
	ResponseHeadersPolicyAccessControlAllowMethodsValuesGet ResponseHeadersPolicyAccessControlAllowMethodsValues = "GET"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesPost ResponseHeadersPolicyAccessControlAllowMethodsValues = "POST"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesOptions ResponseHeadersPolicyAccessControlAllowMethodsValues = "OPTIONS"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesPut ResponseHeadersPolicyAccessControlAllowMethodsValues = "PUT"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesDelete ResponseHeadersPolicyAccessControlAllowMethodsValues = "DELETE"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesPatch ResponseHeadersPolicyAccessControlAllowMethodsValues = "PATCH"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesHead ResponseHeadersPolicyAccessControlAllowMethodsValues = "HEAD"
	ResponseHeadersPolicyAccessControlAllowMethodsValuesAll ResponseHeadersPolicyAccessControlAllowMethodsValues = "ALL"
)

// Values returns all known values for
// ResponseHeadersPolicyAccessControlAllowMethodsValues. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResponseHeadersPolicyAccessControlAllowMethodsValues) Values() []ResponseHeadersPolicyAccessControlAllowMethodsValues {
	return []ResponseHeadersPolicyAccessControlAllowMethodsValues{
		"GET",
		"POST",
		"OPTIONS",
		"PUT",
		"DELETE",
		"PATCH",
		"HEAD",
		"ALL",
	}
}

type ResponseHeadersPolicyType string

// Enum values for ResponseHeadersPolicyType
const (
	ResponseHeadersPolicyTypeManaged ResponseHeadersPolicyType = "managed"
	ResponseHeadersPolicyTypeCustom ResponseHeadersPolicyType = "custom"
)

// Values returns all known values for ResponseHeadersPolicyType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ResponseHeadersPolicyType) Values() []ResponseHeadersPolicyType {
	return []ResponseHeadersPolicyType{
		"managed",
		"custom",
	}
}

type SslProtocol string

// Enum values for SslProtocol
const (
	SslProtocolSSLv3 SslProtocol = "SSLv3"
	SslProtocolTLSv1 SslProtocol = "TLSv1"
	SslProtocolTLSv11 SslProtocol = "TLSv1.1"
	SslProtocolTLSv12 SslProtocol = "TLSv1.2"
)

// Values returns all known values for SslProtocol. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SslProtocol) Values() []SslProtocol {
	return []SslProtocol{
		"SSLv3",
		"TLSv1",
		"TLSv1.1",
		"TLSv1.2",
	}
}

type SSLSupportMethod string

// Enum values for SSLSupportMethod
const (
	SSLSupportMethodSniOnly SSLSupportMethod = "sni-only"
	SSLSupportMethodVip SSLSupportMethod = "vip"
	SSLSupportMethodStaticIp SSLSupportMethod = "static-ip"
)

// Values returns all known values for SSLSupportMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SSLSupportMethod) Values() []SSLSupportMethod {
	return []SSLSupportMethod{
		"sni-only",
		"vip",
		"static-ip",
	}
}

type ViewerProtocolPolicy string

// Enum values for ViewerProtocolPolicy
const (
	ViewerProtocolPolicyAllowAll ViewerProtocolPolicy = "allow-all"
	ViewerProtocolPolicyHttpsOnly ViewerProtocolPolicy = "https-only"
	ViewerProtocolPolicyRedirectToHttps ViewerProtocolPolicy = "redirect-to-https"
)

// Values returns all known values for ViewerProtocolPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ViewerProtocolPolicy) Values() []ViewerProtocolPolicy {
	return []ViewerProtocolPolicy{
		"allow-all",
		"https-only",
		"redirect-to-https",
	}
}
