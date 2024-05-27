/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAwemeAuthRenewalV2AuthType
type ToolsAwemeAuthRenewalV2AuthType string

// List of tools_aweme_auth_renewal_v2_auth_type
const (
	AWEME_ACCOUNT_ToolsAwemeAuthRenewalV2AuthType  ToolsAwemeAuthRenewalV2AuthType = "AWEME_ACCOUNT"
	AWEME_HOMEPAGE_ToolsAwemeAuthRenewalV2AuthType ToolsAwemeAuthRenewalV2AuthType = "AWEME_HOMEPAGE"
	LIVE_ACCOUNT_ToolsAwemeAuthRenewalV2AuthType   ToolsAwemeAuthRenewalV2AuthType = "LIVE_ACCOUNT"
	VIDEO_ITEM_ToolsAwemeAuthRenewalV2AuthType     ToolsAwemeAuthRenewalV2AuthType = "VIDEO_ITEM"
)

// All allowed values of ToolsAwemeAuthRenewalV2AuthType enum
var AllowedToolsAwemeAuthRenewalV2AuthTypeEnumValues = []ToolsAwemeAuthRenewalV2AuthType{
	"AWEME_ACCOUNT",
	"AWEME_HOMEPAGE",
	"LIVE_ACCOUNT",
	"VIDEO_ITEM",
}

// NewToolsAwemeAuthRenewalV2AuthTypeFromValue returns a pointer to a valid ToolsAwemeAuthRenewalV2AuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAwemeAuthRenewalV2AuthTypeFromValue(v string) (*ToolsAwemeAuthRenewalV2AuthType, error) {
	ev := ToolsAwemeAuthRenewalV2AuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAwemeAuthRenewalV2AuthType: valid values are %v", v, AllowedToolsAwemeAuthRenewalV2AuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAwemeAuthRenewalV2AuthType) IsValid() bool {
	for _, existing := range AllowedToolsAwemeAuthRenewalV2AuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_aweme_auth_renewal_v2_auth_type value
func (v ToolsAwemeAuthRenewalV2AuthType) Ptr() *ToolsAwemeAuthRenewalV2AuthType {
	return &v
}
