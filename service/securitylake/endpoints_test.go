// Code generated by smithy-go-codegen DO NOT EDIT.


package securitylake

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

// For region us-gov-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase0(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-gov-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-gov-east-1.api.aws")

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
func TestEndpointCase1(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-gov-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-gov-east-1.amazonaws.com")

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
func TestEndpointCase2(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-gov-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-gov-east-1.api.aws")

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
func TestEndpointCase3(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-gov-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-gov-east-1.amazonaws.com")

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
func TestEndpointCase4(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
Region: ptr.String("cn-north-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.cn-north-1.api.amazonwebservices.com.cn")

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
func TestEndpointCase5(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("cn-north-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.cn-north-1.amazonaws.com.cn")

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
func TestEndpointCase6(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("cn-north-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.cn-north-1.api.amazonwebservices.com.cn")

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
func TestEndpointCase7(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("cn-north-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.cn-north-1.amazonaws.com.cn")

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

// For region us-iso-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase8(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-iso-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}

// For region us-iso-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase9(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-iso-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-iso-east-1.c2s.ic.gov")

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

// For region us-iso-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase10(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-iso-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}

// For region us-iso-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase11(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-iso-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-iso-east-1.c2s.ic.gov")

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
func TestEndpointCase12(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-east-1.api.aws")

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
func TestEndpointCase13(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-east-1.amazonaws.com")

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
func TestEndpointCase14(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-east-1.api.aws")

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
func TestEndpointCase15(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-east-1.amazonaws.com")

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

// For region us-isob-east-1 with FIPS enabled and DualStack enabled
func TestEndpointCase16(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-isob-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "FIPS and DualStack are enabled, but this partition does not support one or both", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}

// For region us-isob-east-1 with FIPS enabled and DualStack disabled
func TestEndpointCase17(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-isob-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake-fips.us-isob-east-1.sc2s.sgov.gov")

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

// For region us-isob-east-1 with FIPS disabled and DualStack enabled
func TestEndpointCase18(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-isob-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err == nil {
    t.Fatalf("expect error, got none")
}
if e, a := "DualStack is enabled but this partition does not support DualStack", err.Error(); !strings.Contains(a, e) {
    t.Errorf("expect %v error in %v", e, a)
}
}

// For region us-isob-east-1 with FIPS disabled and DualStack disabled
func TestEndpointCase19(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-isob-east-1"),
}

resolver := NewDefaultEndpointResolverV2()
result, err := resolver.ResolveEndpoint(context.Background(), params)
_, _ = result, err

if err != nil {
    t.Fatalf("expect no error, got %v", err)
}

uri, _ := url.Parse("https://securitylake.us-isob-east-1.sc2s.sgov.gov")

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

// For custom endpoint with fips disabled and dualstack disabled
func TestEndpointCase20(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-east-1"),
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
func TestEndpointCase21(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(false),
UseFIPS: ptr.Bool(true),
Region: ptr.String("us-east-1"),
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
func TestEndpointCase22(t *testing.T) {
    var params = EndpointParameters{
    UseDualStack: ptr.Bool(true),
UseFIPS: ptr.Bool(false),
Region: ptr.String("us-east-1"),
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
