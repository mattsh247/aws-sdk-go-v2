// Code generated by smithy-go-codegen DO NOT EDIT.


package restxml

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"context"
	"net/http"
	"io"
	"io/ioutil"
	"math"
	"github.com/aws/smithy-go/middleware"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/ptr"
	smithydocument "github.com/aws/smithy-go/document"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	"testing"
	"net/url"
)

func TestClient_HttpPayloadTraits_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
	Params *HttpPayloadTraitsInput
	ExpectMethod string
	ExpectURIPath string
	ExpectQuery []smithytesting.QueryItem
	RequireQuery []string
	ForbidQuery []string
	ExpectHeader http.Header
	RequireHeader []string
	ForbidHeader []string
	Host *url.URL
	BodyMediaType string
	BodyAssert func(io.Reader) error
	}{
		// Serializes a blob in the HTTP payload
		"HttpPayloadTraitsWithBlob": {
			Params: &HttpPayloadTraitsInput{
			Foo:
			ptr.String("Foo"),
			Blob:
			[]byte("blobby blob blob"),
			},
			ExpectMethod: "POST",
			ExpectURIPath: "/HttpPayloadTraits",
			ExpectQuery: []smithytesting.QueryItem {
			},
			ExpectHeader: http.Header{
				"X-Foo": []string{"Foo"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyAssert: func(actual io.Reader) error {
			 return smithytesting.CompareReaderBytes(actual, []byte(`blobby blob blob`))
			},
		},
		// Serializes an empty blob in the HTTP payload
		"HttpPayloadTraitsWithNoBlobBody": {
			Params: &HttpPayloadTraitsInput{
			Foo:
			ptr.String("Foo"),
			},
			ExpectMethod: "POST",
			ExpectURIPath: "/HttpPayloadTraits",
			ExpectQuery: []smithytesting.QueryItem {
			},
			ExpectHeader: http.Header{
				"X-Foo": []string{"Foo"},
			},
			BodyAssert: func(actual io.Reader) error {
			   return smithytesting.CompareReaderEmpty(actual)
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
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region: "us-west-2",
			})
			result, err := client.HttpPayloadTraits(context.Background(), c.Params, func(options *Options) {
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

func TestClient_HttpPayloadTraits_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
	StatusCode int
	Header http.Header
	BodyMediaType string
	Body []byte
	ExpectResult *HttpPayloadTraitsOutput
	}{
		// Serializes a blob in the HTTP payload
		"HttpPayloadTraitsWithBlob": {
			StatusCode: 200,
			Header: http.Header{
				"X-Foo": []string{"Foo"},
			},
			Body: []byte(`blobby blob blob`),
			ExpectResult: &HttpPayloadTraitsOutput{
			Foo:
			ptr.String("Foo"),
			Blob:
			[]byte("blobby blob blob"),
			},
		},
		// Serializes an empty blob in the HTTP payload
		"HttpPayloadTraitsWithNoBlobBody": {
			StatusCode: 200,
			Header: http.Header{
				"X-Foo": []string{"Foo"},
			},
			Body: []byte(``),
			ExpectResult: &HttpPayloadTraitsOutput{
			Foo:
			ptr.String("Foo"),
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
						Header: headers,
						Request: r,
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
				Region: "us-west-2",
			})
			var params HttpPayloadTraitsInput
			result, err := client.HttpPayloadTraits(context.Background(), &params)
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
				},cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				},cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
