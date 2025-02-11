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
