/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30DataListMaxPaymentTierRange
type ToolsMicroAppListV30DataListMaxPaymentTierRange string

// List of tools_micro_app_list_v3.0_data_list_max_payment_tier_range
const (
	ABOVE_1000_ToolsMicroAppListV30DataListMaxPaymentTierRange       ToolsMicroAppListV30DataListMaxPaymentTierRange = "ABOVE_1000"
	BELOW_500_ToolsMicroAppListV30DataListMaxPaymentTierRange        ToolsMicroAppListV30DataListMaxPaymentTierRange = "BELOW_500"
	FROM_500_TO_1000_ToolsMicroAppListV30DataListMaxPaymentTierRange ToolsMicroAppListV30DataListMaxPaymentTierRange = "FROM_500_TO_1000"
)

// Ptr returns reference to tools_micro_app_list_v3.0_data_list_max_payment_tier_range value
func (v ToolsMicroAppListV30DataListMaxPaymentTierRange) Ptr() *ToolsMicroAppListV30DataListMaxPaymentTierRange {
	return &v
}