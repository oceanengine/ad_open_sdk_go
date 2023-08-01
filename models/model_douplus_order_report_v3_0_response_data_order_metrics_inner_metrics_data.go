/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderReportV30ResponseDataOrderMetricsInnerMetricsData
type DouplusOrderReportV30ResponseDataOrderMetricsInnerMetricsData struct {
	//
	CustomConvertCost *int64 `json:"custom_convert_cost,omitempty"`
	//
	CustomLike *int64 `json:"custom_like,omitempty"`
	//
	DouplusLiveFollowCount *int64 `json:"douplus_live_follow_count,omitempty"`
	//
	DpTargetConvertCnt *int64 `json:"dp_target_convert_cnt,omitempty"`
	//
	DyComment *int64 `json:"dy_comment,omitempty"`
	//
	DyFollow *int64 `json:"dy_follow,omitempty"`
	//
	DyHomeVisited *int64 `json:"dy_home_visited,omitempty"`
	//
	DyShare *int64 `json:"dy_share,omitempty"`
	//
	LiveClickSourceCnt *int64 `json:"live_click_source_cnt,omitempty"`
	//
	LiveCommentCnt *int64 `json:"live_comment_cnt,omitempty"`
	//
	LiveGiftAmount *int64 `json:"live_gift_amount,omitempty"`
	//
	LiveGiftCnt *int64 `json:"live_gift_cnt,omitempty"`
	//
	LiveGiftUv *int64 `json:"live_gift_uv,omitempty"`
	//
	PlayDuration5sRank *float32 `json:"play_duration_5s_rank,omitempty"`
	//
	ShowCnt *int64 `json:"show_cnt,omitempty"`
	//
	StatCost *int64 `json:"stat_cost,omitempty"`
	//
	TotalPlay *int64 `json:"total_play,omitempty"`
}
