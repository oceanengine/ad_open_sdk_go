/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceCopyV2Request struct for DmpCustomAudienceCopyV2Request
type DmpCustomAudienceCopyV2Request struct {
	// 人群包所属广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 人群包ID
	CustomAudienceId int64 `json:"custom_audience_id"`
	// 推送广告主ID（云图虚拟广告主ID，即virtual_adv_id）
	ToAdvertiserId int64 `json:"to_advertiser_id"`
}
