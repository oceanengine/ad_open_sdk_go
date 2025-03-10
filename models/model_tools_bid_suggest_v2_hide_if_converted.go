/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsBidSuggestV2HideIfConverted
type ToolsBidSuggestV2HideIfConverted string

// List of tools_bid_suggest_v2_hide_if_converted
const (
	APP_ToolsBidSuggestV2HideIfConverted          ToolsBidSuggestV2HideIfConverted = "APP"
	NO_EXCLUDE_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "NO_EXCLUDE"
	ADVERTISER_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "ADVERTISER"
	AD_ToolsBidSuggestV2HideIfConverted           ToolsBidSuggestV2HideIfConverted = "AD"
	CUSTOMER_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CUSTOMER"
	GLOBAL_APP_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "GLOBAL_APP"
	ORGANIZATION_ToolsBidSuggestV2HideIfConverted ToolsBidSuggestV2HideIfConverted = "ORGANIZATION"
	CAMPAIGN_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CAMPAIGN"
)

// Ptr returns reference to tools_bid_suggest_v2_hide_if_converted value
func (v ToolsBidSuggestV2HideIfConverted) Ptr() *ToolsBidSuggestV2HideIfConverted {
	return &v
}
