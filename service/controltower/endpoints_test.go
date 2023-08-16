// Code generated by smithy-go-codegen DO NOT EDIT.


package controltower

import (
	"github.com/google/go-cmp/cmp"
	"context"
	"net/http"
	"github.com/aws/smithy-go/ptr"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"strings"
	"testing"
	"net/url"
)

// For region ap-northeast-1 with FIPS disabled and DualStack disabled
func TestEndpointCase0(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ap-northeast-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ap-northeast-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ap-northeast-2 with FIPS disabled and DualStack disabled
func TestEndpointCase1(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ap-northeast-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ap-northeast-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ap-south-1 with FIPS disabled and DualStack disabled
func TestEndpointCase2(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ap-south-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ap-south-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ap-southeast-1 with FIPS disabled and DualStack disabled
func TestEndpointCase3(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ap-southeast-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ap-southeast-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ap-southeast-2 with FIPS disabled and DualStack disabled
func TestEndpointCase4(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ap-southeast-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ap-southeast-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ca-central-1 with FIPS disabled and DualStack disabled
func TestEndpointCase5(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ca-central-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.ca-central-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region ca-central-1 with FIPS enabled and DualStack disabled
func TestEndpointCase6(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("ca-central-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.ca-central-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region eu-central-1 with FIPS disabled and DualStack disabled
func TestEndpointCase7(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("eu-central-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.eu-central-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region eu-north-1 with FIPS disabled and DualStack disabled
func TestEndpointCase8(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("eu-north-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.eu-north-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region eu-west-1 with FIPS disabled and DualStack disabled
func TestEndpointCase9(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("eu-west-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.eu-west-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region eu-west-2 with FIPS disabled and DualStack disabled
func TestEndpointCase10(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("eu-west-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.eu-west-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region eu-west-3 with FIPS disabled and DualStack disabled
func TestEndpointCase11(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("eu-west-3"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.eu-west-3.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region sa-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase12(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("sa-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.sa-east-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase13(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-east-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase14(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-east-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-2 with FIPS disabled and DualStack disabled
func TestEndpointCase15(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-east-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-2 with FIPS enabled and DualStack disabled
func TestEndpointCase16(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-east-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-west-2 with FIPS disabled and DualStack disabled
func TestEndpointCase17(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-west-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-west-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-west-2 with FIPS enabled and DualStack disabled
func TestEndpointCase18(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-west-2"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-west-2.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase19(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-east-1.api.aws")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase20(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-east-1.api.aws")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region cn-north-1 with FIPS enabled and DualStack enabled
func TestEndpointCase21(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("cn-north-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.cn-north-1.api.amazonwebservices.com.cn")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region cn-north-1 with FIPS enabled and DualStack disabled
func TestEndpointCase22(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("cn-north-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.cn-north-1.amazonaws.com.cn")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region cn-north-1 with FIPS disabled and DualStack enabled
func TestEndpointCase23(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("cn-north-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.cn-north-1.api.amazonwebservices.com.cn")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region cn-north-1 with FIPS disabled and DualStack disabled
func TestEndpointCase24(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("cn-north-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.cn-north-1.amazonaws.com.cn")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-gov-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase25(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-gov-east-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-gov-east-1.api.aws")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-gov-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase26(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-gov-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-gov-east-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-gov-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase27(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-gov-east-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-gov-east-1.api.aws")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-gov-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase28(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-gov-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-gov-east-1.amazonaws.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-iso-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase29(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-iso-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-iso-east-1.c2s.ic.gov")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-iso-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase30(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-iso-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-iso-east-1.c2s.ic.gov")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-isob-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase31(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-isob-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower-fips.us-isob-east-1.sc2s.sgov.gov")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For region us-isob-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase32(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-isob-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://controltower.us-isob-east-1.sc2s.sgov.gov")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For custom endpoint with region set and fips disabled and dualstack disabled
func TestEndpointCase33(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Endpoint: ptr.String("https://example.com"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://example.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For custom endpoint with region not set and fips disabled and dualstack disabled
func TestEndpointCase34(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Endpoint: ptr.String("https://example.com"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://example.com")

expectEndpoint := smithyendpoints.Endpoint{
    URI: *uri,
    Headers: http.Header{},
    Properties: smithy.Properties{},
}

if e, a := expectEndpoint.URI, result.URI; e != a{
    t.Errorf("expect %v URI, got %v", e, a)
}

if diff := cmp.Diff(expectEndpoint.Headers, result.Headers); diff != "" {
    t.Errorf("expect headers to match\n%s", diff)
}

if diff := cmp.Diff(expectEndpoint.Properties, result.Properties,
    cmp.AllowUnexported(smithy.Properties{}),
); diff != "" {
    t.Errorf("expect properties to match\n%s", diff)
}
}

// For custom endpoint with fips enabled and dualstack disabled
func TestEndpointCase35(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Endpoint: ptr.String("https://example.com"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "Invalid Configuration: FIPS and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}

// For custom endpoint with fips disabled and dualstack enabled
func TestEndpointCase36(t *testing.T) {
    var params = EndpointParameters{
    Region: ptr.String("us-east-1"),
UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Endpoint: ptr.String("https://example.com"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "Invalid Configuration: Dualstack and custom endpoint are not supported", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}
