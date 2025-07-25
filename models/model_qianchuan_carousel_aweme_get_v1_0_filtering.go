/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselAwemeGetV10Filtering
type QianchuanCarouselAwemeGetV10Filtering struct {
	// 挂车商品指定的筛选主页图文id
	AwemeCarouselId *int64 `json:"aweme_carousel_id,omitempty"`
	// 分享链接
	AwemeItemUrl *string `json:"aweme_item_url,omitempty"`
	// 图文id，抖音主页视频用来投放才会有，限制0-50。同时注意material_ids和aweme_carousel_id智能选择一个进行过滤，否则可能差不到数据。
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id"`
}
