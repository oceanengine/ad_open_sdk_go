/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DiagnosisTaskAgentCreateV2Request struct for DiagnosisTaskAgentCreateV2Request
type DiagnosisTaskAgentCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AgentId        int64                                            `json:"agent_id"`
	DiagnoseConfig *DiagnosisTaskAgentCreateV2RequestDiagnoseConfig `json:"diagnose_config,omitempty"`
	// 参考广告id，可复用该广告id的setting进行前测，非空情况下会忽略diagnose_config。 和ref_ad_id二选一
	RefAdId *int64 `json:"ref_ad_id,omitempty"`
	// 参考广告id，可复用该广告id的setting进行前测，非空情况下会忽略diagnose_config。 和ref_ad_id二选一
	RefPromotionId *int64 `json:"ref_promotion_id,omitempty"`
	// 视频id
	VideoIds []string `json:"video_ids,omitempty"`
}
