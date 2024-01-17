/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignUpdateV2Request struct for CampaignUpdateV2Request
type CampaignUpdateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     *float64                    `json:"budget,omitempty"`
	BudgetMode *CampaignUpdateV2BudgetMode `json:"budget_mode,omitempty"`
	//
	CampaignId int64 `json:"campaign_id"`
	//
	CampaignName *string `json:"campaign_name,omitempty"`
	//
	ModifyTime *string `json:"modify_time,omitempty"`
}
