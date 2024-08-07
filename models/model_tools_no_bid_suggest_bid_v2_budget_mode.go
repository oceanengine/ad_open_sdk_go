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

// ToolsNoBidSuggestBidV2BudgetMode
type ToolsNoBidSuggestBidV2BudgetMode string

// List of tools_no_bid_suggest_bid_v2_budget_mode
const (
	BUDGET_MODE_DAY_ToolsNoBidSuggestBidV2BudgetMode      ToolsNoBidSuggestBidV2BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ToolsNoBidSuggestBidV2BudgetMode ToolsNoBidSuggestBidV2BudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_ToolsNoBidSuggestBidV2BudgetMode    ToolsNoBidSuggestBidV2BudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of ToolsNoBidSuggestBidV2BudgetMode enum
var AllowedToolsNoBidSuggestBidV2BudgetModeEnumValues = []ToolsNoBidSuggestBidV2BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewToolsNoBidSuggestBidV2BudgetModeFromValue returns a pointer to a valid ToolsNoBidSuggestBidV2BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsNoBidSuggestBidV2BudgetModeFromValue(v string) (*ToolsNoBidSuggestBidV2BudgetMode, error) {
	ev := ToolsNoBidSuggestBidV2BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsNoBidSuggestBidV2BudgetMode: valid values are %v", v, AllowedToolsNoBidSuggestBidV2BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsNoBidSuggestBidV2BudgetMode) IsValid() bool {
	for _, existing := range AllowedToolsNoBidSuggestBidV2BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_no_bid_suggest_bid_v2_budget_mode value
func (v ToolsNoBidSuggestBidV2BudgetMode) Ptr() *ToolsNoBidSuggestBidV2BudgetMode {
	return &v
}
