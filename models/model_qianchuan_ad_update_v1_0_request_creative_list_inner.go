/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestCreativeListInner struct for QianchuanAdUpdateV10RequestCreativeListInner
type QianchuanAdUpdateV10RequestCreativeListInner struct {
	CarouselMaterial *QianchuanAdUpdateV10RequestCreativeListInnerCarouselMaterial `json:"carousel_material,omitempty"`
	//
	CreativeId            *int64                                                             `json:"creative_id,omitempty"`
	ImageMaterial         *QianchuanAdUpdateV10RequestCreativeListInnerImageMaterial         `json:"image_material,omitempty"`
	ImageMode             QianchuanAdUpdateV10CreativeListImageMode                          `json:"image_mode"`
	PromotionCardMaterial *QianchuanAdUpdateV10RequestCreativeListInnerPromotionCardMaterial `json:"promotion_card_material,omitempty"`
	TitleMaterial         *QianchuanAdUpdateV10RequestCreativeListInnerTitleMaterial         `json:"title_material,omitempty"`
	VideoMaterial         *QianchuanAdUpdateV10RequestCreativeListInnerVideoMaterial         `json:"video_material,omitempty"`
}
