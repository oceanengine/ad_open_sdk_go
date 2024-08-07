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

// ToolsAwemeBannedCreateV30BannedType
type ToolsAwemeBannedCreateV30BannedType string

// List of tools_aweme_banned_create_v3.0_banned_type
const (
	AWEME_TYPE_ToolsAwemeBannedCreateV30BannedType  ToolsAwemeBannedCreateV30BannedType = "AWEME_TYPE"
	CUSTOM_TYPE_ToolsAwemeBannedCreateV30BannedType ToolsAwemeBannedCreateV30BannedType = "CUSTOM_TYPE"
)

// All allowed values of ToolsAwemeBannedCreateV30BannedType enum
var AllowedToolsAwemeBannedCreateV30BannedTypeEnumValues = []ToolsAwemeBannedCreateV30BannedType{
	"AWEME_TYPE",
	"CUSTOM_TYPE",
}

// NewToolsAwemeBannedCreateV30BannedTypeFromValue returns a pointer to a valid ToolsAwemeBannedCreateV30BannedType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeBannedCreateV30BannedTypeFromValue(v string) (*ToolsAwemeBannedCreateV30BannedType, error) {
	ev := ToolsAwemeBannedCreateV30BannedType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeBannedCreateV30BannedType: valid values are %v", v, AllowedToolsAwemeBannedCreateV30BannedTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeBannedCreateV30BannedType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeBannedCreateV30BannedTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_banned_create_v3.0_banned_type value
func (v ToolsAwemeBannedCreateV30BannedType) Ptr() *ToolsAwemeBannedCreateV30BannedType {
	return &v
}
