/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectUpdateV30ResponseData
type ToolsPrivativeWordProjectUpdateV30ResponseData struct {
	//
	ErrorList []*ToolsPrivativeWordProjectUpdateV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 更新失败的项目id列表
	ProjectErrorList []int64 `json:"project_error_list,omitempty"`
	// 更新成功的项目列表
	ProjectList []map[string]interface{} `json:"project_list,omitempty"`
}
