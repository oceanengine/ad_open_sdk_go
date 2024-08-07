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

// ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType
type ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_v2_data_error_list_all_account_by_company_account_type
const (
	AD_ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType   ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType = "AD"
	BP_ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType   ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType = "BP"
	STAR_ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType = "STAR"
)

// All allowed values of ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType{
	"AD",
	"BP",
	"STAR",
}

// NewToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_data_error_list_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareV2DataErrorListAllAccountByCompanyAccountType {
	return &v
}
