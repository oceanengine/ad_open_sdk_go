/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletListV30DataListMaxPaymentTierRange
type ToolsWechatAppletListV30DataListMaxPaymentTierRange string

// List of tools_wechat_applet_list_v3.0_data_list_max_payment_tier_range
const (
	ABOVE_1000_ToolsWechatAppletListV30DataListMaxPaymentTierRange       ToolsWechatAppletListV30DataListMaxPaymentTierRange = "ABOVE_1000"
	BELOW_500_ToolsWechatAppletListV30DataListMaxPaymentTierRange        ToolsWechatAppletListV30DataListMaxPaymentTierRange = "BELOW_500"
	FROM_500_TO_1000_ToolsWechatAppletListV30DataListMaxPaymentTierRange ToolsWechatAppletListV30DataListMaxPaymentTierRange = "FROM_500_TO_1000"
)

// Ptr returns reference to tools_wechat_applet_list_v3.0_data_list_max_payment_tier_range value
func (v ToolsWechatAppletListV30DataListMaxPaymentTierRange) Ptr() *ToolsWechatAppletListV30DataListMaxPaymentTierRange {
	return &v
}