/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInner struct {
	ImageMode *PromotionListV30DataListPromotionMaterialsImageMaterialListImageMode `json:"image_mode,omitempty"`
	//
	Images []*PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInnerImagesInner `json:"images,omitempty"`
}
