/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType
type ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType string

// List of tools_site_template_get_v2_data_list_bricks_text_link_dto_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType    ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType       ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_text_link_dto_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksTextLinkDtoLinkType {
	return &v
}
