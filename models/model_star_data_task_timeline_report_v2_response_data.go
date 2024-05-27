/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDataTaskTimelineReportV2ResponseData
type StarDataTaskTimelineReportV2ResponseData struct {
	// 项目数据
	ProjectData *map[string]StarDataTaskTimelineReportV2ResponseDataProjectDataValue `json:"project_data,omitempty"`
	// 任务数据
	TaskData *map[string]StarDataTaskTimelineReportV2ResponseDataTaskDataValue `json:"task_data,omitempty"`
}
