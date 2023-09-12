/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoShoppingCartAnchor 购物车锚点
type NativeAnchorCreateV30RequestAnchorInfoShoppingCartAnchor struct {
	// 必填，落地页链接
	ExternalUrl *string `json:"external_url,omitempty"`
	// 直达链接，选填
	OpenUrl *string `json:"open_url,omitempty"`
	// 必填，商品图片，比例1:1，至少1张
	ProductImages []*NativeAnchorCreateV30RequestAnchorInfoGameAnchorIconImagesInner `json:"product_images,omitempty"`
	// 必填，商品价格，最多两位小数
	ProductPrice *float64 `json:"product_price,omitempty"`
	// 必填，商品来源，只支持填写淘宝/天猫/京东/拼多多/唯品会/得物
	ProductSource *string `json:"product_source,omitempty"`
	// 必填，商品标题，不超过35个字
	ProductTitle *string `json:"product_title,omitempty"`
	// 必填，购物车小标题，不超过10个字
	Title *string `json:"title,omitempty"`
}
