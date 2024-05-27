/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestRelatedProduct
type ProjectCreateV30RequestRelatedProduct struct {
	//
	ProductId *string `json:"product_id,omitempty"`
	//
	ProductPlatformId *int64                                        `json:"product_platform_id,omitempty"`
	ProductSetting    *ProjectCreateV30RelatedProductProductSetting `json:"product_setting,omitempty"`
	//
	Products []*ProjectCreateV30RequestRelatedProductProductsInner `json:"products,omitempty"`
	//
	UniqueProductId *int64 `json:"unique_product_id,omitempty"`
}
