/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestPromotionMaterialsProductInfo
type PromotionUpdateV30RequestPromotionMaterialsProductInfo struct {
	//
	ImageIds []string `json:"image_ids,omitempty"`
	//
	ProductImageFields []string                                                         `json:"product_image_fields,omitempty"`
	ProductImageType   *PromotionUpdateV30PromotionMaterialsProductInfoProductImageType `json:"product_image_type,omitempty"`
	//
	ProductNameFields []string                                                        `json:"product_name_fields,omitempty"`
	ProductNameType   *PromotionUpdateV30PromotionMaterialsProductInfoProductNameType `json:"product_name_type,omitempty"`
	//
	ProductSellingPointFields []string                                                                `json:"product_selling_point_fields,omitempty"`
	ProductSellingPointType   *PromotionUpdateV30PromotionMaterialsProductInfoProductSellingPointType `json:"product_selling_point_type,omitempty"`
	//
	SellingPoints []string `json:"selling_points,omitempty"`
	//
	Titles []string `json:"titles,omitempty"`
}
