/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameCreateV30MidPaymentTierRange
type ToolsWechatGameCreateV30MidPaymentTierRange string

// List of tools_wechat_game_create_v3.0_mid_payment_tier_range
const (
	ABOVE_500_ToolsWechatGameCreateV30MidPaymentTierRange       ToolsWechatGameCreateV30MidPaymentTierRange = "ABOVE_500"
	BELOW_100_ToolsWechatGameCreateV30MidPaymentTierRange       ToolsWechatGameCreateV30MidPaymentTierRange = "BELOW_100"
	FROM_100_TO_500_ToolsWechatGameCreateV30MidPaymentTierRange ToolsWechatGameCreateV30MidPaymentTierRange = "FROM_100_TO_500"
)

// Ptr returns reference to tools_wechat_game_create_v3.0_mid_payment_tier_range value
func (v ToolsWechatGameCreateV30MidPaymentTierRange) Ptr() *ToolsWechatGameCreateV30MidPaymentTierRange {
	return &v
}
