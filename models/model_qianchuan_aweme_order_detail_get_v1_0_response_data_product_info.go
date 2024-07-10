/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataProductInfo 商品信息
type QianchuanAwemeOrderDetailGetV10ResponseDataProductInfo struct {
	// 折扣价区间最大值，单位为元
	DiscountHigherPrice *float64 `json:"discount_higher_price,omitempty"`
	// 折扣价区间最小值，单位为元
	DiscountLowerPrice *float64 `json:"discount_lower_price,omitempty"`
	// 商品id
	Id *int64 `json:"id,omitempty"`
	// 商品主图
	Img *string `json:"img,omitempty"`
	// 商品名称
	Name *string `json:"name,omitempty"`
}
