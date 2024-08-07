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

// ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType
type ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType string

// List of tools_bp_asset_management_share_v3.0_data_error_list_account_info_account_type
const (
	AD_ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType = "AD"
	BP_ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType = "BP"
)

// All allowed values of ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType enum
var AllowedToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountTypeEnumValues = []ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType{
	"AD",
	"BP",
}

// NewToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType, error) {
	ev := ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_data_error_list_account_info_account_type value
func (v ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType) Ptr() *ToolsBpAssetManagementShareV30DataErrorListAccountInfoAccountType {
	return &v
}
