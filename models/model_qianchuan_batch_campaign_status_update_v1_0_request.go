/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanBatchCampaignStatusUpdateV10Request struct for QianchuanBatchCampaignStatusUpdateV10Request
type QianchuanBatchCampaignStatusUpdateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CampaignIds []int64                                        `json:"campaign_ids"`
	OptStatus   QianchuanBatchCampaignStatusUpdateV10OptStatus `json:"opt_status"`
}
