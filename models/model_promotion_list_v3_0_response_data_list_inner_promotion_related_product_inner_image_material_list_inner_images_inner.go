/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionRelatedProductInnerImageMaterialListInnerImagesInner struct for PromotionListV30ResponseDataListInnerPromotionRelatedProductInnerImageMaterialListInnerImagesInner
type PromotionListV30ResponseDataListInnerPromotionRelatedProductInnerImageMaterialListInnerImagesInner struct {
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	MaterialId        *int64                                                                                   `json:"material_id,omitempty"`
	MaterialOptStatus *PromotionListV30DataListPromotionRelatedProductImageMaterialListImagesMaterialOptStatus `json:"material_opt_status,omitempty"`
	MaterialStatus    *PromotionListV30DataListPromotionRelatedProductImageMaterialListImagesMaterialStatus    `json:"material_status,omitempty"`
}
