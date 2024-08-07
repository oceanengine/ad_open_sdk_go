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

// ToolsBidsSuggestV30Pricing
type ToolsBidsSuggestV30Pricing string

// List of tools_bids_suggest_v3.0_pricing
const (
	PRICING_CPC_ToolsBidsSuggestV30Pricing  ToolsBidsSuggestV30Pricing = "PRICING_CPC"
	PRICING_CPM_ToolsBidsSuggestV30Pricing  ToolsBidsSuggestV30Pricing = "PRICING_CPM"
	PRICING_OCPC_ToolsBidsSuggestV30Pricing ToolsBidsSuggestV30Pricing = "PRICING_OCPC"
	PRICING_OCPM_ToolsBidsSuggestV30Pricing ToolsBidsSuggestV30Pricing = "PRICING_OCPM"
)

// All allowed values of ToolsBidsSuggestV30Pricing enum
var AllowedToolsBidsSuggestV30PricingEnumValues = []ToolsBidsSuggestV30Pricing{
	"PRICING_CPC",
	"PRICING_CPM",
	"PRICING_OCPC",
	"PRICING_OCPM",
}

// NewToolsBidsSuggestV30PricingFromValue returns a pointer to a valid ToolsBidsSuggestV30Pricing
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidsSuggestV30PricingFromValue(v string) (*ToolsBidsSuggestV30Pricing, error) {
	ev := ToolsBidsSuggestV30Pricing(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidsSuggestV30Pricing: valid values are %v", v, AllowedToolsBidsSuggestV30PricingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidsSuggestV30Pricing) IsValid() bool {
	for _, existing := range AllowedToolsBidsSuggestV30PricingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bids_suggest_v3.0_pricing value
func (v ToolsBidsSuggestV30Pricing) Ptr() *ToolsBidsSuggestV30Pricing {
	return &v
}
