// Code generated by smithy-go-codegen DO NOT EDIT.


package backup

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
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Returns backup job details for the specified BackupJobId .
func (c *Client) DescribeBackupJob(ctx context.Context, params *DescribeBackupJobInput, optFns ...func(*Options)) (*DescribeBackupJobOutput, error) {
	if params == nil { params = &DescribeBackupJobInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DescribeBackupJob", params, optFns, c.addOperationDescribeBackupJobMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DescribeBackupJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupJobInput struct {
	
	// Uniquely identifies a request to Backup to back up a resource.
	//
	// This member is required.
	BackupJobId *string
	
	noSmithyDocumentSerde
}

type DescribeBackupJobOutput struct {
	
	// Returns the account ID that owns the backup job.
	AccountId *string
	
	// Uniquely identifies a request to Backup to back up a resource.
	BackupJobId *string
	
	// Represents the options specified as part of backup plan or on-demand backup job.
	BackupOptions map[string]string
	
	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64
	
	// Represents the actual backup type selected for a backup job. For example, if a
	// successful Windows Volume Shadow Copy Service (VSS) backup was taken, BackupType
	// returns "WindowsVSS" . If BackupType is empty, then the backup type was a
	// regular backup.
	BackupType *string
	
	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault .
	BackupVaultArn *string
	
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	BackupVaultName *string
	
	// The size in bytes transferred to a backup vault at the time that the job status
	// was queried.
	BytesTransferred *int64
	
	// This returns the statistics of the included child (nested) backup jobs.
	ChildJobsInState map[string]int64
	
	// The date and time that a job to create a backup job is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CompletionDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time
	
	// Contains identifying information about the creation of a backup job, including
	// the BackupPlanArn , BackupPlanId , BackupPlanVersion , and BackupRuleId of the
	// backup plan that is used to create it.
	CreatedBy *types.RecoveryPointCreator
	
	// The date and time that a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time
	
	// The date and time that a job to back up resources is expected to be completed,
	// in Unix format and Coordinated Universal Time (UTC). The value of
	// ExpectedCompletionDate is accurate to milliseconds. For example, the value
	// 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	ExpectedCompletionDate *time.Time
	
	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access .
	IamRoleArn *string
	
	// This returns the boolean value that a backup job is a parent (composite) job.
	IsParent bool
	
	// This returns the number of child (nested) backup jobs.
	NumberOfChildJobs *int64
	
	// This returns the parent (composite) resource backup job ID.
	ParentJobId *string
	
	// Contains an estimated percentage that is complete of a job at the time the job
	// status was queried.
	PercentDone *string
	
	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string
	
	// An ARN that uniquely identifies a saved resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string
	
	// This is the non-unique name of the resource that belongs to the specified
	// backup.
	ResourceName *string
	
	// The type of Amazon Web Services resource to be backed up; for example, an
	// Amazon Elastic Block Store (Amazon EBS) volume or an Amazon Relational Database
	// Service (Amazon RDS) database.
	ResourceType *string
	
	// Specifies the time in Unix format and Coordinated Universal Time (UTC) when a
	// backup job must be started before it is canceled. The value is calculated by
	// adding the start window to the scheduled time. So if the scheduled time were
	// 6:00 PM and the start window is 2 hours, the StartBy time would be 8:00 PM on
	// the date specified. The value of StartBy is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	StartBy *time.Time
	
	// The current state of a backup job.
	State types.BackupJobState
	
	// A detailed message explaining the status of the job to back up a resource.
	StatusMessage *string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBackupJob{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBackupJob{}, middleware.After)
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
	if err = addDescribeBackupJobResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDescribeBackupJobValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackupJob(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBackupJob(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "backup",
	OperationName: "DescribeBackupJob",
	}
}

type opDescribeBackupJobResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDescribeBackupJobResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDescribeBackupJobResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "backup"
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
	        signingName = "backup"
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
	        v4aScheme.SigningName = aws.String("backup")
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

func addDescribeBackupJobResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDescribeBackupJobResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
