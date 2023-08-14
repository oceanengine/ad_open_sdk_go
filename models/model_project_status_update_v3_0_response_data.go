/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectStatusUpdateV30ResponseData
type ProjectStatusUpdateV30ResponseData struct {
	// 失败原因
	Errors []*ProjectStatusUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新失败的项目ID列表
	ProjectIds []int64 `json:"project_ids,omitempty"`
}
