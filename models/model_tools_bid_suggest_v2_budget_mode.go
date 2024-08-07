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

// ToolsBidSuggestV2BudgetMode
type ToolsBidSuggestV2BudgetMode string

// List of tools_bid_suggest_v2_budget_mode
const (
	BUDGET_MODE_DAY_ToolsBidSuggestV2BudgetMode      ToolsBidSuggestV2BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_ToolsBidSuggestV2BudgetMode ToolsBidSuggestV2BudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_ToolsBidSuggestV2BudgetMode    ToolsBidSuggestV2BudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of ToolsBidSuggestV2BudgetMode enum
var AllowedToolsBidSuggestV2BudgetModeEnumValues = []ToolsBidSuggestV2BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewToolsBidSuggestV2BudgetModeFromValue returns a pointer to a valid ToolsBidSuggestV2BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2BudgetModeFromValue(v string) (*ToolsBidSuggestV2BudgetMode, error) {
	ev := ToolsBidSuggestV2BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2BudgetMode: valid values are %v", v, AllowedToolsBidSuggestV2BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2BudgetMode) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_budget_mode value
func (v ToolsBidSuggestV2BudgetMode) Ptr() *ToolsBidSuggestV2BudgetMode {
	return &v
}
