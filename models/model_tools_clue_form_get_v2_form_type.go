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

// ToolsClueFormGetV2FormType
type ToolsClueFormGetV2FormType string

// List of tools_clue_form_get_v2_form_type
const (
	ADVANCED_CREATIVE_FORM_ToolsClueFormGetV2FormType ToolsClueFormGetV2FormType = "ADVANCED_CREATIVE_FORM"
	NATIVE_FORM_ToolsClueFormGetV2FormType            ToolsClueFormGetV2FormType = "NATIVE_FORM"
	NORMAL_FORM_ToolsClueFormGetV2FormType            ToolsClueFormGetV2FormType = "NORMAL_FORM"
)

// All allowed values of ToolsClueFormGetV2FormType enum
var AllowedToolsClueFormGetV2FormTypeEnumValues = []ToolsClueFormGetV2FormType{
	"ADVANCED_CREATIVE_FORM",
	"NATIVE_FORM",
	"NORMAL_FORM",
}

// NewToolsClueFormGetV2FormTypeFromValue returns a pointer to a valid ToolsClueFormGetV2FormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormGetV2FormTypeFromValue(v string) (*ToolsClueFormGetV2FormType, error) {
	ev := ToolsClueFormGetV2FormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormGetV2FormType: valid values are %v", v, AllowedToolsClueFormGetV2FormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormGetV2FormType) IsValid() bool {
	for _, existing := range AllowedToolsClueFormGetV2FormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_get_v2_form_type value
func (v ToolsClueFormGetV2FormType) Ptr() *ToolsClueFormGetV2FormType {
	return &v
}
