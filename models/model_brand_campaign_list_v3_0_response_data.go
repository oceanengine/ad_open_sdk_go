/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignListV30ResponseData
type BrandCampaignListV30ResponseData struct {
	// 广告组列表
	Campaigns []*BrandCampaignListV30ResponseDataCampaignsInner `json:"campaigns,omitempty"`
	PageInfo  *BrandCampaignListV30ResponseDataPageInfo         `json:"page_info,omitempty"`
}
