/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeListV2ResponseDataChallengeTasksInner struct for StarChallengeListV2ResponseDataChallengeTasksInner
type StarChallengeListV2ResponseDataChallengeTasksInner struct {
	//
	AuthorTaskName *string `json:"author_task_name,omitempty"`
	//
	ChallengeAuditStatus *int64 `json:"challenge_audit_status,omitempty"`
	//
	ChallengeTaskId *int64 `json:"challenge_task_id,omitempty"`
	//
	ChallengeTaskStatus *int64 `json:"challenge_task_status,omitempty"`
	//
	CreateTime *int64 `json:"create_time,omitempty"`
	//
	DemandName *string `json:"demand_name,omitempty"`
	//
	EndTime *int64 `json:"end_time,omitempty"`
	//
	StartTime *int64 `json:"start_time,omitempty"`
}
