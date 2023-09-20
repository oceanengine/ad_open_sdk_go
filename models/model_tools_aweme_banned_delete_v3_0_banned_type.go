/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeBannedDeleteV30BannedType
type ToolsAwemeBannedDeleteV30BannedType string

// List of tools_aweme_banned_delete_v3.0_banned_type
const (
	AWEME_TYPE_ToolsAwemeBannedDeleteV30BannedType  ToolsAwemeBannedDeleteV30BannedType = "AWEME_TYPE"
	CUSTOM_TYPE_ToolsAwemeBannedDeleteV30BannedType ToolsAwemeBannedDeleteV30BannedType = "CUSTOM_TYPE"
)

// All allowed values of ToolsAwemeBannedDeleteV30BannedType enum
var AllowedToolsAwemeBannedDeleteV30BannedTypeEnumValues = []ToolsAwemeBannedDeleteV30BannedType{
	"AWEME_TYPE",
	"CUSTOM_TYPE",
}

// NewToolsAwemeBannedDeleteV30BannedTypeFromValue returns a pointer to a valid ToolsAwemeBannedDeleteV30BannedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeBannedDeleteV30BannedTypeFromValue(v string) (*ToolsAwemeBannedDeleteV30BannedType, error) {
	ev := ToolsAwemeBannedDeleteV30BannedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeBannedDeleteV30BannedType: valid values are %v", v, AllowedToolsAwemeBannedDeleteV30BannedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeBannedDeleteV30BannedType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeBannedDeleteV30BannedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_banned_delete_v3.0_banned_type value
func (v ToolsAwemeBannedDeleteV30BannedType) Ptr() *ToolsAwemeBannedDeleteV30BannedType {
	return &v
}
