/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPoiGetV30ResponseData
type LocalPoiGetV30ResponseData struct {
	PageInfo *LocalPoiGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	// 门店列表
	PoiList []*LocalPoiGetV30ResponseDataPoiListInner `json:"poi_list,omitempty"`
}
