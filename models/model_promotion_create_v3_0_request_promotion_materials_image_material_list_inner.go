/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterialsImageMaterialListInner struct for PromotionCreateV30RequestPromotionMaterialsImageMaterialListInner
type PromotionCreateV30RequestPromotionMaterialsImageMaterialListInner struct {
	ImageMode PromotionCreateV30PromotionMaterialsImageMaterialListImageMode `json:"image_mode"`
	//
	Images []*PromotionCreateV30RequestPromotionMaterialsImageMaterialListInnerImagesInner `json:"images"`
}
