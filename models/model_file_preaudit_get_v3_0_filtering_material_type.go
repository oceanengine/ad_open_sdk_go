/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FilePreauditGetV30FilteringMaterialType
type FilePreauditGetV30FilteringMaterialType string

// List of file_preaudit_get_v3.0_filtering_material_type
const (
	VIDEO_FilePreauditGetV30FilteringMaterialType FilePreauditGetV30FilteringMaterialType = "VIDEO"
)

// All allowed values of FilePreauditGetV30FilteringMaterialType enum
var AllowedFilePreauditGetV30FilteringMaterialTypeEnumValues = []FilePreauditGetV30FilteringMaterialType{
	"VIDEO",
}

// NewFilePreauditGetV30FilteringMaterialTypeFromValue returns a pointer to a valid FilePreauditGetV30FilteringMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditGetV30FilteringMaterialTypeFromValue(v string) (*FilePreauditGetV30FilteringMaterialType, error) {
	ev := FilePreauditGetV30FilteringMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditGetV30FilteringMaterialType: valid values are %v", v, AllowedFilePreauditGetV30FilteringMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditGetV30FilteringMaterialType) IsValid() bool {
	for _, existing := range AllowedFilePreauditGetV30FilteringMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_get_v3.0_filtering_material_type value
func (v FilePreauditGetV30FilteringMaterialType) Ptr() *FilePreauditGetV30FilteringMaterialType {
	return &v
}
