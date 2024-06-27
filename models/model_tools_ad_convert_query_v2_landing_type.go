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

// ToolsAdConvertQueryV2LandingType
type ToolsAdConvertQueryV2LandingType string

// List of tools_ad_convert_query_v2_landing_type
const (
	APP_ToolsAdConvertQueryV2LandingType       ToolsAdConvertQueryV2LandingType = "APP"
	LINK_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "LINK"
	DPA_ToolsAdConvertQueryV2LandingType       ToolsAdConvertQueryV2LandingType = "DPA"
	AWEME_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "AWEME"
	QUICK_APP_ToolsAdConvertQueryV2LandingType ToolsAdConvertQueryV2LandingType = "QUICK_APP"
	SHOP_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "SHOP"
	ARTICLE_ToolsAdConvertQueryV2LandingType   ToolsAdConvertQueryV2LandingType = "ARTICLE"
	STORE_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "STORE"
	LIVE_ToolsAdConvertQueryV2LandingType      ToolsAdConvertQueryV2LandingType = "LIVE"
	GOODS_ToolsAdConvertQueryV2LandingType     ToolsAdConvertQueryV2LandingType = "GOODS"
)

// All allowed values of ToolsAdConvertQueryV2LandingType enum
var AllowedToolsAdConvertQueryV2LandingTypeEnumValues = []ToolsAdConvertQueryV2LandingType{
	"APP",
	"LINK",
	"DPA",
	"AWEME",
	"QUICK_APP",
	"SHOP",
	"ARTICLE",
	"STORE",
	"LIVE",
	"GOODS",
}

// NewToolsAdConvertQueryV2LandingTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2LandingTypeFromValue(v string) (*ToolsAdConvertQueryV2LandingType, error) {
	ev := ToolsAdConvertQueryV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2LandingType: valid values are %v", v, AllowedToolsAdConvertQueryV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2LandingType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_landing_type value
func (v ToolsAdConvertQueryV2LandingType) Ptr() *ToolsAdConvertQueryV2LandingType {
	return &v
}
