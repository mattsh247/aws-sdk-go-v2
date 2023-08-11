// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_QueryMaps_awsAwsquerySerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *QueryMapsInput
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
		// Serializes query maps
		"QuerySimpleQueryMaps": {
			Params: &QueryMapsInput{
				MapArg: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&MapArg.entry.1.key=bar&MapArg.entry.1.value=Bar&MapArg.entry.2.key=foo&MapArg.entry.2.value=Foo`))
			},
		},
		// Serializes query maps and uses xmlName
		"QuerySimpleQueryMapsWithXmlName": {
			Params: &QueryMapsInput{
				RenamedMapArg: map[string]string{
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&Foo.entry.1.key=foo&Foo.entry.1.value=Foo`))
			},
		},
		// Serializes complex query maps
		"QueryComplexQueryMaps": {
			Params: &QueryMapsInput{
				ComplexMapArg: map[string]types.GreetingStruct{
					"bar": {
						Hi: ptr.String("Bar"),
					},
					"foo": {
						Hi: ptr.String("Foo"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&ComplexMapArg.entry.1.key=bar&ComplexMapArg.entry.1.value.hi=Bar&ComplexMapArg.entry.2.key=foo&ComplexMapArg.entry.2.value.hi=Foo`))
			},
		},
		// Does not serialize empty query maps
		"QueryEmptyQueryMaps": {
			Params: &QueryMapsInput{
				MapArg: map[string]string{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08`))
			},
		},
		// Serializes query maps where the member has an xmlName trait
		"QueryQueryMapWithMemberXmlName": {
			Params: &QueryMapsInput{
				MapWithXmlMemberName: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&MapWithXmlMemberName.entry.1.K=bar&MapWithXmlMemberName.entry.1.V=Bar&MapWithXmlMemberName.entry.2.K=foo&MapWithXmlMemberName.entry.2.V=Foo`))
			},
		},
		// Serializes flattened query maps
		"QueryFlattenedQueryMaps": {
			Params: &QueryMapsInput{
				FlattenedMap: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&FlattenedMap.1.key=bar&FlattenedMap.1.value=Bar&FlattenedMap.2.key=foo&FlattenedMap.2.value=Foo`))
			},
		},
		// Serializes flattened query maps that use an xmlName
		"QueryFlattenedQueryMapsWithXmlName": {
			Params: &QueryMapsInput{
				FlattenedMapWithXmlName: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&Hi.1.K=bar&Hi.1.V=Bar&Hi.2.K=foo&Hi.2.V=Foo`))
			},
		},
		// Serializes query map of lists
		"QueryQueryMapOfLists": {
			Params: &QueryMapsInput{
				MapOfLists: map[string][]string{
					"bar": {
						"C",
						"D",
					},
					"foo": {
						"A",
						"B",
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&MapOfLists.entry.1.key=bar&MapOfLists.entry.1.value.member.1=C&MapOfLists.entry.1.value.member.2=D&MapOfLists.entry.2.key=foo&MapOfLists.entry.2.value.member.1=A&MapOfLists.entry.2.value.member.2=B`))
			},
		},
		// Serializes nested struct with map member
		"QueryNestedStructWithMap": {
			Params: &QueryMapsInput{
				NestedStructWithMap: &types.NestedStructWithMap{
					MapArg: map[string]string{
						"bar": "Bar",
						"foo": "Foo",
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps&Version=2020-01-08&NestedStructWithMap.MapArg.entry.1.key=bar&NestedStructWithMap.MapArg.entry.1.value=Bar&NestedStructWithMap.MapArg.entry.2.key=foo&NestedStructWithMap.MapArg.entry.2.value=Foo`))
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
			result, err := client.QueryMaps(context.Background(), c.Params, func(options *Options) {
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
