/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignListGetV10ResponseDataListInner struct for QianchuanCampaignListGetV10ResponseDataListInner
type QianchuanCampaignListGetV10ResponseDataListInner struct {
	//
	Budget     *float64                                      `json:"budget,omitempty"`
	BudgetMode QianchuanCampaignListGetV10DataListBudgetMode `json:"budget_mode"`
	//
	CreateDate *string `json:"create_date,omitempty"`
	//
	Id             int64                                              `json:"id"`
	MarketingGoal  QianchuanCampaignListGetV10DataListMarketingGoal   `json:"marketing_goal"`
	MarketingScene *QianchuanCampaignListGetV10DataListMarketingScene `json:"marketing_scene,omitempty"`
	//
	Name   string                                    `json:"name"`
	Status QianchuanCampaignListGetV10DataListStatus `json:"status"`
}
