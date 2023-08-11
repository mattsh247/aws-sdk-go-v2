// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_PostUnionWithJsonName_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *PostUnionWithJsonNameInput
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
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameRequest1": {
			Params: &PostUnionWithJsonNameInput{
				Value: &types.UnionWithJsonNameMemberFoo{Value: "hi"},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/PostUnionWithJsonName",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "value": {
			        "FOO": "hi"
			    }
			}`))
			},
		},
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameRequest2": {
			Params: &PostUnionWithJsonNameInput{
				Value: &types.UnionWithJsonNameMemberBaz{Value: "hi"},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/PostUnionWithJsonName",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "value": {
			        "_baz": "hi"
			    }
			}`))
			},
		},
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameRequest3": {
			Params: &PostUnionWithJsonNameInput{
				Value: &types.UnionWithJsonNameMemberBar{Value: "hi"},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/PostUnionWithJsonName",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "value": {
			        "bar": "hi"
			    }
			}`))
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
			result, err := client.PostUnionWithJsonName(context.Background(), c.Params, func(options *Options) {
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

func TestClient_PostUnionWithJsonName_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *PostUnionWithJsonNameOutput
	}{
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameResponse1": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "value": {
			        "FOO": "hi"
			    }
			}`),
			ExpectResult: &PostUnionWithJsonNameOutput{
				Value: &types.UnionWithJsonNameMemberFoo{Value: "hi"},
			},
		},
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameResponse2": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "value": {
			        "_baz": "hi"
			    }
			}`),
			ExpectResult: &PostUnionWithJsonNameOutput{
				Value: &types.UnionWithJsonNameMemberBaz{Value: "hi"},
			},
		},
		// Tests that jsonName works with union members.
		"PostUnionWithJsonNameResponse3": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "value": {
			        "bar": "hi"
			    }
			}`),
			ExpectResult: &PostUnionWithJsonNameOutput{
				Value: &types.UnionWithJsonNameMemberBar{Value: "hi"},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
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
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params PostUnionWithJsonNameInput
			result, err := client.PostUnionWithJsonName(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
