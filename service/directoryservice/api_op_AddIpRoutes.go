// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If the DNS server for your self-managed domain uses a publicly addressable IP
// address, you must add a CIDR address block to correctly route traffic to and
// from your Microsoft AD on Amazon Web Services. AddIpRoutes adds this address
// block. You can also use AddIpRoutes to facilitate routing traffic that uses
// public IP ranges from your Microsoft AD on Amazon Web Services to a peer VPC.
// Before you call AddIpRoutes, ensure that all of the required permissions have
// been explicitly granted through a policy. For details about what permissions are
// required to run the AddIpRoutes operation, see Directory Service API
// Permissions: Actions, Resources, and Conditions Reference (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html)
// .
func (c *Client) AddIpRoutes(ctx context.Context, params *AddIpRoutesInput, optFns ...func(*Options)) (*AddIpRoutesOutput, error) {
	if params == nil {
		params = &AddIpRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddIpRoutes", params, optFns, c.addOperationAddIpRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddIpRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddIpRoutesInput struct {

	// Identifier (ID) of the directory to which to add the address block.
	//
	// This member is required.
	DirectoryId *string

	// IP address blocks, using CIDR format, of the traffic to route. This is often
	// the IP address block of the DNS server used for your self-managed domain.
	//
	// This member is required.
	IpRoutes []types.IpRoute

	// If set to true, updates the inbound and outbound rules of the security group
	// that has the description: "Amazon Web Services created security group for
	// directory ID directory controllers." Following are the new rules: Inbound:
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 88, Source: 0.0.0.0/0
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 123, Source: 0.0.0.0/0
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 138, Source: 0.0.0.0/0
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 389, Source: 0.0.0.0/0
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 464, Source: 0.0.0.0/0
	//   - Type: Custom UDP Rule, Protocol: UDP, Range: 445, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 88, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 135, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 445, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 464, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 636, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 1024-65535, Source: 0.0.0.0/0
	//   - Type: Custom TCP Rule, Protocol: TCP, Range: 3268-33269, Source: 0.0.0.0/0
	//   - Type: DNS (UDP), Protocol: UDP, Range: 53, Source: 0.0.0.0/0
	//   - Type: DNS (TCP), Protocol: TCP, Range: 53, Source: 0.0.0.0/0
	//   - Type: LDAP, Protocol: TCP, Range: 389, Source: 0.0.0.0/0
	//   - Type: All ICMP, Protocol: All, Range: N/A, Source: 0.0.0.0/0
	// Outbound:
	//   - Type: All traffic, Protocol: All, Range: All, Destination: 0.0.0.0/0
	// These security rules impact an internal network interface that is not exposed
	// publicly.
	UpdateSecurityGroupForDirectoryControllers bool

	noSmithyDocumentSerde
}

type AddIpRoutesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddIpRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddIpRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddIpRoutes{}, middleware.After)
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
	if err = addAddIpRoutesResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAddIpRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddIpRoutes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddIpRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "AddIpRoutes",
	}
}

type opAddIpRoutesResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opAddIpRoutesResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opAddIpRoutesResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ds"
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
				signingName = "ds"
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
				v4aScheme.SigningName = aws.String("ds")
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

func addAddIpRoutesResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opAddIpRoutesResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
