/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30ResponseDataOrdersInnerScheduleInfo 投放排期
type BrandOrderListV30ResponseDataOrdersInnerScheduleInfo struct {
	// 投放日期和投放时段 {\"20200102\": [\"08:00\", \"09:00\"]}
	DateWithPeriod *map[string][]string `json:"date_with_period,omitempty"`
	// 分天投放日期 [20200101, 20200102]
	Dates []string `json:"dates,omitempty"`
	// 投放结束时间 \"2020-03-01\"
	EndTime *string `json:"end_time,omitempty"`
	// 投放开始时间 \"2020-03-01\"
	StartTime *string `json:"start_time,omitempty"`
}
