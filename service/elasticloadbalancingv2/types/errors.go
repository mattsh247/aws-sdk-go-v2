// Code generated by smithy-go-codegen DO NOT EDIT.


package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified allocation ID does not exist.
type AllocationIdNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AllocationIdNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AllocationIdNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AllocationIdNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AllocationIdNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *AllocationIdNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified ALPN policy is not supported.
type ALPNPolicyNotSupportedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ALPNPolicyNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ALPNPolicyNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ALPNPolicyNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ALPNPolicyNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ALPNPolicyNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified Availability Zone is not supported.
type AvailabilityZoneNotSupportedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *AvailabilityZoneNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AvailabilityZoneNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AvailabilityZoneNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AvailabilityZoneNotSupported"
	}
	return *e.ErrorCodeOverride
}
func (e *AvailabilityZoneNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified certificate does not exist.
type CertificateNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *CertificateNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "CertificateNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *CertificateNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A listener with the specified port already exists.
type DuplicateListenerException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *DuplicateListenerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateListenerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateListenerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateListener"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateListenerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A load balancer with the specified name already exists.
type DuplicateLoadBalancerNameException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *DuplicateLoadBalancerNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateLoadBalancerNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateLoadBalancerNameException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateLoadBalancerName"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateLoadBalancerNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A tag key was specified more than once.
type DuplicateTagKeysException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *DuplicateTagKeysException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateTagKeysException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateTagKeysException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateTagKeys"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateTagKeysException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A target group with the specified name already exists.
type DuplicateTargetGroupNameException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *DuplicateTargetGroupNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateTargetGroupNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateTargetGroupNameException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateTargetGroupName"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateTargetGroupNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The health of the specified targets could not be retrieved due to an internal
// error.
type HealthUnavailableException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *HealthUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HealthUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HealthUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "HealthUnavailable"
	}
	return *e.ErrorCodeOverride
}
func (e *HealthUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The specified configuration is not valid with this protocol.
type IncompatibleProtocolsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *IncompatibleProtocolsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleProtocolsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleProtocolsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IncompatibleProtocols"
	}
	return *e.ErrorCodeOverride
}
func (e *IncompatibleProtocolsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested configuration is not valid.
type InvalidConfigurationRequestException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidConfigurationRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidConfigurationRequest"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidConfigurationRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested action is not valid.
type InvalidLoadBalancerActionException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidLoadBalancerActionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLoadBalancerActionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLoadBalancerActionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidLoadBalancerAction"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidLoadBalancerActionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested scheme is not valid.
type InvalidSchemeException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSchemeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSchemeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSchemeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidScheme"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSchemeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified security group does not exist.
type InvalidSecurityGroupException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSecurityGroupException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSecurityGroupException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSecurityGroupException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSecurityGroup"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSecurityGroupException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified subnet is out of available addresses.
type InvalidSubnetException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidSubnetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSubnetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSubnetException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSubnet"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSubnetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified target does not exist, is not in the same VPC as the target
// group, or has an unsupported instance type.
type InvalidTargetException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *InvalidTargetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTargetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTargetException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTarget"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTargetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified listener does not exist.
type ListenerNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ListenerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListenerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListenerNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ListenerNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *ListenerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified load balancer does not exist.
type LoadBalancerNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *LoadBalancerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LoadBalancerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LoadBalancerNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LoadBalancerNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *LoadBalancerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This operation is not allowed.
type OperationNotPermittedException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *OperationNotPermittedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationNotPermittedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationNotPermittedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "OperationNotPermitted"
	}
	return *e.ErrorCodeOverride
}
func (e *OperationNotPermittedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified priority is in use.
type PriorityInUseException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *PriorityInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PriorityInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PriorityInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PriorityInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *PriorityInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A specified resource is in use.
type ResourceInUseException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceInUse"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified rule does not exist.
type RuleNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *RuleNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RuleNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RuleNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "RuleNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *RuleNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified SSL policy does not exist.
type SSLPolicyNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SSLPolicyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SSLPolicyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SSLPolicyNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SSLPolicyNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *SSLPolicyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified subnet does not exist.
type SubnetNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *SubnetNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SubnetNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SubnetNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SubnetNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *SubnetNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of load balancers per target group.
type TargetGroupAssociationLimitException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TargetGroupAssociationLimitException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TargetGroupAssociationLimitException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TargetGroupAssociationLimitException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TargetGroupAssociationLimit"
	}
	return *e.ErrorCodeOverride
}
func (e *TargetGroupAssociationLimitException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified target group does not exist.
type TargetGroupNotFoundException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TargetGroupNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TargetGroupNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TargetGroupNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TargetGroupNotFound"
	}
	return *e.ErrorCodeOverride
}
func (e *TargetGroupNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of actions per rule.
type TooManyActionsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyActionsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyActionsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyActionsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyActions"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyActionsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of certificates per load balancer.
type TooManyCertificatesException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyCertificatesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyCertificatesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyCertificatesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyCertificates"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyCertificatesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of listeners per load balancer.
type TooManyListenersException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyListenersException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyListenersException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyListenersException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyListeners"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyListenersException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of load balancers for your Amazon Web
// Services account.
type TooManyLoadBalancersException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyLoadBalancersException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyLoadBalancersException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyLoadBalancersException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyLoadBalancers"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyLoadBalancersException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of times a target can be registered with
// a load balancer.
type TooManyRegistrationsForTargetIdException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyRegistrationsForTargetIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRegistrationsForTargetIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRegistrationsForTargetIdException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyRegistrationsForTargetId"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRegistrationsForTargetIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of rules per load balancer.
type TooManyRulesException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyRulesException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRulesException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRulesException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyRules"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyRulesException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of tags per load balancer.
type TooManyTagsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTags"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of target groups for your Amazon Web
// Services account.
type TooManyTargetGroupsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyTargetGroupsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTargetGroupsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTargetGroupsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTargetGroups"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTargetGroupsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of targets.
type TooManyTargetsException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyTargetsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTargetsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTargetsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyTargets"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyTargetsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit on the number of unique target groups per load
// balancer across all listeners. If a target group is used by multiple actions for
// a load balancer, it is counted as only one use.
type TooManyUniqueTargetGroupsPerLoadBalancerException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *TooManyUniqueTargetGroupsPerLoadBalancerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyUniqueTargetGroupsPerLoadBalancerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyUniqueTargetGroupsPerLoadBalancerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TooManyUniqueTargetGroupsPerLoadBalancer"
	}
	return *e.ErrorCodeOverride
}
func (e *TooManyUniqueTargetGroupsPerLoadBalancerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified protocol is not supported.
type UnsupportedProtocolException struct {
	Message *string
	
	ErrorCodeOverride *string
	
	noSmithyDocumentSerde
}

func (e *UnsupportedProtocolException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedProtocolException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedProtocolException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedProtocol"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedProtocolException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
