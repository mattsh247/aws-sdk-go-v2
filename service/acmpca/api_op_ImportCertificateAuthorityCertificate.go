// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports a signed private CA certificate into Amazon Web Services Private CA.
// This action is used when you are using a chain of trust whose root is located
// outside Amazon Web Services Private CA. Before you can call this action, the
// following preparations must in place:
//   - In Amazon Web Services Private CA, call the CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
//     action to create the private CA that you plan to back with the imported
//     certificate.
//   - Call the GetCertificateAuthorityCsr (https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCsr.html)
//     action to generate a certificate signing request (CSR).
//   - Sign the CSR using a root or intermediate CA hosted by either an
//     on-premises PKI hierarchy or by a commercial CA.
//   - Create a certificate chain and copy the signed certificate and the
//     certificate chain to your working directory.
//
// Amazon Web Services Private CA supports three scenarios for installing a CA
// certificate:
//   - Installing a certificate for a root CA hosted by Amazon Web Services
//     Private CA.
//   - Installing a subordinate CA certificate whose parent authority is hosted by
//     Amazon Web Services Private CA.
//   - Installing a subordinate CA certificate whose parent authority is
//     externally hosted.
//
// The following additional requirements apply when you import a CA certificate.
//   - Only a self-signed certificate can be imported as a root CA.
//   - A self-signed certificate cannot be imported as a subordinate CA.
//   - Your certificate chain must not include the private CA certificate that you
//     are importing.
//   - Your root CA must be the last certificate in your chain. The subordinate
//     certificate, if any, that your root CA signed must be next to last. The
//     subordinate certificate signed by the preceding subordinate CA must come next,
//     and so on until your chain is built.
//   - The chain must be PEM-encoded.
//   - The maximum allowed size of a certificate is 32 KB.
//   - The maximum allowed size of a certificate chain is 2 MB.
//
// Enforcement of Critical Constraints Amazon Web Services Private CA allows the
// following extensions to be marked critical in the imported CA certificate or
// chain.
//   - Basic constraints (must be marked critical)
//   - Subject alternative names
//   - Key usage
//   - Extended key usage
//   - Authority key identifier
//   - Subject key identifier
//   - Issuer alternative name
//   - Subject directory attributes
//   - Subject information access
//   - Certificate policies
//   - Policy mappings
//   - Inhibit anyPolicy
//
// Amazon Web Services Private CA rejects the following extensions when they are
// marked critical in an imported CA certificate or chain.
//   - Name constraints
//   - Policy constraints
//   - CRL distribution points
//   - Authority information access
//   - Freshest CRL
//   - Any other extension
func (c *Client) ImportCertificateAuthorityCertificate(ctx context.Context, params *ImportCertificateAuthorityCertificateInput, optFns ...func(*Options)) (*ImportCertificateAuthorityCertificateOutput, error) {
	if params == nil {
		params = &ImportCertificateAuthorityCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportCertificateAuthorityCertificate", params, optFns, c.addOperationImportCertificateAuthorityCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportCertificateAuthorityCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateAuthorityCertificateInput struct {

	// The PEM-encoded certificate for a private CA. This may be a self-signed
	// certificate in the case of a root CA, or it may be signed by another CA that you
	// control.
	//
	// This member is required.
	Certificate []byte

	// The Amazon Resource Name (ARN) that was returned when you called
	// CreateCertificateAuthority (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
	// . This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// A PEM-encoded file that contains all of your certificates, other than the
	// certificate you're importing, chaining up to your root CA. Your Amazon Web
	// Services Private CA-hosted or on-premises root certificate is the last in the
	// chain, and each certificate in the chain signs the one preceding. This parameter
	// must be supplied when you import a subordinate CA. When you import a root CA,
	// there is no chain.
	CertificateChain []byte

	noSmithyDocumentSerde
}

type ImportCertificateAuthorityCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateAuthorityCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificateAuthorityCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificateAuthorityCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addImportCertificateAuthorityCertificateResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpImportCertificateAuthorityCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificateAuthorityCertificate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opImportCertificateAuthorityCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "ImportCertificateAuthorityCertificate",
	}
}

type opImportCertificateAuthorityCertificateResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opImportCertificateAuthorityCertificateResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opImportCertificateAuthorityCertificateResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "acm-pca"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "acm-pca"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("acm-pca")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addImportCertificateAuthorityCertificateResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opImportCertificateAuthorityCertificateResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
