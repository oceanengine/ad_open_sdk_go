/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletUpdateV30MidPaymentTierRange
type ToolsWechatAppletUpdateV30MidPaymentTierRange string

// List of tools_wechat_applet_update_v3.0_mid_payment_tier_range
const (
	ABOVE_500_ToolsWechatAppletUpdateV30MidPaymentTierRange       ToolsWechatAppletUpdateV30MidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsWechatAppletUpdateV30MidPaymentTierRange       ToolsWechatAppletUpdateV30MidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsWechatAppletUpdateV30MidPaymentTierRange ToolsWechatAppletUpdateV30MidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_wechat_applet_update_v3.0_mid_payment_tier_range value
func (v ToolsWechatAppletUpdateV30MidPaymentTierRange) Ptr() *ToolsWechatAppletUpdateV30MidPaymentTierRange {
	return &v
}
