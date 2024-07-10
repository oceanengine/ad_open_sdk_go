/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAppManagementShareAccountListV2SearchType
type ToolsAppManagementShareAccountListV2SearchType string

// List of tools_app_management_share_account_list_v2_search_type
const (
	OTHER_ToolsAppManagementShareAccountListV2SearchType              ToolsAppManagementShareAccountListV2SearchType = "OTHER"
	ORGANIZATION_SHARE_ToolsAppManagementShareAccountListV2SearchType ToolsAppManagementShareAccountListV2SearchType = "ORGANIZATION_SHARE"
)

// All allowed values of ToolsAppManagementShareAccountListV2SearchType enum
var AllowedToolsAppManagementShareAccountListV2SearchTypeEnumValues = []ToolsAppManagementShareAccountListV2SearchType{
	"OTHER",
	"ORGANIZATION_SHARE",
}

// NewToolsAppManagementShareAccountListV2SearchTypeFromValue returns a pointer to a valid ToolsAppManagementShareAccountListV2SearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementShareAccountListV2SearchTypeFromValue(v string) (*ToolsAppManagementShareAccountListV2SearchType, error) {
	ev := ToolsAppManagementShareAccountListV2SearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementShareAccountListV2SearchType: valid values are %v", v, AllowedToolsAppManagementShareAccountListV2SearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementShareAccountListV2SearchType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementShareAccountListV2SearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_share_account_list_v2_search_type value
func (v ToolsAppManagementShareAccountListV2SearchType) Ptr() *ToolsAppManagementShareAccountListV2SearchType {
	return &v
}
