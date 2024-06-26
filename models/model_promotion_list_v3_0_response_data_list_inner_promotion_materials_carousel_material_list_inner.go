/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsCarouselMaterialListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsCarouselMaterialListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsCarouselMaterialListInner struct {
	//
	AudioId      *string                                                                     `json:"audio_id,omitempty"`
	CarouselType *PromotionListV30DataListPromotionMaterialsCarouselMaterialListCarouselType `json:"carousel_type,omitempty"`
	//
	ImageId []string `json:"image_id,omitempty"`
	// 图片主题
	ImageSubjects []*PromotionListV30ResponseDataListInnerPromotionMaterialsCarouselMaterialListInnerImageSubjectsInner `json:"image_subjects,omitempty"`
	//
	IsBlueFlowRecommendMaterial *bool `json:"is_blue_flow_recommend_material,omitempty"`
	//
	MaterialId        *int64                                                                           `json:"material_id,omitempty"`
	MaterialStatus    *PromotionListV30DataListPromotionMaterialsCarouselMaterialListMaterialStatus    `json:"material_status,omitempty"`
	VideoHpVisibility *PromotionListV30DataListPromotionMaterialsCarouselMaterialListVideoHpVisibility `json:"video_hp_visibility,omitempty"`
	//
	VisibleEndDate *string `json:"visible_end_date,omitempty"`
}
