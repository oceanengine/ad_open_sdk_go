/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdCreateV10RequestMultiProductCreativeListInnerCreativeCard
type QianchuanUniAwemeAdCreateV10RequestMultiProductCreativeListInnerCreativeCard struct {
	PromotionCardActionButton *QianchuanUniAwemeAdCreateV10MultiProductCreativeListCreativeCardPromotionCardActionButton `json:"promotion_card_action_button,omitempty"`
	//
	PromotionCardButtonSmartOptimization *bool `json:"promotion_card_button_smart_optimization,omitempty"`
	//
	PromotionCardImageId string `json:"promotion_card_image_id"`
	//
	PromotionCardSellingPoints []string `json:"promotion_card_selling_points"`
	//
	PromotionCardTitle *string `json:"promotion_card_title,omitempty"`
}
