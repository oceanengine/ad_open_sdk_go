/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30CampaignFilterCampaignStatus
type BrandOrderListV30CampaignFilterCampaignStatus string

// List of brand_order_list_v3.0_campaign_filter_campaign_status
const (
	DISABLE_BrandOrderListV30CampaignFilterCampaignStatus           BrandOrderListV30CampaignFilterCampaignStatus = "DISABLE"
	DONE_BrandOrderListV30CampaignFilterCampaignStatus              BrandOrderListV30CampaignFilterCampaignStatus = "DONE"
	ENABLE_BrandOrderListV30CampaignFilterCampaignStatus            BrandOrderListV30CampaignFilterCampaignStatus = "ENABLE"
	IS_DEL_BrandOrderListV30CampaignFilterCampaignStatus            BrandOrderListV30CampaignFilterCampaignStatus = "IS_DEL"
	RESUBMITTING_BrandOrderListV30CampaignFilterCampaignStatus      BrandOrderListV30CampaignFilterCampaignStatus = "RESUBMITTING"
	RESUBMIT_FAILED_BrandOrderListV30CampaignFilterCampaignStatus   BrandOrderListV30CampaignFilterCampaignStatus = "RESUBMIT_FAILED"
	SUBMITTING_BrandOrderListV30CampaignFilterCampaignStatus        BrandOrderListV30CampaignFilterCampaignStatus = "SUBMITTING"
	SUBMIT_FAILED_BrandOrderListV30CampaignFilterCampaignStatus     BrandOrderListV30CampaignFilterCampaignStatus = "SUBMIT_FAILED"
	WAIT_RESUBMITTING_BrandOrderListV30CampaignFilterCampaignStatus BrandOrderListV30CampaignFilterCampaignStatus = "WAIT_RESUBMITTING"
)

// Ptr returns reference to brand_order_list_v3.0_campaign_filter_campaign_status value
func (v BrandOrderListV30CampaignFilterCampaignStatus) Ptr() *BrandOrderListV30CampaignFilterCampaignStatus {
	return &v
}