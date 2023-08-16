// Code generated by smithy-go-codegen DO NOT EDIT.


package timestreamwrite

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/timestreamwrite/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns information about the database, including the database name, time that
// the database was created, and the total number of tables found within the
// database. Service quotas apply (https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html)
// . See code sample (https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.describe-db.html)
// for details.
func (c *Client) DescribeDatabase(ctx context.Context, params *DescribeDatabaseInput, optFns ...func(*Options)) (*DescribeDatabaseOutput, error) {
	if params == nil { params = &DescribeDatabaseInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeDatabase", params, optFns, c.addOperationDescribeDatabaseMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatabaseInput struct {
	
	// The name of the Timestream database.
	//
	// This member is required.
	DatabaseName *string
	
	noSmithyDocumentSerde
}

type DescribeDatabaseOutput struct {
	
	// The name of the Timestream table.
	Database *types.Database
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeDatabase{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeDatabase{}, middleware.After)
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
	if err = addOpDescribeDatabaseDiscoverEndpointMiddleware(stack, options, c); err != nil {
	return err
	}
	if err = addDescribeDatabaseResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeDatabaseValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatabase(options.Region, ), middleware.Before); err != nil {
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

func addOpDescribeDatabaseDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func (opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation: c.fetchOpDescribeDatabaseDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired: true,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpDescribeDatabaseDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*DescribeDatabaseInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in
	
	identifierMap := make(map[string]string, 0)
	
	key := fmt.Sprintf("Timestream Write.%v", identifierMap)
	
	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}
	
	discoveryOperationInput := &DescribeEndpointsInput{
	}
	
	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}
	
	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	if err != nil { return internalEndpointDiscovery.WeightedAddress{}, err }
	
	weighted, ok := endpoint.GetValidAddress()
	if !ok { return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")}
	return weighted, nil
}

func newServiceMetadataMiddleware_opDescribeDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "timestream",
	OperationName: "DescribeDatabase",
	}
}

type opDescribeDatabaseResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeDatabaseResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeDatabaseResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "timestream"
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
	        signingName = "timestream"
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
	        v4aScheme.SigningName = aws.String("timestream")
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

func addDescribeDatabaseResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeDatabaseResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
