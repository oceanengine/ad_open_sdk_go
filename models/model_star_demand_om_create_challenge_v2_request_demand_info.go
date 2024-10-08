/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmCreateChallengeV2RequestDemandInfo
type StarDemandOmCreateChallengeV2RequestDemandInfo struct {
	// 参考素材 list<string> 最多3个
	AttatchmentsUrl []string `json:"attatchments_url,omitempty"`
	// 任务名称 字符串，最长17
	DemandName string `json:"demand_name"`
	// 任务开始时间 秒级时间戳，默认任务创建时间+1h
	ExpirationTime *int64 `json:"expiration_time,omitempty"`
	// 任务截止时间 秒级时间戳，最长365天
	ExpirationTimeEnd int64 `json:"expiration_time_end"`
}
