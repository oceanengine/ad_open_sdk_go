/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignUpdateV10Request struct for QianchuanCampaignUpdateV10Request
type QianchuanCampaignUpdateV10Request struct {
	// 千川广告账户id
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget     *float64                              `json:"budget,omitempty"`
	BudgetMode *QianchuanCampaignUpdateV10BudgetMode `json:"budget_mode,omitempty"`
	//
	CampaignId int64 `json:"campaign_id"`
	// 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	CampaignName *string `json:"campaign_name,omitempty"`
}
