/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetrics 详细数据
type ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetrics struct {
	// 点击数
	ClickCnt *string `json:"click_cnt,omitempty"`
	// 综合指标获胜概率
	CompositeProb *string `json:"composite_prob,omitempty"`
	// 转化成本
	ConversionCost *string `json:"conversion_cost,omitempty"`
	// 转化成本获胜概率
	ConversionCostProb      *string                                                                                      `json:"conversion_cost_prob,omitempty"`
	ConversionCostVariation *ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetricsConversionCostVariation `json:"conversion_cost_variation,omitempty"`
	// 转化率
	ConversionRate *string `json:"conversion_rate,omitempty"`
	// 转化数
	ConvertCnt *string `json:"convert_cnt,omitempty"`
	// 平均点击单价
	CpcPlatform *string `json:"cpc_platform,omitempty"`
	// 平均千次展现费用
	CpmPlatform *string `json:"cpm_platform,omitempty"`
	// 点击率
	Ctr *string `json:"ctr,omitempty"`
	// ctr获胜概率
	CtrProb      *string                                                                           `json:"ctr_prob,omitempty"`
	CtrVariation *ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetricsCtrVariation `json:"ctr_variation,omitempty"`
	// cvr获胜概率
	CvrProb      *string                                                                           `json:"cvr_prob,omitempty"`
	CvrVariation *ToolsAbTestInfoGetV2ResponseDataTestReportsInnerDayStatsInnerMetricsCvrVariation `json:"cvr_variation,omitempty"`
	// 展示数
	ShowCnt *string `json:"show_cnt,omitempty"`
	// 消耗
	StatCost *string `json:"stat_cost,omitempty"`
}
