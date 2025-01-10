/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2Pricing
type ToolsBidSuggestV2Pricing string

// List of tools_bid_suggest_v2_pricing
const (
	PRICING_OCPC_ToolsBidSuggestV2Pricing     ToolsBidSuggestV2Pricing = "PRICING_OCPC"
	PRICING_CPM_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPM"
	PRICING_CPC_OCPM_ToolsBidSuggestV2Pricing ToolsBidSuggestV2Pricing = "PRICING_CPC_OCPM"
	PRICING_CPA_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPA"
	PRICING_OCPM_ToolsBidSuggestV2Pricing     ToolsBidSuggestV2Pricing = "PRICING_OCPM"
	PRICING_CPC_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPC"
	PRICING_CPV_ToolsBidSuggestV2Pricing      ToolsBidSuggestV2Pricing = "PRICING_CPV"
)

// Ptr returns reference to tools_bid_suggest_v2_pricing value
func (v ToolsBidSuggestV2Pricing) Ptr() *ToolsBidSuggestV2Pricing {
	return &v
}
