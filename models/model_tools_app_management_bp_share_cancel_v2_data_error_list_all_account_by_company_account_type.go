/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType
type ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType string

// List of tools_app_management_bp_share_cancel_v2_data_error_list_all_account_by_company_account_type
const (
	STAR_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType = "STAR"
	BP_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType = "BP"
	AD_ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType   ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType = "AD"
)

// All allowed values of ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType enum
var AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountTypeEnumValues = []ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType{
	"STAR",
	"BP",
	"AD",
}

// NewToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountTypeFromValue(v string) (*ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType, error) {
	ev := ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_cancel_v2_data_error_list_all_account_by_company_account_type value
func (v ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType) Ptr() *ToolsAppManagementBpShareCancelV2DataErrorListAllAccountByCompanyAccountType {
	return &v
}