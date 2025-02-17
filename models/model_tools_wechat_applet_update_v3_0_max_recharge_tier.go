/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletUpdateV30MaxRechargeTier
type ToolsWechatAppletUpdateV30MaxRechargeTier string

// List of tools_wechat_applet_update_v3.0_max_recharge_tier
const (
	ABOVE_200_ToolsWechatAppletUpdateV30MaxRechargeTier             ToolsWechatAppletUpdateV30MaxRechargeTier = "ABOVE_200"
	FROM_100_TO_200_ToolsWechatAppletUpdateV30MaxRechargeTier       ToolsWechatAppletUpdateV30MaxRechargeTier = "FROM_100_TO_200"
	FROM_FIFTY_TO_HUNDRED_ToolsWechatAppletUpdateV30MaxRechargeTier ToolsWechatAppletUpdateV30MaxRechargeTier = "FROM_FIFTY_TO_HUNDRED"
	FROM_ONE_TO_FIFTY_ToolsWechatAppletUpdateV30MaxRechargeTier     ToolsWechatAppletUpdateV30MaxRechargeTier = "FROM_ONE_TO_FIFTY"
)

// Ptr returns reference to tools_wechat_applet_update_v3.0_max_recharge_tier value
func (v ToolsWechatAppletUpdateV30MaxRechargeTier) Ptr() *ToolsWechatAppletUpdateV30MaxRechargeTier {
	return &v
}
