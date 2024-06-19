/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerScheduleInfo 投放时间信息
type BrandAdGetV30ResponseDataAdsInnerScheduleInfo struct {
	// 本次直播用到，投放日期和投放时段，{\"20200102\": [0800, 0900]}
	DateWithPeriod *map[string][]string `json:"date_with_period,omitempty"`
	// 投放结束时间 \"2020-03-01\"
	EndTime *string `json:"end_time,omitempty"`
	// 投放开始时间 \"2020-03-01\"
	StartTime *string `json:"start_time,omitempty"`
}
