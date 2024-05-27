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

// FilePreauditGetV30DataListMaterialType
type FilePreauditGetV30DataListMaterialType string

// List of file_preaudit_get_v3.0_data_list_material_type
const (
	IMAGE_FilePreauditGetV30DataListMaterialType FilePreauditGetV30DataListMaterialType = "IMAGE"
	VIDEO_FilePreauditGetV30DataListMaterialType FilePreauditGetV30DataListMaterialType = "VIDEO"
)

// All allowed values of FilePreauditGetV30DataListMaterialType enum
var AllowedFilePreauditGetV30DataListMaterialTypeEnumValues = []FilePreauditGetV30DataListMaterialType{
	"IMAGE",
	"VIDEO",
}

// NewFilePreauditGetV30DataListMaterialTypeFromValue returns a pointer to a valid FilePreauditGetV30DataListMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilePreauditGetV30DataListMaterialTypeFromValue(v string) (*FilePreauditGetV30DataListMaterialType, error) {
	ev := FilePreauditGetV30DataListMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilePreauditGetV30DataListMaterialType: valid values are %v", v, AllowedFilePreauditGetV30DataListMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilePreauditGetV30DataListMaterialType) IsValid() bool {
	for _, existing := range AllowedFilePreauditGetV30DataListMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_preaudit_get_v3.0_data_list_material_type value
func (v FilePreauditGetV30DataListMaterialType) Ptr() *FilePreauditGetV30DataListMaterialType {
	return &v
}
