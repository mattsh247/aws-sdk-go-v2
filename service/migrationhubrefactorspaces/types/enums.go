// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type ApiGatewayEndpointType string

// Enum values for ApiGatewayEndpointType
const (
	ApiGatewayEndpointTypeRegional ApiGatewayEndpointType = "REGIONAL"
	ApiGatewayEndpointTypePrivate ApiGatewayEndpointType = "PRIVATE"
)

// Values returns all known values for ApiGatewayEndpointType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApiGatewayEndpointType) Values() []ApiGatewayEndpointType {
	return []ApiGatewayEndpointType{
		"REGIONAL",
		"PRIVATE",
	}
}

type ApplicationState string

// Enum values for ApplicationState
const (
	ApplicationStateCreating ApplicationState = "CREATING"
	ApplicationStateActive ApplicationState = "ACTIVE"
	ApplicationStateDeleting ApplicationState = "DELETING"
	ApplicationStateFailed ApplicationState = "FAILED"
	ApplicationStateUpdating ApplicationState = "UPDATING"
)

// Values returns all known values for ApplicationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApplicationState) Values() []ApplicationState {
	return []ApplicationState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

type EnvironmentState string

// Enum values for EnvironmentState
const (
	EnvironmentStateCreating EnvironmentState = "CREATING"
	EnvironmentStateActive EnvironmentState = "ACTIVE"
	EnvironmentStateDeleting EnvironmentState = "DELETING"
	EnvironmentStateFailed EnvironmentState = "FAILED"
)

// Values returns all known values for EnvironmentState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentState) Values() []EnvironmentState {
	return []EnvironmentState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeInvalidResourceState ErrorCode = "INVALID_RESOURCE_STATE"
	ErrorCodeResourceLimitExceeded ErrorCode = "RESOURCE_LIMIT_EXCEEDED"
	ErrorCodeResourceCreationFailure ErrorCode = "RESOURCE_CREATION_FAILURE"
	ErrorCodeResourceUpdateFailure ErrorCode = "RESOURCE_UPDATE_FAILURE"
	ErrorCodeServiceEndpointHealthCheckFailure ErrorCode = "SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE"
	ErrorCodeResourceDeletionFailure ErrorCode = "RESOURCE_DELETION_FAILURE"
	ErrorCodeResourceRetrievalFailure ErrorCode = "RESOURCE_RETRIEVAL_FAILURE"
	ErrorCodeResourceInUse ErrorCode = "RESOURCE_IN_USE"
	ErrorCodeResourceNotFound ErrorCode = "RESOURCE_NOT_FOUND"
	ErrorCodeStateTransitionFailure ErrorCode = "STATE_TRANSITION_FAILURE"
	ErrorCodeRequestLimitExceeded ErrorCode = "REQUEST_LIMIT_EXCEEDED"
	ErrorCodeNotAuthorized ErrorCode = "NOT_AUTHORIZED"
)

// Values returns all known values for ErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"INVALID_RESOURCE_STATE",
		"RESOURCE_LIMIT_EXCEEDED",
		"RESOURCE_CREATION_FAILURE",
		"RESOURCE_UPDATE_FAILURE",
		"SERVICE_ENDPOINT_HEALTH_CHECK_FAILURE",
		"RESOURCE_DELETION_FAILURE",
		"RESOURCE_RETRIEVAL_FAILURE",
		"RESOURCE_IN_USE",
		"RESOURCE_NOT_FOUND",
		"STATE_TRANSITION_FAILURE",
		"REQUEST_LIMIT_EXCEEDED",
		"NOT_AUTHORIZED",
	}
}

type ErrorResourceType string

// Enum values for ErrorResourceType
const (
	ErrorResourceTypeEnvironment ErrorResourceType = "ENVIRONMENT"
	ErrorResourceTypeApplication ErrorResourceType = "APPLICATION"
	ErrorResourceTypeRoute ErrorResourceType = "ROUTE"
	ErrorResourceTypeService ErrorResourceType = "SERVICE"
	ErrorResourceTypeTransitGateway ErrorResourceType = "TRANSIT_GATEWAY"
	ErrorResourceTypeTransitGatewayAttachment ErrorResourceType = "TRANSIT_GATEWAY_ATTACHMENT"
	ErrorResourceTypeApiGateway ErrorResourceType = "API_GATEWAY"
	ErrorResourceTypeNlb ErrorResourceType = "NLB"
	ErrorResourceTypeTargetGroup ErrorResourceType = "TARGET_GROUP"
	ErrorResourceTypeLoadBalancerListener ErrorResourceType = "LOAD_BALANCER_LISTENER"
	ErrorResourceTypeVpcLink ErrorResourceType = "VPC_LINK"
	ErrorResourceTypeLambda ErrorResourceType = "LAMBDA"
	ErrorResourceTypeVpc ErrorResourceType = "VPC"
	ErrorResourceTypeSubnet ErrorResourceType = "SUBNET"
	ErrorResourceTypeRouteTable ErrorResourceType = "ROUTE_TABLE"
	ErrorResourceTypeSecurityGroup ErrorResourceType = "SECURITY_GROUP"
	ErrorResourceTypeVpcEndpointServiceConfiguration ErrorResourceType = "VPC_ENDPOINT_SERVICE_CONFIGURATION"
	ErrorResourceTypeResourceShare ErrorResourceType = "RESOURCE_SHARE"
	ErrorResourceTypeIamRole ErrorResourceType = "IAM_ROLE"
)

// Values returns all known values for ErrorResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ErrorResourceType) Values() []ErrorResourceType {
	return []ErrorResourceType{
		"ENVIRONMENT",
		"APPLICATION",
		"ROUTE",
		"SERVICE",
		"TRANSIT_GATEWAY",
		"TRANSIT_GATEWAY_ATTACHMENT",
		"API_GATEWAY",
		"NLB",
		"TARGET_GROUP",
		"LOAD_BALANCER_LISTENER",
		"VPC_LINK",
		"LAMBDA",
		"VPC",
		"SUBNET",
		"ROUTE_TABLE",
		"SECURITY_GROUP",
		"VPC_ENDPOINT_SERVICE_CONFIGURATION",
		"RESOURCE_SHARE",
		"IAM_ROLE",
	}
}

type HttpMethod string

// Enum values for HttpMethod
const (
	HttpMethodDelete HttpMethod = "DELETE"
	HttpMethodGet HttpMethod = "GET"
	HttpMethodHead HttpMethod = "HEAD"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodPatch HttpMethod = "PATCH"
	HttpMethodPost HttpMethod = "POST"
	HttpMethodPut HttpMethod = "PUT"
)

// Values returns all known values for HttpMethod. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (HttpMethod) Values() []HttpMethod {
	return []HttpMethod{
		"DELETE",
		"GET",
		"HEAD",
		"OPTIONS",
		"PATCH",
		"POST",
		"PUT",
	}
}

type NetworkFabricType string

// Enum values for NetworkFabricType
const (
	NetworkFabricTypeTransitGateway NetworkFabricType = "TRANSIT_GATEWAY"
	NetworkFabricTypeNone NetworkFabricType = "NONE"
)

// Values returns all known values for NetworkFabricType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NetworkFabricType) Values() []NetworkFabricType {
	return []NetworkFabricType{
		"TRANSIT_GATEWAY",
		"NONE",
	}
}

type ProxyType string

// Enum values for ProxyType
const (
	ProxyTypeApiGateway ProxyType = "API_GATEWAY"
)

// Values returns all known values for ProxyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ProxyType) Values() []ProxyType {
	return []ProxyType{
		"API_GATEWAY",
	}
}

type RouteActivationState string

// Enum values for RouteActivationState
const (
	RouteActivationStateActive RouteActivationState = "ACTIVE"
	RouteActivationStateInactive RouteActivationState = "INACTIVE"
)

// Values returns all known values for RouteActivationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RouteActivationState) Values() []RouteActivationState {
	return []RouteActivationState{
		"ACTIVE",
		"INACTIVE",
	}
}

type RouteState string

// Enum values for RouteState
const (
	RouteStateCreating RouteState = "CREATING"
	RouteStateActive RouteState = "ACTIVE"
	RouteStateDeleting RouteState = "DELETING"
	RouteStateFailed RouteState = "FAILED"
	RouteStateUpdating RouteState = "UPDATING"
	RouteStateInactive RouteState = "INACTIVE"
)

// Values returns all known values for RouteState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RouteState) Values() []RouteState {
	return []RouteState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"UPDATING",
		"INACTIVE",
	}
}

type RouteType string

// Enum values for RouteType
const (
	RouteTypeDefault RouteType = "DEFAULT"
	RouteTypeUriPath RouteType = "URI_PATH"
)

// Values returns all known values for RouteType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RouteType) Values() []RouteType {
	return []RouteType{
		"DEFAULT",
		"URI_PATH",
	}
}

type ServiceEndpointType string

// Enum values for ServiceEndpointType
const (
	ServiceEndpointTypeLambda ServiceEndpointType = "LAMBDA"
	ServiceEndpointTypeUrl ServiceEndpointType = "URL"
)

// Values returns all known values for ServiceEndpointType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceEndpointType) Values() []ServiceEndpointType {
	return []ServiceEndpointType{
		"LAMBDA",
		"URL",
	}
}

type ServiceState string

// Enum values for ServiceState
const (
	ServiceStateCreating ServiceState = "CREATING"
	ServiceStateActive ServiceState = "ACTIVE"
	ServiceStateDeleting ServiceState = "DELETING"
	ServiceStateFailed ServiceState = "FAILED"
)

// Values returns all known values for ServiceState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceState) Values() []ServiceState {
	return []ServiceState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
	}
}
