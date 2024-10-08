/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsOrangeSiteGetV30Status
type ToolsOrangeSiteGetV30Status string

// List of tools_orange_site_get_v3.0_status
const (
	SITE_ALL_ToolsOrangeSiteGetV30Status      ToolsOrangeSiteGetV30Status = "SITE_ALL"
	SITE_OFFLINE_ToolsOrangeSiteGetV30Status  ToolsOrangeSiteGetV30Status = "SITE_OFFLINE"
	SITE_ONLINE_ToolsOrangeSiteGetV30Status   ToolsOrangeSiteGetV30Status = "SITE_ONLINE"
	SITE_REJECTED_ToolsOrangeSiteGetV30Status ToolsOrangeSiteGetV30Status = "SITE_REJECTED"
	SITE_TRASH_ToolsOrangeSiteGetV30Status    ToolsOrangeSiteGetV30Status = "SITE_TRASH"
)

// Ptr returns reference to tools_orange_site_get_v3.0_status value
func (v ToolsOrangeSiteGetV30Status) Ptr() *ToolsOrangeSiteGetV30Status {
	return &v
}
