/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignCreateV10Request struct for QianchuanCampaignCreateV10Request
type QianchuanCampaignCreateV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     *float64                             `json:"budget,omitempty"`
	BudgetMode QianchuanCampaignCreateV10BudgetMode `json:"budget_mode"`
	//
	CampaignName   string                                   `json:"campaign_name"`
	MarketingGoal  QianchuanCampaignCreateV10MarketingGoal  `json:"marketing_goal"`
	MarketingScene QianchuanCampaignCreateV10MarketingScene `json:"marketing_scene"`
}
