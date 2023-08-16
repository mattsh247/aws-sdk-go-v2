// Code generated by smithy-go-codegen DO NOT EDIT.


package transfer

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
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates the connector, which captures the parameters for a connection for the
// AS2 or SFTP protocol. For AS2, the connector is required for sending files to an
// externally hosted AS2 server. For SFTP, the connector is required when sending
// files to an SFTP server or receiving files from an SFTP server. For more details
// about connectors, see Create AS2 connectors (https://docs.aws.amazon.com/transfer/latest/userguide/create-b2b-server.html#configure-as2-connector)
// and Create SFTP connectors (https://docs.aws.amazon.com/transfer/latest/userguide/configure-sftp-connector.html)
// . You must specify exactly one configuration object: either for AS2 ( As2Config
// ) or SFTP ( SftpConfig ).
func (c *Client) CreateConnector(ctx context.Context, params *CreateConnectorInput, optFns ...func(*Options)) (*CreateConnectorOutput, error) {
	if params == nil { params = &CreateConnectorInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateConnector", params, optFns, c.addOperationCreateConnectorMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorInput struct {
	
	// Connectors are used to send files using either the AS2 or SFTP protocol. For
	// the access role, provide the Amazon Resource Name (ARN) of the Identity and
	// Access Management role to use. For AS2 connectors With AS2, you can send files
	// by calling StartFileTransfer and specifying the file paths in the request
	// parameter, SendFilePaths . We use the file’s parent directory (for example, for
	// --send-file-paths /bucket/dir/file.txt , parent directory is /bucket/dir/ ) to
	// temporarily store a processed AS2 message file, store the MDN when we receive
	// them from the partner, and write a final JSON file containing relevant metadata
	// of the transmission. So, the AccessRole needs to provide read and write access
	// to the parent directory of the file location used in the StartFileTransfer
	// request. Additionally, you need to provide read and write access to the parent
	// directory of the files that you intend to send with StartFileTransfer . If you
	// are using Basic authentication for your AS2 connector, the access role requires
	// the secretsmanager:GetSecretValue permission for the secret. If the secret is
	// encrypted using a customer-managed key instead of the Amazon Web Services
	// managed key in Secrets Manager, then the role also needs the kms:Decrypt
	// permission for that key. For SFTP connectors Make sure that the access role
	// provides read and write access to the parent directory of the file location
	// that's used in the StartFileTransfer request. Additionally, make sure that the
	// role provides secretsmanager:GetSecretValue permission to Secrets Manager.
	//
	// This member is required.
	AccessRole *string
	
	// The URL of the partner's AS2 or SFTP endpoint.
	//
	// This member is required.
	Url *string
	
	// A structure that contains the parameters for an AS2 connector object.
	As2Config *types.As2ConnectorConfig
	
	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that allows a connector to turn on CloudWatch logging for Amazon S3 events. When
	// set, you can view connector activity in your CloudWatch logs.
	LoggingRole *string
	
	// A structure that contains the parameters for an SFTP connector object.
	SftpConfig *types.SftpConnectorConfig
	
	// Key-value pairs that can be used to group and search for connectors. Tags are
	// metadata attached to connectors for any purpose.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateConnectorOutput struct {
	
	// The unique identifier for the connector, returned after the API call succeeds.
	//
	// This member is required.
	ConnectorId *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConnector{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConnector{}, middleware.After)
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
	if err = addCreateConnectorResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateConnectorValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnector(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "transfer",
	OperationName: "CreateConnector",
	}
}

type opCreateConnectorResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateConnectorResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateConnectorResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "transfer"
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
	        signingName = "transfer"
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
	        v4aScheme.SigningName = aws.String("transfer")
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

func addCreateConnectorResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateConnectorResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
