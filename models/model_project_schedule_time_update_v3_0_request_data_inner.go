/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectScheduleTimeUpdateV30RequestDataInner struct for ProjectScheduleTimeUpdateV30RequestDataInner
type ProjectScheduleTimeUpdateV30RequestDataInner struct {
	// 结束的投放时间
	EndTime *int64 `json:"end_time,omitempty"`
	// 项目ID
	ProjectId    int64                                        `json:"project_id"`
	ScheduleType ProjectScheduleTimeUpdateV30DataScheduleType `json:"schedule_type"`
}
