/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserUnassignV2ResponseData
type AgentAdvertiserUnassignV2ResponseData struct {
	// 分配失败账户数
	FailNum *int64 `json:"fail_num,omitempty"`
	// 分配结果
	Results []*AgentAdvertiserUnassignV2ResponseDataResultsInner `json:"results,omitempty"`
	// 分配成功账户数
	SuccessNum *int64 `json:"success_num,omitempty"`
}
