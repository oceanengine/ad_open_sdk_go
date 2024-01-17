/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementExtendPackageUpdateV2V2AccountType
type ToolsAppManagementExtendPackageUpdateV2V2AccountType string

// List of tools_app_management_extend_package_update_v2_v2_account_type
const (
	AD_ToolsAppManagementExtendPackageUpdateV2V2AccountType   ToolsAppManagementExtendPackageUpdateV2V2AccountType = "AD"
	BP_ToolsAppManagementExtendPackageUpdateV2V2AccountType   ToolsAppManagementExtendPackageUpdateV2V2AccountType = "BP"
	STAR_ToolsAppManagementExtendPackageUpdateV2V2AccountType ToolsAppManagementExtendPackageUpdateV2V2AccountType = "STAR"
)

// All allowed values of ToolsAppManagementExtendPackageUpdateV2V2AccountType enum
var AllowedToolsAppManagementExtendPackageUpdateV2V2AccountTypeEnumValues = []ToolsAppManagementExtendPackageUpdateV2V2AccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementExtendPackageUpdateV2V2AccountTypeFromValue returns a pointer to a valid ToolsAppManagementExtendPackageUpdateV2V2AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementExtendPackageUpdateV2V2AccountTypeFromValue(v string) (*ToolsAppManagementExtendPackageUpdateV2V2AccountType, error) {
	ev := ToolsAppManagementExtendPackageUpdateV2V2AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementExtendPackageUpdateV2V2AccountType: valid values are %v", v, AllowedToolsAppManagementExtendPackageUpdateV2V2AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementExtendPackageUpdateV2V2AccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementExtendPackageUpdateV2V2AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_extend_package_update_v2_v2_account_type value
func (v ToolsAppManagementExtendPackageUpdateV2V2AccountType) Ptr() *ToolsAppManagementExtendPackageUpdateV2V2AccountType {
	return &v
}
