/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult
type FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult string

// List of file_video_material_clear_task_result_get_v2_data_clear_result_clear_result
const (
	FAIL_FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult         FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult = "FAIL"
	PART_SUCCESS_FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult = "PART_SUCCESS"
	SUCCESS_FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult      FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult = "SUCCESS"
)

// All allowed values of FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult enum
var AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearResultEnumValues = []FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult{
	"FAIL",
	"PART_SUCCESS",
	"SUCCESS",
}

// NewFileVideoMaterialClearTaskResultGetV2DataClearResultClearResultFromValue returns a pointer to a valid FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskResultGetV2DataClearResultClearResultFromValue(v string) (*FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult, error) {
	ev := FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult: valid values are %v", v, AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskResultGetV2DataClearResultClearResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_result_get_v2_data_clear_result_clear_result value
func (v FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult) Ptr() *FileVideoMaterialClearTaskResultGetV2DataClearResultClearResult {
	return &v
}
