/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType
type CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType string

// List of creative_detail_get_v3.0_data_creative_list_video_material_dpa_video_template_type
const (
	DPA_VIDEO_TEMPLATE_CUSTOM_CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType     CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_DEPRECATED_CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_SMART_CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType      CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType enum
var AllowedCreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateTypeEnumValues = []CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewCreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateTypeFromValue returns a pointer to a valid CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateTypeFromValue(v string) (*CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType, error) {
	ev := CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType: valid values are %v", v, AllowedCreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_creative_list_video_material_dpa_video_template_type value
func (v CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType) Ptr() *CreativeDetailGetV30DataCreativeListVideoMaterialDpaVideoTemplateType {
	return &v
}
