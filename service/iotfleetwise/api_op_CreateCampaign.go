// Code generated by smithy-go-codegen DO NOT EDIT.


package iotfleetwise

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
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates an orchestration of data collection rules. The Amazon Web Services IoT
// FleetWise Edge Agent software running in vehicles uses campaigns to decide how
// to collect and transfer data to the cloud. You create campaigns in the cloud.
// After you or your team approve campaigns, Amazon Web Services IoT FleetWise
// automatically deploys them to vehicles. For more information, see Collect and
// transfer data with campaigns (https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html)
// in the Amazon Web Services IoT FleetWise Developer Guide.
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil { params = &CreateCampaignInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, c.addOperationCreateCampaignMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCampaignInput struct {
	
	// The data collection scheme associated with the campaign. You can specify a
	// scheme that collects data based on time or an event.
	//
	// This member is required.
	CollectionScheme types.CollectionScheme
	
	// The name of the campaign to create.
	//
	// This member is required.
	Name *string
	
	// (Optional) The Amazon Resource Name (ARN) of the signal catalog to associate
	// with the campaign.
	//
	// This member is required.
	SignalCatalogArn *string
	
	// The ARN of the vehicle or fleet to deploy a campaign to.
	//
	// This member is required.
	TargetArn *string
	
	// (Optional) Whether to compress signals before transmitting data to Amazon Web
	// Services IoT FleetWise. If you don't want to compress the signals, use OFF . If
	// it's not specified, SNAPPY is used. Default: SNAPPY
	Compression types.Compression
	
	// The destination where the campaign sends data. You can choose to send data to
	// be stored in Amazon S3 or Amazon Timestream. Amazon S3 optimizes the cost of
	// data storage and provides additional mechanisms to use vehicle data, such as
	// data lakes, centralized data storage, data processing pipelines, and analytics.
	// You can use Amazon Timestream to access and analyze time series data, and
	// Timestream to query vehicle data so that you can identify trends and patterns.
	DataDestinationConfigs []types.DataDestinationConfig
	
	// (Optional) A list of vehicle attributes to associate with a campaign. Enrich
	// the data with specified vehicle attributes. For example, add make and model to
	// the campaign, and Amazon Web Services IoT FleetWise will associate the data with
	// those attributes as dimensions in Amazon Timestream. You can then query the data
	// against make and model . Default: An empty array
	DataExtraDimensions []string
	
	// An optional description of the campaign to help identify its purpose.
	Description *string
	
	// (Optional) Option for a vehicle to send diagnostic trouble codes to Amazon Web
	// Services IoT FleetWise. If you want to send diagnostic trouble codes, use
	// SEND_ACTIVE_DTCS . If it's not specified, OFF is used. Default: OFF
	DiagnosticsMode types.DiagnosticsMode
	
	// (Optional) The time the campaign expires, in seconds since epoch (January 1,
	// 1970 at midnight UTC time). Vehicle data isn't collected after the campaign
	// expires. Default: 253402214400 (December 31, 9999, 00:00:00 UTC)
	ExpiryTime *time.Time
	
	// (Optional) How long (in milliseconds) to collect raw data after a triggering
	// event initiates the collection. If it's not specified, 0 is used. Default: 0
	PostTriggerCollectionDuration *int64
	
	// (Optional) A number indicating the priority of one campaign over another
	// campaign for a certain vehicle or fleet. A campaign with the lowest value is
	// deployed to vehicles before any other campaigns. If it's not specified, 0 is
	// used. Default: 0
	Priority *int32
	
	// (Optional) A list of information about signals to collect.
	SignalsToCollect []types.SignalInformation
	
	// (Optional) Whether to store collected data after a vehicle lost a connection
	// with the cloud. After a connection is re-established, the data is automatically
	// forwarded to Amazon Web Services IoT FleetWise. If you want to store collected
	// data when a vehicle loses connection with the cloud, use TO_DISK . If it's not
	// specified, OFF is used. Default: OFF
	SpoolingMode types.SpoolingMode
	
	// (Optional) The time, in milliseconds, to deliver a campaign after it was
	// approved. If it's not specified, 0 is used. Default: 0
	StartTime *time.Time
	
	// Metadata that can be used to manage the campaign.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateCampaignOutput struct {
	
	// The ARN of the created campaign.
	Arn *string
	
	// The name of the created campaign.
	Name *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateCampaign{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateCampaign{}, middleware.After)
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
	if err = addCreateCampaignResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateCampaignValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "iotfleetwise",
	OperationName: "CreateCampaign",
	}
}

type opCreateCampaignResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateCampaignResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateCampaignResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "iotfleetwise"
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
	        signingName = "iotfleetwise"
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
	        v4aScheme.SigningName = aws.String("iotfleetwise")
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

func addCreateCampaignResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateCampaignResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
