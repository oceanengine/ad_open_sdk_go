/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppUpdateV30MidPaymentTierRange
type ToolsMicroAppUpdateV30MidPaymentTierRange string

// List of tools_micro_app_update_v3.0_mid_payment_tier_range
const (
	ABOVE_500_ToolsMicroAppUpdateV30MidPaymentTierRange       ToolsMicroAppUpdateV30MidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsMicroAppUpdateV30MidPaymentTierRange       ToolsMicroAppUpdateV30MidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsMicroAppUpdateV30MidPaymentTierRange ToolsMicroAppUpdateV30MidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_micro_app_update_v3.0_mid_payment_tier_range value
func (v ToolsMicroAppUpdateV30MidPaymentTierRange) Ptr() *ToolsMicroAppUpdateV30MidPaymentTierRange {
	return &v
}
