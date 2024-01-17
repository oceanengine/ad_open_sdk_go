/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2DataListExternalActionsBelong
type ToolsAdConvertQueryV2DataListExternalActionsBelong string

// List of tools_ad_convert_query_v2_data_list_external_actions_belong
const (
	BELONG_ADVANCED_CREATIVE_ToolsAdConvertQueryV2DataListExternalActionsBelong ToolsAdConvertQueryV2DataListExternalActionsBelong = "BELONG_ADVANCED_CREATIVE"
	BELONG_MICRO_APP_ToolsAdConvertQueryV2DataListExternalActionsBelong         ToolsAdConvertQueryV2DataListExternalActionsBelong = "BELONG_MICRO_APP"
	BELONG_EXTERNAL_URL_ToolsAdConvertQueryV2DataListExternalActionsBelong      ToolsAdConvertQueryV2DataListExternalActionsBelong = "BELONG_EXTERNAL_URL"
)

// All allowed values of ToolsAdConvertQueryV2DataListExternalActionsBelong enum
var AllowedToolsAdConvertQueryV2DataListExternalActionsBelongEnumValues = []ToolsAdConvertQueryV2DataListExternalActionsBelong{
	"BELONG_ADVANCED_CREATIVE",
	"BELONG_MICRO_APP",
	"BELONG_EXTERNAL_URL",
}

// NewToolsAdConvertQueryV2DataListExternalActionsBelongFromValue returns a pointer to a valid ToolsAdConvertQueryV2DataListExternalActionsBelong
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DataListExternalActionsBelongFromValue(v string) (*ToolsAdConvertQueryV2DataListExternalActionsBelong, error) {
	ev := ToolsAdConvertQueryV2DataListExternalActionsBelong(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DataListExternalActionsBelong: valid values are %v", v, AllowedToolsAdConvertQueryV2DataListExternalActionsBelongEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DataListExternalActionsBelong) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DataListExternalActionsBelongEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_data_list_external_actions_belong value
func (v ToolsAdConvertQueryV2DataListExternalActionsBelong) Ptr() *ToolsAdConvertQueryV2DataListExternalActionsBelong {
	return &v
}
