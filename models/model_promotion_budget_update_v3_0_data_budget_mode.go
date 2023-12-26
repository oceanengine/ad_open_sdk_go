/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionBudgetUpdateV30DataBudgetMode
type PromotionBudgetUpdateV30DataBudgetMode string

// List of promotion_budget_update_v3.0_data_budget_mode
const (
	BUDGET_MODE_DAY_PromotionBudgetUpdateV30DataBudgetMode   PromotionBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_PromotionBudgetUpdateV30DataBudgetMode PromotionBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of PromotionBudgetUpdateV30DataBudgetMode enum
var AllowedPromotionBudgetUpdateV30DataBudgetModeEnumValues = []PromotionBudgetUpdateV30DataBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
}

// NewPromotionBudgetUpdateV30DataBudgetModeFromValue returns a pointer to a valid PromotionBudgetUpdateV30DataBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionBudgetUpdateV30DataBudgetModeFromValue(v string) (*PromotionBudgetUpdateV30DataBudgetMode, error) {
	ev := PromotionBudgetUpdateV30DataBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionBudgetUpdateV30DataBudgetMode: valid values are %v", v, AllowedPromotionBudgetUpdateV30DataBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionBudgetUpdateV30DataBudgetMode) IsValid() bool {
	for _, existing := range AllowedPromotionBudgetUpdateV30DataBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_budget_update_v3.0_data_budget_mode value
func (v PromotionBudgetUpdateV30DataBudgetMode) Ptr() *PromotionBudgetUpdateV30DataBudgetMode {
	return &v
}
