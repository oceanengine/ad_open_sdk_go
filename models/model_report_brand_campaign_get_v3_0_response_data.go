/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandCampaignGetV30ResponseData
type ReportBrandCampaignGetV30ResponseData struct {
	// 报表列表
	DataReports []*ReportBrandCampaignGetV30ResponseDataDataReportsInner `json:"data_reports,omitempty"`
	// 结果总数
	TotalCount *int64 `json:"total_count,omitempty"`
}
