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

// ToolsSiteTemplateGetV2DataListBricksButtonEventType
type ToolsSiteTemplateGetV2DataListBricksButtonEventType string

// List of tools_site_template_get_v2_data_list_bricks_button_event_type
const (
	APPOINT_EVENT_ToolsSiteTemplateGetV2DataListBricksButtonEventType   ToolsSiteTemplateGetV2DataListBricksButtonEventType = "APPOINT_EVENT"
	DOWNLOAD_EVENT_ToolsSiteTemplateGetV2DataListBricksButtonEventType  ToolsSiteTemplateGetV2DataListBricksButtonEventType = "DOWNLOAD_EVENT"
	LINK_EVENT_ToolsSiteTemplateGetV2DataListBricksButtonEventType      ToolsSiteTemplateGetV2DataListBricksButtonEventType = "LINK_EVENT"
	TELEPHONE_EVENT_ToolsSiteTemplateGetV2DataListBricksButtonEventType ToolsSiteTemplateGetV2DataListBricksButtonEventType = "TELEPHONE_EVENT"
)

// All allowed values of ToolsSiteTemplateGetV2DataListBricksButtonEventType enum
var AllowedToolsSiteTemplateGetV2DataListBricksButtonEventTypeEnumValues = []ToolsSiteTemplateGetV2DataListBricksButtonEventType{
	"APPOINT_EVENT",
	"DOWNLOAD_EVENT",
	"LINK_EVENT",
	"TELEPHONE_EVENT",
}

// NewToolsSiteTemplateGetV2DataListBricksButtonEventTypeFromValue returns a pointer to a valid ToolsSiteTemplateGetV2DataListBricksButtonEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateGetV2DataListBricksButtonEventTypeFromValue(v string) (*ToolsSiteTemplateGetV2DataListBricksButtonEventType, error) {
	ev := ToolsSiteTemplateGetV2DataListBricksButtonEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateGetV2DataListBricksButtonEventType: valid values are %v", v, AllowedToolsSiteTemplateGetV2DataListBricksButtonEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateGetV2DataListBricksButtonEventType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateGetV2DataListBricksButtonEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_get_v2_data_list_bricks_button_event_type value
func (v ToolsSiteTemplateGetV2DataListBricksButtonEventType) Ptr() *ToolsSiteTemplateGetV2DataListBricksButtonEventType {
	return &v
}
