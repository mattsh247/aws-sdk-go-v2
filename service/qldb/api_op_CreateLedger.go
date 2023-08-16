// Code generated by smithy-go-codegen DO NOT EDIT.


package qldb

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
	"time"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a new ledger in your Amazon Web Services account in the current Region.
func (c *Client) CreateLedger(ctx context.Context, params *CreateLedgerInput, optFns ...func(*Options)) (*CreateLedgerOutput, error) {
	if params == nil { params = &CreateLedgerInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateLedger", params, optFns, c.addOperationCreateLedgerMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLedgerInput struct {
	
	// The name of the ledger that you want to create. The name must be unique among
	// all of the ledgers in your Amazon Web Services account in the current Region.
	// Naming constraints for ledger names are defined in Quotas in Amazon QLDB (https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming)
	// in the Amazon QLDB Developer Guide.
	//
	// This member is required.
	Name *string
	
	// The permissions mode to assign to the ledger that you want to create. This
	// parameter can have one of the following values:
	//   - ALLOW_ALL : A legacy permissions mode that enables access control with
	//   API-level granularity for ledgers. This mode allows users who have the
	//   SendCommand API permission for this ledger to run all PartiQL commands (hence,
	//   ALLOW_ALL ) on any tables in the specified ledger. This mode disregards any
	//   table-level or command-level IAM permissions policies that you create for the
	//   ledger.
	//   - STANDARD : (Recommended) A permissions mode that enables access control with
	//   finer granularity for ledgers, tables, and PartiQL commands. By default, this
	//   mode denies all user requests to run any PartiQL commands on any tables in this
	//   ledger. To allow PartiQL commands to run, you must create IAM permissions
	//   policies for specific table resources and PartiQL actions, in addition to the
	//   SendCommand API permission for the ledger. For information, see Getting
	//   started with the standard permissions mode (https://docs.aws.amazon.com/qldb/latest/developerguide/getting-started-standard-mode.html)
	//   in the Amazon QLDB Developer Guide.
	// We strongly recommend using the STANDARD permissions mode to maximize the
	// security of your ledger data.
	//
	// This member is required.
	PermissionsMode types.PermissionsMode
	
	// Specifies whether the ledger is protected from being deleted by any user. If
	// not defined during ledger creation, this feature is enabled ( true ) by default.
	// If deletion protection is enabled, you must first disable it before you can
	// delete the ledger. You can disable it by calling the UpdateLedger operation to
	// set this parameter to false .
	DeletionProtection *bool
	
	// The key in Key Management Service (KMS) to use for encryption of data at rest
	// in the ledger. For more information, see Encryption at rest (https://docs.aws.amazon.com/qldb/latest/developerguide/encryption-at-rest.html)
	// in the Amazon QLDB Developer Guide. Use one of the following options to specify
	// this parameter:
	//   - AWS_OWNED_KMS_KEY : Use an KMS key that is owned and managed by Amazon Web
	//   Services on your behalf.
	//   - Undefined: By default, use an Amazon Web Services owned KMS key.
	//   - A valid symmetric customer managed KMS key: Use the specified symmetric
	//   encryption KMS key in your account that you create, own, and manage. Amazon QLDB
	//   does not support asymmetric keys. For more information, see Using symmetric
	//   and asymmetric keys (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
	//   in the Key Management Service Developer Guide.
	// To specify a customer managed KMS key, you can use its key ID, Amazon Resource
	// Name (ARN), alias name, or alias ARN. When using an alias name, prefix it with
	// "alias/" . To specify a key in a different Amazon Web Services account, you must
	// use the key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// For more information, see Key identifiers (KeyId) (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id)
	// in the Key Management Service Developer Guide.
	KmsKey *string
	
	// The key-value pairs to add as tags to the ledger that you want to create. Tag
	// keys are case sensitive. Tag values are case sensitive and can be null.
	Tags map[string]*string
	
	noSmithyDocumentSerde
}

type CreateLedgerOutput struct {
	
	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string
	
	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time
	
	// Specifies whether the ledger is protected from being deleted by any user. If
	// not defined during ledger creation, this feature is enabled ( true ) by default.
	// If deletion protection is enabled, you must first disable it before you can
	// delete the ledger. You can disable it by calling the UpdateLedger operation to
	// set this parameter to false .
	DeletionProtection *bool
	
	// The ARN of the customer managed KMS key that the ledger uses for encryption at
	// rest. If this parameter is undefined, the ledger uses an Amazon Web Services
	// owned KMS key for encryption.
	KmsKeyArn *string
	
	// The name of the ledger.
	Name *string
	
	// The permissions mode of the ledger that you created.
	PermissionsMode types.PermissionsMode
	
	// The current status of the ledger.
	State types.LedgerState
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLedger{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLedger{}, middleware.After)
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
	if err = addCreateLedgerResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateLedgerValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLedger(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "qldb",
	OperationName: "CreateLedger",
	}
}

type opCreateLedgerResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateLedgerResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateLedgerResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "qldb"
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
	        signingName = "qldb"
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
	        v4aScheme.SigningName = aws.String("qldb")
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

func addCreateLedgerResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateLedgerResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
