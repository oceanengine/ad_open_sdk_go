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

// CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType
type CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType string

// List of creative_auto_generate_config_create_v2_data_templates_template_type
const (
	HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType = "HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO"
	VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO_CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType  CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType = "VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO"
)

// All allowed values of CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType enum
var AllowedCreativeAutoGenerateConfigCreateV2DataTemplatesTemplateTypeEnumValues = []CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType{
	"HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO",
	"VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO",
}

// NewCreativeAutoGenerateConfigCreateV2DataTemplatesTemplateTypeFromValue returns a pointer to a valid CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeAutoGenerateConfigCreateV2DataTemplatesTemplateTypeFromValue(v string) (*CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType, error) {
	ev := CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType: valid values are %v", v, AllowedCreativeAutoGenerateConfigCreateV2DataTemplatesTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeAutoGenerateConfigCreateV2DataTemplatesTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_auto_generate_config_create_v2_data_templates_template_type value
func (v CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType) Ptr() *CreativeAutoGenerateConfigCreateV2DataTemplatesTemplateType {
	return &v
}
