/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerCampaignInfo 广告组维度信息
type AdlabGroupListV30ResponseDataGroupsInnerCampaignInfo struct {
	// 广告组（管家项目）日预算
	Budget             *float64                                                   `json:"budget,omitempty"`
	BudgetMode         *AdlabGroupListV30DataGroupsCampaignInfoBudgetMode         `json:"budget_mode,omitempty"`
	DeliveryRelatedNum *AdlabGroupListV30DataGroupsCampaignInfoDeliveryRelatedNum `json:"delivery_related_num,omitempty"`
	LandingType        *AdlabGroupListV30DataGroupsCampaignInfoLandingType        `json:"landing_type,omitempty"`
	MarketingPurpose   *AdlabGroupListV30DataGroupsCampaignInfoMarketingPurpose   `json:"marketing_purpose,omitempty"`
	// 广告组(管家项目)名称
	Name *string `json:"name,omitempty"`
}
