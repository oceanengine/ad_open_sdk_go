/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30ResponseData
type BrandOrderListV30ResponseData struct {
	// 预定单列表
	Orders   []*BrandOrderListV30ResponseDataOrdersInner `json:"orders,omitempty"`
	PageInfo *BrandOrderListV30ResponseDataPageInfo      `json:"page_info,omitempty"`
}
