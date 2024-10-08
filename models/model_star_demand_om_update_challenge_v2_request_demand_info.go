/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmUpdateChallengeV2RequestDemandInfo
type StarDemandOmUpdateChallengeV2RequestDemandInfo struct {
	// 参考素材 最多3个
	AttatchmentsUrl []string `json:"attatchments_url,omitempty"`
	// 任务名称
	DemandName string `json:"demand_name"`
	// 任务截止时间 秒级时间戳，即调整任务周期，不能早于当前时间，最长不能超过最初创建任务之后的365天
	ExpirationTimeEnd *int64 `json:"expiration_time_end,omitempty"`
}
