/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportBrandCreativeGetV30ResponseDataDataReportsInner struct for ReportBrandCreativeGetV30ResponseDataDataReportsInner
type ReportBrandCreativeGetV30ResponseDataDataReportsInner struct {
	// 计划ID
	AdId *string `json:"ad_id,omitempty"`
	// 计划名称
	AdName *string `json:"ad_name,omitempty"`
	// 广告组ID
	CampaignId *string `json:"campaign_id,omitempty"`
	// 广告组名称
	CampaignName *string `json:"campaign_name,omitempty"`
	// 创意ID
	CreativeId *string                                                          `json:"creative_id,omitempty"`
	DataReport *ReportBrandCreativeGetV30ResponseDataDataReportsInnerDataReport `json:"data_report,omitempty"`
	// 日期
	Date *string `json:"date,omitempty"`
}
