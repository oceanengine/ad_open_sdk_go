/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInnerImagesInner struct for PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInnerImagesInner
type PromotionListV30ResponseDataListInnerPromotionMaterialsImageMaterialListInnerImagesInner struct {
	//
	ImageId *string `json:"image_id,omitempty"`
	//
	MaterialId        *int64                                                                              `json:"material_id,omitempty"`
	MaterialOptStatus *PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialOptStatus `json:"material_opt_status,omitempty"`
	MaterialStatus    *PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus    `json:"material_status,omitempty"`
	//
	TemplateId *int64 `json:"template_id,omitempty"`
}
