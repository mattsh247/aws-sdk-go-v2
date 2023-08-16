// Code generated by smithy-go-codegen DO NOT EDIT.


package rds

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"context"
	"net/http"
	"github.com/aws/smithy-go/middleware"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"testing"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/unit"
)

func TestClientCreateDBInstanceReadReplica_presignURLCustomization(t *testing.T) {
	cases := map[string]struct {
		Input              *CreateDBInstanceReadReplicaInput
		ClientRegion       string
		ExpectPresignedURL string
		ExpectPresignedURLDestinationRegion string
		ExpectRequestURL   string
		ExpectErr          string
	}{
		"have presigned URL no auto fill": {
			Input: &CreateDBInstanceReadReplicaInput{
				PreSignedUrl: aws.String("https://example.aws/signed-url"),
			},
			ClientRegion:       "mock-region",
			ExpectPresignedURL: "https://example.aws/signed-url",
			ExpectRequestURL:   "https://service.mock-region.amazonaws.com/",
		},

		"no source region no auto fill": {
			Input:            &CreateDBInstanceReadReplicaInput{},
			ClientRegion:     "mock-region",
			ExpectRequestURL: "https://service.mock-region.amazonaws.com/",
		},

		"auto fill presign URL matching region": {
			Input: &CreateDBInstanceReadReplicaInput{
				SourceRegion: aws.String("mock-region"),
			},
			ClientRegion:       "mock-region",
			ExpectPresignedURL: "https://service.mock-region.amazonaws.com/",
			ExpectPresignedURLDestinationRegion: "DestinationRegion=mock-region",
			ExpectRequestURL:   "https://service.mock-region.amazonaws.com/",
		},

		"auto fill presign URL different region": {
			Input: &CreateDBInstanceReadReplicaInput{
				SourceRegion: aws.String("mock-other-region"),
			},
			ClientRegion:       "mock-region",
			ExpectPresignedURL: "https://service.mock-other-region.amazonaws.com/",
			ExpectPresignedURLDestinationRegion: "DestinationRegion=mock-region",
			ExpectRequestURL:   "https://service.mock-region.amazonaws.com/",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := New(Options{
				Region:      c.ClientRegion,
				Credentials: unit.StubCredentialsProvider{},
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					if e, a := c.ExpectRequestURL, r.URL.String(); !strings.HasPrefix(a, e) {
						t.Errorf("expect presigned URL to contain %v, got %v", e, a)
					}
					return smithyhttp.NopClient{}.Do(r)
				}),
				EndpointResolver: EndpointResolverFunc(
					func(region string, options EndpointResolverOptions) (aws.Endpoint, error) {
						return aws.Endpoint{
							URL:           "https://service." + region + ".amazonaws.com",
							SigningRegion: c.ClientRegion,
						}, nil
					}),
			})

			_, err := client.CreateDBInstanceReadReplica(context.Background(), c.Input,
				func(o *Options) {
					o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) (err error) {
						_, err = stack.Initialize.Remove("OperationInputValidation")
						if err != nil {
							return err
						}

						return stack.Serialize.Add(middleware.SerializeMiddlewareFunc(t.Name(),
							func(
								ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
							) (
								out middleware.SerializeOutput, metadata middleware.Metadata, err error,
							) {
								input, ok := in.Parameters.(*CreateDBInstanceReadReplicaInput)
								if !ok {
									t.Fatalf("expect CreateDBInstanceReadReplicaInput, got %T", in.Parameters)
								}

								// Switch based on if presign flow or not
								if presignedurlcust.GetIsPresigning(ctx) {
									// Presign Flow
									if v := input.PreSignedUrl; v != nil {
										t.Errorf("expect no presigned URL, got %v", *v)
									}
									if input.destinationRegion == nil {
										t.Fatalf("expect destination region to be set")
									}
									if e, a := c.ClientRegion, *input.destinationRegion; e != a {
										t.Errorf("expect %v destination region, got %v", e, a)
									}

								} else {
									// Operation flow
									if v := input.destinationRegion; v != nil {
										t.Errorf("expect no destination region, got %v", *v)
									}

									if len(c.ExpectPresignedURL) != 0 {
										if input.PreSignedUrl == nil {
											t.Fatalf("expect presigned URL, got none")
										}

										if e, a := c.ExpectPresignedURL, *input.PreSignedUrl; !strings.HasPrefix(a, e) {
											t.Errorf("expect presigned URL to contain %v, got %v", e, a)
										}
										if e, a := c.ExpectPresignedURLDestinationRegion, *input.PreSignedUrl; !strings.Contains(a, e) {
											t.Errorf("expect presigned URL destination region to contain %v, got %v", e, a)
										}
										return next.HandleSerialize(ctx, in)
									}
									if v := input.PreSignedUrl; v != nil {
										t.Errorf("expect no presigned url, got %v", *v)
									}
								}

								return next.HandleSerialize(ctx, in)
							},
						), middleware.After)
					})
				},
			)

			if len(c.ExpectErr) != 0 {
				if err == nil {
					t.Fatalf("expect error, got none")
				}
				if e, a := c.ExpectErr, err.Error(); !strings.Contains(a, e) {
					t.Fatalf("expect error to contain %v, got %v", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
		})
	}
}
