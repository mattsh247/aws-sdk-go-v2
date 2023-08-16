// Code generated by smithy-go-codegen DO NOT EDIT.


package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
)

func ExampleManifestOverridesPayload_outputUsage() {
	var union types.ManifestOverridesPayload
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.ManifestOverridesPayloadMemberPayloadData:
			_ = v.Value // Value is string
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string

func ExampleManifestPayload_outputUsage() {
	var union types.ManifestPayload
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.ManifestPayloadMemberPayloadData:
			_ = v.Value // Value is string
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string
