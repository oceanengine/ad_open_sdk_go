/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionCreateV30RequestCustomerMaterialListInner struct for LocalPromotionCreateV30RequestCustomerMaterialListInner
type LocalPromotionCreateV30RequestCustomerMaterialListInner struct {
	ImageMode     LocalPromotionCreateV30CustomerMaterialListImageMode                  `json:"image_mode"`
	TitleMaterial *LocalPromotionCreateV30RequestCustomerMaterialListInnerTitleMaterial `json:"title_material,omitempty"`
	VideoMaterial *LocalPromotionCreateV30RequestCustomerMaterialListInnerVideoMaterial `json:"video_material,omitempty"`
}