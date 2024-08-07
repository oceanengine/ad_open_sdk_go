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

// ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType
type ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_success_list_account_info_account_type
const (
	AD_ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType = "AD"
	BP_ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType   ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType = "BP"
	STAR_ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_success_list_account_info_account_type value
func (v ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataSuccessListAccountInfoAccountType {
	return &v
}
