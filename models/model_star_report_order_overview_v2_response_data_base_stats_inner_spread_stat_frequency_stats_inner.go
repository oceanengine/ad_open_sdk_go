/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStatFrequencyStatsInner struct for StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStatFrequencyStatsInner
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStatFrequencyStatsInner struct {
	// 触达频次 eq_1 -> eq_6 代表触达1 - 6次 ge_7 代表触达7次及以上
	Frequency *string `json:"frequency,omitempty"`
	// 周期
	Period *int64 `json:"period,omitempty"`
	// 触达人数占总人数的比例
	Proportion *float64 `json:"proportion,omitempty"`
	// 触达人数
	Uv *int64 `json:"uv,omitempty"`
}