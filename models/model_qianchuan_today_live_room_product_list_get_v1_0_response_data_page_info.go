/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomProductListGetV10ResponseDataPageInfo
type QianchuanTodayLiveRoomProductListGetV10ResponseDataPageInfo struct {
	//
	Page int32 `json:"page"`
	//
	PageSize int32 `json:"page_size"`
	//
	TotalNum *int32 `json:"total_num,omitempty"`
	//
	TotalPage *int32 `json:"total_page,omitempty"`
}
