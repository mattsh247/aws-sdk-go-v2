// Code generated by smithy-go-codegen DO NOT EDIT.


package acm

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Imports a certificate into Certificate Manager (ACM) to use with services that
// are integrated with ACM. Note that integrated services (https://docs.aws.amazon.com/acm/latest/userguide/acm-services.html)
// allow only certificate types and keys they support to be associated with their
// resources. Further, their support differs depending on whether the certificate
// is imported into IAM or into ACM. For more information, see the documentation
// for each service. For more information about importing certificates into ACM,
// see Importing Certificates (https://docs.aws.amazon.com/acm/latest/userguide/import-certificate.html)
// in the Certificate Manager User Guide. ACM does not provide managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html)
// for certificates that you import. Note the following guidelines when importing
// third party certificates:
//   - You must enter the private key that matches the certificate you are
//   importing.
//   - The private key must be unencrypted. You cannot import a private key that
//   is protected by a password or a passphrase.
//   - The private key must be no larger than 5 KB (5,120 bytes).
//   - If the certificate you are importing is not self-signed, you must enter its
//   certificate chain.
//   - If a certificate chain is included, the issuer must be the subject of one
//   of the certificates in the chain.
//   - The certificate, private key, and certificate chain must be PEM-encoded.
//   - The current time must be between the Not Before and Not After certificate
//   fields.
//   - The Issuer field must not be empty.
//   - The OCSP authority URL, if present, must not exceed 1000 characters.
//   - To import a new certificate, omit the CertificateArn argument. Include this
//   argument only when you want to replace a previously imported certificate.
//   - When you import a certificate by using the CLI, you must specify the
//   certificate, the certificate chain, and the private key by their file names
//   preceded by fileb:// . For example, you can specify a certificate saved in the
//   C:\temp folder as fileb://C:\temp\certificate_to_import.pem . If you are
//   making an HTTP or HTTPS Query request, include these arguments as BLOBs.
//   - When you import a certificate by using an SDK, you must specify the
//   certificate, the certificate chain, and the private key files in the manner
//   required by the programming language you're using.
//   - The cryptographic algorithm of an imported certificate must match the
//   algorithm of the signing CA. For example, if the signing CA key type is RSA,
//   then the certificate key type must also be RSA.
// This operation returns the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// of the imported certificate.
func (c *Client) ImportCertificate(ctx context.Context, params *ImportCertificateInput, optFns ...func(*Options)) (*ImportCertificateOutput, error) {
	if params == nil { params = &ImportCertificateInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ImportCertificate", params, optFns, c.addOperationImportCertificateMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ImportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportCertificateInput struct {
	
	// The certificate to import.
	//
	// This member is required.
	Certificate []byte
	
	// The private key that matches the public key in the certificate.
	//
	// This member is required.
	PrivateKey []byte
	
	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of an imported certificate to replace. To import a new certificate, omit this
	// field.
	CertificateArn *string
	
	// The PEM encoded certificate chain.
	CertificateChain []byte
	
	// One or more resource tags to associate with the imported certificate. Note: You
	// cannot apply tags when reimporting a certificate.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type ImportCertificateOutput struct {
	
	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the imported certificate.
	CertificateArn *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationImportCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportCertificate{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportCertificate{}, middleware.After)
	if err != nil { return err }
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
	if err = addImportCertificateResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpImportCertificateValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportCertificate(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "acm",
	OperationName: "ImportCertificate",
	}
}

type opImportCertificateResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opImportCertificateResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opImportCertificateResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "acm"
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
	        signingName = "acm"
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
	        v4aScheme.SigningName = aws.String("acm")
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

func addImportCertificateResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opImportCertificateResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
