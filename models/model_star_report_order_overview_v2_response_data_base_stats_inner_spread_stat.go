/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStat 传播价值指标
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStat struct {
	// 评论播放时长
	AvgPlayTime *float64 `json:"avg_play_time,omitempty"`
	// 互动成本
	Cpe *float64 `json:"cpe,omitempty"`
	// 千次播放成本
	Cpm *float64 `json:"cpm,omitempty"`
	// 完播率
	FinishRate *float64 `json:"finish_rate,omitempty"`
	// 触达频次
	FrequencyStats []*StarReportOrderOverviewV2ResponseDataBaseStatsInnerSpreadStatFrequencyStatsInner `json:"frequency_stats,omitempty"`
	// 互动率
	InteractRate *float64 `json:"interact_rate,omitempty"`
	// 评论量
	ItemCommentCnt *int64 `json:"item_comment_cnt,omitempty"`
	// 点赞量
	ItemLikeCnt *int64 `json:"item_like_cnt,omitempty"`
	// 播放量
	ItemPlayCnt *int64 `json:"item_play_cnt,omitempty"`
	// 分享量
	ItemShareCnt *int64 `json:"item_share_cnt,omitempty"`
}
