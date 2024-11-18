/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2AutoExtendTargets
type ToolsBidSuggestV2AutoExtendTargets string

// List of tools_bid_suggest_v2_auto_extend_targets
const (
	CUSTOM_AUDIENCE_ToolsBidSuggestV2AutoExtendTargets ToolsBidSuggestV2AutoExtendTargets = "CUSTOM_AUDIENCE"
	AGE_ToolsBidSuggestV2AutoExtendTargets             ToolsBidSuggestV2AutoExtendTargets = "AGE"
	GENDER_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "GENDER"
	REGION_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "REGION"
	INTEREST_TAG_ToolsBidSuggestV2AutoExtendTargets    ToolsBidSuggestV2AutoExtendTargets = "INTEREST_TAG"
	AD_TAG_ToolsBidSuggestV2AutoExtendTargets          ToolsBidSuggestV2AutoExtendTargets = "AD_TAG"
)

// Ptr returns reference to tools_bid_suggest_v2_auto_extend_targets value
func (v ToolsBidSuggestV2AutoExtendTargets) Ptr() *ToolsBidSuggestV2AutoExtendTargets {
	return &v
}
