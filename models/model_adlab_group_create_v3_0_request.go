/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30Request struct for AdlabGroupCreateV30Request
type AdlabGroupCreateV30Request struct {
	AdInfo AdlabGroupCreateV30RequestAdInfo `json:"ad_info"`
	// 广告主id
	AdvertiserId int64                                  `json:"advertiser_id"`
	CampaignInfo AdlabGroupCreateV30RequestCampaignInfo `json:"campaign_info"`
	CreativeInfo AdlabGroupCreateV30RequestCreativeInfo `json:"creative_info"`
}
