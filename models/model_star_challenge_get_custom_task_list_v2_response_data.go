/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeGetCustomTaskListV2ResponseData
type StarChallengeGetCustomTaskListV2ResponseData struct {
	PageInfo *StarChallengeGetCustomTaskListV2ResponseDataPageInfo `json:"page_info,omitempty"`
	// 任务列表
	TaskList []*StarChallengeGetCustomTaskListV2ResponseDataTaskListInner `json:"task_list,omitempty"`
}