/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2AppBehaviorTarget
type ToolsBidSuggestV2AppBehaviorTarget string

// List of tools_bid_suggest_v2_app_behavior_target
const (
	NONE_ToolsBidSuggestV2AppBehaviorTarget     ToolsBidSuggestV2AppBehaviorTarget = "NONE"
	APP_ToolsBidSuggestV2AppBehaviorTarget      ToolsBidSuggestV2AppBehaviorTarget = "APP"
	CATEGORY_ToolsBidSuggestV2AppBehaviorTarget ToolsBidSuggestV2AppBehaviorTarget = "CATEGORY"
)

// Ptr returns reference to tools_bid_suggest_v2_app_behavior_target value
func (v ToolsBidSuggestV2AppBehaviorTarget) Ptr() *ToolsBidSuggestV2AppBehaviorTarget {
	return &v
}
