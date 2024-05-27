/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupUpdateV2ResponseDataSitesInner struct for ToolsLandingGroupUpdateV2ResponseDataSitesInner
type ToolsLandingGroupUpdateV2ResponseDataSitesInner struct {
	// 成员 ID，即站点在落地页组中的唯一标识
	MemberId *string `json:"member_id,omitempty"`
	// 站点 ID
	SiteId *string `json:"site_id,omitempty"`
	// 站点启用状态。
	SiteOptStatus *string `json:"site_opt_status,omitempty"`
	// 站点URL
	SiteUrl *string `json:"site_url,omitempty"`
}
