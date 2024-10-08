/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorUpdateV30RequestAnchorInfoShoppingCartAnchor 购物车锚点，当anchor_type=SHOPPING_CART：购物车锚点时必填
type NativeAnchorUpdateV30RequestAnchorInfoShoppingCartAnchor struct {
	// 落地页链接
	ExternalUrl *string                                                    `json:"external_url,omitempty"`
	LinkType    *NativeAnchorUpdateV30AnchorInfoShoppingCartAnchorLinkType `json:"link_type,omitempty"`
	// 直达链接
	OpenUrl *string `json:"open_url,omitempty"`
	// 商品图片，比例1:1，至少1张
	ProductImages []*NativeAnchorUpdateV30RequestAnchorInfoShoppingCartAnchorProductImagesInner `json:"product_images,omitempty"`
	// 商品价格，最多两位小数
	ProductPrice *float64 `json:"product_price,omitempty"`
	// 商品来源，只支持填写淘宝/天猫/京东/拼多多/唯品会/得物
	ProductSource *string `json:"product_source,omitempty"`
	// 商品标题，不超过35个字
	ProductTitle *string `json:"product_title,omitempty"`
	// 购物车小标题，不超过10个字
	Title *string `json:"title,omitempty"`
}
