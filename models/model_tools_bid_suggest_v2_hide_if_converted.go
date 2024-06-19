/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsBidSuggestV2HideIfConverted
type ToolsBidSuggestV2HideIfConverted string

// List of tools_bid_suggest_v2_hide_if_converted
const (
	CUSTOMER_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CUSTOMER"
	AD_ToolsBidSuggestV2HideIfConverted           ToolsBidSuggestV2HideIfConverted = "AD"
	NO_EXCLUDE_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "NO_EXCLUDE"
	APP_ToolsBidSuggestV2HideIfConverted          ToolsBidSuggestV2HideIfConverted = "APP"
	CAMPAIGN_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CAMPAIGN"
	ADVERTISER_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "ADVERTISER"
	ORGANIZATION_ToolsBidSuggestV2HideIfConverted ToolsBidSuggestV2HideIfConverted = "ORGANIZATION"
)

// All allowed values of ToolsBidSuggestV2HideIfConverted enum
var AllowedToolsBidSuggestV2HideIfConvertedEnumValues = []ToolsBidSuggestV2HideIfConverted{
	"CUSTOMER",
	"AD",
	"NO_EXCLUDE",
	"APP",
	"CAMPAIGN",
	"ADVERTISER",
	"ORGANIZATION",
}

// NewToolsBidSuggestV2HideIfConvertedFromValue returns a pointer to a valid ToolsBidSuggestV2HideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidSuggestV2HideIfConvertedFromValue(v string) (*ToolsBidSuggestV2HideIfConverted, error) {
	ev := ToolsBidSuggestV2HideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidSuggestV2HideIfConverted: valid values are %v", v, AllowedToolsBidSuggestV2HideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidSuggestV2HideIfConverted) IsValid() bool {
	for _, existing := range AllowedToolsBidSuggestV2HideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bid_suggest_v2_hide_if_converted value
func (v ToolsBidSuggestV2HideIfConverted) Ptr() *ToolsBidSuggestV2HideIfConverted {
	return &v
}
