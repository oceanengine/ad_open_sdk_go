/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsActionTextGetV2AdvancedCreativeType
type ToolsActionTextGetV2AdvancedCreativeType string

// List of tools_action_text_get_v2_advanced_creative_type
const (
	ATTACHED_CREATIVE_VOTE_INTERACT_ToolsActionTextGetV2AdvancedCreativeType   ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_VOTE_INTERACT"
	ATTACHED_CREATIVE_CARD_ToolsActionTextGetV2AdvancedCreativeType            ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_CARD"
	ATTACHED_CREATIVE_GAME_SUBSCRIBE_ToolsActionTextGetV2AdvancedCreativeType  ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_SUBSCRIBE"
	ATTACHED_CREATIVE_PHONE_ToolsActionTextGetV2AdvancedCreativeType           ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_PHONE"
	ATTACHED_CREATIVE_COUPON_ToolsActionTextGetV2AdvancedCreativeType          ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON"
	ATTACHED_CREATIVE_COMMERCE_CARD_ToolsActionTextGetV2AdvancedCreativeType   ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COMMERCE_CARD"
	ATTACHED_CREATIVE_FORM_ToolsActionTextGetV2AdvancedCreativeType            ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_FORM"
	ATTACHED_CREATIVE_INTERACT_ToolsActionTextGetV2AdvancedCreativeType        ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_INTERACT"
	ATTACHED_CREATIVE_GAME_FORM_ToolsActionTextGetV2AdvancedCreativeType       ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_FORM"
	ATTACHED_CREATIVE_SMART_PHONE_ToolsActionTextGetV2AdvancedCreativeType     ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_SMART_PHONE"
	ATTACHED_CREATIVE_DOWNLOAD_CARD_ToolsActionTextGetV2AdvancedCreativeType   ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_DOWNLOAD_CARD"
	ATTACHED_CREATIVE_GAME_PACKAGE_ToolsActionTextGetV2AdvancedCreativeType    ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_PACKAGE"
	ATTACHED_CREATIVE_LIVE_CARD_ToolsActionTextGetV2AdvancedCreativeType       ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_LIVE_CARD"
	ATTACHED_CREATIVE_CONSULTANT_ToolsActionTextGetV2AdvancedCreativeType      ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_CONSULTANT"
	ATTACHED_CREATIVE_NONE_ToolsActionTextGetV2AdvancedCreativeType            ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_NONE"
	ATTACHED_CREATIVE_APP_ToolsActionTextGetV2AdvancedCreativeType             ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_APP"
	ATTACHED_CREATIVE_COUPON_INTERACT_ToolsActionTextGetV2AdvancedCreativeType ToolsActionTextGetV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON_INTERACT"
)

// All allowed values of ToolsActionTextGetV2AdvancedCreativeType enum
var AllowedToolsActionTextGetV2AdvancedCreativeTypeEnumValues = []ToolsActionTextGetV2AdvancedCreativeType{
	"ATTACHED_CREATIVE_VOTE_INTERACT",
	"ATTACHED_CREATIVE_CARD",
	"ATTACHED_CREATIVE_GAME_SUBSCRIBE",
	"ATTACHED_CREATIVE_PHONE",
	"ATTACHED_CREATIVE_COUPON",
	"ATTACHED_CREATIVE_COMMERCE_CARD",
	"ATTACHED_CREATIVE_FORM",
	"ATTACHED_CREATIVE_INTERACT",
	"ATTACHED_CREATIVE_GAME_FORM",
	"ATTACHED_CREATIVE_SMART_PHONE",
	"ATTACHED_CREATIVE_DOWNLOAD_CARD",
	"ATTACHED_CREATIVE_GAME_PACKAGE",
	"ATTACHED_CREATIVE_LIVE_CARD",
	"ATTACHED_CREATIVE_CONSULTANT",
	"ATTACHED_CREATIVE_NONE",
	"ATTACHED_CREATIVE_APP",
	"ATTACHED_CREATIVE_COUPON_INTERACT",
}

// NewToolsActionTextGetV2AdvancedCreativeTypeFromValue returns a pointer to a valid ToolsActionTextGetV2AdvancedCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsActionTextGetV2AdvancedCreativeTypeFromValue(v string) (*ToolsActionTextGetV2AdvancedCreativeType, error) {
	ev := ToolsActionTextGetV2AdvancedCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsActionTextGetV2AdvancedCreativeType: valid values are %v", v, AllowedToolsActionTextGetV2AdvancedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsActionTextGetV2AdvancedCreativeType) IsValid() bool {
	for _, existing := range AllowedToolsActionTextGetV2AdvancedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_action_text_get_v2_advanced_creative_type value
func (v ToolsActionTextGetV2AdvancedCreativeType) Ptr() *ToolsActionTextGetV2AdvancedCreativeType {
	return &v
}