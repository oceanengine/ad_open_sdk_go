/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskSharingListV30ResponseData
type StardeliveryTaskSharingListV30ResponseData struct {
	PageInfo *StardeliveryTaskSharingListV30ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	SharingAdvertiserInfo []*StardeliveryTaskSharingListV30ResponseDataSharingAdvertiserInfoInner `json:"sharing_advertiser_info,omitempty"`
}
