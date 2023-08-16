// Code generated by smithy-go-codegen DO NOT EDIT.


package mgn

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
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Updates multiple ReplicationConfigurationTemplates by ID.
func (c *Client) UpdateReplicationConfigurationTemplate(ctx context.Context, params *UpdateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationTemplateOutput, error) {
	if params == nil { params = &UpdateReplicationConfigurationTemplateInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationConfigurationTemplate", params, optFns, c.addOperationUpdateReplicationConfigurationTemplateMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationConfigurationTemplateInput struct {
	
	// Update replication configuration template template ID request.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string
	
	// Update replication configuration template ARN request.
	Arn *string
	
	// Update replication configuration template associate default Application
	// Migration Service Security group request.
	AssociateDefaultSecurityGroup *bool
	
	// Update replication configuration template bandwidth throttling request.
	BandwidthThrottling int64
	
	// Update replication configuration template create Public IP request.
	CreatePublicIP *bool
	
	// Update replication configuration template data plane routing request.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting
	
	// Update replication configuration template use default large Staging Disk type
	// request.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType
	
	// Update replication configuration template EBS encryption request.
	EbsEncryption types.ReplicationConfigurationEbsEncryption
	
	// Update replication configuration template EBS encryption key ARN request.
	EbsEncryptionKeyArn *string
	
	// Update replication configuration template Replication Server instance type
	// request.
	ReplicationServerInstanceType *string
	
	// Update replication configuration template Replication Server Security groups
	// IDs request.
	ReplicationServersSecurityGroupsIDs []string
	
	// Update replication configuration template Staging Area subnet ID request.
	StagingAreaSubnetId *string
	
	// Update replication configuration template Staging Area Tags request.
	StagingAreaTags map[string]string
	
	// Update replication configuration template use dedicated Replication Server
	// request.
	UseDedicatedReplicationServer *bool
	
	// Update replication configuration template use Fips Endpoint request.
	UseFipsEndpoint *bool
	
	noSmithyDocumentSerde
}

type UpdateReplicationConfigurationTemplateOutput struct {
	
	// Replication Configuration template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string
	
	// Replication Configuration template ARN.
	Arn *string
	
	// Replication Configuration template associate default Application Migration
	// Service Security group.
	AssociateDefaultSecurityGroup *bool
	
	// Replication Configuration template bandwidth throttling.
	BandwidthThrottling int64
	
	// Replication Configuration template create Public IP.
	CreatePublicIP *bool
	
	// Replication Configuration template data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting
	
	// Replication Configuration template use default large Staging Disk type.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType
	
	// Replication Configuration template EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption
	
	// Replication Configuration template EBS encryption key ARN.
	EbsEncryptionKeyArn *string
	
	// Replication Configuration template server instance type.
	ReplicationServerInstanceType *string
	
	// Replication Configuration template server Security Groups IDs.
	ReplicationServersSecurityGroupsIDs []string
	
	// Replication Configuration template Staging Area subnet ID.
	StagingAreaSubnetId *string
	
	// Replication Configuration template Staging Area Tags.
	StagingAreaTags map[string]string
	
	// Replication Configuration template Tags.
	Tags map[string]string
	
	// Replication Configuration template use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool
	
	// Replication Configuration template use Fips Endpoint.
	UseFipsEndpoint *bool
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
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
	if err = addUpdateReplicationConfigurationTemplateResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "mgn",
	OperationName: "UpdateReplicationConfigurationTemplate",
	}
}

type opUpdateReplicationConfigurationTemplateResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateReplicationConfigurationTemplateResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateReplicationConfigurationTemplateResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "mgn"
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
	        signingName = "mgn"
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
	        v4aScheme.SigningName = aws.String("mgn")
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

func addUpdateReplicationConfigurationTemplateResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateReplicationConfigurationTemplateResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
