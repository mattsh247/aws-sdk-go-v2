// Code generated by smithy-go-codegen DO NOT EDIT.


package awsrestjson

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"context"
	"net/http"
	"io"
	"github.com/aws/smithy-go/middleware"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/ptr"
	smithyprivateprotocol "github.com/aws/smithy-go/private/protocol"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	"testing"
	"net/url"
)

func TestClient_TestNoPayload_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
	Params *TestNoPayloadInput
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
		// Serializes a GET request with no modeled body
		"RestJsonHttpWithNoModeledBody": {
			Params: &TestNoPayloadInput{
			},
			ExpectMethod: "GET",
			ExpectURIPath: "/no_payload",
			ExpectQuery: []smithytesting.QueryItem {
			},
			ForbidHeader: []string{
				"Content-Length",
				"Content-Type",
			},
			BodyAssert: func(actual io.Reader) error {
			   return smithytesting.CompareReaderEmpty(actual)
			},
		},
		// Serializes a GET request with header member but no modeled body
		"RestJsonHttpWithHeaderMemberNoModeledBody": {
			Params: &TestNoPayloadInput{
			TestId:
			ptr.String("t-12345"),
			},
			ExpectMethod: "GET",
			ExpectURIPath: "/no_payload",
			ExpectQuery: []smithytesting.QueryItem {
			},
			ExpectHeader: http.Header{
				"X-Amz-Test-Id": []string{"t-12345"},
			},
			ForbidHeader: []string{
				"Content-Length",
				"Content-Type",
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
			result, err := client.TestNoPayload(context.Background(), c.Params, func(options *Options) {
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
