/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignModifyV30Request struct for BrandCampaignModifyV30Request
type BrandCampaignModifyV30Request struct {
	// 广告主
	AdvertiserId int64 `json:"advertiser_id"`
	// 组ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
}
