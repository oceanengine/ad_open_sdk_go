/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus
type ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus string

// List of tools_app_management_android_basic_package_get_v2_data_current_version_status
const (
	ALL_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus            ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "ALL"
	AUDIT_ACCEPTED_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus    ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "AUDIT_REJECTED"
	ENABLE_ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus         ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus = "ENABLE"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus enum
var AllowedToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatusEnumValues = []ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus{
	"ALL",
	"AUDIT_ACCEPTED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"ENABLE",
}

// NewToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatusFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatusFromValue(v string) (*ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus, error) {
	ev := ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_get_v2_data_current_version_status value
func (v ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus) Ptr() *ToolsAppManagementAndroidBasicPackageGetV2DataCurrentVersionStatus {
	return &v
}
