// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"math"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestClient_AllQueryStringTypes_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *AllQueryStringTypesInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes query string parameters with all supported types
		"RestJsonAllQueryStringTypes": {
			Params: &AllQueryStringTypesInput{
				QueryString: ptr.String("Hello there"),
				QueryStringList: []string{
					"a",
					"b",
					"c",
				},
				QueryStringSet: []string{
					"a",
					"b",
					"c",
				},
				QueryByte:    ptr.Int8(1),
				QueryShort:   ptr.Int16(2),
				QueryInteger: ptr.Int32(3),
				QueryIntegerList: []int32{
					1,
					2,
					3,
				},
				QueryIntegerSet: []int32{
					1,
					2,
					3,
				},
				QueryLong:   ptr.Int64(4),
				QueryFloat:  ptr.Float32(1.1),
				QueryDouble: ptr.Float64(1.1),
				QueryDoubleList: []float64{
					1.1,
					2.1,
					3.1,
				},
				QueryBoolean: ptr.Bool(true),
				QueryBooleanList: []bool{
					true,
					false,
					true,
				},
				QueryTimestamp: ptr.Time(smithytime.ParseEpochSeconds(1)),
				QueryTimestampList: []time.Time{
					smithytime.ParseEpochSeconds(1),
					smithytime.ParseEpochSeconds(2),
					smithytime.ParseEpochSeconds(3),
				},
				QueryEnum: types.FooEnum("Foo"),
				QueryEnumList: []types.FooEnum{
					types.FooEnum("Foo"),
					types.FooEnum("Baz"),
					types.FooEnum("Bar"),
				},
				QueryIntegerEnum: 1,
				QueryIntegerEnumList: []types.IntegerEnum{
					1,
					2,
					3,
				},
				QueryParamsMapOfStringList: map[string][]string{
					"String": {
						"Hello there",
					},
					"StringList": {
						"a",
						"b",
						"c",
					},
					"StringSet": {
						"a",
						"b",
						"c",
					},
					"Byte": {
						"1",
					},
					"Short": {
						"2",
					},
					"Integer": {
						"3",
					},
					"IntegerList": {
						"1",
						"2",
						"3",
					},
					"IntegerSet": {
						"1",
						"2",
						"3",
					},
					"Long": {
						"4",
					},
					"Float": {
						"1.1",
					},
					"Double": {
						"1.1",
					},
					"DoubleList": {
						"1.1",
						"2.1",
						"3.1",
					},
					"Boolean": {
						"true",
					},
					"BooleanList": {
						"true",
						"false",
						"true",
					},
					"Timestamp": {
						"1970-01-01T00:00:01Z",
					},
					"TimestampList": {
						"1970-01-01T00:00:01Z",
						"1970-01-01T00:00:02Z",
						"1970-01-01T00:00:03Z",
					},
					"Enum": {
						"Foo",
					},
					"EnumList": {
						"Foo",
						"Baz",
						"Bar",
					},
					"IntegerEnum": {
						"1",
					},
					"IntegerEnumList": {
						"1",
						"2",
						"3",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "String", Value: "Hello%20there"},
				{Key: "StringList", Value: "a"},
				{Key: "StringList", Value: "b"},
				{Key: "StringList", Value: "c"},
				{Key: "StringSet", Value: "a"},
				{Key: "StringSet", Value: "b"},
				{Key: "StringSet", Value: "c"},
				{Key: "Byte", Value: "1"},
				{Key: "Short", Value: "2"},
				{Key: "Integer", Value: "3"},
				{Key: "IntegerList", Value: "1"},
				{Key: "IntegerList", Value: "2"},
				{Key: "IntegerList", Value: "3"},
				{Key: "IntegerSet", Value: "1"},
				{Key: "IntegerSet", Value: "2"},
				{Key: "IntegerSet", Value: "3"},
				{Key: "Long", Value: "4"},
				{Key: "Float", Value: "1.1"},
				{Key: "Double", Value: "1.1"},
				{Key: "DoubleList", Value: "1.1"},
				{Key: "DoubleList", Value: "2.1"},
				{Key: "DoubleList", Value: "3.1"},
				{Key: "Boolean", Value: "true"},
				{Key: "BooleanList", Value: "true"},
				{Key: "BooleanList", Value: "false"},
				{Key: "BooleanList", Value: "true"},
				{Key: "Timestamp", Value: "1970-01-01T00%3A00%3A01Z"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A01Z"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A02Z"},
				{Key: "TimestampList", Value: "1970-01-01T00%3A00%3A03Z"},
				{Key: "Enum", Value: "Foo"},
				{Key: "EnumList", Value: "Foo"},
				{Key: "EnumList", Value: "Baz"},
				{Key: "EnumList", Value: "Bar"},
				{Key: "IntegerEnum", Value: "1"},
				{Key: "IntegerEnumList", Value: "1"},
				{Key: "IntegerEnumList", Value: "2"},
				{Key: "IntegerEnumList", Value: "3"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Handles query string maps
		"RestJsonQueryStringMap": {
			Params: &AllQueryStringTypesInput{
				QueryParamsMapOfStringList: map[string][]string{
					"QueryParamsStringKeyA": {
						"Foo",
					},
					"QueryParamsStringKeyB": {
						"Bar",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "QueryParamsStringKeyA", Value: "Foo"},
				{Key: "QueryParamsStringKeyB", Value: "Bar"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Handles escaping all required characters in the query string.
		"RestJsonQueryStringEscaping": {
			Params: &AllQueryStringTypesInput{
				QueryString: ptr.String("%:/?#[]@!$&'()*+,;=😹"),
				QueryParamsMapOfStringList: map[string][]string{
					"String": {
						"%:/?#[]@!$&'()*+,;=😹",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "String", Value: "%25%3A%2F%3F%23%5B%5D%40%21%24%26%27%28%29%2A%2B%2C%3B%3D%F0%9F%98%B9"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling NaN float query values.
		"RestJsonSupportsNaNFloatQueryValues": {
			Params: &AllQueryStringTypesInput{
				QueryFloat:  ptr.Float32(float32(math.NaN())),
				QueryDouble: ptr.Float64(math.NaN()),
				QueryParamsMapOfStringList: map[string][]string{
					"Float": {
						"NaN",
					},
					"Double": {
						"NaN",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "Float", Value: "NaN"},
				{Key: "Double", Value: "NaN"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling Infinity float query values.
		"RestJsonSupportsInfinityFloatQueryValues": {
			Params: &AllQueryStringTypesInput{
				QueryFloat:  ptr.Float32(float32(math.Inf(1))),
				QueryDouble: ptr.Float64(math.Inf(1)),
				QueryParamsMapOfStringList: map[string][]string{
					"Float": {
						"Infinity",
					},
					"Double": {
						"Infinity",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "Float", Value: "Infinity"},
				{Key: "Double", Value: "Infinity"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Supports handling -Infinity float query values.
		"RestJsonSupportsNegativeInfinityFloatQueryValues": {
			Params: &AllQueryStringTypesInput{
				QueryFloat:  ptr.Float32(float32(math.Inf(-1))),
				QueryDouble: ptr.Float64(math.Inf(-1)),
				QueryParamsMapOfStringList: map[string][]string{
					"Float": {
						"-Infinity",
					},
					"Double": {
						"-Infinity",
					},
				},
			},
			ExpectMethod:  "GET",
			ExpectURIPath: "/AllQueryStringTypesInput",
			ExpectQuery: []smithytesting.QueryItem{
				{Key: "Float", Value: "-Infinity"},
				{Key: "Double", Value: "-Infinity"},
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderEmpty(actual)
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               protocoltesthttp.NewClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			capturedReq := &http.Request{}
			result, err := client.AllQueryStringTypes(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyhttp.AddCaptureRequestMiddleware(stack, capturedReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, capturedReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, capturedReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(capturedReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, capturedReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, capturedReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, capturedReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(capturedReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
