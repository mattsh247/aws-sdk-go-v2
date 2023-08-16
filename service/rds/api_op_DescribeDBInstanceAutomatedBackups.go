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

// Displays backups for both current and deleted instances. For example, use this
// operation to find details about automated backups for previously deleted
// instances. Current instances with retention periods greater than zero (0) are
// returned for both the DescribeDBInstanceAutomatedBackups and DescribeDBInstances
// operations. All parameters are optional.
func (c *Client) DescribeDBInstanceAutomatedBackups(ctx context.Context, params *DescribeDBInstanceAutomatedBackupsInput, optFns ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error) {
	if params == nil { params = &DescribeDBInstanceAutomatedBackupsInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeDBInstanceAutomatedBackups", params, optFns, c.addOperationDescribeDBInstanceAutomatedBackupsMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeDBInstanceAutomatedBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Parameter input for DescribeDBInstanceAutomatedBackups.
type DescribeDBInstanceAutomatedBackupsInput struct {
	
	// The Amazon Resource Name (ARN) of the replicated automated backups, for
	// example,
	// arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE .
	// This setting doesn't apply to RDS Custom.
	DBInstanceAutomatedBackupsArn *string
	
	// (Optional) The user-supplied instance identifier. If this parameter is
	// specified, it must match the identifier of an existing DB instance. It returns
	// information from the specific DB instance's automated backup. This parameter
	// isn't case-sensitive.
	DBInstanceIdentifier *string
	
	// The resource ID of the DB instance that is the source of the automated backup.
	// This parameter isn't case-sensitive.
	DbiResourceId *string
	
	// A filter that specifies which resources to return based on status. Supported
	// filters are the following:
	//   - status
	//   - active - Automated backups for current instances.
	//   - creating - Automated backups that are waiting for the first automated
	//   snapshot to be available.
	//   - retained - Automated backups for deleted instances and after backup
	//   replication is stopped.
	//   - db-instance-id - Accepts DB instance identifiers and Amazon Resource Names
	//   (ARNs). The results list includes only information about the DB instance
	//   automated backups identified by these ARNs.
	//   - dbi-resource-id - Accepts DB resource identifiers and Amazon Resource Names
	//   (ARNs). The results list includes only information about the DB instance
	//   resources identified by these ARNs.
	// Returns all resources by default. The status for each resource is specified in
	// the response.
	Filters []types.Filter
	
	// The pagination token provided in the previous request. If this parameter is
	// specified the response includes only records beyond the marker, up to MaxRecords
	// .
	Marker *string
	
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	MaxRecords *int32
	
	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeDBInstanceAutomatedBackups action.
type DescribeDBInstanceAutomatedBackupsOutput struct {
	
	// A list of DBInstanceAutomatedBackup instances.
	DBInstanceAutomatedBackups []types.DBInstanceAutomatedBackup
	
	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBInstanceAutomatedBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBInstanceAutomatedBackups{}, middleware.After)
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
	if err = addDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeDBInstanceAutomatedBackupsValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(options.Region, ), middleware.Before); err != nil {
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

// DescribeDBInstanceAutomatedBackupsAPIClient is a client that implements the
// DescribeDBInstanceAutomatedBackups operation.
type DescribeDBInstanceAutomatedBackupsAPIClient interface {
	DescribeDBInstanceAutomatedBackups(context.Context, *DescribeDBInstanceAutomatedBackupsInput, ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error)
}

var _ DescribeDBInstanceAutomatedBackupsAPIClient = (*Client)(nil)

// DescribeDBInstanceAutomatedBackupsPaginatorOptions is the paginator options for
// DescribeDBInstanceAutomatedBackups
type DescribeDBInstanceAutomatedBackupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	Limit int32
	
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBInstanceAutomatedBackupsPaginator is a paginator for
// DescribeDBInstanceAutomatedBackups
type DescribeDBInstanceAutomatedBackupsPaginator struct {
    options DescribeDBInstanceAutomatedBackupsPaginatorOptions
    client DescribeDBInstanceAutomatedBackupsAPIClient
    params *DescribeDBInstanceAutomatedBackupsInput
    nextToken *string
    firstPage bool
}

// NewDescribeDBInstanceAutomatedBackupsPaginator returns a new
// DescribeDBInstanceAutomatedBackupsPaginator
func NewDescribeDBInstanceAutomatedBackupsPaginator(client DescribeDBInstanceAutomatedBackupsAPIClient, params *DescribeDBInstanceAutomatedBackupsInput, optFns ...func(*DescribeDBInstanceAutomatedBackupsPaginatorOptions)) *DescribeDBInstanceAutomatedBackupsPaginator {
	if params == nil {
	    params = &DescribeDBInstanceAutomatedBackupsInput{}
	}
	
	options := DescribeDBInstanceAutomatedBackupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}
	
	for _, fn := range optFns {
	    fn(&options)
	}
	
	return &DescribeDBInstanceAutomatedBackupsPaginator{
	    options: options,
	    client: client,
	    params: params,
	    firstPage: true,
	    nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBInstanceAutomatedBackupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0 )
}

// NextPage retrieves the next DescribeDBInstanceAutomatedBackups page.
func (p *DescribeDBInstanceAutomatedBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBInstanceAutomatedBackupsOutput, error) {
	if !p.HasMorePages() {
	    return nil, fmt.Errorf("no more pages available")
	}
	
	params := *p.params
	params.Marker = p.nextToken
	
	var limit *int32
	if p.options.Limit > 0 {
	    limit = &p.options.Limit
	}
	params.MaxRecords = limit
	
	result, err := p.client.DescribeDBInstanceAutomatedBackups(ctx, &params, optFns...)
	if err != nil {
	    return nil, err
	}
	p.firstPage = false
	
	prevToken := p.nextToken
	p.nextToken = result.Marker
	
	if p.options.StopOnDuplicateToken &&
	    prevToken != nil &&
	    p.nextToken != nil &&
	    *prevToken == *p.nextToken {
	    p.nextToken = nil
	}
	
	return result, nil
}

func newServiceMetadataMiddleware_opDescribeDBInstanceAutomatedBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "rds",
	OperationName: "DescribeDBInstanceAutomatedBackups",
	}
}

type opDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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

func addDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeDBInstanceAutomatedBackupsResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
