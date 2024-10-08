/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPoiGetV30Filtering
type LocalPoiGetV30Filtering struct {
	// 根据门店所在城市id筛选
	City []int64 `json:"city,omitempty"`
	// 商品ID，当需要获取商品适用门店时填入
	ProductId *int64 `json:"product_id,omitempty"`
	// 根据门店所在省份id筛选
	Province []int64 `json:"province,omitempty"`
	// 根据门店名称/门店ID筛选，门店名称支持模糊搜索，门店ID支持精确搜索。
	SearchKeyWord *string `json:"search_key_word,omitempty"`
}
