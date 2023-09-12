/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType
type ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType string

// List of tools_site_template_get_v2_data_list_bricks_picture_group_content_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType    ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType       ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_picture_group_content_link_dto_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksPictureGroupContentLinkDtoLinkType {
	return &v
}
