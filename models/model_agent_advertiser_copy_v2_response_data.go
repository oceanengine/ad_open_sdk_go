/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserCopyV2ResponseData
type AgentAdvertiserCopyV2ResponseData struct {
	CopyStatus *AgentAdvertiserCopyV2DataCopyStatus `json:"copy_status,omitempty"`
	// 复制结果，包括单个账户的失败原因
	Item []*AgentAdvertiserCopyV2ResponseDataItemInner `json:"item,omitempty"`
}
