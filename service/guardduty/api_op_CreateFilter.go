// Code generated by smithy-go-codegen DO NOT EDIT.


package guardduty

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
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a filter using the specified finding criteria. The maximum number of
// saved filters per Amazon Web Services account per Region is 100. For more
// information, see Quotas for GuardDuty (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_limits.html)
// .
func (c *Client) CreateFilter(ctx context.Context, params *CreateFilterInput, optFns ...func(*Options)) (*CreateFilterOutput, error) {
	if params == nil { params = &CreateFilterInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateFilter", params, optFns, c.addOperationCreateFilterMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFilterInput struct {
	
	// The ID of the detector belonging to the GuardDuty account that you want to
	// create a filter for.
	//
	// This member is required.
	DetectorId *string
	
	// Represents the criteria to be used in the filter for querying findings. You can
	// only use the following attributes to query findings:
	//   - accountId
	//   - id
	//   - region
	//   - severity To filter on the basis of severity, the API and CLI use the
	//   following input list for the FindingCriteria (https://docs.aws.amazon.com/guardduty/latest/APIReference/API_FindingCriteria.html)
	//   condition:
	//   - Low: ["1", "2", "3"]
	//   - Medium: ["4", "5", "6"]
	//   - High: ["7", "8", "9"] For more information, see Severity levels for
	//   GuardDuty findings (https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_findings.html#guardduty_findings-severity)
	//   .
	//   - type
	//   - updatedAt Type: ISO 8601 string format: YYYY-MM-DDTHH:MM:SS.SSSZ or
	//   YYYY-MM-DDTHH:MM:SSZ depending on whether the value contains milliseconds.
	//   - resource.accessKeyDetails.accessKeyId
	//   - resource.accessKeyDetails.principalId
	//   - resource.accessKeyDetails.userName
	//   - resource.accessKeyDetails.userType
	//   - resource.instanceDetails.iamInstanceProfile.id
	//   - resource.instanceDetails.imageId
	//   - resource.instanceDetails.instanceId
	//   - resource.instanceDetails.tags.key
	//   - resource.instanceDetails.tags.value
	//   - resource.instanceDetails.networkInterfaces.ipv6Addresses
	//   -
	//   resource.instanceDetails.networkInterfaces.privateIpAddresses.privateIpAddress
	//   - resource.instanceDetails.networkInterfaces.publicDnsName
	//   - resource.instanceDetails.networkInterfaces.publicIp
	//   - resource.instanceDetails.networkInterfaces.securityGroups.groupId
	//   - resource.instanceDetails.networkInterfaces.securityGroups.groupName
	//   - resource.instanceDetails.networkInterfaces.subnetId
	//   - resource.instanceDetails.networkInterfaces.vpcId
	//   - resource.instanceDetails.outpostArn
	//   - resource.resourceType
	//   - resource.s3BucketDetails.publicAccess.effectivePermissions
	//   - resource.s3BucketDetails.name
	//   - resource.s3BucketDetails.tags.key
	//   - resource.s3BucketDetails.tags.value
	//   - resource.s3BucketDetails.type
	//   - service.action.actionType
	//   - service.action.awsApiCallAction.api
	//   - service.action.awsApiCallAction.callerType
	//   - service.action.awsApiCallAction.errorCode
	//   - service.action.awsApiCallAction.remoteIpDetails.city.cityName
	//   - service.action.awsApiCallAction.remoteIpDetails.country.countryName
	//   - service.action.awsApiCallAction.remoteIpDetails.ipAddressV4
	//   - service.action.awsApiCallAction.remoteIpDetails.organization.asn
	//   - service.action.awsApiCallAction.remoteIpDetails.organization.asnOrg
	//   - service.action.awsApiCallAction.serviceName
	//   - service.action.dnsRequestAction.domain
	//   - service.action.networkConnectionAction.blocked
	//   - service.action.networkConnectionAction.connectionDirection
	//   - service.action.networkConnectionAction.localPortDetails.port
	//   - service.action.networkConnectionAction.protocol
	//   - service.action.networkConnectionAction.remoteIpDetails.city.cityName
	//   - service.action.networkConnectionAction.remoteIpDetails.country.countryName
	//   - service.action.networkConnectionAction.remoteIpDetails.ipAddressV4
	//   - service.action.networkConnectionAction.remoteIpDetails.organization.asn
	//   - service.action.networkConnectionAction.remoteIpDetails.organization.asnOrg
	//   - service.action.networkConnectionAction.remotePortDetails.port
	//   - service.action.awsApiCallAction.remoteAccountDetails.affiliated
	//   - service.action.kubernetesApiCallAction.remoteIpDetails.ipAddressV4
	//   - service.action.kubernetesApiCallAction.requestUri
	//   - service.action.networkConnectionAction.localIpDetails.ipAddressV4
	//   - service.action.networkConnectionAction.protocol
	//   - service.action.awsApiCallAction.serviceName
	//   - service.action.awsApiCallAction.remoteAccountDetails.accountId
	//   - service.additionalInfo.threatListName
	//   - service.resourceRole
	//   - resource.eksClusterDetails.name
	//   - resource.kubernetesDetails.kubernetesWorkloadDetails.name
	//   - resource.kubernetesDetails.kubernetesWorkloadDetails.namespace
	//   - resource.kubernetesDetails.kubernetesUserDetails.username
	//   - resource.kubernetesDetails.kubernetesWorkloadDetails.containers.image
	//   - resource.kubernetesDetails.kubernetesWorkloadDetails.containers.imagePrefix
	//   - service.ebsVolumeScanDetails.scanId
	//   -
	//   service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.name
	//
	//   -
	//   service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.severity
	//
	//   -
	//   service.ebsVolumeScanDetails.scanDetections.threatDetectedByName.threatNames.filePaths.hash
	//
	//   - resource.ecsClusterDetails.name
	//   - resource.ecsClusterDetails.taskDetails.containers.image
	//   - resource.ecsClusterDetails.taskDetails.definitionArn
	//   - resource.containerDetails.image
	//   - resource.rdsDbInstanceDetails.dbInstanceIdentifier
	//   - resource.rdsDbInstanceDetails.dbClusterIdentifier
	//   - resource.rdsDbInstanceDetails.engine
	//   - resource.rdsDbUserDetails.user
	//   - resource.rdsDbInstanceDetails.tags.key
	//   - resource.rdsDbInstanceDetails.tags.value
	//   - service.runtimeDetails.process.executableSha256
	//   - service.runtimeDetails.process.name
	//   - service.runtimeDetails.process.name
	//   - resource.lambdaDetails.functionName
	//   - resource.lambdaDetails.functionArn
	//   - resource.lambdaDetails.tags.key
	//   - resource.lambdaDetails.tags.value
	//
	// This member is required.
	FindingCriteria *types.FindingCriteria
	
	// The name of the filter. Valid characters include period (.), underscore (_),
	// dash (-), and alphanumeric characters. A whitespace is considered to be an
	// invalid character.
	//
	// This member is required.
	Name *string
	
	// Specifies the action that is to be applied to the findings that match the
	// filter.
	Action types.FilterAction
	
	// The idempotency token for the create request.
	ClientToken *string
	
	// The description of the filter. Valid characters include alphanumeric
	// characters, and special characters such as hyphen, period, colon, underscore,
	// parentheses ( { } , [ ] , and ( ) ), forward slash, horizontal tab, vertical
	// tab, newline, form feed, return, and whitespace.
	Description *string
	
	// Specifies the position of the filter in the list of current filters. Also
	// specifies the order in which this filter is applied to the findings.
	Rank int32
	
	// The tags to be added to a new filter resource.
	Tags map[string]string
	
	noSmithyDocumentSerde
}

type CreateFilterOutput struct {
	
	// The name of the successfully created filter.
	//
	// This member is required.
	Name *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFilter{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFilter{}, middleware.After)
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
	if err = addCreateFilterResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addIdempotencyToken_opCreateFilterMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateFilterValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFilter(options.Region, ), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFilter struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFilter) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFilter) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}
	
	input, ok := in.Parameters.(*CreateFilterInput)
	if !ok { return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFilterInput ")}
	
	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		 if err != nil { return out, metadata, err }
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFilterMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFilter{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "guardduty",
	OperationName: "CreateFilter",
	}
}

type opCreateFilterResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateFilterResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateFilterResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "guardduty"
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
	        signingName = "guardduty"
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
	        v4aScheme.SigningName = aws.String("guardduty")
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

func addCreateFilterResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateFilterResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
