/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanTodayLiveRoomProductListGetV10ResponseData
type QianchuanTodayLiveRoomProductListGetV10ResponseData struct {
	//
	List     []*QianchuanTodayLiveRoomProductListGetV10ResponseDataListInner `json:"list,omitempty"`
	PageInfo *QianchuanTodayLiveRoomProductListGetV10ResponseDataPageInfo    `json:"page_info,omitempty"`
}
