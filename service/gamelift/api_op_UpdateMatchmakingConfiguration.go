// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates settings for a FlexMatch matchmaking configuration. These changes
// affect all matches and game sessions that are created after the update. To
// update settings, specify the configuration name to be updated and provide the
// new settings. Learn more Design a FlexMatch matchmaker (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html)
func (c *Client) UpdateMatchmakingConfiguration(ctx context.Context, params *UpdateMatchmakingConfigurationInput, optFns ...func(*Options)) (*UpdateMatchmakingConfigurationOutput, error) {
	if params == nil {
		params = &UpdateMatchmakingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMatchmakingConfiguration", params, optFns, c.addOperationUpdateMatchmakingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

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
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMatchmakingConfiguration{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateMatchmakingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMatchmakingConfiguration(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateMatchmakingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateMatchmakingConfiguration",
	}
}
