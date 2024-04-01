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

// BudgetGroupListV30FilteringAdType
type BudgetGroupListV30FilteringAdType string

// List of budget_group_list_v3.0_filtering_ad_type
const (
	ALL_BudgetGroupListV30FilteringAdType    BudgetGroupListV30FilteringAdType = "ALL"
	SEARCH_BudgetGroupListV30FilteringAdType BudgetGroupListV30FilteringAdType = "SEARCH"
)

// All allowed values of BudgetGroupListV30FilteringAdType enum
var AllowedBudgetGroupListV30FilteringAdTypeEnumValues = []BudgetGroupListV30FilteringAdType{
	"ALL",
	"SEARCH",
}

// NewBudgetGroupListV30FilteringAdTypeFromValue returns a pointer to a valid BudgetGroupListV30FilteringAdType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30FilteringAdTypeFromValue(v string) (*BudgetGroupListV30FilteringAdType, error) {
	ev := BudgetGroupListV30FilteringAdType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30FilteringAdType: valid values are %v", v, AllowedBudgetGroupListV30FilteringAdTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30FilteringAdType) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30FilteringAdTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_filtering_ad_type value
func (v BudgetGroupListV30FilteringAdType) Ptr() *BudgetGroupListV30FilteringAdType {
	return &v
}
