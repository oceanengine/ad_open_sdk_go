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

// ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus
type ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus string

// List of tools_app_management_android_basic_package_get_v2_data_next_version_status
const (
	ALL_ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus            ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus = "ALL"
	AUDIT_ACCEPTED_ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus = "AUDIT_ACCEPTED"
	AUDIT_DOING_ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus    ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus = "AUDIT_DOING"
	AUDIT_REJECTED_ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus = "AUDIT_REJECTED"
	ENABLE_ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus         ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus = "ENABLE"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus enum
var AllowedToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatusEnumValues = []ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus{
	"ALL",
	"AUDIT_ACCEPTED",
	"AUDIT_DOING",
	"AUDIT_REJECTED",
	"ENABLE",
}

// NewToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatusFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatusFromValue(v string) (*ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus, error) {
	ev := ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_get_v2_data_next_version_status value
func (v ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus) Ptr() *ToolsAppManagementAndroidBasicPackageGetV2DataNextVersionStatus {
	return &v
}
