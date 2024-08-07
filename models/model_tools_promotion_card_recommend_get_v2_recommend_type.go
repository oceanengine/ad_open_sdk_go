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

// ToolsPromotionCardRecommendGetV2RecommendType
type ToolsPromotionCardRecommendGetV2RecommendType string

// List of tools_promotion_card_recommend_get_v2_recommend_type
const (
	CALL_TO_ACTION_ToolsPromotionCardRecommendGetV2RecommendType      ToolsPromotionCardRecommendGetV2RecommendType = "CALL_TO_ACTION"
	SELLING_POINTS_ToolsPromotionCardRecommendGetV2RecommendType      ToolsPromotionCardRecommendGetV2RecommendType = "SELLING_POINTS"
	PRODUCT_DESCRIPTION_ToolsPromotionCardRecommendGetV2RecommendType ToolsPromotionCardRecommendGetV2RecommendType = "PRODUCT_DESCRIPTION"
)

// All allowed values of ToolsPromotionCardRecommendGetV2RecommendType enum
var AllowedToolsPromotionCardRecommendGetV2RecommendTypeEnumValues = []ToolsPromotionCardRecommendGetV2RecommendType{
	"CALL_TO_ACTION",
	"SELLING_POINTS",
	"PRODUCT_DESCRIPTION",
}

// NewToolsPromotionCardRecommendGetV2RecommendTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendGetV2RecommendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendGetV2RecommendTypeFromValue(v string) (*ToolsPromotionCardRecommendGetV2RecommendType, error) {
	ev := ToolsPromotionCardRecommendGetV2RecommendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendGetV2RecommendType: valid values are %v", v, AllowedToolsPromotionCardRecommendGetV2RecommendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendGetV2RecommendType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendGetV2RecommendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_get_v2_recommend_type value
func (v ToolsPromotionCardRecommendGetV2RecommendType) Ptr() *ToolsPromotionCardRecommendGetV2RecommendType {
	return &v
}
