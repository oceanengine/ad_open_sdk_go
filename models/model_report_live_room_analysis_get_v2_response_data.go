/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomAnalysisGetV2ResponseData
type ReportLiveRoomAnalysisGetV2ResponseData struct {
	//
	List     []*ReportLiveRoomAnalysisGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *ReportLiveRoomAnalysisGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
