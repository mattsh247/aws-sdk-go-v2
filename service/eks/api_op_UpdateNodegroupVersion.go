// Code generated by smithy-go-codegen DO NOT EDIT.


package eks

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
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Updates the Kubernetes version or AMI version of an Amazon EKS managed node
// group. You can update a node group using a launch template only if the node
// group was originally deployed with a launch template. If you need to update a
// custom AMI in a node group that was deployed with a launch template, then update
// your custom AMI, specify the new ID in a new version of the launch template, and
// then update the node group to the new version of the launch template. If you
// update without a launch template, then you can update to the latest available
// AMI version of a node group's current Kubernetes version by not specifying a
// Kubernetes version in the request. You can update to the latest AMI version of
// your cluster's current Kubernetes version by specifying your cluster's
// Kubernetes version in the request. For information about Linux versions, see
// Amazon EKS optimized Amazon Linux AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
// in the Amazon EKS User Guide. For information about Windows versions, see
// Amazon EKS optimized Windows AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html)
// in the Amazon EKS User Guide. You cannot roll back a node group to an earlier
// Kubernetes version or AMI version. When a node in a managed node group is
// terminated due to a scaling action or update, the pods in that node are drained
// first. Amazon EKS attempts to drain the nodes gracefully and will fail if it is
// unable to do so. You can force the update if Amazon EKS is unable to drain the
// nodes as a result of a pod disruption budget issue.
func (c *Client) UpdateNodegroupVersion(ctx context.Context, params *UpdateNodegroupVersionInput, optFns ...func(*Options)) (*UpdateNodegroupVersionOutput, error) {
	if params == nil { params = &UpdateNodegroupVersionInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateNodegroupVersion", params, optFns, c.addOperationUpdateNodegroupVersionMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateNodegroupVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodegroupVersionInput struct {
	
	// The name of the Amazon EKS cluster that is associated with the managed node
	// group to update.
	//
	// This member is required.
	ClusterName *string
	
	// The name of the managed node group to update.
	//
	// This member is required.
	NodegroupName *string
	
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string
	
	// Force the update if the existing node group's pods are unable to be drained due
	// to a pod disruption budget issue. If an update fails because pods could not be
	// drained, you can force the update after it fails to terminate the old node
	// whether or not any pods are running on the node.
	Force bool
	
	// An object representing a node group's launch template specification. You can
	// only update a node group using a launch template if the node group was
	// originally deployed with a launch template.
	LaunchTemplate *types.LaunchTemplateSpecification
	
	// The AMI version of the Amazon EKS optimized AMI to use for the update. By
	// default, the latest available AMI version for the node group's Kubernetes
	// version is used. For information about Linux versions, see Amazon EKS optimized
	// Amazon Linux AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html)
	// in the Amazon EKS User Guide. Amazon EKS managed node groups support the
	// November 2022 and later releases of the Windows AMIs. For information about
	// Windows versions, see Amazon EKS optimized Windows AMI versions (https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html)
	// in the Amazon EKS User Guide. If you specify launchTemplate , and your launch
	// template uses a custom AMI, then don't specify releaseVersion , or the node
	// group update will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	ReleaseVersion *string
	
	// The Kubernetes version to update to. If no version is specified, then the
	// Kubernetes version of the node group does not change. You can specify the
	// Kubernetes version of the cluster to update the node group to the latest AMI
	// version of the cluster's Kubernetes version. If you specify launchTemplate , and
	// your launch template uses a custom AMI, then don't specify version , or the node
	// group update will fail. For more information about using launch templates with
	// Amazon EKS, see Launch template support (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html)
	// in the Amazon EKS User Guide.
	Version *string
	
	noSmithyDocumentSerde
}

type UpdateNodegroupVersionOutput struct {
	
	// An object representing an asynchronous update.
	Update *types.Update
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNodegroupVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNodegroupVersion{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNodegroupVersion{}, middleware.After)
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
	if err = addUpdateNodegroupVersionResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateNodegroupVersionValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNodegroupVersion(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateNodegroupVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateNodegroupVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateNodegroupVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*UpdateNodegroupVersionInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateNodegroupVersionInput ")}
	
	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateNodegroupVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateNodegroupVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateNodegroupVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "eks",
	OperationName: "UpdateNodegroupVersion",
	}
}

type opUpdateNodegroupVersionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateNodegroupVersionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateNodegroupVersionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "eks"
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
	        signingName = "eks"
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
	        v4aScheme.SigningName = aws.String("eks")
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

func addUpdateNodegroupVersionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateNodegroupVersionResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
