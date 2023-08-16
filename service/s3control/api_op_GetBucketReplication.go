// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation gets an Amazon S3 on Outposts bucket's replication
// configuration. To get an S3 bucket's replication configuration, see
// GetBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html)
// in the Amazon S3 API Reference. Returns the replication configuration of an S3
// on Outposts bucket. For more information about S3 on Outposts, see Using Amazon
// S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
// in the Amazon S3 User Guide. For information about S3 replication on Outposts
// configuration, see Replicating objects for S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html)
// in the Amazon S3 User Guide. It can take a while to propagate PUT or DELETE
// requests for a replication configuration to all S3 on Outposts systems.
// Therefore, the replication configuration that's returned by a GET request soon
// after a PUT or DELETE request might return a more recent result than what's on
// the Outpost. If an Outpost is offline, the delay in updating the replication
// configuration on that Outpost can be significant. This action requires
// permissions for the s3-outposts:GetReplicationConfiguration action. The
// Outposts bucket owner has this permission by default and can grant it to others.
// For more information about permissions, see Setting up IAM with S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html)
// and Managing access to S3 on Outposts bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html)
// in the Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for
// this action require an additional parameter of x-amz-outpost-id to be passed
// with the request. In addition, you must use an S3 on Outposts endpoint hostname
// prefix instead of s3-control . For an example of the request syntax for Amazon
// S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html#API_control_GetBucketReplication_Examples)
// section. If you include the Filter element in a replication configuration, you
// must also include the DeleteMarkerReplication , Status , and Priority elements.
// The response also returns those elements. For information about S3 on Outposts
// replication failure reasons, see Replication failure reasons (https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes)
// in the Amazon S3 User Guide. The following operations are related to
// GetBucketReplication :
//   - PutBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html)
//   - DeleteBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html)
func (c *Client) GetBucketReplication(ctx context.Context, params *GetBucketReplicationInput, optFns ...func(*Options)) (*GetBucketReplicationOutput, error) {
	if params == nil {
		params = &GetBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketReplication", params, optFns, c.addOperationGetBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketReplicationInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the bucket to get the replication information for. For using this
	// parameter with Amazon S3 on Outposts with the REST API, you must specify the
	// name and the x-amz-outpost-id as well. For using this parameter with S3 on
	// Outposts with the Amazon Web Services SDK and CLI, you must specify the ARN of
	// the bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/ . For
	// example, to access the bucket reports through Outpost my-outpost owned by
	// account 123456789012 in Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

type GetBucketReplicationOutput struct {

	// A container for one or more replication rules. A replication configuration must
	// have at least one rule and you can add up to 100 rules. The maximum size of a
	// replication configuration is 128 KB.
	ReplicationConfiguration *types.ReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketReplication{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetBucketReplicationMiddleware(stack); err != nil {
		return err
	}
	if err = addGetBucketReplicationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketReplication(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addGetBucketReplicationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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

func (m *GetBucketReplicationInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *GetBucketReplicationInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opGetBucketReplicationMiddleware struct {
}

func (*endpointPrefix_opGetBucketReplicationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketReplicationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetBucketReplicationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetBucketReplicationMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketReplication",
	}
}

func copyGetBucketReplicationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetBucketReplicationInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetBucketReplicationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getGetBucketReplicationARNMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketReplicationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setGetBucketReplicationARNMember(input interface{}, v string) error {
	in := input.(*GetBucketReplicationInput)
	in.Bucket = &v
	return nil
}
func backFillGetBucketReplicationAccountID(input interface{}, v string) error {
	in := input.(*GetBucketReplicationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetBucketReplicationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getGetBucketReplicationARNMember,
			BackfillAccountID: backFillGetBucketReplicationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setGetBucketReplicationARNMember,
			CopyInput:         copyGetBucketReplicationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opGetBucketReplicationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetBucketReplicationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetBucketReplicationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.AccountId = input.AccountId

	params.Bucket = input.Bucket

	params.RequiresAccountId = ptr.Bool(true)

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
			signingName := "s3"
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
				signingName = "s3"
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
				v4aScheme.SigningName = aws.String("s3")
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

	ctx = smithyhttp.DisableEndpointHostPrefix(ctx, true)

	return next.HandleSerialize(ctx, in)
}

func addGetBucketReplicationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetBucketReplicationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:     options.BaseEndpoint,
			UseArnRegion: options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
