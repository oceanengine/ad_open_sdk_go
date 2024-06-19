/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdGetV10ResponseDataListInner struct for QianchuanAdGetV10ResponseDataListInner
type QianchuanAdGetV10ResponseDataListInner struct {
	//
	AdCreateTime *string `json:"ad_create_time,omitempty"`
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdModifyTime *string `json:"ad_modify_time,omitempty"`
	//
	AwemeInfo []*QianchuanAdGetV10ResponseDataListInnerAwemeInfoInner `json:"aweme_info,omitempty"`
	//
	CampaignId      *int64                                                 `json:"campaign_id,omitempty"`
	CampaignScene   *QianchuanAdGetV10DataListCampaignScene                `json:"campaign_scene,omitempty"`
	DeliverySetting *QianchuanAdGetV10ResponseDataListInnerDeliverySetting `json:"delivery_setting,omitempty"`
	// 抢首屏ROI目标降低幅度
	GrabFirstScreenBidAdjustRate *int64 `json:"grab_first_screen_bid_adjust_rate,omitempty"`
	// 抢首屏开关
	GrabFirstScreenSwitch *bool                                    `json:"grab_first_screen_switch,omitempty"`
	LabAdType             *QianchuanAdGetV10DataListLabAdType      `json:"lab_ad_type,omitempty"`
	MarketingGoal         *QianchuanAdGetV10DataListMarketingGoal  `json:"marketing_goal,omitempty"`
	MarketingScene        *QianchuanAdGetV10DataListMarketingScene `json:"marketing_scene,omitempty"`
	//
	Name      *string                             `json:"name,omitempty"`
	OptStatus *QianchuanAdGetV10DataListOptStatus `json:"opt_status,omitempty"`
	//
	ProductInfo []*QianchuanAdGetV10ResponseDataListInnerProductInfoInner `json:"product_info,omitempty"`
	Status      *QianchuanAdGetV10DataListStatus                          `json:"status,omitempty"`
}
