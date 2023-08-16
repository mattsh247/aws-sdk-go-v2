// Code generated by smithy-go-codegen DO NOT EDIT.


package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
)

func ExampleAnnotationValue_outputUsage() {
	var union types.AnnotationValue
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.AnnotationValueMemberBooleanValue:
			_ = v.Value // Value is bool
		
		case *types.AnnotationValueMemberNumberValue:
			_ = v.Value // Value is float64
		
		case *types.AnnotationValueMemberStringValue:
			_ = v.Value // Value is string
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string
var _ *bool
var _ *float64
