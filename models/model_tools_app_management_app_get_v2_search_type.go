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

// ToolsAppManagementAppGetV2SearchType
type ToolsAppManagementAppGetV2SearchType string

// List of tools_app_management_app_get_v2_search_type
const (
	ALL_ToolsAppManagementAppGetV2SearchType         ToolsAppManagementAppGetV2SearchType = "ALL"
	CREATE_ONLY_ToolsAppManagementAppGetV2SearchType ToolsAppManagementAppGetV2SearchType = "CREATE_ONLY"
	SHARED_ONLY_ToolsAppManagementAppGetV2SearchType ToolsAppManagementAppGetV2SearchType = "SHARED_ONLY"
)

// All allowed values of ToolsAppManagementAppGetV2SearchType enum
var AllowedToolsAppManagementAppGetV2SearchTypeEnumValues = []ToolsAppManagementAppGetV2SearchType{
	"ALL",
	"CREATE_ONLY",
	"SHARED_ONLY",
}

// NewToolsAppManagementAppGetV2SearchTypeFromValue returns a pointer to a valid ToolsAppManagementAppGetV2SearchType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAppManagementAppGetV2SearchTypeFromValue(v string) (*ToolsAppManagementAppGetV2SearchType, error) {
	ev := ToolsAppManagementAppGetV2SearchType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAppManagementAppGetV2SearchType: valid values are %v", v, AllowedToolsAppManagementAppGetV2SearchTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAppManagementAppGetV2SearchType) IsValid() bool {
	for _, existing := range AllowedToolsAppManagementAppGetV2SearchTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_app_management_app_get_v2_search_type value
func (v ToolsAppManagementAppGetV2SearchType) Ptr() *ToolsAppManagementAppGetV2SearchType {
	return &v
}
