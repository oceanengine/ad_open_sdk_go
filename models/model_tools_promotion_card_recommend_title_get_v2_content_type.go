/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionCardRecommendTitleGetV2ContentType
type ToolsPromotionCardRecommendTitleGetV2ContentType string

// List of tools_promotion_card_recommend_title_get_v2_content_type
const (
	GAME_SUBSCRIBE_ToolsPromotionCardRecommendTitleGetV2ContentType ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_SUBSCRIBE"
	GAME_FORM_ToolsPromotionCardRecommendTitleGetV2ContentType      ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_FORM"
	CONSULT_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "CONSULT"
	CARD_ToolsPromotionCardRecommendTitleGetV2ContentType           ToolsPromotionCardRecommendTitleGetV2ContentType = "CARD"
	DOWNLOAD_ToolsPromotionCardRecommendTitleGetV2ContentType       ToolsPromotionCardRecommendTitleGetV2ContentType = "DOWNLOAD"
	PHONE_ToolsPromotionCardRecommendTitleGetV2ContentType          ToolsPromotionCardRecommendTitleGetV2ContentType = "PHONE"
	GAME_PACKAGE_ToolsPromotionCardRecommendTitleGetV2ContentType   ToolsPromotionCardRecommendTitleGetV2ContentType = "GAME_PACKAGE"
	UNKNOWN_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "UNKNOWN"
	TABLE_ToolsPromotionCardRecommendTitleGetV2ContentType          ToolsPromotionCardRecommendTitleGetV2ContentType = "TABLE"
	LANDING_ToolsPromotionCardRecommendTitleGetV2ContentType        ToolsPromotionCardRecommendTitleGetV2ContentType = "LANDING"
)

// All allowed values of ToolsPromotionCardRecommendTitleGetV2ContentType enum
var AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues = []ToolsPromotionCardRecommendTitleGetV2ContentType{
	"GAME_SUBSCRIBE",
	"GAME_FORM",
	"CONSULT",
	"CARD",
	"DOWNLOAD",
	"PHONE",
	"GAME_PACKAGE",
	"UNKNOWN",
	"TABLE",
	"LANDING",
}

// NewToolsPromotionCardRecommendTitleGetV2ContentTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendTitleGetV2ContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendTitleGetV2ContentTypeFromValue(v string) (*ToolsPromotionCardRecommendTitleGetV2ContentType, error) {
	ev := ToolsPromotionCardRecommendTitleGetV2ContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendTitleGetV2ContentType: valid values are %v", v, AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendTitleGetV2ContentType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendTitleGetV2ContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_title_get_v2_content_type value
func (v ToolsPromotionCardRecommendTitleGetV2ContentType) Ptr() *ToolsPromotionCardRecommendTitleGetV2ContentType {
	return &v
}
