/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30FilteringDeliveryType
type BudgetGroupListV30FilteringDeliveryType string

// List of budget_group_list_v3.0_filtering_delivery_type
const (
	MANUAL_BudgetGroupListV30FilteringDeliveryType     BudgetGroupListV30FilteringDeliveryType = "MANUAL"
	PROCEDURAL_BudgetGroupListV30FilteringDeliveryType BudgetGroupListV30FilteringDeliveryType = "PROCEDURAL"
)

// Ptr returns reference to budget_group_list_v3.0_filtering_delivery_type value
func (v BudgetGroupListV30FilteringDeliveryType) Ptr() *BudgetGroupListV30FilteringDeliveryType {
	return &v
}
