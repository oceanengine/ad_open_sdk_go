/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupGetV2DataListSitesSiteAuditStatus
type ToolsLandingGroupGetV2DataListSitesSiteAuditStatus string

// List of tools_landing_group_get_v2_data_list_sites_site_audit_status
const (
	CREATED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus  ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "CREATED"
	ACCEPTED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "ACCEPTED"
	UNKNOWN_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus  ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "UNKNOWN"
	REJECTED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "REJECTED"
	AUDITING_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "AUDITING"
	BANNED_ToolsLandingGroupGetV2DataListSitesSiteAuditStatus   ToolsLandingGroupGetV2DataListSitesSiteAuditStatus = "BANNED"
)

// Ptr returns reference to tools_landing_group_get_v2_data_list_sites_site_audit_status value
func (v ToolsLandingGroupGetV2DataListSitesSiteAuditStatus) Ptr() *ToolsLandingGroupGetV2DataListSitesSiteAuditStatus {
	return &v
}
