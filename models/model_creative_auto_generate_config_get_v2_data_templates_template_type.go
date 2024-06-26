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

// CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType
type CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType string

// List of creative_auto_generate_config_get_v2_data_templates_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType  CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType enum
var AllowedCreativeAutoGenerateConfigGetV2DataTemplatesTemplateTypeEnumValues = []CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewCreativeAutoGenerateConfigGetV2DataTemplatesTemplateTypeFromValue returns a pointer to a valid CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigGetV2DataTemplatesTemplateTypeFromValue(v string) (*CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType, error) {
	ev := CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType: valid values are %v", v, AllowedCreativeAutoGenerateConfigGetV2DataTemplatesTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigGetV2DataTemplatesTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_get_v2_data_templates_template_type value
func (v CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType) Ptr() *CreativeAutoGenerateConfigGetV2DataTemplatesTemplateType {
	return &v
}
