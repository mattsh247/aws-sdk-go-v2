// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This action puts tags on an Amazon S3 on Outposts bucket. To put tags on an S3
// bucket, see PutBucketTagging (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html)
// in the Amazon S3 API Reference. Sets the tags for an S3 on Outposts bucket. For
// more information, see Using Amazon S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
// in the Amazon S3 User Guide. Use tags to organize your Amazon Web Services bill
// to reflect your own cost structure. To do this, sign up to get your Amazon Web
// Services account bill with tag key values included. Then, to see the cost of
// combined resources, organize your billing information according to resources
// with the same tag key values. For example, you can tag several resources with a
// specific application name, and then organize your billing information to see the
// total cost of that application across several services. For more information,
// see Cost allocation and tagging (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// . Within a bucket, if you add a tag that has the same key as an existing tag,
// the new value overwrites the old value. For more information, see Using cost
// allocation in Amazon S3 bucket tags (https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html)
// . To use this action, you must have permissions to perform the
// s3-outposts:PutBucketTagging action. The Outposts bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see Permissions Related to Bucket Subresource
// Operations (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing access permissions to your Amazon S3 resources (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// . PutBucketTagging has the following special errors:
//   - Error code: InvalidTagError
//   - Description: The tag provided was not a valid tag. This error can occur if
//     the tag did not pass input validation. For information about tag restrictions,
//     see User-Defined Tag Restrictions (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
//     and Amazon Web Services-Generated Cost Allocation Tag Restrictions (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html)
//     .
//   - Error code: MalformedXMLError
//   - Description: The XML provided does not match the schema.
//   - Error code: OperationAbortedError
//   - Description: A conflicting conditional action is currently in progress
//     against this resource. Try again.
//   - Error code: InternalError
//   - Description: The service was unable to apply the provided tag to the
//     bucket.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html#API_control_PutBucketTagging_Examples)
// section. The following actions are related to PutBucketTagging :
//   - GetBucketTagging (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html)
//   - DeleteBucketTagging (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html)
func (c *Client) PutBucketTagging(ctx context.Context, params *PutBucketTaggingInput, optFns ...func(*Options)) (*PutBucketTaggingOutput, error) {
	if params == nil {
		params = &PutBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketTagging", params, optFns, c.addOperationPutBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketTaggingInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The Amazon Resource Name (ARN) of the bucket. For using this parameter with
	// Amazon S3 on Outposts with the REST API, you must specify the name and the
	// x-amz-outpost-id as well. For using this parameter with S3 on Outposts with the
	// Amazon Web Services SDK and CLI, you must specify the ARN of the bucket accessed
	// in the format arn:aws:s3-outposts:::outpost//bucket/ . For example, to access
	// the bucket reports through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	//
	//
	// This member is required.
	Tagging *types.Tagging

	noSmithyDocumentSerde
}

type PutBucketTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketTagging{}, middleware.After)
	if err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketTaggingResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutBucketTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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
	if err = addPutBucketTaggingEndpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketTagging",
	}
}

func copyPutBucketTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutBucketTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutBucketTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getPutBucketTaggingARNMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketTaggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setPutBucketTaggingARNMember(input interface{}, v string) error {
	in := input.(*PutBucketTaggingInput)
	in.Bucket = &v
	return nil
}
func backFillPutBucketTaggingAccountID(input interface{}, v string) error {
	in := input.(*PutBucketTaggingInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutBucketTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutBucketTaggingARNMember,
			BackfillAccountID: backFillPutBucketTaggingAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutBucketTaggingARNMember,
			CopyInput:         copyPutBucketTaggingInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}

type opPutBucketTaggingEndpointDisableHTTPSMiddleware struct {
	EndpointDisableHTTPS bool
}

func (*opPutBucketTaggingEndpointDisableHTTPSMiddleware) ID() string {
	return "opPutBucketTaggingEndpointDisableHTTPSMiddleware"
}

func (m *opPutBucketTaggingEndpointDisableHTTPSMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointDisableHTTPS {
		req.URL.Scheme = "http"
	}

	return next.HandleSerialize(ctx, in)

}
func addPutBucketTaggingEndpointDisableHTTPSMiddleware(stack *middleware.Stack, o Options) error {
	return stack.Serialize.Insert(&opPutBucketTaggingEndpointDisableHTTPSMiddleware{
		EndpointDisableHTTPS: o.EndpointOptions.DisableHTTPS,
	}, "opPutBucketTaggingResolveEndpointMiddleware", middleware.After)
}

type opPutBucketTaggingResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opPutBucketTaggingResolveEndpointMiddleware) ID() string {
	return "opPutBucketTaggingResolveEndpointMiddleware"
}

func (m *opPutBucketTaggingResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketTaggingInput)
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
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
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
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
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
			var signingName string
			if v4aScheme.SigningName == nil {
				signingName = "s3"
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addPutBucketTaggingResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opPutBucketTaggingResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:       options.Region,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:     options.BaseEndpoint,
			UseArnRegion: options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
