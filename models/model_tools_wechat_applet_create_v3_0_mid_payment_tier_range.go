/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletCreateV30MidPaymentTierRange
type ToolsWechatAppletCreateV30MidPaymentTierRange string

// List of tools_wechat_applet_create_v3.0_mid_payment_tier_range
const (
	ABOVE_500_ToolsWechatAppletCreateV30MidPaymentTierRange       ToolsWechatAppletCreateV30MidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsWechatAppletCreateV30MidPaymentTierRange       ToolsWechatAppletCreateV30MidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsWechatAppletCreateV30MidPaymentTierRange ToolsWechatAppletCreateV30MidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_wechat_applet_create_v3.0_mid_payment_tier_range value
func (v ToolsWechatAppletCreateV30MidPaymentTierRange) Ptr() *ToolsWechatAppletCreateV30MidPaymentTierRange {
	return &v
}
