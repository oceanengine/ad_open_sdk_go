/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsNoBidSuggestBidV2FilteringPricing
type ToolsNoBidSuggestBidV2FilteringPricing string

// List of tools_no_bid_suggest_bid_v2_filtering_pricing
const (
	PRICING_OCPM_ToolsNoBidSuggestBidV2FilteringPricing     ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_OCPM"
	PRICING_CPM_ToolsNoBidSuggestBidV2FilteringPricing      ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_CPM"
	PRICING_OCPC_ToolsNoBidSuggestBidV2FilteringPricing     ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_OCPC"
	PRICING_CPV_ToolsNoBidSuggestBidV2FilteringPricing      ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_CPV"
	PRICING_CPC_ToolsNoBidSuggestBidV2FilteringPricing      ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_CPC"
	PRICING_CPC_OCPM_ToolsNoBidSuggestBidV2FilteringPricing ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_CPC_OCPM"
	PRICING_CPA_ToolsNoBidSuggestBidV2FilteringPricing      ToolsNoBidSuggestBidV2FilteringPricing = "PRICING_CPA"
)

// All allowed values of ToolsNoBidSuggestBidV2FilteringPricing enum
var AllowedToolsNoBidSuggestBidV2FilteringPricingEnumValues = []ToolsNoBidSuggestBidV2FilteringPricing{
	"PRICING_OCPM",
	"PRICING_CPM",
	"PRICING_OCPC",
	"PRICING_CPV",
	"PRICING_CPC",
	"PRICING_CPC_OCPM",
	"PRICING_CPA",
}

// NewToolsNoBidSuggestBidV2FilteringPricingFromValue returns a pointer to a valid ToolsNoBidSuggestBidV2FilteringPricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsNoBidSuggestBidV2FilteringPricingFromValue(v string) (*ToolsNoBidSuggestBidV2FilteringPricing, error) {
	ev := ToolsNoBidSuggestBidV2FilteringPricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsNoBidSuggestBidV2FilteringPricing: valid values are %v", v, AllowedToolsNoBidSuggestBidV2FilteringPricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsNoBidSuggestBidV2FilteringPricing) IsValid() bool {
	for _, existing := range AllowedToolsNoBidSuggestBidV2FilteringPricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_no_bid_suggest_bid_v2_filtering_pricing value
func (v ToolsNoBidSuggestBidV2FilteringPricing) Ptr() *ToolsNoBidSuggestBidV2FilteringPricing {
	return &v
}
