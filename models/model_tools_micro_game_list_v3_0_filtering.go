/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameListV30Filtering
type ToolsMicroGameListV30Filtering struct {
	AuditStatus *ToolsMicroGameListV30FilteringAuditStatus `json:"audit_status,omitempty"`
	CreateTime  *ToolsMicroGameListV30FilteringCreateTime  `json:"create_time,omitempty"`
	//
	SearchKey  *string                                   `json:"search_key,omitempty"`
	SearchType *ToolsMicroGameListV30FilteringSearchType `json:"search_type,omitempty"`
}
