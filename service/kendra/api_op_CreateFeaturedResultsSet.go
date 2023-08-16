// Code generated by smithy-go-codegen DO NOT EDIT.


package kendra

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
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Creates a set of featured results to display at the top of the search results
// page. Featured results are placed above all other results for certain queries.
// You map specific queries to specific documents for featuring in the results. If
// a query contains an exact match, then one or more specific documents are
// featured in the search results. You can create up to 50 sets of featured results
// per index. You can request to increase this limit by contacting Support (http://aws.amazon.com/contact-us/)
// .
func (c *Client) CreateFeaturedResultsSet(ctx context.Context, params *CreateFeaturedResultsSetInput, optFns ...func(*Options)) (*CreateFeaturedResultsSetOutput, error) {
	if params == nil { params = &CreateFeaturedResultsSetInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "CreateFeaturedResultsSet", params, optFns, c.addOperationCreateFeaturedResultsSetMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*CreateFeaturedResultsSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFeaturedResultsSetInput struct {
	
	// A name for the set of featured results.
	//
	// This member is required.
	FeaturedResultsSetName *string
	
	// The identifier of the index that you want to use for featuring results.
	//
	// This member is required.
	IndexId *string
	
	// A token that you provide to identify the request to create a set of featured
	// results. Multiple calls to the CreateFeaturedResultsSet API with the same
	// client token will create only one featured results set.
	ClientToken *string
	
	// A description for the set of featured results.
	Description *string
	
	// A list of document IDs for the documents you want to feature at the top of the
	// search results page. For more information on the list of documents, see
	// FeaturedResultsSet (https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html)
	// .
	FeaturedDocuments []types.FeaturedDocument
	
	// A list of queries for featuring results. For more information on the list of
	// queries, see FeaturedResultsSet (https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html)
	// .
	QueryTexts []string
	
	// The current status of the set of featured results. When the value is ACTIVE ,
	// featured results are ready for use. You can still configure your settings before
	// setting the status to ACTIVE . You can set the status to ACTIVE or INACTIVE
	// using the UpdateFeaturedResultsSet (https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateFeaturedResultsSet.html)
	// API. The queries you specify for featured results must be unique per featured
	// results set for each index, whether the status is ACTIVE or INACTIVE .
	Status types.FeaturedResultsSetStatus
	
	// A list of key-value pairs that identify or categorize the featured results set.
	// You can also use tags to help control access to the featured results set. Tag
	// keys and values can consist of Unicode letters, digits, white space, and any of
	// the following symbols:_ . : / = + - @.
	Tags []types.Tag
	
	noSmithyDocumentSerde
}

type CreateFeaturedResultsSetOutput struct {
	
	// Information on the set of featured results. This includes the identifier of the
	// featured results set, whether the featured results set is active or inactive,
	// when the featured results set was created, and more.
	FeaturedResultsSet *types.FeaturedResultsSet
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFeaturedResultsSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFeaturedResultsSet{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFeaturedResultsSet{}, middleware.After)
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
	if err = addCreateFeaturedResultsSetResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpCreateFeaturedResultsSetValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFeaturedResultsSet(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFeaturedResultsSet(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "kendra",
	OperationName: "CreateFeaturedResultsSet",
	}
}

type opCreateFeaturedResultsSetResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opCreateFeaturedResultsSetResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opCreateFeaturedResultsSetResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "kendra"
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
	        signingName = "kendra"
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
	        v4aScheme.SigningName = aws.String("kendra")
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

func addCreateFeaturedResultsSetResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opCreateFeaturedResultsSetResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
