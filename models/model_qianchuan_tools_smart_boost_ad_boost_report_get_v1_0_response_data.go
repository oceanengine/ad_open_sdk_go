/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanToolsSmartBoostAdBoostReportGetV10ResponseData
type QianchuanToolsSmartBoostAdBoostReportGetV10ResponseData struct {
	// 起量版本号
	AdRaiseVersion *int64                                                           `json:"ad_raise_version,omitempty"`
	PageInfo       *QianchuanToolsSmartBoostAdBoostReportGetV10ResponseDataPageInfo `json:"page_info,omitempty"`
	// json返回值
	RaiseResults []*QianchuanToolsSmartBoostAdBoostReportGetV10ResponseDataRaiseResultsInner `json:"raise_results,omitempty"`
}
