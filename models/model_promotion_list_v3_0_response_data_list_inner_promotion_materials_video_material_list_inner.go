/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsVideoMaterialListInner struct {
	ImageMode *PromotionListV30DataListPromotionMaterialsVideoMaterialListImageMode `json:"image_mode,omitempty"`
	//
	ItemId *int64 `json:"item_id,omitempty"`
	//
	MaterialId     *int64                                                                     `json:"material_id,omitempty"`
	MaterialStatus *PromotionListV30DataListPromotionMaterialsVideoMaterialListMaterialStatus `json:"material_status,omitempty"`
	//
	VideoCoverId *string `json:"video_cover_id,omitempty"`
	//
	VideoId *string `json:"video_id,omitempty"`
	//
	VideoTaskIds      []string                                                                      `json:"video_task_ids,omitempty"`
	VideoTemplateType *PromotionListV30DataListPromotionMaterialsVideoMaterialListVideoTemplateType `json:"video_template_type,omitempty"`
}
