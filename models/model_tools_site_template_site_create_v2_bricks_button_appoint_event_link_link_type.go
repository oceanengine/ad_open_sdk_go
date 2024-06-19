/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType
type ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType string

// List of tools_site_template_site_create_v2_bricks_button_appoint_event_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType    ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType       ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_button_appoint_event_link_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksButtonAppointEventLinkLinkType {
	return &v
}
