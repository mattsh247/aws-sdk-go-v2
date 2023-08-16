// Code generated by smithy-go-codegen DO NOT EDIT.


package types

type AlgorithmicStemming string

// Enum values for AlgorithmicStemming
const (
	AlgorithmicStemmingNone AlgorithmicStemming = "none"
	AlgorithmicStemmingMinimal AlgorithmicStemming = "minimal"
	AlgorithmicStemmingLight AlgorithmicStemming = "light"
	AlgorithmicStemmingFull AlgorithmicStemming = "full"
)

// Values returns all known values for AlgorithmicStemming. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AlgorithmicStemming) Values() []AlgorithmicStemming {
	return []AlgorithmicStemming{
		"none",
		"minimal",
		"light",
		"full",
	}
}

type AnalysisSchemeLanguage string

// Enum values for AnalysisSchemeLanguage
const (
	AnalysisSchemeLanguageAr AnalysisSchemeLanguage = "ar"
	AnalysisSchemeLanguageBg AnalysisSchemeLanguage = "bg"
	AnalysisSchemeLanguageCa AnalysisSchemeLanguage = "ca"
	AnalysisSchemeLanguageCs AnalysisSchemeLanguage = "cs"
	AnalysisSchemeLanguageDa AnalysisSchemeLanguage = "da"
	AnalysisSchemeLanguageDe AnalysisSchemeLanguage = "de"
	AnalysisSchemeLanguageEl AnalysisSchemeLanguage = "el"
	AnalysisSchemeLanguageEn AnalysisSchemeLanguage = "en"
	AnalysisSchemeLanguageEs AnalysisSchemeLanguage = "es"
	AnalysisSchemeLanguageEu AnalysisSchemeLanguage = "eu"
	AnalysisSchemeLanguageFa AnalysisSchemeLanguage = "fa"
	AnalysisSchemeLanguageFi AnalysisSchemeLanguage = "fi"
	AnalysisSchemeLanguageFr AnalysisSchemeLanguage = "fr"
	AnalysisSchemeLanguageGa AnalysisSchemeLanguage = "ga"
	AnalysisSchemeLanguageGl AnalysisSchemeLanguage = "gl"
	AnalysisSchemeLanguageHe AnalysisSchemeLanguage = "he"
	AnalysisSchemeLanguageHi AnalysisSchemeLanguage = "hi"
	AnalysisSchemeLanguageHu AnalysisSchemeLanguage = "hu"
	AnalysisSchemeLanguageHy AnalysisSchemeLanguage = "hy"
	AnalysisSchemeLanguageId AnalysisSchemeLanguage = "id"
	AnalysisSchemeLanguageIt AnalysisSchemeLanguage = "it"
	AnalysisSchemeLanguageJa AnalysisSchemeLanguage = "ja"
	AnalysisSchemeLanguageKo AnalysisSchemeLanguage = "ko"
	AnalysisSchemeLanguageLv AnalysisSchemeLanguage = "lv"
	AnalysisSchemeLanguageMul AnalysisSchemeLanguage = "mul"
	AnalysisSchemeLanguageNl AnalysisSchemeLanguage = "nl"
	AnalysisSchemeLanguageNo AnalysisSchemeLanguage = "no"
	AnalysisSchemeLanguagePt AnalysisSchemeLanguage = "pt"
	AnalysisSchemeLanguageRo AnalysisSchemeLanguage = "ro"
	AnalysisSchemeLanguageRu AnalysisSchemeLanguage = "ru"
	AnalysisSchemeLanguageSv AnalysisSchemeLanguage = "sv"
	AnalysisSchemeLanguageTh AnalysisSchemeLanguage = "th"
	AnalysisSchemeLanguageTr AnalysisSchemeLanguage = "tr"
	AnalysisSchemeLanguageZhHans AnalysisSchemeLanguage = "zh-Hans"
	AnalysisSchemeLanguageZhHant AnalysisSchemeLanguage = "zh-Hant"
)

// Values returns all known values for AnalysisSchemeLanguage. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalysisSchemeLanguage) Values() []AnalysisSchemeLanguage {
	return []AnalysisSchemeLanguage{
		"ar",
		"bg",
		"ca",
		"cs",
		"da",
		"de",
		"el",
		"en",
		"es",
		"eu",
		"fa",
		"fi",
		"fr",
		"ga",
		"gl",
		"he",
		"hi",
		"hu",
		"hy",
		"id",
		"it",
		"ja",
		"ko",
		"lv",
		"mul",
		"nl",
		"no",
		"pt",
		"ro",
		"ru",
		"sv",
		"th",
		"tr",
		"zh-Hans",
		"zh-Hant",
	}
}

type IndexFieldType string

// Enum values for IndexFieldType
const (
	IndexFieldTypeInt IndexFieldType = "int"
	IndexFieldTypeDouble IndexFieldType = "double"
	IndexFieldTypeLiteral IndexFieldType = "literal"
	IndexFieldTypeText IndexFieldType = "text"
	IndexFieldTypeDate IndexFieldType = "date"
	IndexFieldTypeLatlon IndexFieldType = "latlon"
	IndexFieldTypeIntArray IndexFieldType = "int-array"
	IndexFieldTypeDoubleArray IndexFieldType = "double-array"
	IndexFieldTypeLiteralArray IndexFieldType = "literal-array"
	IndexFieldTypeTextArray IndexFieldType = "text-array"
	IndexFieldTypeDateArray IndexFieldType = "date-array"
)

// Values returns all known values for IndexFieldType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IndexFieldType) Values() []IndexFieldType {
	return []IndexFieldType{
		"int",
		"double",
		"literal",
		"text",
		"date",
		"latlon",
		"int-array",
		"double-array",
		"literal-array",
		"text-array",
		"date-array",
	}
}

type OptionState string

// Enum values for OptionState
const (
	OptionStateRequiresIndexDocuments OptionState = "RequiresIndexDocuments"
	OptionStateProcessing OptionState = "Processing"
	OptionStateActive OptionState = "Active"
	OptionStateFailedToValidate OptionState = "FailedToValidate"
)

// Values returns all known values for OptionState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OptionState) Values() []OptionState {
	return []OptionState{
		"RequiresIndexDocuments",
		"Processing",
		"Active",
		"FailedToValidate",
	}
}

type PartitionInstanceType string

// Enum values for PartitionInstanceType
const (
	PartitionInstanceTypeSearchM1Small PartitionInstanceType = "search.m1.small"
	PartitionInstanceTypeSearchM1Large PartitionInstanceType = "search.m1.large"
	PartitionInstanceTypeSearchM2Xlarge PartitionInstanceType = "search.m2.xlarge"
	PartitionInstanceTypeSearchM22xlarge PartitionInstanceType = "search.m2.2xlarge"
	PartitionInstanceTypeSearchM3Medium PartitionInstanceType = "search.m3.medium"
	PartitionInstanceTypeSearchM3Large PartitionInstanceType = "search.m3.large"
	PartitionInstanceTypeSearchM3Xlarge PartitionInstanceType = "search.m3.xlarge"
	PartitionInstanceTypeSearchM32xlarge PartitionInstanceType = "search.m3.2xlarge"
	PartitionInstanceTypeSearchSmall PartitionInstanceType = "search.small"
	PartitionInstanceTypeSearchMedium PartitionInstanceType = "search.medium"
	PartitionInstanceTypeSearchLarge PartitionInstanceType = "search.large"
	PartitionInstanceTypeSearchXlarge PartitionInstanceType = "search.xlarge"
	PartitionInstanceTypeSearch2xlarge PartitionInstanceType = "search.2xlarge"
	PartitionInstanceTypeSearchPreviousgenerationSmall PartitionInstanceType = "search.previousgeneration.small"
	PartitionInstanceTypeSearchPreviousgenerationLarge PartitionInstanceType = "search.previousgeneration.large"
	PartitionInstanceTypeSearchPreviousgenerationXlarge PartitionInstanceType = "search.previousgeneration.xlarge"
	PartitionInstanceTypeSearchPreviousgeneration2xlarge PartitionInstanceType = "search.previousgeneration.2xlarge"
)

// Values returns all known values for PartitionInstanceType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PartitionInstanceType) Values() []PartitionInstanceType {
	return []PartitionInstanceType{
		"search.m1.small",
		"search.m1.large",
		"search.m2.xlarge",
		"search.m2.2xlarge",
		"search.m3.medium",
		"search.m3.large",
		"search.m3.xlarge",
		"search.m3.2xlarge",
		"search.small",
		"search.medium",
		"search.large",
		"search.xlarge",
		"search.2xlarge",
		"search.previousgeneration.small",
		"search.previousgeneration.large",
		"search.previousgeneration.xlarge",
		"search.previousgeneration.2xlarge",
	}
}

type SuggesterFuzzyMatching string

// Enum values for SuggesterFuzzyMatching
const (
	SuggesterFuzzyMatchingNone SuggesterFuzzyMatching = "none"
	SuggesterFuzzyMatchingLow SuggesterFuzzyMatching = "low"
	SuggesterFuzzyMatchingHigh SuggesterFuzzyMatching = "high"
)

// Values returns all known values for SuggesterFuzzyMatching. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuggesterFuzzyMatching) Values() []SuggesterFuzzyMatching {
	return []SuggesterFuzzyMatching{
		"none",
		"low",
		"high",
	}
}

type TLSSecurityPolicy string

// Enum values for TLSSecurityPolicy
const (
	TLSSecurityPolicyPolicyMinTls10201907 TLSSecurityPolicy = "Policy-Min-TLS-1-0-2019-07"
	TLSSecurityPolicyPolicyMinTls12201907 TLSSecurityPolicy = "Policy-Min-TLS-1-2-2019-07"
)

// Values returns all known values for TLSSecurityPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TLSSecurityPolicy) Values() []TLSSecurityPolicy {
	return []TLSSecurityPolicy{
		"Policy-Min-TLS-1-0-2019-07",
		"Policy-Min-TLS-1-2-2019-07",
	}
}
