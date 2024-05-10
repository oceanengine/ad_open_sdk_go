/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType
type ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType string

// List of tools_bp_asset_management_share_v3.0_all_accounts_by_company_account_type
const (
	AD_ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType = "AD"
)

// All allowed values of ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType enum
var AllowedToolsBpAssetManagementShareV30AllAccountsByCompanyAccountTypeEnumValues = []ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType{
	"AD",
}

// NewToolsBpAssetManagementShareV30AllAccountsByCompanyAccountTypeFromValue returns a pointer to a valid ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBpAssetManagementShareV30AllAccountsByCompanyAccountTypeFromValue(v string) (*ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType, error) {
	ev := ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType: valid values are %v", v, AllowedToolsBpAssetManagementShareV30AllAccountsByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsBpAssetManagementShareV30AllAccountsByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bp_asset_management_share_v3.0_all_accounts_by_company_account_type value
func (v ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType) Ptr() *ToolsBpAssetManagementShareV30AllAccountsByCompanyAccountType {
	return &v
}
