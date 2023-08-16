// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the Savings Plans recommendations for your account. First use
// StartSavingsPlansPurchaseRecommendationGeneration to generate a new set of
// recommendations, and then use GetSavingsPlansPurchaseRecommendation to retrieve
// them.
func (c *Client) GetSavingsPlansPurchaseRecommendation(ctx context.Context, params *GetSavingsPlansPurchaseRecommendationInput, optFns ...func(*Options)) (*GetSavingsPlansPurchaseRecommendationOutput, error) {
	if params == nil {
		params = &GetSavingsPlansPurchaseRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansPurchaseRecommendation", params, optFns, c.addOperationGetSavingsPlansPurchaseRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansPurchaseRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansPurchaseRecommendationInput struct {

	// The lookback period that's used to generate the recommendation.
	//
	// This member is required.
	LookbackPeriodInDays types.LookbackPeriodInDays

	// The payment option that's used to generate these recommendations.
	//
	// This member is required.
	PaymentOption types.PaymentOption

	// The Savings Plans recommendation type that's requested.
	//
	// This member is required.
	SavingsPlansType types.SupportedSavingsPlansType

	// The savings plan recommendation term that's used to generate these
	// recommendations.
	//
	// This member is required.
	TermInYears types.TermInYears

	// The account scope that you want your recommendations for. Amazon Web Services
	// calculates recommendations including the management account and member accounts
	// if the value is set to PAYER . If the value is LINKED , recommendations are
	// calculated for individual member accounts only.
	AccountScope types.AccountScope

	// You can filter your recommendations by Account ID with the LINKED_ACCOUNT
	// dimension. To filter your recommendations by Account ID, specify Key as
	// LINKED_ACCOUNT and Value as the comma-separated Acount ID(s) that you want to
	// see Savings Plans purchase recommendations for. For
	// GetSavingsPlansPurchaseRecommendation, the Filter doesn't include CostCategories
	// or Tags . It only includes Dimensions . With Dimensions , Key must be
	// LINKED_ACCOUNT and Value can be a single Account ID or multiple comma-separated
	// Account IDs that you want to see Savings Plans Purchase Recommendations for. AND
	// and OR operators are not supported.
	Filter *types.Expression

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize int32

	noSmithyDocumentSerde
}

type GetSavingsPlansPurchaseRecommendationOutput struct {

	// Information that regards this specific recommendation set.
	Metadata *types.SavingsPlansPurchaseRecommendationMetadata

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// Contains your request parameters, Savings Plan Recommendations Summary, and
	// Details.
	SavingsPlansPurchaseRecommendation *types.SavingsPlansPurchaseRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSavingsPlansPurchaseRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansPurchaseRecommendation{}, middleware.After)
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
	if err = addGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetSavingsPlansPurchaseRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSavingsPlansPurchaseRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetSavingsPlansPurchaseRecommendation",
	}
}

type opGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
			signingName := "ce"
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
				signingName = "ce"
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
				v4aScheme.SigningName = aws.String("ce")
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

func addGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opGetSavingsPlansPurchaseRecommendationResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
