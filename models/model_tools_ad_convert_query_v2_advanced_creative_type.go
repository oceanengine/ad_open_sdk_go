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

// ToolsAdConvertQueryV2AdvancedCreativeType
type ToolsAdConvertQueryV2AdvancedCreativeType string

// List of tools_ad_convert_query_v2_advanced_creative_type
const (
	ATTACHED_CREATIVE_CONSULTANT_ToolsAdConvertQueryV2AdvancedCreativeType      ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_CONSULTANT"
	ATTACHED_CREATIVE_VOTE_INTERACT_ToolsAdConvertQueryV2AdvancedCreativeType   ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_VOTE_INTERACT"
	ATTACHED_CREATIVE_GAME_SUBSCRIBE_ToolsAdConvertQueryV2AdvancedCreativeType  ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_SUBSCRIBE"
	ATTACHED_CREATIVE_NONE_ToolsAdConvertQueryV2AdvancedCreativeType            ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_NONE"
	ATTACHED_CREATIVE_CARD_ToolsAdConvertQueryV2AdvancedCreativeType            ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_CARD"
	ATTACHED_CREATIVE_PHONE_ToolsAdConvertQueryV2AdvancedCreativeType           ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_PHONE"
	ATTACHED_CREATIVE_COUPON_INTERACT_ToolsAdConvertQueryV2AdvancedCreativeType ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON_INTERACT"
	ATTACHED_CREATIVE_APP_ToolsAdConvertQueryV2AdvancedCreativeType             ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_APP"
	ATTACHED_CREATIVE_COUPON_ToolsAdConvertQueryV2AdvancedCreativeType          ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_COUPON"
	ATTACHED_CREATIVE_GAME_PACKAGE_ToolsAdConvertQueryV2AdvancedCreativeType    ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_PACKAGE"
	ATTACHED_CREATIVE_GAME_FORM_ToolsAdConvertQueryV2AdvancedCreativeType       ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_GAME_FORM"
	ATTACHED_CREATIVE_DOWNLOAD_CARD_ToolsAdConvertQueryV2AdvancedCreativeType   ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_DOWNLOAD_CARD"
	ATTACHED_CREATIVE_COMMERCE_CARD_ToolsAdConvertQueryV2AdvancedCreativeType   ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_COMMERCE_CARD"
	ATTACHED_CREATIVE_LIVE_CARD_ToolsAdConvertQueryV2AdvancedCreativeType       ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_LIVE_CARD"
	ATTACHED_CREATIVE_SMART_PHONE_ToolsAdConvertQueryV2AdvancedCreativeType     ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_SMART_PHONE"
	ATTACHED_CREATIVE_INTERACT_ToolsAdConvertQueryV2AdvancedCreativeType        ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_INTERACT"
	ATTACHED_CREATIVE_FORM_ToolsAdConvertQueryV2AdvancedCreativeType            ToolsAdConvertQueryV2AdvancedCreativeType = "ATTACHED_CREATIVE_FORM"
)

// All allowed values of ToolsAdConvertQueryV2AdvancedCreativeType enum
var AllowedToolsAdConvertQueryV2AdvancedCreativeTypeEnumValues = []ToolsAdConvertQueryV2AdvancedCreativeType{
	"ATTACHED_CREATIVE_CONSULTANT",
	"ATTACHED_CREATIVE_VOTE_INTERACT",
	"ATTACHED_CREATIVE_GAME_SUBSCRIBE",
	"ATTACHED_CREATIVE_NONE",
	"ATTACHED_CREATIVE_CARD",
	"ATTACHED_CREATIVE_PHONE",
	"ATTACHED_CREATIVE_COUPON_INTERACT",
	"ATTACHED_CREATIVE_APP",
	"ATTACHED_CREATIVE_COUPON",
	"ATTACHED_CREATIVE_GAME_PACKAGE",
	"ATTACHED_CREATIVE_GAME_FORM",
	"ATTACHED_CREATIVE_DOWNLOAD_CARD",
	"ATTACHED_CREATIVE_COMMERCE_CARD",
	"ATTACHED_CREATIVE_LIVE_CARD",
	"ATTACHED_CREATIVE_SMART_PHONE",
	"ATTACHED_CREATIVE_INTERACT",
	"ATTACHED_CREATIVE_FORM",
}

// NewToolsAdConvertQueryV2AdvancedCreativeTypeFromValue returns a pointer to a valid ToolsAdConvertQueryV2AdvancedCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2AdvancedCreativeTypeFromValue(v string) (*ToolsAdConvertQueryV2AdvancedCreativeType, error) {
	ev := ToolsAdConvertQueryV2AdvancedCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2AdvancedCreativeType: valid values are %v", v, AllowedToolsAdConvertQueryV2AdvancedCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2AdvancedCreativeType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2AdvancedCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_advanced_creative_type value
func (v ToolsAdConvertQueryV2AdvancedCreativeType) Ptr() *ToolsAdConvertQueryV2AdvancedCreativeType {
	return &v
}
