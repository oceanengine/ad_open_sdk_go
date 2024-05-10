/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestDemandInfoAdSyncConfOceanEngine 产出物推送巨量引擎配置信息
type StarDemandCreateChallengeV2RequestDemandInfoAdSyncConfOceanEngine struct {
	// 期望投放时长（单位：天） 大于0的整数，默认30天
	SyncDuration *int64 `json:"sync_duration,omitempty"`
}
