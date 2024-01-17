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

// ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType
type ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType string

// List of tools_site_template_site_create_v2_bricks_text_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType    ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType       ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_text_link_dto_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksTextLinkDtoLinkType {
	return &v
}
