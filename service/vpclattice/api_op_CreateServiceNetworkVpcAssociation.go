// Code generated by smithy-go-codegen DO NOT EDIT.


package vpclattice

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
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Associates a VPC with a service network. When you associate a VPC with the
// service network, it enables all the resources within that VPC to be clients and
// communicate with other services in the service network. For more information,
// see Manage VPC associations (https://docs.aws.amazon.com/vpc-lattice/latest/ug/service-network-associations.html#service-network-vpc-associations)
// in the Amazon VPC Lattice User Guide. You can't use this operation if there is a
// disassociation in progress. If the association fails, retry by deleting the
// association and recreating it. As a result of this operation, the association
// gets created in the service network account and the VPC owner account. If you
// add a security group to the service network and VPC association, the association
// must continue to always have at least one security group. You can add or edit
// security groups at any time. However, to remove all security groups, you must
// first delete the association and recreate it without security groups.
func (c *Client) CreateServiceNetworkVpcAssociation(ctx context.Context, params *CreateServiceNetworkVpcAssociationInput, optFns ...func(*Options)) (*CreateServiceNetworkVpcAssociationOutput, error) {
	if params == nil { params = &CreateServiceNetworkVpcAssociationInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateServiceNetworkVpcAssociation", params, optFns, c.addOperationCreateServiceNetworkVpcAssociationMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateServiceNetworkVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceNetworkVpcAssociationInput struct {
	
	// The ID or Amazon Resource Name (ARN) of the service network. You must use the
	// ARN when the resources specified in the operation are in different accounts.
	//
	// This member is required.
	ServiceNetworkIdentifier *string
	
	// The ID of the VPC.
	//
	// This member is required.
	VpcIdentifier *string
	
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you retry a request that completed successfully using the
	// same client token and parameters, the retry succeeds without performing any
	// actions. If the parameters aren't identical, the retry fails.
	ClientToken *string
	
	// The IDs of the security groups. Security groups aren't added by default. You
	// can add a security group to apply network level controls to control which
	// resources in a VPC are allowed to access the service network and its services.
	// For more information, see Control traffic to resources using security groups (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html)
	// in the Amazon VPC User Guide.
	SecurityGroupIds []string
	
	// The tags for the association.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

type CreateServiceNetworkVpcAssociationOutput struct {
	
	// The Amazon Resource Name (ARN) of the association.
	Arn *string
	
	// The account that created the association.
	CreatedBy *string
	
	// The ID of the association.
	Id *string
	
	// The IDs of the security groups.
	SecurityGroupIds []string
	
	// The operation's status.
	Status types.ServiceNetworkVpcAssociationStatus
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceNetworkVpcAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateServiceNetworkVpcAssociation{}, middleware.After)
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
	if err = addCreateServiceNetworkVpcAssociationResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opCreateServiceNetworkVpcAssociationMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateServiceNetworkVpcAssociationValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceNetworkVpcAssociation(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*CreateServiceNetworkVpcAssociationInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceNetworkVpcAssociationInput ")}
	
	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceNetworkVpcAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceNetworkVpcAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceNetworkVpcAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "vpc-lattice",
	OperationName: "CreateServiceNetworkVpcAssociation",
	}
}

type opCreateServiceNetworkVpcAssociationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateServiceNetworkVpcAssociationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateServiceNetworkVpcAssociationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "vpc-lattice"
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
	        signingName = "vpc-lattice"
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
	        v4aScheme.SigningName = aws.String("vpc-lattice")
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

func addCreateServiceNetworkVpcAssociationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateServiceNetworkVpcAssociationResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
