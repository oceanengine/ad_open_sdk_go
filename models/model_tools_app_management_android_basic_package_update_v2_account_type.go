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

// ToolsAppManagementAndroidBasicPackageUpdateV2AccountType
type ToolsAppManagementAndroidBasicPackageUpdateV2AccountType string

// List of tools_app_management_android_basic_package_update_v2_account_type
const (
	AD_ToolsAppManagementAndroidBasicPackageUpdateV2AccountType ToolsAppManagementAndroidBasicPackageUpdateV2AccountType = "AD"
	BP_ToolsAppManagementAndroidBasicPackageUpdateV2AccountType ToolsAppManagementAndroidBasicPackageUpdateV2AccountType = "BP"
)

// All allowed values of ToolsAppManagementAndroidBasicPackageUpdateV2AccountType enum
var AllowedToolsAppManagementAndroidBasicPackageUpdateV2AccountTypeEnumValues = []ToolsAppManagementAndroidBasicPackageUpdateV2AccountType{
	"AD",
	"BP",
}

// NewToolsAppManagementAndroidBasicPackageUpdateV2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementAndroidBasicPackageUpdateV2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAndroidBasicPackageUpdateV2AccountTypeFromValue(v string) (*ToolsAppManagementAndroidBasicPackageUpdateV2AccountType, error) {
	ev := ToolsAppManagementAndroidBasicPackageUpdateV2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAndroidBasicPackageUpdateV2AccountType: valid values are %v", v, AllowedToolsAppManagementAndroidBasicPackageUpdateV2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAndroidBasicPackageUpdateV2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAndroidBasicPackageUpdateV2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_android_basic_package_update_v2_account_type value
func (v ToolsAppManagementAndroidBasicPackageUpdateV2AccountType) Ptr() *ToolsAppManagementAndroidBasicPackageUpdateV2AccountType {
	return &v
}
