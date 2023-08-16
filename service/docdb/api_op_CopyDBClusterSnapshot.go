// Code generated by smithy-go-codegen DO NOT EDIT.


package docdb

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Copies a snapshot of a cluster. To copy a cluster snapshot from a shared manual
// cluster snapshot, SourceDBClusterSnapshotIdentifier must be the Amazon Resource
// Name (ARN) of the shared cluster snapshot. You can only copy a shared DB cluster
// snapshot, whether encrypted or not, in the same Amazon Web Services Region. To
// cancel the copy operation after it is in progress, delete the target cluster
// snapshot identified by TargetDBClusterSnapshotIdentifier while that cluster
// snapshot is in the copying status.
func (c *Client) CopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*Options)) (*CopyDBClusterSnapshotOutput, error) {
	if params == nil { params = &CopyDBClusterSnapshotInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CopyDBClusterSnapshot", params, optFns, c.addOperationCopyDBClusterSnapshotMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CopyDBClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to CopyDBClusterSnapshot .
type CopyDBClusterSnapshotInput struct {
	
	// The identifier of the cluster snapshot to copy. This parameter is not case
	// sensitive. Constraints:
	//   - Must specify a valid system snapshot in the available state.
	//   - If the source snapshot is in the same Amazon Web Services Region as the
	//   copy, specify a valid snapshot identifier.
	//   - If the source snapshot is in a different Amazon Web Services Region than
	//   the copy, specify a valid cluster snapshot ARN.
	// Example: my-cluster-snapshot1
	//
	// This member is required.
	SourceDBClusterSnapshotIdentifier *string
	
	// The identifier of the new cluster snapshot to create from the source cluster
	// snapshot. This parameter is not case sensitive. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: my-cluster-snapshot2
	//
	// This member is required.
	TargetDBClusterSnapshotIdentifier *string
	
	// Set to true to copy all tags from the source cluster snapshot to the target
	// cluster snapshot, and otherwise false . The default is false .
	CopyTags *bool
	
	// The KMS key ID for an encrypted cluster snapshot. The KMS key ID is the Amazon
	// Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS
	// encryption key. If you copy an encrypted cluster snapshot from your Amazon Web
	// Services account, you can specify a value for KmsKeyId to encrypt the copy with
	// a new KMS encryption key. If you don't specify a value for KmsKeyId , then the
	// copy of the cluster snapshot is encrypted with the same KMS key as the source
	// cluster snapshot. If you copy an encrypted cluster snapshot that is shared from
	// another Amazon Web Services account, then you must specify a value for KmsKeyId
	// . To copy an encrypted cluster snapshot to another Amazon Web Services Region,
	// set KmsKeyId to the KMS key ID that you want to use to encrypt the copy of the
	// cluster snapshot in the destination Region. KMS encryption keys are specific to
	// the Amazon Web Services Region that they are created in, and you can't use
	// encryption keys from one Amazon Web Services Region in another Amazon Web
	// Services Region. If you copy an unencrypted cluster snapshot and specify a value
	// for the KmsKeyId parameter, an error is returned.
	KmsKeyId *string
	
	// The URL that contains a Signature Version 4 signed request for the
	// CopyDBClusterSnapshot API action in the Amazon Web Services Region that contains
	// the source cluster snapshot to copy. You must use the PreSignedUrl parameter
	// when copying a cluster snapshot from another Amazon Web Services Region. If you
	// are using an Amazon Web Services SDK tool or the CLI, you can specify
	// SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl
	// manually. Specifying SourceRegion autogenerates a pre-signed URL that is a
	// valid request for the operation that can be executed in the source Amazon Web
	// Services Region. The presigned URL must be a valid request for the
	// CopyDBClusterSnapshot API action that can be executed in the source Amazon Web
	// Services Region that contains the cluster snapshot to be copied. The presigned
	// URL request must contain the following parameter values:
	//   - SourceRegion - The ID of the region that contains the snapshot to be copied.
	//   - SourceDBClusterSnapshotIdentifier - The identifier for the the encrypted
	//   cluster snapshot to be copied. This identifier must be in the Amazon Resource
	//   Name (ARN) format for the source Amazon Web Services Region. For example, if you
	//   are copying an encrypted cluster snapshot from the us-east-1 Amazon Web Services
	//   Region, then your SourceDBClusterSnapshotIdentifier looks something like the
	//   following:
	//   arn:aws:rds:us-east-1:12345678012:sample-cluster:sample-cluster-snapshot .
	//   - TargetDBClusterSnapshotIdentifier - The identifier for the new cluster
	//   snapshot to be created. This parameter isn't case sensitive.
	PreSignedUrl *string
	
	// The AWS region the resource is in. The presigned URL will be created with this
	// region, if the PresignURL member is empty set.
	SourceRegion *string
	
	// The tags to be assigned to the cluster snapshot.
	Tags []types.Tag
	
	// Used by the SDK's PresignURL autofill customization to specify the region the
	// of the client's request.
	destinationRegion *string
	
	noSmithyDocumentSerde
}

type CopyDBClusterSnapshotOutput struct {
	
	// Detailed information about a cluster snapshot.
	DBClusterSnapshot *types.DBClusterSnapshot
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCopyDBClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopyDBClusterSnapshot{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopyDBClusterSnapshot{}, middleware.After)
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
	if err = addCopyDBClusterSnapshotPresignURLMiddleware(stack, options); err != nil {
	return err
	}
	if err = addCopyDBClusterSnapshotResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCopyDBClusterSnapshotValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopyDBClusterSnapshot(options.Region, ), middleware.Before); err != nil {
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

func copyCopyDBClusterSnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopyDBClusterSnapshotPreSignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.PreSignedUrl == nil || len(*input.PreSignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PreSignedUrl, true, nil
}
func getCopyDBClusterSnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopyDBClusterSnapshotPreSignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.PreSignedUrl = &value
	return nil
}
func setCopyDBClusterSnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}
func addCopyDBClusterSnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL: getCopyDBClusterSnapshotPreSignedUrl,
			
			GetSourceRegion: getCopyDBClusterSnapshotSourceRegion,
			
			CopyInput: copyCopyDBClusterSnapshotInputForPresign,
			
			SetDestinationRegion: setCopyDBClusterSnapshotdestinationRegion,
			
			SetPresignedURL: setCopyDBClusterSnapshotPreSignedUrl,
			
		},
		Presigner: &presignAutoFillCopyDBClusterSnapshotClient{ client: NewPresignClient(New(options))},
		
	})
}
type presignAutoFillCopyDBClusterSnapshotClient struct {
	client *PresignClient
}

// PresignURL is a middleware accessor that satisfies URLPresigner interface.
func (c *presignAutoFillCopyDBClusterSnapshotClient) PresignURL(ctx context.Context, srcRegion string, params interface{}) (*v4.PresignedHTTPRequest, error) {
	input, ok := params.(*CopyDBClusterSnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopyDBClusterSnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = srcRegion
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	presignOptFn := WithPresignClientFromClientOptions(optFn)
	return c.client.PresignCopyDBClusterSnapshot(ctx, input, presignOptFn)
}

func newServiceMetadataMiddleware_opCopyDBClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rds",
	OperationName: "CopyDBClusterSnapshot",
	}
}

// PresignCopyDBClusterSnapshot is used to generate a presigned HTTP Request which
// contains presigned URL, signed headers and HTTP method used.
func (c *PresignClient) PresignCopyDBClusterSnapshot(ctx context.Context, params *CopyDBClusterSnapshotInput, optFns ...func(*PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	if params == nil { params = &CopyDBClusterSnapshotInput{} }
	options := c.options.copy()
	for _, fn := range optFns {
		fn(&options)
	}
	clientOptFns := append(options.ClientOptions, withNopHTTPClientAPIOption)
	
	result, _, err := c.client.invokeOperation(ctx, "CopyDBClusterSnapshot", params, clientOptFns,
		c.client.addOperationCopyDBClusterSnapshotMiddlewares,
		presignConverter(options).convertToPresignMiddleware,
	)
	if err != nil { return nil, err }
	
	out := result.(*v4.PresignedHTTPRequest)
	return out, nil
}

type opCopyDBClusterSnapshotResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCopyDBClusterSnapshotResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCopyDBClusterSnapshotResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "rds"
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
	        signingName = "rds"
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
	        v4aScheme.SigningName = aws.String("rds")
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

func addCopyDBClusterSnapshotResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCopyDBClusterSnapshotResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
