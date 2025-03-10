/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2Age
type ToolsBidSuggestV2Age string

// List of tools_bid_suggest_v2_age
const (
	AGE_BELOW_18_ToolsBidSuggestV2Age      ToolsBidSuggestV2Age = "AGE_BELOW_18"
	AGE_BETWEEN_41_49_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_24_30_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_18_23_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_18_23"
	AGE_ABOVE_50_ToolsBidSuggestV2Age      ToolsBidSuggestV2Age = "AGE_ABOVE_50"
	AGE_BETWEEN_31_40_ToolsBidSuggestV2Age ToolsBidSuggestV2Age = "AGE_BETWEEN_31_40"
)

// Ptr returns reference to tools_bid_suggest_v2_age value
func (v ToolsBidSuggestV2Age) Ptr() *ToolsBidSuggestV2Age {
	return &v
}
