/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatAppletListV30Filtering
type ToolsWechatAppletListV30Filtering struct {
	AuditStatus *ToolsWechatAppletListV30FilteringAuditStatus `json:"audit_status,omitempty"`
	CreateTime  *ToolsWechatAppletListV30FilteringCreateTime  `json:"create_time,omitempty"`
	//
	Name       *string                                      `json:"name,omitempty"`
	SearchType *ToolsWechatAppletListV30FilteringSearchType `json:"search_type,omitempty"`
}
