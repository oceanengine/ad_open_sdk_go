/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportTodayLiveRoomDataGetV10ResponseData
type QianchuanReportTodayLiveRoomDataGetV10ResponseData struct {
	PageInfo *QianchuanReportTodayLiveRoomDataGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Rows []*QianchuanReportTodayLiveRoomDataGetV10ResponseDataRowsInner `json:"rows,omitempty"`
}