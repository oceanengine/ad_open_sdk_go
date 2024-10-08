/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataCreativeListInner struct for QianchuanAdDetailGetV10ResponseDataCreativeListInner
type QianchuanAdDetailGetV10ResponseDataCreativeListInner struct {
	CarouselMaterial *QianchuanAdDetailGetV10ResponseDataCreativeListInnerCarouselMaterial `json:"carousel_material,omitempty"`
	//
	CreativeCreateTime *string `json:"creative_create_time,omitempty"`
	//
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	CreativeModifyTime    *string                                                                    `json:"creative_modify_time,omitempty"`
	ImageMaterial         *QianchuanAdDetailGetV10ResponseDataCreativeListInnerImageMaterial         `json:"image_material,omitempty"`
	ImageMode             *QianchuanAdDetailGetV10DataCreativeListImageMode                          `json:"image_mode,omitempty"`
	PromotionCardMaterial *QianchuanAdDetailGetV10ResponseDataCreativeListInnerPromotionCardMaterial `json:"promotion_card_material,omitempty"`
	TitleMaterial         *QianchuanAdDetailGetV10ResponseDataCreativeListInnerTitleMaterial         `json:"title_material,omitempty"`
	VideoMaterial         *QianchuanAdDetailGetV10ResponseDataCreativeListInnerVideoMaterial         `json:"video_material,omitempty"`
}
