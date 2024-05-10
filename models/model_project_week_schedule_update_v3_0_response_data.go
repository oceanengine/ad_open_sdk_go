/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectWeekScheduleUpdateV30ResponseData
type ProjectWeekScheduleUpdateV30ResponseData struct {
	// 更新失败的项目ID列表
	Errors []*ProjectWeekScheduleUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的项目ID集合
	ProjectIds []int64 `json:"project_ids,omitempty"`
}
