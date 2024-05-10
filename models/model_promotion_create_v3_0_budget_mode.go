/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30BudgetMode
type PromotionCreateV30BudgetMode string

// List of promotion_create_v3.0_budget_mode
const (
	BUDGET_MODE_DAY_PromotionCreateV30BudgetMode   PromotionCreateV30BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_PromotionCreateV30BudgetMode PromotionCreateV30BudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of PromotionCreateV30BudgetMode enum
var AllowedPromotionCreateV30BudgetModeEnumValues = []PromotionCreateV30BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_TOTAL",
}

// NewPromotionCreateV30BudgetModeFromValue returns a pointer to a valid PromotionCreateV30BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30BudgetModeFromValue(v string) (*PromotionCreateV30BudgetMode, error) {
	ev := PromotionCreateV30BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30BudgetMode: valid values are %v", v, AllowedPromotionCreateV30BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30BudgetMode) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_budget_mode value
func (v PromotionCreateV30BudgetMode) Ptr() *PromotionCreateV30BudgetMode {
	return &v
}
