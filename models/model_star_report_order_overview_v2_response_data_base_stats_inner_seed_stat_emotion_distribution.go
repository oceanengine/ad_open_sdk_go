/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatEmotionDistribution 舆情占比
type StarReportOrderOverviewV2ResponseDataBaseStatsInnerSeedStatEmotionDistribution struct {
	// 评论舆情负向
	NegCnt *int64 `json:"neg_cnt,omitempty"`
	// 负向表现
	NegRate *float64 `json:"neg_rate,omitempty"`
	// 评论舆情中立
	NeuCnt *int64 `json:"neu_cnt,omitempty"`
	// 中立表现
	NeuRate *float64 `json:"neu_rate,omitempty"`
	// 评论舆情正向
	PosCnt *int64 `json:"pos_cnt,omitempty"`
	// 正向表现
	PosRate *float64 `json:"pos_rate,omitempty"`
}
