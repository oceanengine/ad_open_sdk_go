/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPreAuditGetV2FilterMaterialType
type ToolsPreAuditGetV2FilterMaterialType string

// List of tools_pre_audit_get_v2_filter_material_type
const (
	IMG_ToolsPreAuditGetV2FilterMaterialType   ToolsPreAuditGetV2FilterMaterialType = "IMG"
	SITE_ToolsPreAuditGetV2FilterMaterialType  ToolsPreAuditGetV2FilterMaterialType = "SITE"
	TITLE_ToolsPreAuditGetV2FilterMaterialType ToolsPreAuditGetV2FilterMaterialType = "TITLE"
	VIDEO_ToolsPreAuditGetV2FilterMaterialType ToolsPreAuditGetV2FilterMaterialType = "VIDEO"
)

// Ptr returns reference to tools_pre_audit_get_v2_filter_material_type value
func (v ToolsPreAuditGetV2FilterMaterialType) Ptr() *ToolsPreAuditGetV2FilterMaterialType {
	return &v
}
