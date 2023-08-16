// Code generated by smithy-go-codegen DO NOT EDIT.


package gamelift

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
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Updates settings for a FlexMatch matchmaking configuration. These changes
// affect all matches and game sessions that are created after the update. To
// update settings, specify the configuration name to be updated and provide the
// new settings. Learn more Design a FlexMatch matchmaker (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
func (c *Client) UpdateMatchmakingConfiguration(ctx context.Context, params *UpdateMatchmakingConfigurationInput, optFns ...func(*Options)) (*UpdateMatchmakingConfigurationOutput, error) {
	if params == nil { params = &UpdateMatchmakingConfigurationInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "UpdateMatchmakingConfiguration", params, optFns, c.addOperationUpdateMatchmakingConfigurationMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*UpdateMatchmakingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMatchmakingConfigurationInput struct {
	
	// A unique identifier for the matchmaking configuration to update. You can use
	// either the configuration name or ARN value.
	//
	// This member is required.
	Name *string
	
	// A flag that indicates whether a match that was created with this configuration
	// must be accepted by the matched players. To require acceptance, set to TRUE.
	// With this option enabled, matchmaking tickets use the status REQUIRES_ACCEPTANCE
	// to indicate when a completed potential match is waiting for player acceptance.
	AcceptanceRequired *bool
	
	// The length of time (in seconds) to wait for players to accept a proposed match,
	// if acceptance is required.
	AcceptanceTimeoutSeconds *int32
	
	// The number of player slots in a match to keep open for future players. For
	// example, if the configuration's rule set specifies a match for a single
	// 10-person team, and the additional player count is set to 2, 10 players will be
	// selected for the match and 2 more player slots will be open for future players.
	// This parameter is not used if FlexMatchMode is set to STANDALONE .
	AdditionalPlayerCount *int32
	
	// The method that is used to backfill game sessions created with this matchmaking
	// configuration. Specify MANUAL when your game manages backfill requests manually
	// or does not use the match backfill feature. Specify AUTOMATIC to have GameLift
	// create a match backfill request whenever a game session has one or more open
	// slots. Learn more about manual and automatic backfill in Backfill Existing
	// Games with FlexMatch (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-backfill.html)
	// . Automatic backfill is not available when FlexMatchMode is set to STANDALONE .
	BackfillMode types.BackfillMode
	
	// Information to add to all events related to the matchmaking configuration.
	CustomEventData *string
	
	// A description for the matchmaking configuration.
	Description *string
	
	// Indicates whether this matchmaking configuration is being used with Amazon
	// GameLift hosting or as a standalone matchmaking solution.
	//   - STANDALONE - FlexMatch forms matches and returns match information,
	//   including players and team assignments, in a MatchmakingSucceeded (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-events.html#match-events-matchmakingsucceeded)
	//   event.
	//   - WITH_QUEUE - FlexMatch forms matches and uses the specified Amazon GameLift
	//   queue to start a game session for the match.
	FlexMatchMode types.FlexMatchMode
	
	// A set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process with a request to start a
	// new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
	// ). This information is added to the new GameSession object that is created for
	// a successful match. This parameter is not used if FlexMatchMode is set to
	// STANDALONE .
	GameProperties []types.GameProperty
	
	// A set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process with a request to start a new game
	// session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)
	// ). This information is added to the game session that is created for a
	// successful match. This parameter is not used if FlexMatchMode is set to
	// STANDALONE .
	GameSessionData *string
	
	// The Amazon Resource Name ( ARN (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)
	// ) that is assigned to a Amazon GameLift game session queue resource and uniquely
	// identifies it. ARNs are unique across all Regions. Format is
	// arn:aws:gamelift:::gamesessionqueue/ . Queues can be located in any Region.
	// Queues are used to start new Amazon GameLift-hosted game sessions for matches
	// that are created with this matchmaking configuration. If FlexMatchMode is set
	// to STANDALONE , do not set this parameter.
	GameSessionQueueArns []string
	
	// An SNS topic ARN that is set up to receive matchmaking notifications. See
	// Setting up notifications for matchmaking (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html)
	// for more information.
	NotificationTarget *string
	
	// The maximum duration, in seconds, that a matchmaking ticket can remain in
	// process before timing out. Requests that fail due to timing out can be
	// resubmitted as needed.
	RequestTimeoutSeconds *int32
	
	// A unique identifier for the matchmaking rule set to use with this
	// configuration. You can use either the rule set name or ARN value. A matchmaking
	// configuration can only use rule sets that are defined in the same Region.
	RuleSetName *string
	
	noSmithyDocumentSerde
}

type UpdateMatchmakingConfigurationOutput struct {
	
	// The updated matchmaking configuration.
	Configuration *types.MatchmakingConfiguration
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMatchmakingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMatchmakingConfiguration{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMatchmakingConfiguration{}, middleware.After)
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
	if err = addUpdateMatchmakingConfigurationResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpUpdateMatchmakingConfigurationValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMatchmakingConfiguration(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMatchmakingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "gamelift",
	OperationName: "UpdateMatchmakingConfiguration",
	}
}

type opUpdateMatchmakingConfigurationResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opUpdateMatchmakingConfigurationResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateMatchmakingConfigurationResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "gamelift"
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
	        signingName = "gamelift"
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
	        v4aScheme.SigningName = aws.String("gamelift")
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

func addUpdateMatchmakingConfigurationResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opUpdateMatchmakingConfigurationResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
