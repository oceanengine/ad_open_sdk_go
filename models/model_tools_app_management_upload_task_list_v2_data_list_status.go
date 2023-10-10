/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementUploadTaskListV2DataListStatus
type ToolsAppManagementUploadTaskListV2DataListStatus string

// List of tools_app_management_upload_task_list_v2_data_list_status
const (
	FAILED_ToolsAppManagementUploadTaskListV2DataListStatus  ToolsAppManagementUploadTaskListV2DataListStatus = "FAILED"
	RUNNING_ToolsAppManagementUploadTaskListV2DataListStatus ToolsAppManagementUploadTaskListV2DataListStatus = "RUNNING"
	SUCCESS_ToolsAppManagementUploadTaskListV2DataListStatus ToolsAppManagementUploadTaskListV2DataListStatus = "SUCCESS"
	WAITING_ToolsAppManagementUploadTaskListV2DataListStatus ToolsAppManagementUploadTaskListV2DataListStatus = "WAITING"
)

// All allowed values of ToolsAppManagementUploadTaskListV2DataListStatus enum
var AllowedToolsAppManagementUploadTaskListV2DataListStatusEnumValues = []ToolsAppManagementUploadTaskListV2DataListStatus{
	"FAILED",
	"RUNNING",
	"SUCCESS",
	"WAITING",
}

// NewToolsAppManagementUploadTaskListV2DataListStatusFromValue returns a pointer to a valid ToolsAppManagementUploadTaskListV2DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementUploadTaskListV2DataListStatusFromValue(v string) (*ToolsAppManagementUploadTaskListV2DataListStatus, error) {
	ev := ToolsAppManagementUploadTaskListV2DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementUploadTaskListV2DataListStatus: valid values are %v", v, AllowedToolsAppManagementUploadTaskListV2DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementUploadTaskListV2DataListStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementUploadTaskListV2DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_upload_task_list_v2_data_list_status value
func (v ToolsAppManagementUploadTaskListV2DataListStatus) Ptr() *ToolsAppManagementUploadTaskListV2DataListStatus {
	return &v
}
