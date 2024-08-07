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

// ToolsCountryInfoV2Language
type ToolsCountryInfoV2Language string

// List of tools_country_info_v2_language
const (
	ZH_CN_ToolsCountryInfoV2Language     ToolsCountryInfoV2Language = "ZH_CN"
	ZH_CN_GOV_ToolsCountryInfoV2Language ToolsCountryInfoV2Language = "ZH_CN_GOV"
)

// All allowed values of ToolsCountryInfoV2Language enum
var AllowedToolsCountryInfoV2LanguageEnumValues = []ToolsCountryInfoV2Language{
	"ZH_CN",
	"ZH_CN_GOV",
}

// NewToolsCountryInfoV2LanguageFromValue returns a pointer to a valid ToolsCountryInfoV2Language
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsCountryInfoV2LanguageFromValue(v string) (*ToolsCountryInfoV2Language, error) {
	ev := ToolsCountryInfoV2Language(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsCountryInfoV2Language: valid values are %v", v, AllowedToolsCountryInfoV2LanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsCountryInfoV2Language) IsValid() bool {
	for _, existing := range AllowedToolsCountryInfoV2LanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_country_info_v2_language value
func (v ToolsCountryInfoV2Language) Ptr() *ToolsCountryInfoV2Language {
	return &v
}
