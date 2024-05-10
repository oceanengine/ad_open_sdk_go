/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueContactLogOverviewQueryV2ResponseData
type ToolsClueContactLogOverviewQueryV2ResponseData struct {
	// 接通话单平均通话时长。单位秒。时间粒度为话单创建时间。
	ConnectedAverageDuration *int64 `json:"connected_average_duration,omitempty"`
	// 接通话单量。时间粒度为话单创建时间。
	ConnectedContactLogCnt *int64 `json:"connected_contact_log_cnt,omitempty"`
	// 话单接通率。时间粒度为话单创建时间。
	ConnectedContactLogRate *float64 `json:"connected_contact_log_rate,omitempty"`
	// 总话单量。时间粒度为话单创建时间。
	ContactLogCnt *int64 `json:"contact_log_cnt,omitempty"`
	// 深度沟通（通话时长大于3min）线索量。时间粒度为话单创建时间。
	DeepTalkClueCnt *int64 `json:"deep_talk_clue_cnt,omitempty"`
	// 深度沟通（通话时长大于3min）话单量。时间粒度为话单创建时间。
	DeepTalkContactLogCnt *int64 `json:"deep_talk_contact_log_cnt,omitempty"`
	// 通话总时长。单位秒。时间粒度为话单创建时间。
	TotalDuration *int64 `json:"total_duration,omitempty"`
	// 未接通话单平均拨打时长。单位秒。时间粒度为话单创建时间。
	UnconnectedAverageTryDuration *int64 `json:"unconnected_average_try_duration,omitempty"`
	// 未接通话单量。时间粒度为话单创建时间。
	UnconnectedContactLogCnt *int64 `json:"unconnected_contact_log_cnt,omitempty"`
}
