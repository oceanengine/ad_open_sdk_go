/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2DataSuccessListMatchType
type KeywordUpdateV2V2DataSuccessListMatchType string

// List of keyword_update_v2_v2_data_success_list_match_type
const (
	EXTENSIVE_KeywordUpdateV2V2DataSuccessListMatchType KeywordUpdateV2V2DataSuccessListMatchType = "EXTENSIVE"
	PHRASE_KeywordUpdateV2V2DataSuccessListMatchType    KeywordUpdateV2V2DataSuccessListMatchType = "PHRASE"
	PRECISION_KeywordUpdateV2V2DataSuccessListMatchType KeywordUpdateV2V2DataSuccessListMatchType = "PRECISION"
)

// Ptr returns reference to keyword_update_v2_v2_data_success_list_match_type value
func (v KeywordUpdateV2V2DataSuccessListMatchType) Ptr() *KeywordUpdateV2V2DataSuccessListMatchType {
	return &v
}
