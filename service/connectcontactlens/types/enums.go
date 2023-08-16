// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type SentimentValue string

// Enum values for SentimentValue
const (
	SentimentValuePositive SentimentValue = "POSITIVE"
	SentimentValueNeutral SentimentValue = "NEUTRAL"
	SentimentValueNegative SentimentValue = "NEGATIVE"
)

// Values returns all known values for SentimentValue. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SentimentValue) Values() []SentimentValue {
	return []SentimentValue{
		"POSITIVE",
		"NEUTRAL",
		"NEGATIVE",
	}
}
