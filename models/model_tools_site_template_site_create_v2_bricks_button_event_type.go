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

// ToolsSiteTemplateSiteCreateV2BricksButtonEventType
type ToolsSiteTemplateSiteCreateV2BricksButtonEventType string

// List of tools_site_template_site_create_v2_bricks_button_event_type
const (
	APPOINT_EVENT_ToolsSiteTemplateSiteCreateV2BricksButtonEventType   ToolsSiteTemplateSiteCreateV2BricksButtonEventType = "APPOINT_EVENT"
	DOWNLOAD_EVENT_ToolsSiteTemplateSiteCreateV2BricksButtonEventType  ToolsSiteTemplateSiteCreateV2BricksButtonEventType = "DOWNLOAD_EVENT"
	LINK_EVENT_ToolsSiteTemplateSiteCreateV2BricksButtonEventType      ToolsSiteTemplateSiteCreateV2BricksButtonEventType = "LINK_EVENT"
	TELEPHONE_EVENT_ToolsSiteTemplateSiteCreateV2BricksButtonEventType ToolsSiteTemplateSiteCreateV2BricksButtonEventType = "TELEPHONE_EVENT"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksButtonEventType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksButtonEventTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksButtonEventType{
	"APPOINT_EVENT",
	"DOWNLOAD_EVENT",
	"LINK_EVENT",
	"TELEPHONE_EVENT",
}

// NewToolsSiteTemplateSiteCreateV2BricksButtonEventTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksButtonEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksButtonEventTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksButtonEventType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksButtonEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksButtonEventType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksButtonEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksButtonEventType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksButtonEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_button_event_type value
func (v ToolsSiteTemplateSiteCreateV2BricksButtonEventType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksButtonEventType {
	return &v
}
