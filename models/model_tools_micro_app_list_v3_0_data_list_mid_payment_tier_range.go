/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30DataListMidPaymentTierRange
type ToolsMicroAppListV30DataListMidPaymentTierRange string

// List of tools_micro_app_list_v3.0_data_list_mid_payment_tier_range
const (
	ABOVE_500_ToolsMicroAppListV30DataListMidPaymentTierRange       ToolsMicroAppListV30DataListMidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsMicroAppListV30DataListMidPaymentTierRange       ToolsMicroAppListV30DataListMidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsMicroAppListV30DataListMidPaymentTierRange ToolsMicroAppListV30DataListMidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_micro_app_list_v3.0_data_list_mid_payment_tier_range value
func (v ToolsMicroAppListV30DataListMidPaymentTierRange) Ptr() *ToolsMicroAppListV30DataListMidPaymentTierRange {
	return &v
}
