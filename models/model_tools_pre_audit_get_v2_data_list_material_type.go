/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPreAuditGetV2DataListMaterialType
type ToolsPreAuditGetV2DataListMaterialType string

// List of tools_pre_audit_get_v2_data_list_material_type
const (
	IMG_ToolsPreAuditGetV2DataListMaterialType   ToolsPreAuditGetV2DataListMaterialType = "IMG"
	SITE_ToolsPreAuditGetV2DataListMaterialType  ToolsPreAuditGetV2DataListMaterialType = "SITE"
	TITLE_ToolsPreAuditGetV2DataListMaterialType ToolsPreAuditGetV2DataListMaterialType = "TITLE"
	VIDEO_ToolsPreAuditGetV2DataListMaterialType ToolsPreAuditGetV2DataListMaterialType = "VIDEO"
)

// All allowed values of ToolsPreAuditGetV2DataListMaterialType enum
var AllowedToolsPreAuditGetV2DataListMaterialTypeEnumValues = []ToolsPreAuditGetV2DataListMaterialType{
	"IMG",
	"SITE",
	"TITLE",
	"VIDEO",
}

// NewToolsPreAuditGetV2DataListMaterialTypeFromValue returns a pointer to a valid ToolsPreAuditGetV2DataListMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPreAuditGetV2DataListMaterialTypeFromValue(v string) (*ToolsPreAuditGetV2DataListMaterialType, error) {
	ev := ToolsPreAuditGetV2DataListMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPreAuditGetV2DataListMaterialType: valid values are %v", v, AllowedToolsPreAuditGetV2DataListMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPreAuditGetV2DataListMaterialType) IsValid() bool {
	for _, existing := range AllowedToolsPreAuditGetV2DataListMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pre_audit_get_v2_data_list_material_type value
func (v ToolsPreAuditGetV2DataListMaterialType) Ptr() *ToolsPreAuditGetV2DataListMaterialType {
	return &v
}
