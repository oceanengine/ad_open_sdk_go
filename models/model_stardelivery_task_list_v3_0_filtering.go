/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskListV30Filtering
type StardeliveryTaskListV30Filtering struct {
	// 任务id过滤，精确搜索，最长50
	StarTaskIds []int64 `json:"star_task_ids,omitempty"`
	// 任务名称过滤，长度1-50个字符（两个英文字符占1个字，该字段采取模糊查询的方式）
	StarTaskName   *string                                         `json:"star_task_name,omitempty"`
	StarTaskSource *StardeliveryTaskListV30FilteringStarTaskSource `json:"star_task_source,omitempty"`
	StarTaskStatus *StardeliveryTaskListV30FilteringStarTaskStatus `json:"star_task_status,omitempty"`
	// 任务更新时间，结束时间，格式YYYY-MM-DD hh:mm:ss
	TaskModifyEndTime *string `json:"task_modify_end_time,omitempty"`
	// 任务更新时间，起始时间，格式YYYY-MM-DD hh:mm:ss
	TaskModifyStartTime *string `json:"task_modify_start_time,omitempty"`
}
