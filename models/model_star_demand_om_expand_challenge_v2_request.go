/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmExpandChallengeV2Request struct for StarDemandOmExpandChallengeV2Request
type StarDemandOmExpandChallengeV2Request struct {
	//
	ChallengeEndTime *int64 `json:"challenge_end_time,omitempty"`
	// 任务ID 19位数字
	ChallengeTaskId int64 `json:"challenge_task_id"`
	//
	DeveloperId              *int64                                                        `json:"developer_id,omitempty"`
	OmParticipateAuthorRange *StarDemandOmExpandChallengeV2RequestOmParticipateAuthorRange `json:"om_participate_author_range,omitempty"`
	// 客户星图ID
	StarId int64 `json:"star_id"`
}
