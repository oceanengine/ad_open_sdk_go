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

// ToolsClueFormGetV2DataListFormType
type ToolsClueFormGetV2DataListFormType string

// List of tools_clue_form_get_v2_data_list_form_type
const (
	ADVANCED_CREATIVE_FORM_ToolsClueFormGetV2DataListFormType ToolsClueFormGetV2DataListFormType = "ADVANCED_CREATIVE_FORM"
	NATIVE_FORM_ToolsClueFormGetV2DataListFormType            ToolsClueFormGetV2DataListFormType = "NATIVE_FORM"
	NORMAL_FORM_ToolsClueFormGetV2DataListFormType            ToolsClueFormGetV2DataListFormType = "NORMAL_FORM"
)

// All allowed values of ToolsClueFormGetV2DataListFormType enum
var AllowedToolsClueFormGetV2DataListFormTypeEnumValues = []ToolsClueFormGetV2DataListFormType{
	"ADVANCED_CREATIVE_FORM",
	"NATIVE_FORM",
	"NORMAL_FORM",
}

// NewToolsClueFormGetV2DataListFormTypeFromValue returns a pointer to a valid ToolsClueFormGetV2DataListFormType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsClueFormGetV2DataListFormTypeFromValue(v string) (*ToolsClueFormGetV2DataListFormType, error) {
	ev := ToolsClueFormGetV2DataListFormType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsClueFormGetV2DataListFormType: valid values are %v", v, AllowedToolsClueFormGetV2DataListFormTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsClueFormGetV2DataListFormType) IsValid() bool {
	for _, existing := range AllowedToolsClueFormGetV2DataListFormTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_clue_form_get_v2_data_list_form_type value
func (v ToolsClueFormGetV2DataListFormType) Ptr() *ToolsClueFormGetV2DataListFormType {
	return &v
}
