/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignListGetV10FilterMarketingGoal
type QianchuanCampaignListGetV10FilterMarketingGoal string

// List of qianchuan_campaign_list_get_v1.0_filter_marketing_goal
const (
	LIVE_PROM_GOODS_QianchuanCampaignListGetV10FilterMarketingGoal  QianchuanCampaignListGetV10FilterMarketingGoal = "LIVE_PROM_GOODS"
	VIDEO_PROM_GOODS_QianchuanCampaignListGetV10FilterMarketingGoal QianchuanCampaignListGetV10FilterMarketingGoal = "VIDEO_PROM_GOODS"
)

// All allowed values of QianchuanCampaignListGetV10FilterMarketingGoal enum
var AllowedQianchuanCampaignListGetV10FilterMarketingGoalEnumValues = []QianchuanCampaignListGetV10FilterMarketingGoal{
	"LIVE_PROM_GOODS",
	"VIDEO_PROM_GOODS",
}

// NewQianchuanCampaignListGetV10FilterMarketingGoalFromValue returns a pointer to a valid QianchuanCampaignListGetV10FilterMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10FilterMarketingGoalFromValue(v string) (*QianchuanCampaignListGetV10FilterMarketingGoal, error) {
	ev := QianchuanCampaignListGetV10FilterMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10FilterMarketingGoal: valid values are %v", v, AllowedQianchuanCampaignListGetV10FilterMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10FilterMarketingGoal) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10FilterMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_filter_marketing_goal value
func (v QianchuanCampaignListGetV10FilterMarketingGoal) Ptr() *QianchuanCampaignListGetV10FilterMarketingGoal {
	return &v
}
