// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type FooEnum string

// Enum values for FooEnum
const (
	FooEnumFoo FooEnum = "Foo"
	FooEnumBaz FooEnum = "Baz"
	FooEnumBar FooEnum = "Bar"
	FooEnumOne FooEnum = "1"
	FooEnumZero FooEnum = "0"
)

// Values returns all known values for FooEnum. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (FooEnum) Values() []FooEnum {
	return []FooEnum{
		"Foo",
		"Baz",
		"Bar",
		"1",
		"0",
	}
}

type IntegerEnum = int32

// Enum values for IntegerEnum
const (
	IntegerEnumA IntegerEnum = 1
	IntegerEnumB IntegerEnum = 2
	IntegerEnumC IntegerEnum = 3
)
