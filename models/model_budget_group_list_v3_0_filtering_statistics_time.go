/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupListV30FilteringStatisticsTime 统计时间范围
type BudgetGroupListV30FilteringStatisticsTime struct {
	// 结束时间，格式：yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"end_time,omitempty"`
	// 开始时间，格式：yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"start_time,omitempty"`
}
