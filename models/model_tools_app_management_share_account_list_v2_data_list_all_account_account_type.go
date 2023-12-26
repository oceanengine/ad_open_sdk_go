/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementShareAccountListV2DataListAllAccountAccountType
type ToolsAppManagementShareAccountListV2DataListAllAccountAccountType string

// List of tools_app_management_share_account_list_v2_data_list_all_account_account_type
const (
	AD_ToolsAppManagementShareAccountListV2DataListAllAccountAccountType   ToolsAppManagementShareAccountListV2DataListAllAccountAccountType = "AD"
	STAR_ToolsAppManagementShareAccountListV2DataListAllAccountAccountType ToolsAppManagementShareAccountListV2DataListAllAccountAccountType = "STAR"
	BP_ToolsAppManagementShareAccountListV2DataListAllAccountAccountType   ToolsAppManagementShareAccountListV2DataListAllAccountAccountType = "BP"
)

// All allowed values of ToolsAppManagementShareAccountListV2DataListAllAccountAccountType enum
var AllowedToolsAppManagementShareAccountListV2DataListAllAccountAccountTypeEnumValues = []ToolsAppManagementShareAccountListV2DataListAllAccountAccountType{
	"AD",
	"STAR",
	"BP",
}

// NewToolsAppManagementShareAccountListV2DataListAllAccountAccountTypeFromValue returns a pointer to a valid ToolsAppManagementShareAccountListV2DataListAllAccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementShareAccountListV2DataListAllAccountAccountTypeFromValue(v string) (*ToolsAppManagementShareAccountListV2DataListAllAccountAccountType, error) {
	ev := ToolsAppManagementShareAccountListV2DataListAllAccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementShareAccountListV2DataListAllAccountAccountType: valid values are %v", v, AllowedToolsAppManagementShareAccountListV2DataListAllAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementShareAccountListV2DataListAllAccountAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementShareAccountListV2DataListAllAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_share_account_list_v2_data_list_all_account_account_type value
func (v ToolsAppManagementShareAccountListV2DataListAllAccountAccountType) Ptr() *ToolsAppManagementShareAccountListV2DataListAllAccountAccountType {
	return &v
}
