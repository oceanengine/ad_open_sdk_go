/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupCreateV30DeliveryMode
type BudgetGroupCreateV30DeliveryMode string

// List of budget_group_create_v3.0_delivery_mode
const (
	MANUAL_BudgetGroupCreateV30DeliveryMode     BudgetGroupCreateV30DeliveryMode = "MANUAL"
	PROCEDURAL_BudgetGroupCreateV30DeliveryMode BudgetGroupCreateV30DeliveryMode = "PROCEDURAL"
)

// Ptr returns reference to budget_group_create_v3.0_delivery_mode value
func (v BudgetGroupCreateV30DeliveryMode) Ptr() *BudgetGroupCreateV30DeliveryMode {
	return &v
}
