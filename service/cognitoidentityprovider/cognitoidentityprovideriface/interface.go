// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cognitoidentityprovideriface provides an interface to enable mocking the Amazon Cognito Identity Provider service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitoidentityprovideriface

import (
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

// CognitoIdentityProviderAPI provides an interface to enable mocking the
// cognitoidentityprovider.CognitoIdentityProvider service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Identity Provider.
//    func myFunc(svc cognitoidentityprovideriface.CognitoIdentityProviderAPI) bool {
//        // Make svc.AddCustomAttributes request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := cognitoidentityprovider.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoIdentityProviderClient struct {
//        cognitoidentityprovideriface.CognitoIdentityProviderAPI
//    }
//    func (m *mockCognitoIdentityProviderClient) AddCustomAttributes(input *cognitoidentityprovider.AddCustomAttributesInput) (*cognitoidentityprovider.AddCustomAttributesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoIdentityProviderClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CognitoIdentityProviderAPI interface {
	AddCustomAttributesRequest(*cognitoidentityprovider.AddCustomAttributesInput) cognitoidentityprovider.AddCustomAttributesRequest

	AdminAddUserToGroupRequest(*cognitoidentityprovider.AdminAddUserToGroupInput) cognitoidentityprovider.AdminAddUserToGroupRequest

	AdminConfirmSignUpRequest(*cognitoidentityprovider.AdminConfirmSignUpInput) cognitoidentityprovider.AdminConfirmSignUpRequest

	AdminCreateUserRequest(*cognitoidentityprovider.AdminCreateUserInput) cognitoidentityprovider.AdminCreateUserRequest

	AdminDeleteUserRequest(*cognitoidentityprovider.AdminDeleteUserInput) cognitoidentityprovider.AdminDeleteUserRequest

	AdminDeleteUserAttributesRequest(*cognitoidentityprovider.AdminDeleteUserAttributesInput) cognitoidentityprovider.AdminDeleteUserAttributesRequest

	AdminDisableProviderForUserRequest(*cognitoidentityprovider.AdminDisableProviderForUserInput) cognitoidentityprovider.AdminDisableProviderForUserRequest

	AdminDisableUserRequest(*cognitoidentityprovider.AdminDisableUserInput) cognitoidentityprovider.AdminDisableUserRequest

	AdminEnableUserRequest(*cognitoidentityprovider.AdminEnableUserInput) cognitoidentityprovider.AdminEnableUserRequest

	AdminForgetDeviceRequest(*cognitoidentityprovider.AdminForgetDeviceInput) cognitoidentityprovider.AdminForgetDeviceRequest

	AdminGetDeviceRequest(*cognitoidentityprovider.AdminGetDeviceInput) cognitoidentityprovider.AdminGetDeviceRequest

	AdminGetUserRequest(*cognitoidentityprovider.AdminGetUserInput) cognitoidentityprovider.AdminGetUserRequest

	AdminInitiateAuthRequest(*cognitoidentityprovider.AdminInitiateAuthInput) cognitoidentityprovider.AdminInitiateAuthRequest

	AdminLinkProviderForUserRequest(*cognitoidentityprovider.AdminLinkProviderForUserInput) cognitoidentityprovider.AdminLinkProviderForUserRequest

	AdminListDevicesRequest(*cognitoidentityprovider.AdminListDevicesInput) cognitoidentityprovider.AdminListDevicesRequest

	AdminListGroupsForUserRequest(*cognitoidentityprovider.AdminListGroupsForUserInput) cognitoidentityprovider.AdminListGroupsForUserRequest

	AdminListUserAuthEventsRequest(*cognitoidentityprovider.AdminListUserAuthEventsInput) cognitoidentityprovider.AdminListUserAuthEventsRequest

	AdminRemoveUserFromGroupRequest(*cognitoidentityprovider.AdminRemoveUserFromGroupInput) cognitoidentityprovider.AdminRemoveUserFromGroupRequest

	AdminResetUserPasswordRequest(*cognitoidentityprovider.AdminResetUserPasswordInput) cognitoidentityprovider.AdminResetUserPasswordRequest

	AdminRespondToAuthChallengeRequest(*cognitoidentityprovider.AdminRespondToAuthChallengeInput) cognitoidentityprovider.AdminRespondToAuthChallengeRequest

	AdminSetUserMFAPreferenceRequest(*cognitoidentityprovider.AdminSetUserMFAPreferenceInput) cognitoidentityprovider.AdminSetUserMFAPreferenceRequest

	AdminSetUserSettingsRequest(*cognitoidentityprovider.AdminSetUserSettingsInput) cognitoidentityprovider.AdminSetUserSettingsRequest

	AdminUpdateAuthEventFeedbackRequest(*cognitoidentityprovider.AdminUpdateAuthEventFeedbackInput) cognitoidentityprovider.AdminUpdateAuthEventFeedbackRequest

	AdminUpdateDeviceStatusRequest(*cognitoidentityprovider.AdminUpdateDeviceStatusInput) cognitoidentityprovider.AdminUpdateDeviceStatusRequest

	AdminUpdateUserAttributesRequest(*cognitoidentityprovider.AdminUpdateUserAttributesInput) cognitoidentityprovider.AdminUpdateUserAttributesRequest

	AdminUserGlobalSignOutRequest(*cognitoidentityprovider.AdminUserGlobalSignOutInput) cognitoidentityprovider.AdminUserGlobalSignOutRequest

	AssociateSoftwareTokenRequest(*cognitoidentityprovider.AssociateSoftwareTokenInput) cognitoidentityprovider.AssociateSoftwareTokenRequest

	ChangePasswordRequest(*cognitoidentityprovider.ChangePasswordInput) cognitoidentityprovider.ChangePasswordRequest

	ConfirmDeviceRequest(*cognitoidentityprovider.ConfirmDeviceInput) cognitoidentityprovider.ConfirmDeviceRequest

	ConfirmForgotPasswordRequest(*cognitoidentityprovider.ConfirmForgotPasswordInput) cognitoidentityprovider.ConfirmForgotPasswordRequest

	ConfirmSignUpRequest(*cognitoidentityprovider.ConfirmSignUpInput) cognitoidentityprovider.ConfirmSignUpRequest

	CreateGroupRequest(*cognitoidentityprovider.CreateGroupInput) cognitoidentityprovider.CreateGroupRequest

	CreateIdentityProviderRequest(*cognitoidentityprovider.CreateIdentityProviderInput) cognitoidentityprovider.CreateIdentityProviderRequest

	CreateResourceServerRequest(*cognitoidentityprovider.CreateResourceServerInput) cognitoidentityprovider.CreateResourceServerRequest

	CreateUserImportJobRequest(*cognitoidentityprovider.CreateUserImportJobInput) cognitoidentityprovider.CreateUserImportJobRequest

	CreateUserPoolRequest(*cognitoidentityprovider.CreateUserPoolInput) cognitoidentityprovider.CreateUserPoolRequest

	CreateUserPoolClientRequest(*cognitoidentityprovider.CreateUserPoolClientInput) cognitoidentityprovider.CreateUserPoolClientRequest

	CreateUserPoolDomainRequest(*cognitoidentityprovider.CreateUserPoolDomainInput) cognitoidentityprovider.CreateUserPoolDomainRequest

	DeleteGroupRequest(*cognitoidentityprovider.DeleteGroupInput) cognitoidentityprovider.DeleteGroupRequest

	DeleteIdentityProviderRequest(*cognitoidentityprovider.DeleteIdentityProviderInput) cognitoidentityprovider.DeleteIdentityProviderRequest

	DeleteResourceServerRequest(*cognitoidentityprovider.DeleteResourceServerInput) cognitoidentityprovider.DeleteResourceServerRequest

	DeleteUserRequest(*cognitoidentityprovider.DeleteUserInput) cognitoidentityprovider.DeleteUserRequest

	DeleteUserAttributesRequest(*cognitoidentityprovider.DeleteUserAttributesInput) cognitoidentityprovider.DeleteUserAttributesRequest

	DeleteUserPoolRequest(*cognitoidentityprovider.DeleteUserPoolInput) cognitoidentityprovider.DeleteUserPoolRequest

	DeleteUserPoolClientRequest(*cognitoidentityprovider.DeleteUserPoolClientInput) cognitoidentityprovider.DeleteUserPoolClientRequest

	DeleteUserPoolDomainRequest(*cognitoidentityprovider.DeleteUserPoolDomainInput) cognitoidentityprovider.DeleteUserPoolDomainRequest

	DescribeIdentityProviderRequest(*cognitoidentityprovider.DescribeIdentityProviderInput) cognitoidentityprovider.DescribeIdentityProviderRequest

	DescribeResourceServerRequest(*cognitoidentityprovider.DescribeResourceServerInput) cognitoidentityprovider.DescribeResourceServerRequest

	DescribeRiskConfigurationRequest(*cognitoidentityprovider.DescribeRiskConfigurationInput) cognitoidentityprovider.DescribeRiskConfigurationRequest

	DescribeUserImportJobRequest(*cognitoidentityprovider.DescribeUserImportJobInput) cognitoidentityprovider.DescribeUserImportJobRequest

	DescribeUserPoolRequest(*cognitoidentityprovider.DescribeUserPoolInput) cognitoidentityprovider.DescribeUserPoolRequest

	DescribeUserPoolClientRequest(*cognitoidentityprovider.DescribeUserPoolClientInput) cognitoidentityprovider.DescribeUserPoolClientRequest

	DescribeUserPoolDomainRequest(*cognitoidentityprovider.DescribeUserPoolDomainInput) cognitoidentityprovider.DescribeUserPoolDomainRequest

	ForgetDeviceRequest(*cognitoidentityprovider.ForgetDeviceInput) cognitoidentityprovider.ForgetDeviceRequest

	ForgotPasswordRequest(*cognitoidentityprovider.ForgotPasswordInput) cognitoidentityprovider.ForgotPasswordRequest

	GetCSVHeaderRequest(*cognitoidentityprovider.GetCSVHeaderInput) cognitoidentityprovider.GetCSVHeaderRequest

	GetDeviceRequest(*cognitoidentityprovider.GetDeviceInput) cognitoidentityprovider.GetDeviceRequest

	GetGroupRequest(*cognitoidentityprovider.GetGroupInput) cognitoidentityprovider.GetGroupRequest

	GetIdentityProviderByIdentifierRequest(*cognitoidentityprovider.GetIdentityProviderByIdentifierInput) cognitoidentityprovider.GetIdentityProviderByIdentifierRequest

	GetSigningCertificateRequest(*cognitoidentityprovider.GetSigningCertificateInput) cognitoidentityprovider.GetSigningCertificateRequest

	GetUICustomizationRequest(*cognitoidentityprovider.GetUICustomizationInput) cognitoidentityprovider.GetUICustomizationRequest

	GetUserRequest(*cognitoidentityprovider.GetUserInput) cognitoidentityprovider.GetUserRequest

	GetUserAttributeVerificationCodeRequest(*cognitoidentityprovider.GetUserAttributeVerificationCodeInput) cognitoidentityprovider.GetUserAttributeVerificationCodeRequest

	GetUserPoolMfaConfigRequest(*cognitoidentityprovider.GetUserPoolMfaConfigInput) cognitoidentityprovider.GetUserPoolMfaConfigRequest

	GlobalSignOutRequest(*cognitoidentityprovider.GlobalSignOutInput) cognitoidentityprovider.GlobalSignOutRequest

	InitiateAuthRequest(*cognitoidentityprovider.InitiateAuthInput) cognitoidentityprovider.InitiateAuthRequest

	ListDevicesRequest(*cognitoidentityprovider.ListDevicesInput) cognitoidentityprovider.ListDevicesRequest

	ListGroupsRequest(*cognitoidentityprovider.ListGroupsInput) cognitoidentityprovider.ListGroupsRequest

	ListIdentityProvidersRequest(*cognitoidentityprovider.ListIdentityProvidersInput) cognitoidentityprovider.ListIdentityProvidersRequest

	ListResourceServersRequest(*cognitoidentityprovider.ListResourceServersInput) cognitoidentityprovider.ListResourceServersRequest

	ListUserImportJobsRequest(*cognitoidentityprovider.ListUserImportJobsInput) cognitoidentityprovider.ListUserImportJobsRequest

	ListUserPoolClientsRequest(*cognitoidentityprovider.ListUserPoolClientsInput) cognitoidentityprovider.ListUserPoolClientsRequest

	ListUserPoolsRequest(*cognitoidentityprovider.ListUserPoolsInput) cognitoidentityprovider.ListUserPoolsRequest

	ListUsersRequest(*cognitoidentityprovider.ListUsersInput) cognitoidentityprovider.ListUsersRequest

	ListUsersInGroupRequest(*cognitoidentityprovider.ListUsersInGroupInput) cognitoidentityprovider.ListUsersInGroupRequest

	ResendConfirmationCodeRequest(*cognitoidentityprovider.ResendConfirmationCodeInput) cognitoidentityprovider.ResendConfirmationCodeRequest

	RespondToAuthChallengeRequest(*cognitoidentityprovider.RespondToAuthChallengeInput) cognitoidentityprovider.RespondToAuthChallengeRequest

	SetRiskConfigurationRequest(*cognitoidentityprovider.SetRiskConfigurationInput) cognitoidentityprovider.SetRiskConfigurationRequest

	SetUICustomizationRequest(*cognitoidentityprovider.SetUICustomizationInput) cognitoidentityprovider.SetUICustomizationRequest

	SetUserMFAPreferenceRequest(*cognitoidentityprovider.SetUserMFAPreferenceInput) cognitoidentityprovider.SetUserMFAPreferenceRequest

	SetUserPoolMfaConfigRequest(*cognitoidentityprovider.SetUserPoolMfaConfigInput) cognitoidentityprovider.SetUserPoolMfaConfigRequest

	SetUserSettingsRequest(*cognitoidentityprovider.SetUserSettingsInput) cognitoidentityprovider.SetUserSettingsRequest

	SignUpRequest(*cognitoidentityprovider.SignUpInput) cognitoidentityprovider.SignUpRequest

	StartUserImportJobRequest(*cognitoidentityprovider.StartUserImportJobInput) cognitoidentityprovider.StartUserImportJobRequest

	StopUserImportJobRequest(*cognitoidentityprovider.StopUserImportJobInput) cognitoidentityprovider.StopUserImportJobRequest

	UpdateAuthEventFeedbackRequest(*cognitoidentityprovider.UpdateAuthEventFeedbackInput) cognitoidentityprovider.UpdateAuthEventFeedbackRequest

	UpdateDeviceStatusRequest(*cognitoidentityprovider.UpdateDeviceStatusInput) cognitoidentityprovider.UpdateDeviceStatusRequest

	UpdateGroupRequest(*cognitoidentityprovider.UpdateGroupInput) cognitoidentityprovider.UpdateGroupRequest

	UpdateIdentityProviderRequest(*cognitoidentityprovider.UpdateIdentityProviderInput) cognitoidentityprovider.UpdateIdentityProviderRequest

	UpdateResourceServerRequest(*cognitoidentityprovider.UpdateResourceServerInput) cognitoidentityprovider.UpdateResourceServerRequest

	UpdateUserAttributesRequest(*cognitoidentityprovider.UpdateUserAttributesInput) cognitoidentityprovider.UpdateUserAttributesRequest

	UpdateUserPoolRequest(*cognitoidentityprovider.UpdateUserPoolInput) cognitoidentityprovider.UpdateUserPoolRequest

	UpdateUserPoolClientRequest(*cognitoidentityprovider.UpdateUserPoolClientInput) cognitoidentityprovider.UpdateUserPoolClientRequest

	VerifySoftwareTokenRequest(*cognitoidentityprovider.VerifySoftwareTokenInput) cognitoidentityprovider.VerifySoftwareTokenRequest

	VerifyUserAttributeRequest(*cognitoidentityprovider.VerifyUserAttributeInput) cognitoidentityprovider.VerifyUserAttributeRequest
}

var _ CognitoIdentityProviderAPI = (*cognitoidentityprovider.CognitoIdentityProvider)(nil)
