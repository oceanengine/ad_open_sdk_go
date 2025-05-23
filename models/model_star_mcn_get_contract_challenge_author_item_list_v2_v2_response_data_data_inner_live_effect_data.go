/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerLiveEffectData 结算状态
type StarMcnGetContractChallengeAuthorItemListV2V2ResponseDataDataInnerLiveEffectData struct {
	// 涨粉人数(只看点击关注的人数，不关心取关)
	FollowUv *int64 `json:"follow_uv,omitempty"`
	// 分享率(0.XXXX)
	ShareRate *float64 `json:"share_rate,omitempty"`
	// 直播观看次数
	WatchCnt *int64 `json:"watch_cnt,omitempty"`
	// #直播平均观看时长-s (总观看时长/观看人数)
	WatchDurationAvgS *float64 `json:"watch_duration_avg_s,omitempty"`
	// 直播观看人数
	WatchUv *int64 `json:"watch_uv,omitempty"`
}
