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

// ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType
type ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType string

// List of tools_bp_asset_management_share_cancel_v3.0_data_error_list_account_info_account_type
const (
	AD_ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType = "AD"
	BP_ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType = "BP"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType enum
var AllowedToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues = []ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType{
	"AD",
	"BP",
}

// NewToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType, error) {
	ev := ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_data_error_list_account_info_account_type value
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType) Ptr() *ToolsBpAssetManagementShareCancelV30DataErrorListAccountInfoAccountType {
	return &v
}
