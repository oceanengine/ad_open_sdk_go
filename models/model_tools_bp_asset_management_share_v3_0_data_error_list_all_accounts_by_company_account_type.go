/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType
type ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType string

// List of tools_bp_asset_management_share_v3.0_data_error_list_all_accounts_by_company_account_type
const (
	AD_ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType = "AD"
)

// All allowed values of ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType enum
var AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues = []ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType{
	"AD",
}

// NewToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType, error) {
	ev := ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_data_error_list_all_accounts_by_company_account_type value
func (v ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType) Ptr() *ToolsBpAssetManagementShareV30DataErrorListAllAccountsByCompanyAccountType {
	return &v
}