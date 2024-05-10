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

// ToolsAwemeAuthCancelV2AuthType
type ToolsAwemeAuthCancelV2AuthType string

// List of tools_aweme_auth_cancel_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthCancelV2AuthType ToolsAwemeAuthCancelV2AuthType = "AWEME_ACCOUNT"
	LIVE_ACCOUNT_ToolsAwemeAuthCancelV2AuthType  ToolsAwemeAuthCancelV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthCancelV2AuthType    ToolsAwemeAuthCancelV2AuthType = "VIDEO_ITEM"
)

// All allowed values of ToolsAwemeAuthCancelV2AuthType enum
var AllowedToolsAwemeAuthCancelV2AuthTypeEnumValues = []ToolsAwemeAuthCancelV2AuthType{
	"AWEME_ACCOUNT",
	"LIVE_ACCOUNT",
	"VIDEO_ITEM",
}

// NewToolsAwemeAuthCancelV2AuthTypeFromValue returns a pointer to a valid ToolsAwemeAuthCancelV2AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthCancelV2AuthTypeFromValue(v string) (*ToolsAwemeAuthCancelV2AuthType, error) {
	ev := ToolsAwemeAuthCancelV2AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthCancelV2AuthType: valid values are %v", v, AllowedToolsAwemeAuthCancelV2AuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthCancelV2AuthType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthCancelV2AuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_cancel_v2_auth_type value
func (v ToolsAwemeAuthCancelV2AuthType) Ptr() *ToolsAwemeAuthCancelV2AuthType {
	return &v
}
