// Code generated by smithy-go-codegen DO NOT EDIT.


package apigateway

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
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// A feature of the API Gateway control service for creating a new API from an
// external API definition file.
func (c *Client) ImportRestApi(ctx context.Context, params *ImportRestApiInput, optFns ...func(*Options)) (*ImportRestApiOutput, error) {
	if params == nil { params = &ImportRestApiInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "ImportRestApi", params, optFns, c.addOperationImportRestApiMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*ImportRestApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A POST request to import an API to API Gateway using an input of an API
// definition file.
type ImportRestApiInput struct {
	
	// The POST request body containing external API definitions. Currently, only
	// OpenAPI definition JSON/YAML files are supported. The maximum size of the API
	// definition file is 6MB.
	//
	// This member is required.
	Body []byte
	
	// A query parameter to indicate whether to rollback the API creation ( true ) or
	// not ( false ) when a warning is encountered. The default value is false .
	FailOnWarnings bool
	
	// A key-value map of context-specific query string parameters specifying the
	// behavior of different API importing operations. The following shows
	// operation-specific parameters and their supported values. To exclude
	// DocumentationParts from the import, set parameters as ignore=documentation . To
	// configure the endpoint type, set parameters as endpointConfigurationTypes=EDGE ,
	// endpointConfigurationTypes=REGIONAL , or endpointConfigurationTypes=PRIVATE .
	// The default endpoint type is EDGE . To handle imported basepath , set parameters
	// as basepath=ignore , basepath=prepend or basepath=split . For example, the AWS
	// CLI command to exclude documentation from the imported API is: The AWS CLI
	// command to set the regional endpoint on the imported API is:
	Parameters map[string]string
	
	noSmithyDocumentSerde
}

// Represents a REST API.
type ImportRestApiOutput struct {
	
	// The source of the API key for metering requests according to a usage plan.
	// Valid values are: > HEADER to read the API key from the X-API-Key header of a
	// request. AUTHORIZER to read the API key from the UsageIdentifierKey from a
	// custom authorizer.
	ApiKeySource types.ApiKeySourceType
	
	// The list of binary media types supported by the RestApi. By default, the
	// RestApi supports only UTF-8-encoded text payloads.
	BinaryMediaTypes []string
	
	// The timestamp when the API was created.
	CreatedDate *time.Time
	
	// The API's description.
	Description *string
	
	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint bool
	
	// The endpoint configuration of this RestApi showing the endpoint types of the
	// API.
	EndpointConfiguration *types.EndpointConfiguration
	
	// The API's identifier. This identifier is unique across all of your APIs in API
	// Gateway.
	Id *string
	
	// A nullable integer that is used to enable compression (with non-negative
	// between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a
	// null value) on an API. When compression is enabled, compression or decompression
	// is not applied on the payload if the payload size is smaller than this value.
	// Setting it to zero allows compression for any payload size.
	MinimumCompressionSize *int32
	
	// The API's name.
	Name *string
	
	// A stringified JSON policy document that applies to this RestApi regardless of
	// the caller and Method configuration.
	Policy *string
	
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string
	
	// A version identifier for the API.
	Version *string
	
	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationImportRestApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportRestApi{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportRestApi{}, middleware.After)
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
	if err = addImportRestApiResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpImportRestApiValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportRestApi(options.Region, ), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opImportRestApi(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "apigateway",
	OperationName: "ImportRestApi",
	}
}

type opImportRestApiResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opImportRestApiResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opImportRestApiResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "apigateway"
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
	        signingName = "apigateway"
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
	        v4aScheme.SigningName = aws.String("apigateway")
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

func addImportRestApiResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opImportRestApiResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
