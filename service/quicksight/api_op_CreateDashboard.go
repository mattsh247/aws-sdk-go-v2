// Code generated by smithy-go-codegen DO NOT EDIT.


package quicksight

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
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a dashboard from either a template or directly with a
// DashboardDefinition . To first create a template, see the CreateTemplate (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateTemplate.html)
// API operation. A dashboard is an entity in Amazon QuickSight that identifies
// Amazon QuickSight reports, created from analyses. You can share Amazon
// QuickSight dashboards. With the right permissions, you can create scheduled
// email reports from them. If you have the correct permissions, you can create a
// dashboard from a template that exists in a different Amazon Web Services
// account.
func (c *Client) CreateDashboard(ctx context.Context, params *CreateDashboardInput, optFns ...func(*Options)) (*CreateDashboardOutput, error) {
	if params == nil { params = &CreateDashboardInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateDashboard", params, optFns, c.addOperationCreateDashboardMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDashboardInput struct {
	
	// The ID of the Amazon Web Services account where you want to create the
	// dashboard.
	//
	// This member is required.
	AwsAccountId *string
	
	// The ID for the dashboard, also added to the IAM policy.
	//
	// This member is required.
	DashboardId *string
	
	// The display name of the dashboard.
	//
	// This member is required.
	Name *string
	
	// Options for publishing the dashboard when you create it:
	//   - AvailabilityStatus for AdHocFilteringOption - This status can be either
	//   ENABLED or DISABLED . When this is set to DISABLED , Amazon QuickSight
	//   disables the left filter pane on the published dashboard, which can be used for
	//   ad hoc (one-time) filtering. This option is ENABLED by default.
	//   - AvailabilityStatus for ExportToCSVOption - This status can be either ENABLED
	//   or DISABLED . The visual option to export data to .CSV format isn't enabled
	//   when this is set to DISABLED . This option is ENABLED by default.
	//   - VisibilityState for SheetControlsOption - This visibility state can be
	//   either COLLAPSED or EXPANDED . This option is COLLAPSED by default.
	DashboardPublishOptions *types.DashboardPublishOptions
	
	// The definition of a dashboard. A definition is the data model of all features
	// in a Dashboard, Template, or Analysis. Either a SourceEntity or a Definition
	// must be provided in order for the request to be valid.
	Definition *types.DashboardVersionDefinition
	
	// The parameters for the creation of the dashboard, which you want to use to
	// override the default settings. A dashboard can have any type of parameters, and
	// some parameters might accept multiple values.
	Parameters *types.Parameters
	
	// A structure that contains the permissions of the dashboard. You can use this
	// structure for granting permissions by providing a list of IAM action information
	// for each principal ARN. To specify no permissions, omit the permissions list.
	Permissions []types.ResourcePermission
	
	// The entity that you are using as a source when you create the dashboard. In
	// SourceEntity , you specify the type of object you're using as source. You can
	// only create a dashboard from a template, so you use a SourceTemplate entity. If
	// you need to create a dashboard from an analysis, first convert the analysis to a
	// template by using the CreateTemplate (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_CreateTemplate.html)
	// API operation. For SourceTemplate , specify the Amazon Resource Name (ARN) of
	// the source template. The SourceTemplate ARN can contain any Amazon Web Services
	// account and any Amazon QuickSight-supported Amazon Web Services Region. Use the
	// DataSetReferences entity within SourceTemplate to list the replacement datasets
	// for the placeholders listed in the original. The schema in each dataset must
	// match its placeholder. Either a SourceEntity or a Definition must be provided
	// in order for the request to be valid.
	SourceEntity *types.DashboardSourceEntity
	
	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the dashboard.
	Tags []types.Tag
	
	// The Amazon Resource Name (ARN) of the theme that is being used for this
	// dashboard. If you add a value for this field, it overrides the value that is
	// used in the source entity. The theme ARN must exist in the same Amazon Web
	// Services account where you create the dashboard.
	ThemeArn *string
	
	// A description for the first version of the dashboard being created.
	VersionDescription *string
	
	noSmithyDocumentSerde
}

type CreateDashboardOutput struct {
	
	// The ARN of the dashboard.
	Arn *string
	
	// The status of the dashboard creation request.
	CreationStatus types.ResourceStatus
	
	// The ID for the dashboard.
	DashboardId *string
	
	// The Amazon Web Services request ID for this operation.
	RequestId *string
	
	// The HTTP status of the request.
	Status int32
	
	// The ARN of the dashboard, including the version number of the first version
	// that is created.
	VersionArn *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDashboard{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDashboard{}, middleware.After)
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
	if err = addCreateDashboardResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateDashboardValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDashboard(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "quicksight",
	OperationName: "CreateDashboard",
	}
}

type opCreateDashboardResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateDashboardResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateDashboardResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "quicksight"
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
	        signingName = "quicksight"
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
	        v4aScheme.SigningName = aws.String("quicksight")
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

func addCreateDashboardResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateDashboardResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
