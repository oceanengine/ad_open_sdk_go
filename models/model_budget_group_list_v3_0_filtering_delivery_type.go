/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BudgetGroupListV30FilteringDeliveryType
type BudgetGroupListV30FilteringDeliveryType string

// List of budget_group_list_v3.0_filtering_delivery_type
const (
	MANUAL_BudgetGroupListV30FilteringDeliveryType     BudgetGroupListV30FilteringDeliveryType = "MANUAL"
	PROCEDURAL_BudgetGroupListV30FilteringDeliveryType BudgetGroupListV30FilteringDeliveryType = "PROCEDURAL"
)

// All allowed values of BudgetGroupListV30FilteringDeliveryType enum
var AllowedBudgetGroupListV30FilteringDeliveryTypeEnumValues = []BudgetGroupListV30FilteringDeliveryType{
	"MANUAL",
	"PROCEDURAL",
}

// NewBudgetGroupListV30FilteringDeliveryTypeFromValue returns a pointer to a valid BudgetGroupListV30FilteringDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetGroupListV30FilteringDeliveryTypeFromValue(v string) (*BudgetGroupListV30FilteringDeliveryType, error) {
	ev := BudgetGroupListV30FilteringDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetGroupListV30FilteringDeliveryType: valid values are %v", v, AllowedBudgetGroupListV30FilteringDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetGroupListV30FilteringDeliveryType) IsValid() bool {
	for _, existing := range AllowedBudgetGroupListV30FilteringDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budget_group_list_v3.0_filtering_delivery_type value
func (v BudgetGroupListV30FilteringDeliveryType) Ptr() *BudgetGroupListV30FilteringDeliveryType {
	return &v
}
