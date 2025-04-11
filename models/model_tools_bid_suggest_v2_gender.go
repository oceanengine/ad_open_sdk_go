/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2Gender
type ToolsBidSuggestV2Gender string

// List of tools_bid_suggest_v2_gender
const (
	GENDER_MALE_ToolsBidSuggestV2Gender      ToolsBidSuggestV2Gender = "GENDER_MALE"
	NONE_ToolsBidSuggestV2Gender             ToolsBidSuggestV2Gender = "NONE"
	GENDER_UNLIMITED_ToolsBidSuggestV2Gender ToolsBidSuggestV2Gender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ToolsBidSuggestV2Gender    ToolsBidSuggestV2Gender = "GENDER_FEMALE"
)

// Ptr returns reference to tools_bid_suggest_v2_gender value
func (v ToolsBidSuggestV2Gender) Ptr() *ToolsBidSuggestV2Gender {
	return &v
}
