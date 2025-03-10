/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2ActionScene
type ToolsBidSuggestV2ActionScene string

// List of tools_bid_suggest_v2_action_scene
const (
	AD_ToolsBidSuggestV2ActionScene         ToolsBidSuggestV2ActionScene = "AD"
	E_COMMERCE_ToolsBidSuggestV2ActionScene ToolsBidSuggestV2ActionScene = "E-COMMERCE"
	NEWS_ToolsBidSuggestV2ActionScene       ToolsBidSuggestV2ActionScene = "NEWS"
	APP_ToolsBidSuggestV2ActionScene        ToolsBidSuggestV2ActionScene = "APP"
	SEARCH_ToolsBidSuggestV2ActionScene     ToolsBidSuggestV2ActionScene = "SEARCH"
)

// Ptr returns reference to tools_bid_suggest_v2_action_scene value
func (v ToolsBidSuggestV2ActionScene) Ptr() *ToolsBidSuggestV2ActionScene {
	return &v
}
