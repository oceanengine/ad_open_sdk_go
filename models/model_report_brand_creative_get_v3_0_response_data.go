/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandCreativeGetV30ResponseData
type ReportBrandCreativeGetV30ResponseData struct {
	//
	DataReports []*ReportBrandCreativeGetV30ResponseDataDataReportsInner `json:"data_reports,omitempty"`
	// 总结果数
	TotalCount *int64 `json:"total_count,omitempty"`
}
