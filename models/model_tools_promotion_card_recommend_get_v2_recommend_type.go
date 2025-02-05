/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionCardRecommendGetV2RecommendType
type ToolsPromotionCardRecommendGetV2RecommendType string

// List of tools_promotion_card_recommend_get_v2_recommend_type
const (
	SELLING_POINTS_ToolsPromotionCardRecommendGetV2RecommendType      ToolsPromotionCardRecommendGetV2RecommendType = "SELLING_POINTS"
	PRODUCT_DESCRIPTION_ToolsPromotionCardRecommendGetV2RecommendType ToolsPromotionCardRecommendGetV2RecommendType = "PRODUCT_DESCRIPTION"
	CALL_TO_ACTION_ToolsPromotionCardRecommendGetV2RecommendType      ToolsPromotionCardRecommendGetV2RecommendType = "CALL_TO_ACTION"
)

// Ptr returns reference to tools_promotion_card_recommend_get_v2_recommend_type value
func (v ToolsPromotionCardRecommendGetV2RecommendType) Ptr() *ToolsPromotionCardRecommendGetV2RecommendType {
	return &v
}
