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

// ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType
type ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType string

// List of tools_site_template_site_create_v2_bricks_picture_group_content_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType    ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType       ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_picture_group_content_link_dto_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksPictureGroupContentLinkDtoLinkType {
	return &v
}
