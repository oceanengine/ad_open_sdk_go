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

// ToolsAppManagementAndroidBasicPackageGetV2AccountType
type ToolsAppManagementAndroidBasicPackageGetV2AccountType string

// List of tools_app_management_android_basic_package_get_v2_account_type
const (
	AD_ToolsAppManagementAndroidBasicPackageGetV2AccountType ToolsAppManagementAndroidBasicPackageGetV2AccountType = "AD"
	BP_ToolsAppManagementAndroidBasicPackageGetV2AccountType ToolsAppManagementAndroidBasicPackageGetV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageGetV2AccountType enum
var AllowedToolsAppManagementAndroidBasicPackageGetV2AccountTypeEnumValues = []ToolsAppManagementAndroidBasicPackageGetV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementAndroidBasicPackageGetV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageGetV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageGetV2AccountTypeFromValue(v string) (*ToolsAppManagementAndroidBasicPackageGetV2AccountType, error) {
	ev := ToolsAppManagementAndroidBasicPackageGetV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageGetV2AccountType: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageGetV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageGetV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageGetV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_get_v2_account_type value
func (v ToolsAppManagementAndroidBasicPackageGetV2AccountType) Ptr() *ToolsAppManagementAndroidBasicPackageGetV2AccountType {
	return &v
}
