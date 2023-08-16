// Code generated by smithy-go-codegen DO NOT EDIT.


package restxml

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"context"
	"net/http"
	"io/ioutil"
	"math"
	"github.com/aws/smithy-go/middleware"
	smithydocument "github.com/aws/smithy-go/document"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	"testing"
)

func TestClient_XmlEmptyBlobs_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
	StatusCode int
	Header http.Header
	BodyMediaType string
	Body []byte
	ExpectResult *XmlEmptyBlobsOutput
	}{
		// Empty blobs are deserialized as empty string
		"XmlEmptyBlobs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlBlobsInputOutput>
			    <data></data>
			</XmlBlobsInputOutput>
			`),
			ExpectResult: &XmlEmptyBlobsOutput{
			Data:
			[]byte(""),
			},
		},
		// Empty self closed blobs are deserialized as empty string
		"XmlEmptySelfClosedBlobs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<XmlBlobsInputOutput>
			    <data/>
			</XmlBlobsInputOutput>
			`),
			ExpectResult: &XmlEmptyBlobsOutput{
			Data:
			[]byte(""),
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
			var params XmlEmptyBlobsInput
			result, err := client.XmlEmptyBlobs(context.Background(), &params)
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
