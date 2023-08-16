// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	"github.com/aws/smithy-go/ptr"
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

func TestClient_SimpleScalarProperties_awsAwsjson10Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *SimpleScalarPropertiesInput
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
		// Supports handling NaN float values.
		"AwsJson10SupportsNaNFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.NaN())),
				DoubleValue: ptr.Float64(math.NaN()),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.SimpleScalarProperties"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "floatValue": "NaN",
			    "doubleValue": "NaN"
			}`))
			},
		},
		// Supports handling Infinity float values.
		"AwsJson10SupportsInfinityFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.Inf(1))),
				DoubleValue: ptr.Float64(math.Inf(1)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.SimpleScalarProperties"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "floatValue": "Infinity",
			    "doubleValue": "Infinity"
			}`))
			},
		},
		// Supports handling -Infinity float values.
		"AwsJson10SupportsNegativeInfinityFloatInputs": {
			Params: &SimpleScalarPropertiesInput{
				FloatValue:  ptr.Float32(float32(math.Inf(-1))),
				DoubleValue: ptr.Float64(math.Inf(-1)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
				"X-Amz-Target": []string{"JsonRpc10.SimpleScalarProperties"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "floatValue": "-Infinity",
			    "doubleValue": "-Infinity"
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			actualReq := &http.Request{}
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
				HTTPClient: protocoltesthttp.NewClient(),
				Region:     "us-west-2",
			})
			result, err := client.SimpleScalarProperties(context.Background(), c.Params, func(options *Options) {
				options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
					return smithyprivateprotocol.AddCaptureRequestMiddleware(stack, actualReq)
				})
			})
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_SimpleScalarProperties_awsAwsjson10Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarPropertiesOutput
	}{
		// Supports handling NaN float values.
		"AwsJson10SupportsNaNFloatInputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "floatValue": "NaN",
			    "doubleValue": "NaN"
			}`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.NaN())),
				DoubleValue: ptr.Float64(math.NaN()),
			},
		},
		// Supports handling Infinity float values.
		"AwsJson10SupportsInfinityFloatInputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "floatValue": "Infinity",
			    "doubleValue": "Infinity"
			}`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(1))),
				DoubleValue: ptr.Float64(math.Inf(1)),
			},
		},
		// Supports handling -Infinity float values.
		"AwsJson10SupportsNegativeInfinityFloatInputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.0"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "floatValue": "-Infinity",
			    "doubleValue": "-Infinity"
			}`),
			ExpectResult: &SimpleScalarPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(-1))),
				DoubleValue: ptr.Float64(math.Inf(-1)),
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
				Region: "us-west-2",
			})
			var params SimpleScalarPropertiesInput
			result, err := client.SimpleScalarProperties(context.Background(), &params)
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
