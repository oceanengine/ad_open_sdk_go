/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdUpdateV10RequestMultiProductCreativeListInnerCreativeCard
type QianchuanUniAwemeAdUpdateV10RequestMultiProductCreativeListInnerCreativeCard struct {
	PromotionCardActionButton *QianchuanUniAwemeAdUpdateV10MultiProductCreativeListCreativeCardPromotionCardActionButton `json:"promotion_card_action_button,omitempty"`
	//
	PromotionCardButtonSmartOptimization *bool `json:"promotion_card_button_smart_optimization,omitempty"`
	//
	PromotionCardImageId *string `json:"promotion_card_image_id,omitempty"`
	//
	PromotionCardSellingPoints []string `json:"promotion_card_selling_points,omitempty"`
	//
	PromotionCardTitle *string `json:"promotion_card_title,omitempty"`
}
