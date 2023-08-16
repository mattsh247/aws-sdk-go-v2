// Code generated by smithy-go-codegen DO NOT EDIT.


package lexruntimeservice

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
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

// Sends user input to Amazon Lex. Client applications can use this API to send
// requests to Amazon Lex at runtime. Amazon Lex then interprets the user input
// using the machine learning model it built for the bot. In response, Amazon Lex
// returns the next message to convey to the user an optional responseCard to
// display. Consider the following example messages:
//   - For a user input "I would like a pizza", Amazon Lex might return a response
//   with a message eliciting slot data (for example, PizzaSize): "What size pizza
//   would you like?"
//   - After the user provides all of the pizza order information, Amazon Lex
//   might return a response with a message to obtain user confirmation "Proceed with
//   the pizza order?".
//   - After the user replies to a confirmation prompt with a "yes", Amazon Lex
//   might return a conclusion statement: "Thank you, your cheese pizza has been
//   ordered.".
// Not all Amazon Lex messages require a user response. For example, a conclusion
// statement does not require a response. Some messages require only a "yes" or
// "no" user response. In addition to the message , Amazon Lex provides additional
// context about the message in the response that you might use to enhance client
// behavior, for example, to display the appropriate client user interface. These
// are the slotToElicit , dialogState , intentName , and slots fields in the
// response. Consider the following examples:
//   - If the message is to elicit slot data, Amazon Lex returns the following
//   context information:
//   - dialogState set to ElicitSlot
//   - intentName set to the intent name in the current context
//   - slotToElicit set to the slot name for which the message is eliciting
//   information
//   - slots set to a map of slots, configured for the intent, with currently known
//   values
//   - If the message is a confirmation prompt, the dialogState is set to
//   ConfirmIntent and SlotToElicit is set to null.
//   - If the message is a clarification prompt (configured for the intent) that
//   indicates that user intent is not understood, the dialogState is set to
//   ElicitIntent and slotToElicit is set to null.
// In addition, Amazon Lex also returns your application-specific sessionAttributes
// . For more information, see Managing Conversation Context (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html)
// .
func (c *Client) PostText(ctx context.Context, params *PostTextInput, optFns ...func(*Options)) (*PostTextOutput, error) {
	if params == nil { params = &PostTextInput{} }
	
	result, metadata, err := c.invokeOperation(ctx, "PostText", params, optFns, c.addOperationPostTextMiddlewares)
	if err != nil { return nil, err }
	
	out := result.(*PostTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PostTextInput struct {
	
	// The alias of the Amazon Lex bot.
	//
	// This member is required.
	BotAlias *string
	
	// The name of the Amazon Lex bot.
	//
	// This member is required.
	BotName *string
	
	// The text that the user entered (Amazon Lex interprets this text).
	//
	// This member is required.
	InputText *string
	
	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot. At runtime, each request must contain the
	// userID field. To decide the user ID to use for your application, consider the
	// following factors.
	//   - The userID field must not contain any personally identifiable information of
	//   the user, for example, name, personal identification numbers, or other end user
	//   personal information.
	//   - If you want a user to start a conversation on one device and continue on
	//   another device, use a user-specific identifier.
	//   - If you want the same user to be able to have two independent conversations
	//   on two different devices, choose a device-specific identifier.
	//   - A user can't have two independent conversations with two different versions
	//   of the same bot. For example, a user can't have a conversation with the PROD and
	//   BETA versions of the same bot. If you anticipate that a user will need to have
	//   conversation with two different versions, for example, while testing, include
	//   the bot alias in the user ID to separate the two conversations.
	//
	// This member is required.
	UserId *string
	
	// A list of contexts active for the request. A context can be activated when a
	// previous intent is fulfilled, or by including the context in the request, If you
	// don't specify a list of contexts, Amazon Lex will use the current list of
	// contexts for the session. If you specify an empty list, all contexts for the
	// session are cleared.
	ActiveContexts []types.ActiveContext
	
	// Request-specific information passed between Amazon Lex and a client
	// application. The namespace x-amz-lex: is reserved for special attributes. Don't
	// create any request attributes with the prefix x-amz-lex: . For more information,
	// see Setting Request Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs)
	// .
	RequestAttributes map[string]string
	
	// Application-specific information passed between Amazon Lex and a client
	// application. For more information, see Setting Session Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs)
	// .
	SessionAttributes map[string]string
	
	noSmithyDocumentSerde
}

type PostTextOutput struct {
	
	// A list of active contexts for the session. A context can be set when an intent
	// is fulfilled or by calling the PostContent , PostText , or PutSession
	// operation. You can use a context to control the intents that can follow up an
	// intent, or to modify the operation of your application.
	ActiveContexts []types.ActiveContext
	
	// One to four alternative intents that may be applicable to the user's intent.
	// Each alternative includes a score that indicates how confident Amazon Lex is
	// that the intent matches the user's intent. The intents are sorted by the
	// confidence score.
	AlternativeIntents []types.PredictedIntent
	
	// The version of the bot that responded to the conversation. You can use this
	// information to help determine if one version of a bot is performing better than
	// another version.
	BotVersion *string
	
	// Identifies the current state of the user interaction. Amazon Lex returns one of
	// the following values as dialogState . The client can optionally use this
	// information to customize the user interface.
	//   - ElicitIntent - Amazon Lex wants to elicit user intent. For example, a user
	//   might utter an intent ("I want to order a pizza"). If Amazon Lex cannot infer
	//   the user intent from this utterance, it will return this dialogState.
	//   - ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response. For
	//   example, Amazon Lex wants user confirmation before fulfilling an intent. Instead
	//   of a simple "yes" or "no," a user might respond with additional information. For
	//   example, "yes, but make it thick crust pizza" or "no, I want to order a drink".
	//   Amazon Lex can process such additional information (in these examples, update
	//   the crust type slot value, or change intent from OrderPizza to OrderDrink).
	//   - ElicitSlot - Amazon Lex is expecting a slot value for the current intent.
	//   For example, suppose that in the response Amazon Lex sends this message: "What
	//   size pizza would you like?". A user might reply with the slot value (e.g.,
	//   "medium"). The user might also provide additional information in the response
	//   (e.g., "medium thick crust pizza"). Amazon Lex can process such additional
	//   information appropriately.
	//   - Fulfilled - Conveys that the Lambda function configured for the intent has
	//   successfully fulfilled the intent.
	//   - ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	//   - Failed - Conveys that the conversation with the user failed. This can happen
	//   for various reasons including that the user did not provide an appropriate
	//   response to prompts from the service (you can configure how many times Amazon
	//   Lex can prompt a user for specific information), or the Lambda function failed
	//   to fulfill the intent.
	DialogState types.DialogState
	
	// The current user intent that Amazon Lex is aware of.
	IntentName *string
	
	// The message to convey to the user. The message can come from the bot's
	// configuration or from a Lambda function. If the intent is not configured with a
	// Lambda function, or if the Lambda function returned Delegate as the
	// dialogAction.type its response, Amazon Lex decides on the next course of action
	// and selects an appropriate message from the bot's configuration based on the
	// current interaction context. For example, if Amazon Lex isn't able to understand
	// user input, it uses a clarification prompt message. When you create an intent
	// you can assign messages to groups. When messages are assigned to groups Amazon
	// Lex returns one message from each group in the response. The message field is an
	// escaped JSON string containing the messages. For more information about the
	// structure of the JSON string returned, see msg-prompts-formats . If the Lambda
	// function returns a message, Amazon Lex passes it to the client in its response.
	Message *string
	
	// The format of the response message. One of the following values:
	//   - PlainText - The message contains plain UTF-8 text.
	//   - CustomPayload - The message is a custom format defined by the Lambda
	//   function.
	//   - SSML - The message contains text formatted for voice output.
	//   - Composite - The message contains an escaped JSON object containing one or
	//   more messages from the groups that messages were assigned to when the intent was
	//   created.
	MessageFormat types.MessageFormatType
	
	// Provides a score that indicates how confident Amazon Lex is that the returned
	// intent is the one that matches the user's intent. The score is between 0.0 and
	// 1.0. For more information, see Confidence Scores (https://docs.aws.amazon.com/lex/latest/dg/confidence-scores.html)
	// . The score is a relative score, not an absolute score. The score may change
	// based on improvements to Amazon Lex.
	NluIntentConfidence *types.IntentConfidence
	
	// Represents the options that the user has to respond to the current prompt.
	// Response Card can come from the bot configuration (in the Amazon Lex console,
	// choose the settings button next to a slot) or from a code hook (Lambda
	// function).
	ResponseCard *types.ResponseCard
	
	// The sentiment expressed in and utterance. When the bot is configured to send
	// utterances to Amazon Comprehend for sentiment analysis, this field contains the
	// result of the analysis.
	SentimentResponse *types.SentimentResponse
	
	// A map of key-value pairs representing the session-specific context information.
	SessionAttributes map[string]string
	
	// A unique identifier for the session.
	SessionId *string
	
	// If the dialogState value is ElicitSlot , returns the name of the slot for which
	// Amazon Lex is eliciting a value.
	SlotToElicit *string
	
	// The intent slots that Amazon Lex detected from the user input in the
	// conversation. Amazon Lex creates a resolution list containing likely values for
	// a slot. The value that it returns is determined by the valueSelectionStrategy
	// selected when the slot type was created or updated. If valueSelectionStrategy
	// is set to ORIGINAL_VALUE , the value provided by the user is returned, if the
	// user value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list or, if
	// there is no resolution list, null. If you don't specify a valueSelectionStrategy
	// , the default is ORIGINAL_VALUE .
	Slots map[string]string
	
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
	
	noSmithyDocumentSerde
}

func (c *Client) addOperationPostTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPostText{}, middleware.After)
	if err != nil { return err }
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPostText{}, middleware.After)
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
	if err = addPostTextResolveEndpointMiddleware(stack, options); err != nil {
	return err
	}
	if err = addOpPostTextValidationMiddleware(stack); err != nil {
	return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPostText(options.Region, ), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPostText(region string) *awsmiddleware.RegisterServiceMetadata {
	 return &awsmiddleware.RegisterServiceMetadata{
	Region: region,
	ServiceID: ServiceID,
	SigningName: "lex",
	OperationName: "PostText",
	}
}

type opPostTextResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver builtInParameterResolver
}

func (*opPostTextResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opPostTextResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
	            signingName := "lex"
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
	        signingName = "lex"
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
	        v4aScheme.SigningName = aws.String("lex")
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

func addPostTextResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
    return stack.Serialize.Insert(&opPostTextResolveEndpointMiddleware{
        EndpointResolver: options.EndpointResolverV2,
        BuiltInResolver: &builtInResolver{
        Region: options.Region,
UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
UseFIPS: options.EndpointOptions.UseFIPSEndpoint,
Endpoint: options.BaseEndpoint,
    },

    }, "ResolveEndpoint", middleware.After)
}
