/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppCreateV30MidPaymentTierRange
type ToolsMicroAppCreateV30MidPaymentTierRange string

// List of tools_micro_app_create_v3.0_mid_payment_tier_range
const (
	ABOVE_500_ToolsMicroAppCreateV30MidPaymentTierRange       ToolsMicroAppCreateV30MidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsMicroAppCreateV30MidPaymentTierRange       ToolsMicroAppCreateV30MidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsMicroAppCreateV30MidPaymentTierRange ToolsMicroAppCreateV30MidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_micro_app_create_v3.0_mid_payment_tier_range value
func (v ToolsMicroAppCreateV30MidPaymentTierRange) Ptr() *ToolsMicroAppCreateV30MidPaymentTierRange {
	return &v
}
