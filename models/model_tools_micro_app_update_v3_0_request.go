/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppUpdateV30Request struct for ToolsMicroAppUpdateV30Request
type ToolsMicroAppUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AppPage []*ToolsMicroAppUpdateV30RequestAppPageInner `json:"app_page"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	Remark *string `json:"remark,omitempty"`
}
