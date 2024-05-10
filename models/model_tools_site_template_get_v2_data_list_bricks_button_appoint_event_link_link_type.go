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

// ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType
type ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType string

// List of tools_site_template_get_v2_data_list_bricks_button_appoint_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType    ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType       ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType enum
var AllowedToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_button_appoint_event_link_link_type value
func (v ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType) Ptr() *ToolsSiteTemplateGetV2DataListBricksButtonAppointEventLinkLinkType {
	return &v
}
