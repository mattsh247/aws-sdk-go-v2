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
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual cluster snapshot. To share a manual cluster snapshot with other Amazon
// Web Services accounts, specify restore as the AttributeName , and use the
// ValuesToAdd parameter to add a list of IDs of the Amazon Web Services accounts
// that are authorized to restore the manual cluster snapshot. Use the value all
// to make the manual cluster snapshot public, which means that it can be copied or
// restored by all Amazon Web Services accounts. Do not add the all value for any
// manual cluster snapshots that contain private information that you don't want
// available to all Amazon Web Services accounts. If a manual cluster snapshot is
// encrypted, it can be shared, but only by specifying a list of authorized Amazon
// Web Services account IDs for the ValuesToAdd parameter. You can't use all as a
// value for that parameter in this case.
func (c *Client) ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) {
	if params == nil { params = &ModifyDBClusterSnapshotAttributeInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterSnapshotAttribute", params, optFns, c.addOperationModifyDBClusterSnapshotAttributeMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ModifyDBClusterSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBClusterSnapshotAttribute .
type ModifyDBClusterSnapshotAttributeInput struct {
	
	// The name of the cluster snapshot attribute to modify. To manage authorization
	// for other Amazon Web Services accounts to copy or restore a manual cluster
	// snapshot, set this value to restore .
	//
	// This member is required.
	AttributeName *string
	
	// The identifier for the cluster snapshot to modify the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string
	
	// A list of cluster snapshot attributes to add to the attribute specified by
	// AttributeName . To authorize other Amazon Web Services accounts to copy or
	// restore a manual cluster snapshot, set this list to include one or more Amazon
	// Web Services account IDs. To make the manual cluster snapshot restorable by any
	// Amazon Web Services account, set it to all . Do not add the all value for any
	// manual cluster snapshots that contain private information that you don't want to
	// be available to all Amazon Web Services accounts.
	ValuesToAdd []string
	
	// A list of cluster snapshot attributes to remove from the attribute specified by
	// AttributeName . To remove authorization for other Amazon Web Services accounts
	// to copy or restore a manual cluster snapshot, set this list to include one or
	// more Amazon Web Services account identifiers. To remove authorization for any
	// Amazon Web Services account to copy or restore the cluster snapshot, set it to
	// all . If you specify all , an Amazon Web Services account whose account ID is
	// explicitly added to the restore attribute can still copy or restore a manual
	// cluster snapshot.
	ValuesToRemove []string
	
	noSmithyDocumentSerde
}

type ModifyDBClusterSnapshotAttributeOutput struct {
	
	// Detailed information about the attributes that are associated with a cluster
	// snapshot.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
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
	if err = addModifyDBClusterSnapshotAttributeResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpModifyDBClusterSnapshotAttributeValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rds",
	OperationName: "ModifyDBClusterSnapshotAttribute",
	}
}

type opModifyDBClusterSnapshotAttributeResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opModifyDBClusterSnapshotAttributeResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opModifyDBClusterSnapshotAttributeResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addModifyDBClusterSnapshotAttributeResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opModifyDBClusterSnapshotAttributeResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
