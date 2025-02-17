/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAipThirdSiteGetV2DataAuditStatus
type ToolsAipThirdSiteGetV2DataAuditStatus string

// List of tools_aip_third_site_get_v2_data_audit_status
const (
	AUDIT_UNKNOW_ToolsAipThirdSiteGetV2DataAuditStatus   ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_UNKNOW"
	AUDIT_ACCEPTED_ToolsAipThirdSiteGetV2DataAuditStatus ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_ToolsAipThirdSiteGetV2DataAuditStatus ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_REJECTED"
	AUDIT_BANNED_ToolsAipThirdSiteGetV2DataAuditStatus   ToolsAipThirdSiteGetV2DataAuditStatus = "AUDIT_BANNED"
	AUDITING_ToolsAipThirdSiteGetV2DataAuditStatus       ToolsAipThirdSiteGetV2DataAuditStatus = "AUDITING"
	AWAIT_AUDIT_ToolsAipThirdSiteGetV2DataAuditStatus    ToolsAipThirdSiteGetV2DataAuditStatus = "AWAIT_AUDIT"
)

// Ptr returns reference to tools_aip_third_site_get_v2_data_audit_status value
func (v ToolsAipThirdSiteGetV2DataAuditStatus) Ptr() *ToolsAipThirdSiteGetV2DataAuditStatus {
	return &v
}
