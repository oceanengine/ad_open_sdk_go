/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthCancelV2AuthType
type ToolsAwemeAuthCancelV2AuthType string

// List of tools_aweme_auth_cancel_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthCancelV2AuthType  ToolsAwemeAuthCancelV2AuthType = "AWEME_ACCOUNT"
	AWEME_HOMEPAGE_ToolsAwemeAuthCancelV2AuthType ToolsAwemeAuthCancelV2AuthType = "AWEME_HOMEPAGE"
	LIVE_ACCOUNT_ToolsAwemeAuthCancelV2AuthType   ToolsAwemeAuthCancelV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthCancelV2AuthType     ToolsAwemeAuthCancelV2AuthType = "VIDEO_ITEM"
)

// Ptr returns reference to tools_aweme_auth_cancel_v2_auth_type value
func (v ToolsAwemeAuthCancelV2AuthType) Ptr() *ToolsAwemeAuthCancelV2AuthType {
	return &v
}
