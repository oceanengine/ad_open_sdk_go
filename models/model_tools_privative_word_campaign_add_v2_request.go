/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordCampaignAddV2Request struct for ToolsPrivativeWordCampaignAddV2Request
type ToolsPrivativeWordCampaignAddV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CampaignList []*ToolsPrivativeWordCampaignAddV2RequestCampaignListInner `json:"campaign_list"`
}
