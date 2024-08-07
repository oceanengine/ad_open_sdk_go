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

// ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp
type ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp string

// List of tools_bp_asset_management_share_v3.0_data_error_list_all_accounts_by_bp
const (
	AD_ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp = "AD"
)

// All allowed values of ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp enum
var AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByBpEnumValues = []ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp{
	"AD",
}

// NewToolsBpAssetManagementShareV30DataErrorListAllAccountsByBpFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30DataErrorListAllAccountsByBpFromValue(v string) (*ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp, error) {
	ev := ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp: valid values are %v", v, AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByBpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByBpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_data_error_list_all_accounts_by_bp value
func (v ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp) Ptr() *ToolsBpAssetManagementShareV30DataErrorListAllAccountsByBp {
	return &v
}
