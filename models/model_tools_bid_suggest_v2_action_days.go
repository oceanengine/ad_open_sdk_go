/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2ActionDays
type ToolsBidSuggestV2ActionDays int64

// List of tools_bid_suggest_v2_action_days
const (
	Enum_7_ToolsBidSuggestV2ActionDays   ToolsBidSuggestV2ActionDays = 7
	Enum_15_ToolsBidSuggestV2ActionDays  ToolsBidSuggestV2ActionDays = 15
	Enum_30_ToolsBidSuggestV2ActionDays  ToolsBidSuggestV2ActionDays = 30
	Enum_60_ToolsBidSuggestV2ActionDays  ToolsBidSuggestV2ActionDays = 60
	Enum_90_ToolsBidSuggestV2ActionDays  ToolsBidSuggestV2ActionDays = 90
	Enum_180_ToolsBidSuggestV2ActionDays ToolsBidSuggestV2ActionDays = 180
	Enum_365_ToolsBidSuggestV2ActionDays ToolsBidSuggestV2ActionDays = 365
)

// Ptr returns reference to tools_bid_suggest_v2_action_days value
func (v ToolsBidSuggestV2ActionDays) Ptr() *ToolsBidSuggestV2ActionDays {
	return &v
}
