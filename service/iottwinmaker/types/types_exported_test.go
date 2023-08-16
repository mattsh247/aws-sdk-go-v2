// Code generated by smithy-go-codegen DO NOT EDIT.


package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
)

func ExampleListComponentTypesFilter_outputUsage() {
	var union types.ListComponentTypesFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.ListComponentTypesFilterMemberExtendsFrom:
			_ = v.Value // Value is string
		
		case *types.ListComponentTypesFilterMemberIsAbstract:
			_ = v.Value // Value is bool
		
		case *types.ListComponentTypesFilterMemberNamespace:
			_ = v.Value // Value is string
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string
var _ *string
var _ *bool

func ExampleListEntitiesFilter_outputUsage() {
	var union types.ListEntitiesFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.ListEntitiesFilterMemberComponentTypeId:
			_ = v.Value // Value is string
		
		case *types.ListEntitiesFilterMemberExternalId:
			_ = v.Value // Value is string
		
		case *types.ListEntitiesFilterMemberParentEntityId:
			_ = v.Value // Value is string
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string
var _ *string
var _ *string

func ExampleSyncResourceFilter_outputUsage() {
	var union types.SyncResourceFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
		case *types.SyncResourceFilterMemberExternalId:
			_ = v.Value // Value is string
		
		case *types.SyncResourceFilterMemberResourceId:
			_ = v.Value // Value is string
		
		case *types.SyncResourceFilterMemberResourceType:
			_ = v.Value // Value is types.SyncResourceType
		
		case *types.SyncResourceFilterMemberState:
			_ = v.Value // Value is types.SyncResourceState
		
		case *types.UnknownUnionMember:
			fmt.Println("unknown tag:", v.Tag)
		
		default:
			fmt.Println("union is nil or unknown type")
		
	}
}

var _ *string
var _ types.SyncResourceType
var _ types.SyncResourceState
