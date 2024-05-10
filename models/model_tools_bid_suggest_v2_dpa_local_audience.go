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

// ToolsBidSuggestV2DpaLocalAudience
type ToolsBidSuggestV2DpaLocalAudience int64

// List of tools_bid_suggest_v2_dpa_local_audience
const (
	Enum_0_ToolsBidSuggestV2DpaLocalAudience ToolsBidSuggestV2DpaLocalAudience = 0
	Enum_1_ToolsBidSuggestV2DpaLocalAudience ToolsBidSuggestV2DpaLocalAudience = 1
)

// All allowed values of ToolsBidSuggestV2DpaLocalAudience enum
var AllowedToolsBidSuggestV2DpaLocalAudienceEnumValues = []ToolsBidSuggestV2DpaLocalAudience{
	0,
	1,
}

// NewToolsBidSuggestV2DpaLocalAudienceFromValue returns a pointer to a valid ToolsBidSuggestV2DpaLocalAudience
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2DpaLocalAudienceFromValue(v int64) (*ToolsBidSuggestV2DpaLocalAudience, error) {
	ev := ToolsBidSuggestV2DpaLocalAudience(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2DpaLocalAudience: valid values are %v", v, AllowedToolsBidSuggestV2DpaLocalAudienceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2DpaLocalAudience) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2DpaLocalAudienceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_dpa_local_audience value
func (v ToolsBidSuggestV2DpaLocalAudience) Ptr() *ToolsBidSuggestV2DpaLocalAudience {
	return &v
}
