// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	protocoltesthttp "github.com/aws/aws-sdk-go-v2/internal/protocoltest"
	"github.com/aws/smithy-go/middleware"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_HostWithPathOperation_awsAwsjson10Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *HostWithPathOperationInput
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
		// Custom endpoints supplied by users can have paths
		"AwsJson10HostWithPath": {
			Params:        &HostWithPathOperationInput{},
			ExpectMethod:  "POST",
			ExpectURIPath: "/custom/",
			ExpectQuery:   []smithytesting.QueryItem{},
			Host: func() *url.URL {
				host := "https://example.com/custom"
				if len(host) == 0 {
					return nil
				}
				u, err := url.Parse(host)
				if err != nil {
					panic(err)
				}
				return u
			}(),
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderBytes(actual, []byte(`{}`))
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
				HTTPClient: protocoltesthttp.NewClient(),
				Region:     "us-west-2",
			})
			capturedReq := &http.Request{}
			result, err := client.HostWithPathOperation(context.Background(), c.Params, func(options *Options) {
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
