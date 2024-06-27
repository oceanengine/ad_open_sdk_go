/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerConvertStat
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerConvertStat struct {
	// 组件点击量
	Click *int64 `json:"click,omitempty"`
	// 转化量
	ConvertCnt *int64 `json:"convert_cnt,omitempty"`
	// 转化目标
	ConvertType *int64 `json:"convert_type,omitempty"`
	// 点击率
	Ctr *float64 `json:"ctr,omitempty"`
	// 转化率
	Cvr *float64 `json:"cvr,omitempty"`
	// 展现转化率
	Pvr *float64 `json:"pvr,omitempty"`
	// 组件曝光
	Show *int64 `json:"show,omitempty"`
}
