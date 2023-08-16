// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptographydata

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates pin-related data such as PIN, PIN Verification Value (PVV), PIN
// Block, and PIN Offset during new card issuance or reissuance. For more
// information, see Generate PIN data (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/generate-pin-data.html)
// in the Amazon Web Services Payment Cryptography User Guide. PIN data is never
// transmitted in clear to or from Amazon Web Services Payment Cryptography. This
// operation generates PIN, PVV, or PIN Offset and then encrypts it using Pin
// Encryption Key (PEK) to create an EncryptedPinBlock for transmission from
// Amazon Web Services Payment Cryptography. This operation uses a separate Pin
// Verification Key (PVK) for VISA PVV generation. For information about valid keys
// for this operation, see Understanding key attributes (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-validattributes.html)
// and Key types for specific data operations (https://docs.aws.amazon.com/payment-cryptography/latest/userguide/crypto-ops-validkeys-ops.html)
// in the Amazon Web Services Payment Cryptography User Guide. Cross-account use:
// This operation can't be used across different Amazon Web Services accounts.
// Related operations:
//   - GenerateCardValidationData
//   - TranslatePinData
//   - VerifyPinData
func (c *Client) GeneratePinData(ctx context.Context, params *GeneratePinDataInput, optFns ...func(*Options)) (*GeneratePinDataOutput, error) {
	if params == nil {
		params = &GeneratePinDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GeneratePinData", params, optFns, c.addOperationGeneratePinDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GeneratePinDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GeneratePinDataInput struct {

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses to
	// encrypt the PIN Block.
	//
	// This member is required.
	EncryptionKeyIdentifier *string

	// The attributes and values to use for PIN, PVV, or PIN Offset generation.
	//
	// This member is required.
	GenerationAttributes types.PinGenerationAttributes

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses for
	// pin data generation.
	//
	// This member is required.
	GenerationKeyIdentifier *string

	// The PIN encoding format for pin data generation as specified in ISO 9564.
	// Amazon Web Services Payment Cryptography supports ISO_Format_0 and ISO_Format_3
	// . The ISO_Format_0 PIN block format is equivalent to the ANSI X9.8, VISA-1, and
	// ECI-1 PIN block formats. It is similar to a VISA-4 PIN block format. It supports
	// a PIN from 4 to 12 digits in length. The ISO_Format_3 PIN block format is the
	// same as ISO_Format_0 except that the fill digits are random values from 10 to
	// 15.
	//
	// This member is required.
	PinBlockFormat types.PinBlockFormatForPinData

	// The Primary Account Number (PAN), a unique identifier for a payment credit or
	// debit card that associates the card with a specific account holder.
	//
	// This member is required.
	PrimaryAccountNumber *string

	// The length of PIN under generation.
	PinDataLength *int32

	noSmithyDocumentSerde
}

type GeneratePinDataOutput struct {

	// The PIN block encrypted under PEK from Amazon Web Services Payment
	// Cryptography. The encrypted PIN block is a composite of PAN (Primary Account
	// Number) and PIN (Personal Identification Number), generated in accordance with
	// ISO 9564 standard.
	//
	// This member is required.
	EncryptedPinBlock *string

	// The keyARN of the PEK that Amazon Web Services Payment Cryptography uses for
	// encrypted pin block generation.
	//
	// This member is required.
	EncryptionKeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed. Amazon Web Services Payment Cryptography calculates the KCV by using
	// standard algorithms, typically by encrypting 8 or 16 bytes or "00" or "01" and
	// then truncating the result to the first 3 bytes, or 6 hex digits, of the
	// resulting cryptogram.
	//
	// This member is required.
	EncryptionKeyCheckValue *string

	// The keyARN of the pin data generation key that Amazon Web Services Payment
	// Cryptography uses for PIN, PVV or PIN Offset generation.
	//
	// This member is required.
	GenerationKeyArn *string

	// The key check value (KCV) of the encryption key. The KCV is used to check if
	// all parties holding a given key have the same key or to detect that a key has
	// changed. Amazon Web Services Payment Cryptography calculates the KCV by using
	// standard algorithms, typically by encrypting 8 or 16 bytes or "00" or "01" and
	// then truncating the result to the first 3 bytes, or 6 hex digits, of the
	// resulting cryptogram.
	//
	// This member is required.
	GenerationKeyCheckValue *string

	// The attributes and values Amazon Web Services Payment Cryptography uses for pin
	// data generation.
	//
	// This member is required.
	PinData types.PinData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGeneratePinDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGeneratePinData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGeneratePinData{}, middleware.After)
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
	if err = addGeneratePinDataResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGeneratePinDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGeneratePinData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGeneratePinData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "payment-cryptography",
		OperationName: "GeneratePinData",
	}
}

type opGeneratePinDataResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGeneratePinDataResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGeneratePinDataResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "payment-cryptography"
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
				signingName = "payment-cryptography"
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
				v4aScheme.SigningName = aws.String("payment-cryptography")
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

func addGeneratePinDataResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGeneratePinDataResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
