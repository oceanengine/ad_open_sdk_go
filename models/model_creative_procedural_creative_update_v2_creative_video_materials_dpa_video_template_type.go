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

// CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType
type CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType string

// List of creative_procedural_creative_update_v2_creative_video_materials_dpa_video_template_type
const (
	DPA_VIDEO_TEMPLATE_SMART_CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType      CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_SMART"
	DPA_VIDEO_TEMPLATE_CUSTOM_CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType     CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_CUSTOM"
	DPA_VIDEO_TEMPLATE_DEPRECATED_CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType = "DPA_VIDEO_TEMPLATE_DEPRECATED"
)

// All allowed values of CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType enum
var AllowedCreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues = []CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType{
	"DPA_VIDEO_TEMPLATE_SMART",
	"DPA_VIDEO_TEMPLATE_CUSTOM",
	"DPA_VIDEO_TEMPLATE_DEPRECATED",
}

// NewCreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateTypeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateTypeFromValue(v string) (*CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType, error) {
	ev := CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_creative_video_materials_dpa_video_template_type value
func (v CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType) Ptr() *CreativeProceduralCreativeUpdateV2CreativeVideoMaterialsDpaVideoTemplateType {
	return &v
}
