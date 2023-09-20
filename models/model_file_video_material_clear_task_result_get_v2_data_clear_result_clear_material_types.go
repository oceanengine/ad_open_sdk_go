/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes
type FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes string

// List of file_video_material_clear_task_result_get_v2_data_clear_result_clear_material_types
const (
	INEFFICIENT_MATERIAL_FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes   FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes = "INEFFICIENT_MATERIAL"
	SIMILAR_MATERIAL_FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes       FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes = "SIMILAR_MATERIAL"
	SIMILAR_QUEUE_MATERIAL_FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes = "SIMILAR_QUEUE_MATERIAL"
)

// All allowed values of FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes enum
var AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypesEnumValues = []FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes{
	"INEFFICIENT_MATERIAL",
	"SIMILAR_MATERIAL",
	"SIMILAR_QUEUE_MATERIAL",
}

// NewFileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypesFromValue returns a pointer to a valid FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypesFromValue(v string) (*FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes, error) {
	ev := FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes: valid values are %v", v, AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_result_get_v2_data_clear_result_clear_material_types value
func (v FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes) Ptr() *FileVideoMaterialClearTaskResultGetV2DataClearResultClearMaterialTypes {
	return &v
}
