/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomProductGetV2ResponseData
type ReportLiveRoomProductGetV2ResponseData struct {
	//
	List     []*ReportLiveRoomProductGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportLiveRoomProductGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
