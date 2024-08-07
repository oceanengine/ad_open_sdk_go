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

// ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType
type ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType string

// List of tools_app_management_bp_share_v2_data_error_list_all_account_account_type
const (
	AD_ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType   ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType = "AD"
	BP_ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType   ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType = "BP"
	STAR_ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType enum
var AllowedToolsAppManagementBpShareV2DataErrorListAllAccountAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementBpShareV2DataErrorListAllAccountAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataErrorListAllAccountAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataErrorListAllAccountAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataErrorListAllAccountAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_error_list_all_account_account_type value
func (v ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType) Ptr() *ToolsAppManagementBpShareV2DataErrorListAllAccountAccountType {
	return &v
}
