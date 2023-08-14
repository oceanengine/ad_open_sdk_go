/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordProjectAddV30ResponseData
type ToolsPrivativeWordProjectAddV30ResponseData struct {
	//
	ErrorList []*ToolsPrivativeWordProjectAddV30ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 添加失败的项目id列表
	ProjectErrorList []int64 `json:"project_error_list,omitempty"`
	// 添加成功的项目列表
	ProjectList []map[string]interface{} `json:"project_list,omitempty"`
}
