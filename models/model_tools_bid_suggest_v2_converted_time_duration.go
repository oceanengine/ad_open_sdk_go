/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2ConvertedTimeDuration
type ToolsBidSuggestV2ConvertedTimeDuration string

// List of tools_bid_suggest_v2_converted_time_duration
const (
	NONE_ToolsBidSuggestV2ConvertedTimeDuration         ToolsBidSuggestV2ConvertedTimeDuration = "NONE"
	ONE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "ONE_MONTH"
	TODAY_ToolsBidSuggestV2ConvertedTimeDuration        ToolsBidSuggestV2ConvertedTimeDuration = "TODAY"
	SIX_MONTH_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "SIX_MONTH"
	THREE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration  ToolsBidSuggestV2ConvertedTimeDuration = "THREE_MONTH"
	SEVEN_DAY_ToolsBidSuggestV2ConvertedTimeDuration    ToolsBidSuggestV2ConvertedTimeDuration = "SEVEN_DAY"
	TWELVE_MONTH_ToolsBidSuggestV2ConvertedTimeDuration ToolsBidSuggestV2ConvertedTimeDuration = "TWELVE_MONTH"
)

// Ptr returns reference to tools_bid_suggest_v2_converted_time_duration value
func (v ToolsBidSuggestV2ConvertedTimeDuration) Ptr() *ToolsBidSuggestV2ConvertedTimeDuration {
	return &v
}
