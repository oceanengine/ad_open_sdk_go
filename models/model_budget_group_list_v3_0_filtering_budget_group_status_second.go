/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BudgetGroupListV30FilteringBudgetGroupStatusSecond
type BudgetGroupListV30FilteringBudgetGroupStatusSecond string

// List of budget_group_list_v3.0_filtering_budget_group_status_second
const (
	ACCOUNT_EXCEEDED_BudgetGroupListV30FilteringBudgetGroupStatusSecond BudgetGroupListV30FilteringBudgetGroupStatusSecond = "ACCOUNT_EXCEEDED"
	GROUP_EXCEEDED_BudgetGroupListV30FilteringBudgetGroupStatusSecond   BudgetGroupListV30FilteringBudgetGroupStatusSecond = "GROUP_EXCEEDED"
)

// All allowed values of BudgetGroupListV30FilteringBudgetGroupStatusSecond enum
var AllowedBudgetGroupListV30FilteringBudgetGroupStatusSecondEnumValues = []BudgetGroupListV30FilteringBudgetGroupStatusSecond{
	"ACCOUNT_EXCEEDED",
	"GROUP_EXCEEDED",
}

// NewBudgetGroupListV30FilteringBudgetGroupStatusSecondFromValue returns a pointer to a valid BudgetGroupListV30FilteringBudgetGroupStatusSecond
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30FilteringBudgetGroupStatusSecondFromValue(v string) (*BudgetGroupListV30FilteringBudgetGroupStatusSecond, error) {
	ev := BudgetGroupListV30FilteringBudgetGroupStatusSecond(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30FilteringBudgetGroupStatusSecond: valid values are %v", v, AllowedBudgetGroupListV30FilteringBudgetGroupStatusSecondEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30FilteringBudgetGroupStatusSecond) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30FilteringBudgetGroupStatusSecondEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_filtering_budget_group_status_second value
func (v BudgetGroupListV30FilteringBudgetGroupStatusSecond) Ptr() *BudgetGroupListV30FilteringBudgetGroupStatusSecond {
	return &v
}