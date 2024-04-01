/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2FilterOwnAwemeFans
type ToolsBidSuggestV2FilterOwnAwemeFans int64

// List of tools_bid_suggest_v2_filter_own_aweme_fans
const (
	Enum_0_ToolsBidSuggestV2FilterOwnAwemeFans ToolsBidSuggestV2FilterOwnAwemeFans = 0
	Enum_1_ToolsBidSuggestV2FilterOwnAwemeFans ToolsBidSuggestV2FilterOwnAwemeFans = 1
)

// All allowed values of ToolsBidSuggestV2FilterOwnAwemeFans enum
var AllowedToolsBidSuggestV2FilterOwnAwemeFansEnumValues = []ToolsBidSuggestV2FilterOwnAwemeFans{
	0,
	1,
}

// NewToolsBidSuggestV2FilterOwnAwemeFansFromValue returns a pointer to a valid ToolsBidSuggestV2FilterOwnAwemeFans
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2FilterOwnAwemeFansFromValue(v int64) (*ToolsBidSuggestV2FilterOwnAwemeFans, error) {
	ev := ToolsBidSuggestV2FilterOwnAwemeFans(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2FilterOwnAwemeFans: valid values are %v", v, AllowedToolsBidSuggestV2FilterOwnAwemeFansEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2FilterOwnAwemeFans) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2FilterOwnAwemeFansEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_filter_own_aweme_fans value
func (v ToolsBidSuggestV2FilterOwnAwemeFans) Ptr() *ToolsBidSuggestV2FilterOwnAwemeFans {
	return &v
}
