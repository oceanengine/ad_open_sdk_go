/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementBpShareV2AllAccountsAccountType
type ToolsAppManagementBpShareV2AllAccountsAccountType string

// List of tools_app_management_bp_share_v2_all_accounts_account_type
const (
	AD_ToolsAppManagementBpShareV2AllAccountsAccountType   ToolsAppManagementBpShareV2AllAccountsAccountType = "AD"
	STAR_ToolsAppManagementBpShareV2AllAccountsAccountType ToolsAppManagementBpShareV2AllAccountsAccountType = "STAR"
	BP_ToolsAppManagementBpShareV2AllAccountsAccountType   ToolsAppManagementBpShareV2AllAccountsAccountType = "BP"
)

// All allowed values of ToolsAppManagementBpShareV2AllAccountsAccountType enum
var AllowedToolsAppManagementBpShareV2AllAccountsAccountTypeEnumValues = []ToolsAppManagementBpShareV2AllAccountsAccountType{
	"AD",
	"STAR",
	"BP",
}

// NewToolsAppManagementBpShareV2AllAccountsAccountTypeFromValue returns a pointer to a valid ToolsAppManagementBpShareV2AllAccountsAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementBpShareV2AllAccountsAccountTypeFromValue(v string) (*ToolsAppManagementBpShareV2AllAccountsAccountType, error) {
	ev := ToolsAppManagementBpShareV2AllAccountsAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementBpShareV2AllAccountsAccountType: valid values are %v", v, AllowedToolsAppManagementBpShareV2AllAccountsAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementBpShareV2AllAccountsAccountType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementBpShareV2AllAccountsAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_bp_share_v2_all_accounts_account_type value
func (v ToolsAppManagementBpShareV2AllAccountsAccountType) Ptr() *ToolsAppManagementBpShareV2AllAccountsAccountType {
	return &v
}
