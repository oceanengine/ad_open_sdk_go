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

// ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType
type ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType string

// List of tools_site_template_site_create_v2_bricks_button_download_event_android_link_link_type
const (
	QUICK_APP_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType = "QUICK_APP"
	SCHEME_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType    ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType = "SCHEME"
	URL_ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType       ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType = "URL"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType{
	"QUICK_APP",
	"SCHEME",
	"URL",
}

// NewToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_button_download_event_android_link_link_type value
func (v ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksButtonDownloadEventAndroidLinkLinkType {
	return &v
}
