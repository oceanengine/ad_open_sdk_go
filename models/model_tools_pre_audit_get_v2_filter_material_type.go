/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPreAuditGetV2FilterMaterialType
type ToolsPreAuditGetV2FilterMaterialType string

// List of tools_pre_audit_get_v2_filter_material_type
const (
	IMG_ToolsPreAuditGetV2FilterMaterialType   ToolsPreAuditGetV2FilterMaterialType = "IMG"
	SITE_ToolsPreAuditGetV2FilterMaterialType  ToolsPreAuditGetV2FilterMaterialType = "SITE"
	TITLE_ToolsPreAuditGetV2FilterMaterialType ToolsPreAuditGetV2FilterMaterialType = "TITLE"
	VIDEO_ToolsPreAuditGetV2FilterMaterialType ToolsPreAuditGetV2FilterMaterialType = "VIDEO"
)

// All allowed values of ToolsPreAuditGetV2FilterMaterialType enum
var AllowedToolsPreAuditGetV2FilterMaterialTypeEnumValues = []ToolsPreAuditGetV2FilterMaterialType{
	"IMG",
	"SITE",
	"TITLE",
	"VIDEO",
}

// NewToolsPreAuditGetV2FilterMaterialTypeFromValue returns a pointer to a valid ToolsPreAuditGetV2FilterMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPreAuditGetV2FilterMaterialTypeFromValue(v string) (*ToolsPreAuditGetV2FilterMaterialType, error) {
	ev := ToolsPreAuditGetV2FilterMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPreAuditGetV2FilterMaterialType: valid values are %v", v, AllowedToolsPreAuditGetV2FilterMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPreAuditGetV2FilterMaterialType) IsValid() bool {
	for _, existing := range AllowedToolsPreAuditGetV2FilterMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_pre_audit_get_v2_filter_material_type value
func (v ToolsPreAuditGetV2FilterMaterialType) Ptr() *ToolsPreAuditGetV2FilterMaterialType {
	return &v
}
