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

// ToolsAwemeAuthListV2FilteringAuthType
type ToolsAwemeAuthListV2FilteringAuthType string

// List of tools_aweme_auth_list_v2_filtering_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthListV2FilteringAuthType  ToolsAwemeAuthListV2FilteringAuthType = "AWEME_ACCOUNT"
	AWEME_HOMEPAGE_ToolsAwemeAuthListV2FilteringAuthType ToolsAwemeAuthListV2FilteringAuthType = "AWEME_HOMEPAGE"
	LIVE_ACCOUNT_ToolsAwemeAuthListV2FilteringAuthType   ToolsAwemeAuthListV2FilteringAuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthListV2FilteringAuthType     ToolsAwemeAuthListV2FilteringAuthType = "VIDEO_ITEM"
)

// All allowed values of ToolsAwemeAuthListV2FilteringAuthType enum
var AllowedToolsAwemeAuthListV2FilteringAuthTypeEnumValues = []ToolsAwemeAuthListV2FilteringAuthType{
	"AWEME_ACCOUNT",
	"AWEME_HOMEPAGE",
	"LIVE_ACCOUNT",
	"VIDEO_ITEM",
}

// NewToolsAwemeAuthListV2FilteringAuthTypeFromValue returns a pointer to a valid ToolsAwemeAuthListV2FilteringAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthListV2FilteringAuthTypeFromValue(v string) (*ToolsAwemeAuthListV2FilteringAuthType, error) {
	ev := ToolsAwemeAuthListV2FilteringAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthListV2FilteringAuthType: valid values are %v", v, AllowedToolsAwemeAuthListV2FilteringAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthListV2FilteringAuthType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthListV2FilteringAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_list_v2_filtering_auth_type value
func (v ToolsAwemeAuthListV2FilteringAuthType) Ptr() *ToolsAwemeAuthListV2FilteringAuthType {
	return &v
}
