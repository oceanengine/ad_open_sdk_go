/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCreativeGetV10ResponseDataListInner struct for QianchuanCreativeGetV10ResponseDataListInner
type QianchuanCreativeGetV10ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	// 图文素材
	CarouselMaterialList []*QianchuanCreativeGetV10ResponseDataListInnerCarouselMaterialListInner `json:"carousel_material_list,omitempty"`
	//
	CreativeCreateTime *string `json:"creative_create_time,omitempty"`
	//
	CreativeId           *int64                                               `json:"creative_id,omitempty"`
	CreativeMaterialMode *QianchuanCreativeGetV10DataListCreativeMaterialMode `json:"creative_material_mode,omitempty"`
	//
	CreativeModifyTime *string                                                    `json:"creative_modify_time,omitempty"`
	ImageMaterial      *QianchuanCreativeGetV10ResponseDataListInnerImageMaterial `json:"image_material,omitempty"`
	//
	ImageMaterialList     []*QianchuanCreativeGetV10ResponseDataListInnerImageMaterialListInner `json:"image_material_list,omitempty"`
	ImageMode             *QianchuanCreativeGetV10DataListImageMode                             `json:"image_mode,omitempty"`
	LabAdType             *QianchuanCreativeGetV10DataListLabAdType                             `json:"lab_ad_type,omitempty"`
	MarketingScene        *QianchuanCreativeGetV10DataListMarketingScene                        `json:"marketing_scene,omitempty"`
	OptStatus             *QianchuanCreativeGetV10DataListOptStatus                             `json:"opt_status,omitempty"`
	PromotionCardMaterial *QianchuanCreativeGetV10ResponseDataListInnerPromotionCardMaterial    `json:"promotion_card_material,omitempty"`
	Status                *QianchuanCreativeGetV10DataListStatus                                `json:"status,omitempty"`
	TitleMaterial         *QianchuanCreativeGetV10ResponseDataListInnerTitleMaterial            `json:"title_material,omitempty"`
	//
	TitleMaterialList []*QianchuanCreativeGetV10ResponseDataListInnerTitleMaterialListInner `json:"title_material_list,omitempty"`
	VideoMaterial     *QianchuanCreativeGetV10ResponseDataListInnerVideoMaterial            `json:"video_material,omitempty"`
	//
	VideoMaterialList []*QianchuanCreativeGetV10ResponseDataListInnerVideoMaterialListInner `json:"video_material_list,omitempty"`
}
