/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileAutoGenerateSourceGetV2MaterialsMaterialType
type FileAutoGenerateSourceGetV2MaterialsMaterialType string

// List of file_auto_generate_source_get_v2_materials_material_type
const (
	VIDEO_FileAutoGenerateSourceGetV2MaterialsMaterialType FileAutoGenerateSourceGetV2MaterialsMaterialType = "video"
	IMAGE_FileAutoGenerateSourceGetV2MaterialsMaterialType FileAutoGenerateSourceGetV2MaterialsMaterialType = "image"
)

// All allowed values of FileAutoGenerateSourceGetV2MaterialsMaterialType enum
var AllowedFileAutoGenerateSourceGetV2MaterialsMaterialTypeEnumValues = []FileAutoGenerateSourceGetV2MaterialsMaterialType{
	"video",
	"image",
}

// NewFileAutoGenerateSourceGetV2MaterialsMaterialTypeFromValue returns a pointer to a valid FileAutoGenerateSourceGetV2MaterialsMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileAutoGenerateSourceGetV2MaterialsMaterialTypeFromValue(v string) (*FileAutoGenerateSourceGetV2MaterialsMaterialType, error) {
	ev := FileAutoGenerateSourceGetV2MaterialsMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileAutoGenerateSourceGetV2MaterialsMaterialType: valid values are %v", v, AllowedFileAutoGenerateSourceGetV2MaterialsMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileAutoGenerateSourceGetV2MaterialsMaterialType) IsValid() bool {
	for _, existing := range AllowedFileAutoGenerateSourceGetV2MaterialsMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_auto_generate_source_get_v2_materials_material_type value
func (v FileAutoGenerateSourceGetV2MaterialsMaterialType) Ptr() *FileAutoGenerateSourceGetV2MaterialsMaterialType {
	return &v
}
