/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileMaterialAuditResultGetV2MaterialInfosMaterialType
type FileMaterialAuditResultGetV2MaterialInfosMaterialType string

// List of file_material_audit_result_get_v2_material_infos_material_type
const (
	VIDEO_FileMaterialAuditResultGetV2MaterialInfosMaterialType   FileMaterialAuditResultGetV2MaterialInfosMaterialType = "VIDEO"
	PICTURE_FileMaterialAuditResultGetV2MaterialInfosMaterialType FileMaterialAuditResultGetV2MaterialInfosMaterialType = "PICTURE"
	TITLE_FileMaterialAuditResultGetV2MaterialInfosMaterialType   FileMaterialAuditResultGetV2MaterialInfosMaterialType = "TITLE"
)

// All allowed values of FileMaterialAuditResultGetV2MaterialInfosMaterialType enum
var AllowedFileMaterialAuditResultGetV2MaterialInfosMaterialTypeEnumValues = []FileMaterialAuditResultGetV2MaterialInfosMaterialType{
	"VIDEO",
	"PICTURE",
	"TITLE",
}

// NewFileMaterialAuditResultGetV2MaterialInfosMaterialTypeFromValue returns a pointer to a valid FileMaterialAuditResultGetV2MaterialInfosMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileMaterialAuditResultGetV2MaterialInfosMaterialTypeFromValue(v string) (*FileMaterialAuditResultGetV2MaterialInfosMaterialType, error) {
	ev := FileMaterialAuditResultGetV2MaterialInfosMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileMaterialAuditResultGetV2MaterialInfosMaterialType: valid values are %v", v, AllowedFileMaterialAuditResultGetV2MaterialInfosMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileMaterialAuditResultGetV2MaterialInfosMaterialType) IsValid() bool {
	for _, existing := range AllowedFileMaterialAuditResultGetV2MaterialInfosMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_material_audit_result_get_v2_material_infos_material_type value
func (v FileMaterialAuditResultGetV2MaterialInfosMaterialType) Ptr() *FileMaterialAuditResultGetV2MaterialInfosMaterialType {
	return &v
}
