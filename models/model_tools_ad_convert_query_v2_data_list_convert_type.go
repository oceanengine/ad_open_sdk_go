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

// ToolsAdConvertQueryV2DataListConvertType
type ToolsAdConvertQueryV2DataListConvertType string

// List of tools_ad_convert_query_v2_data_list_convert_type
const (
	AD_CONVERT_SOURCE_TYPE_SDK_ToolsAdConvertQueryV2DataListConvertType                 ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_SDK"
	AD_CONVERT_SOURCE_TYPE_XPATH_ToolsAdConvertQueryV2DataListConvertType               ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_XPATH"
	AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT_ToolsAdConvertQueryV2DataListConvertType       ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT"
	AD_CONVERT_TYPE_MICRO_APP_SDK_ToolsAdConvertQueryV2DataListConvertType              ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_TYPE_MICRO_APP_SDK"
	AD_CONVERT_TYPE_NATIVE_PROMOTION_ToolsAdConvertQueryV2DataListConvertType           ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_TYPE_NATIVE_PROMOTION"
	AD_CONVERT_SOURCE_TYPE_JS_ToolsAdConvertQueryV2DataListConvertType                  ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_JS"
	AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD_ToolsAdConvertQueryV2DataListConvertType ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD"
	AD_CONVERT_SOURCE_TYPE_INTERNAL_ToolsAdConvertQueryV2DataListConvertType            ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_INTERNAL"
	AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI_ToolsAdConvertQueryV2DataListConvertType       ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI"
	AD_CONVERT_TYPE_MICRO_APP_API_ToolsAdConvertQueryV2DataListConvertType              ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_TYPE_MICRO_APP_API"
	AD_CONVERT_SOURCE_TYPE_H5_API_ToolsAdConvertQueryV2DataListConvertType              ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_H5_API"
	AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD_ToolsAdConvertQueryV2DataListConvertType        ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD"
	AD_CONVERT_SOURCE_TYPE_API_ToolsAdConvertQueryV2DataListConvertType                 ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_API"
	AD_CONVERT_SOURCE_TYPE_CONFIG_ToolsAdConvertQueryV2DataListConvertType              ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_CONFIG"
	AD_CONVERT_SOURCE_TYPE_CPS_GAME_ToolsAdConvertQueryV2DataListConvertType            ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_CPS_GAME"
	AD_CONVERT_SOURCE_TYPE_OPEN_URL_ToolsAdConvertQueryV2DataListConvertType            ToolsAdConvertQueryV2DataListConvertType = "AD_CONVERT_SOURCE_TYPE_OPEN_URL"
)

// All allowed values of ToolsAdConvertQueryV2DataListConvertType enum
var AllowedToolsAdConvertQueryV2DataListConvertTypeEnumValues = []ToolsAdConvertQueryV2DataListConvertType{
	"AD_CONVERT_SOURCE_TYPE_SDK",
	"AD_CONVERT_SOURCE_TYPE_XPATH",
	"AD_CONVERT_SOURCE_TYPE_MULTI_CONVERT",
	"AD_CONVERT_TYPE_MICRO_APP_SDK",
	"AD_CONVERT_TYPE_NATIVE_PROMOTION",
	"AD_CONVERT_SOURCE_TYPE_JS",
	"AD_CONVERT_SOURCE_TYPE_NORMAL_APP_DOWNLOAD",
	"AD_CONVERT_SOURCE_TYPE_INTERNAL",
	"AD_CONVERT_SOURCE_TYPE_APP_API_TEMAI",
	"AD_CONVERT_TYPE_MICRO_APP_API",
	"AD_CONVERT_SOURCE_TYPE_H5_API",
	"AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD",
	"AD_CONVERT_SOURCE_TYPE_API",
	"AD_CONVERT_SOURCE_TYPE_CONFIG",
	"AD_CONVERT_SOURCE_TYPE_CPS_GAME",
	"AD_CONVERT_SOURCE_TYPE_OPEN_URL",
}

// NewToolsAdConvertQueryV2DataListConvertTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2DataListConvertType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2DataListConvertTypeFromValue(v string) (*ToolsAdConvertQueryV2DataListConvertType, error) {
	ev := ToolsAdConvertQueryV2DataListConvertType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2DataListConvertType: valid values are %v", v, AllowedToolsAdConvertQueryV2DataListConvertTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2DataListConvertType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2DataListConvertTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_data_list_convert_type value
func (v ToolsAdConvertQueryV2DataListConvertType) Ptr() *ToolsAdConvertQueryV2DataListConvertType {
	return &v
}
