/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignGetV30ResponseData
type BrandCampaignGetV30ResponseData struct {
	// 广告组列表
	Campaigns []*BrandCampaignGetV30ResponseDataCampaignsInner `json:"campaigns,omitempty"`
	// 广告组总数
	TotalCount *int64 `json:"total_count,omitempty"`
}
