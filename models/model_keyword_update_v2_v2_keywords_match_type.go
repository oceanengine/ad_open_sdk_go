/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2KeywordsMatchType
type KeywordUpdateV2V2KeywordsMatchType string

// List of keyword_update_v2_v2_keywords_match_type
const (
	EXTENSIVE_KeywordUpdateV2V2KeywordsMatchType KeywordUpdateV2V2KeywordsMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV2V2KeywordsMatchType    KeywordUpdateV2V2KeywordsMatchType = "PHRASE"
	PRECISION_KeywordUpdateV2V2KeywordsMatchType KeywordUpdateV2V2KeywordsMatchType = "PRECISION"
)

// Ptr returns reference to keyword_update_v2_v2_keywords_match_type value
func (v KeywordUpdateV2V2KeywordsMatchType) Ptr() *KeywordUpdateV2V2KeywordsMatchType {
	return &v
}
