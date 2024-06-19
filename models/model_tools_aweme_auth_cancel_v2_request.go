/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthCancelV2Request struct for ToolsAwemeAuthCancelV2Request
type ToolsAwemeAuthCancelV2Request struct {
	//
	AdvertiserId int64                          `json:"advertiser_id"`
	AuthType     ToolsAwemeAuthCancelV2AuthType `json:"auth_type"`
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音短视频ID
	ItemId *int64 `json:"item_id,omitempty"`
}
