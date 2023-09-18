/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsSiteTemplateSiteCreateV2BricksType
type ToolsSiteTemplateSiteCreateV2BricksType string

// List of tools_site_template_site_create_v2_bricks_type
const (
	BUTTON_ToolsSiteTemplateSiteCreateV2BricksType        ToolsSiteTemplateSiteCreateV2BricksType = "BUTTON"
	COUPON_ToolsSiteTemplateSiteCreateV2BricksType        ToolsSiteTemplateSiteCreateV2BricksType = "COUPON"
	FORM_ToolsSiteTemplateSiteCreateV2BricksType          ToolsSiteTemplateSiteCreateV2BricksType = "FORM"
	PICTURE_ToolsSiteTemplateSiteCreateV2BricksType       ToolsSiteTemplateSiteCreateV2BricksType = "PICTURE"
	PICTURE_GROUP_ToolsSiteTemplateSiteCreateV2BricksType ToolsSiteTemplateSiteCreateV2BricksType = "PICTURE_GROUP"
	RICH_TEXT_ToolsSiteTemplateSiteCreateV2BricksType     ToolsSiteTemplateSiteCreateV2BricksType = "RICH_TEXT"
	SIMPLE_TEXT_ToolsSiteTemplateSiteCreateV2BricksType   ToolsSiteTemplateSiteCreateV2BricksType = "SIMPLE_TEXT"
	VIDEO_ToolsSiteTemplateSiteCreateV2BricksType         ToolsSiteTemplateSiteCreateV2BricksType = "VIDEO"
	WECHAT_APPLET_ToolsSiteTemplateSiteCreateV2BricksType ToolsSiteTemplateSiteCreateV2BricksType = "WECHAT_APPLET"
	WECHAT_GAME_ToolsSiteTemplateSiteCreateV2BricksType   ToolsSiteTemplateSiteCreateV2BricksType = "WECHAT_GAME"
)

// All allowed values of ToolsSiteTemplateSiteCreateV2BricksType enum
var AllowedToolsSiteTemplateSiteCreateV2BricksTypeEnumValues = []ToolsSiteTemplateSiteCreateV2BricksType{
	"BUTTON",
	"COUPON",
	"FORM",
	"PICTURE",
	"PICTURE_GROUP",
	"RICH_TEXT",
	"SIMPLE_TEXT",
	"VIDEO",
	"WECHAT_APPLET",
	"WECHAT_GAME",
}

// NewToolsSiteTemplateSiteCreateV2BricksTypeFromValue returns a pointer to a valid ToolsSiteTemplateSiteCreateV2BricksType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsSiteTemplateSiteCreateV2BricksTypeFromValue(v string) (*ToolsSiteTemplateSiteCreateV2BricksType, error) {
	ev := ToolsSiteTemplateSiteCreateV2BricksType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsSiteTemplateSiteCreateV2BricksType: valid values are %v", v, AllowedToolsSiteTemplateSiteCreateV2BricksTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsSiteTemplateSiteCreateV2BricksType) IsValid() bool {
	for _, existing := range AllowedToolsSiteTemplateSiteCreateV2BricksTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_site_template_site_create_v2_bricks_type value
func (v ToolsSiteTemplateSiteCreateV2BricksType) Ptr() *ToolsSiteTemplateSiteCreateV2BricksType {
	return &v
}