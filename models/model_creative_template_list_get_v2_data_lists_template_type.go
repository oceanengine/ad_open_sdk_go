/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeTemplateListGetV2DataListsTemplateType
type CreativeTemplateListGetV2DataListsTemplateType string

// List of creative_template_list_get_v2_data_lists_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_CreativeTemplateListGetV2DataListsTemplateType CreativeTemplateListGetV2DataListsTemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_CreativeTemplateListGetV2DataListsTemplateType CreativeTemplateListGetV2DataListsTemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_CreativeTemplateListGetV2DataListsTemplateType  CreativeTemplateListGetV2DataListsTemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of CreativeTemplateListGetV2DataListsTemplateType enum
var AllowedCreativeTemplateListGetV2DataListsTemplateTypeEnumValues = []CreativeTemplateListGetV2DataListsTemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewCreativeTemplateListGetV2DataListsTemplateTypeFromValue returns a pointer to a valid CreativeTemplateListGetV2DataListsTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeTemplateListGetV2DataListsTemplateTypeFromValue(v string) (*CreativeTemplateListGetV2DataListsTemplateType, error) {
	ev := CreativeTemplateListGetV2DataListsTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeTemplateListGetV2DataListsTemplateType: valid values are %v", v, AllowedCreativeTemplateListGetV2DataListsTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeTemplateListGetV2DataListsTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeTemplateListGetV2DataListsTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_template_list_get_v2_data_lists_template_type value
func (v CreativeTemplateListGetV2DataListsTemplateType) Ptr() *CreativeTemplateListGetV2DataListsTemplateType {
	return &v
}
