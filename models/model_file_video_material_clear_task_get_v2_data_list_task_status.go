/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileVideoMaterialClearTaskGetV2DataListTaskStatus
type FileVideoMaterialClearTaskGetV2DataListTaskStatus string

// List of file_video_material_clear_task_get_v2_data_list_task_status
const (
	DONE_FileVideoMaterialClearTaskGetV2DataListTaskStatus    FileVideoMaterialClearTaskGetV2DataListTaskStatus = "DONE"
	RUNNING_FileVideoMaterialClearTaskGetV2DataListTaskStatus FileVideoMaterialClearTaskGetV2DataListTaskStatus = "RUNNING"
	TIMEOUT_FileVideoMaterialClearTaskGetV2DataListTaskStatus FileVideoMaterialClearTaskGetV2DataListTaskStatus = "TIMEOUT"
)

// All allowed values of FileVideoMaterialClearTaskGetV2DataListTaskStatus enum
var AllowedFileVideoMaterialClearTaskGetV2DataListTaskStatusEnumValues = []FileVideoMaterialClearTaskGetV2DataListTaskStatus{
	"DONE",
	"RUNNING",
	"TIMEOUT",
}

// NewFileVideoMaterialClearTaskGetV2DataListTaskStatusFromValue returns a pointer to a valid FileVideoMaterialClearTaskGetV2DataListTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileVideoMaterialClearTaskGetV2DataListTaskStatusFromValue(v string) (*FileVideoMaterialClearTaskGetV2DataListTaskStatus, error) {
	ev := FileVideoMaterialClearTaskGetV2DataListTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileVideoMaterialClearTaskGetV2DataListTaskStatus: valid values are %v", v, AllowedFileVideoMaterialClearTaskGetV2DataListTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileVideoMaterialClearTaskGetV2DataListTaskStatus) IsValid() bool {
	for _, existing := range AllowedFileVideoMaterialClearTaskGetV2DataListTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_video_material_clear_task_get_v2_data_list_task_status value
func (v FileVideoMaterialClearTaskGetV2DataListTaskStatus) Ptr() *FileVideoMaterialClearTaskGetV2DataListTaskStatus {
	return &v
}
