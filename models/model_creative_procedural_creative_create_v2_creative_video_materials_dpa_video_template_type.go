/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType
type CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType string

// List of creative_procedural_creative_create_v2_creative_video_materials_dpa_video_template_type
const (
	DPA_VIDEO_TEMPLATE_SMART_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType      CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
	DPA_VIDEO_TEMPLATE_DEPRECATED_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
	DPA_VIDEO_TEMPLATE_CUSTOM_CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType     CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
)

// All allowed values of CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType enum
var AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues = []CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_SMART",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
	"DPA_VIDEO_TEMPLATE_CUSTOM",
}

// NewCreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateTypeFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateTypeFromValue(v string) (*CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType, error) {
	ev := CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_creative_video_materials_dpa_video_template_type value
func (v CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType) Ptr() *CreativeProceduralCreativeCreateV2CreativeVideoMaterialsDpaVideoTemplateType {
	return &v
}
