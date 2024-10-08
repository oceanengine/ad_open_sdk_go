/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsUnionFlowPackagePromotionReportV30Filter
type ToolsUnionFlowPackagePromotionReportV30Filter struct {
	// 结束时间，格式为\"yyyy-mm-dd\" 默认昨天，即不指定起止时间获取最近7天数据
	EndTime *string `json:"end_time,omitempty"`
	// 消耗金额上限，单位元
	HighCost    *string                                                   `json:"high_cost,omitempty"`
	LandingType *ToolsUnionFlowPackagePromotionReportV30FilterLandingType `json:"landing_type,omitempty"`
	// 消耗金额下限，单位元
	LowCost *string `json:"low_cost,omitempty"`
	// 广告 id 列表，最大数量限制：1000
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
	// 广告位列表，最大数量限制：1000
	Rits []int64 `json:"rits,omitempty"`
	// 开始时间，格式为\"yyyy-mm-dd\" 限制范围100天内 默认7天前（不包括当天），即不指定起止时间获取最近7天数据
	StartTime *string `json:"start_time,omitempty"`
}
