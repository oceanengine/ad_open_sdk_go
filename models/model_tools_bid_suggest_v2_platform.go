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

// ToolsBidSuggestV2Platform
type ToolsBidSuggestV2Platform string

// List of tools_bid_suggest_v2_platform
const (
	IPAD_ToolsBidSuggestV2Platform    ToolsBidSuggestV2Platform = "IPAD"
	NONE_ToolsBidSuggestV2Platform    ToolsBidSuggestV2Platform = "NONE"
	ANDROID_ToolsBidSuggestV2Platform ToolsBidSuggestV2Platform = "ANDROID"
	PC_ToolsBidSuggestV2Platform      ToolsBidSuggestV2Platform = "PC"
	IOS_ToolsBidSuggestV2Platform     ToolsBidSuggestV2Platform = "IOS"
	WAP_ToolsBidSuggestV2Platform     ToolsBidSuggestV2Platform = "WAP"
)

// All allowed values of ToolsBidSuggestV2Platform enum
var AllowedToolsBidSuggestV2PlatformEnumValues = []ToolsBidSuggestV2Platform{
	"IPAD",
	"NONE",
	"ANDROID",
	"PC",
	"IOS",
	"WAP",
}

// NewToolsBidSuggestV2PlatformFromValue returns a pointer to a valid ToolsBidSuggestV2Platform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2PlatformFromValue(v string) (*ToolsBidSuggestV2Platform, error) {
	ev := ToolsBidSuggestV2Platform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2Platform: valid values are %v", v, AllowedToolsBidSuggestV2PlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2Platform) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2PlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_platform value
func (v ToolsBidSuggestV2Platform) Ptr() *ToolsBidSuggestV2Platform {
	return &v
}
