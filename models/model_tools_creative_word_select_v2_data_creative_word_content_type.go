/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsCreativeWordSelectV2DataCreativeWordContentType
type ToolsCreativeWordSelectV2DataCreativeWordContentType string

// List of tools_creative_word_select_v2_data_creative_word_content_type
const (
	AGE_ToolsCreativeWordSelectV2DataCreativeWordContentType      ToolsCreativeWordSelectV2DataCreativeWordContentType = "AGE"
	CARRIER_ToolsCreativeWordSelectV2DataCreativeWordContentType  ToolsCreativeWordSelectV2DataCreativeWordContentType = "CARRIER"
	DATE_ToolsCreativeWordSelectV2DataCreativeWordContentType     ToolsCreativeWordSelectV2DataCreativeWordContentType = "DATE"
	DEVICE_ToolsCreativeWordSelectV2DataCreativeWordContentType   ToolsCreativeWordSelectV2DataCreativeWordContentType = "DEVICE"
	KEYWORD_ToolsCreativeWordSelectV2DataCreativeWordContentType  ToolsCreativeWordSelectV2DataCreativeWordContentType = "KEYWORD"
	NETWORK_ToolsCreativeWordSelectV2DataCreativeWordContentType  ToolsCreativeWordSelectV2DataCreativeWordContentType = "NETWORK"
	PROVINCE_ToolsCreativeWordSelectV2DataCreativeWordContentType ToolsCreativeWordSelectV2DataCreativeWordContentType = "PROVINCE"
	REGION_ToolsCreativeWordSelectV2DataCreativeWordContentType   ToolsCreativeWordSelectV2DataCreativeWordContentType = "REGION"
	RICE_ToolsCreativeWordSelectV2DataCreativeWordContentType     ToolsCreativeWordSelectV2DataCreativeWordContentType = "RICE"
	WEEK_ToolsCreativeWordSelectV2DataCreativeWordContentType     ToolsCreativeWordSelectV2DataCreativeWordContentType = "WEEK"
)

// All allowed values of ToolsCreativeWordSelectV2DataCreativeWordContentType enum
var AllowedToolsCreativeWordSelectV2DataCreativeWordContentTypeEnumValues = []ToolsCreativeWordSelectV2DataCreativeWordContentType{
	"AGE",
	"CARRIER",
	"DATE",
	"DEVICE",
	"KEYWORD",
	"NETWORK",
	"PROVINCE",
	"REGION",
	"RICE",
	"WEEK",
}

// NewToolsCreativeWordSelectV2DataCreativeWordContentTypeFromValue returns a pointer to a valid ToolsCreativeWordSelectV2DataCreativeWordContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCreativeWordSelectV2DataCreativeWordContentTypeFromValue(v string) (*ToolsCreativeWordSelectV2DataCreativeWordContentType, error) {
	ev := ToolsCreativeWordSelectV2DataCreativeWordContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCreativeWordSelectV2DataCreativeWordContentType: valid values are %v", v, AllowedToolsCreativeWordSelectV2DataCreativeWordContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCreativeWordSelectV2DataCreativeWordContentType) IsValid() bool {
	for _, existing := range AllowedToolsCreativeWordSelectV2DataCreativeWordContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_creative_word_select_v2_data_creative_word_content_type value
func (v ToolsCreativeWordSelectV2DataCreativeWordContentType) Ptr() *ToolsCreativeWordSelectV2DataCreativeWordContentType {
	return &v
}
