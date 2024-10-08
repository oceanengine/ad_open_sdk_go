/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10Filtering
type QianchuanAdGetV10Filtering struct {
	//
	AdCreateEndDate *string `json:"ad_create_end_date,omitempty"`
	//
	AdCreateStartDate *string `json:"ad_create_start_date,omitempty"`
	//
	AdModifyTime *string `json:"ad_modify_time,omitempty"`
	//
	AdName           *string                                     `json:"ad_name,omitempty"`
	AutoManageFilter *QianchuanAdGetV10FilteringAutoManageFilter `json:"auto_manage_filter,omitempty"`
	//
	AwemeId *int64 `json:"aweme_id,omitempty"`
	//
	CampaignId *int64 `json:"campaign_id,omitempty"`
	//
	CampaignScene []*QianchuanAdGetV10FilteringCampaignScene `json:"campaign_scene,omitempty"`
	//
	Ids            []int64                                   `json:"ids,omitempty"`
	MarketingGoal  QianchuanAdGetV10FilteringMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanAdGetV10FilteringMarketingScene `json:"marketing_scene,omitempty"`
	Status         *QianchuanAdGetV10FilteringStatus         `json:"status,omitempty"`
}
