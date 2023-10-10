/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareCancelV30AllAccountsByBp
type ToolsBpAssetManagementShareCancelV30AllAccountsByBp string

// List of tools_bp_asset_management_share_cancel_v3.0_all_accounts_by_bp
const (
	AD_ToolsBpAssetManagementShareCancelV30AllAccountsByBp ToolsBpAssetManagementShareCancelV30AllAccountsByBp = "AD"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30AllAccountsByBp enum
var AllowedToolsBpAssetManagementShareCancelV30AllAccountsByBpEnumValues = []ToolsBpAssetManagementShareCancelV30AllAccountsByBp{
	"AD",
}

// NewToolsBpAssetManagementShareCancelV30AllAccountsByBpFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30AllAccountsByBp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30AllAccountsByBpFromValue(v string) (*ToolsBpAssetManagementShareCancelV30AllAccountsByBp, error) {
	ev := ToolsBpAssetManagementShareCancelV30AllAccountsByBp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30AllAccountsByBp: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30AllAccountsByBpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30AllAccountsByBp) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30AllAccountsByBpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_all_accounts_by_bp value
func (v ToolsBpAssetManagementShareCancelV30AllAccountsByBp) Ptr() *ToolsBpAssetManagementShareCancelV30AllAccountsByBp {
	return &v
}
