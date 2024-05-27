/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameUpdateV30Request struct for ToolsMicroGameUpdateV30Request
type ToolsMicroGameUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	GameLink []*ToolsMicroGameUpdateV30RequestGameLinkInner `json:"game_link"`
	//
	InstanceId int64 `json:"instance_id"`
	//
	Remark *string `json:"remark,omitempty"`
}
