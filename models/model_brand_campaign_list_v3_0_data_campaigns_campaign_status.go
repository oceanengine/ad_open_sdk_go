/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignListV30DataCampaignsCampaignStatus
type BrandCampaignListV30DataCampaignsCampaignStatus string

// List of brand_campaign_list_v3.0_data_campaigns_campaign_status
const (
	DISABLE_BrandCampaignListV30DataCampaignsCampaignStatus           BrandCampaignListV30DataCampaignsCampaignStatus = "DISABLE"
	DONE_BrandCampaignListV30DataCampaignsCampaignStatus              BrandCampaignListV30DataCampaignsCampaignStatus = "DONE"
	ENABLE_BrandCampaignListV30DataCampaignsCampaignStatus            BrandCampaignListV30DataCampaignsCampaignStatus = "ENABLE"
	IS_DEL_BrandCampaignListV30DataCampaignsCampaignStatus            BrandCampaignListV30DataCampaignsCampaignStatus = "IS_DEL"
	RESUBMITTING_BrandCampaignListV30DataCampaignsCampaignStatus      BrandCampaignListV30DataCampaignsCampaignStatus = "RESUBMITTING"
	RESUBMIT_FAILED_BrandCampaignListV30DataCampaignsCampaignStatus   BrandCampaignListV30DataCampaignsCampaignStatus = "RESUBMIT_FAILED"
	SUBMITTING_BrandCampaignListV30DataCampaignsCampaignStatus        BrandCampaignListV30DataCampaignsCampaignStatus = "SUBMITTING"
	SUBMIT_FAILED_BrandCampaignListV30DataCampaignsCampaignStatus     BrandCampaignListV30DataCampaignsCampaignStatus = "SUBMIT_FAILED"
	WAIT_RESUBMITTING_BrandCampaignListV30DataCampaignsCampaignStatus BrandCampaignListV30DataCampaignsCampaignStatus = "WAIT_RESUBMITTING"
)

// Ptr returns reference to brand_campaign_list_v3.0_data_campaigns_campaign_status value
func (v BrandCampaignListV30DataCampaignsCampaignStatus) Ptr() *BrandCampaignListV30DataCampaignsCampaignStatus {
	return &v
}
