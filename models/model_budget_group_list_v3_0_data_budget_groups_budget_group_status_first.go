/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst
type BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst string

// List of budget_group_list_v3.0_data_budget_groups_budget_group_status_first
const (
	DELETED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst      BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst = "DELETED"
	ENABLE_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst       BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst = "ENABLE"
	UNDELIVERIED_BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst = "UNDELIVERIED"
)

// All allowed values of BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst enum
var AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirstEnumValues = []BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst{
	"DELETED",
	"ENABLE",
	"UNDELIVERIED",
}

// NewBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirstFromValue returns a pointer to a valid BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirstFromValue(v string) (*BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst, error) {
	ev := BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst: valid values are %v", v, AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirstEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirstEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_data_budget_groups_budget_group_status_first value
func (v BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst) Ptr() *BudgetGroupListV30DataBudgetGroupsBudgetGroupStatusFirst {
	return &v
}
