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

// ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType
type ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType string

// List of tools_site_template_site_create_v2_bricks_picture_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType    ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType       ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_picture_link_dto_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksPictureLinkDtoLinkType {
	return &v
}
