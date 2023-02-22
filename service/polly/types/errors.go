// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// This engine is not compatible with the voice that you have designated. Choose a
// new voice that is compatible with the engine or change the engine and restart
// the operation.
type EngineNotSupportedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EngineNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EngineNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EngineNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EngineNotSupportedException"
	}
	return *e.ErrorCodeOverride
}
func (e *EngineNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Polly can't find the specified lexicon. Verify that the lexicon's name is
// spelled correctly, and then try again.
type InvalidLexiconException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidLexiconException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidLexiconException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidLexiconException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidLexiconException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidLexiconException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The NextToken is invalid. Verify that it's spelled correctly, and then try
// again.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided Amazon S3 bucket name is invalid. Please check your input with S3
// bucket naming requirements and try again.
type InvalidS3BucketException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidS3BucketException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3BucketException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3BucketException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidS3BucketException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidS3BucketException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided Amazon S3 key prefix is invalid. Please provide a valid S3 object
// key name.
type InvalidS3KeyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidS3KeyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3KeyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3KeyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidS3KeyException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidS3KeyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified sample rate is not valid.
type InvalidSampleRateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSampleRateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSampleRateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSampleRateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSampleRateException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSampleRateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided SNS topic ARN is invalid. Please provide a valid SNS topic ARN and
// try again.
type InvalidSnsTopicArnException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSnsTopicArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSnsTopicArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSnsTopicArnException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSnsTopicArnException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSnsTopicArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The SSML you provided is invalid. Verify the SSML syntax, spelling of tags and
// values, and then try again.
type InvalidSsmlException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidSsmlException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSsmlException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSsmlException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidSsmlException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidSsmlException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The provided Task ID is not valid. Please provide a valid Task ID and try again.
type InvalidTaskIdException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidTaskIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTaskIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTaskIdException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidTaskIdException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidTaskIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The language specified is not currently supported by Amazon Polly in this
// capacity.
type LanguageNotSupportedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LanguageNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LanguageNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LanguageNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LanguageNotSupportedException"
	}
	return *e.ErrorCodeOverride
}
func (e *LanguageNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Polly can't find the specified lexicon. This could be caused by a lexicon
// that is missing, its name is misspelled or specifying a lexicon that is in a
// different region. Verify that the lexicon exists, is in the region (see
// ListLexicons) and that you spelled its name is spelled correctly. Then try
// again.
type LexiconNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LexiconNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LexiconNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LexiconNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LexiconNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *LexiconNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum size of the specified lexicon would be exceeded by this operation.
type LexiconSizeExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LexiconSizeExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LexiconSizeExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LexiconSizeExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LexiconSizeExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LexiconSizeExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Speech marks are not supported for the OutputFormat selected. Speech marks are
// only available for content in json format.
type MarksNotSupportedForFormatException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MarksNotSupportedForFormatException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MarksNotSupportedForFormatException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MarksNotSupportedForFormatException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MarksNotSupportedForFormatException"
	}
	return *e.ErrorCodeOverride
}
func (e *MarksNotSupportedForFormatException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The maximum size of the lexeme would be exceeded by this operation.
type MaxLexemeLengthExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MaxLexemeLengthExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MaxLexemeLengthExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MaxLexemeLengthExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MaxLexemeLengthExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *MaxLexemeLengthExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The maximum number of lexicons would be exceeded by this operation.
type MaxLexiconsNumberExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *MaxLexiconsNumberExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MaxLexiconsNumberExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MaxLexiconsNumberExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "MaxLexiconsNumberExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *MaxLexiconsNumberExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// An unknown condition has caused a service failure.
type ServiceFailureException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ServiceFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceFailureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ServiceFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *ServiceFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// SSML speech marks are not supported for plain text-type input.
type SsmlMarksNotSupportedForTextTypeException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SsmlMarksNotSupportedForTextTypeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SsmlMarksNotSupportedForTextTypeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SsmlMarksNotSupportedForTextTypeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SsmlMarksNotSupportedForTextTypeException"
	}
	return *e.ErrorCodeOverride
}
func (e *SsmlMarksNotSupportedForTextTypeException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The Speech Synthesis task with requested Task ID cannot be found.
type SynthesisTaskNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *SynthesisTaskNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SynthesisTaskNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SynthesisTaskNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SynthesisTaskNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *SynthesisTaskNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value of the "Text" parameter is longer than the accepted limits. For the
// SynthesizeSpeech API, the limit for input text is a maximum of 6000 characters
// total, of which no more than 3000 can be billed characters. For the
// StartSpeechSynthesisTask API, the maximum is 200,000 characters, of which no
// more than 100,000 can be billed characters. SSML tags are not counted as billed
// characters.
type TextLengthExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TextLengthExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TextLengthExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TextLengthExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TextLengthExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *TextLengthExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The alphabet specified by the lexicon is not a supported alphabet. Valid values
// are x-sampa and ipa.
type UnsupportedPlsAlphabetException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedPlsAlphabetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedPlsAlphabetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedPlsAlphabetException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedPlsAlphabetException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedPlsAlphabetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The language specified in the lexicon is unsupported. For a list of supported
// languages, see Lexicon Attributes
// (https://docs.aws.amazon.com/polly/latest/dg/API_LexiconAttributes.html).
type UnsupportedPlsLanguageException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *UnsupportedPlsLanguageException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedPlsLanguageException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedPlsLanguageException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedPlsLanguageException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedPlsLanguageException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
