/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeExpandRangeV2Request struct for StarChallengeExpandRangeV2Request
type StarChallengeExpandRangeV2Request struct {
	//
	ChallengeEndTime *int64 `json:"challenge_end_time,omitempty"`
	//
	ChallengeTaskId int64 `json:"challenge_task_id"`
	//
	DeveloperId            *int64                                                   `json:"developer_id,omitempty"`
	ParticipateAuthorRange *StarChallengeExpandRangeV2RequestParticipateAuthorRange `json:"participate_author_range,omitempty"`
	//
	StarId int64 `json:"star_id"`
}
