/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeV2ResponseDataTaskInfo
type StarDemandOmGetChallengeV2ResponseDataTaskInfo struct {
	ChallengeInfo *StarDemandOmGetChallengeV2ResponseDataTaskInfoChallengeInfo `json:"challenge_info,omitempty"`
	// 任务ID
	ChallengeTaskId *int64                                                    `json:"challenge_task_id,omitempty"`
	DemandInfo      *StarDemandOmGetChallengeV2ResponseDataTaskInfoDemandInfo `json:"demand_info,omitempty"`
}
