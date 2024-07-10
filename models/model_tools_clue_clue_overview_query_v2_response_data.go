/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueClueOverviewQueryV2ResponseData
type ToolsClueClueOverviewQueryV2ResponseData struct {
	// 拨打线索量。时间粒度为线索创建时间。
	CalledClueCnt *int64 `json:"called_clue_cnt,omitempty"`
	// 线索拨打率。时间粒度为线索创建时间。
	CalledClueRate *float64 `json:"called_clue_rate,omitempty"`
	// 落入4h内拨打线索量。时间粒度为线索创建时间与话单创建时间。
	CalledFourHourClueCnt *int64 `json:"called_four_hour_clue_cnt,omitempty"`
	// 4h线索拨打率。时间粒度为线索创建时间与话单创建时间。
	CalledFourHourClueRate *float64 `json:"called_four_hour_clue_rate,omitempty"`
	// 当天落入并拨打线索量。时间粒度为线索创建时间与话单创建时间。
	CalledTodayClueCnt *int64 `json:"called_today_clue_cnt,omitempty"`
	// 线索平均跟进时效。时间粒度为线索创建时间。
	ClueAverageCallInterval *float64 `json:"clue_average_call_interval,omitempty"`
	// 落入线索量。时间粒度为线索创建时间。
	ClueCnt *int64 `json:"clue_cnt,omitempty"`
	// 接通线索量。时间粒度为线索创建时间。
	ConnectedClueCnt *int64 `json:"connected_clue_cnt,omitempty"`
	// 线索接通率。时间粒度为线索创建时间。
	ConnectedClueRate *float64 `json:"connected_clue_rate,omitempty"`
	// 落入4h内拨打且接通线索量。时间粒度为线索创建时间与话单创建时间。
	ConnectedFourHourClueCnt *int64 `json:"connected_four_hour_clue_cnt,omitempty"`
	// 4h线索接通率。时间粒度为线索创建时间与话单创建时间。
	ConnectedFourHourClueRate *float64 `json:"connected_four_hour_clue_rate,omitempty"`
	// 当天落入并拨打且接通线索量。时间粒度为线索创建时间与话单创建时间。
	ConnectedTodayClueCnt *int64 `json:"connected_today_clue_cnt,omitempty"`
	// 未接通线索平均拨打轮次。时间粒度为线索创建时间。
	UnconnectedClueAverageCallTimes *float64 `json:"unconnected_clue_average_call_times,omitempty"`
	// 拨打一次且未接通线索量。时间粒度为线索创建时间。
	UnconnectedOneCallClueCnt *int64 `json:"unconnected_one_call_clue_cnt,omitempty"`
}
