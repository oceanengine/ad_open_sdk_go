/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueInfoUpdateV2ResponseData
type ToolsClueInfoUpdateV2ResponseData struct {
	// 更新失败的线索id列表
	Errors []*ToolsClueInfoUpdateV2ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的线索id列表
	Successes []int64 `json:"successes,omitempty"`
}
