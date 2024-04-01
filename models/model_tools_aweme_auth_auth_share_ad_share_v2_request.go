/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAwemeAuthAuthShareAdShareV2Request struct for ToolsAwemeAuthAuthShareAdShareV2Request
type ToolsAwemeAuthAuthShareAdShareV2Request struct {
	// 广告主账户id
	AdvertiserId int64 `json:"advertiser_id"`
	// 共享账号范围，一次最多操作20个
	AdvertiserIds []int64                                  `json:"advertiser_ids"`
	AuthType      ToolsAwemeAuthAuthShareAdShareV2AuthType `json:"auth_type"`
	// 抖音号
	AwemeId *string `json:"aweme_id,omitempty"`
	// 抖音号视频id
	ItemId *int64 `json:"item_id,omitempty"`
}
