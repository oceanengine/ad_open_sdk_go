/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2DataErrorListMatchType
type KeywordUpdateV2V2DataErrorListMatchType string

// List of keyword_update_v2_v2_data_error_list_match_type
const (
	EXTENSIVE_KeywordUpdateV2V2DataErrorListMatchType KeywordUpdateV2V2DataErrorListMatchType = "EXTENSIVE"
	PRECISION_KeywordUpdateV2V2DataErrorListMatchType KeywordUpdateV2V2DataErrorListMatchType = "PRECISION"
	PHRASE_KeywordUpdateV2V2DataErrorListMatchType    KeywordUpdateV2V2DataErrorListMatchType = "PHRASE"
)

// Ptr returns reference to keyword_update_v2_v2_data_error_list_match_type value
func (v KeywordUpdateV2V2DataErrorListMatchType) Ptr() *KeywordUpdateV2V2DataErrorListMatchType {
	return &v
}
