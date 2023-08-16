// Code generated by smithy-go-codegen DO NOT EDIT.


package ssm

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
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// ServiceSetting is an account-level setting for an Amazon Web Services service.
// This setting defines how a user interacts with or uses a service or a feature of
// a service. For example, if an Amazon Web Services service charges money to the
// account based on feature or service usage, then the Amazon Web Services service
// team might create a default setting of "false". This means the user can't use
// this feature unless they change the setting to "true" and intentionally opt in
// for a paid feature. Services map a SettingId object to a setting value. Amazon
// Web Services services teams define the default value for a SettingId . You can't
// create a new SettingId , but you can overwrite the default value if you have the
// ssm:UpdateServiceSetting permission for the setting. Use the GetServiceSetting
// API operation to view the current value. Or, use the ResetServiceSetting to
// change the value back to the original value defined by the Amazon Web Services
// service team. Update the service setting for the account.
func (c *Client) UpdateServiceSetting(ctx context.Context, params *UpdateServiceSettingInput, optFns ...func(*Options)) (*UpdateServiceSettingOutput, error) {
	if params == nil { params = &UpdateServiceSettingInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceSetting", params, optFns, c.addOperationUpdateServiceSettingMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateServiceSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body of the UpdateServiceSetting API operation.
type UpdateServiceSettingInput struct {
	
	// The Amazon Resource Name (ARN) of the service setting to update. For example,
	// arn:aws:ssm:us-east-1:111122223333:servicesetting/ssm/parameter-store/high-throughput-enabled
	// . The setting ID can be one of the following.
	//   - /ssm/managed-instance/default-ec2-instance-management-role
	//   - /ssm/automation/customer-script-log-destination
	//   - /ssm/automation/customer-script-log-group-name
	//   - /ssm/documents/console/public-sharing-permission
	//   - /ssm/managed-instance/activation-tier
	//   - /ssm/opsinsights/opscenter
	//   - /ssm/parameter-store/default-parameter-tier
	//   - /ssm/parameter-store/high-throughput-enabled
	// Permissions to update the
	// /ssm/managed-instance/default-ec2-instance-management-role setting should only
	// be provided to administrators. Implement least privilege access when allowing
	// individuals to configure or modify the Default Host Management Configuration.
	//
	// This member is required.
	SettingId *string
	
	// The new value to specify for the service setting. The following list specifies
	// the available values for each setting.
	//   - /ssm/managed-instance/default-ec2-instance-management-role: The name of an
	//   IAM role
	//   - /ssm/automation/customer-script-log-destination : CloudWatch
	//   - /ssm/automation/customer-script-log-group-name : The name of an Amazon
	//   CloudWatch Logs log group
	//   - /ssm/documents/console/public-sharing-permission : Enable or Disable
	//   - /ssm/managed-instance/activation-tier : standard or advanced
	//   - /ssm/opsinsights/opscenter : Enabled or Disabled
	//   - /ssm/parameter-store/default-parameter-tier : Standard , Advanced ,
	//   Intelligent-Tiering
	//   - /ssm/parameter-store/high-throughput-enabled : true or false
	//
	// This member is required.
	SettingValue *string
	
	noSmithyDocumentSerde
}

// The result body of the UpdateServiceSetting API operation.
type UpdateServiceSettingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServiceSetting{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServiceSetting{}, middleware.After)
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
	if err = addUpdateServiceSettingResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateServiceSettingValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSetting(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "ssm",
	OperationName: "UpdateServiceSetting",
	}
}

type opUpdateServiceSettingResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateServiceSettingResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateServiceSettingResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "ssm"
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
	        signingName = "ssm"
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
	        v4aScheme.SigningName = aws.String("ssm")
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

func addUpdateServiceSettingResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateServiceSettingResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
