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

// ToolsBidSuggestV2HideIfConverted
type ToolsBidSuggestV2HideIfConverted string

// List of tools_bid_suggest_v2_hide_if_converted
const (
	ORGANIZATION_ToolsBidSuggestV2HideIfConverted ToolsBidSuggestV2HideIfConverted = "ORGANIZATION"
	CAMPAIGN_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CAMPAIGN"
	ADVERTISER_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "ADVERTISER"
	APP_ToolsBidSuggestV2HideIfConverted          ToolsBidSuggestV2HideIfConverted = "APP"
	CUSTOMER_ToolsBidSuggestV2HideIfConverted     ToolsBidSuggestV2HideIfConverted = "CUSTOMER"
	NO_EXCLUDE_ToolsBidSuggestV2HideIfConverted   ToolsBidSuggestV2HideIfConverted = "NO_EXCLUDE"
	AD_ToolsBidSuggestV2HideIfConverted           ToolsBidSuggestV2HideIfConverted = "AD"
)

// All allowed values of ToolsBidSuggestV2HideIfConverted enum
var AllowedToolsBidSuggestV2HideIfConvertedEnumValues = []ToolsBidSuggestV2HideIfConverted{
	"ORGANIZATION",
	"CAMPAIGN",
	"ADVERTISER",
	"APP",
	"CUSTOMER",
	"NO_EXCLUDE",
	"AD",
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
