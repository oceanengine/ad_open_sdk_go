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

// ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType
type ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType string

// List of tools_app_management_android_basic_package_update_v2_files_file_type
const (
	APK_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType   ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType = "APK"
	IMAGE_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType = "IMAGE"
	VIDEO_ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType = "VIDEO"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType enum
var AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTypeEnumValues = []ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType{
	"APK",
	"IMAGE",
	"VIDEO",
}

// NewToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTypeFromValue(v string) (*ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType, error) {
	ev := ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageUpdateV2FilesFileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_update_v2_files_file_type value
func (v ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType) Ptr() *ToolsAppManagementAndroidBasicPackageUpdateV2FilesFileType {
	return &v
}
