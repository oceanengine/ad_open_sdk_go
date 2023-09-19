/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType
type ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType string

// List of tools_site_template_create_v2_data_bricks_picture_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType    ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType       ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType enum
var AllowedToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_create_v2_data_bricks_picture_link_dto_link_type value
func (v ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType) Ptr() *ToolsSiteTemplateCreateV2DataBricksPictureLinkDtoLinkType {
	return &v
}
