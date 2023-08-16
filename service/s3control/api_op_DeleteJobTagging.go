// Code generated by smithy-go-codegen DO NOT EDIT.


package s3control

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"context"
	"errors"
	"fmt"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Removes the entire tag set from the specified S3 Batch Operations job. To use
// the DeleteJobTagging operation, you must have permission to perform the
// s3:DeleteJobTagging action. For more information, see Controlling access and
// labeling jobs using tags (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
// in the Amazon S3 User Guide. Related actions include:
//   - CreateJob (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//   - GetJobTagging (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html)
//   - PutJobTagging (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html)
func (c *Client) DeleteJobTagging(ctx context.Context, params *DeleteJobTaggingInput, optFns ...func(*Options)) (*DeleteJobTaggingOutput, error) {
	if params == nil { params = &DeleteJobTaggingInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "DeleteJobTagging", params, optFns, c.addOperationDeleteJobTaggingMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*DeleteJobTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobTaggingInput struct {
	
	// The Amazon Web Services account ID associated with the S3 Batch Operations job.
	//
	// This member is required.
	AccountId *string
	
	// The ID for the S3 Batch Operations job whose tags you want to delete.
	//
	// This member is required.
	JobId *string
	
	noSmithyDocumentSerde
}

type DeleteJobTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteJobTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteJobTagging{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteJobTagging{}, middleware.After)
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
	return err
	}
	if err = addEndpointPrefix_opDeleteJobTaggingMiddleware(stack); err != nil {
	return err
	}
	if err = addDeleteJobTaggingResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpDeleteJobTaggingValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobTagging(options.Region, ), middleware.Before); err != nil {
	return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
	return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
	return err
	}
	if err = addDeleteJobTaggingUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opDeleteJobTaggingMiddleware struct {
}

func (*endpointPrefix_opDeleteJobTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteJobTaggingMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}
	
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}
	
	input, ok := in.Parameters.(*DeleteJobTaggingInput)
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
func addEndpointPrefix_opDeleteJobTaggingMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteJobTaggingMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteJobTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "s3",
	OperationName: "DeleteJobTagging",
	}
}

func copyDeleteJobTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteJobTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteJobTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillDeleteJobTaggingAccountID (input interface{}, v string) error {
in := input.(*DeleteJobTaggingInput)
if in.AccountId != nil {
if !strings.EqualFold(*in.AccountId, v) {
return fmt.Errorf("error backfilling account id") }
return nil }
in.AccountId = &v
return nil }
func addDeleteJobTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
	Accessor : s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
	 BackfillAccountID: nopBackfillAccountIDAccessor,
	GetOutpostIDInput: nopGetOutpostIDFromInput,
	 UpdateARNField: nopSetARNAccessor,
	 CopyInput: copyDeleteJobTaggingInputForUpdateEndpoint,
	 },
	EndpointResolver: options.EndpointResolver,
	 EndpointResolverOptions: options.EndpointOptions,
	UseARNRegion: options.UseARNRegion,
	 })
}

type opDeleteJobTaggingResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opDeleteJobTaggingResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDeleteJobTaggingResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
	            return next.HandleSerialize(ctx, in)
	        }
	
	req, ok := in.Request.(*smithyhttp.Request)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	input, ok := in.Parameters.(*DeleteJobTaggingInput)
	    if !ok {
	        return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	    }
	
	if m.EndpointResolver == nil {
	        return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	    }
	
	params := EndpointParameters{}
	
	m.BuiltInResolver.ResolveBuiltIns(&params)
	
	params.AccountId = input.AccountId
	
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

func addDeleteJobTaggingResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opDeleteJobTaggingResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
Endpoint: options.BaseEndpoint,
UseArnRegion: options.UseARNRegion,
    },

    }, "ResolveEndpoint", middleware.After)
}
