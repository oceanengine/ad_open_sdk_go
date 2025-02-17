/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsSiteUpdateStatusV2ResponseData
type ToolsSiteUpdateStatusV2ResponseData struct {
	// 更新失败的site_id的list及失败的原因。如果全部更新成功，len(list) = 0
	Fail []*ToolsSiteUpdateStatusV2ResponseDataFailInner `json:"fail,omitempty"`
	// 更新成功的site_id 列表。如果全部更新失败，len(list) = 0
	Success []string `json:"success,omitempty"`
}
