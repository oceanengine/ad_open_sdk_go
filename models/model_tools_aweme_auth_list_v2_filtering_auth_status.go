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

// ToolsAwemeAuthListV2FilteringAuthStatus
type ToolsAwemeAuthListV2FilteringAuthStatus string

// List of tools_aweme_auth_list_v2_filtering_auth_status
const (
	AUTHRIZED_ToolsAwemeAuthListV2FilteringAuthStatus  ToolsAwemeAuthListV2FilteringAuthStatus = "AUTHRIZED"
	AUTHRIZING_ToolsAwemeAuthListV2FilteringAuthStatus ToolsAwemeAuthListV2FilteringAuthStatus = "AUTHRIZING"
	INVALID_ToolsAwemeAuthListV2FilteringAuthStatus    ToolsAwemeAuthListV2FilteringAuthStatus = "INVALID"
)

// All allowed values of ToolsAwemeAuthListV2FilteringAuthStatus enum
var AllowedToolsAwemeAuthListV2FilteringAuthStatusEnumValues = []ToolsAwemeAuthListV2FilteringAuthStatus{
	"AUTHRIZED",
	"AUTHRIZING",
	"INVALID",
}

// NewToolsAwemeAuthListV2FilteringAuthStatusFromValue returns a pointer to a valid ToolsAwemeAuthListV2FilteringAuthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2FilteringAuthStatusFromValue(v string) (*ToolsAwemeAuthListV2FilteringAuthStatus, error) {
	ev := ToolsAwemeAuthListV2FilteringAuthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2FilteringAuthStatus: valid values are %v", v, AllowedToolsAwemeAuthListV2FilteringAuthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2FilteringAuthStatus) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2FilteringAuthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_filtering_auth_status value
func (v ToolsAwemeAuthListV2FilteringAuthStatus) Ptr() *ToolsAwemeAuthListV2FilteringAuthStatus {
	return &v
}
