/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2KeywordsMatchType
type KeywordCreateV2V2KeywordsMatchType string

// List of keyword_create_v2_v2_keywords_match_type
const (
	PHRASE_KeywordCreateV2V2KeywordsMatchType    KeywordCreateV2V2KeywordsMatchType = "PHRASE"
	PRECISION_KeywordCreateV2V2KeywordsMatchType KeywordCreateV2V2KeywordsMatchType = "PRECISION"
	EXTENSIVE_KeywordCreateV2V2KeywordsMatchType KeywordCreateV2V2KeywordsMatchType = "EXTENSIVE"
)

// Ptr returns reference to keyword_create_v2_v2_keywords_match_type value
func (v KeywordCreateV2V2KeywordsMatchType) Ptr() *KeywordCreateV2V2KeywordsMatchType {
	return &v
}
