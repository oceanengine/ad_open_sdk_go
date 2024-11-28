/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30CampaignFilter
type BrandMaterialListV30CampaignFilter struct {
	// 广告组ID
	CampaignId *string `json:"campaign_id,omitempty"`
	// 广告组名称
	CampaignName *string `json:"campaign_name,omitempty"`
	// 广告组状态
	CampaignStatus []*BrandMaterialListV30CampaignFilterCampaignStatus `json:"campaign_status,omitempty"`
	// 广告组创建截止时间
	CreateEndTime *string `json:"create_end_time,omitempty"`
	// 广告组创建起始时间
	CreateStartTime *string `json:"create_start_time,omitempty"`
	// 广告组截止投放时间
	EndTime *string `json:"end_time,omitempty"`
	// 广告组起始投放时间
	StartTime *string `json:"start_time,omitempty"`
}
