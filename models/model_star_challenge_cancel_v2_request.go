/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeCancelV2Request struct for StarChallengeCancelV2Request
type StarChallengeCancelV2Request struct {
	// 投稿任务ID
	ChallengeTaskId int64 `json:"challenge_task_id"`
	// 客户星图ID
	StarId int64 `json:"star_id"`
}
