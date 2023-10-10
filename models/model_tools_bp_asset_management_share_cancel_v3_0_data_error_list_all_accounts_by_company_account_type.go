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

// ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType
type ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType string

// List of tools_bp_asset_management_share_cancel_v3.0_data_error_list_all_accounts_by_company_account_type
const (
	AD_ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType = "AD"
)

// All allowed values of ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType enum
var AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues = []ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType{
	"AD",
}

// NewToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType, error) {
	ev := ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_cancel_v3.0_data_error_list_all_accounts_by_company_account_type value
func (v ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType) Ptr() *ToolsBpAssetManagementShareCancelV30DataErrorListAllAccountsByCompanyAccountType {
	return &v
}
