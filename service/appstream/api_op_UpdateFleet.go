// Code generated by smithy-go-codegen DO NOT EDIT.


package appstream

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
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Updates the specified fleet. If the fleet is in the STOPPED state, you can
// update any attribute except the fleet name. If the fleet is in the RUNNING
// state, you can update the following based on the fleet type:
//   - Always-On and On-Demand fleet types You can update the DisplayName ,
//   ComputeCapacity , ImageARN , ImageName , IdleDisconnectTimeoutInSeconds , and
//   DisconnectTimeoutInSeconds attributes.
//   - Elastic fleet type You can update the DisplayName ,
//   IdleDisconnectTimeoutInSeconds , DisconnectTimeoutInSeconds ,
//   MaxConcurrentSessions , SessionScriptS3Location and UsbDeviceFilterStrings
//   attributes.
// If the fleet is in the STARTING or STOPPED state, you can't update it.
func (c *Client) UpdateFleet(ctx context.Context, params *UpdateFleetInput, optFns ...func(*Options)) (*UpdateFleetOutput, error) {
	if params == nil { params = &UpdateFleetInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateFleet", params, optFns, c.addOperationUpdateFleetMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetInput struct {
	
	// The fleet attributes to delete.
	AttributesToDelete []types.FleetAttribute
	
	// The desired capacity for the fleet. This is not allowed for Elastic fleets.
	ComputeCapacity *types.ComputeCapacity
	
	// Deletes the VPC association for the specified fleet.
	//
	// Deprecated: This member has been deprecated.
	DeleteVpcConfig bool
	
	// The description to display.
	Description *string
	
	// The amount of time that a streaming session remains active after users
	// disconnect. If users try to reconnect to the streaming session after a
	// disconnection or network interruption within this time interval, they are
	// connected to their previous session. Otherwise, they are connected to a new
	// session with a new streaming instance. Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *int32
	
	// The fleet name to display.
	DisplayName *string
	
	// The name of the directory and organizational unit (OU) to use to join the fleet
	// to a Microsoft Active Directory domain.
	DomainJoinInfo *types.DomainJoinInfo
	
	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess *bool
	
	// The Amazon Resource Name (ARN) of the IAM role to apply to the fleet. To assume
	// a role, a fleet instance calls the AWS Security Token Service (STS) AssumeRole
	// API operation and passes the ARN of the role to use. The operation creates a new
	// session with temporary credentials. AppStream 2.0 retrieves the temporary
	// credentials and creates the appstream_machine_role credential profile on the
	// instance. For more information, see Using an IAM Role to Grant Permissions to
	// Applications and Scripts Running on AppStream 2.0 Streaming Instances (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string
	
	// The amount of time that users can be idle (inactive) before they are
	// disconnected from their streaming session and the DisconnectTimeoutInSeconds
	// time interval begins. Users are notified before they are disconnected due to
	// inactivity. If users try to reconnect to the streaming session before the time
	// interval specified in DisconnectTimeoutInSeconds elapses, they are connected to
	// their previous session. Users are considered idle when they stop providing
	// keyboard or mouse input during their streaming session. File uploads and
	// downloads, audio in, audio out, and pixels changing do not qualify as user
	// activity. If users continue to be idle after the time interval in
	// IdleDisconnectTimeoutInSeconds elapses, they are disconnected. To prevent users
	// from being disconnected due to inactivity, specify a value of 0. Otherwise,
	// specify a value between 60 and 3600. The default value is 0. If you enable this
	// feature, we recommend that you specify a value that corresponds exactly to a
	// whole number of minutes (for example, 60, 120, and 180). If you don't do this,
	// the value is rounded to the nearest minute. For example, if you specify a value
	// of 70, users are disconnected after 1 minute of inactivity. If you specify a
	// value that is at the midpoint between two different minutes, the value is
	// rounded up. For example, if you specify a value of 90, users are disconnected
	// after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *int32
	
	// The ARN of the public, private, or shared image to use.
	ImageArn *string
	
	// The name of the image used to create the fleet.
	ImageName *string
	
	// The instance type to use when launching fleet instances. The following instance
	// types are available:
	//   - stream.standard.small
	//   - stream.standard.medium
	//   - stream.standard.large
	//   - stream.standard.xlarge
	//   - stream.standard.2xlarge
	//   - stream.compute.large
	//   - stream.compute.xlarge
	//   - stream.compute.2xlarge
	//   - stream.compute.4xlarge
	//   - stream.compute.8xlarge
	//   - stream.memory.large
	//   - stream.memory.xlarge
	//   - stream.memory.2xlarge
	//   - stream.memory.4xlarge
	//   - stream.memory.8xlarge
	//   - stream.memory.z1d.large
	//   - stream.memory.z1d.xlarge
	//   - stream.memory.z1d.2xlarge
	//   - stream.memory.z1d.3xlarge
	//   - stream.memory.z1d.6xlarge
	//   - stream.memory.z1d.12xlarge
	//   - stream.graphics-design.large
	//   - stream.graphics-design.xlarge
	//   - stream.graphics-design.2xlarge
	//   - stream.graphics-design.4xlarge
	//   - stream.graphics-desktop.2xlarge
	//   - stream.graphics.g4dn.xlarge
	//   - stream.graphics.g4dn.2xlarge
	//   - stream.graphics.g4dn.4xlarge
	//   - stream.graphics.g4dn.8xlarge
	//   - stream.graphics.g4dn.12xlarge
	//   - stream.graphics.g4dn.16xlarge
	//   - stream.graphics-pro.4xlarge
	//   - stream.graphics-pro.8xlarge
	//   - stream.graphics-pro.16xlarge
	// The following instance types are available for Elastic fleets:
	//   - stream.standard.small
	//   - stream.standard.medium
	//   - stream.standard.large
	//   - stream.standard.xlarge
	//   - stream.standard.2xlarge
	InstanceType *string
	
	// The maximum number of concurrent sessions for a fleet.
	MaxConcurrentSessions *int32
	
	// The maximum amount of time that a streaming session can remain active, in
	// seconds. If users are still connected to a streaming instance five minutes
	// before this limit is reached, they are prompted to save any open documents
	// before being disconnected. After this time elapses, the instance is terminated
	// and replaced by a new instance. Specify a value between 600 and 432000.
	MaxUserDurationInSeconds *int32
	
	// A unique name for the fleet.
	Name *string
	
	// The platform of the fleet. WINDOWS_SERVER_2019 and AMAZON_LINUX2 are supported
	// for Elastic fleets.
	Platform types.PlatformType
	
	// The S3 location of the session scripts configuration zip file. This only
	// applies to Elastic fleets.
	SessionScriptS3Location *types.S3Location
	
	// The AppStream 2.0 view that is displayed to your users when they stream from
	// the fleet. When APP is specified, only the windows of applications opened by
	// users display. When DESKTOP is specified, the standard desktop that is provided
	// by the operating system displays. The default value is APP .
	StreamView types.StreamView
	
	// The USB device filter strings that specify which USB devices a user can
	// redirect to the fleet streaming session, when using the Windows native client.
	// This is allowed but not required for Elastic fleets.
	UsbDeviceFilterStrings []string
	
	// The VPC configuration for the fleet. This is required for Elastic fleets, but
	// not required for other fleet types. Elastic fleets require that you specify at
	// least two subnets in different availability zones.
	VpcConfig *types.VpcConfig
	
	noSmithyDocumentSerde
}

type UpdateFleetOutput struct {
	
	// Information about the fleet.
	Fleet *types.Fleet
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleet{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleet{}, middleware.After)
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
	if err = addUpdateFleetResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateFleetValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleet(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "appstream",
	OperationName: "UpdateFleet",
	}
}

type opUpdateFleetResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateFleetResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateFleetResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "appstream"
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
	        signingName = "appstream"
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
	        v4aScheme.SigningName = aws.String("appstream")
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

func addUpdateFleetResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateFleetResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
