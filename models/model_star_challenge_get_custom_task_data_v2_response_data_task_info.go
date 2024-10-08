/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeGetCustomTaskDataV2ResponseDataTaskInfo
type StarChallengeGetCustomTaskDataV2ResponseDataTaskInfo struct {
	// 任务ID
	ChallengeTaskId *int64 `json:"challenge_task_id,omitempty"`
	// 任务状态（2，3为可选稿状态）
	Status *int64 `json:"status,omitempty"`
	// 任务类型（4:短视频，57:直播）
	TaskCategory *int64 `json:"task_category,omitempty"`
}
