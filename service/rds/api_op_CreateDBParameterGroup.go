// Code generated by smithy-go-codegen DO NOT EDIT.


package rds

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
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a new DB parameter group. A DB parameter group is initially created
// with the default parameters for the database engine used by the DB instance. To
// provide custom values for any of the parameters, you must modify the group after
// creating it using ModifyDBParameterGroup . Once you've created a DB parameter
// group, you need to associate it with your DB instance using ModifyDBInstance .
// When you associate a new DB parameter group with a running DB instance, you need
// to reboot the DB instance without failover for the new DB parameter group and
// associated settings to take effect. This command doesn't apply to RDS Custom.
// After you create a DB parameter group, you should wait at least 5 minutes before
// creating your first DB instance that uses that DB parameter group as the default
// parameter group. This allows Amazon RDS to fully complete the create action
// before the parameter group is used as the default for a new DB instance. This is
// especially important for parameters that are critical when creating the default
// database for a DB instance, such as the character set for the default database
// defined by the character_set_database parameter. You can use the Parameter
// Groups option of the Amazon RDS console (https://console.aws.amazon.com/rds/)
// or the DescribeDBParameters command to verify that your DB parameter group has
// been created or modified.
func (c *Client) CreateDBParameterGroup(ctx context.Context, params *CreateDBParameterGroupInput, optFns ...func(*Options)) (*CreateDBParameterGroupOutput, error) {
	if params == nil { params = &CreateDBParameterGroupInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateDBParameterGroup", params, optFns, c.addOperationCreateDBParameterGroupMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBParameterGroupInput struct {
	
	// The DB parameter group family name. A DB parameter group can be associated with
	// one and only one DB parameter group family, and can be applied only to a DB
	// instance running a database engine and engine version compatible with that DB
	// parameter group family. To list all of the available parameter group families
	// for a DB engine, use the following command: aws rds describe-db-engine-versions
	// --query "DBEngineVersions[].DBParameterGroupFamily" --engine For example, to
	// list all of the available parameter group families for the MySQL DB engine, use
	// the following command: aws rds describe-db-engine-versions --query
	// "DBEngineVersions[].DBParameterGroupFamily" --engine mysql The output contains
	// duplicates. The following are the valid DB engine values:
	//   - aurora-mysql
	//   - aurora-postgresql
	//   - mariadb
	//   - mysql
	//   - oracle-ee
	//   - oracle-ee-cdb
	//   - oracle-se2
	//   - oracle-se2-cdb
	//   - postgres
	//   - sqlserver-ee
	//   - sqlserver-se
	//   - sqlserver-ex
	//   - sqlserver-web
	//
	// This member is required.
	DBParameterGroupFamily *string
	
	// The name of the DB parameter group. Constraints:
	//   - Must be 1 to 255 letters, numbers, or hyphens.
	//   - First character must be a letter
	//   - Can't end with a hyphen or contain two consecutive hyphens
	// This value is stored as a lowercase string.
	//
	// This member is required.
	DBParameterGroupName *string
	
	// The description for the DB parameter group.
	//
	// This member is required.
	Description *string
	
	// Tags to assign to the DB parameter group.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateDBParameterGroupOutput struct {
	
	// Contains the details of an Amazon RDS DB parameter group. This data type is
	// used as a response element in the DescribeDBParameterGroups action.
	DBParameterGroup *types.DBParameterGroup
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBParameterGroup{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBParameterGroup{}, middleware.After)
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
	if err = addCreateDBParameterGroupResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateDBParameterGroupValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBParameterGroup(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rds",
	OperationName: "CreateDBParameterGroup",
	}
}

type opCreateDBParameterGroupResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateDBParameterGroupResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateDBParameterGroupResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "rds"
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
	        signingName = "rds"
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
	        v4aScheme.SigningName = aws.String("rds")
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

func addCreateDBParameterGroupResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateDBParameterGroupResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
