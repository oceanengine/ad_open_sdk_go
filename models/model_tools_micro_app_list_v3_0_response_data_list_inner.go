/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30ResponseDataListInner struct for ToolsMicroAppListV30ResponseDataListInner
type ToolsMicroAppListV30ResponseDataListInner struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AppId       *string                                  `json:"app_id,omitempty"`
	AuditStatus *ToolsMicroAppListV30DataListAuditStatus `json:"audit_status,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Reason *string `json:"reason,omitempty"`
	//
	Remark *string `json:"remark,omitempty"`
}
