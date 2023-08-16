// Code generated by smithy-go-codegen DO NOT EDIT.


package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
)

func ExampleLogSourceResource_outputUsage() {
	var union types.LogSourceResource
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.LogSourceResourceMemberAwsLogSource:
			_ = v.Value // Value is types.AwsLogSourceResource
		
		case *types.LogSourceResourceMemberCustomLogSource:
			_ = v.Value // Value is types.CustomLogSourceResource
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *types.AwsLogSourceResource
var _ *types.CustomLogSourceResource

func ExampleNotificationConfiguration_outputUsage() {
	var union types.NotificationConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.NotificationConfigurationMemberHttpsNotificationConfiguration:
			_ = v.Value // Value is types.HttpsNotificationConfiguration
		
		case *types.NotificationConfigurationMemberSqsNotificationConfiguration:
			_ = v.Value // Value is types.SqsNotificationConfiguration
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *types.SqsNotificationConfiguration
var _ *types.HttpsNotificationConfiguration
