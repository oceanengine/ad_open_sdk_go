/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType
type CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType string

// List of creative_custom_creative_create_v2_creative_list_video_material_dpa_video_template_type
const (
	DPA_VIDEO_TEMPLATE_DEPRECATED_CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_CUSTOM_CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType     CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_SMART_CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType      CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
)

// All allowed values of CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType enum
var AllowedCreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateTypeEnumValues = []CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_SMART",
}

// NewCreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateTypeFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateTypeFromValue(v string) (*CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType, error) {
	ev := CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_creative_list_video_material_dpa_video_template_type value
func (v CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType) Ptr() *CreativeCustomCreativeCreateV2CreativeListVideoMaterialDpaVideoTemplateType {
	return &v
}
