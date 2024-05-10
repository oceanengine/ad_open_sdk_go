/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselAwemeGetV10Filtering
type QianchuanCarouselAwemeGetV10Filtering struct {
	// 挂车商品指定的筛选主页图文id
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	// 分享链接
	AwemeItemUrl *string `json:"aweme_item_url,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id"`
}
