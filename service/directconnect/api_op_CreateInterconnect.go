// Code generated by smithy-go-codegen DO NOT EDIT.


package directconnect

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
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates an interconnect between an Direct Connect Partner's network and a
// specific Direct Connect location. An interconnect is a connection that is
// capable of hosting other connections. The Direct Connect Partner can use an
// interconnect to provide Direct Connect hosted connections to customers through
// their own network services. Like a standard connection, an interconnect links
// the partner's network to an Direct Connect location over a standard Ethernet
// fiber-optic cable. One end is connected to the partner's router, the other to an
// Direct Connect router. You can automatically add the new interconnect to a link
// aggregation group (LAG) by specifying a LAG ID in the request. This ensures that
// the new interconnect is allocated on the same Direct Connect endpoint that hosts
// the specified LAG. If there are no available ports on the endpoint, the request
// fails and no interconnect is created. For each end customer, the Direct Connect
// Partner provisions a connection on their interconnect by calling
// AllocateHostedConnection . The end customer can then connect to Amazon Web
// Services resources by creating a virtual interface on their connection, using
// the VLAN assigned to them by the Direct Connect Partner. Intended for use by
// Direct Connect Partners only.
func (c *Client) CreateInterconnect(ctx context.Context, params *CreateInterconnectInput, optFns ...func(*Options)) (*CreateInterconnectOutput, error) {
	if params == nil { params = &CreateInterconnectInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateInterconnect", params, optFns, c.addOperationCreateInterconnectMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInterconnectInput struct {
	
	// The port bandwidth, in Gbps. The possible values are 1 and 10.
	//
	// This member is required.
	Bandwidth *string
	
	// The name of the interconnect.
	//
	// This member is required.
	InterconnectName *string
	
	// The location of the interconnect.
	//
	// This member is required.
	Location *string
	
	// The ID of the LAG.
	LagId *string
	
	// The name of the service provider associated with the interconnect.
	ProviderName *string
	
	// The tags to associate with the interconnect.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

// Information about an interconnect.
type CreateInterconnectOutput struct {
	
	// The Direct Connect endpoint on which the physical connection terminates.
	//
	// Deprecated: This member has been deprecated.
	AwsDevice *string
	
	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string
	
	// The Direct Connect endpoint that terminates the logical connection. This device
	// might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string
	
	// The bandwidth of the connection.
	Bandwidth *string
	
	// Indicates whether the interconnect supports a secondary BGP in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy
	
	// The ID of the interconnect.
	InterconnectId *string
	
	// The name of the interconnect.
	InterconnectName *string
	
	// The state of the interconnect. The following are the possible values:
	//   - requested : The initial state of an interconnect. The interconnect stays in
	//   the requested state until the Letter of Authorization (LOA) is sent to the
	//   customer.
	//   - pending : The interconnect is approved, and is being initialized.
	//   - available : The network link is up, and the interconnect is ready for use.
	//   - down : The network link is down.
	//   - deleting : The interconnect is being deleted.
	//   - deleted : The interconnect is deleted.
	//   - unknown : The state of the interconnect is not available.
	InterconnectState types.InterconnectState
	
	// Indicates whether jumbo frames are supported.
	JumboFrameCapable *bool
	
	// The ID of the LAG.
	LagId *string
	
	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time
	
	// The location of the connection.
	Location *string
	
	// The name of the service provider associated with the interconnect.
	ProviderName *string
	
	// The Amazon Web Services Region where the connection is located.
	Region *string
	
	// The tags associated with the interconnect.
	Tags []types.Tag
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInterconnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInterconnect{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInterconnect{}, middleware.After)
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
	if err = addCreateInterconnectResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateInterconnectValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInterconnect(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInterconnect(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "directconnect",
	OperationName: "CreateInterconnect",
	}
}

type opCreateInterconnectResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateInterconnectResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateInterconnectResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "directconnect"
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
	        signingName = "directconnect"
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
	        v4aScheme.SigningName = aws.String("directconnect")
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

func addCreateInterconnectResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateInterconnectResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
