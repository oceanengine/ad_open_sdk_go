/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsIsSupportUniversalGetV2LandingType
type ToolsIsSupportUniversalGetV2LandingType string

// List of tools_is_support_universal_get_v2_landing_type
const (
	APP_ToolsIsSupportUniversalGetV2LandingType       ToolsIsSupportUniversalGetV2LandingType = "APP"
	ARTICLE_ToolsIsSupportUniversalGetV2LandingType   ToolsIsSupportUniversalGetV2LandingType = "ARTICLE"
	AWEME_ToolsIsSupportUniversalGetV2LandingType     ToolsIsSupportUniversalGetV2LandingType = "AWEME"
	DPA_ToolsIsSupportUniversalGetV2LandingType       ToolsIsSupportUniversalGetV2LandingType = "DPA"
	GOODS_ToolsIsSupportUniversalGetV2LandingType     ToolsIsSupportUniversalGetV2LandingType = "GOODS"
	LINK_ToolsIsSupportUniversalGetV2LandingType      ToolsIsSupportUniversalGetV2LandingType = "LINK"
	QUICK_APP_ToolsIsSupportUniversalGetV2LandingType ToolsIsSupportUniversalGetV2LandingType = "QUICK_APP"
	SHOP_ToolsIsSupportUniversalGetV2LandingType      ToolsIsSupportUniversalGetV2LandingType = "SHOP"
)

// All allowed values of ToolsIsSupportUniversalGetV2LandingType enum
var AllowedToolsIsSupportUniversalGetV2LandingTypeEnumValues = []ToolsIsSupportUniversalGetV2LandingType{
	"APP",
	"ARTICLE",
	"AWEME",
	"DPA",
	"GOODS",
	"LINK",
	"QUICK_APP",
	"SHOP",
}

// NewToolsIsSupportUniversalGetV2LandingTypeFromValue returns a pointer to a valid ToolsIsSupportUniversalGetV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIsSupportUniversalGetV2LandingTypeFromValue(v string) (*ToolsIsSupportUniversalGetV2LandingType, error) {
	ev := ToolsIsSupportUniversalGetV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIsSupportUniversalGetV2LandingType: valid values are %v", v, AllowedToolsIsSupportUniversalGetV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIsSupportUniversalGetV2LandingType) IsValid() bool {
	for _, existing := range AllowedToolsIsSupportUniversalGetV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_is_support_universal_get_v2_landing_type value
func (v ToolsIsSupportUniversalGetV2LandingType) Ptr() *ToolsIsSupportUniversalGetV2LandingType {
	return &v
}
