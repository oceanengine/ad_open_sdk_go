/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthV2AuthType
type ToolsAwemeAuthV2AuthType string

// List of tools_aweme_auth_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthV2AuthType  ToolsAwemeAuthV2AuthType = "AWEME_ACCOUNT"
	AWEME_HOMEPAGE_ToolsAwemeAuthV2AuthType ToolsAwemeAuthV2AuthType = "AWEME_HOMEPAGE"
	LIVE_ACCOUNT_ToolsAwemeAuthV2AuthType   ToolsAwemeAuthV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthV2AuthType     ToolsAwemeAuthV2AuthType = "VIDEO_ITEM"
)

// Ptr returns reference to tools_aweme_auth_v2_auth_type value
func (v ToolsAwemeAuthV2AuthType) Ptr() *ToolsAwemeAuthV2AuthType {
	return &v
}
