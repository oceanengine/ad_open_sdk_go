/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner struct for PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner
type PromotionUpdateV30RequestPromotionMaterialsVideoMaterialListInner struct {
	ImageMode PromotionUpdateV30PromotionMaterialsVideoMaterialListImageMode `json:"image_mode"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoTaskIds      []string                                                                `json:"video_task_ids,omitempty"`
	VideoTemplateType *PromotionUpdateV30PromotionMaterialsVideoMaterialListVideoTemplateType `json:"video_template_type,omitempty"`
}
