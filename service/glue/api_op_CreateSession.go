// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new session.
func (c *Client) CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) {
	if params == nil {
		params = &CreateSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSession", params, optFns, c.addOperationCreateSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to create a new session.
type CreateSessionInput struct {

	// The SessionCommand that runs the job.
	//
	// This member is required.
	Command *types.SessionCommand

	// The ID of the session request.
	//
	// This member is required.
	Id *string

	// The IAM Role ARN
	//
	// This member is required.
	Role *string

	// The number of connections to use for the session.
	Connections *types.ConnectionsList

	// A map array of key-value pairs. Max is 75 pairs.
	DefaultArguments map[string]string

	// The description of the session.
	Description *string

	// The Glue version determines the versions of Apache Spark and Python that Glue
	// supports. The GlueVersion must be greater than 2.0.
	GlueVersion *string

	// The number of minutes when idle before session times out. Default for Spark ETL
	// jobs is value of Timeout. Consult the documentation for other job types.
	IdleTimeout *int32

	// The number of Glue data processing units (DPUs) that can be allocated when the
	// job runs. A DPU is a relative measure of processing power that consists of 4
	// vCPUs of compute capacity and 16 GB memory.
	MaxCapacity *float64

	// The number of workers of a defined WorkerType to use for the session.
	NumberOfWorkers *int32

	// The origin of the request.
	RequestOrigin *string

	// The name of the SecurityConfiguration structure to be used with the session
	SecurityConfiguration *string

	// The map of key value pairs (tags) belonging to the session.
	Tags map[string]string

	// The number of minutes before session times out. Default for Spark ETL jobs is
	// 48 hours (2880 minutes), the maximum session lifetime for this job type. Consult
	// the documentation for other job types.
	Timeout *int32

	// The type of predefined worker that is allocated when a job runs. Accepts a
	// value of G.1X, G.2X, G.4X, or G.8X for Spark jobs. Accepts the value Z.2X for
	// Ray notebooks.
	//   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPUs, 16 GB of
	//   memory) with 84GB disk (approximately 34GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPUs, 32 GB of
	//   memory) with 128GB disk (approximately 77GB free), and provides 1 executor per
	//   worker. We recommend this worker type for workloads such as data transforms,
	//   joins, and queries, to offers a scalable and cost effective way to run most
	//   jobs.
	//   - For the G.4X worker type, each worker maps to 4 DPU (16 vCPUs, 64 GB of
	//   memory) with 256GB disk (approximately 235GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs in the following
	//   Amazon Web Services Regions: US East (Ohio), US East (N. Virginia), US West
	//   (Oregon), Asia Pacific (Singapore), Asia Pacific (Sydney), Asia Pacific (Tokyo),
	//   Canada (Central), Europe (Frankfurt), Europe (Ireland), and Europe (Stockholm).
	//   - For the G.8X worker type, each worker maps to 8 DPU (32 vCPUs, 128 GB of
	//   memory) with 512GB disk (approximately 487GB free), and provides 1 executor per
	//   worker. We recommend this worker type for jobs whose workloads contain your most
	//   demanding transforms, aggregations, joins, and queries. This worker type is
	//   available only for Glue version 3.0 or later Spark ETL jobs, in the same Amazon
	//   Web Services Regions as supported for the G.4X worker type.
	//   - For the Z.2X worker type, each worker maps to 2 M-DPU (8vCPUs, 64 GB of
	//   memory) with 128 GB disk (approximately 120GB free), and provides up to 8 Ray
	//   workers based on the autoscaler.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type CreateSessionOutput struct {

	// Returns the session object in the response.
	Session *types.Session

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSession{}, middleware.After)
	if err != nil {
		return err
	}
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
	if err = addCreateSessionResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateSession",
	}
}

type opCreateSessionResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opCreateSessionResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateSessionResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "glue"
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
				signingName = "glue"
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
				v4aScheme.SigningName = aws.String("glue")
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

func addCreateSessionResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opCreateSessionResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
